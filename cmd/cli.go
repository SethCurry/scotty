package cmd

import (
	"context"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/scotty"
	"go.uber.org/zap"
)

type Context struct {
	Logger  *zap.Logger
	Config  *scotty.Config
	DB      *ent.Client
	Context context.Context
}

type CLI struct {
	DB      DBCommands `cmd:"db" help:"Interact with Scotty's database"`
	Discord Discord    `cmd:"discord" help:"Interact with the Scotty Discord bot"`
}
