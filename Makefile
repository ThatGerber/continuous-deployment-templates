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
TMPLSRCS = $(shell find $(TMPLDIR) -name '*.tf')
TMPLSRCS += $(shell find $(TMPLDIR) -name '*.tfvars')
TMPLSRCS += $(shell find $(TMPLDIR) -name '*.json')
TMPLSRCS += $(shell find $(TMPLDIR) -name 'Makefile')

SRCS += $(TMPLDIR)/*/assets.go

CLEAN_TARGETS = $(TARGET) $(GOBIN)/$(TARGET)

GO = go
GOFLAGS =

BINDATAPKG = jteeuwen/go-bindata
BINDATA_EXE = $(GOPATH)/bin/go-bindata

# $(call go,cmd)
define go
$(GO) $(1) $(GOFLAGS) $(2)
endef

define templates
find $(1) \
\( -name "*.tf" -o -name "*.tfvars" -o -name "*.json" -o -name "Makefile" \) \
-print
endef

.PHONY: all clean fmt install run test

all : $(TARGET)

PKG_TEMPLATE_DIRS = `cat $(PACKAGE_TEMPLATES_FILE)`

$(TARGET) : GOFLAGS = -o $(TARGET)
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

get : $(SRCS) $(BINDATA_EXE)
	$(call go,$@ -t,./...)

test : $(SRCS)
	$(call go,$@,./...)

TEMPLATE_DIRS = `find $(TMPLDIR) -type d -depth '1' -print0 | xargs -0 -n1 -x basename`

$(BINDATA_EXE) : GOFLAGS = -v -u
$(BINDATA_EXE) :
	$(call go,get,github.com/$(BINDATAPKG)/...)

$(TMPLDIR)/*/assets.go : $(TMPLSRCS) | $(BINDATA_EXE)
	$(info Compiling static assets...)
	@for d in $(TEMPLATE_DIRS); do \
		templates="`$(call templates,$(TMPLDIR)/$$d)`"; \
		assets="`echo $$templates | tr "\n" " "`"; \
		$(BINDATA_EXE) -pkg $$d -prefix $(TMPLDIR)/$$d -o $(TMPLDIR)/$$d/assets.go $$assets; \
	done;
