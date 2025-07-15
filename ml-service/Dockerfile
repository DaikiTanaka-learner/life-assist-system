# ベースイメージとして公式のPythonイメージを使用
FROM python:3.11-slim

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係ファイルを先にコピーしてインストール
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# プロジェクトのソースコードをすべてコピー
COPY . .

# 8000ポートを公開
EXPOSE 8000

# コンテナ起動時に実行するコマンド (uvicornサーバーを起動)
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]