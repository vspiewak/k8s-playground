    # If you do not already have a .dockercfg file, create a dockercfg secret directly
    kubectl create secret docker-registry registry-credentials --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-password=DOCKER_PASSWORD --docker-email=DOCKER_EMAIL

    # Create a new secret named my-secret from ~/.docker/config.json
    kubectl create secret docker-registry registry-credentials --from-file=.docker/config.json

```
apiVersion: v1
kind: Pod
metadata:
  name: private-reg
spec:
  containers:
  - name: private-reg-container
    image: <your-private-image>
  imagePullSecrets:
  - name: registry-credentials
```

```
# use image digest instead of tag version
kubectl -n kube-system get pod kube-apiserver-minikube -o yaml | grep -C2 imageID
  - containerID: docker://faaa00235efe4ba3bbd6f323811bf066f032000aec1e19d5afa472cd5a57bbc6
    image: registry.k8s.io/kube-apiserver:v1.28.3
    imageID: docker-pullable://registry.k8s.io/kube-apiserver@sha256:8db46adefb0f251da210504e2ce268c36a5a7c630667418ea4601f63c9057a2d
    lastState:
      terminated:
```