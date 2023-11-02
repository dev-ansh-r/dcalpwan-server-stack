

SHELL = bash
GIT = git
NPM = npm
GO = go
GOBIN = $(PWD)/bin
export GOBIN

.PHONY: default
default: validate

.PHONY: deps
deps:
	$(NPM) install
	pushd ./tools/runscript && $(GO) install . && popd
	pushd ./tools/validate-image && $(GO) install . && popd

.PHONY: deps.update
deps.update:
	$(GIT) subtree pull -P lib/draft https://github.com/lorawan-schema/draft-devices.git v1

.PHONY: validate
validate:
	$(NPM) run validate

.PHONY: fmt
fmt:
	$(NPM) run format

# vim: ft=make
