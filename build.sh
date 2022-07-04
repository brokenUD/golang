#!/usr/bin/env bash


TARGET=$1

CURPATH=`pwd`
export GO111MODULE=on
mkdir -p ./bin

if [[ ${TARGET} == "gateway" ]]
then
	cd gateway
	go build -mod vendor -o ${CURPATH}/bin/gateway
	cd -
fi
