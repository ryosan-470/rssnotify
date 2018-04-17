# Running on Kubernetes

1. You must create a configuration file like `config.yaml` and create `ConfigMap` like this:
```console
% kubectl create configmap rssnotify-configmap --from-file config.yaml
```

1. Deploy a cronjob
```console
% kubectl create -f ./cronjob.yaml
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