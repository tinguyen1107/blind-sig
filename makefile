.PHONY: run

run-bank:
	@cd bank && go build -o ../cbank main.go
	@chmod +x cbank
	@./cbank $(filter-out $@,$(MAKECMDGOALS))

prepare:
	mkdir bankdb clientdb storedb
	echo "[]" > ./clientdb/data.json
	echo '[ {"user_id": "john", "balance": 100000000 }, { "user_id": "store", "balance": 900000000 } ]' > ./bankdb/balance.json

init-bank:
	@echo "[Bank]: Generate keypair"
	./cbank genKey
	@echo "[Bank]: Broadcase PublicKey"
	cp ./bankdb/public_key.pem ./clientdb/
	cp ./bankdb/public_key.pem ./storedb/

build:
	@cd bank && go build -o ../cbank main.go
	@cd client && go build -o ../cclient main.go
	@cd store && go build -o ../cstore main.go

test-flow: build init-bank
	./cclient genTicket 100000
