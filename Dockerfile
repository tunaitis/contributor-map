FROM golang:1.17 as builder

WORKDIR /src

ENV CGO_ENABLED=0

COPY . .

RUN go get -d -v ./...
RUN go build -o /bin/action cmd/contributor-map/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/data /data
COPY --from=builder /src/template /template
COPY --from=builder /bin/action /
ENTRYPOINT ["/action"]