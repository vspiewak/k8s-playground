Allow all traffic inside sandbox or internet.    
Deny traffic from/to other namespaces.

    kubectl -n sandbox exec  shell -- curl -s nginx.sandbox # ok
    kubectl -n sandbox exec  shell -- curl -s google.com # ok
    
    kubectl -n sandbox exec  shell -- curl -s nginx.spotify # ko
    kubectl -n spotify exec  shell -- curl -s nginx.sandbox # ko


> [!TIP]
> Helpful advice for doing things better or more easily.

CIDR
* 0.0.0.0/0 => every ip
* 1.2.3.4/32 => only ip 1.2.3.4