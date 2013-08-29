#!/usr/bin/env sh
#-*-mode: Shell-script; coding: utf-8;-*-
script=$(basename $0)
dir=$(cd $(dirname $0); pwd)
iam=${dir}/${script}

git_hash=$(git rev-parse --short HEAD 2> /dev/null)
git_hash=${git_hash:=unknown}

cat <<FIN
package main

const releaseGitVersionHash = "$git_hash"
FIN
