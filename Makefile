.PHONY: go get github.com/kardianos/govendor

test-cover:
	govendor test -coverprofile=coverage.out
	govendor tool cover -html=coverage.out

test:
	govendor test -cover +local

run:
	go run cmd/main.go

install:
	godep restore

swagger-validate:
	swagger validate ./swagger.json

swagger-serve:
	swagger serve --flavor=swagger ./swagger.json

swagger:
	swagger validate ./swagger.json
	swagger serve --flavor=swagger ./swagger.json
