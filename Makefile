.DEFAULT_GOAL := build

# Build app
build:
	docker build -t race-result .
.PHONY: build

# Dowload dependencies
dependencies:
	docker-compose run dependencies
.PHONY: dependencies

# Run app
run:
	docker-compose run app
.PHONY: run

#  Execute tests
test:
	docker-compose run tests
.PHONY: test

#  Generate coverage
coverage:
	docker-compose run coverage
	 go tool cover -html=c.out
.PHONY: coverage