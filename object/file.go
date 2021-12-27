package object

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type File struct {
	Filename string
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Handle   *os.File
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return fmt.Sprintf("<file:%s>", f.Filename) }
func (f *File) Open(mode string) error {
	if f.Filename == "!STDIN!" {
		f.Reader = bufio.NewReader(os.Stdin)
		return nil
	}
	if f.Filename == "!STDOUT!" {
		f.Writer = bufio.NewWriter(os.Stdout)
		return nil
	}
	if f.Filename == "!STDERR!" {
		f.Writer = bufio.NewWriter(os.Stderr)
		return nil
	}

	md := os.O_RDONLY

	if mode == "w" {
		md = os.O_WRONLY

		os.Remove(f.Filename)
	} else {
		if strings.Contains(mode, "w") && strings.Contains(mode, "a") {
			md = os.O_WRONLY
			md |= os.O_APPEND
		}
	}

	file, err := os.OpenFile(f.Filename, os.O_CREATE|md, 0644)
	if err != nil {
		return err
	}

	f.Handle = file

	if md == os.O_RDONLY {
		f.Reader = bufio.NewReader(file)
	} else {
		f.Writer = bufio.NewWriter(file)
	}

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
			description: "If successfull, returns all lines of the file as array elements, otherwise `null`.",
			returnPattern: [][]string{
				[]string{ARRAY_OBJ, NULL_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				f := o.(*File)
				if f.Reader == nil {
					return (&Null{})
				}

				var lines []string
				for {
					line, err := f.Reader.ReadString('\n')
					if err != nil {
						break
					}
					lines = append(lines, line)
				}

				l := len(lines)
				result := make([]Object, l)
				for i, txt := range lines {
					result[i] = &String{Value: txt}
				}
				return &Array{Elements: result}
			},
		},
		"read": ObjectMethod{
			description: "Reads content of the file and returns it.",
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				f := o.(*File)
				if f.Reader == nil {
					return (&String{Value: ""})
				}

				line, err := f.Reader.ReadString('\n')
				if err != nil {
					return (&String{Value: ""})
				}
				return &String{Value: line}
			},
		},
		"rewind": ObjectMethod{
			description: "Resets the read pointer back to position `0`. Always returns `true`.",
			returnPattern: [][]string{
				[]string{BOOLEAN_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				f := o.(*File)
				f.Handle.Seek(0, 0)
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

				if f.Writer == nil {
					return (&Null{})
				}

				_, err := f.Writer.Write([]byte(args[0].(*String).Value))
				if err == nil {
					f.Writer.Flush()
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
