# This makefile was inspired by https://github.com/containous/traefik/blob/master/Makefile

SERVER_BIN=rok4go_server
BUILD_OPTS=--tags dist
COVERFILE=${DIST_DIR}/cover.out
TEST_OPTS= -coverprofile=${COVERFILE}

DIST_DIR=dist
RESOURCES_DIR=resources
UI_DIR=${RESOURCES_DIR}/ui

all: server

install-dependencies: vendor ${UI_DIR}/node_modules

vendor:
	dep ensure

server: dist server-ui server-bin ## build only server

server-bin: install-dependencies
	go generate ./... && \
	go build ${BUILD_OPTS} -o $(DIST_DIR)/$(SERVER_BIN) ./cli/server.go

server-bin-dev: install-dependencies
	go build -o $(DIST_DIR)/$(SERVER_BIN) ./cli/server.go
server-bin-watch: install-dependencies ## watch, build and run server, using raw assets
	nodemon -e go -x "sh -c" "make server-bin-dev && ./$(DIST_DIR)/$(SERVER_BIN) --mongodb mongo:27017"

server-ui: ${UI_DIR}/node_modules
	cd ${UI_DIR} && \
	yarn run webpack
server-ui-watch: ${UI_DIR}/node_modules ## run webpack in watching mode for ui development
	cd ${UI_DIR} && \
	yarn run webpack-dev


${UI_DIR}/node_modules:
	cd ${UI_DIR} && \
	yarn

unittestcover: unittest ## show unit tests coverage
	go tool cover -func=${COVERFILE}

test: unittest ## run tests

unittest: dist install-dependencies
	go test --tags test $(TEST_OPTS) ./...

dist:
	mkdir -p $(DIST_DIR)

clean: ## cleanup generated files
	rm -rf $(DIST_DIR) && \
	rm -f ${RESOURCES_DIR}/assets_vfsdata.go && \
	rm -rf ${UI_DIR}/dist
	

clean-full: clean ## cleanup generated files + node_modules and vendor directory
	rm -r ${UI_DIR}/node_modules && \
	rm -r vendor

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

%-docker:
	docker-compose exec bin make $*