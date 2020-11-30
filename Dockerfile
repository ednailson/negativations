FROM golang:1.15
RUN apt-get update && apt-get install -y \
  ca-certificates
WORKDIR /app
COPY ./ /app/
RUN go mod tidy
RUN go build -o ./ ./
ENTRYPOINT ["/app/serasa-challenge", "run"]