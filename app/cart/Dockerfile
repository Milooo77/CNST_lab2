FROM golang:1.21.13-bullseye
COPY . /app/gomall
WORKDIR /app/gomall/app/cart

RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN cd /app/gomall/app/cart
RUN bash build.sh
RUN go build -o output/bin/cart
CMD bash output/bootstrap.sh