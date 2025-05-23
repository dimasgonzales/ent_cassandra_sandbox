// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent_cassandra_sandbox/ent/shoppingcart"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ShoppingCart is the model entity for the ShoppingCart schema.
type ShoppingCart struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ItemCount holds the value of the "item_count" field.
	ItemCount int `json:"item_count,omitempty"`
	// LastUpdateTimestamp holds the value of the "last_update_timestamp" field.
	LastUpdateTimestamp time.Time `json:"last_update_timestamp,omitempty"`
	selectValues        sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShoppingCart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shoppingcart.FieldID, shoppingcart.FieldItemCount:
			values[i] = new(sql.NullInt64)
		case shoppingcart.FieldUserID:
			values[i] = new(sql.NullString)
		case shoppingcart.FieldLastUpdateTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShoppingCart fields.
func (sc *ShoppingCart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shoppingcart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case shoppingcart.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sc.UserID = value.String
			}
		case shoppingcart.FieldItemCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field item_count", values[i])
			} else if value.Valid {
				sc.ItemCount = int(value.Int64)
			}
		case shoppingcart.FieldLastUpdateTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_update_timestamp", values[i])
			} else if value.Valid {
				sc.LastUpdateTimestamp = value.Time
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShoppingCart.
// This includes values selected through modifiers, order, etc.
func (sc *ShoppingCart) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// Update returns a builder for updating this ShoppingCart.
// Note that you need to call ShoppingCart.Unwrap() before calling this method if this ShoppingCart
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *ShoppingCart) Update() *ShoppingCartUpdateOne {
	return NewShoppingCartClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the ShoppingCart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *ShoppingCart) Unwrap() *ShoppingCart {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShoppingCart is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *ShoppingCart) String() string {
	var builder strings.Builder
	builder.WriteString("ShoppingCart(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("user_id=")
	builder.WriteString(sc.UserID)
	builder.WriteString(", ")
	builder.WriteString("item_count=")
	builder.WriteString(fmt.Sprintf("%v", sc.ItemCount))
	builder.WriteString(", ")
	builder.WriteString("last_update_timestamp=")
	builder.WriteString(sc.LastUpdateTimestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ShoppingCarts is a parsable slice of ShoppingCart.
type ShoppingCarts []*ShoppingCart
