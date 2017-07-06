.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -e

# Resulting binary
TARGET = generate

# Source Folder
SOURCEDIR = src
SRCS := $(SOURCEDIR)/main.go
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

CLEAN_TARGETS = $(TARGET) $(GOBIN)/$(TARGET)

GO = go
GOFLAGS =

# $(call go,cmd)
define go
	$(GO) $(1) $(GOFLAGS) $(2)
endef

.PHONY: all clean fmt install run test

all : $(TARGET)

$(TARGET) : GOFLAGS = -x -o $(TARGET)
$(TARGET) : $(SRCS)
	$(call go,build,$<)

install : GOFLAGS = -x -i -o $(GOBIN)/$(TARGET)
install : $(SRCS)
	$(info Removing local binary $(TARGET))
	rm -f $(TARGET)
	$(call go,build,$<)

clean:
	rm -rf $(CLEAN_TARGETS)

fmt : $(SRCS)
	@gofmt -w -s .

run : $(SRCS)
	$(call go,$@,$<)

test : $(SRCS)
	$(call go,$@,.../.)
