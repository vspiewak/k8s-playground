[k8s - Security Context - Capabilities](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container)

```
kubectl run app --image=ubuntu --dry-run=client -o yaml -- sh -c "apt update && apt install iptables && iptables -L && sleep infinity"
```

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: app
  name: app
spec:
  containers:
  - args:
    - sh
    - -c
    - apt update && apt install iptables && iptables -L && sleep infinity
    image: ubuntu
    name: app
    securityContext:
      capabilities:
        add: ["NET_ADMIN"]
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```