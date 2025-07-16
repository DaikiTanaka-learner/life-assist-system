# ml-service/main.py

from fastapi import FastAPI, UploadFile, File
import whisper
import torch
import os
import uvicorn

# --- 1. 初期設定 ---
app = FastAPI()
TEMP_AUDIO_PATH = "temp_audio.wav"

print("✅ Python AI Engine with Whisper is starting...")

# --- 2. Whisperモデルのロード ---
# アプリケーション起動時に一度だけモデルをロードする
print("   Loading Whisper model...")
try:
    # PCにGPUがあれば、GPUを使うように自動で設定される
    model = whisper.load_model("base")
    print("   Whisper model loaded successfully.")
except Exception as e:
    print(f"   Error loading Whisper model: {e}")
    model = None

# --- 3. APIエンドポイントの定義 ---

# ルートパスへのアクセス（動作確認用）
@app.get("/")
def read_root():
    return {"message": "Whisper AI Engine is running."}

# 音声ファイルを受け取って文字起こしするエンドポイント
@app.post("/v1/transcribe")
async def transcribe_audio(audio_file: UploadFile = File(...)):
    if not model:
        return {"error": "Whisper model is not available."}

    print("   Receiving audio file for transcription...")
    
    # アップロードされた音声ファイルを一時的に保存
    with open(TEMP_AUDIO_PATH, "wb") as buffer:
        buffer.write(await audio_file.read())

    # Whisperで文字起こしを実行
    try:
        # fp16=torch.cuda.is_available() は、CUDAが使えるGPUがあれば半精度浮動小数点数を使って高速化する設定
        result = model.transcribe(TEMP_AUDIO_PATH, fp16=torch.cuda.is_available(), language='ja')
        transcribed_text = result["text"].strip()
        print(f"   Transcription result: 「{transcribed_text}」")
    except Exception as e:
        print(f"   Error during transcription: {e}")
        return {"error": f"Transcription failed: {e}"}
    finally:
        # 処理が終わったら一時ファイルを必ず削除する
        if os.path.exists(TEMP_AUDIO_PATH):
            os.remove(TEMP_AUDIO_PATH)

    return {"transcribed_text": transcribed_text}

# --- 4. サーバーの起動 ---
if __name__ == "__main__":
    # "0.0.0.0"はコンテナの全てのネットワークインターフェースで待ち受けるという意味
    # port=8000はコンテナ内の8000番ポートで起動するという意味
    uvicorn.run(app, host="0.0.0.0", port=8000)
