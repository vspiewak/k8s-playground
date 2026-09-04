Allow all traffic inside sandbox or internet.    
Deny traffic from/to other namespaces.

    kubectl -n sandbox exec  shell -- curl -s nginx.sandbox # ok
    kubectl -n sandbox exec  shell -- curl -s google.com # ok
    
    kubectl -n sandbox exec  shell -- curl -s nginx.spotify # ko
    kubectl -n spotify exec  shell -- curl -s nginx.sandbox # ko


> [!TIP]
> NetworkPolicies are additive and allow-only : there is no deny rule.
> As soon as *one* policy with `policyTypes: Egress` selects a pod, everything not explicitly
> allowed is dropped — including DNS to CoreDNS. That is why `04` breaks name resolution
> and `05` has to give port 53 back (TCP *and* UDP).

CIDR
* 0.0.0.0/0 => every ip
* 1.2.3.4/32 => only ip 1.2.3.4