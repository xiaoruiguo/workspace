VERSION=1.0.0
COMPILE=$(shell date -u "+%Y-%m-%d/%H:%M:%S")

.PHONY: build 

build:
	go build -ldflags "-X github.com/lijiaocn/GoPkgs/version.VERSION=$(VERSION) -X github.com/lijiaocn/GoPkgs/version.COMPILE=$(COMPILE) "
