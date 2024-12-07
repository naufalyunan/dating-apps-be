# List of services
SERVICES = date-service api-gateway logs-service payment-service users-service profiles-service

# Target to build each service by running make in its folder
.PHONY: all build_services docker_compose

compose: build_services docker_compose

# runs make file in each service
build_services:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		$(MAKE) -C $$service protoc; \
	done

# Target to run docker-compose
docker_compose:
	@echo "Running docker-compose..."
	docker compose up --build

cloud:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		$(MAKE) -C $$service cloud; \
	done

cloud_run:
	@for service in $(SERVICES); do \
		echo "deploy run $$service..."; \
		$(MAKE) -C $$service cloud_run; \
	done