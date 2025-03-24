SHELL = /bin/bash

dev-compose-start:
	docker-compose --env-file .env.prod -f docker-compose.dev.yml up --build

prod-compose-start:
	docker pull kjw2262/witchs_lounge_backend:latest
	docker-compose --env-file .env.prod -f docker-compose.prod.yml up --build

manual-docker-hub-push:
	docker build -f Dockerfile.prod -t kjw2262/witchs_lounge_backend:latest .
	docker push kjw2262/witchs_lounge_backend:latest

update-image:
	docker pull kjw2262/witchs_lounge_backend:latest