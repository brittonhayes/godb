FROM golang

LABEL maintainer="Britton Hayes"
LABEL application="godb"
LABEL version="0.0.1"

WORKDIR /go/src/godb

COPY . .

RUN go get -u github.com/brittonhayes/godb
RUN go get -u github.com/nanobox-io/golang-scribble
RUN go get -u github.com/spf13/cobra
RUN go install github.com/brittonhayes/godb

CMD ["godb"]