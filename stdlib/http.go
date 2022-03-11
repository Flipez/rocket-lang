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
			fmt.Println("Error shutting down the net server: %v\n", err)
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

		// httpRequest := object.NewMap(map[string]interface{}{
		// 	"method":        request.Method,
		// 	"host":          request.Host,
		// 	"contentLength": request.ContentLength,
		// 	"protocol":      request.Proto,
		// 	"protocolMajor": request.ProtoMajor,
		// 	"protocolMinor": request.ProtoMinor,
		// 	"body":          requestBodyBuf.String(),
		// })

		// callbackArgs := make([]object.Object, 0)
		// callbackArgs = append(callbackArgs, httpRequest)

		callback := args[1].(*object.Function)
		object.Evaluator(callback.Body, env)
	})

	return nil
}
