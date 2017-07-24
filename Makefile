.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -e

.SUFFIXES:

GO = go
GOFLAGS =

# Binary
BIN_FILE = generate

# Source
SOURCEDIR = src
SRCS := $(SOURCEDIR)/main.go
SRCS += $(shell find $(SOURCEDIR) -name '*.go' | grep -v '$(SOURCEDIR)/main.go')

# Templates
TEMPLATE_DIR = templates
TEMPLATE_SRCS = $(call templates,$(TEMPLATE_DIR))
TEMPLATE_ASSET_FILES = $(TEMPLATE_DIR)/*/assets.go

# Generated files
CLEAN_TARGETS = $(BIN_FILE) $(GOBIN)/$(BIN_FILE) $(TEMPLATE_ASSET_FILES)

ALL_FILES = $(SRCS) $(TEMPLATE_ASSET_FILES)

# $(call go,cmd)
define go
$(info > Go $(1): $(@))
$(strip $(GO) $(1) $(GOFLAGS) $(2) )
endef

TEMPLATE_SUFFIXES := *.tfvars *.json Makefile
FIND_FLAGS=$(foreach search,$(TEMPLATE_SUFFIXES),-name "$(search)" -o)

define templates
$(shell find $(patsubst %/,%,$(1)) \
\( $(FIND_FLAGS) -name "*.tf" \) -print )
endef

BUILD_ID=$(shell echo "`git rev-parse HEAD | cut -c1-10`.`date "+%s"`")
INSTALL_DIR := .

$(BIN_FILE) : GOFLAGS = -o $(BIN_FILE) -ldflags="-X main.buildID=$(BUILD_ID)"
$(BIN_FILE) : $(ALL_FILES)
	$(MAKE) vet
	$(MAKE) fmt
	$(call go,build,$<)

$(GOBIN)/$(BIN_FILE) : $(ALL_FILES)
	$(call go,build,$<)

$(SRCS) : ;

$(TEMPLATE_ASSET_FILES) : $(TEMPLATE_SRCS) | $(BINDATA_EXE)
	$(info Building $@)
	@$(BINDATA_EXE) -prefix $(dir $@) -o $@ \
		-pkg $(filter-out $(TEMPLATE_DIR) $(notdir $@),$(subst /, ,$@)) \
		$(call templates,$(dir $@));

$(TEMPLATE_SRCS): ;
BINDATAPKG = jteeuwen/go-bindata
BINDATA_EXE = $(GOPATH)/bin/go-bindata

$(BINDATA_EXE) : GOFLAGS = -v -u
$(BINDATA_EXE) :
	$(call go,get,github.com/$(BINDATAPKG)/...)

.PHONY: all clean fmt get run test vet

install : GOFLAGS += -i -o $(GOBIN)/$(BIN_FILE) -ldflags="-X main.buildID=$(BUILD_ID)"
install : $(GOBIN)/$(BIN_FILE)

run :
	$(call go,$@,$<)

get : $(BINDATA_EXE)
	$(call go,$@ -t,./...)

test :
	$(call go,$@,./...)

fmt :
	$(info > Go $@:)
	@gofmt -w -s .

vet :
	$(call go,tool vet,$(SOURCEDIR))
	$(call go,tool vet,$(TEMPLATE_DIR))

clean:
	rm -rf $(CLEAN_TARGETS)
