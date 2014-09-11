#!/bin/bash

tooling_gopath=/tmp/tooling-gopath

ensure_go_binary() {
  binary_name=$1
  url=$2

  if ! which "${binary_name}" > /dev/null ; then
    echo "${binary_name} isn't installed globally. Installing it would make this process a bit faster."
    mkdir -p "${tooling_gopath}"
    echo "+ go get ${url}"
    GOPATH="${tooling_gopath}" go get "${url}"
    export PATH="${PATH}:${tooling_gopath}/bin"
  fi
}


