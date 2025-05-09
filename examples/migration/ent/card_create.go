// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/payment"
	"entgo.io/ent/examples/migration/ent/user"
	"entgo.io/ent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (_c *CardCreate) SetType(v string) *CardCreate {
	_c.mutation.SetType(v)
	return _c
}

// SetNillableType sets the "type" field if the given value is not nil.
func (_c *CardCreate) SetNillableType(v *string) *CardCreate {
	if v != nil {
		_c.SetType(*v)
	}
	return _c
}

// SetNumberHash sets the "number_hash" field.
func (_c *CardCreate) SetNumberHash(v string) *CardCreate {
	_c.mutation.SetNumberHash(v)
	return _c
}

// SetCvvHash sets the "cvv_hash" field.
func (_c *CardCreate) SetCvvHash(v string) *CardCreate {
	_c.mutation.SetCvvHash(v)
	return _c
}

// SetExpiresAt sets the "expires_at" field.
func (_c *CardCreate) SetExpiresAt(v time.Time) *CardCreate {
	_c.mutation.SetExpiresAt(v)
	return _c
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (_c *CardCreate) SetNillableExpiresAt(v *time.Time) *CardCreate {
	if v != nil {
		_c.SetExpiresAt(*v)
	}
	return _c
}

// SetOwnerID sets the "owner_id" field.
func (_c *CardCreate) SetOwnerID(v int) *CardCreate {
	_c.mutation.SetOwnerID(v)
	return _c
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (_c *CardCreate) SetNillableOwnerID(v *int) *CardCreate {
	if v != nil {
		_c.SetOwnerID(*v)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *CardCreate) SetOwner(v *User) *CardCreate {
	return _c.SetOwnerID(v.ID)
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (_c *CardCreate) AddPaymentIDs(ids ...int) *CardCreate {
	_c.mutation.AddPaymentIDs(ids...)
	return _c
}

// AddPayments adds the "payments" edges to the Payment entity.
func (_c *CardCreate) AddPayments(v ...*Payment) *CardCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddPaymentIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (_c *CardCreate) Mutation() *CardMutation {
	return _c.mutation
}

// Save creates the Card in the database.
func (_c *CardCreate) Save(ctx context.Context) (*Card, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CardCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CardCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *CardCreate) defaults() {
	if _, ok := _c.mutation.GetType(); !ok {
		v := card.DefaultType
		_c.mutation.SetType(v)
	}
	if _, ok := _c.mutation.OwnerID(); !ok {
		v := card.DefaultOwnerID
		_c.mutation.SetOwnerID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *CardCreate) check() error {
	if _, ok := _c.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Card.type"`)}
	}
	if _, ok := _c.mutation.NumberHash(); !ok {
		return &ValidationError{Name: "number_hash", err: errors.New(`ent: missing required field "Card.number_hash"`)}
	}
	if _, ok := _c.mutation.CvvHash(); !ok {
		return &ValidationError{Name: "cvv_hash", err: errors.New(`ent: missing required field "Card.cvv_hash"`)}
	}
	if _, ok := _c.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Card.owner_id"`)}
	}
	if len(_c.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Card.owner"`)}
	}
	return nil
}

func (_c *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
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

func (_c *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(card.Table, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.GetType(); ok {
		_spec.SetField(card.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := _c.mutation.NumberHash(); ok {
		_spec.SetField(card.FieldNumberHash, field.TypeString, value)
		_node.NumberHash = value
	}
	if value, ok := _c.mutation.CvvHash(); ok {
		_spec.SetField(card.FieldCvvHash, field.TypeString, value)
		_node.CvvHash = value
	}
	if value, ok := _c.mutation.ExpiresAt(); ok {
		_spec.SetField(card.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if nodes := _c.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.PaymentsTable,
			Columns: []string{card.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	err      error
	builders []*CardCreate
}

// Save creates the Card entities in the database.
func (_c *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Card, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
func (_c *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CardCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
