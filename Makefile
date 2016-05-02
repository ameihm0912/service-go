PROJS = stest
GO = GOPATH=$(shell pwd):$(shell go env GOROOT)/bin go

all: $(PROJS)

stest:
	$(GO) install stest

clean:
	rm -rf pkg
	rm -f bin/*
