// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/product"
	"github.com/witchs-lounge_backend/ent/user"
	"github.com/witchs-lounge_backend/ent/userpurchase"
)

// UserPurchase is the model entity for the UserPurchase schema.
type UserPurchase struct {
	config `json:"-"`
	// ID of the ent.
	// Global custom UUID ID
	ID uuid.UUID `json:"id,omitempty"`
	// Created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID uuid.UUID `json:"product_id,omitempty"`
	// PurchaseDate holds the value of the "purchase_date" field.
	PurchaseDate time.Time `json:"purchase_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserPurchaseQuery when eager-loading is set.
	Edges        UserPurchaseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserPurchaseEdges holds the relations/edges for other nodes in the graph.
type UserPurchaseEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPurchaseEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPurchaseEdges) ProductOrErr() (*Product, error) {
	if e.Product != nil {
		return e.Product, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: product.Label}
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPurchase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpurchase.FieldCreatedAt, userpurchase.FieldUpdatedAt, userpurchase.FieldPurchaseDate:
			values[i] = new(sql.NullTime)
		case userpurchase.FieldID, userpurchase.FieldUserID, userpurchase.FieldProductID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPurchase fields.
func (up *UserPurchase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpurchase.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				up.ID = *value
			}
		case userpurchase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				up.CreatedAt = value.Time
			}
		case userpurchase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				up.UpdatedAt = value.Time
			}
		case userpurchase.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				up.UserID = *value
			}
		case userpurchase.FieldProductID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value != nil {
				up.ProductID = *value
			}
		case userpurchase.FieldPurchaseDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_date", values[i])
			} else if value.Valid {
				up.PurchaseDate = value.Time
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserPurchase.
// This includes values selected through modifiers, order, etc.
func (up *UserPurchase) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserPurchase entity.
func (up *UserPurchase) QueryUser() *UserQuery {
	return NewUserPurchaseClient(up.config).QueryUser(up)
}

// QueryProduct queries the "product" edge of the UserPurchase entity.
func (up *UserPurchase) QueryProduct() *ProductQuery {
	return NewUserPurchaseClient(up.config).QueryProduct(up)
}

// Update returns a builder for updating this UserPurchase.
// Note that you need to call UserPurchase.Unwrap() before calling this method if this UserPurchase
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPurchase) Update() *UserPurchaseUpdateOne {
	return NewUserPurchaseClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserPurchase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPurchase) Unwrap() *UserPurchase {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserPurchase is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPurchase) String() string {
	var builder strings.Builder
	builder.WriteString("UserPurchase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("created_at=")
	builder.WriteString(up.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(up.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", up.UserID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", up.ProductID))
	builder.WriteString(", ")
	builder.WriteString("purchase_date=")
	builder.WriteString(up.PurchaseDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserPurchases is a parsable slice of UserPurchase.
type UserPurchases []*UserPurchase
