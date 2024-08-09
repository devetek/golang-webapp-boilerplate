################################################################################
##@ PLAYGROUND
################################################################################

run: .bin-validator ## Run service development
	@export ENV=development && go run cmd/webapp/*.go

run-hot: .bin-validator ## Run service development with hot reload
	@air
	
css-dev: ## Run css development
	@npm run watch
