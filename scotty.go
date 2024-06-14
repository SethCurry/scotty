package main

import (
	"context"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/scotty"
	"github.com/alecthomas/kong"
	"go.uber.org/zap"
)

type MigrateDBCommand struct{}

func (m MigrateDBCommand) Run(ctx *Context) error {
	return ctx.DB.Schema.Create(ctx.Context)
}

type DBCommands struct {
	Migrate MigrateDBCommand `cmd:"migrate" help:"Migrate the database"`
}

type StartCommand struct{}

type CLI struct {
	DB    DBCommands   `cmd:"db" help:"Interact with Scotty's database"`
	Start StartCommand `cmd:"start" help:"Start the bot"`
}

type Context struct {
	Logger  *zap.Logger
	Config  *scotty.Config
	DB      *ent.Client
	Context context.Context
}

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

	cli := &CLI{}

	ctx := kong.Parse(cli)

	err = ctx.Run(&Context{Logger: logger, Config: cfg, DB: dbConn, Context: context.Background()})
	if err != nil {
		logger.Fatal("failed to run", zap.Error(err))
	}
}
