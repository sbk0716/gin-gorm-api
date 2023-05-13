up:
		docker-compose up

down:
		docker-compose down --rmi all

up/prod:
		docker-compose -f docker-compose.production.yml up

down/prod:
		docker-compose -f docker-compose.production.yml down --rmi all