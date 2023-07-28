#!/bin/bash

set -e -o pipefail

if [ -z "$1" ]; then
    echo "Usage: $0 export/to/path.md"
    exit 1
fi

TMP_FILE=$(mktemp)
trap "rm $TMP_FILE" EXIT

# use the generated proto documentation as source
cp README.md $TMP_FILE

# cleanup the generated documentation for the website
# remove the title
sed -i '/^# tetragon$/d' $TMP_FILE
# remove the badges
sed -i '/^\!\[Version:/d' $TMP_FILE
# remove the "Helm chart for Tetragon mention"
sed -i '/^Helm chart for Tetragon$/d' $TMP_FILE
# remove unnecessary line
sed -i '/^----------------------------------------------$/d' $TMP_FILE
# remove autogenerated copyright
sed -i '/^Autogenerated from chart metadata using \[helm-docs /d' $TMP_FILE
# cleanup unecessary consecutive whitelines
sed -i '/^[[:space:]]*$/N;/^[[:space:]]*\n[[:space:]]*$/D' $TMP_FILE


# add a frontmatter and a small introduction
echo '---
title: "Helm chart"
description: "This reference is generated from the Tetragon Helm chart values."
---

{{< comment >}}
This page was generated with github.io/cilium/tetragon/install/kubernetes/export-doc.sh,
please do not edit directly.
{{< /comment >}}

The Tetragon Helm chart source is available under 
[github.io/cilium/tetragon/install/kubernetes](https://github.com/cilium/tetragon/tree/main/install/kubernetes)
and is distributed from the Cilium helm charts repository [helm.cilium.io](https://helm.cilium.io).

To deploy Tetragon using this Helm chart you can run the following commands:
```shell-session
helm repo add cilium https://helm.cilium.io
helm repo update
helm install tetragon cilium/tetragon -n kube-system
```

To use [the values available](#values), with `helm install` or `helm upgrade`, use `--set key=value`.' | cat - $TMP_FILE > $1

