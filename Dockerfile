FROM golang:1.23.6

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD [ "./main" ]
