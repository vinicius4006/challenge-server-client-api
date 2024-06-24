package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println("Iniciando chamada...")
	time.Sleep(1 * time.Second)
	ctx := context.Background()
	client := New(ctx)
	client.Run()
}
