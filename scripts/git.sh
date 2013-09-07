#!/usr/bin/env sh
#-*-mode: Shell-script; coding: utf-8;-*-
script=$(basename $0)
dir=$(cd $(dirname $0); pwd)
iam=${dir}/${script}

git_hash=$(git rev-parse --short HEAD 2> /dev/null)
git_hash=${git_hash:=unknown}

git_tag=$(git describe --tag HEAD 2> /dev/null)
git_tag=${git_tag:=unknown}

git_build_date=$(${dir}/iso8601)
git_build_date=${git_build_date:=unknown_build_date}

cat <<FIN
package main

const (
  GitHash      = "${git_hash}"
  GitTag       = "${git_tag}"
  GitBuildDate = "${git_build_date}"
)
FIN
