FROM golang:1.21.13-bullseye AS builder
WORKDIR /gomall
COPY . .
# set GOPROXY to use the proxy server in China
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN chmod +x /gomall/app/order/build.sh && cd /gomall/app/order && ./build.sh

FROM debian:bullseye-slim AS runtime
WORKDIR /app
COPY --from=builder /gomall/app/order/output /app
CMD ["bash", "/app/bootstrap.sh"]