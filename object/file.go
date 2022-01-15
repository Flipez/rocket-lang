package object

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Filename string
	Handle   *os.File
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return fmt.Sprintf("<file:%s>", f.Filename) }
func (f *File) Open(mode string, perm string) error {
	if f.Filename == "!STDIN!" {
		f.Handle = os.Stdin
		return nil
	}
	if f.Filename == "!STDOUT!" {
		f.Handle = os.Stdout
		return nil
	}
	if f.Filename == "!STDERR!" {
		f.Handle = os.Stderr
		return nil
	}

	md := os.O_RDONLY

	switch mode {
	case "r":
	case "w":
		md = os.O_WRONLY
	case "wa":
		md = os.O_WRONLY | os.O_APPEND
	case "rw":
		md = os.O_RDWR
	case "rwa":
		md = os.O_RDWR | os.O_APPEND
	default:
		return fmt.Errorf("invalid file mode, got `%s`", mode)
	}

	if md != os.O_RDONLY {
		md = md | os.O_CREATE
	}

	filePerm, err := strconv.ParseUint(perm, 10, 32)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(f.Filename, md, fs.FileMode(filePerm))
	if err != nil {
		return err
	}

	f.Handle = file

	return nil
}

func init() {
	objectMethods[FILE_OBJ] = map[string]ObjectMethod{
		"close": ObjectMethod{
			description: "Closes the file pointer. Returns always `true`.",
			returnPattern: [][]string{
				[]string{BOOLEAN_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				f := o.(*File)
				f.Handle.Close()
				return &Boolean{Value: true}
			},
		},
		"lines": ObjectMethod{
			description: "If successfull, returns all lines of the file as array elements, otherwise `null`. Resets the position to 0 after read.",
			returnPattern: [][]string{
				[]string{ARRAY_OBJ, ERROR_OBJ},
			},
			method: func(o Object, oo []Object) Object {
				file := readFile(o, oo)
				fileString := file.(*String)
				lines := strings.Split(fileString.Value, "\n")

				result := make([]Object, len(lines), len(lines))

				for i, line := range lines {
					result[i] = &String{Value: line}
				}

				return &Array{Elements: result}
			},
		},
		"read": ObjectMethod{
			description: "Reads content of the file and returns it. Resets the position to 0 after read.",
			returnPattern: [][]string{
				[]string{STRING_OBJ, ERROR_OBJ},
			},
			method: readFile,
		},
		"seek": ObjectMethod{
			description: "Seeks the file handle relative from the given position.",
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
				[]string{INTEGER_OBJ},
			},
			returnPattern: [][]string{
				[]string{BOOLEAN_OBJ},
			},
			method: func(o Object, args []Object) Object {
				f := o.(*File)
				seekAmount := args[0].(*Integer).Value
				seekRelative := args[1].(*Integer).Value
				f.Handle.Seek(seekAmount, int(seekRelative))
				return &Boolean{Value: true}
			},
		},
		"write": ObjectMethod{
			description: "Writes the given string to the file. Returns `true` on success, `false` on failure and `null` if pointer is invalid.",
			returnPattern: [][]string{
				[]string{BOOLEAN_OBJ, NULL_OBJ},
			},
			argPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, args []Object) Object {
				f := o.(*File)

				if f.Handle == nil {
					return &Error{Message: "Invalid file handle."}
				}

				_, err := f.Handle.Write([]byte(args[0].(*String).Value))
				if err == nil {
					return &Boolean{Value: true}
				}

				return &Boolean{Value: false}
			},
		},
	}
}

func (f *File) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(f, method, args)
}

func readFile(o Object, _ []Object) Object {
	f := o.(*File)
	if f.Handle == nil {
		return &Error{Message: "Invalid file handle."}
	}

	file, err := ioutil.ReadAll(f.Handle)
	if err != nil {
		return &Error{Message: err.Error()}
	}

	f.Handle.Seek(0, 0)
	return &String{Value: string(file)}
}
