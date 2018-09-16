FROM golang:1.10-alpine

RUN apk update && \
    apk add make yarn git openssh-client curl && \
    yarn global add webpack-cli nodemon && \
    (curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh) && \
    git config --global url."https://github.com/".insteadOf "root@github.com:"