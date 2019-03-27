MYSQL_PORT_3306_TCP_ADDR ?= 127.0.0.1
MYSQL_PORT_3306_TCP_PORT ?= 3306
MYSQL_ADDRESS=${MYSQL_PORT_3306_TCP_ADDR}:${MYSQL_PORT_3306_TCP_PORT}

bootstrap:
	@echo "=== Preparing databases ==="
	MYSQL_ADDRESS=${MYSQL_ADDRESS} go run tests/bootstrap.go

docker-boostrap:
	@echo "=== Boostrapping MySQL container with data ==="
	docker exec -it routes-api go run /go/src/github.com/romeufcrosa/best-route-finder/tests/bootstrap.go
docker-build:
	@echo "=== Building docker image ==="
	docker build -t route-finder .

docker-local-run: local-env
	@echo "=== Starting Docker container ==="
	sleep 15
	@echo "MySQL should be ready... this needs to be improved"
	docker run --name routes-api -d route-finder
	@echo "Populating database with some data"
	docker exec -it routes-api go run /go/src/github.com/romeufcrosa/best-route-finder/tests/bootstrap.go

docker-test:
	@echo "=== Running tests in Docker container ==="
	docker exec -it routes-api bash -c "cd /go/src/github.com/romeufcrosa/best-route-finder && go test -cover ./..."

cleanup:
	@echo "=== Flushing databases ==="
	MYSQL_ADDRESS=${MYSQL_ADDRESS} go run tests/cleanup.go

install:
	echo "=== Installing dependencies ==="
	dep ensure -v
	echo "Done"

local-env:
	@echo "=== Starting local services ==="
	docker-compose up -d

local-run: bootstrap
	@echo "=== Running application ==="
	ENV="tests" SQL_ADDRESS="${MYSQL_ADDRESS}"\
		go run cmd/api/main.go

stop-docker: stop-local-env
	@echo "=== Stopping docker containers ==="
	docker stop routes-api

stop-local-env:
	@echo "=== Stopping local services ==="
	docker-compose stop

test: cleanup
	@echo "=== Running tests ==="
	ENV=tests SQL_ADDRESS=${MYSQL_ADDRESS}\
		go test -cover ./...
