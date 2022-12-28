.PHONY : all
all: build

run:
	go run ./api/v1/internal/cmd/main.go

run-production:
	export KRAIKUB_ENV=production && \
	make run

run-production-ps:
	. ./run-production.ps1

image:
	docker build -t kraikub/mail-service -f ./build/docker/Dockerfile .


