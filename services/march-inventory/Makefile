# Makefile

# Define variables
PRISMA_CMD = go run github.com/steebchen/prisma-client-go
dev:
	nodemon --exec "go run" cmd/app/server.go --signal SIGTERM

start:
	go run cmd/app/server.go

gen:
	go run github.com/99designs/gqlgen generate

prisma-push:
	go run github.com/steebchen/prisma-client-go db push --schema=./cmd/app/prisma/schema.prisma

prisma-gen:
	go run github.com/steebchen/prisma-client-go generate --schema=./cmd/app/prisma/schema.prisma

prisma-migrate:
	go run github.com/steebchen/prisma-client-go migrate dev --schema=./cmd/app/prisma/schema.prisma