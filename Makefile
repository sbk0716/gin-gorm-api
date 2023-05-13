up:
		docker-compose up

down:
		docker-compose down --rmi all

up/prod:
		docker-compose -f docker-compose.production.yaml up

down/prod:
		docker-compose -f docker-compose.production.yaml down --rmi all