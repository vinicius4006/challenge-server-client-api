package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type client struct {
	ctx context.Context
}

func New(ctx context.Context) *client {
	return &client{
		ctx,
	}
}

func (c client) Run() {
	ht := http.Client{}
	ctxCancel, cancel := context.WithTimeout(c.ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctxCancel, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Panicln(err.Error())
	}

	res, err := ht.Do(req)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		log.Println("Cotação recebida!")
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err.Error())
	}

	err = saveFile(bytes)
	if err != nil {
		log.Panicln(err.Error())
	}

	time.Sleep(1 * time.Second)
	log.Println("Arquivo criado com sucesso")
}

func saveFile(content []byte) error {
	content = convertJsonToSave(content)
	f, err := os.Create(filepath.Join("files", filepath.Base("cotacao.txt")))
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	defer f.Close()
	return nil
}

func convertJsonToSave(data []byte) []byte {
	obj := struct {
		Bid string `json:"bid"`
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Panicln(err.Error())
	}

	bytes := []byte(fmt.Sprintf("Dólar:{%s}", obj.Bid))
	return bytes
}
