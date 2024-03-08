
    kubectl apply -f https://raw.githubusercontent.com/aquasecurity/kube-bench/main/job.yaml

    # config
    cat /etc/kube-bench/cfg/config.yaml

    # run
    kube-bench run

    # run specific target
    kube-bench run --targets=master

    # run specific checks
    kube-bench -c 1.1.12
    kube-bench -c 1.1.12,1.3.2