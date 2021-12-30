.PHONY: \
	docker-compose-build \
	docker-compose-up \
	docker-compose-down \
	apiserver-genswagger \
	apiserver-run \
	apiserver-build \
	worker-run

ARGS :=
ifdef DEV
	ARGS += "--dev"
endif

ifdef DEV_APPS
	ARGS += "--dev-apps"
endif

ifdef PURGE
	ARGS += "--purge"
endif

# Docker-compose
docker-compose-build:
	@$(CURDIR)/hack/docker-compose-build.sh $(ARGS)
docker-compose-up:
	@$(CURDIR)/hack/docker-compose-up.sh $(ARGS)
docker-compose-down:
	@$(CURDIR)/hack/docker-compose-down.sh $(ARGS)

# API Server
apiserver-genswagger:
	@$(CURDIR)/hack/gen-swagger.sh
apiserver-build:
	@$(CURDIR)/hack/go-build-apiserver.sh

# Worker
worker-build:
	@$(CURDIR)/hack/go-build-worker.sh
