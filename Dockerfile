ARG  BUILDER_IMAGE=golang:bookworm
ARG  DISTROLESS_IMAGE=gcr.io/distroless/static

FROM ${BUILDER_IMAGE} as builder

RUN update-ca-certificates

WORKDIR /go/bin

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/run cmd/api/main.go


FROM ${DISTROLESS_IMAGE}

COPY --from=builder /go/bin/run /go/bin/run

ENTRYPOINT ["/go/bin/run"]