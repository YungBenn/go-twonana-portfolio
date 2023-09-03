FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

CMD [ "/binary" ]