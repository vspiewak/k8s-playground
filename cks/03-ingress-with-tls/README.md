    # install ingress-nginx with service type NodePort
    helm upgrade --install ingress-nginx ingress-nginx \
    --repo https://kubernetes.github.io/ingress-nginx \
    --namespace ingress-nginx --create-namespace \
    --set controller.service.type=NodePort

    # check ingress-nginx controller service
    kubectl -n ingress-nginx get service/ingress-nginx-controller

    # check HTTP service1 & service2
    curl ec2-13-38-47-153.eu-west-3.compute.amazonaws.com:32182/service1
    curl ec2-13-38-47-153.eu-west-3.compute.amazonaws.com:32182/service2

    # check HTTPS service1 & service2 (CN=Kubernetes Ingress Controller Fake Certificate)
    curl https://ec2-13-38-47-153.eu-west-3.compute.amazonaws.com:31234/service1 -kv
    curl https://ec2-13-38-47-153.eu-west-3.compute.amazonaws.com:31234/service2 -kv

    # generate certificate
    openssl req -x509 -newkey rsa:4096 \
    -keyout key.pem -out cert.pem -days 365 -nodes \
    -subj '/CN=secure-domain.com'

    # create tls-secret
    kubectl -n spotify create secret tls tls-secret --cert=cert.pem --key=key.pem

    # check HTTPS (CN=secure-domain.com)
    curl https://secure-domain.com:31234/service1 --resolve secure-domain.com:31234:13.38.47.153 -kv
