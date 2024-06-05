package odoo

import (
	"fmt"
)

// LunchOrder represents lunch.order model.
type LunchOrder struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Active          *Bool      `xmlrpc:"active,omptempty"`
	CategoryId      *Many2One  `xmlrpc:"category_id,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date            *Time      `xmlrpc:"date,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	DisplayToppings *String    `xmlrpc:"display_toppings,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	Note            *String    `xmlrpc:"note,omptempty"`
	Price           *Float     `xmlrpc:"price,omptempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omptempty"`
	Quantity        *Float     `xmlrpc:"quantity,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	SupplierId      *Many2One  `xmlrpc:"supplier_id,omptempty"`
	ToppingIds1     *Relation  `xmlrpc:"topping_ids_1,omptempty"`
	ToppingIds2     *Relation  `xmlrpc:"topping_ids_2,omptempty"`
	ToppingIds3     *Relation  `xmlrpc:"topping_ids_3,omptempty"`
	UserId          *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LunchOrders represents array of lunch.order model.
type LunchOrders []LunchOrder

// LunchOrderModel is the odoo model name.
const LunchOrderModel = "lunch.order"

// Many2One convert LunchOrder to *Many2One.
func (lo *LunchOrder) Many2One() *Many2One {
	return NewMany2One(lo.Id.Get(), "")
}

// CreateLunchOrder creates a new lunch.order model and returns its id.
func (c *Client) CreateLunchOrder(lo *LunchOrder) (int64, error) {
	ids, err := c.CreateLunchOrders([]*LunchOrder{lo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchOrder creates a new lunch.order model and returns its id.
func (c *Client) CreateLunchOrders(los []*LunchOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range los {
		vv = append(vv, v)
	}
	return c.Create(LunchOrderModel, vv)
}

// UpdateLunchOrder updates an existing lunch.order record.
func (c *Client) UpdateLunchOrder(lo *LunchOrder) error {
	return c.UpdateLunchOrders([]int64{lo.Id.Get()}, lo)
}

// UpdateLunchOrders updates existing lunch.order records.
// All records (represented by ids) will be updated by lo values.
func (c *Client) UpdateLunchOrders(ids []int64, lo *LunchOrder) error {
	return c.Update(LunchOrderModel, ids, lo)
}

// DeleteLunchOrder deletes an existing lunch.order record.
func (c *Client) DeleteLunchOrder(id int64) error {
	return c.DeleteLunchOrders([]int64{id})
}

// DeleteLunchOrders deletes existing lunch.order records.
func (c *Client) DeleteLunchOrders(ids []int64) error {
	return c.Delete(LunchOrderModel, ids)
}

// GetLunchOrder gets lunch.order existing record.
func (c *Client) GetLunchOrder(id int64) (*LunchOrder, error) {
	los, err := c.GetLunchOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if los != nil && len(*los) > 0 {
		return &((*los)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.order not found", id)
}

// GetLunchOrders gets lunch.order existing records.
func (c *Client) GetLunchOrders(ids []int64) (*LunchOrders, error) {
	los := &LunchOrders{}
	if err := c.Read(LunchOrderModel, ids, nil, los); err != nil {
		return nil, err
	}
	return los, nil
}

// FindLunchOrder finds lunch.order record by querying it with criteria.
func (c *Client) FindLunchOrder(criteria *Criteria) (*LunchOrder, error) {
	los := &LunchOrders{}
	if err := c.SearchRead(LunchOrderModel, criteria, NewOptions().Limit(1), los); err != nil {
		return nil, err
	}
	if los != nil && len(*los) > 0 {
		return &((*los)[0]), nil
	}
	return nil, fmt.Errorf("lunch.order was not found with criteria %v", criteria)
}

// FindLunchOrders finds lunch.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrders(criteria *Criteria, options *Options) (*LunchOrders, error) {
	los := &LunchOrders{}
	if err := c.SearchRead(LunchOrderModel, criteria, options, los); err != nil {
		return nil, err
	}
	return los, nil
}

// FindLunchOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchOrderId finds record id by querying it with criteria.
func (c *Client) FindLunchOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.order was not found with criteria %v and options %v", criteria, options)
}
