FROM golang:1.16.5 as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
COPY Makefile ./
RUN make build

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/bin/slow-http ./
CMD /app/slow-http