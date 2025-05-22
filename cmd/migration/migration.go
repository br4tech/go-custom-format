package main

import (
	"github.com/br4tech/go-custom-format/config"
	"github.com/br4tech/go-custom-format/internal/adapter/postgres"
	"github.com/br4tech/go-custom-format/internal/core/port"
	"github.com/br4tech/go-custom-format/internal/model"
)

func main() {
	cfg := config.GetConfig()
	db := postgres.NewPostgresAdapter(&cfg)

	ExecuteMigrate(db)
}

func ExecuteMigrate(db port.IPostgresAdapter) {
	db.GetDb().Migrator().CreateTable(
		&model.Product{},
	)
}
