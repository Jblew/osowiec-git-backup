FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/
RUN go build -o "/bin/gitbackup" *.go
RUN go test ./...

ENV SSH_KNOWN_HOSTS="/ssh_known_hosts"
# Below default value for github
RUN echo "github.com ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==" \
  > /ssh_known_hosts

ENV REPOSITORIES_LIST_FILE="/config/repositories.lst"
ENV REPOSITORIES_DIR="/mnt/repos"
ENV SSH_PRIVATE_KEY_PATH="/ssh_key"
ENV MONITORING_ENDPOINT_LOG=""
ENV MONITORING_ENDPOINT_PING_SUCCESS=""
ENV MONITORING_ENDPOINT_PING_FAILURE=""

CMD ["/bin/gitbackup"]
