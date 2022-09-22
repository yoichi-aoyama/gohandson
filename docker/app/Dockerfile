FROM golang:1.19.0 as builder
WORKDIR /workspace
COPY . /workspace

# FIXME アプリケーションビルドをdocker imageから分離させる
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go && chmod +x ./main

FROM alpine:3.15
WORKDIR /app
RUN apk --no-cache add ca-certificates
# root ユーザだとなんでもできてしまうため appuser を作成する
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
COPY --from=builder /workspace/main ./
# コンテナを立ち上げたとき、勝手にWeb サーバーを立ち上げる
CMD ["./main"]
