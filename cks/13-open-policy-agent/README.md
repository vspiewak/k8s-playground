## Open Policy Agent

* [Documentation](https://open-policy-agent.github.io/gatekeeper/website/docs/)
* [Policy Library](https://open-policy-agent.github.io/gatekeeper-library/website/)
* [Rego Playground](https://play.openpolicyagent.org/)


### installation

    # Helm install
    helm repo add gatekeeper https://open-policy-agent.github.io/gatekeeper/charts
    helm install gatekeeper/gatekeeper --name-template=gatekeeper --namespace gatekeeper-system --create-namespace

    # add constraint template disallowedtags
    kubectl apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper-library/master/library/general/disallowedtags/template.yaml

    # view it
    kubectl get constrainttemplates
    NAME                AGE
    k8sdisallowedtags   39s

    # disallow latest image
    kubectl apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper-library/master/library/general/disallowedtags/samples/container-image-must-not-have-latest-tag/constraint.yaml

    # view it
    kubectl get k8sdisallowedtags.constraints.gatekeeper.sh
    NAME                                       ENFORCEMENT-ACTION   TOTAL-VIOLATIONS
    container-image-must-not-have-latest-tag   

    # test it 
    kubectl run nginx-latest --image=nginx:latest
    Error from server (Forbidden): admission webhook "validation.gatekeeper.sh" denied the request: [container-image-must-not-have-latest-tag] container <nginx-latest> uses a disallowed tag <nginx:latest>; disallowed tags are ["latest"]

### view violations

```
kubectl describe k8sdisallowedtags.constraints.gatekeeper.sh
Name:         container-image-must-not-have-latest-tag
Namespace:
Labels:       <none>
Annotations:  <none>
API Version:  constraints.gatekeeper.sh/v1beta1
Kind:         K8sDisallowedTags
Metadata:
  Creation Timestamp:  2024-03-04T20:51:13Z
  Generation:          1
  Resource Version:    3084
  UID:                 bbbf6db6-7eb6-4df8-bded-190d0b0e5683
Spec:
  Match:
    Kinds:
      API Groups:

      Kinds:
        Pod
    Namespaces:
      default
  Parameters:
    Exempt Images:
      openpolicyagent/opa-exp:latest
      openpolicyagent/opa-exp2:latest
    Tags:
      latest
Status:
  Audit Timestamp:  2024-03-04T20:53:40Z
  By Pod:
    Constraint UID:       bbbf6db6-7eb6-4df8-bded-190d0b0e5683
    Enforced:             true
    Id:                   gatekeeper-audit-64b9df68c6-pdqlt
    Observed Generation:  1
    Operations:
      audit
      mutation-status
      status
    Constraint UID:       bbbf6db6-7eb6-4df8-bded-190d0b0e5683
    Enforced:             true
    Id:                   gatekeeper-controller-manager-8b9df7bf5-2hhqw
    Observed Generation:  1
    Operations:
      mutation-webhook
      webhook
    Constraint UID:       bbbf6db6-7eb6-4df8-bded-190d0b0e5683
    Enforced:             true
    Id:                   gatekeeper-controller-manager-8b9df7bf5-8qpth
    Observed Generation:  1
    Operations:
      mutation-webhook
      webhook
    Constraint UID:       bbbf6db6-7eb6-4df8-bded-190d0b0e5683
    Enforced:             true
    Id:                   gatekeeper-controller-manager-8b9df7bf5-rzvtp
    Observed Generation:  1
    Operations:
      mutation-webhook
      webhook
  Total Violations:  1
  Violations:
    Enforcement Action:  deny
    Group:
    Kind:                Pod
    Message:             container <app> didn't specify an image tag <ubuntu>
    Name:                app
    Namespace:           default
    Version:             v1
Events:                  <none>
```