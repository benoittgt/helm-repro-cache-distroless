```sh
» docker build -t helm-distroless-repro .
» docker run --rm helm-distroless-repro

2025/07/05 08:52:04 Using custom cache folder: /tmp/helm-cache1000833173
2025/07/05 08:52:04 Downloading repository index...
2025/07/05 08:52:06 Index downloaded successfully to custom cache
2025/07/05 08:52:06 Attempting to pull chart...
2025/07/05 08:52:10 Pull succeeded (no bugs?)
2025/07/05 08:52:10 Cache folder contents:
2025/07/05 08:52:10 Cache folder: /tmp/helm-cache1000833173
2025/07/05 08:52:10 - nginx-19.0.0.tgz
2025/07/05 08:52:10 - temp-repo-charts.txt
2025/07/05 08:52:10 - temp-repo-index.yaml
```
