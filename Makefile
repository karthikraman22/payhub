ARGS = $(filter-out $@,$(MAKECMDGOALS))
MAKEFLAGS += --silent
PACKAGE = achuala.in/payhub
PROTOC_GEN_VALIDATE_VERSION=v0.6.1

.PHONY: build

list:
	sh -c "echo; $(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | grep -v 'Makefile'| sort"

pbgen: 
	protoc -I. -I./proto -I$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@${PROTOC_GEN_VALIDATE_VERSION} \
	--go_out="module=${PACKAGE}:." \
	--validate_out="module=${PACKAGE},lang=go:." \
	./proto/*.proto*

clean:
	rm -rf pbgen
#############################
# Argument fix workaround
#############################
%:
	@: