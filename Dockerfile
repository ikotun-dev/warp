FROM golang:latest AS builder 
LABEL maintainer="Ikotun danlogan2003@gmail.com"
LABEL description="Warp."
WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY ./pkg/ . 
RUN go build -o warp . 

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/warp .

EXPOSE 8080

CMD ["./warp"]
