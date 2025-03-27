# ビルド用のベースイメージ
FROM golang:1.23

# 作業ディレクトリの作成
WORKDIR /go/src/app

# 依存関係のファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# Swaggerのインストール
RUN go install github.com/swaggo/swag/cmd/swag && \
  go install github.com/swaggo/gin-swagger && \
  go install github.com/swaggo/files

# airのインストール
RUN go install github.com/air-verse/air@latest

# PATHを更新
ENV PATH=$PATH:/go/bin

# swaggerの初期化
RUN swag init

# airの実行
CMD ["air", "-c", ".air.toml"]
