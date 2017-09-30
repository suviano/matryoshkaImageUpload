.PHONY: go get github.com/kardianos/govendor

test-cover:
	godep get
	godep restore
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	godep get
	godep restore
	go test -cover

install:
	godep get
	godep restore

run:
