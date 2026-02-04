package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/WadeGulbrandsen/gator/internal/config"
	"github.com/WadeGulbrandsen/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v", err)
		os.Exit(0)
	}
	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		os.Exit(0)
	}
	dbQueries := database.New(db)
	s := state{db: dbQueries, cfg: &cfg}
	commands := commands{cmds: map[string]func(*state, command) error{}}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not command given")
		os.Exit(1)
	}
	cmd := command{name: args[1], args: args[2:]}
	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
