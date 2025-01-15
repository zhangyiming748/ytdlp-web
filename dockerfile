FROM golang:1.23.4-alpine3.21
COPY . /app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app
RUN go vet && go mod tidy && go mod vendor && go build -o gin main.go


RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories

RUN apk add build-base py3-pip python3 ffmpeg
RUN pip config set global.index-url https://mirrors4.tuna.tsinghua.edu.cn/pypi/web/simple
RUN pip install yt-dlp --break-system-packages


ENTRYPOINT ["/app/gin"]


# docker run  --name ytdlp -v C:\Users\zen\Githea\ytdlp-web\videos:/videos -p 8192:9001 --rm ytdlp:latest
# docker build --debug -t ytdlp:latest .