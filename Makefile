.PHONY: migrate migrate_down
DB?=football
MODEL?=

migrate:
	@echo "Running migrations"
	goose -dir sql/schema postgres postgres://postgres:@localhost:5432/$(DB) up

migrate_down:
	@echo "Running down migrations"
	goose -dir sql/schema postgres postgres://postgres:@localhost:5432/$(DB) down

genmodel:
	@echo "Generating ent schema for model"
	go run -mod=mod entgo.io/ent/cmd/ent new $(MODEL)

gencode:
	@echo "Generating ent code"
	go generate ./ent