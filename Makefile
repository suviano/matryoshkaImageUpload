.PHONY: go get github.com/kardianos/govendor

EXECUTABLE_FILES=$(ls -h *.go | grep -v '_test')

test-cover:
	govendor test -coverprofile=coverage.out
	govendor tool cover -html=coverage.out

test:
	govendor test -cover +local

run:
	go install
	./nameless-storage-image-saver

