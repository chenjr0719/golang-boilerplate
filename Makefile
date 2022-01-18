.PHONY: \
	docker-compose-build \
	docker-compose-up \
	docker-compose-down \
	docker-compose-purge \
	docker-compose-dev-up \
	docker-compose-dev-down \
	docker-compose-dev-purge \
	apiserver-genswagger \
	apiserver-run \
	apiserver-build \
	worker-run \
	test

# Docker-compose
docker-compose-build:
	@$(CURDIR)/hack/docker-compose-build.sh
docker-compose-up:
	@$(CURDIR)/hack/docker-compose-up.sh
docker-compose-down:
	@$(CURDIR)/hack/docker-compose-down.sh
docker-compose-purge:
	@$(CURDIR)/hack/docker-compose-down.sh --purge
docker-compose-dev-up:
	@$(CURDIR)/hack/docker-compose-up.sh --dev
docker-compose-debug-up:
	@$(CURDIR)/hack/docker-compose-up.sh --dev --debug
docker-compose-dev-down:
	@$(CURDIR)/hack/docker-compose-down.sh --dev
docker-compose-dev-purge:
	@$(CURDIR)/hack/docker-compose-down.sh --dev --purge

# API Server
apiserver-genswagger:
	@$(CURDIR)/hack/gen-swagger.sh
apiserver-build:
	@$(CURDIR)/hack/go-build-apiserver.sh

# Worker
worker-build:
	@$(CURDIR)/hack/go-build-worker.sh

test:
	@$(CURDIR)/hack/go-test.sh
