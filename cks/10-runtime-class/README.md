[k8s - Runtime Class](https://kubernetes.io/docs/concepts/containers/runtime-class/)

```
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: gvisor 
handler: runsc 
```

```
$ kubectl get runtimeclass
NAME     HANDLER   AGE
gvisor   runsc     17s

$ kubectl get runtimeclass gvisor -o yaml
apiVersion: node.k8s.io/v1
handler: runsc
kind: RuntimeClass
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"node.k8s.io/v1","handler":"runsc","kind":"RuntimeClass","metadata":{"annotations":{},"name":"gvisor"}}
  creationTimestamp: "2024-03-04T10:01:21Z"
  name: gvisor
  resourceVersion: "19797"
  uid: f81282d0-9820-4650-addc-640410ca9bb4
```

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: gvisor
  name: gvisor
spec:
  runtimeClassName: gvisor
  containers:
  - image: nginx
    name: gvisor
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```
