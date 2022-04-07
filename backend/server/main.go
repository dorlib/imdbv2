package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"imdbv2/ent"
	"imdbv2/ent/migrate"
	"imdbv2/graphql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var cli struct {
		Addr  string `name:"address" defualt:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}

	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()
	ctx := context.Background()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	log.Println("listening on", cli.Addr)
	if err := http.ListenAndServe(cli.Addr, nil); err != nil {
		log.Fatalf("error running server (%s)", err)
	}
}
