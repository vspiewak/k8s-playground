    $ curl https://localhost:6443 -k
    {
    "kind": "Status",
    "apiVersion": "v1",
    "metadata": {},
    "status": "Failure",
    "message": "forbidden: User \"system:anonymous\" cannot get path \"/\"",
    "reason": "Forbidden",
    "details": {},
    "code": 403
    }

    kube-apiserver --anonymous-auth=false

    $ curl https://localhost:6443 -k
    {
    "kind": "Status",
    "apiVersion": "v1",
    "metadata": {},
    "status": "Failure",
    "message": "Unauthorized",
    "reason": "Unauthorized",
    "code": 401
    }


    #
    # manually curl apiserver
    #

    # ca
    kubectl config view -o json --raw \
    | jq -r '.clusters[0].cluster."certificate-authority-data"' \
    | base64 -d > ca

    # cert
    kubectl config view -o json --raw \
    | jq -r '.users[0].user."client-certificate-data"' \
    | base64 -d > cert

    # key
    kubectl config view -o json --raw \
    | jq -r '.users[0].user."client-key-data"' \
    | base64 -d > key

    # server
    kubectl config view -o json --raw | jq -r '.clusters[0].cluster.server'
    https://10.0.10.110:6443

    # curl API
    curl https://10.0.10.110:6443 # SSL certificate problem:
    curl https://10.0.10.110:6443 --cacert ca
    curl https://10.0.10.110:6443 --cacert ca --cert cert --key key


    kube-apiserver --enable-admission-plugins=NodeRestriction

    KUBECONFIG=/etc/kubernetes/kubelet.conf kubectl get nodes
    NAME             STATUS   ROLES           AGE   VERSION
    ip-10-0-0-48     Ready    <none>          13m   v1.29.0
    ip-10-0-10-110   Ready    control-plane   23m   v1.29.0
    ip-10-0-15-199   Ready    <none>          13m   v1.29.0

    KUBECONFIG=/etc/kubernetes/kubelet.conf kubectl get ns
    Error from server (Forbidden): namespaces is forbidden: User "system:node:ip-10-0-0-48" cannot list resource "namespaces" in API group "" at the cluster scope

    KUBECONFIG=/etc/kubernetes/kubelet.conf kubectl label node ip-10-0-0-48 cks/test=yes
    node/ip-10-0-0-48 labeled

    KUBECONFIG=/etc/kubernetes/kubelet.conf kubectl label node ip-10-0-0-48 node-restriction.kubernetes.io/test=yes
    Error from server (Forbidden): nodes "ip-10-0-0-48" is forbidden: is not allowed to modify labels: node-restriction.kubernetes.io/test