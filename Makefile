.default: build

build: lint
	go install ./...

lint:
	go fmt ./...
	golint -set_exit_status $$(go list ./... | grep -v /vendor/)

test: build
	go test -timeout 1s ./...

cover: test
	goverage -covermode=set -coverprofile=cov.out `go list ./...`
	gocov convert cov.out | gocov report

coverhtml: cover
	go tool cover -html=cov.out
