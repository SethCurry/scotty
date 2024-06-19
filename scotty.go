package main

import (
	"context"

	"github.com/SethCurry/scotty/cmd"
	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/scotty"
	"github.com/alecthomas/kong"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	cfg, err := scotty.LoadDefaultConfig()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	dbConn, err := ent.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		logger.Fatal("failed to open database", zap.Error(err))
	}

	cli := &cmd.CLI{}

	ctx := kong.Parse(cli)

	err = ctx.Run(&cmd.Context{Logger: logger, Config: cfg, DB: dbConn, Context: context.Background()})
	if err != nil {
		logger.Fatal("failed to run", zap.Error(err))
	}
}
