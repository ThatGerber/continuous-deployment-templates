SHELL = /bin/bash
#.SHELLFLAGS = -e

GO = go
GOFLAGS =

# Resulting binary
EXE_FILE = generate

TARGET = $(EXE_FILE)
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

BINDATAPKG = jteeuwen/go-bindata
BINDATA_EXE = $(GOPATH)/bin/go-bindata

# $(call go,cmd)
define go
$(strip $(GO) $(1) $(GOFLAGS) $(2) )
endef

define templates
find $(1) \
\( -name "*.tf" -o -name "*.tfvars" -o -name "*.json" -o -name "Makefile" \) \
-print
endef

.PHONY: all clean fmt get install run test vet

all : $(TARGET)

$(TARGET) : GOFLAGS += -o $(TARGET)
$(TARGET) : $(SRCS) vet fmt
	$(info > Go Build: $(TARGET))
	$(call go,build,$<)

$(GOBIN)/$(EXE_FILE) : TARGET = $(GOBIN)/$(EXE_FILE)
$(GOBIN)/$(EXE_FILE) : GOFLAGS = -i
$(GOBIN)/$(EXE_FILE) : $(SRCS) $(TARGET)

$(BINDATA_EXE) : GOFLAGS = -v -u
$(BINDATA_EXE) :
	$(info > Go Get: $(BINDATAPKG))
	$(call go,get,github.com/$(BINDATAPKG)/...)

TEMPLATE_DIRS = `find $(TMPLDIR) -type d -depth '1' -print0 | xargs -0 -n1 -x basename`

$(TMPLDIR)/*/assets.go : $(TMPLSRCS) | $(BINDATA_EXE)
	$(info Compiling Go files from assets...)
	@for d in $(TEMPLATE_DIRS); do \
		templates="`$(call templates,$(TMPLDIR)/$$d)`"; \
		assets="`echo $$templates | tr "\n" " "`"; \
		outfile="$(TMPLDIR)/$$d/assets.go"; \
		$(BINDATA_EXE) -pkg $$d -prefix $(TMPLDIR)/$$d -o $(TMPLDIR)/$$d/assets.go $$assets; \
		echo "Created asset file $$outfile"; \
	done;

install : $(GOBIN)/$(EXE_FILE)

clean:
	rm -rf $(CLEAN_TARGETS)

fmt : $(SRCS)
	$(info > Go $@: $<)
	@gofmt -w -s .

run : $(SRCS)
	$(info > Go $@:)
	$(call go,$@,$<)

get : $(SRCS) $(BINDATA_EXE)
	$(info > Go $@:)
	$(call go,$@ -t,./...)

test : $(SRCS)
	$(info > Go $@:)
	$(call go,$@,./...)

vet : $(SRCS)
	$(info > Go $@: $(SOURCEDIR)/ $(TMPLDIR)/)
	$(shell go tool vet $(SOURCEDIR))
	$(shell go tool vet $(TMPLDIR))
