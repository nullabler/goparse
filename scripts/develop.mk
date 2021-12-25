##@ Develop:

sh: ## start shell in backend
	docker-compose --env-file configs/.env -f build/docker-compose.yml exec parser sh
