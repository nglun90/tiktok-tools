// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/4meepo/tiktok-tools/ent/predicate"
	"github.com/4meepo/tiktok-tools/ent/tiktokcreator"
)

// TiktokCreatorUpdate is the builder for updating TiktokCreator entities.
type TiktokCreatorUpdate struct {
	config
	hooks    []Hook
	mutation *TiktokCreatorMutation
}

// Where appends a list predicates to the TiktokCreatorUpdate builder.
func (tcu *TiktokCreatorUpdate) Where(ps ...predicate.TiktokCreator) *TiktokCreatorUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetUpdateTime sets the "update_time" field.
func (tcu *TiktokCreatorUpdate) SetUpdateTime(t time.Time) *TiktokCreatorUpdate {
	tcu.mutation.SetUpdateTime(t)
	return tcu
}

// SetCreatorID sets the "creator_id" field.
func (tcu *TiktokCreatorUpdate) SetCreatorID(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetCreatorID(s)
	return tcu
}

// SetCreatorName sets the "creator_name" field.
func (tcu *TiktokCreatorUpdate) SetCreatorName(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetCreatorName(s)
	return tcu
}

// SetCreatorNickname sets the "creator_nickname" field.
func (tcu *TiktokCreatorUpdate) SetCreatorNickname(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetCreatorNickname(s)
	return tcu
}

// SetRegion sets the "region" field.
func (tcu *TiktokCreatorUpdate) SetRegion(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetRegion(s)
	return tcu
}

// SetProductCategories sets the "product_categories" field.
func (tcu *TiktokCreatorUpdate) SetProductCategories(s []string) *TiktokCreatorUpdate {
	tcu.mutation.SetProductCategories(s)
	return tcu
}

// AppendProductCategories appends s to the "product_categories" field.
func (tcu *TiktokCreatorUpdate) AppendProductCategories(s []string) *TiktokCreatorUpdate {
	tcu.mutation.AppendProductCategories(s)
	return tcu
}

// SetFollowerCount sets the "follower_count" field.
func (tcu *TiktokCreatorUpdate) SetFollowerCount(u uint32) *TiktokCreatorUpdate {
	tcu.mutation.ResetFollowerCount()
	tcu.mutation.SetFollowerCount(u)
	return tcu
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableFollowerCount(u *uint32) *TiktokCreatorUpdate {
	if u != nil {
		tcu.SetFollowerCount(*u)
	}
	return tcu
}

// AddFollowerCount adds u to the "follower_count" field.
func (tcu *TiktokCreatorUpdate) AddFollowerCount(u int32) *TiktokCreatorUpdate {
	tcu.mutation.AddFollowerCount(u)
	return tcu
}

// SetVideoAvgViewCnt sets the "video_avg_view_cnt" field.
func (tcu *TiktokCreatorUpdate) SetVideoAvgViewCnt(u uint32) *TiktokCreatorUpdate {
	tcu.mutation.ResetVideoAvgViewCnt()
	tcu.mutation.SetVideoAvgViewCnt(u)
	return tcu
}

// SetNillableVideoAvgViewCnt sets the "video_avg_view_cnt" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableVideoAvgViewCnt(u *uint32) *TiktokCreatorUpdate {
	if u != nil {
		tcu.SetVideoAvgViewCnt(*u)
	}
	return tcu
}

// AddVideoAvgViewCnt adds u to the "video_avg_view_cnt" field.
func (tcu *TiktokCreatorUpdate) AddVideoAvgViewCnt(u int32) *TiktokCreatorUpdate {
	tcu.mutation.AddVideoAvgViewCnt(u)
	return tcu
}

// SetVideoPubCnt sets the "video_pub_cnt" field.
func (tcu *TiktokCreatorUpdate) SetVideoPubCnt(u uint32) *TiktokCreatorUpdate {
	tcu.mutation.ResetVideoPubCnt()
	tcu.mutation.SetVideoPubCnt(u)
	return tcu
}

