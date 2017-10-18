build:
	docker-compose up --build -d

shell:
	docker run --name gokafka_shell -p 0.0.0.0:3005:3000 -v $(PWD)/src:/go/src -it gokafka_mygoapp:latest bash

attach:
	@docker attach gokafka_shell ||:

start:
	docker run --name gokafka_shell -p 0.0.0.0:3005:3000 -v $(PWD)/src:/go/src -it gokafka_mygoapp:latest /go/src/start_server.sh
