### Create secrets

    kubectl -n spotify create secret generic backend-secret \
    --from-literal=username='admin' --from-literal=password='1234'

    kubectl -n spotify create secret generic db-secret \
    --from-literal=username='pguser' --from-literal=password='pg1234'

    kubectl -n spotify create secret generic ssh-secret \
    --from-file=authorized-keys=/home/ubuntu/.ssh/authorized_keys

### Secrets usage in a Pod
```
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx-with-secrets
  name: nginx-with-secrets
  namespace: spotify
spec:
  containers:
  - image: nginx
    name: nginx-with-secrets
    env:
    # specific env var
    - name: BACKEND_USERNAME
      valueFrom:
        secretKeyRef:
          name: backend-secret
          key: username
    # all key-values as env var
    envFrom:
    - secretRef:
        name: db-secret 
    # secret as volume
    volumeMounts:
    - name: secret-volume
      readOnly: true
      mountPath: "/etc/secret-volume"    
  volumes:
  - name: secret-volume
    secret:
      secretName: ssh-secret   
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

    $ kubectl -n spotify exec nginx-with-secrets -- env | grep -i user
    BACKEND_USERNAME=admin
    username=pguser

    $ kubectl -n spotify exec nginx-with-secrets -- cat /etc/secret-volume/authorized-keys
    ssh-rsa <redacted> aws-ec2


### crictl env vars
```
$ crictl ps --name nginx-with-secrets
CONTAINER           IMAGE               CREATED             STATE               NAME                 ATTEMPT             POD ID              POD
c6f0a6a3e1e54       e4720093a3c13       10 minutes ago      Running             nginx-with-secrets   0                   52374c42af2d2       nginx-with-secrets
```

    $ crictl inspect c6f0a6a3e1e54 --o json | jq -r '.info.config.envs' | grep -i -C2 username

    $ ps aux | grep nginx
    $ cat /proc/26888/root/etc/secret-volume/authorized-keys


### ETCD secrets

    $ grep etcd /etc/kubernetes/manifests/kube-apiserver.yaml
    - --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt
    - --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt
    - --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key
    - --etcd-servers=https://127.0.0.1:2379

    $ ETCDCTL_API=3 etcdctl \
    --cacert=/etc/kubernetes/pki/etcd/ca.crt \
    --cert=/etc/kubernetes/pki/apiserver-etcd-client.crt \
    --key=/etc/kubernetes/pki/apiserver-etcd-client.key \
    endpoint health
    127.0.0.1:2379 is healthy: successfully committed proposal: took = 11.928331ms

```
ETCDCTL_API=3 etcdctl \
--cacert=/etc/kubernetes/pki/etcd/ca.crt \
--cert=/etc/kubernetes/pki/apiserver-etcd-client.crt \
--key=/etc/kubernetes/pki/apiserver-etcd-client.key \
get /registry/secrets/spotify/backend-secret

/registry/secrets/spotify/backend-secret
k8s


v1Secret�
�
backend-secretspotify"*$07f21b6f-453d-462e-9760-e9a20dd3f4bc2�����u
kubectl-createUpdatev����FieldsV1:A
?{"f:data":{".":{},"f:password":{},"f:username":{}},"f:type":{}}B
password1234
usernameadminOpaque"
```

### ETCD Encryption

    # generate secret
    head -c 32 /dev/urandom | base64
    SNm9k9xJiEc1SPvHi7P5R2mkahMFjSNHhQCqZqvRO/4=

```
# /etc/kubernetes/enc/enc.yaml
apiVersion: apiserver.config.k8s.io/v1
kind: EncryptionConfiguration
resources:
  - resources:
      - secrets
      - configmaps
      - pandas.awesome.bears.example
    providers:
      - aescbc:
          keys:
            - name: key1
              # See the following text for more details about the secret value
              secret: SNm9k9xJiEc1SPvHi7P5R2mkahMFjSNHhQCqZqvRO/4=
      - identity: {} # this fallback allows reading unencrypted secrets;
                     # for example, during initial migration
```

```
    kube-apiserver --encryption-provider-config=/etc/kubernetes/enc/enc.yaml
    ...
    - mountPath: /etc/kubernetes/enc
      name: k8s-enc
      readOnly: true
    ...
    - hostPath:
        path: /etc/kubernetes/enc
        type: DirectoryOrCreate
      name: k8s-enc    
```

```
kubectl create secret generic new-secret --from-literal=dummy=1234
```

```
ETCDCTL_API=3 etcdctl \
--cacert=/etc/kubernetes/pki/etcd/ca.crt \
--cert=/etc/kubernetes/pki/apiserver-etcd-client.crt \
--key=/etc/kubernetes/pki/apiserver-etcd-client.key \
get /registry/secrets/default/new-secret
/registry/secrets/default/new-secret
k8s:enc:aescbc:v1:key1:>��M��[�Z��V��8
Ė/xƓ��\�v�Q��-��}q�J�
                     �/��6���A����j!�d����c�e�^R�
                                                 }`H)	���cf�X��8��z�uzK�l�y�DF�e
                                                                                  �)r*
�O�9�חM/�d��W2�34=ݼ�r���z�C�K���.��^&��DL��?pO�I\������Z�u
                                                           �vd]��÷Y������6ez�=]Ku�m1>�x�k
                                                                                         ���/��D&柳�&m,ц!
```

```
# re-create all secrets with encryption
kubectl get secrets -A -o json | kubectl replace -f -
```

```
ETCDCTL_API=3 etcdctl \
--cacert=/etc/kubernetes/pki/etcd/ca.crt \
--cert=/etc/kubernetes/pki/apiserver-etcd-client.crt \
--key=/etc/kubernetes/pki/apiserver-etcd-client.key \
get /registry/secrets/spotify/backend-secret
/registry/secrets/spotify/backend-secret
k8s:enc:aescbc:v1:key1:S���J�����;W	�<�6̢��^�^xN_x�5�
��Te��	i���&-x�sԹ��/����˙����=p��L�d�~���
                                           ˌlZ}J�d�4:�M�O�����x�'�U<�
                                                                      ��� B�͔)��w*��x���r�^�Iw��kV����<N����CxK��
E������%��VB|���e`��H�O���W�3���`װ�(��!��~�����l����ݛ@�9	�M�d�M�5����ve{�{V�M~��Yvy3�C�-�9>,
```

### get secret from k8s api inside a pod
```
curl -k https://kubernetes.default/api/v1/namespaces/restricted/secrets \
-H "Authorization: Bearer $(cat /run/secrets/kubernetes.io/serviceaccount/token)"
```