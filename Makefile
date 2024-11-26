GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./coverage/cover.out -covermode=set -coverpkg=./...
	go tool cover -html=./coverage/cover.out -o=./coverage/cover.html
	${GOBIN}/go-test-coverage --config=./config/.testcoverage.yml
