package odoo

import (
	"fmt"
)

// LunchOrderTemp represents lunch.order.temp model.
type LunchOrderTemp struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AvailableToppings1 *Bool      `xmlrpc:"available_toppings_1,omptempty"`
	AvailableToppings2 *Bool      `xmlrpc:"available_toppings_2,omptempty"`
	AvailableToppings3 *Bool      `xmlrpc:"available_toppings_3,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Edit               *Bool      `xmlrpc:"edit,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	Image128           *String    `xmlrpc:"image_128,omptempty"`
	Image1920          *String    `xmlrpc:"image_1920,omptempty"`
	Note               *String    `xmlrpc:"note,omptempty"`
	PriceTotal         *Float     `xmlrpc:"price_total,omptempty"`
	ProductCategory    *Many2One  `xmlrpc:"product_category,omptempty"`
	ProductDescription *String    `xmlrpc:"product_description,omptempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductName        *String    `xmlrpc:"product_name,omptempty"`
	Quantity           *Float     `xmlrpc:"quantity,omptempty"`
	ToppingIds1        *Relation  `xmlrpc:"topping_ids_1,omptempty"`
	ToppingIds2        *Relation  `xmlrpc:"topping_ids_2,omptempty"`
	ToppingIds3        *Relation  `xmlrpc:"topping_ids_3,omptempty"`
	ToppingLabel1      *String    `xmlrpc:"topping_label_1,omptempty"`
	ToppingLabel2      *String    `xmlrpc:"topping_label_2,omptempty"`
	ToppingLabel3      *String    `xmlrpc:"topping_label_3,omptempty"`
	ToppingQuantity1   *Selection `xmlrpc:"topping_quantity_1,omptempty"`
	ToppingQuantity2   *Selection `xmlrpc:"topping_quantity_2,omptempty"`
	ToppingQuantity3   *Selection `xmlrpc:"topping_quantity_3,omptempty"`
	UserId             *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LunchOrderTemps represents array of lunch.order.temp model.
type LunchOrderTemps []LunchOrderTemp

// LunchOrderTempModel is the odoo model name.
const LunchOrderTempModel = "lunch.order.temp"

// Many2One convert LunchOrderTemp to *Many2One.
func (lot *LunchOrderTemp) Many2One() *Many2One {
	return NewMany2One(lot.Id.Get(), "")
}

// CreateLunchOrderTemp creates a new lunch.order.temp model and returns its id.
func (c *Client) CreateLunchOrderTemp(lot *LunchOrderTemp) (int64, error) {
	ids, err := c.CreateLunchOrderTemps([]*LunchOrderTemp{lot})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchOrderTemp creates a new lunch.order.temp model and returns its id.
func (c *Client) CreateLunchOrderTemps(lots []*LunchOrderTemp) ([]int64, error) {
	var vv []interface{}
	for _, v := range lots {
		vv = append(vv, v)
	}
	return c.Create(LunchOrderTempModel, vv)
}

// UpdateLunchOrderTemp updates an existing lunch.order.temp record.
func (c *Client) UpdateLunchOrderTemp(lot *LunchOrderTemp) error {
	return c.UpdateLunchOrderTemps([]int64{lot.Id.Get()}, lot)
}

// UpdateLunchOrderTemps updates existing lunch.order.temp records.
// All records (represented by ids) will be updated by lot values.
func (c *Client) UpdateLunchOrderTemps(ids []int64, lot *LunchOrderTemp) error {
	return c.Update(LunchOrderTempModel, ids, lot)
}

// DeleteLunchOrderTemp deletes an existing lunch.order.temp record.
func (c *Client) DeleteLunchOrderTemp(id int64) error {
	return c.DeleteLunchOrderTemps([]int64{id})
}

// DeleteLunchOrderTemps deletes existing lunch.order.temp records.
func (c *Client) DeleteLunchOrderTemps(ids []int64) error {
	return c.Delete(LunchOrderTempModel, ids)
}

// GetLunchOrderTemp gets lunch.order.temp existing record.
func (c *Client) GetLunchOrderTemp(id int64) (*LunchOrderTemp, error) {
	lots, err := c.GetLunchOrderTemps([]int64{id})
	if err != nil {
		return nil, err
	}
	if lots != nil && len(*lots) > 0 {
		return &((*lots)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.order.temp not found", id)
}

// GetLunchOrderTemps gets lunch.order.temp existing records.
func (c *Client) GetLunchOrderTemps(ids []int64) (*LunchOrderTemps, error) {
	lots := &LunchOrderTemps{}
	if err := c.Read(LunchOrderTempModel, ids, nil, lots); err != nil {
		return nil, err
	}
	return lots, nil
}

// FindLunchOrderTemp finds lunch.order.temp record by querying it with criteria.
func (c *Client) FindLunchOrderTemp(criteria *Criteria) (*LunchOrderTemp, error) {
	lots := &LunchOrderTemps{}
	if err := c.SearchRead(LunchOrderTempModel, criteria, NewOptions().Limit(1), lots); err != nil {
		return nil, err
	}
	if lots != nil && len(*lots) > 0 {
		return &((*lots)[0]), nil
	}
	return nil, fmt.Errorf("lunch.order.temp was not found with criteria %v", criteria)
}

// FindLunchOrderTemps finds lunch.order.temp records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrderTemps(criteria *Criteria, options *Options) (*LunchOrderTemps, error) {
	lots := &LunchOrderTemps{}
	if err := c.SearchRead(LunchOrderTempModel, criteria, options, lots); err != nil {
		return nil, err
	}
	return lots, nil
}

// FindLunchOrderTempIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrderTempIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchOrderTempModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchOrderTempId finds record id by querying it with criteria.
func (c *Client) FindLunchOrderTempId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchOrderTempModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.order.temp was not found with criteria %v and options %v", criteria, options)
}
