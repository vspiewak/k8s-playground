[kubesec.io](https://kubesec.io/)

```
cat <<EOF > kubesec-test.yaml
apiVersion: v1
kind: Pod
metadata:
  name: kubesec-demo
spec:
  containers:
  - name: kubesec-demo
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      readOnlyRootFilesystem: true
EOF
```

```
docker run -i kubesec/kubesec:512c5e0 scan /dev/stdin < kubesec-test.yaml
```