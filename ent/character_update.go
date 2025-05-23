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
	"github.com/witchs-lounge_backend/ent/character"
	"github.com/witchs-lounge_backend/ent/predicate"
	"github.com/witchs-lounge_backend/ent/product"
	"github.com/witchs-lounge_backend/ent/record"
)

// CharacterUpdate is the builder for updating Character entities.
type CharacterUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterMutation
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cu *CharacterUpdate) Where(ps ...predicate.Character) *CharacterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CharacterUpdate) SetUpdatedAt(t time.Time) *CharacterUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *CharacterUpdate) SetName(s string) *CharacterUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableName(s *string) *CharacterUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *CharacterUpdate) SetDescription(s string) *CharacterUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableDescription(s *string) *CharacterUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CharacterUpdate) ClearDescription() *CharacterUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetSource sets the "source" field.
func (cu *CharacterUpdate) SetSource(s string) *CharacterUpdate {
	cu.mutation.SetSource(s)
	return cu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableSource(s *string) *CharacterUpdate {
	if s != nil {
		cu.SetSource(*s)
	}
	return cu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (cu *CharacterUpdate) AddProductIDs(ids ...uuid.UUID) *CharacterUpdate {
	cu.mutation.AddProductIDs(ids...)
	return cu
}

// AddProducts adds the "products" edges to the Product entity.
func (cu *CharacterUpdate) AddProducts(p ...*Product) *CharacterUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProductIDs(ids...)
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (cu *CharacterUpdate) AddRecordIDs(ids ...uuid.UUID) *CharacterUpdate {
	cu.mutation.AddRecordIDs(ids...)
	return cu
}

// AddRecords adds the "records" edges to the Record entity.
func (cu *CharacterUpdate) AddRecords(r ...*Record) *CharacterUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddRecordIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cu *CharacterUpdate) Mutation() *CharacterMutation {
	return cu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (cu *CharacterUpdate) ClearProducts() *CharacterUpdate {
	cu.mutation.ClearProducts()
	return cu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (cu *CharacterUpdate) RemoveProductIDs(ids ...uuid.UUID) *CharacterUpdate {
	cu.mutation.RemoveProductIDs(ids...)
	return cu
}

// RemoveProducts removes "products" edges to Product entities.
func (cu *CharacterUpdate) RemoveProducts(p ...*Product) *CharacterUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProductIDs(ids...)
}

// ClearRecords clears all "records" edges to the Record entity.
func (cu *CharacterUpdate) ClearRecords() *CharacterUpdate {
	cu.mutation.ClearRecords()
	return cu
}

// RemoveRecordIDs removes the "records" edge to Record entities by IDs.
func (cu *CharacterUpdate) RemoveRecordIDs(ids ...uuid.UUID) *CharacterUpdate {
	cu.mutation.RemoveRecordIDs(ids...)
	return cu
}

// RemoveRecords removes "records" edges to Record entities.
func (cu *CharacterUpdate) RemoveRecords(r ...*Record) *CharacterUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveRecordIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CharacterUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CharacterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CharacterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CharacterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CharacterUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := character.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CharacterUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := character.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Character.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CharacterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(character.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(character.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(character.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Source(); ok {
		_spec.SetField(character.FieldSource, field.TypeString, value)
	}
	if cu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !cu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
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
	if nodes := cu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
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
	if cu.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRecordsIDs(); len(nodes) > 0 && !cu.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CharacterUpdateOne is the builder for updating a single Character entity.
type CharacterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CharacterUpdateOne) SetUpdatedAt(t time.Time) *CharacterUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CharacterUpdateOne) SetName(s string) *CharacterUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableName(s *string) *CharacterUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CharacterUpdateOne) SetDescription(s string) *CharacterUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableDescription(s *string) *CharacterUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CharacterUpdateOne) ClearDescription() *CharacterUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetSource sets the "source" field.
func (cuo *CharacterUpdateOne) SetSource(s string) *CharacterUpdateOne {
	cuo.mutation.SetSource(s)
	return cuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableSource(s *string) *CharacterUpdateOne {
	if s != nil {
		cuo.SetSource(*s)
	}
	return cuo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (cuo *CharacterUpdateOne) AddProductIDs(ids ...uuid.UUID) *CharacterUpdateOne {
	cuo.mutation.AddProductIDs(ids...)
	return cuo
}

// AddProducts adds the "products" edges to the Product entity.
func (cuo *CharacterUpdateOne) AddProducts(p ...*Product) *CharacterUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProductIDs(ids...)
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (cuo *CharacterUpdateOne) AddRecordIDs(ids ...uuid.UUID) *CharacterUpdateOne {
	cuo.mutation.AddRecordIDs(ids...)
	return cuo
}

// AddRecords adds the "records" edges to the Record entity.
func (cuo *CharacterUpdateOne) AddRecords(r ...*Record) *CharacterUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddRecordIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cuo *CharacterUpdateOne) Mutation() *CharacterMutation {
	return cuo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (cuo *CharacterUpdateOne) ClearProducts() *CharacterUpdateOne {
	cuo.mutation.ClearProducts()
	return cuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (cuo *CharacterUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *CharacterUpdateOne {
	cuo.mutation.RemoveProductIDs(ids...)
	return cuo
}

// RemoveProducts removes "products" edges to Product entities.
func (cuo *CharacterUpdateOne) RemoveProducts(p ...*Product) *CharacterUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProductIDs(ids...)
}

// ClearRecords clears all "records" edges to the Record entity.
func (cuo *CharacterUpdateOne) ClearRecords() *CharacterUpdateOne {
	cuo.mutation.ClearRecords()
	return cuo
}

// RemoveRecordIDs removes the "records" edge to Record entities by IDs.
func (cuo *CharacterUpdateOne) RemoveRecordIDs(ids ...uuid.UUID) *CharacterUpdateOne {
	cuo.mutation.RemoveRecordIDs(ids...)
	return cuo
}

// RemoveRecords removes "records" edges to Record entities.
func (cuo *CharacterUpdateOne) RemoveRecords(r ...*Record) *CharacterUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveRecordIDs(ids...)
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cuo *CharacterUpdateOne) Where(ps ...predicate.Character) *CharacterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CharacterUpdateOne) Select(field string, fields ...string) *CharacterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Character entity.
func (cuo *CharacterUpdateOne) Save(ctx context.Context) (*Character, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CharacterUpdateOne) SaveX(ctx context.Context) *Character {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CharacterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CharacterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CharacterUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := character.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CharacterUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := character.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Character.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CharacterUpdateOne) sqlSave(ctx context.Context) (_node *Character, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Character.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for _, f := range fields {
			if !character.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != character.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(character.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(character.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(character.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Source(); ok {
		_spec.SetField(character.FieldSource, field.TypeString, value)
	}
	if cuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !cuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
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
	if nodes := cuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.ProductsTable,
			Columns: []string{character.ProductsColumn},
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
	if cuo.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRecordsIDs(); len(nodes) > 0 && !cuo.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   character.RecordsTable,
			Columns: []string{character.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Character{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
