FROM golang:1.14-alpine AS build

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOOS=linux
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /gohit

# gcc/g++ are required to build SASS libraries for extended version
RUN apk update && \
    apk add --no-cache gcc g++