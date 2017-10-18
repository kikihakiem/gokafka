FROM golang:1.9

ENV APP_HOME $GOPATH/src
WORKDIR $APP_HOME
RUN go get github.com/gin-gonic/gin
EXPOSE 3000
