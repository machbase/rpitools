.PHONY: all httpcamsvr lcdctl

all: httpcamsvr lcdctl

httpcamsvr:
	GO111MODULE=on CGO_ENABLED=1 go build -o tmp/httpcamsvr httpcamsvr/*

lcdctl:
	GO111MODULE=on CGO_ENABLED=1 go build -o tmp/lcdctl lcdctl/*
