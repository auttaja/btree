package btree

import "github.com/andersfylling/snowflake"

type (
	TreeValue struct {
		key  snowflake.Snowflake // Assume snowflake?
		item interface{}
	}
)

type Tree interface {
	Insert(value *TreeValue)
	Find(key snowflake.Snowflake) interface{}
	Delete(key snowflake.Snowflake)
}
