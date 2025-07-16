// api-server/cmd/api/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Go API server on port 8080...")

	// 元々のハンドラ
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go API Server in Docker!")
	})

	// PythonのAIエンジンを呼び出す新しいハンドラ
	http.HandleFunc("/api/ask-ai", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request for /api/ask-ai. Calling Python ML service...")

		// Docker Composeのネットワーク内では、サービス名でコンテナにアクセスできる
		// Pythonサービスはポート8000で動いている
		resp, err := http.Post("http://ml-service:8000/v1/transcribe", "application/json", nil)
		if err != nil {
			http.Error(w, "Failed to call ML service", http.StatusInternalServerError)
			log.Printf("Error calling ML service: %v", err)
			return
		}
		defer resp.Body.Close()

		// Pythonからの応答をそのままクライアントに返す
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, resp.Body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
