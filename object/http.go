package object

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type HTTP struct {
	mux             *http.ServeMux
	registeredPaths []string
	quitChannel     chan (os.Signal)
	raisedError     *Error
}

func (h *HTTP) Type() ObjectType { return HTTP_OBJ }
func (h *HTTP) Inspect() string  { return "HTTP" }
func (h *HTTP) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(h, method, env, args)
}

func (h *HTTP) raiseErrorAndInterrupt(error string) {
	h.raisedError = NewError(error)
	h.quitChannel <- os.Interrupt
}

func init() {
	objectMethods[HTTP_OBJ] = map[string]ObjectMethod{
		"new": ObjectMethod{
			returnPattern: [][]string{
				[]string{HTTP_OBJ},
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return &HTTP{mux: http.NewServeMux()}
			},
		},
		"listen": ObjectMethod{
			description: "Starts a blocking webserver on the given port.",
			example:     `🚀 > HTTP.listen(3000)`,
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			returnPattern: [][]string{
				[]string{NIL_OBJ, ERROR_OBJ},
			},
			method: func(o Object, args []Object, env Environment) Object {
				if o.(*HTTP).mux == nil {
					return NewError("Invalid handler. Call only supported on instance.")
				}

				var returnError *Error

				port := strconv.FormatInt(args[0].(*Integer).Value, 10)
				server := &http.Server{
					Handler: o.(*HTTP).mux,
					Addr:    ":" + port,
				}

				done := make(chan bool)
				o.(*HTTP).quitChannel = make(chan os.Signal, 1)
				signal.Notify(o.(*HTTP).quitChannel, os.Interrupt)

				go func() {
					<-o.(*HTTP).quitChannel
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()

					server.SetKeepAlivesEnabled(false)

					if err := server.Shutdown(ctx); err != nil {
						returnError = NewErrorFormat("Error shutting down the net server: %v", err)
					}

					close(done)
				}()

				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					o.(*HTTP).quitChannel <- os.Interrupt
					returnError = NewErrorFormat("listening on port %s: %v", port, err)
				}

				<-done

				if returnError != nil {
					return returnError
				}

				if o.(*HTTP).raisedError != nil {
					return o.(*HTTP).raisedError
				}

				return NIL
			},
		},
		"handle": ObjectMethod{
			description: `Adds a handle to the global HTTP server. Needs to be done before starting one via .listen().
Inside the function a variable called "request" will be populated which is a hash with information about the request.

Also a variable called "response" will be created which will be returned automatically as a response to the client.
The response can be adjusted to the needs. It is a HASH supports the following content:

- "status" needs to be an INTEGER (eg. 200, 400, 500). Default is 200.
- "body" needs to be a STRING. Default ""
- "headers" needs to be a HASH(STRING:STRING) eg. headers["Content-Type"] = "text/plain". Default is {"Content-Type": "text/plain"}`,
			example: `🚀 > HTTP.handle("/", callback_func)`,
			argPattern: [][]string{
				[]string{STRING_OBJ},
				[]string{FUNCTION_OBJ},
			},
			returnPattern: [][]string{
				[]string{NIL_OBJ, ERROR_OBJ},
			},
			method: func(o Object, args []Object, env Environment) Object {
				if o.(*HTTP).mux == nil {
					return NewError("Invalid handler. Call only supported on instance.")
				}

				path := args[0].(*String).Value

				for _, regPath := range o.(*HTTP).registeredPaths {
					if path == regPath {
						return NewErrorFormat("Already registered path `%s`", path)
					}
				}

				o.(*HTTP).mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
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

					httpResponse := NewHash(map[HashKey]HashPair{
						NewString("status").HashKey(): HashPair{
							Key:   NewString("status"),
							Value: NewInteger(200),
						},
						NewString("body").HashKey(): HashPair{
							Key:   NewString("body"),
							Value: NewString(""),
						},
						NewString("headers").HashKey(): HashPair{
							Key: NewString("headers"),
							Value: NewHash(map[HashKey]HashPair{
								NewString("Content-Type").HashKey(): HashPair{
									Key:   NewString("Content-Type"),
									Value: NewString("text/plain"),
								},
							}),
						},
					})

					env.Set("request", httpRequest)
					env.Set("response", httpResponse)
					callback := args[1].(*Function)
					Evaluator(callback.Body, &env)
					userReponse, ok := env.Get("response")
					if !ok {
						o.(*HTTP).raiseErrorAndInterrupt("unable to extract response variable.")
						return
					}

					if userReponse.Type() != HASH_OBJ {
						o.(*HTTP).raiseErrorAndInterrupt("response is not a HASH")
						return
					}

					if userHeaders, ok := userReponse.(*Hash).Get("headers"); ok {
						if userHeaders.Type() != HASH_OBJ {
							o.(*HTTP).raiseErrorAndInterrupt("response headers is not a HASH")
							return
						}

						for _, pair := range userHeaders.(*Hash).Pairs {
							writer.Header().Set(pair.Key.(*String).Value, pair.Value.(*String).Value)
						}
					}

					userBody, bodyOk := userReponse.(*Hash).Get("body")
					if bodyOk {
						if userBody.Type() != STRING_OBJ {
							o.(*HTTP).raiseErrorAndInterrupt("body is not STRING")
							return
						}

						if writer.Header().Get("Content-Length") == "" {
							writer.Header().Set("Content-Length", fmt.Sprint(len(userBody.(*String).Value)))
						}
					}

					if userStatus, ok := userReponse.(*Hash).Get("status"); ok {
						if userStatus.Type() != INTEGER_OBJ {
							o.(*HTTP).raiseErrorAndInterrupt("status is not INTEGER")
							return
						}

						writer.WriteHeader(int(userStatus.(*Integer).Value))

					}

					if bodyOk {
						writer.Write([]byte(userBody.(*String).Value))
					}
				})

				o.(*HTTP).registeredPaths = append(o.(*HTTP).registeredPaths, path)
				return NIL
			},
		},
	}
}
