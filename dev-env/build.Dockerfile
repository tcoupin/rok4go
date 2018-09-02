FROM golang:1.10-alpine

RUN apk update && \
    apk add make yarn && \
    yarn global add webpack-cli nodemon