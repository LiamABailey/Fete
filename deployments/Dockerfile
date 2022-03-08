# dockerfile for running api
# app is built on-container so build
# happens on runtime env
FROM golang:1.16-alpine
WORKDIR ./fête
COPY api ./api
COPY cmd ./cmd

WORKDIR ./cmd
RUN go build
CMD ["./cmd"]
