run-db:
	docker-compose -f build/db-docker-compose.yml up -d --remove-orphans

run-tests:
	go test $$(go list ./... | grep -v /data/) -coverprofile=cover.out.tmp && cat ./cover.out.tmp | grep -v "mock.go" > ./cover.out && go tool cover -html=cover.out 

run-app:
	go run cmd/client/main.go

test-build-bake:
	docker build -t product-service . -f build/Dockerfile