FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/
RUN go build -o "/bin/gitbackup" *.go
RUN go test ./...


ENV REPOSITORIES_LIST_FILE="/config/repositories.lst"
ENV REPOSITORIES_DIR="/mnt/repos"
ENV SSH_PRIVATE_KEY_PATH="/ssh_key"
ENV MONITORING_ENDPOINT_LOG=""
ENV MONITORING_ENDPOINT_PING_SUCCESS=""
ENV MONITORING_ENDPOINT_PING_FAILURE=""

CMD ["/bin/gitbackup"]
