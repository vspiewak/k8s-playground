[Trivy v0.28.1](https://aquasecurity.github.io/trivy/v0.28.1/getting-started/overview/)

### scanning public images
```
# scan image nginx
docker run aquasec/trivy:0.28.1 image nginx

# scan image nginx for CRITICAL vulnerabilites
docker run aquasec/trivy:0.28.1 image -s CRITICAL nginx:latest
```