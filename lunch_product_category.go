package odoo

import (
	"fmt"
)

// LunchProductCategory represents lunch.product.category model.
type LunchProductCategory struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId        *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId       *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Image1024        *String    `xmlrpc:"image_1024,omptempty"`
	Image128         *String    `xmlrpc:"image_128,omptempty"`
	Image1920        *String    `xmlrpc:"image_1920,omptempty"`
	Image256         *String    `xmlrpc:"image_256,omptempty"`
	Image512         *String    `xmlrpc:"image_512,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	ProductCount     *Int       `xmlrpc:"product_count,omptempty"`
	ToppingIds1      *Relation  `xmlrpc:"topping_ids_1,omptempty"`
	ToppingIds2      *Relation  `xmlrpc:"topping_ids_2,omptempty"`
	ToppingIds3      *Relation  `xmlrpc:"topping_ids_3,omptempty"`
	ToppingLabel1    *String    `xmlrpc:"topping_label_1,omptempty"`
	ToppingLabel2    *String    `xmlrpc:"topping_label_2,omptempty"`
	ToppingLabel3    *String    `xmlrpc:"topping_label_3,omptempty"`
	ToppingQuantity1 *Selection `xmlrpc:"topping_quantity_1,omptempty"`
	ToppingQuantity2 *Selection `xmlrpc:"topping_quantity_2,omptempty"`
	ToppingQuantity3 *Selection `xmlrpc:"topping_quantity_3,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LunchProductCategorys represents array of lunch.product.category model.
type LunchProductCategorys []LunchProductCategory

// LunchProductCategoryModel is the odoo model name.
const LunchProductCategoryModel = "lunch.product.category"

// Many2One convert LunchProductCategory to *Many2One.
func (lpc *LunchProductCategory) Many2One() *Many2One {
	return NewMany2One(lpc.Id.Get(), "")
}

// CreateLunchProductCategory creates a new lunch.product.category model and returns its id.
func (c *Client) CreateLunchProductCategory(lpc *LunchProductCategory) (int64, error) {
	ids, err := c.CreateLunchProductCategorys([]*LunchProductCategory{lpc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchProductCategory creates a new lunch.product.category model and returns its id.
func (c *Client) CreateLunchProductCategorys(lpcs []*LunchProductCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range lpcs {
		vv = append(vv, v)
	}
	return c.Create(LunchProductCategoryModel, vv)
}

// UpdateLunchProductCategory updates an existing lunch.product.category record.
func (c *Client) UpdateLunchProductCategory(lpc *LunchProductCategory) error {
	return c.UpdateLunchProductCategorys([]int64{lpc.Id.Get()}, lpc)
}

// UpdateLunchProductCategorys updates existing lunch.product.category records.
// All records (represented by ids) will be updated by lpc values.
func (c *Client) UpdateLunchProductCategorys(ids []int64, lpc *LunchProductCategory) error {
	return c.Update(LunchProductCategoryModel, ids, lpc)
}

// DeleteLunchProductCategory deletes an existing lunch.product.category record.
func (c *Client) DeleteLunchProductCategory(id int64) error {
	return c.DeleteLunchProductCategorys([]int64{id})
}

// DeleteLunchProductCategorys deletes existing lunch.product.category records.
func (c *Client) DeleteLunchProductCategorys(ids []int64) error {
	return c.Delete(LunchProductCategoryModel, ids)
}

// GetLunchProductCategory gets lunch.product.category existing record.
func (c *Client) GetLunchProductCategory(id int64) (*LunchProductCategory, error) {
	lpcs, err := c.GetLunchProductCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if lpcs != nil && len(*lpcs) > 0 {
		return &((*lpcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.product.category not found", id)
}

// GetLunchProductCategorys gets lunch.product.category existing records.
func (c *Client) GetLunchProductCategorys(ids []int64) (*LunchProductCategorys, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.Read(LunchProductCategoryModel, ids, nil, lpcs); err != nil {
		return nil, err
	}
	return lpcs, nil
}

// FindLunchProductCategory finds lunch.product.category record by querying it with criteria.
func (c *Client) FindLunchProductCategory(criteria *Criteria) (*LunchProductCategory, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.SearchRead(LunchProductCategoryModel, criteria, NewOptions().Limit(1), lpcs); err != nil {
		return nil, err
	}
	if lpcs != nil && len(*lpcs) > 0 {
		return &((*lpcs)[0]), nil
	}
	return nil, fmt.Errorf("lunch.product.category was not found with criteria %v", criteria)
}

// FindLunchProductCategorys finds lunch.product.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductCategorys(criteria *Criteria, options *Options) (*LunchProductCategorys, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.SearchRead(LunchProductCategoryModel, criteria, options, lpcs); err != nil {
		return nil, err
	}
	return lpcs, nil
}

// FindLunchProductCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchProductCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchProductCategoryId finds record id by querying it with criteria.
func (c *Client) FindLunchProductCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchProductCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.product.category was not found with criteria %v and options %v", criteria, options)
}
