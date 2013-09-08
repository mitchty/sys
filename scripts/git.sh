#!/usr/bin/env sh
#-*-mode: Shell-script; coding: utf-8;-*-
script=$(basename $0)
dir=$(cd $(dirname $0); pwd)
iam=${dir}/${script}

git_hash=$(git rev-parse --short HEAD 2> /dev/null)
git_hash=${git_hash:=no_git_hash}

git_tag=$(git describe --tag HEAD 2> /dev/null)
git_tag=${git_tag:=no_git_tag}

git_build_date=$(${dir}/iso8601)
git_build_date=${git_build_date:=no_build_date}

cat <<FIN
package git

const (
  Hash      = "${git_hash}"
  Tag       = "${git_tag}"
  Date      = "${git_build_date}"
)
FIN
