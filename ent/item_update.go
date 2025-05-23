// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/item"
	"github.com/witchs-lounge_backend/ent/predicate"
	"github.com/witchs-lounge_backend/ent/product"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ItemUpdate) SetUpdatedAt(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetName sets the "name" field.
func (iu *ItemUpdate) SetName(s string) *ItemUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableName(s *string) *ItemUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *ItemUpdate) SetDescription(s string) *ItemUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableDescription(s *string) *ItemUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// ClearDescription clears the value of the "description" field.
func (iu *ItemUpdate) ClearDescription() *ItemUpdate {
	iu.mutation.ClearDescription()
	return iu
}

// SetEffectID sets the "effect_id" field.
func (iu *ItemUpdate) SetEffectID(s string) *ItemUpdate {
	iu.mutation.SetEffectID(s)
	return iu
}

// SetNillableEffectID sets the "effect_id" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableEffectID(s *string) *ItemUpdate {
	if s != nil {
		iu.SetEffectID(*s)
	}
	return iu
}

// ClearEffectID clears the value of the "effect_id" field.
func (iu *ItemUpdate) ClearEffectID() *ItemUpdate {
	iu.mutation.ClearEffectID()
	return iu
}

// SetType sets the "type" field.
func (iu *ItemUpdate) SetType(i item.Type) *ItemUpdate {
	iu.mutation.SetType(i)
	return iu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableType(i *item.Type) *ItemUpdate {
	if i != nil {
		iu.SetType(*i)
	}
	return iu
}

// SetSource sets the "source" field.
func (iu *ItemUpdate) SetSource(s string) *ItemUpdate {
	iu.mutation.SetSource(s)
	return iu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableSource(s *string) *ItemUpdate {
	if s != nil {
		iu.SetSource(*s)
	}
	return iu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (iu *ItemUpdate) AddProductIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.AddProductIDs(ids...)
	return iu
}

// AddProducts adds the "products" edges to the Product entity.
func (iu *ItemUpdate) AddProducts(p ...*Product) *ItemUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.AddProductIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (iu *ItemUpdate) ClearProducts() *ItemUpdate {
	iu.mutation.ClearProducts()
	return iu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (iu *ItemUpdate) RemoveProductIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.RemoveProductIDs(ids...)
	return iu
}

// RemoveProducts removes "products" edges to Product entities.
func (iu *ItemUpdate) RemoveProducts(p ...*Product) *ItemUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := item.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Item.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.GetType(); ok {
		if err := item.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Item.type": %w`, err)}
		}
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(item.FieldDescription, field.TypeString, value)
	}
	if iu.mutation.DescriptionCleared() {
		_spec.ClearField(item.FieldDescription, field.TypeString)
	}
	if value, ok := iu.mutation.EffectID(); ok {
		_spec.SetField(item.FieldEffectID, field.TypeString, value)
	}
	if iu.mutation.EffectIDCleared() {
		_spec.ClearField(item.FieldEffectID, field.TypeString)
	}
	if value, ok := iu.mutation.GetType(); ok {
		_spec.SetField(item.FieldType, field.TypeEnum, value)
	}
	if value, ok := iu.mutation.Source(); ok {
		_spec.SetField(item.FieldSource, field.TypeString, value)
	}
	if iu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !iu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ItemUpdateOne) SetUpdatedAt(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetName sets the "name" field.
func (iuo *ItemUpdateOne) SetName(s string) *ItemUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableName(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *ItemUpdateOne) SetDescription(s string) *ItemUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableDescription(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// ClearDescription clears the value of the "description" field.
func (iuo *ItemUpdateOne) ClearDescription() *ItemUpdateOne {
	iuo.mutation.ClearDescription()
	return iuo
}

// SetEffectID sets the "effect_id" field.
func (iuo *ItemUpdateOne) SetEffectID(s string) *ItemUpdateOne {
	iuo.mutation.SetEffectID(s)
	return iuo
}

// SetNillableEffectID sets the "effect_id" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableEffectID(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetEffectID(*s)
	}
	return iuo
}

// ClearEffectID clears the value of the "effect_id" field.
func (iuo *ItemUpdateOne) ClearEffectID() *ItemUpdateOne {
	iuo.mutation.ClearEffectID()
	return iuo
}

// SetType sets the "type" field.
func (iuo *ItemUpdateOne) SetType(i item.Type) *ItemUpdateOne {
	iuo.mutation.SetType(i)
	return iuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableType(i *item.Type) *ItemUpdateOne {
	if i != nil {
		iuo.SetType(*i)
	}
	return iuo
}

// SetSource sets the "source" field.
func (iuo *ItemUpdateOne) SetSource(s string) *ItemUpdateOne {
	iuo.mutation.SetSource(s)
	return iuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableSource(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetSource(*s)
	}
	return iuo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (iuo *ItemUpdateOne) AddProductIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.AddProductIDs(ids...)
	return iuo
}

// AddProducts adds the "products" edges to the Product entity.
func (iuo *ItemUpdateOne) AddProducts(p ...*Product) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.AddProductIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (iuo *ItemUpdateOne) ClearProducts() *ItemUpdateOne {
	iuo.mutation.ClearProducts()
	return iuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (iuo *ItemUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.RemoveProductIDs(ids...)
	return iuo
}

// RemoveProducts removes "products" edges to Product entities.
func (iuo *ItemUpdateOne) RemoveProducts(p ...*Product) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := item.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Item.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.GetType(); ok {
		if err := item.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Item.type": %w`, err)}
		}
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(item.FieldDescription, field.TypeString, value)
	}
	if iuo.mutation.DescriptionCleared() {
		_spec.ClearField(item.FieldDescription, field.TypeString)
	}
	if value, ok := iuo.mutation.EffectID(); ok {
		_spec.SetField(item.FieldEffectID, field.TypeString, value)
	}
	if iuo.mutation.EffectIDCleared() {
		_spec.ClearField(item.FieldEffectID, field.TypeString)
	}
	if value, ok := iuo.mutation.GetType(); ok {
		_spec.SetField(item.FieldType, field.TypeEnum, value)
	}
	if value, ok := iuo.mutation.Source(); ok {
		_spec.SetField(item.FieldSource, field.TypeString, value)
	}
	if iuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !iuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ProductsTable,
			Columns: []string{item.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
