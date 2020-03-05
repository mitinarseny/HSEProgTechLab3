.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -timeout 0 -bench .

.PHONY: report
report:
	docker-compose up
