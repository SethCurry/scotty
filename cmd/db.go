package cmd

type DB struct {
	Migrate MigrateDB `cmd:"migrate" help:"Migrate the database"`
}

type MigrateDB struct{}

func (m MigrateDB) Run(ctx *Context) error {
	return ctx.DB.Schema.Create(ctx.Context)
}
