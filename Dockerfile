FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV PORT :8000

RUN CGO_ENABLED=0 GOOS=linux go build -o /pokedex-backend

CMD ["/pokedex-backend"]