## Purpose

In Kubernetes `Job`, containers should be terminated to become successful.
There are no straight way to terminate sidecar containers because Process can't send signal in another container(PID namespace is not supported).

The `guillotine` manages the lifetime of child process to solve this problem.

## Usage

```
GUILLOTINE_WATCHED_FILE=/tmp/job-success guillotine /cloud_sql_proxy -fuse
```

The `guillotine` launch child process by args and poll the file exists periodically.
When `/tmp/job-success` exists, `cloud_sql_proxy` is killed.
