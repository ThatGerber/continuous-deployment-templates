.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -e

# Resulting binary
TARGET = generate

# Source Folder
SOURCEDIR = src
SRCS := main.go
SRCS += $(shell find $(SOURCEDIR) -name '*.go')
SRCS += $(shell find $(SOURCEDIR) -name '*.tf')
SRCS += $(shell find $(SOURCEDIR) -name '*.tfvars')
SRCS += $(shell find $(SOURCEDIR) -name '*.json')

# Templates Folder
TMPLDIR = templates
SRCS += $(shell find $(TMPLDIR) -name '*.go')
SRCS += $(shell find $(TMPLDIR) -name '*.tf')
SRCS += $(shell find $(TMPLDIR) -name '*.tfvars')
SRCS += $(shell find $(TMPLDIR) -name '*.json')

CLEAN_TARGETS = $(TARGET)

SANDBOX_DIR = sandbox

CLEAN_TARGETS += $(SANDBOX_DIR)

GO := go
GOFLAGS =

# $(call go,cmd)
define go
	$(GO) $(1) $(GOFLAGS) $(2)
endef

.PHONY: all clean fmt run

all : $(SANDBOX_DIR)/$(TARGET)

$(SANDBOX_DIR)/$(TARGET) : $(SANDBOX_DIR) $(TARGET)
	cp $(TARGET) $@

$(TARGET) : GOFLAGS = -o $(TARGET)
$(TARGET) : $(SRCS)
	$(call go,build,$<)

$(SANDBOX_DIR) :
	mkdir -p $@

clean: $(CLEAN_TARGETS)
	rm -rf $^

fmt : $(SRCS)
	@gofmt -w -s .

run : $(SRCS)
	$(call go,run,$<)

test : $(SRCS)
	$(call go,test,.../.)
