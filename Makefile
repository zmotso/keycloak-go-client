## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
OAPICODEGEN ?= $(LOCALBIN)/oapi-codegen

.PHONY: oapi-codegen
oapi-codegen: $(OAPICODEGEN) ## Download oapi-codegen locally if necessary.
$(OAPICODEGEN): $(LOCALBIN)
	$(call go-install,$(OAPICODEGEN),github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen,latest)

.PHONY: generate-keycloak-go-client
generate-keycloak-go-client: oapi-codegen
	#curl -o openapi.yaml https://www.keycloak.org/docs-api/latest/rest-api/openapi.yaml
	$(OAPICODEGEN) -config oapicfg.yaml openapi.yaml

# go-install will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef
