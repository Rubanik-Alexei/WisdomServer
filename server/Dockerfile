FROM golang:1.17.3 AS builder 

WORKDIR /server

COPY ./server .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest                       
WORKDIR /server
COPY --from=builder /server .  

EXPOSE 8080

CMD [ "/server/main" ] 