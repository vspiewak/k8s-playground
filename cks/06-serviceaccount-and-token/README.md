    
    # create service account
    kubectl -n spotify create sa accessor

    # create token
    kubectl -n spotify create token accessor

```
# kubectl -n spotify run accessor-pod --image=nginx --dry-run=client -o yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: accessor-pod
  name: accessor-pod
  namespace: spotify
spec:
  automountServiceAccountToken: false # add
  serviceAccountName: accessor # add
  containers:
  - image: nginx
    name: accessor-pod
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

    # token is not mounted
    kubectl -n spotify exec accessor-pod -- mount | grep token

    $ kubectl auth can-i get secrets --as system:serviceaccount:spotify:accessor
    no

    $ kubectl create clusterrolebinding accessor --clusterrole edit --serviceaccount spotify:accessor
    clusterrolebinding.rbac.authorization.k8s.io/accessor created

    $ kubectl auth can-i get secrets --as system:serviceaccount:spotify:accessor
    yes