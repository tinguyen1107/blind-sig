.PHONY: run

run-bank:
	@cd bank && go build -o ../cbank main.go
	@chmod +x cbank
	@./cbank $(filter-out $@,$(MAKECMDGOALS))

