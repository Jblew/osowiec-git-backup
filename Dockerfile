FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/
RUN CGO_ENABLED=0 go build -o "/bin/osowiec-git-backup" *.go
RUN go test ./...

FROM alpine:3.15
COPY --from=builder /bin/osowiec-git-backup /bin/osowiec-git-backup

ENV REPOSITORIES_LIST_FILE="/config/repositories.lst"
ENV REPOSITORIES_DIR="/mnt/repos"
ENV SSH_PRIVATE_KEY_PATH="/ssh_key"
ENV MONITORING_ENDPOINT_PING_SUCCESS=""
ENV MONITORING_ENDPOINT_PING_FAILURE=""
ENV PROMETHEUS_PUSHGATEWAY_URL=""
ENV PROMETHEUS_PUSHGATEWAY_JOBNAME="osowiec-git-backup"

CMD ["/bin/osowiec-git-backup"]