// SetNillableVideoPubCnt sets the "video_pub_cnt" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableVideoPubCnt(u *uint32) *TiktokCreatorUpdate {
	if u != nil {
		tcu.SetVideoPubCnt(*u)
	}
	return tcu
}

// AddVideoPubCnt adds u to the "video_pub_cnt" field.
func (tcu *TiktokCreatorUpdate) AddVideoPubCnt(u int32) *TiktokCreatorUpdate {
	tcu.mutation.AddVideoPubCnt(u)
	return tcu
}

// SetEcVideoAvgViewCnt sets the "ec_video_avg_view_cnt" field.
func (tcu *TiktokCreatorUpdate) SetEcVideoAvgViewCnt(u uint32) *TiktokCreatorUpdate {
	tcu.mutation.ResetEcVideoAvgViewCnt()
	tcu.mutation.SetEcVideoAvgViewCnt(u)
	return tcu
}

// SetNillableEcVideoAvgViewCnt sets the "ec_video_avg_view_cnt" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableEcVideoAvgViewCnt(u *uint32) *TiktokCreatorUpdate {
	if u != nil {
		tcu.SetEcVideoAvgViewCnt(*u)
	}
	return tcu
}

// AddEcVideoAvgViewCnt adds u to the "ec_video_avg_view_cnt" field.
func (tcu *TiktokCreatorUpdate) AddEcVideoAvgViewCnt(u int32) *TiktokCreatorUpdate {
	tcu.mutation.AddEcVideoAvgViewCnt(u)
	return tcu
}

// SetCreatorOecuid sets the "creator_oecuid" field.
func (tcu *TiktokCreatorUpdate) SetCreatorOecuid(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetCreatorOecuid(s)
	return tcu
}

// SetNillableCreatorOecuid sets the "creator_oecuid" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableCreatorOecuid(s *string) *TiktokCreatorUpdate {
	if s != nil {
		tcu.SetCreatorOecuid(*s)
	}
	return tcu
}

// SetCreatorTtuid sets the "creator_ttuid" field.
func (tcu *TiktokCreatorUpdate) SetCreatorTtuid(s string) *TiktokCreatorUpdate {
	tcu.mutation.SetCreatorTtuid(s)
	return tcu
}

// SetNillableCreatorTtuid sets the "creator_ttuid" field if the given value is not nil.
func (tcu *TiktokCreatorUpdate) SetNillableCreatorTtuid(s *string) *TiktokCreatorUpdate {
	if s != nil {
		tcu.SetCreatorTtuid(*s)
	}
	return tcu
}

