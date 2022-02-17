# osowiec-git-backup

Automatic CRON based backup

Features:

1. Repository list fetched from endpoint (eg. Github raw content)
2. Pulls all branches (actually fetches)
3. Clones if repository doesnt exist
4. Stores repository in bare mode
5. Authenticates with ssh key (both RSA and ED25519)
6. Does not require ssh agent installed
7. Does not require git installed (uses go-git)
8. Sends logs to configurable log endpoint
9. Sends success_ping/failure_ping to ping endpoint (eg. healthchecks.io)
10. May push metrics to prometheus pushgateway
11. Docker image available
12. Works well with kubernetes (example below)

## Executable usage

```bash
# Clone repository
$ git clone https://github.com/Jblew/osowiec-git-backup.git

go build -o

# Move the binary to your own dir
$ mv osowiec-git-backup/dist/gitbackup /var/gitbackup

# Setup cron on your own
```

## Configuration options

All configuration options are required: Configuration is stored in `project.config.sh`

```bash
export REPOSITORIES_LIST_ENDPOINT="https://raw.githubusercontent.com/You/repository-with-list-of-repos/master/repositories.lst"
export REPOSITORIES_DIR="/mnt/repos"
export SSH_PRIVATE_KEY_PATH="/root/.ssh/id_rsa"
export MONITORING_ENDPOINT_LOG="https://hc-ping.com/xxx-xxxx-xxxxx" # I use my own logging endpoint
export MONITORING_ENDPOINT_PING_SUCCESS="https://hc-ping.com/xxx-xxxx-xxxxx"
export MONITORING_ENDPOINT_PING_FAILURE="https://hc-ping.com/xxx-xxxx-xxxxx/fail"
```

---

Made with ❤️ by [Jędrzej Lewandowski](https://jedrzej.lewandowski.doctor/).
