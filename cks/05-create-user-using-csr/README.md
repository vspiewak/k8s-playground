[k8s - CSR - normal user](https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/#normal-user)

    # create a private key
    openssl genrsa -out jane.key 2048
    openssl req -new -key jane.key -out jane.csr -subj "/CN=jane"

    # encode csr in base64
    cat jane.csr | base64 -w 0

```
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: jane
spec:
  request: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1ZEQ0NBVHdDQVFBd0R6RU5NQXNHQTFVRUF3d0VhbUZ1WlRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRApnZ0VQQURDQ0FRb0NnZ0VCQU1SZ2tsM1R3S3QyQ3hYU1lvR2s2Z1A3QjdDbmRRZkxBYjU3VVdVTVZXcmJTS2xjCkMreG03c0tVVjBuQSszWVFxMno2RURmNWZ3bXlOc0Z6cUFGRTc1azdaQlNOTWR5MkJwYmVScVo2Mmxta1NSZ3cKUTR2OGZ5bGxTa0xLQTAwbnhLOUkxeHQybzJPSWRhenZkalBubVc1V1l6MzUyVDVoLzhHbmxGYzJJR3k2bFVSWgpVQlhZQlZRdmxCZW5KTVFpU0lGVkxwZlZvN2ZJZkYyWGRwa2tsMmxucXdnTHE3TTkyUTZxempxZ1NMdkdJUWxiClZNRU5NNzhJQjJxdmx0am1PUUxmUzJSbzY2STZtT3lZOFhVRHY2MzBObFF4OWQ2SEFmMDV3WGlBRjJqdHJjcTQKRytSNmhDRGN4emdTWkRYMHpJQXArcExwbG1rdkozMlZOcnhvM2VVQ0F3RUFBYUFBTUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQlRoUUhkc3o0cndmL0k2cFJlTVR1YmYyU2xHdTJkME5uOW5YRUNsdnJOampnemNnMUVObXFyCk11Nm9DSksvRy9MMnp5TjllVE1HR0VIMVYveE9LYWxHZzIybVRRUVBEdTFHT0RGbjhpTk9McTBlMm40YkxRcTUKN01lM1liWFh3UkpWb2s2UkM2ZzhOY1ByRThrNkh1L09MQ0drSmk4NGZxK1Q1ZVpNV3EzSW1qZm4rcjBpS0VseAphNnpQUitnbGRSTXNPdThFaFZBVUFybS83bnVyekhDa0c5ajVwYTFFdEpqMjJ1ZXN3WGRyWnNYM1gxNEEzYlF2CldEQ3pyYWdzMW4wcDljRDVJSFBrWFgrZittWTF4Mm0xVWZyNWlwT3dQb0ZjWmo1cmJWa2MzNXN6d3RXQUMxQ2cKZllVQUtnTm5seDh2KytpbWVISTQrOHR3ekx6bFd1LzAKLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tCg==
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400  # one day
  usages:
  - client auth
EOF
```

    # jane csr is in pending state
    kubectl get csr jane

    # approve jane csr
    kubectl certificate approve jane

    # show certificate
    kubectl get csr/jane -o yaml

    # get crt
    kubectl get csr jane -o jsonpath='{.status.certificate}'| base64 -d > jane.crt

    # add credentials
    kubectl config set-credentials jane --client-key=jane.key --client-certificate=jane.crt --embed-certs=true

    # create context
    kubectl config set-context jane --cluster=kubernetes --user=jane

    # add view clusterrole to jane
    kubectl create clusterrolebinding jane-view --clusterrole=view --user=jane

    # use jane context
    kubectl config use-context jane

    # use default kubernets-admin context
    kubectl config use-context kubernetes-admin@kubernetes