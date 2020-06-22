SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:

init:
	(cd .. && make install)
	terraform init

test: init
	terraform plan

clean:
	rm -Rf .terraform *.tfstate*
