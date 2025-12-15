.PHONY: build seed_wallets

build: 
	go build -o build/client ./internal/client/

seed_wallets:
	@cd data && rm -rf * && cd - > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/alice -P "" > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/bob -P "" > /dev/null
	@ssh-keygen -t ed25519 -m PEM -f data/charlie -P "" > /dev/null
