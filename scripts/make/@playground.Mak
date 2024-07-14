################################################################################
##@ PLAYGROUND
################################################################################

init: .bin-validator ## Init development runtime
	@touch sql_db
	@sleep 5
	@make migrate

run: ## Run service development
	@export ENV=development && go run cmd/webapp/*.go

css-dev: ## Run css development
	@npm run watch

run-hot: ## Run service development with hot reload
	@air

migrate: ## Migrate database "make migrate env=beta"
	@export ENV=$(env) && go run cmd/migrate/*.go