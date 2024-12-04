# バイナリ作成ステージ
FROM golang:1.23-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-s -w" -o app

# デプロイステージ
FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app", "80"]

# ホットリロードステージ
FROM golang:1.23 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air", "80"]
