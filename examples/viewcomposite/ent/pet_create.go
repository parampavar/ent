// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/viewcomposite/ent/pet"
	"entgo.io/ent/schema/field"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (_c *PetCreate) SetName(v string) *PetCreate {
	_c.mutation.SetName(v)
	return _c
}

// Mutation returns the PetMutation object of the builder.
func (_c *PetCreate) Mutation() *PetMutation {
	return _c.mutation
}

// Save creates the Pet in the database.
func (_c *PetCreate) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *PetCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *PetCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *PetCreate) check() error {
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	return nil
}

func (_c *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	err      error
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (_c *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Pet, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *PetCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
