FROM golang:latest

LABEL maintainer="Jérémie MONSINJON <jeremie.monsinjon@gmail.com>"

WORKDIR $GOPATH/src/github.com/jmonsinjon/go-true-french
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["go-true-french"]