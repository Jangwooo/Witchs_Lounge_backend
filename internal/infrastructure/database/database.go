package database

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/witchs-lounge_backend/ent"
)

func NewEntClient(config *Config) (*ent.Client, error) {
	drv, err := sql.Open("postgres", config.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}
	defer drv.Close()

	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
