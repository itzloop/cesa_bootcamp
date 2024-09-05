.SILENT:

ENV_FILE=.env
include $(ENV_FILE)

DB_SCRIPTS_PATH="scripts/db"
# DB_CREATE_FILE="$(DB_SCRIPTS_PATH)/create_db.sql"
COMPOSE_FILE_PATH="docker-compose.yml"
DB_MIGRATIONS_PATH="$(DB_SCRIPTS_PATH)/migrations"

PG_CONN_STR="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)"
DB_CONN_STR="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

REMOTE_IMAGE_TAG=$(DOCKER_HUB_USERNAME)/$(IMAGE_TAG):0.0.1

pg-up:
	@docker compose -f $(COMPOSE_FILE_PATH) --env-file $(ENV_FILE) up -d

pg-down:
	@docker compose -f $(COMPOSE_FILE_PATH) --env-file $(ENV_FILE) down

pg-downv:
	@docker compose -f $(COMPOSE_FILE_PATH) --env-file $(ENV_FILE) down -v

# db-init:
# 	@psql $(PG_CONN_STR) -f $(DB_CREATE_FILE)

mig-create:
	@migrate create -ext sql -dir $(DB_MIGRATIONS_PATH) -seq -digits 8 $f

mig-up:
	@migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_CONN_STR) up $c

mig-down:
	@migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_CONN_STR) down $c $a

mig-force:
	@migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_CONN_STR) force $v

login:
	@docker login -u $(DOCKER_HUB_USERNAME)

build-local:
	@docker build -t $(IMAGE_TAG) .

build-remote:
	@docker build -t $(REMOTE_IMAGE_TAG) .

push-remote:
	@docker push $(REMOTE_IMAGE_TAG)

pull-remote:
	@docker pull $(REMOTE_IMAGE_TAG)

# compose-up:
# 	@docker compose --env-file $(ENV_FILE) up postgres server

# compose-down:
# 	@docker compose --env-file $(ENV_FILE) down

# compose-down-v:
# 	@docker compose --env-file $(ENV_FILE) down --volumes