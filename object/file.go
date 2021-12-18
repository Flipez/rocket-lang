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

func (f *File) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "close":
		f.Handle.Close()
		return &Boolean{Value: true}
	case "lines":
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
	case "read":
		if f.Reader == nil {
			return (&String{Value: ""})
		}

		line, err := f.Reader.ReadString('\n')
		if err != nil {
			return (&String{Value: ""})
		}
		return &String{Value: line}
	case "rewind":
		f.Handle.Seek(0, 0)
		return &Boolean{Value: true}
	case "write":
		if len(args) < 1 {
			return &Error{Message: "Missing argument to write()!"}
		}

		if f.Writer == nil {
			return (&Null{})
		}

		txt := args[0].Inspect()
		_, err := f.Writer.Write([]byte(txt))
		if err == nil {
			f.Writer.Flush()
			return &Boolean{Value: true}
		}

		return &Boolean{Value: false}
	}
	return nil
}
