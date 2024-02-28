    # get kube-apiserver running version
    kubectl -n kube-system get pods -l component=kube-apiserver -o yaml | grep image
        image: registry.k8s.io/kube-apiserver:v1.29.2
        imagePullPolicy: IfNotPresent
        image: registry.k8s.io/kube-apiserver:v1.29.2
        imageID: registry.k8s.io/kube-apiserver@sha256:fe4196cd9fa06bd75b5fb437be89bbccc277e83f3e0296c30b71485ce4834461


    # download official release and check sum
    curl -O -L https://dl.k8s.io/v1.29.2/kubernetes-server-linux-amd64.tar.gz
    echo 'd5575da7f28a5284d4ffb40ca1b597213e03c381e161c1ec2bdadd7fe0532d62f41c758443ecefed70f484fb770e0bac53218f0a429587ac983469a39e56979b kubernetes-server-linux-amd64.tar.gz' | sha512sum -c
    kubernetes-server-linux-amd64.tar.gz: OK

    # extract downloaded archive
    tar xvzf kubernetes-server-linux-amd64.tar.gz

    # get official checksum of kube-apiserver
    sha512sum kubernetes/server/bin/kube-apiserver
    30c64cdbdb323852d8f3992008664cbec9ab7044b4c31ce7edff27c202cb35d895b6fae704d80376b9dd53bb40c562481bfb7f5f6d37ea023d29c40fe67a9428  kubernetes/server/bin/kube-apiserver

    # get kube-apiserver pid
    ps aux | grep kube-apiserver
    root       23803  6.2  9.3 1694440 372888 ?      Ssl  11:26   7:27 kube-apiserver --advertise-address=10.0.9.163 --allow-privileged=true --authorization-mode=Node,RBAC --client-ca-file=/etc/kubernetes/pki/ca.crt --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key --etcd-servers=https://127.0.0.1:2379 --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --requestheader-allowed-names=front-proxy-client --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group --requestheader-username-headers=X-Remote-User --secure-port=6443 --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=10.96.0.0/12 --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
    ubuntu     61957  0.0  0.0   7008  2304 pts/0    S+   13:26   0:00 grep --color=auto kube-apiserver

    # get checksum of container binary
    find /proc/23803/root/ | grep kube-api
    /proc/23803/root/usr/local/bin/kube-apiserver
    
    sha512sum /proc/23803/root/usr/local/bin/kube-apiserver
    30c64cdbdb323852d8f3992008664cbec9ab7044b4c31ce7edff27c202cb35d895b6fae704d80376b9dd53bb40c562481bfb7f5f6d37ea023d29c40fe67a9428  /proc/23803/root/usr/local/bin/kube-apiserver

    