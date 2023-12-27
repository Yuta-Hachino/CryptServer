FROM golang:latest
WORKDIR /app
RUN go mod init github.com/Yuta-Hachino/CryptServer # ご自身のモジュールで
CMD [ "go", "run", "main.go" ]
