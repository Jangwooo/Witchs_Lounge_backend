// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/character"
	"github.com/witchs-lounge_backend/ent/item"
	"github.com/witchs-lounge_backend/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	// Global custom UUID ID
	ID uuid.UUID `json:"id,omitempty"`
	// Created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Type holds the value of the "type" field.
	Type product.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges        ProductEdges `json:"edges"`
	item_id      *uuid.UUID
	character_id *uuid.UUID
	selectValues sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// Character holds the value of the character edge.
	Character *Character `json:"character,omitempty"`
	// Purchasers holds the value of the purchasers edge.
	Purchasers []*User `json:"purchasers,omitempty"`
	// UserPurchases holds the value of the user_purchases edge.
	UserPurchases []*UserPurchase `json:"user_purchases,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ItemOrErr() (*Item, error) {
	if e.Item != nil {
		return e.Item, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: item.Label}
	}
	return nil, &NotLoadedError{edge: "item"}
}

// CharacterOrErr returns the Character value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) CharacterOrErr() (*Character, error) {
	if e.Character != nil {
		return e.Character, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: character.Label}
	}
	return nil, &NotLoadedError{edge: "character"}
}

// PurchasersOrErr returns the Purchasers value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) PurchasersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Purchasers, nil
	}
	return nil, &NotLoadedError{edge: "purchasers"}
}

// UserPurchasesOrErr returns the UserPurchases value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) UserPurchasesOrErr() ([]*UserPurchase, error) {
	if e.loadedTypes[3] {
		return e.UserPurchases, nil
	}
	return nil, &NotLoadedError{edge: "user_purchases"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case product.FieldName, product.FieldDescription, product.FieldType:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case product.FieldID:
			values[i] = new(uuid.UUID)
		case product.ForeignKeys[0]: // item_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case product.ForeignKeys[1]: // character_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.Float64
			}
		case product.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pr.Type = product.Type(value.String)
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field item_id", values[i])
			} else if value.Valid {
				pr.item_id = new(uuid.UUID)
				*pr.item_id = *value.S.(*uuid.UUID)
			}
		case product.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field character_id", values[i])
			} else if value.Valid {
				pr.character_id = new(uuid.UUID)
				*pr.character_id = *value.S.(*uuid.UUID)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryItem queries the "item" edge of the Product entity.
func (pr *Product) QueryItem() *ItemQuery {
	return NewProductClient(pr.config).QueryItem(pr)
}

// QueryCharacter queries the "character" edge of the Product entity.
func (pr *Product) QueryCharacter() *CharacterQuery {
	return NewProductClient(pr.config).QueryCharacter(pr)
}

// QueryPurchasers queries the "purchasers" edge of the Product entity.
func (pr *Product) QueryPurchasers() *UserQuery {
	return NewProductClient(pr.config).QueryPurchasers(pr)
}

// QueryUserPurchases queries the "user_purchases" edge of the Product entity.
func (pr *Product) QueryUserPurchases() *UserPurchaseQuery {
	return NewProductClient(pr.config).QueryUserPurchases(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pr.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product
