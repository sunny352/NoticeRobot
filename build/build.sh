#!/usr/bin/env bash

emulate bash

if [[ -z "${VERSION}" ]]; then
  export VERSION="0.0.1"
fi

echo "配置代理"
export GOPROXY=https://goproxy.cn

#兼容不同运行路径
if [[ ! -f go.mod ]]; then
  echo "没有找到go.mod，向上一级"
  cd ..
fi

export current_folder=${PWD}

if [[ ! -d GoPath ]]; then
  mkdir GoPath
fi

echo "配置项目GOPATH"
export GOPATH=${current_folder}/GoPath
export PATH=${GOPATH}/bin:$PATH

echo "处理工程依赖"
go mod tidy -compat=1.17

echo "启动"
go run cmd/main.go
