// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-core/rpc/ent/api"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/rpc/ent/position"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/ent/token"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type APIPager struct {
	Order  api.Order
	Filter func(*APIQuery) (*APIQuery, error)
}

// APIPaginateOption enables pagination customization.
type APIPaginateOption func(*APIPager)

// DefaultAPIOrder is the default ordering of API.
var DefaultAPIOrder = Desc(api.FieldID)

func newAPIPager(opts []APIPaginateOption) (*APIPager, error) {
	pager := &APIPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAPIOrder
	}
	return pager, nil
}

func (p *APIPager) ApplyFilter(query *APIQuery) (*APIQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// APIPageList is API PageList result.
type APIPageList struct {
	List        []*API       `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (a *APIQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...APIPaginateOption,
) (*APIPageList, error) {

	pager, err := newAPIPager(opts)
	if err != nil {
		return nil, err
	}

	if a, err = pager.ApplyFilter(a); err != nil {
		return nil, err
	}

	ret := &APIPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := a.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		a = a.Order(pager.Order)
	} else {
		a = a.Order(DefaultAPIOrder)
	}

	a = a.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type DepartmentPager struct {
	Order  department.Order
	Filter func(*DepartmentQuery) (*DepartmentQuery, error)
}

// DepartmentPaginateOption enables pagination customization.
type DepartmentPaginateOption func(*DepartmentPager)

// DefaultDepartmentOrder is the default ordering of Department.
var DefaultDepartmentOrder = Desc(department.FieldID)

func newDepartmentPager(opts []DepartmentPaginateOption) (*DepartmentPager, error) {
	pager := &DepartmentPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultDepartmentOrder
	}
	return pager, nil
}

func (p *DepartmentPager) ApplyFilter(query *DepartmentQuery) (*DepartmentQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// DepartmentPageList is Department PageList result.
type DepartmentPageList struct {
	List        []*Department `json:"list"`
	PageDetails *PageDetails  `json:"pageDetails"`
}

func (d *DepartmentQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DepartmentPaginateOption,
) (*DepartmentPageList, error) {

	pager, err := newDepartmentPager(opts)
	if err != nil {
		return nil, err
	}

	if d, err = pager.ApplyFilter(d); err != nil {
		return nil, err
	}

	ret := &DepartmentPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := d.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		d = d.Order(pager.Order)
	} else {
		d = d.Order(DefaultDepartmentOrder)
	}

	d = d.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := d.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type DictionaryPager struct {
	Order  dictionary.Order
	Filter func(*DictionaryQuery) (*DictionaryQuery, error)
}

// DictionaryPaginateOption enables pagination customization.
type DictionaryPaginateOption func(*DictionaryPager)

// DefaultDictionaryOrder is the default ordering of Dictionary.
var DefaultDictionaryOrder = Desc(dictionary.FieldID)

func newDictionaryPager(opts []DictionaryPaginateOption) (*DictionaryPager, error) {
	pager := &DictionaryPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultDictionaryOrder
	}
	return pager, nil
}

func (p *DictionaryPager) ApplyFilter(query *DictionaryQuery) (*DictionaryQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// DictionaryPageList is Dictionary PageList result.
type DictionaryPageList struct {
	List        []*Dictionary `json:"list"`
	PageDetails *PageDetails  `json:"pageDetails"`
}

func (d *DictionaryQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DictionaryPaginateOption,
) (*DictionaryPageList, error) {

	pager, err := newDictionaryPager(opts)
	if err != nil {
		return nil, err
	}

	if d, err = pager.ApplyFilter(d); err != nil {
		return nil, err
	}

	ret := &DictionaryPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := d.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		d = d.Order(pager.Order)
	} else {
		d = d.Order(DefaultDictionaryOrder)
	}

	d = d.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := d.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type DictionaryDetailPager struct {
	Order  dictionarydetail.Order
	Filter func(*DictionaryDetailQuery) (*DictionaryDetailQuery, error)
}

// DictionaryDetailPaginateOption enables pagination customization.
type DictionaryDetailPaginateOption func(*DictionaryDetailPager)

// DefaultDictionaryDetailOrder is the default ordering of DictionaryDetail.
var DefaultDictionaryDetailOrder = Desc(dictionarydetail.FieldID)

func newDictionaryDetailPager(opts []DictionaryDetailPaginateOption) (*DictionaryDetailPager, error) {
	pager := &DictionaryDetailPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultDictionaryDetailOrder
	}
	return pager, nil
}

func (p *DictionaryDetailPager) ApplyFilter(query *DictionaryDetailQuery) (*DictionaryDetailQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// DictionaryDetailPageList is DictionaryDetail PageList result.
type DictionaryDetailPageList struct {
	List        []*DictionaryDetail `json:"list"`
	PageDetails *PageDetails        `json:"pageDetails"`
}

func (dd *DictionaryDetailQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DictionaryDetailPaginateOption,
) (*DictionaryDetailPageList, error) {

	pager, err := newDictionaryDetailPager(opts)
	if err != nil {
		return nil, err
	}

	if dd, err = pager.ApplyFilter(dd); err != nil {
		return nil, err
	}

	ret := &DictionaryDetailPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := dd.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		dd = dd.Order(pager.Order)
	} else {
		dd = dd.Order(DefaultDictionaryDetailOrder)
	}

	dd = dd.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := dd.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type MenuPager struct {
	Order  menu.Order
	Filter func(*MenuQuery) (*MenuQuery, error)
}

// MenuPaginateOption enables pagination customization.
type MenuPaginateOption func(*MenuPager)

// DefaultMenuOrder is the default ordering of Menu.
var DefaultMenuOrder = Desc(menu.FieldID)

func newMenuPager(opts []MenuPaginateOption) (*MenuPager, error) {
	pager := &MenuPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultMenuOrder
	}
	return pager, nil
}

func (p *MenuPager) ApplyFilter(query *MenuQuery) (*MenuQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// MenuPageList is Menu PageList result.
type MenuPageList struct {
	List        []*Menu      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (m *MenuQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...MenuPaginateOption,
) (*MenuPageList, error) {

	pager, err := newMenuPager(opts)
	if err != nil {
		return nil, err
	}

	if m, err = pager.ApplyFilter(m); err != nil {
		return nil, err
	}

	ret := &MenuPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := m.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		m = m.Order(pager.Order)
	} else {
		m = m.Order(DefaultMenuOrder)
	}

	m = m.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OauthProviderPager struct {
	Order  oauthprovider.Order
	Filter func(*OauthProviderQuery) (*OauthProviderQuery, error)
}

// OauthProviderPaginateOption enables pagination customization.
type OauthProviderPaginateOption func(*OauthProviderPager)

// DefaultOauthProviderOrder is the default ordering of OauthProvider.
var DefaultOauthProviderOrder = Desc(oauthprovider.FieldID)

func newOauthProviderPager(opts []OauthProviderPaginateOption) (*OauthProviderPager, error) {
	pager := &OauthProviderPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOauthProviderOrder
	}
	return pager, nil
}

func (p *OauthProviderPager) ApplyFilter(query *OauthProviderQuery) (*OauthProviderQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OauthProviderPageList is OauthProvider PageList result.
type OauthProviderPageList struct {
	List        []*OauthProvider `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (op *OauthProviderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OauthProviderPaginateOption,
) (*OauthProviderPageList, error) {

	pager, err := newOauthProviderPager(opts)
	if err != nil {
		return nil, err
	}

	if op, err = pager.ApplyFilter(op); err != nil {
		return nil, err
	}

	ret := &OauthProviderPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := op.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		op = op.Order(pager.Order)
	} else {
		op = op.Order(DefaultOauthProviderOrder)
	}

	op = op.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := op.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type PositionPager struct {
	Order  position.Order
	Filter func(*PositionQuery) (*PositionQuery, error)
}

// PositionPaginateOption enables pagination customization.
type PositionPaginateOption func(*PositionPager)

// DefaultPositionOrder is the default ordering of Position.
var DefaultPositionOrder = Desc(position.FieldID)

func newPositionPager(opts []PositionPaginateOption) (*PositionPager, error) {
	pager := &PositionPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultPositionOrder
	}
	return pager, nil
}

func (p *PositionPager) ApplyFilter(query *PositionQuery) (*PositionQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// PositionPageList is Position PageList result.
type PositionPageList struct {
	List        []*Position  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (po *PositionQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...PositionPaginateOption,
) (*PositionPageList, error) {

	pager, err := newPositionPager(opts)
	if err != nil {
		return nil, err
	}

	if po, err = pager.ApplyFilter(po); err != nil {
		return nil, err
	}

	ret := &PositionPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := po.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		po = po.Order(pager.Order)
	} else {
		po = po.Order(DefaultPositionOrder)
	}

	po = po.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := po.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RolePager struct {
	Order  role.Order
	Filter func(*RoleQuery) (*RoleQuery, error)
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*RolePager)

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = Desc(role.FieldID)

func newRolePager(opts []RolePaginateOption) (*RolePager, error) {
	pager := &RolePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *RolePager) ApplyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RolePageList is Role PageList result.
type RolePageList struct {
	List        []*Role      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (r *RoleQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RolePaginateOption,
) (*RolePageList, error) {

	pager, err := newRolePager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.ApplyFilter(r); err != nil {
		return nil, err
	}

	ret := &RolePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := r.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		r = r.Order(pager.Order)
	} else {
		r = r.Order(DefaultRoleOrder)
	}

	r = r.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type TokenPager struct {
	Order  token.Order
	Filter func(*TokenQuery) (*TokenQuery, error)
}

// TokenPaginateOption enables pagination customization.
type TokenPaginateOption func(*TokenPager)

// DefaultTokenOrder is the default ordering of Token.
var DefaultTokenOrder = Desc(token.FieldID)

func newTokenPager(opts []TokenPaginateOption) (*TokenPager, error) {
	pager := &TokenPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultTokenOrder
	}
	return pager, nil
}

func (p *TokenPager) ApplyFilter(query *TokenQuery) (*TokenQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// TokenPageList is Token PageList result.
type TokenPageList struct {
	List        []*Token     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (t *TokenQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...TokenPaginateOption,
) (*TokenPageList, error) {

	pager, err := newTokenPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.ApplyFilter(t); err != nil {
		return nil, err
	}

	ret := &TokenPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := t.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		t = t.Order(pager.Order)
	} else {
		t = t.Order(DefaultTokenOrder)
	}

	t = t.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type UserPager struct {
	Order  user.Order
	Filter func(*UserQuery) (*UserQuery, error)
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*UserPager)

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = Desc(user.FieldID)

func newUserPager(opts []UserPaginateOption) (*UserPager, error) {
	pager := &UserPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultUserOrder
	}
	return pager, nil
}

func (p *UserPager) ApplyFilter(query *UserQuery) (*UserQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// UserPageList is User PageList result.
type UserPageList struct {
	List        []*User      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (u *UserQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...UserPaginateOption,
) (*UserPageList, error) {

	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.ApplyFilter(u); err != nil {
		return nil, err
	}

	ret := &UserPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := u.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		u = u.Order(pager.Order)
	} else {
		u = u.Order(DefaultUserOrder)
	}

	u = u.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
