GO=go
GOPATH=${CURDIR}
export GOPATH

.PHONY:	all
all:	
	$(GO) install ./...

.PHONY:	clean
clean:
	$(GO) clean


.PHONY:	check
check:
	$(GO) test ./...
