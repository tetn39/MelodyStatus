# ビルド用のベースイメージ
FROM golang:1.23

# 作業ディレクトリの作成
WORKDIR /backend/src

# 依存関係のファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# Swaggerのインストール
RUN go install github.com/swaggo/swag/cmd/swag && \
  go install github.com/swaggo/gin-swagger && \
  go install github.com/swaggo/files

# PATHを更新
ENV PATH=$PATH:/go/bin

# Swaggerドキュメントの生成
RUN swag init || (echo "swag init failed" && exit 1) && ls -la

# アプリケーションの実行
CMD ["go", "run", "main.go"]
