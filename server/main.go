package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func cpuHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

const addr = "127.0.0.1:8080"

const shutdownTimeout = 10 * time.Second

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, fmt.Sprintf(" %s:%d", frame.File, frame.Line)
		},
	})
	log.SetReportCaller(true)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "%d\n", rand.Int())
		if err != nil {
			log.Errorf("Root: %v", err)
		}
	})

	sys := r.PathPrefix("/system").Subrouter()
	sys.HandleFunc("/cpu", cpuHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		log.Infof("Starting server at: http://%s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Serve: %v", err)
		}
	}()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	<-exitSignal

	log.Info("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Shutdown: %v", err)
	}
}
