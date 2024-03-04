[conftest](https://www.conftest.dev)

```
mkdir -p policy

cat <<EOF > policy/base.rego
package main

denylist = [
  "ubuntu"
]

deny[msg] {
  input[i].Cmd == "from"
  val := input[i].Value
  contains(val[i], denylist[_])

  msg = sprintf("unallowed image found %s", [val])
}
EOF

cat <<EOF > policy/commands.rego
package commands

denylist = [
  "apk",
  "apt",
  "pip",
  "curl",
  "wget",
]

deny[msg] {
  input[i].Cmd == "run"
  val := input[i].Value
  contains(val[_], denylist[_])

  msg = sprintf("unallowed commands found %s", [val])
}
EOF

cat <<EOF > Dockerfile
FROM ubuntu
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y golang-go
COPY app.go .
RUN go build app.go
CMD ["./app"]
EOF
```

```
docker run --rm -v $(pwd):/project openpolicyagent/conftest test Dockerfile --all-namespaces
```