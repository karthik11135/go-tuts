package main

import (
	"io"
	"net"
	"net/http"
	"context"
	"errors"
	"fmt"
)

func main() {
	// creating two servers in a single program
	twoServers()


	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr: ":3000",
		Handler: mux,
		BaseContext: func (l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "myKey", l.Addr().String())
			return ctx
		},
	}
	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "myKey", l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<- ctx.Done()

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the root endpoint")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is hello endpoint")
}