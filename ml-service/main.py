from fastapi import FastAPI

# FastAPIインスタンスを作成
app = FastAPI()

@app.get("/")
def read_root():
    return {"message": "Hello from ML Service! 🧠"}

# ダミーの予測エンドポイント
@app.get("/ml/predict")
def predict():
    # 本来はここで services/stt_service.py などを呼び出し、
    # 音声認識などのML処理を行う
    return {"prediction": "This is a dummy ML prediction."}