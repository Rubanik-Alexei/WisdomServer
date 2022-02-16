FROM golang:1.17.3 AS builder 

WORKDIR /client

COPY ./client .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest                       
WORKDIR /client
COPY --from=builder /client .  

CMD [ "/client/main" ] 