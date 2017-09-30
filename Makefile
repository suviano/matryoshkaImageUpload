.PHONY: go get github.com/kardianos/govendor

test-cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test -cover

install:
	godep get

run:
