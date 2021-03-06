// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/dmalykh/tagservice/repository/entgo/ent/namespace"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Namespace is the model entity for the Namespace schema.
type Namespace struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Namespace) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case namespace.FieldID:
			values[i] = new(sql.NullInt64)
		case namespace.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Namespace", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Namespace fields.
func (n *Namespace) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case namespace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case namespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Namespace.
// Note that you need to call Namespace.Unwrap() before calling this method if this Namespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Namespace) Update() *NamespaceUpdateOne {
	return (&NamespaceClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Namespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Namespace) Unwrap() *Namespace {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Namespace is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Namespace) String() string {
	var builder strings.Builder
	builder.WriteString("Namespace(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", name=")
	builder.WriteString(n.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Namespaces is a parsable slice of Namespace.
type Namespaces []*Namespace

func (n Namespaces) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
