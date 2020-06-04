SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:

export CGO_ENABLED?=0
export TERRAFORM_PLUGIN_DIR?=$(HOME)/.terraform.d/plugins

generate:
	rm -Rf pkg/api/*/
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client \
		--name=victorops \
		--spec=pkg/api/victorops-api-v1.yaml \
		--target=pkg/api \
		--with-expand \
		--skip-validation \
		--skip-tag-packages \
		--tags=On-call

build:
	rm -Rf bin
	mkdir -p bin
	go run github.com/mitchellh/gox \
		-osarch="darwin/amd64 linux/amd64 windows/amd64" \
		-ldflags="-s -w" \
		-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

test:
	(gofmt -d -s -e . 2>&1 | tee -a fmt.out) && test ! -s fmt.out
	go run golang.org/x/lint/golint -set_exit_status .
	go test -timeout 10s -v ./pkg/provider/...

install:
	go build -o bin/terraform-provider-victorops
	mv bin/terraform-provider-victorops $(TERRAFORM_PLUGIN_DIR)/
