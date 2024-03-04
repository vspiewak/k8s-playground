[k8s - Security Context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)

### user / group
    
    $ kubectl run pod --image busybox --command -- sleep infinity

    $ kubectl exec pod -- id
    uid=0(root) gid=0(root) groups=0(root),10(wheel)

    $ kubectl delete pod nginx


```
# kubectl run pod --image busybox --command --dry-run=client -o yaml -- sleep infinity
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: pod
  name: pod
spec:
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
  containers:
  - command:
    - sleep
    - infinity
    image: busybox
    name: pod
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

    $ kubectl exec pod -- id
    uid=1000 gid=3000 groups=2000,3000


### privileged

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: privileged
  name: privileged
spec:
  containers:
  - command:
    - sleep
    - infinity
    image: busybox
    name: privileged
    securityContext:
      privileged: true
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

    $ kubectl exec privileged -- sysctl kernel.hostname=attack

    $ kubectl exec privileged -- hostname
    attack

    $ kubectl exec privileged -- cat /proc/1/status | grep NoNewPrivs
    NoNewPrivs:	0