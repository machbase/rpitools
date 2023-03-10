all: tmp/httpcamsvr

tmp/httpcamsvr:
	go build -o tmp/httpcamsvr httpcamsvr/*