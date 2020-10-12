GO:=$(or ${GO_EXEC},go)
MAGE = ./tools/bin/mage

APP=carlware/accounts
DATE=$(or ${BUILD_DATE},"`date '+%Y-%m-%dT%H:%M:%S'`")
VER=$(or ${VERSION},"`git describe --tags --abbrev=0 | cut -c 2-`")
VREF=$(or ${VCS_REF},"`git rev-parse --abbrev-ref HEAD`")
GO_VER=$(or, ${GO_VERSION},"`go version | awk '{print $$3}'`")

TARGETS := $(shell $(MAGE) -l | awk '\
	BEGIN { count = 1 }\
	/^  / { gsub(/:/, ".", $$1); if (sub(/\*$$/, "", $$1) > 0) { targets[0] = $$1 } else { targets[count++] = $$1 } }\
	END   { for (i = 0; i < count; i++) { print targets[i] } }\
')

.PHONY: $(TARGETS)

run:
	$(GO) run cli/main.go serve -c config.yaml

install-tools:
	$(GO) run github.com/magefile/mage -d tools

gounit-init:
	./tools/bin/gounit template add tools/gounit/init
	./tools/bin/gounit template add tools/gounit/model
	./tools/bin/gounit template add tools/gounit/gomock

gounit-clean:
	./tools/bin/gounit template remove init model gomock

gounit:
	./tools/bin/gounit template use $(m)
	./tools/bin/gounit gen -i $(t)

graph:
	go run github.com/99designs/gqlgen generate --config cli/dispatchers/graphql/gqlgen.yml

clean:
	rm -rf bin vendor

clean-all:
	rm -rf tools/vendor tools/bin
	rm -rf bin vendor

docker-build:
	docker build -t "$(APP):latest" -t "$(APP):$(VER)" --build-arg BUILD_DATE=$(DATE) --build-arg VERSION=$(VER) --build-arg VCS_REF=$(VREF) --build-arg GIT_ID --build-arg GIT_TOKEN .

$(TARGETS):
	@$(MAGE) $(subst .,:,$@)
