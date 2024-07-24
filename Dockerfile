FROM golang:bookworm

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /productpromous

EXPOSE 808

CMD [ "/productpromous" ]
