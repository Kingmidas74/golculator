run:
	docker-compose up --build --remove-orphans

start:
	docker-compose up --remove-orphans

clean:
	docker-compose rm --stop --force

test:
	go test ./...