package cmd

import (
	"context"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/scotty"
	"go.uber.org/zap"
)

// Context is a collection of objects that are passed into
// CLI commands to prevent duplicating loading config,
// connecting to the database, etc.
type Context struct {
	// Logger is a pre-configured logger that commands use
	// to log messages.
	Logger *zap.Logger

	// Config is a pre-loaded config from wherever scotty was
	// able to find a config.  This is typically ~/.config/scotty/config.yaml
	Config *scotty.Config

	// DB is a pre-connection database connection.
	DB *ent.Client

	// Context is a context.Context used to provide timeouts for
	// commands where that is relevant.
	Context context.Context
}

// CLI is the root of all of the CLI commands.  It branches into
// subcommands for interacting with the various subsystems of Scotty.
type CLI struct {
	DB      DB      `cmd:"db" help:"Interact with Scotty's database"`
	Discord Discord `cmd:"discord" help:"Interact with the Scotty Discord bot"`
}
