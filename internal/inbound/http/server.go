package http

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer() *http.Server {
	w := logrus.StandardLogger().Writer()
	h := InitHandlers()
	InitDocHandler(h)

	server := &http.Server{
		Addr:         ":80",
		Handler:      h,
		WriteTimeout: 35 * time.Second,
		ReadTimeout:  35 * time.Second,
		IdleTimeout:  time.Second * 60,
		ErrorLog:     log.New(w, "", 0),
	}
	return server
}

func InitPProfServer() *http.Server {
	w := logrus.StandardLogger().Writer()
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 35 * time.Second,
		ReadTimeout:  35 * time.Second,
		IdleTimeout:  time.Second * 60,
		ErrorLog:     log.New(w, "", 0),
	}
	return server
}

func GracefulStop(servers map[string]*http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	switch sig {
	case syscall.SIGQUIT:
		logrus.Info(fmt.Sprintf("Servers stop by SIGQUIT"))
	case syscall.SIGTERM:
		logrus.Info(fmt.Sprintf("Servers stop by SIGTERM"))
	case syscall.SIGINT:
		logrus.Info(fmt.Sprintf("Servers stop by SIGINT"))
	default:
		logrus.Error(fmt.Sprintf("Servers stop by DEFAULT"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for name, s := range servers {
		if err := s.Shutdown(ctx); err != nil {
			logrus.Error(fmt.Sprintf("Server %s listen on %v Shutdown Error: %v\n", name, s.Addr, err))
		}
		logrus.Infof("Server %s listen on %v Shutdown\n", name, s.Addr)
	}
	os.Exit(0)
}
