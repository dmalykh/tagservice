// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/dmalykh/tagservice/repository/entgo/ent/category"
	"github.com/dmalykh/tagservice/repository/entgo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetTitle sets the "title" field.
func (cu *CategoryUpdate) SetTitle(s string) *CategoryUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableTitle(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// ClearTitle clears the value of the "title" field.
func (cu *CategoryUpdate) ClearTitle() *CategoryUpdate {
	cu.mutation.ClearTitle()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDescription(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CategoryUpdate) ClearDescription() *CategoryUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(i int) *CategoryUpdate {
	cu.mutation.ResetParentID()
	cu.mutation.SetParentID(i)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentID(i *int) *CategoryUpdate {
	if i != nil {
		cu.SetParentID(*i)
	}
	return cu
}

// AddParentID adds i to the "parent_id" field.
func (cu *CategoryUpdate) AddParentID(i int) *CategoryUpdate {
	cu.mutation.AddParentID(i)
	return cu
}

// ClearParentID clears the value of the "parent_id" field.
func (cu *CategoryUpdate) ClearParentID() *CategoryUpdate {
	cu.mutation.ClearParentID()
	return cu
}

// AddChildIDs adds the "children" edge to the Category entity by IDs.
func (cu *CategoryUpdate) AddChildIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddChildIDs(ids...)
	return cu
}

// AddChildren adds the "children" edges to the Category entity.
func (cu *CategoryUpdate) AddChildren(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChildIDs(ids...)
}

// AddParentIDs adds the "parent" edge to the Category entity by IDs.
func (cu *CategoryUpdate) AddParentIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddParentIDs(ids...)
	return cu
}

// AddParent adds the "parent" edges to the Category entity.
func (cu *CategoryUpdate) AddParent(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddParentIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearChildren clears all "children" edges to the Category entity.
func (cu *CategoryUpdate) ClearChildren() *CategoryUpdate {
	cu.mutation.ClearChildren()
	return cu
}

// RemoveChildIDs removes the "children" edge to Category entities by IDs.
func (cu *CategoryUpdate) RemoveChildIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveChildIDs(ids...)
	return cu
}

// RemoveChildren removes "children" edges to Category entities.
func (cu *CategoryUpdate) RemoveChildren(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChildIDs(ids...)
}

// ClearParent clears all "parent" edges to the Category entity.
func (cu *CategoryUpdate) ClearParent() *CategoryUpdate {
	cu.mutation.ClearParent()
	return cu
}

// RemoveParentIDs removes the "parent" edge to Category entities by IDs.
func (cu *CategoryUpdate) RemoveParentIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveParentIDs(ids...)
	return cu
}

// RemoveParent removes "parent" edges to Category entities.
func (cu *CategoryUpdate) RemoveParent(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveParentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldTitle,
		})
	}
	if cu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldDescription,
		})
	}
	if cu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldDescription,
		})
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if cu.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: category.FieldParentID,
		})
	}
	if cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedParentIDs(); len(nodes) > 0 && !cu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *CategoryUpdateOne) SetTitle(s string) *CategoryUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableTitle(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// ClearTitle clears the value of the "title" field.
func (cuo *CategoryUpdateOne) ClearTitle() *CategoryUpdateOne {
	cuo.mutation.ClearTitle()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDescription(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CategoryUpdateOne) ClearDescription() *CategoryUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(i int) *CategoryUpdateOne {
	cuo.mutation.ResetParentID()
	cuo.mutation.SetParentID(i)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentID(i *int) *CategoryUpdateOne {
	if i != nil {
		cuo.SetParentID(*i)
	}
	return cuo
}

// AddParentID adds i to the "parent_id" field.
func (cuo *CategoryUpdateOne) AddParentID(i int) *CategoryUpdateOne {
	cuo.mutation.AddParentID(i)
	return cuo
}

// ClearParentID clears the value of the "parent_id" field.
func (cuo *CategoryUpdateOne) ClearParentID() *CategoryUpdateOne {
	cuo.mutation.ClearParentID()
	return cuo
}

// AddChildIDs adds the "children" edge to the Category entity by IDs.
func (cuo *CategoryUpdateOne) AddChildIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddChildIDs(ids...)
	return cuo
}

// AddChildren adds the "children" edges to the Category entity.
func (cuo *CategoryUpdateOne) AddChildren(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChildIDs(ids...)
}

// AddParentIDs adds the "parent" edge to the Category entity by IDs.
func (cuo *CategoryUpdateOne) AddParentIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddParentIDs(ids...)
	return cuo
}

// AddParent adds the "parent" edges to the Category entity.
func (cuo *CategoryUpdateOne) AddParent(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddParentIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearChildren clears all "children" edges to the Category entity.
func (cuo *CategoryUpdateOne) ClearChildren() *CategoryUpdateOne {
	cuo.mutation.ClearChildren()
	return cuo
}

// RemoveChildIDs removes the "children" edge to Category entities by IDs.
func (cuo *CategoryUpdateOne) RemoveChildIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveChildIDs(ids...)
	return cuo
}

// RemoveChildren removes "children" edges to Category entities.
func (cuo *CategoryUpdateOne) RemoveChildren(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChildIDs(ids...)
}

// ClearParent clears all "parent" edges to the Category entity.
func (cuo *CategoryUpdateOne) ClearParent() *CategoryUpdateOne {
	cuo.mutation.ClearParent()
	return cuo
}

// RemoveParentIDs removes the "parent" edge to Category entities by IDs.
func (cuo *CategoryUpdateOne) RemoveParentIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveParentIDs(ids...)
	return cuo
}

// RemoveParent removes "parent" edges to Category entities.
func (cuo *CategoryUpdateOne) RemoveParent(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveParentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
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
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldTitle,
		})
	}
	if cuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldDescription,
		})
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if cuo.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: category.FieldParentID,
		})
	}
	if cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: category.ChildrenPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedParentIDs(); len(nodes) > 0 && !cuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: category.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
