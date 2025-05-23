// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
)

// GroupTag is the model entity for the GroupTag schema.
type GroupTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID int `json:"tag_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupTagQuery when eager-loading is set.
	Edges        GroupTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupTagEdges holds the relations/edges for other nodes in the graph.
type GroupTagEdges struct {
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupTagEdges) TagOrErr() (*Tag, error) {
	if e.Tag != nil {
		return e.Tag, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tag.Label}
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupTagEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case grouptag.FieldID, grouptag.FieldTagID, grouptag.FieldGroupID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupTag fields.
func (_m *GroupTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case grouptag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case grouptag.FieldTagID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				_m.TagID = int(value.Int64)
			}
		case grouptag.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				_m.GroupID = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupTag.
// This includes values selected through modifiers, order, etc.
func (_m *GroupTag) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryTag queries the "tag" edge of the GroupTag entity.
func (_m *GroupTag) QueryTag() *TagQuery {
	return NewGroupTagClient(_m.config).QueryTag(_m)
}

// QueryGroup queries the "group" edge of the GroupTag entity.
func (_m *GroupTag) QueryGroup() *GroupQuery {
	return NewGroupTagClient(_m.config).QueryGroup(_m)
}

// Update returns a builder for updating this GroupTag.
// Note that you need to call GroupTag.Unwrap() before calling this method if this GroupTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *GroupTag) Update() *GroupTagUpdateOne {
	return NewGroupTagClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the GroupTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *GroupTag) Unwrap() *GroupTag {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupTag is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *GroupTag) String() string {
	var builder strings.Builder
	builder.WriteString("GroupTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("tag_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.TagID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.GroupID))
	builder.WriteByte(')')
	return builder.String()
}

// GroupTags is a parsable slice of GroupTag.
type GroupTags []*GroupTag
