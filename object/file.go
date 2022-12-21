package object

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Filename string
	Position int
	Handle   *os.File
}

func NewFile(name string) *File {
	return &File{Filename: name}
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
	f.Position = 0

	return nil
}

func init() {
	objectMethods[FILE_OBJ] = map[string]ObjectMethod{
		"close": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				f := o.(*File)
				f.Handle.Close()
				f.Position = -1
				return TRUE
			},
		},
		"lines": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, oo []Object, env Environment) Object {
				file := readFile(o, oo, env)
				fileString := file.(*String)
				lines := strings.Split(fileString.Value, "\n")

				result := make([]Object, len(lines))

				for i, line := range lines {
					result[i] = NewString(line)
				}

				return NewArray(result)
			},
		},
		"content": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ, ERROR_OBJ),
				),
			},
			method: readFile,
		},
		"position": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				f := o.(*File)
				return NewInteger(int(f.Position))
			},
		},
		"read": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(STRING_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				f := o.(*File)
				bytesAmount := args[0].(*Integer).Value
				if f.Handle == nil {
					return NewError("Invalid file handle.")
				}

				buffer := make([]byte, bytesAmount)
				bytesRealRead, err := f.Handle.Read(buffer)
				f.Position += bytesRealRead

				if err != nil {
					return NewError(err)
				}

				return NewString(string(buffer))
			},
		},
		"seek": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				f := o.(*File)

				if f.Handle == nil {
					return NewError("Invalid file handle.")
				}

				seekAmount := args[0].(*Integer).Value
				seekRelative := args[1].(*Integer).Value
				newOffset, err := f.Handle.Seek(int64(seekAmount), int(seekRelative))
				f.Position = int(newOffset)

				if err != nil {
					return NewError(err)
				}

				return NewInteger(int(f.Position))
			},
		},
		"write": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ERROR_OBJ),
				),
				ArgPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				f := o.(*File)
				content := []byte(args[0].(*String).Value)

				if f.Handle == nil {
					return NewError("Invalid file handle.")
				}

				bytesWritten, err := f.Handle.Write(content)
				f.Position += bytesWritten

				if err != nil {
					return NewError(err)
				}

				return NewInteger(bytesWritten)
			},
		},
	}
}

func (f *File) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(f, method, env, args)
}

func readFile(o Object, _ []Object, _ Environment) Object {
	f := o.(*File)
	if f.Handle == nil {
		return NewError("Invalid file handle.")
	}
	if _, err := f.Handle.Seek(0, 0); err != nil {
		return NewError(err)
	}

	file, err := io.ReadAll(f.Handle)
	if err != nil {
		return NewError(err)
	}

	if _, err := f.Handle.Seek(0, 0); err != nil {
		return NewError(err)
	}
	f.Position = 0
	return NewString(string(file))
}
