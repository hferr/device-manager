# builds and runs the docker container
build-run:
	docker-compose up --build

# runs tests
run-test:
	go test ./... -race

# create new migration
new-migration:
	goose create -dir ./migrations $(f) sql

# generate documentation
gen-docs:
	swag init -g ./cmd/api/main.go
