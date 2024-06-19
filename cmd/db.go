package cmd

type DBCommands struct {
	Migrate MigrateDBCommand `cmd:"migrate" help:"Migrate the database"`
}

type MigrateDBCommand struct{}

func (m MigrateDBCommand) Run(ctx *Context) error {
	return ctx.DB.Schema.Create(ctx.Context)
}
