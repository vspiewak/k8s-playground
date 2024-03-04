[Trivy v0.28.1](https://aquasecurity.github.io/trivy/v0.28.1/getting-started/overview/)

### scanning public images
```
docker run aquasec/trivy:0.28.1 image nginx:latest

docker run aquasec/trivy:0.28.1 image registry.k8s.io/kube-apiserver:v1.29.2
```