// Mutation returns the TiktokCreatorMutation object of the builder.
func (tcu *TiktokCreatorUpdate) Mutation() *TiktokCreatorMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TiktokCreatorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tcu.defaults()
	if len(tcu.hooks) == 0 {
		if err = tcu.check(); err != nil {
			return 0, err
		}
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TiktokCreatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcu.check(); err != nil {
				return 0, err
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TiktokCreatorUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TiktokCreatorUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TiktokCreatorUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TiktokCreatorUpdate) defaults() {
	if _, ok := tcu.mutation.UpdateTime(); !ok {
		v := tiktokcreator.UpdateDefaultUpdateTime()
		tcu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcu *TiktokCreatorUpdate) check() error {
	if v, ok := tcu.mutation.CreatorID(); ok {
		if err := tiktokcreator.CreatorIDValidator(v); err != nil {
			return &ValidationError{Name: "creator_id", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_id": %w`, err)}
		}
	}
	if v, ok := tcu.mutation.CreatorName(); ok {
		if err := tiktokcreator.CreatorNameValidator(v); err != nil {
			return &ValidationError{Name: "creator_name", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_name": %w`, err)}
		}
	}
	if v, ok := tcu.mutation.CreatorNickname(); ok {
		if err := tiktokcreator.CreatorNicknameValidator(v); err != nil {
			return &ValidationError{Name: "creator_nickname", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_nickname": %w`, err)}
		}
	}
	if v, ok := tcu.mutation.Region(); ok {
		if err := tiktokcreator.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.region": %w`, err)}
		}
	}
	return nil
}

func (tcu *TiktokCreatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tiktokcreator.Table,
			Columns: tiktokcreator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tiktokcreator.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.UpdateTime(); ok {
		_spec.SetField(tiktokcreator.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tcu.mutation.CreatorID(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorID, field.TypeString, value)
	}
	if value, ok := tcu.mutation.CreatorName(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorName, field.TypeString, value)
	}
	if value, ok := tcu.mutation.CreatorNickname(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorNickname, field.TypeString, value)
	}
	if value, ok := tcu.mutation.Region(); ok {
		_spec.SetField(tiktokcreator.FieldRegion, field.TypeString, value)
	}
	if value, ok := tcu.mutation.ProductCategories(); ok {
		_spec.SetField(tiktokcreator.FieldProductCategories, field.TypeJSON, value)
	}
	if value, ok := tcu.mutation.AppendedProductCategories(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tiktokcreator.FieldProductCategories, value)
		})
	}
	if value, ok := tcu.mutation.FollowerCount(); ok {
		_spec.SetField(tiktokcreator.FieldFollowerCount, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.AddedFollowerCount(); ok {
		_spec.AddField(tiktokcreator.FieldFollowerCount, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.VideoAvgViewCnt(); ok {
		_spec.SetField(tiktokcreator.FieldVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.AddedVideoAvgViewCnt(); ok {
		_spec.AddField(tiktokcreator.FieldVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.VideoPubCnt(); ok {
		_spec.SetField(tiktokcreator.FieldVideoPubCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.AddedVideoPubCnt(); ok {
		_spec.AddField(tiktokcreator.FieldVideoPubCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.EcVideoAvgViewCnt(); ok {
		_spec.SetField(tiktokcreator.FieldEcVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.AddedEcVideoAvgViewCnt(); ok {
		_spec.AddField(tiktokcreator.FieldEcVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcu.mutation.CreatorOecuid(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorOecuid, field.TypeString, value)
	}
	if value, ok := tcu.mutation.CreatorTtuid(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorTtuid, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tiktokcreator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TiktokCreatorUpdateOne is the builder for updating a single TiktokCreator entity.
type TiktokCreatorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TiktokCreatorMutation
}

// SetUpdateTime sets the "update_time" field.
func (tcuo *TiktokCreatorUpdateOne) SetUpdateTime(t time.Time) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetUpdateTime(t)
	return tcuo
}

// SetCreatorID sets the "creator_id" field.
func (tcuo *TiktokCreatorUpdateOne) SetCreatorID(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetCreatorID(s)
	return tcuo
}

// SetCreatorName sets the "creator_name" field.
func (tcuo *TiktokCreatorUpdateOne) SetCreatorName(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetCreatorName(s)
	return tcuo
}

// SetCreatorNickname sets the "creator_nickname" field.
func (tcuo *TiktokCreatorUpdateOne) SetCreatorNickname(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetCreatorNickname(s)
	return tcuo
}

// SetRegion sets the "region" field.
func (tcuo *TiktokCreatorUpdateOne) SetRegion(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetRegion(s)
	return tcuo
}

// SetProductCategories sets the "product_categories" field.
func (tcuo *TiktokCreatorUpdateOne) SetProductCategories(s []string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetProductCategories(s)
	return tcuo
}

// AppendProductCategories appends s to the "product_categories" field.
func (tcuo *TiktokCreatorUpdateOne) AppendProductCategories(s []string) *TiktokCreatorUpdateOne {
	tcuo.mutation.AppendProductCategories(s)
	return tcuo
}

// SetFollowerCount sets the "follower_count" field.
func (tcuo *TiktokCreatorUpdateOne) SetFollowerCount(u uint32) *TiktokCreatorUpdateOne {
	tcuo.mutation.ResetFollowerCount()
	tcuo.mutation.SetFollowerCount(u)
	return tcuo
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableFollowerCount(u *uint32) *TiktokCreatorUpdateOne {
	if u != nil {
		tcuo.SetFollowerCount(*u)
	}
	return tcuo
}

// AddFollowerCount adds u to the "follower_count" field.
func (tcuo *TiktokCreatorUpdateOne) AddFollowerCount(u int32) *TiktokCreatorUpdateOne {
	tcuo.mutation.AddFollowerCount(u)
	return tcuo
}

// SetVideoAvgViewCnt sets the "video_avg_view_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) SetVideoAvgViewCnt(u uint32) *TiktokCreatorUpdateOne {
	tcuo.mutation.ResetVideoAvgViewCnt()
	tcuo.mutation.SetVideoAvgViewCnt(u)
	return tcuo
}

// SetNillableVideoAvgViewCnt sets the "video_avg_view_cnt" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableVideoAvgViewCnt(u *uint32) *TiktokCreatorUpdateOne {
	if u != nil {
		tcuo.SetVideoAvgViewCnt(*u)
	}
	return tcuo
}

// AddVideoAvgViewCnt adds u to the "video_avg_view_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) AddVideoAvgViewCnt(u int32) *TiktokCreatorUpdateOne {
	tcuo.mutation.AddVideoAvgViewCnt(u)
	return tcuo
}

// SetVideoPubCnt sets the "video_pub_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) SetVideoPubCnt(u uint32) *TiktokCreatorUpdateOne {
	tcuo.mutation.ResetVideoPubCnt()
	tcuo.mutation.SetVideoPubCnt(u)
	return tcuo
}

// SetNillableVideoPubCnt sets the "video_pub_cnt" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableVideoPubCnt(u *uint32) *TiktokCreatorUpdateOne {
	if u != nil {
		tcuo.SetVideoPubCnt(*u)
	}
	return tcuo
}

// AddVideoPubCnt adds u to the "video_pub_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) AddVideoPubCnt(u int32) *TiktokCreatorUpdateOne {
	tcuo.mutation.AddVideoPubCnt(u)
	return tcuo
}

// SetEcVideoAvgViewCnt sets the "ec_video_avg_view_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) SetEcVideoAvgViewCnt(u uint32) *TiktokCreatorUpdateOne {
	tcuo.mutation.ResetEcVideoAvgViewCnt()
	tcuo.mutation.SetEcVideoAvgViewCnt(u)
	return tcuo
}

// SetNillableEcVideoAvgViewCnt sets the "ec_video_avg_view_cnt" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableEcVideoAvgViewCnt(u *uint32) *TiktokCreatorUpdateOne {
	if u != nil {
		tcuo.SetEcVideoAvgViewCnt(*u)
	}
	return tcuo
}

// AddEcVideoAvgViewCnt adds u to the "ec_video_avg_view_cnt" field.
func (tcuo *TiktokCreatorUpdateOne) AddEcVideoAvgViewCnt(u int32) *TiktokCreatorUpdateOne {
	tcuo.mutation.AddEcVideoAvgViewCnt(u)
	return tcuo
}

// SetCreatorOecuid sets the "creator_oecuid" field.
func (tcuo *TiktokCreatorUpdateOne) SetCreatorOecuid(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetCreatorOecuid(s)
	return tcuo
}

// SetNillableCreatorOecuid sets the "creator_oecuid" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableCreatorOecuid(s *string) *TiktokCreatorUpdateOne {
	if s != nil {
		tcuo.SetCreatorOecuid(*s)
	}
	return tcuo
}

// SetCreatorTtuid sets the "creator_ttuid" field.
func (tcuo *TiktokCreatorUpdateOne) SetCreatorTtuid(s string) *TiktokCreatorUpdateOne {
	tcuo.mutation.SetCreatorTtuid(s)
	return tcuo
}

// SetNillableCreatorTtuid sets the "creator_ttuid" field if the given value is not nil.
func (tcuo *TiktokCreatorUpdateOne) SetNillableCreatorTtuid(s *string) *TiktokCreatorUpdateOne {
	if s != nil {
		tcuo.SetCreatorTtuid(*s)
	}
	return tcuo
}

// Mutation returns the TiktokCreatorMutation object of the builder.
func (tcuo *TiktokCreatorUpdateOne) Mutation() *TiktokCreatorMutation {
	return tcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TiktokCreatorUpdateOne) Select(field string, fields ...string) *TiktokCreatorUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TiktokCreator entity.
func (tcuo *TiktokCreatorUpdateOne) Save(ctx context.Context) (*TiktokCreator, error) {
	var (
		err  error
		node *TiktokCreator
	)
	tcuo.defaults()
	if len(tcuo.hooks) == 0 {
		if err = tcuo.check(); err != nil {
			return nil, err
		}
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TiktokCreatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcuo.check(); err != nil {
				return nil, err
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TiktokCreator)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TiktokCreatorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TiktokCreatorUpdateOne) SaveX(ctx context.Context) *TiktokCreator {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TiktokCreatorUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TiktokCreatorUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TiktokCreatorUpdateOne) defaults() {
	if _, ok := tcuo.mutation.UpdateTime(); !ok {
		v := tiktokcreator.UpdateDefaultUpdateTime()
		tcuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcuo *TiktokCreatorUpdateOne) check() error {
	if v, ok := tcuo.mutation.CreatorID(); ok {
		if err := tiktokcreator.CreatorIDValidator(v); err != nil {
			return &ValidationError{Name: "creator_id", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_id": %w`, err)}
		}
	}
	if v, ok := tcuo.mutation.CreatorName(); ok {
		if err := tiktokcreator.CreatorNameValidator(v); err != nil {
			return &ValidationError{Name: "creator_name", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_name": %w`, err)}
		}
	}
	if v, ok := tcuo.mutation.CreatorNickname(); ok {
		if err := tiktokcreator.CreatorNicknameValidator(v); err != nil {
			return &ValidationError{Name: "creator_nickname", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.creator_nickname": %w`, err)}
		}
	}
	if v, ok := tcuo.mutation.Region(); ok {
		if err := tiktokcreator.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "TiktokCreator.region": %w`, err)}
		}
	}
	return nil
}

func (tcuo *TiktokCreatorUpdateOne) sqlSave(ctx context.Context) (_node *TiktokCreator, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tiktokcreator.Table,
			Columns: tiktokcreator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tiktokcreator.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TiktokCreator.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tiktokcreator.FieldID)
		for _, f := range fields {
			if !tiktokcreator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tiktokcreator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.UpdateTime(); ok {
		_spec.SetField(tiktokcreator.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tcuo.mutation.CreatorID(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorID, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.CreatorName(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorName, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.CreatorNickname(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorNickname, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.Region(); ok {
		_spec.SetField(tiktokcreator.FieldRegion, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.ProductCategories(); ok {
		_spec.SetField(tiktokcreator.FieldProductCategories, field.TypeJSON, value)
	}
	if value, ok := tcuo.mutation.AppendedProductCategories(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tiktokcreator.FieldProductCategories, value)
		})
	}
	if value, ok := tcuo.mutation.FollowerCount(); ok {
		_spec.SetField(tiktokcreator.FieldFollowerCount, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.AddedFollowerCount(); ok {
		_spec.AddField(tiktokcreator.FieldFollowerCount, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.VideoAvgViewCnt(); ok {
		_spec.SetField(tiktokcreator.FieldVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.AddedVideoAvgViewCnt(); ok {
		_spec.AddField(tiktokcreator.FieldVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.VideoPubCnt(); ok {
		_spec.SetField(tiktokcreator.FieldVideoPubCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.AddedVideoPubCnt(); ok {
		_spec.AddField(tiktokcreator.FieldVideoPubCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.EcVideoAvgViewCnt(); ok {
		_spec.SetField(tiktokcreator.FieldEcVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.AddedEcVideoAvgViewCnt(); ok {
		_spec.AddField(tiktokcreator.FieldEcVideoAvgViewCnt, field.TypeUint32, value)
	}
	if value, ok := tcuo.mutation.CreatorOecuid(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorOecuid, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.CreatorTtuid(); ok {
		_spec.SetField(tiktokcreator.FieldCreatorTtuid, field.TypeString, value)
	}
	_node = &TiktokCreator{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tiktokcreator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
