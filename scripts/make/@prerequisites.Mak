$(eval WHICH_GAWK := $(strip $(shell which gawk)))
gawk_found = $(strip $(findstring gawk, $(WHICH_GAWK)))

$(eval WHICH_GO := $(strip $(shell which go)))
go_found = $(strip $(findstring go, $(WHICH_GO)))

$(eval WHICH_NODE := $(strip $(shell which node)))
node_found = $(strip $(findstring node, $(WHICH_NODE)))

prerequisite-test: .bin-validator
	@echo "test!"

.bin-validator: ## validate require binaries exist
# exit status code 126 - Command invoked cannot execute
ifneq ($(gawk_found),gawk)
	@echo "$(__prefix) gawk binary not found. install gawk for mac https://formulae.brew.sh/formula/gawk, or install gawk for ubuntu https://howtoinstall.co/en/gawk"
	@exit 126
endif

ifneq ($(go_found),go)
	@echo "$(__prefix) go binary not found, follow golang installation https://go.dev/doc/install"
	@exit 126
endif

ifneq ($(node_found),node)
	@echo "$(__prefix) node binary not found, follow nodeJS installation https://github.com/nvm-sh/nvm?tab=readme-ov-file#installing-and-updating"
	@exit 126
endif