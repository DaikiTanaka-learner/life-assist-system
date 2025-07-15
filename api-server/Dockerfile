# ベースイメージとして公式のGoイメージを使用
FROM golang:1.22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係を先にコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# プロジェクトのソースコードをすべてコピー
COPY . .

# アプリケーションをビルド
# 出力先を /api-server という名前にして、cmd/api ディレクトリをビルド対象にする
RUN go build -o /api-server ./cmd/api

# 8080ポートを公開
EXPOSE 8080

# コンテナ起動時に実行するコマンド
CMD [ "/api-server" ]