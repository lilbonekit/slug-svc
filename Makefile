# -----------------------------
# Postgres local dev environment
# -----------------------------

# Config
PG_HOST       ?= 127.0.0.1
PG_CONTAINER  ?= slug-pg
PG_VOLUME     ?= slug-pgdata
PG_USER       ?= slug
PG_PASSWORD   ?= slug
PG_DB         ?= slugsvc
PG_PORT       ?= 55432

# Service config (for go run with local config)
KV_FILE       ?= ./config.yaml

# -----------------------------
# Helpers
# -----------------------------

define wait_for_pg
	@echo "⏳ Waiting for Postgres to be ready at $(PG_HOST):$(PG_PORT) db=$(PG_DB) user=$(PG_USER)..."
	@for i in $$(seq 1 30); do \
		PGPASSWORD=$(PG_PASSWORD) psql \
			-h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -d $(PG_DB) \
			-c "select 1" >/dev/null 2>&1 && { echo "✅ Postgres is ready."; exit 0; }; \
		sleep 1; \
	done; \
	echo "❌ Postgres did not become ready in time."; \
	PGPASSWORD=$(PG_PASSWORD) psql -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -d $(PG_DB) -c "select 1" || true; \
	exit 1
endef

# -----------------------------
# Commands
# -----------------------------

## Full reset: stop + remove container and volume
db-reset:
	- docker rm -f $(PG_CONTAINER) >/dev/null 2>&1 || true
	- docker volume rm $(PG_VOLUME) >/dev/null 2>&1 || true
	@echo "✅ Reset done."

## Start Postgres container
db-up:
	@echo "🚀 Starting PostgreSQL container..."
	docker run -d --name $(PG_CONTAINER) \
		-e POSTGRES_USER=$(PG_USER) \
		-e POSTGRES_PASSWORD=$(PG_PASSWORD) \
		-e POSTGRES_DB=$(PG_DB) \
		-p $(PG_PORT):5432 \
		-v $(PG_VOLUME):/var/lib/postgresql/data \
		postgres:16 >/dev/null
	@echo "✅ Postgres is running on port $(PG_PORT)"
	$(call wait_for_pg)

## Check connection
db-check:
	@echo "🔍 Checking connection..."
	PGPASSWORD=$(PG_PASSWORD) psql \
		-h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -d $(PG_DB) \
		-c "select now();"

## Container logs
db-logs:
	docker logs -f $(PG_CONTAINER)

## Open interactive psql session
db-psql:
	PGPASSWORD=$(PG_PASSWORD) psql -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -d $(PG_DB)

## Stop and remove container (without removing volume)
db-down:
	@echo "🧹 Removing Postgres container..."
	- docker rm -f $(PG_CONTAINER) >/dev/null 2>&1 || true

## Apply migrations
migrate-up:
	@echo "📦 Running migrations..."
	KV_VIPER_FILE=$(KV_FILE) go run . migrate up

## Rollback migrations
migrate-down:
	@echo "⏪ Rolling back migrations..."
	KV_VIPER_FILE=$(KV_FILE) go run . migrate down

## Start service
run:
	@echo "⚙️  Starting service..."
	KV_VIPER_FILE=$(KV_FILE) go run . run service
