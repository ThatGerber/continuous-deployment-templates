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
PLANTIMESTAMP = $(TF_DIR)/.tfplan
SRCS += $(shell find $(SOURCEDIR) -name '*.tf')
SRCS += $(shell find $(SOURCEDIR) -name '*.tfvars')
SRCS += $(shell find $(SOURCEDIR) -name '*.json')

BACKEND_CONFIG_FILE = backend.tfvars

# Example : $(call tf,version --help)
define tf
$(TF) $(1) $(TF_FLAGS)
endef

CLEAN_TARGETS += $(TF_MODULE_DIR)

## default  : Run $(TF) plan and generate planfile.
# State File
$(STATE_FILE) : $(SRCS) | $(TF_DIR) $(TF_MODULE_DIR) validate
ifeq ($(strip $(IGNORE_INIT)),)
	@$(MAKE) init
else
	$(info Skipping terraform initilization)
endif
	$(call tf,plan)

## init     : Initialize Terraform backend
init : TF_FLAGS = -backend=true \
-backend-config=$(BACKEND_CONFIG_FILE) \
-force-copy \
-input=false
init :
	$(call tf,init)

$(TF_DIR) :
	mkdir -p $@

$(TF_MODULE_DIR) : TF_FLAGS = -update=true
$(TF_MODULE_DIR) : $(SRCS)
	$(call tf,get)

## clean    : Remove generated files like modules
clean : $(CLEAN_TARGETS)
	rm -rf $^

## plan     : Generate terraform plan. Same as `default`
plan : $(STATE_FILE)
	touch $(PLANTIMESTAMP);

define tf-apply
@ERROR=; \
for f in $^; do \
	if [ $$f -nt $(PLANTIMESTAMP) ]; then \
		ERROR=true; \
		echo 'A source file has changed since last plan. Not applying'; \
	fi \
done;
test -z $$ERROR && $(call tf,apply)
endef

## apply    : Apply plan file. Will fail if files have changed since the plan.
apply : $(SRCS)
	$(call tf-apply)

## validate : Run validation on Terraform files.
validate :
	$(call tf,$@)

## fmt      : Format TF files in correct format.
fmt :
	$(call tf,$@)

## help     : This help text.
help : Makefile
	$(info --- Terraform ---)
	@echo ''
	@sed -n 's/^## //p' $<

.PHONY: apply clean fmt help init plan validate
