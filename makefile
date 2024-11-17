.PHONY: run_http_server build_http_server  run_built_http_server test_coverage generate_mocks

run_http_server:
	go run ./cmd/httpserver/

build_http_server:
	go build -o ./target/apphttpserver ./cmd/httpserver/

run_built_http_server:
	./target/apphttpserver

test:
	go test ./internal/usecase/... ./internal/handler/... ./internal/helper/... ./internal/middleware/...
test_coverage:
	mkdir -p ./target && touch ./target/cover.out && go test -coverprofile ./target/cover.out ./internal/usecase/... ./internal/handler/... ./internal/helper/... ./internal/middleware/...

generate_mocks:
	mockery --dir=./internal/logger --name=Logger --output=./mocks/
	mockery --dir=./internal/database --name=Transactor --output=./mocks/

	mockery --dir=./internal/repository --all --output=./mocks/repository

	mockery --dir=./internal/usecase --all --output=./mocks/usecase

	mockery --dir=./internal/util --all --output=./mocks/util
	mockery --dir=./internal/helper --all --output=./mocks/helper

