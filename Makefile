.PHONY: login test up migrate

up: ## Start local server
	docker compose up
login:  ## Login mysql server
	mysql -h 127.0.0.1 -u root ozaki --password=pass
local-migrate: ## Migrate query disignated by CONFIG path
	mysql -h 127.0.0.1 -u root ozaki -p < $(CONFIG)
test:  ## Execute local test
	go test ./...
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
