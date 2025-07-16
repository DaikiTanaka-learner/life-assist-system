    # ベースイメージとして軽量なPython 3.10-slimを選択
    FROM python:3.10-slim

    # --- この部分を追加 ---
    # ffmpegをインストールするためのコマンド
    # apt-get updateでパッケージリストを更新し、ffmpegをインストールする
    RUN apt-get update && apt-get install -y ffmpeg

    # 作業ディレクトリを設定
    WORKDIR /app

    # 最初にライブラリをインストール
    COPY requirements.txt requirements.txt
    RUN pip install --no-cache-dir -r requirements.txt

    # アプリケーションのコードをコピー
    COPY . .

    # コンテナ起動時に実行するコマンド
    CMD ["python", "main.py"]