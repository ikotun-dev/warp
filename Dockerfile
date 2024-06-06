
# Use a compatible base image for Go
FROM golang:1.22-alpine AS builder

WORKDIR /source

COPY go.mod go.sum ./

RUN go mod download

COPY ./pkg ./pkg

# Cross-compile for Linux on x86 architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o warp ./pkg

# Use a lightweight Alpine image
FROM alpine:3.14

COPY --from=builder /source/warp /usr/local/bin/warp

RUN mkdir -p /var/www/html /etc/warp

# Create a default config.yaml file
RUN echo "port: 8080" > /etc/warp/warp.yaml && \
  echo "fallbackDocument: index.html" >> /etc/warp/warp.yaml && \
  echo "root: index.html" >> /etc/warp/warp.yaml

# Create a default index.html file
RUN echo $'<body style="background-color: black; color: white;" ><h1>Hello !!</h1>\n\
  <p>You have successfully setup and started Warp.</p>\n\
  <p>Copy your own config file to <b>/warp.yaml</b> and your static files to the <b>/public</b> directory to serve your files.</p><body>' > /public/index.html


EXPOSE 8080

ENTRYPOINT [ "./warp" ]


