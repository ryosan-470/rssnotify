GOPACKAGES := $(shell go list ./... | grep -v /vendor/)

test:
	go test -v $(GOPACKAGES)