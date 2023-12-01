all :
		go run ../cmd/base/main.go

docker :
		docker-compose -f docker-compose.yml up -d --remove-orphans

clean_docker:
	docker stop WildBerries && docker rm WildBerries
	docker stop NatsStreaming && docker rm NatsStreaming

generator :
		go run ../cmd/generator/main.go