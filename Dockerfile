FROM golang:1.19
RUN mkdir /app
WORKDIR /app
COPY app.go .
RUN go build app.go
ENTRYPOINT ["/app/app"]
