.PHONY: build seed_wallets

build: 
	cd internal/client && go build -o ../../build/client . ; cd - > /dev/null

seed_wallets:
	@cd data && rm -rf * && cd - > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/alice -P "" -C "alice" > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/bob -P "" -C "bob" > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/charlie -P "" -C "charlie" > /dev/null
