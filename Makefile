.PHONY: go get github.com/kardianos/govendor

install:
	godep get
	godep restore

test-cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test -cover
