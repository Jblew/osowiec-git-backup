# osowiec-git-backup

Backup for multiple git repositories. Multi paradigm. Favourite flavour: SSH+Kubernetes+Prometheus.

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



## What is osowiec?

You may wonder what is this `osowiec` prefix? Well, the name is both geographical and philosophical.

Osowiec is a small village in the north-eastern Poland. It is **the only** passage through the very wide river of Biebrza on the distance of ~100km. Biebrza is not a big river itself but it is surrounded by multi-killometer wide bogs and morasses. You can pass the river at osowiec or else, you need to go 60km north or south to find passage. Historically this area had crucial millitary importance because millitary troops could only pass at osowiec and elsewhere the river with morasses could not be passed even with the tanks. Tsar and germans made few attempst to build fortifications and roads and the area is full of massive bunkers that are now home to birds and animals.

So I named the tool osowiec, because for at the time of writing it I have been enjoying kayaking at the Biebrza river and was surprised and astounded when I landed in small village and suddenly found myself in a postapocalyptical land of blownup concrete covered with blossoming nature. Thus osowiec is a symbol of `an optimal way` or `backing up the history` :)

Below some photos from my trip to Biebrza and stopping at osowiec:

![Osowiec](doc/img/osowiec.jpg)

> These bunkers are at least 20m high and extend another 20m down the ground with massive rails for artillery guns. All that filled with birds nests, rats, bees and surrounded by species-rich ecosystem of Biebrzan National Park.



---

Made with ❤️ by [Jędrzej Bogumił Lewandowski](https://jblewandowski.com/).
