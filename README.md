# k8s on AWS EC2

> [!CAUTION]
> Just for CKAD, CKA, CKS purpose. Not production ready

    # debug dynamic inventory
    ansible-inventory -i inventory/aws_ec2.yml --list
    ansible-inventory -i inventory/aws_ec2.yml --graph

    # create AWS infra
    ansible-playbook -i inventory/localhost create-on-aws.yml

    # configure k8s
    ansible-playbook -i inventory/aws_ec2.yml configure-k8s.yml

    # terminate AWS infra
    ansible-playbook -i inventory/aws_ec2.yml terminate-aws.yml
