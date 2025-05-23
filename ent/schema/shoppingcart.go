package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ShoppingCart holds the schema definition for the ShoppingCart entity.
// We will use ent to generate the Go struct, but gocql for DB operations.
type ShoppingCart struct {
	ent.Schema
}

// Fields of the ShoppingCart.
func (ShoppingCart) Fields() []ent.Field {
	return []ent.Field{
		// We use UserID as an ent field, but it will correspond
		// to the 'userid' primary key in Cassandra.
		field.String("user_id").
			Unique(). // Ensures it's treated as a key in ent's view
			Immutable().
			StorageKey("userid"), // Hint for mapping if we were using an ent driver
		field.Int("item_count").
			Default(0).
			StorageKey("item_count"),
		field.Time("last_update_timestamp").
			Default(time.Now).
			UpdateDefault(time.Now). // Ent can manage this for the struct
			StorageKey("last_update_timestamp"),
	}
}

// Edges of the ShoppingCart.
func (ShoppingCart) Edges() []ent.Edge {
	return nil
}
