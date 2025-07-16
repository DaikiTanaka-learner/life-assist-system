# --- ステージ1: ビルド用ステージ ---
# Goのビルド環境として公式のgolangイメージを使用
FROM golang:1.22-alpine as builder

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係のダウンロードをキャッシュするため、go.modファイルを先にコピー
COPY go.mod ./
RUN go mod download

# プロジェクトのソースコードを全てコピー
COPY . .

# アプリケーションをビルド。CGOを無効にし、静的バイナリを作成
RUN CGO_ENABLED=0 GOOS=linux go build -o /server-binary ./cmd/api

# --- ステージ2: 実行用ステージ ---
# 軽量なAlpine Linuxイメージをベースにする
FROM alpine:latest

# 作業ディレクトリを設定
WORKDIR /root/

# ビルド用ステージから、コンパイル済みの実行ファイルのみをコピー
COPY --from=builder /server-binary .

# コンテナがリッスンするポートを8080に指定
EXPOSE 8080

# コンテナ起動時に実行するコマンド
CMD ["./server-binary"]