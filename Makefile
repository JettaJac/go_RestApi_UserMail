#https://www.youtube.com/watch?v=LxJLuW5aUDQ
DB_MAIN = restapi_dev

build:
	go build cmd/main

#tests: build
#	cd tests && go test -v

run: 
	go run cmd/main.go

db: 
	@echo "Создание базы данных $(DB_MAIN)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_MAIN);"	


test: 
	cd internal/app && go test
# -v -race -timeout 30s ./ ...

clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := run
.PHONY: build test db run build