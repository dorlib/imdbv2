package imdbv2

import (
	"context"
	"imdbv2/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
