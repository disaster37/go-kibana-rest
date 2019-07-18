SUDO_DOCKER ?=

all: help

help:
	@echo "make test - run tests"
	@echo "make clean - stop and remove test containers"
	@echo "make pull - pull Docker images on registry"
	@echo "make init - lauch ambari infra for test purpose"
	@echo "make build - compile program"

init:
	${SUDO_DOCKER} docker-compose up -d kibana
	until $$(docker-compose run --rm curl --output /dev/null --silent --head --fail -u elastic:changeme http://kibana:5601); do sleep 5; done

test: clean init
	${SUDO_DOCKER} docker-compose run --rm test


build:
	${SUDO_DOCKER} docker-compose run --rm build
	
pull:
	${SUDO_DOCKER} docker-compose pull --ignore-pull-failures

clean:
	${SUDO_DOCKER} docker-compose logs elasticsearch || exit 0
	${SUDO_DOCKER} docker-compose logs kibana || exit 0
	${SUDO_DOCKER} docker-compose down -v

release:
	@echo "Do nothink"
