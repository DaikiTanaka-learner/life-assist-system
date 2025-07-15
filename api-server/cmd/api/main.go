package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// /api/hello というパスにリクエストが来たら、helloHandlerを呼び出す
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go API Server! 👋")
	})

	fmt.Println("Go API server starting on port 8080...")
	// ポート8080でサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
