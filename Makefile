.ONESHELL:
SHELL = /bin/bash
# Disable implicit rules
.SUFFIXES:

# Using One shell. Errors if anything fails during the process.
.SHELLFLAGS = -e

# Go Config
GO = go

# bin-data package config
BINDATAPKG = jteeuwen/go-bindata
BINDATA_EXE = $(GOPATH)/bin/go-bindata

# Vendor Config
VENDOR_DIR = vendor
GODEP = $(GOBIN)/dep
GODEP_PKG = github.com/golang/dep/cmd/dep
GOPKG_FILE = Gopkg.toml
GOPKG_LOCKFILE = Gopkg.lock

# Application Binary
BIN_FILE = generate
INSTALL_DIR := .

# Application Source
SOURCEDIR = src
SRCS := $(SOURCEDIR)/main.go
SRCS += $(shell find $(SOURCEDIR) -name '*.go' | grep -v '$(SOURCEDIR)/main.go')
SRCS += $(TEMPLATE_SRCS)

# Template Sources
TEMPLATE_DIR = templates
TEMPLATE_SRCS = $(call templates,$(TEMPLATE_DIR))
TEMPLATE_ASSET_FILES = $(TEMPLATE_DIR)/*/assets.go
ADDL_TEMPLATE_SUFFIXES := *.tfvars *.json Makefile
FIND_FLAGS=$(foreach search,$(ADDL_TEMPLATE_SUFFIXES),-name "$(search)" -o)

# Build Vars
BUILD_ID = $(shell echo "`git rev-parse HEAD | cut -c1-10`.`date "+%s"`")

# Generated files
CLEAN_TARGETS = $(BIN_FILE) $(GOBIN)/$(BIN_FILE) $(TEMPLATE_ASSET_FILES)

# All Files needed by the binary
ALL_FILES = $(SRCS) $(TEMPLATE_ASSET_FILES)

# $(call go,cmd)
define go
$(strip $(GO) $(1) $(GOFLAGS) $(2) )
endef

# $(call go,cmd)
define govet
$(strip $(GO) tool vet $(1) )
endef

# Returns the application's templates.
define templates
$(shell find $(1) \( $(FIND_FLAGS) -name "*.tf" \) -print )
endef

# Overriding Config
include config.mk

# ---
$(BIN_FILE) : GOFLAGS = -o $(BIN_FILE) -ldflags="-X main.buildID=$(BUILD_ID)"
$(BIN_FILE) : $(ALL_FILES)
	@echo '::: Go Vet'
	$(call govet,$(filter %.go,$?))
	@echo '::: Go Build'
	$(call go,build,$<)

# ---
$(GOBIN)/$(BIN_FILE) : $(ALL_FILES)
	@echo '::: Go Build'
	$(call go,build,$<)

install : GOFLAGS = -i -o $(GOBIN)/$(BIN_FILE) -ldflags="-X main.buildID=$(BUILD_ID)"
install : $(GOBIN)/$(BIN_FILE)

# --- Empty build for SRC files.
$(SRCS) : ;

# --- Template Static Asset Files
$(TEMPLATE_ASSET_FILES) : $(TEMPLATE_SRCS) | $(BINDATA_EXE)
	@echo '::: Building $@'
	@$(BINDATA_EXE) -prefix $(dir $@) -o $@ \
		-pkg $(filter-out $(TEMPLATE_DIR) $(notdir $@),$(subst /, ,$@)) \
		$(call templates,$(dir $@));

# --- install bindata package
$(BINDATA_EXE) : GOFLAGS = -v -u
$(BINDATA_EXE) :
	$(call go,get,github.com/$(BINDATAPKG)/...)

# --- Install deps
$(GODEP) : GOFLAGS = -v -u
$(GODEP) :
	$(call go,get,$(GODEP_PKG))

$(GOPKG_LOCKFILE) $(VENDOR_DIR): $(GOPKG_FILE) $(GODEP) $(ALL_FILES)
	dep ensure

$(GOPKG_FILE) : $(GODEP)
	dep init

# --- Phonies
.PHONY: all clean fmt install get run test vet

dep :
	dep ensure

run :
	$(call go,$@,$<)

get : $(BINDATA_EXE) $(GODEP)
	@echo "::: Go Get: $@"
	$(call go,$@ -t,./...)

test :
	@echo '::: Go Test'
	$(call go,$@,./...)

fmt :
	@echo '::: Go Fmt'
	@gofmt -w -s .

vet :
	@echo ":: Go Vet $(SOURCE_DIR)"
	$(call go,tool vet,$(SOURCE_DIR))
	@echo ":: Go Vet $(TEMPLATE_DIR)"
	$(call go,tool vet,$(TEMPLATE_DIR))

clean:
	@echo '::: Clean'
	rm -rf $(CLEAN_TARGETS)