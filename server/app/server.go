package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server-go/models"
	"server-go/repository"
	"time"
)

type server struct {
	repo repository.Interface
}

func (s server) request(w http.ResponseWriter, r *http.Request) {
	var c http.Client
	ctxCancel, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctxCancel, "GET", fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s", "USD-BRL"), nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error to get information about %s: %s", "USD-BRL", err.Error()), http.StatusInternalServerError)
	}

	res, err := c.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error to do request: %s", err.Error()), http.StatusInternalServerError)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error to get body: %s", err.Error()), http.StatusInternalServerError)
	}

	var cr models.CurrencyRate
	cr.ConvertJSONToCurrencyRate("USD-BRL", bytes)

	err = s.repo.Save(r.Context(), cr)
	if err != nil {
		http.Error(w, fmt.Sprintf("error to save body on db: %s", err.Error()), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"bid":"%s"}`, cr.Bid)))
}

func (s server) GetAllObjs(w http.ResponseWriter, r *http.Request) {
	infos, err := s.repo.GetCotes()
	if err != nil {
		http.Error(w, "error to get objs from db: "+err.Error(), http.StatusInternalServerError)
	}
	bytes, err := json.Marshal(infos)
	if err != nil {
		http.Error(w, "error to parse objs from db: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	println()
	w.Write(bytes)
}

func (s server) Run() {
	log.Println("Starting server...")
	http.HandleFunc("/cotacao", s.request)
	http.HandleFunc("/getallcotacao", s.GetAllObjs)
	time.Sleep(1 * time.Second)
	log.Println("Server on")
	http.ListenAndServe(":8080", nil)
}

func New() *server {
	repo := repository.New()
	return &server{
		repo,
	}
}
