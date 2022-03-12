package object

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type HTTP struct{}

func (h *HTTP) Type() ObjectType { return HTTP_OBJ }
func (h *HTTP) Inspect() string  { return "HTTP" }
func (h *HTTP) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(h, method, env, args)
}

func init() {
	objectMethods[HTTP_OBJ] = map[string]ObjectMethod{
		"listen": ObjectMethod{
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(_ Object, args []Object, env Environment) Object {
				var returnError *Error

				port := strconv.FormatInt(args[0].(*Integer).Value, 10)
				server := &http.Server{
					Addr: ":" + port,
				}

				done := make(chan bool)
				quit := make(chan os.Signal, 1)
				signal.Notify(quit, os.Interrupt)

				go func() {
					<-quit
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()

					server.SetKeepAlivesEnabled(false)

					if err := server.Shutdown(ctx); err != nil {
						returnError = NewErrorFormat("Error shutting down the net server: %v", err)
					}

					close(done)
				}()

				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					quit <- os.Interrupt
					returnError = NewErrorFormat("listening on port %s: %v", port, err)
				}

				<-done

				if returnError != nil {
					return returnError
				}

				return NULL
			},
		},
		"handle": ObjectMethod{
			argPattern: [][]string{
				[]string{STRING_OBJ},
				[]string{FUNCTION_OBJ},
			},
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(_ Object, args []Object, env Environment) Object {
				path := args[0].(*String).Value

				http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
					requestBodyBuf := new(bytes.Buffer)
					requestBodyBuf.ReadFrom(request.Body)

					httpRequest := NewHash(map[HashKey]HashPair{
						NewString("method").HashKey(): HashPair{
							Key:   NewString("method"),
							Value: NewString(request.Method),
						},
						NewString("host").HashKey(): HashPair{
							Key:   NewString("host"),
							Value: NewString(request.Host),
						},
						NewString("contentLength").HashKey(): HashPair{
							Key:   NewString("contentLength"),
							Value: NewInteger(request.ContentLength),
						},
						NewString("protocol").HashKey(): HashPair{
							Key:   NewString("protocol"),
							Value: NewString(request.Proto),
						},
						NewString("protocolMajor").HashKey(): HashPair{
							Key:   NewString("protocolMajor"),
							Value: NewInteger(int64(request.ProtoMajor)),
						},
						NewString("protocolMinor").HashKey(): HashPair{
							Key:   NewString("protocolMinor"),
							Value: NewInteger(int64(request.ProtoMinor)),
						},
						NewString("body").HashKey(): HashPair{
							Key:   NewString("body"),
							Value: NewString(requestBodyBuf.String()),
						},
					})

					env.Set("request", httpRequest)
					callback := args[1].(*Function)
					writer.Write([]byte(Evaluator(callback.Body, &env).(*ReturnValue).Value.(*String).Value))
				})

				return NULL
			},
		},
	}
}
