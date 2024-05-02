FROM golang:latest AS builder 

WORKDIR /source 

COPY go.mod go.sum ./

RUN go mod download 

COPY ./pkg ./pkg 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o warp ./pkg

FROM alpine:3.19

COPY --from=builder /source/warp .

RUN mkdir public

# create a default config.yaml file 
RUN echo "port: 8080" > /warp.yaml && \
  echo "fallbackDocument: index.html" >> /warp.yaml && \
  echo "root: index.html" >> /warp.yaml

RUN echo $'<body style="background-color: black; color: white;" ><h1>Hello !!</h1>\n\
  <p>You have successfully setup and started Warp.</p>\n\
  <p>Copy your own config file to <b>`/config.yaml`</b> and your static files to the <b>`/frontend`</b> directory o serve your files.</p><body>' > /public/index.html


EXPOSE 8080

ENTRYPOINT [ "./warp" ]





