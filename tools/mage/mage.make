

MAGE ?= tools/bin/mage
MAGE_ABSOLUTE = $(abspath $(MAGE))

$(MAGE): tools/magefile.go $(wildcard tools/mage/*.go)
	@cd tools && GO111MODULE=on go run github.com/magefile/mage -compile $(MAGE_ABSOLUTE)

.PHONY: init
init: $(MAGE)
	@$(MAGE) init
	@echo "Run \"$(MAGE) -l\" for a list of build targets"

.PHONY: git.pre-commit
git.pre-commit: $(MAGE) # NOTE: DO NOT CHANGE - will break previously installed git hooks.
	

.PHONY: git.commit-msg
git.commit-msg: $(MAGE) # NOTE: DO NOT CHANGE - will break previously installed git hooks.
	

.PHONY: git.pre-push
git.pre-push: $(MAGE) # NOTE: DO NOT CHANGE - will break previously installed git hooks.
	@HOOK=pre-push $(MAGE) git:runHook

# vim: ft=make
