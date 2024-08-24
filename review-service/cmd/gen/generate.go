//go:build geninject
// +build geninject

package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../internal/data/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	gormdb, _ := gorm.Open(rawsql.New(rawsql.Config{
		FilePath: []string{
			"../../../sql_tables/reviews.sql",
			"../../../sql_tables/replies.sql",
			"../../../sql_tables/appeals.sql",
		},
	}))
	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
