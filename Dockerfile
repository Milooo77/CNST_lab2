FROM golang:1.21.13-bullseye
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
WORKDIR /app/app/checkout
RUN chmod +x build.sh && ./build.sh
CMD bash output/bootstrap.sh