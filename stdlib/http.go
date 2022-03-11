package stdlib

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/flipez/rocket-lang/object"
)

func httpListenFunction(env *object.Environment, args ...object.Object) object.Object {
	if args[0].Type() != object.INTEGER_OBJ {
		return nil
	}

	port := strconv.FormatInt(args[0].(*object.Integer).Value, 10)
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
			fmt.Printf("Error shutting down the net server: %v\n", err)
		}

		close(done)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error listening on port %s: %v\n", port, err)
		quit <- os.Interrupt
	}

	<-done

	return nil
}

func httpHandleFunction(env *object.Environment, args ...object.Object) object.Object {
	if args[0].Type() != object.STRING_OBJ {
		return nil
	}

	fmt.Printf("%#v", args[0])
	if args[1].Type() != object.FUNCTION_OBJ {
		return nil
	}

	path := args[0].(*object.String).Value

	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		requestBodyBuf := new(bytes.Buffer)
		requestBodyBuf.ReadFrom(request.Body)

		httpRequest := object.NewHash(map[object.HashKey]object.HashPair{
			object.NewString("method").HashKey(): object.HashPair{
				Key:   object.NewString("method"),
				Value: object.NewString(request.Method),
			},
			object.NewString("host").HashKey(): object.HashPair{
				Key:   object.NewString("host"),
				Value: object.NewString(request.Host),
			},
			object.NewString("contentLength").HashKey(): object.HashPair{
				Key:   object.NewString("contentLength"),
				Value: object.NewInteger(request.ContentLength),
			},
			object.NewString("protocol").HashKey(): object.HashPair{
				Key:   object.NewString("protocol"),
				Value: object.NewString(request.Proto),
			},
			object.NewString("protocolMajor").HashKey(): object.HashPair{
				Key:   object.NewString("protocolMajor"),
				Value: object.NewInteger(int64(request.ProtoMajor)),
			},
			object.NewString("protocolMinor").HashKey(): object.HashPair{
				Key:   object.NewString("protocolMinor"),
				Value: object.NewInteger(int64(request.ProtoMinor)),
			},
			object.NewString("body").HashKey(): object.HashPair{
				Key:   object.NewString("body"),
				Value: object.NewString(requestBodyBuf.String()),
			},
		})

		env.Set("request", httpRequest)
		callback := args[1].(*object.Function)
		writer.Write([]byte(object.Evaluator(callback.Body, env).Inspect()))
	})

	return nil
}
