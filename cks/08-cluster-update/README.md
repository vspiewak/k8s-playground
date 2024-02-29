    # cordon + drain
    kubectl drain -l node-role.kubernetes.io/control-plane= --ignore-daemonsets

    # status - Ready,SchedulingDisabled
    kubectl get nodes
    NAME             STATUS                     ROLES           AGE   VERSION
    ip-10-0-0-48     Ready                      <none>          44m   v1.29.0
    ip-10-0-10-110   Ready,SchedulingDisabled   control-plane   54m   v1.29.0
    ip-10-0-15-199   Ready                      <none>          44m   v1.29.0

    # update repos
    apt-get update

    # list all versions
    apt list kubeadm -a
    Listing... Done
    kubeadm/unknown 1.29.2-1.1 amd64 [upgradable from: 1.29.0-1.1]
    kubeadm/unknown 1.29.1-1.1 amd64
    kubeadm/unknown,now 1.29.0-1.1 amd64 [installed,upgradable to: 1.29.2-1.1]

    kubeadm/unknown 1.29.2-1.1 arm64
    kubeadm/unknown 1.29.1-1.1 arm64
    kubeadm/unknown 1.29.0-1.1 arm64

    kubeadm/unknown 1.29.2-1.1 ppc64el
    kubeadm/unknown 1.29.1-1.1 ppc64el
    kubeadm/unknown 1.29.0-1.1 ppc64el

    kubeadm/unknown 1.29.2-1.1 s390x
    kubeadm/unknown 1.29.1-1.1 s390x
    kubeadm/unknown 1.29.0-1.1 s390x

    # hold / unhold
    apt-mark hold kubelet kubectl
    apt-mark unhold kubeadm
    
    # install kubeadm
    apt install kubeadm=1.29.2-1.1

    # check version
    kubeadm version

    # plan
    kubeadm upgrade plan

    # apply
    kubeadm upgrade apply v1.29.2

    # kubelet to update
    kubelet --version
    Kubernetes v1.29.0
    
    # kubectl to update
    kubectl version
    Client Version: v1.29.0
    Kustomize Version: v5.0.4-0.20230601165947-6ce0bf390ce3
    Server Version: v1.29.2

    # hold / unhold
    apt-mark hold kubeadm
    apt-mark unhold kubelet kubectl

    # install
    apt install kubelet=1.29.2-1.1 kubectl=1.29.2-1.1

    # restart kubelet (just in case, seems unecessary)
    service kubelet restart

    # hold back everything
    apt-mark hold kubeadm kubelet kubectl

    # uncordon
    kubectl uncordon -l node-role.kubernetes.io/control-plane=

# TODO: same on workers...