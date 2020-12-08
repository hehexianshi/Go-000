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

func (s *Server) StartServer() error {
	server := &http.Server{Addr: ":8081"}
	s.HttpServe = server
	http.HandleFunc("/abc", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("this is a demo"))
		log.Println("accept client")
	})

	err := server.ListenAndServe()
	if err != nil {
		close(s.Stop)
	}
	return err
}

func (s *Server) StopServer() {
	log.Println("stop server")
	err := s.HttpServe.Shutdown(s.Ctx)
	if err != nil {
		log.Println("stop server error")
	}
}

func (s *Server) StartListenSingle() error {
	signal.Notify(s.Stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println("start single")
	<-s.Stop
	log.Println("stop single")
	s.StopServer()
	return errors.New("stop single")
}

func main() {

	errGroup, ctx := errgroup.WithContext(context.Background())
	server := &Server{Ctx :ctx, Stop: make(chan os.Signal, 1)}
	errGroup.Go(server.StartServer)
	errGroup.Go(server.StartListenSingle)

	err := errGroup.Wait()
	log.Println(err)
	log.Println("error")

}