package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

// ml-serviceのコンテナ内アドレス
const mlServiceURL = "http://ml-service:8000/v1/transcribe"

// speechToTextHandlerは、音声ファイルを受け取り、ml-serviceに転送するハンドラ
func speechToTextHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for speech-to-text")

	// 1. クライアントからのファイルを取得
	// "audio_file"は、クライアントが送信するフォームのフィールド名
	file, header, err := r.FormFile("audio_file")
	if err != nil {
		http.Error(w, "Could not get audio file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. ml-serviceへ転送するための新しいリクエストボディを作成
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 受け取ったファイルを新しいフォームパートとして書き込む
	part, err := writer.CreateFormFile("audio_file", header.Filename)
	if err != nil {
		http.Error(w, "Failed to create form file for forwarding", http.StatusInternalServerError)
		return
	}
	if _, err = io.Copy(part, file); err != nil {
		http.Error(w, "Failed to copy file content", http.StatusInternalServerError)
		return
	}
	writer.Close()

	// 3. ml-serviceへの新しいPOSTリクエストを作成
	req, err := http.NewRequest("POST", mlServiceURL, body)
	if err != nil {
		http.Error(w, "Failed to create new request to ML service", http.StatusInternalServerError)
		return
	}
	// Content-Typeヘッダーを正しく設定（multipart/form-data; boundary=...）
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 4. リクエストを実行し、ml-serviceからの応答を取得
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to call ML service", http.StatusInternalServerError)
		log.Printf("Error calling ML service: %v", err)
		return
	}
	defer resp.Body.Close()

	// 5. ml-serviceからの応答を、そのまま元のクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	fmt.Println("Starting Go API server on port 8080...")

	// 新しいエンドポイント `/v1/speech-to-text` を登録
	http.HandleFunc("/v1/speech-to-text", speechToTextHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
