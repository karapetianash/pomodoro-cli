FROM golang:1.19 AS builder
RUN mkdir /distributing
WORKDIR /distributing
COPY . pomo/
WORKDIR /distributing/pomo
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers

FROM scratch
WORKDIR /
COPY --from=builder /distributing/pomo/pomo .
CMD ["/pomo"]