package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Word Wolf Backend!")
	})

	port := "8080"
	log.Printf("サーバーをポート %s で起動します・・・", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
