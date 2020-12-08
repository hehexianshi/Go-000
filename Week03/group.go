package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	HttpServe *http.Server
	Ctx context.Context
	Stop chan os.Signal
}

func (s *Server) StartServer(ctx context.Context) error {
	server := &http.Server{Addr: ":8081"}
	s.HttpServe = server
	http.HandleFunc("/abc", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("this is a demo"))
		log.Println("accept client")
	})

	go s.StopServer(ctx)
	log.Println("server start")
	return server.ListenAndServe()
}

func (s *Server) StopServer(ctx context.Context) {
	<-ctx.Done()
	log.Println("stop server")
	err := s.HttpServe.Shutdown(s.Ctx)
	if err != nil {
		log.Println("stop server error")
	}
}

func (s *Server) StartListenSingle(ctx context.Context) error {
	signal.Notify(s.Stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println("start single")
	select {
	case <-s.Stop:
		return errors.New("stop single")
	case <-ctx.Done():
		return errors.New("server is stop")
	}

}

func main() {

	errGroup, ctx := errgroup.WithContext(context.Background())
	server := &Server{Ctx :ctx, Stop: make(chan os.Signal, 1)}
	errGroup.Go(func () error {
		return server.StartServer(ctx)
	})
	errGroup.Go(func() error {
		return server.StartListenSingle(ctx)
	})

	err := errGroup.Wait()
	log.Println(err)
}