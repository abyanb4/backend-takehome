DOCKER_COMPOSE = docker-compose
DOCKER_BUILD = $(DOCKER_COMPOSE) build
DOCKER_UP = $(DOCKER_COMPOSE) up

DOCKER_BUILD_PROD = docker-compose -f docker-compose.prod.yml build
DOCKER_UP_PROD = docker-compose -f docker-compose.prod.yml up
DOCKER_DOWN_PROD = docker-compose -f docker-compose.prod.yml down

.PHONY: run
run: build
	@echo "Starting the container..."
	$(DOCKER_UP)

# Build the containers
.PHONY: build
build: 
	@echo "Building the container..."
	$(DOCKER_BUILD)

.PHONY: run-prod
run-prod: build-prod
	@echo "Starting the container for production..."
	$(DOCKER_UP_PROD)

.PHONY: down-prod
down-prod: 
	@echo "Stop the container for production..."
	$(DOCKER_DOWN_PROD)

.PHONY: build-prod
build-prod: 
	@echo "Building the container for production..."
	$(DOCKER_BUILD_PROD)
    


# migrate-up:
#     migrate -path db/migrations -database "$(DATABASE_URL)" up

# migrate-down:
#     migrate -path db/migrations -database "$(DATABASE_URL)" down
