FROM golang:1.21.0-alpine as golang

RUN apk add -U tzdata
RUN apk --update add ca-certificates

WORKDIR /api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

FROM scratch

COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group

COPY --from=golang /api .

CMD [ "./api" ]
LABEL \
    org.opencontainers.image.title="karriarum-ctf-backend" \
    org.opencontainers.image.source="https://github.com/onedr0p/exportarr"