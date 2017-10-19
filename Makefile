build:
	docker-compose up --build -d

shell:
	@docker run --security-opt=seccomp:unconfined --name gokafka_shell -p 0.0.0.0:3005:3000 -v $(PWD)/src:/go/src -it gokafka_mygoapp:latest bash || docker start -i gokafka_shell

attach:
	@docker attach gokafka_shell ||:
