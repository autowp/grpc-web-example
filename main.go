package main

import (
	"context"
	"encoding/json"
	"github.com/NYTimes/gziphandler"
	"github.com/common-nighthawk/go-figure"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	grpc2 "grpc-web-example/grpc"
	"net/http"
	"time"
)

var allowedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

type RestHandler struct {
}

type GRPCServer struct {
	grpc2.UnimplementedExampleServer
}

func (s GRPCServer) Ascii(_ context.Context, request *grpc2.ExampleRequest) (*grpc2.ExampleResult, error) {
	myFigure := figure.NewFigure(request.Query, "", true)
	return &grpc2.ExampleResult{
		Result: myFigure.String(),
	}, nil
}

func (s GRPCServer) AsciiStream(request *grpc2.ExampleRequest, srv grpc2.Example_AsciiStreamServer) error {
	for i := 1; i <= len(request.Query); i++ {
		q := request.Query[:i]
		myFigure := figure.NewFigure(q, "", true)

		_ = srv.Send(&grpc2.ExampleResult{
			Result: myFigure.String(),
		})
		time.Sleep(time.Second)
	}

	return nil
}

func main() {

	h, _ := gziphandler.GzipHandlerWithOpts(
		gziphandler.MinSize(1),
	)

	go func() {
		srv := GRPCServer{}

		grpcServer := grpc.NewServer()
		grpc2.RegisterExampleServer(grpcServer, srv)

		wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(string) bool {
			return true
		}))

		httpServer := &http.Server{Addr: ":8080", Handler: h(wrappedGrpc)}

		err := httpServer.ListenAndServeTLS("./localhost.crt", "./localhost.key")
		if err != nil {
			panic(err)
		}
	}()

	restHttpServer := &http.Server{Addr: ":8081", Handler: h(RestHandler{})}

	err := restHttpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (h RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	}

	if r.URL.Path != "/ascii" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	q := r.URL.Query().Get("query")

	myFigure := figure.NewFigure(q, "", true)

	w.Header().Add("Content-Type", "application/json")
	result := map[string]string{
		"result": myFigure.String(),
	}
	body, _ := json.Marshal(result)
	_, err := w.Write(body)
	if err != nil {
		panic(err)
	}
}
