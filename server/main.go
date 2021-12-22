package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/sirupsen/logrus"
)

const cpuReadDuration = time.Second

type system struct {
	log *logrus.Logger
}

func respondJSON(w http.ResponseWriter, log *logrus.Logger, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Errorf("Write: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *system) cpuHandler(w http.ResponseWriter, _ *http.Request) {
	load, err := cpu.Percent(cpuReadDuration, true)
	if err != nil {
		s.log.Errorf("Read CPU: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, s.log, load)
}

func (s *system) ramHandler(w http.ResponseWriter, _ *http.Request) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("Read RAM: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, s.log, ram)
}

const addr = "127.0.0.1:8080"

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

	systemHandler := system{log: log}

	sys := r.PathPrefix("/system").Subrouter()
	sys.HandleFunc("/cpu", systemHandler.cpuHandler)
	sys.HandleFunc("/ram", systemHandler.ramHandler)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Shutdown: %v", err)
	}
}
