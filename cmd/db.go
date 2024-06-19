package cmd

import "go.uber.org/zap"

// DB is a collection of commands that operate on the database.
type DB struct {
	Migrate MigrateDB `cmd:"migrate" help:"Upgrade the database schema if any migrations are available."`
}

// MigrateDB is a command that upgrades the schema in the database.
type MigrateDB struct{}

// Run runs the migrate command.
func (m MigrateDB) Run(ctx *Context) error {
	err := ctx.DB.Schema.Create(ctx.Context)
	if err != nil {
		ctx.Logger.Error("failed to update database schema", zap.Error(err))
	} else {
		ctx.Logger.Info("successfully migrated schema")
	}

	return err
}
