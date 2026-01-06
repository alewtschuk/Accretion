.PHONY: build proto 

build: 
	cd internal/client && go build -o ../../build/client . ; cd - > /dev/null

proto:
	@./scripts/proto.sh mempool
	@./scripts/proto.sh validator