# Running on Kubernetes

1. You must create a configuration file like `config.yaml` and create `secret` like this:
```console
% make create-secret
```

1. Deploy a cronjob
```console
% make deploy
```

## Configuration
This tool runs every 5 mininutes. Please change the following 2 files to change the interval time.

* cronjob.yaml
```yaml
spec:
  schedule: "*/10 * * * * "
...
```

* config.yaml
```yaml
...
interval: 10
```