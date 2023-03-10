.PHONY: all httpcamsvr

all: httpcamsvr

httpcamsvr:
	GO111MODULE=on CGO_ENABLED=1 go build -o tmp/httpcamsvr httpcamsvr/*