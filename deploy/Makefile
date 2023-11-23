all :
		go run ../cmd/base/main.go

db :
		docker-compose -f docker-compose.yml up -d --remove-orphans


test :
		go run ../cmd/base/main.go &
		go run ../cmd/generator/main.go