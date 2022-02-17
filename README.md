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

## Use as executable

```bash
# Install (newest version is denoted by a tag)
$ go install github.com/jblew/osowiec-git-backup@v1.11.1

# Setup cron on your own
REPOSITORIES_LIST_FILE="./repositories.lst" \
REPOSITORIES_DIR="/mnt/repos" \
SSH_PRIVATE_KEY_PATH="/root/.ssh/id_rsa" \
  osowied-git-backup
```

The repositories list file is just a list of ssh git repositories:

```
# repositories.lst example:
git@github.com:Jblew/inspector-widget.git
git@github.com:Jblew/inspector-widget-osowiec.git
git@github.com:Jblew/osowiec-git-backup.git
```

## Running with docker using shell

```sh
docker run -it \
  --env REPOSITORIES_LIST_FILE="./repositories.lst" \
  --env REPOSITORIES_DIR="/mnt/repos" \
  --env SSH_PRIVATE_KEY_PATH="/root/.ssh/id_rsa" \
  ghcr.io/jblew/osowiec-git-backup:1.11.1
```

## Running with docker using dockerfile

```dockerfile
FROM ghcr.io/jblew/osowiec-git-backup:1.11.1
ENV REPOSITORIES_LIST_FILE="./repositories.lst"
ENV REPOSITORIES_DIR="/mnt/repos"
ENV SSH_PRIVATE_KEY_PATH="/root/.ssh/id_rsa"
```

## Running on kubernetes:

Full example that is working on my cluster can be found in the `k8` dir. Below brief example with cron job:

```yml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: gitbackup-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: gitbackup
spec:
  schedule: "0 3 * * *" # 3:00 every day
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: gitbackup
              image: ghcr.io/jblew/osowiec-git-backup:v1.11.1
              imagePullPolicy: IfNotPresent
              env:
                - name: "REPOSITORIES_LIST_FILE"
                  value: "/config/repositories.lst"
                - name: "REPOSITORIES_DIR"
                  value: "/mnt/repos"
                - name: "SSH_PRIVATE_KEY_PATH"
                  value: "/sshkey/github.key"
                - name: "MONITORING_ENDPOINT_PING_SUCCESS"
                  valueFrom:
                    secretKeyRef:
                      name: ping-endpoint
                      key: success-url
                - name: "MONITORING_ENDPOINT_PING_FAILURE"
                  valueFrom:
                    secretKeyRef:
                      name: ping-endpoint
                      key: failure-url
                - name: "PROMETHEUS_PUSHGATEWAY_URL"
                  value: "http://pushgateway.metrics:9091"
              volumeMounts:
                - name: gitbackup-pv
                  mountPath: /mnt/repos
                - name: gitbackup-config
                  mountPath: /config/
                - name: ssh-key
                  mountPath: /sshkey
                  readOnly: true
          restartPolicy: OnFailure
          volumes:
            - name: gitbackup-pv
              persistentVolumeClaim:
                claimName: gitbackup-pvc
            - name: gitbackup-config
              configMap:
                name: gitbackup-config
            - name: ssh-key
              secret:
                secretName: github-ssh-key
                items:
                  - key: github.key
                    path: github.key
          initContainers:
            - name: data-permission-fix
              image: busybox
              command: ["/bin/chmod", "-R", "777", "/data"]
              volumeMounts:
                - name: gitbackup-pv
                  mountPath: /data
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: gitbackup-config
data:
  repositories.lst: |-
    git@github.com:Jblew/firebase-functions-rate-limiter.git
    git@github.com:Jblew/inspector-widget.git
    git@github.com:Jblew/osowiec-git-backup.git
```

## Environment variables

- `REPOSITORIES_LIST_FILE="/config/repositories.lst"` — (**required**) path to file with list of all repositories to back up. 1 line = 1 ssh link to repository (eg. `git@github.com:Jblew/osowiec-git-backup.git`)
- `REPOSITORIES_DIR="/mnt/repos"` — (**required**) directory where headless repositories will be backed up
- `SSH_PRIVATE_KEY_PATH="/ssh_key"` — (**required**) path to the ssh private key file. Note that you do not need to have ssh installed. Osowiec-git-backup uses golang implementation of ssh.
- `MONITORING_ENDPOINT_LOG="http://logs-endpoint"` — (optional) if set, osowiec-git-backup will POST all the logs as body of POST request to the specified endpoint
- `MONITORING_ENDPOINT_PING_SUCCESS="https://hc-ping.com/xxx-xxx"` — (optional) Will make HEAD request to the specified endpoint on success
- `MONITORING_ENDPOINT_PING_FAILURE="https://hc-ping.com/xxx-xxx/fail"` — (optional) Will make HEAD request to the specified endpoint on failure
- `PROMETHEUS_PUSHGATEWAY_URL="http://pusggateway.namespace/"` — (optional) url of prometheus's pushgateway to send logs to.
- `PROMETHEUS_PUSHGATEWAY_JOBNAME="osowiec-git-backup"` — (optional) job label for pushgateway to attach to pushed metrics

---

Made with ❤️ by [Jędrzej Bogumił Lewandowski](https://jblewandowski.com/).
