FROM golang:1.15.6-alpine3.12

WORKDIR /app/
ENV GO111MODULE=auto
RUN apk --no-cache add --virtual build-deps
RUN apk update && apk add --no-cache bash~=5.0.17-r0 git && rm /var/cache/apk/*
RUN go get -v github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/model
RUN go get -v github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin