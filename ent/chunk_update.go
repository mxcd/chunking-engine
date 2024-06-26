// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"github.com/mxcd/chunking-engine/ent/chunk"
	"github.com/mxcd/chunking-engine/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChunkUpdate is the builder for updating Chunk entities.
type ChunkUpdate struct {
	config
	hooks    []Hook
	mutation *ChunkMutation
}

// Where appends a list predicates to the ChunkUpdate builder.
func (cu *ChunkUpdate) Where(ps ...predicate.Chunk) *ChunkUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChunkUpdate) SetName(s string) *ChunkUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ChunkUpdate) SetNillableName(s *string) *ChunkUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetData sets the "data" field.
func (cu *ChunkUpdate) SetData(b []byte) *ChunkUpdate {
	cu.mutation.SetData(b)
	return cu
}

// Mutation returns the ChunkMutation object of the builder.
func (cu *ChunkUpdate) Mutation() *ChunkMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChunkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChunkUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChunkUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChunkUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChunkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chunk.Table, chunk.Columns, sqlgraph.NewFieldSpec(chunk.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(chunk.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Data(); ok {
		_spec.SetField(chunk.FieldData, field.TypeBytes, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chunk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChunkUpdateOne is the builder for updating a single Chunk entity.
type ChunkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChunkMutation
}

// SetName sets the "name" field.
func (cuo *ChunkUpdateOne) SetName(s string) *ChunkUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ChunkUpdateOne) SetNillableName(s *string) *ChunkUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetData sets the "data" field.
func (cuo *ChunkUpdateOne) SetData(b []byte) *ChunkUpdateOne {
	cuo.mutation.SetData(b)
	return cuo
}

// Mutation returns the ChunkMutation object of the builder.
func (cuo *ChunkUpdateOne) Mutation() *ChunkMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ChunkUpdate builder.
func (cuo *ChunkUpdateOne) Where(ps ...predicate.Chunk) *ChunkUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChunkUpdateOne) Select(field string, fields ...string) *ChunkUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chunk entity.
func (cuo *ChunkUpdateOne) Save(ctx context.Context) (*Chunk, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChunkUpdateOne) SaveX(ctx context.Context) *Chunk {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChunkUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChunkUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChunkUpdateOne) sqlSave(ctx context.Context) (_node *Chunk, err error) {
	_spec := sqlgraph.NewUpdateSpec(chunk.Table, chunk.Columns, sqlgraph.NewFieldSpec(chunk.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chunk.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chunk.FieldID)
		for _, f := range fields {
			if !chunk.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chunk.FieldID {
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
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(chunk.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Data(); ok {
		_spec.SetField(chunk.FieldData, field.TypeBytes, value)
	}
	_node = &Chunk{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chunk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
