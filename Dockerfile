FROM golang:1.24-alpine

# 必要なツール
RUN apk add --no-cache git make

# 作業ディレクトリ
WORKDIR /app

COPY . .

# air, swag をインストール（GOBIN を /usr/local/bin に設定）
ENV GOBIN=/usr/local/bin
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.49.0 && \
    go install github.com/swaggo/swag/cmd/swag@latest

CMD ["make", "dev"]
