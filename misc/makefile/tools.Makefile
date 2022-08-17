# This makefile should be used to hold functions/variables

define github_url
    https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef

# creates a directory bin.
bin:
	@ mkdir -p $@

# ~~~ Tools ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ~~ [ gotestsum ] ~~~ https://github.com/gotestyourself/gotestsum ~~~~~~~~~~~~~~~~~~~~~~~

GOTESTSUM := $(shell command -v gotestsum || echo "bin/gotestsum")
gotestsum: bin/gotestsum ## Installs gotestsum (testing go code)

bin/gotestsum: VERSION := 1.8.1
bin/gotestsum: GITHUB  := gotestyourself/gotestsum
bin/gotestsum: ARCHIVE := gotestsum_$(VERSION)_$(OSTYPE)_amd64.tar.gz
bin/gotestsum: bin
	@ printf "Install gotestsum... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - gotestsum > $@ && chmod +x $@
	@ echo "done."

# ~~ [ tparse ] ~~~ https://github.com/mfridman/tparse ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

TPARSE := $(shell command -v tparse || echo "bin/tparse")
tparse: bin/tparse ## Installs tparse (testing go code)

bin/tparse: VERSION := 0.11.1
bin/tparse: GITHUB  := mfridman/tparse
bin/tparse: ARCHIVE := tparse_$(OSTYPE)_x86_64
bin/tparse: bin
	@ printf "Install tparse... "
	@ echo $(ARCHIVE)
	@ curl -Ls $(call github_url) > $@ && chmod +x $@
	@ echo "done."

# ~~ [ golangci-lint ] ~~~ https://github.com/golangci/golangci-lint ~~~~~~~~~~~~~~~~~~~~~

GOLANGCI := $(shell command -v golangci-lint || echo "bin/golangci-lint")
golangci-lint: bin/golangci-lint ## Installs golangci-lint (linter)

bin/golangci-lint: VERSION := 1.48.0
bin/golangci-lint: GITHUB  := golangci/golangci-lint
bin/golangci-lint: ARCHIVE := golangci-lint-$(VERSION)-$(OSTYPE)-amd64.tar.gz
bin/golangci-lint: bin
	@ printf "Install golangci-linter... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - $(shell printf golangci-lint-$(VERSION)-$(OSTYPE)-amd64/golangci-lint | tr A-Z a-z ) > $@ && chmod +x $@
	@ echo "done."
