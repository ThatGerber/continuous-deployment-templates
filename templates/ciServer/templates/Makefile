.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -e

CLEAN_TARGETS =

TF        = terraform
SOURCEDIR := ./
TF_DIR    = .terraform

IGNORE_INIT   = true
STATE_FILE    = $(TF_DIR)/terraform.tfstate
TF_MODULE_DIR = $(TF_DIR)/modules
SRCS += $(shell find $(SOURCEDIR) -name '*.tf')
SRCS += $(shell find $(SOURCEDIR) -name '*.tfvars')
SRCS += $(shell find $(SOURCEDIR) -name '*.json')

# Example : $(call tf,version --help)
define tf
$(TF) $(1) $(TF_FLAGS)
endef

CLEAN_TARGETS += $(TF_MODULE_DIR)

.DEFAULT_GOAL := $(STATE_FILE)

## init     : Initialize Terraform backend
init : TF_FLAGS = -backend=true \
-get=true \
-force-copy \
-input=false
init :
	$(call tf,init)

$(TF_DIR) :
	mkdir -p $@

$(TF_MODULE_DIR) :
	$(call tf,get)

## clean    : Remove generated files like modules
clean : $(CLEAN_TARGETS)
	rm -rf $^

## plan     : Generate terraform plan. Same as `default`
plan : $(STATE_FILE)

define tf-apply
@ERROR=; \
for f in $^; do \
	if [ $$f -nt $(STATE_FILE) ]; then \
		ERROR=true; \
		echo 'A source file has changed since last plan. Not applying'; \
	fi \
done;
test -z $$ERROR && $(call tf,apply)
endef

## apply    : Apply plan file.
apply : $(SRCS)
	$(call tf-apply)

## validate : Run validation on Terraform files.
validate :
	$(call tf,$@)

## fmt      : Format TF files in correct format.
fmt :
	$(call tf,$@)

## default  : Run $(TF) plan and generate planfile.
# State File
$(STATE_FILE) : $(TF_DIR) $(TF_MODULE_DIR) $(SRCS) validate
ifeq ($(strip $(IGNORE_INIT)),)
	@$(MAKE) init
else
	$(info Skipping terraform initilization)
endif
	$(call tf,plan)

## help     : This help text.
help : Makefile
	$(info --- Terraform ---)
	@echo ''
	@sed -n 's/^## //p' $<

.PHONY: apply clean fmt help init plan validate
