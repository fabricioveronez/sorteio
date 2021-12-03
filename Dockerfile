FROM golang:1.16 as build
WORKDIR /app
COPY main.go ./
COPY go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
WORKDIR /app
COPY --from=build /app/main ./
ENTRYPOINT ["./main"]  
