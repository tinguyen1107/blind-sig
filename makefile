.PHONY: run

run-bank:
	@cd bank && go build -o ../cbank main.go
	@chmod +x cbank
	@./cbank $(filter-out $@,$(MAKECMDGOALS))


init-bank:
	@echo "[Bank]: Generate keypair"
	./cbank genKey
	@echo "[Bank]: Broadcase PublicKey"
	cp ./bankdb/public_key.pem ./clientdb/
	cp ./bankdb/public_key.pem ./storedb/

