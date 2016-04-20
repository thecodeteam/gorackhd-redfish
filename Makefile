export GO15VENDOREXPERIMENT=1

GLIDE := $(GOPATH)/bin/glide
GLIDE_YAML := glide.yaml
GLIDE_LOCK := glide.lock
GLIDE_SWAGGER_YAML := glide-swagger.yaml
GO_GET := .go.get
SWAGGER := $(GOPATH)/bin/swagger
SWAGGER_PKG := github.com/go-swagger/go-swagger
SWAGGER_CMD_SRC := $(SWAGGER_PKG)/cmd/swagger/swagger.go
GOPATH_SWAGGER_CMD_SRC := $(GOPATH)/src/$(SWAGGER_CMD_SRC)
VENDORED_SWAGGER_CMD_SRC := vendor/$(SWAGGER_CMD_SRC)
SPEC_FILE := swagger-spec/redfish.yml
NAME := redfish
CLIENT := client
MODELS := models
CLIENT_BINDINGS := $(CLIENT)/redfish_client.go
GOCOV := $(GOPATH)/bin/gocov
GOVERALLS := $(GOPATH)/bin/goveralls
COVER := $(GOPATH)/bin/cover
COVER_OUT := cover.out

all: install

deps: $(CLIENT_BINDINGS) $(GO_GET)

$(GO_GET): | $(GLIDE_LOCK)
	go get -d $$($(GLIDE) nv) && touch $@

$(GLIDE_LOCK): $(GLIDE_YAML) | $(GLIDE) $(CLIENT_BINDINGS)
	$(GLIDE) --home $(HOME) up

$(GLIDE):
	go get -u github.com/Masterminds/glide

$(GOPATH_SWAGGER_CMD_SRC): $(VENDORED_SWAGGER_CMD_SRC)
$(VENDORED_SWAGGER_CMD_SRC): $(GLIDE_SWAGGER_YAML) | $(GLIDE)
	$(GLIDE) --home $(HOME) --yaml $< up --quick --cache-gopath --use-gopath

client-bindings: $(CLIENT_BINDINGS)
$(CLIENT_BINDINGS): $(VENDORED_SWAGGER_CMD_SRC) $(GOPATH_SWAGGER_CMD_SRC) $(SPEC_FILE)
	go run $(GOPATH_SWAGGER_CMD_SRC) generate client -f $(SPEC_FILE) -A $(NAME)

install: $(CLIENT_BINDINGS) $(GO_GET)
	go install -v $$($(GLIDE) nv)

test: $(COVER_OUT)
$(COVER_OUT):
	go test -v $$($(GLIDE) nv) -cover -coverprofile $@

cover: $(COVER_OUT) | $(GOCOV) $(GOVERALLS) $(COVER)
	$(GOVERALLS) -coverprofile=$<

clean:
	rm -fr $(GLIDE_LOCK) $(GO_GET) $(COVER_OUT) $(CLIENT) $(MODELS)

clobber: clean
	rm -fr vendor

clobber-gopath-swagger:
	rm -fr $(SWAGGER)
	rm -fr $(GOPATH)/src/$(SWAGGER_PKG)
	rm -fr $(GOPATH)/pkg/*/$(SWAGGER_PKG)

.PHONY: all deps install test cover \
		clean clobber clobber-gopath-swagger \
		client-bindings
