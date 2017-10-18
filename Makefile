build:
	docker-compose up --build -d

shell:
	docker run -v $PWD/src:/go/src -it gokafka_mygoapp:latest bash
