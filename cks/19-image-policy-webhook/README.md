### ImagePolicyWebhook

Admission controller asking an external webhook if an image is allowed.

#### Enable on kube-apiserver

    # file: /etc/kubernetes/manifests/kube-apiserver.yaml
    - --enable-admission-plugins=NodeRestriction,ImagePolicyWebhook
    - --admission-control-config-file=/etc/kubernetes/admission/admission-config.yaml

```
# file: /etc/kubernetes/admission/admission-config.yaml
apiVersion: apiserver.config.k8s.io/v1
kind: AdmissionConfiguration
plugins:
- name: ImagePolicyWebhook
  configuration:
    imagePolicy:
      kubeConfigFile: /etc/kubernetes/admission/kubeconf.yaml
      allowTTL: 50
      denyTTL: 50
      retryBackoff: 500
      defaultAllow: false
```

```
# file: /etc/kubernetes/admission/kubeconf.yaml
apiVersion: v1
kind: Config
clusters:
- name: webhook
  cluster:
    certificate-authority: /etc/kubernetes/admission/ca.crt
    server: https://image-policy.webhook.svc:443/image_policy
users:
- name: apiserver
  user:
    client-certificate: /etc/kubernetes/admission/apiserver-client.crt
    client-key: /etc/kubernetes/admission/apiserver-client.key
contexts:
- name: webhook
  context:
    cluster: webhook
    user: apiserver
current-context: webhook
```

#### Mount it in the static pod

```
volumeMounts:
  - mountPath: /etc/kubernetes/admission
    name: admission
    readOnly: true
...
volumes:
- name: admission
  hostPath:
    path: /etc/kubernetes/admission
    type: DirectoryOrCreate
```

> [!CAUTION]
> `defaultAllow: false` means an *unreachable* webhook denies every pod creation.
> Fail-closed is what the exam asks for, but it locks you out of your own cluster
> if the webhook is not running — check the apiserver logs before panicking.

> [!TIP]
> The apiserver is a static pod : editing the manifest restarts it.
> If it never comes back, the mistake is almost always the mount or a yaml typo :
>
>     journalctl -u kubelet -f
>     crictl ps -a | grep apiserver
>     crictl logs <container-id>
