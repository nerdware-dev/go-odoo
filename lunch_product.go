package odoo

import (
	"fmt"
)

// LunchProduct represents lunch.product model.
type LunchProduct struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Active          *Bool     `xmlrpc:"active,omptempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omptempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	Description     *String   `xmlrpc:"description,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	FavoriteUserIds *Relation `xmlrpc:"favorite_user_ids,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Image1024       *String   `xmlrpc:"image_1024,omptempty"`
	Image128        *String   `xmlrpc:"image_128,omptempty"`
	Image1920       *String   `xmlrpc:"image_1920,omptempty"`
	Image256        *String   `xmlrpc:"image_256,omptempty"`
	Image512        *String   `xmlrpc:"image_512,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	NewUntil        *Time     `xmlrpc:"new_until,omptempty"`
	Price           *Float    `xmlrpc:"price,omptempty"`
	SupplierId      *Many2One `xmlrpc:"supplier_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LunchProducts represents array of lunch.product model.
type LunchProducts []LunchProduct

// LunchProductModel is the odoo model name.
const LunchProductModel = "lunch.product"

// Many2One convert LunchProduct to *Many2One.
func (lp *LunchProduct) Many2One() *Many2One {
	return NewMany2One(lp.Id.Get(), "")
}

// CreateLunchProduct creates a new lunch.product model and returns its id.
func (c *Client) CreateLunchProduct(lp *LunchProduct) (int64, error) {
	ids, err := c.CreateLunchProducts([]*LunchProduct{lp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchProduct creates a new lunch.product model and returns its id.
func (c *Client) CreateLunchProducts(lps []*LunchProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range lps {
		vv = append(vv, v)
	}
	return c.Create(LunchProductModel, vv)
}

// UpdateLunchProduct updates an existing lunch.product record.
func (c *Client) UpdateLunchProduct(lp *LunchProduct) error {
	return c.UpdateLunchProducts([]int64{lp.Id.Get()}, lp)
}

// UpdateLunchProducts updates existing lunch.product records.
// All records (represented by ids) will be updated by lp values.
func (c *Client) UpdateLunchProducts(ids []int64, lp *LunchProduct) error {
	return c.Update(LunchProductModel, ids, lp)
}

// DeleteLunchProduct deletes an existing lunch.product record.
func (c *Client) DeleteLunchProduct(id int64) error {
	return c.DeleteLunchProducts([]int64{id})
}

// DeleteLunchProducts deletes existing lunch.product records.
func (c *Client) DeleteLunchProducts(ids []int64) error {
	return c.Delete(LunchProductModel, ids)
}

// GetLunchProduct gets lunch.product existing record.
func (c *Client) GetLunchProduct(id int64) (*LunchProduct, error) {
	lps, err := c.GetLunchProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if lps != nil && len(*lps) > 0 {
		return &((*lps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.product not found", id)
}

// GetLunchProducts gets lunch.product existing records.
func (c *Client) GetLunchProducts(ids []int64) (*LunchProducts, error) {
	lps := &LunchProducts{}
	if err := c.Read(LunchProductModel, ids, nil, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLunchProduct finds lunch.product record by querying it with criteria.
func (c *Client) FindLunchProduct(criteria *Criteria) (*LunchProduct, error) {
	lps := &LunchProducts{}
	if err := c.SearchRead(LunchProductModel, criteria, NewOptions().Limit(1), lps); err != nil {
		return nil, err
	}
	if lps != nil && len(*lps) > 0 {
		return &((*lps)[0]), nil
	}
	return nil, fmt.Errorf("lunch.product was not found with criteria %v", criteria)
}

// FindLunchProducts finds lunch.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProducts(criteria *Criteria, options *Options) (*LunchProducts, error) {
	lps := &LunchProducts{}
	if err := c.SearchRead(LunchProductModel, criteria, options, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLunchProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchProductId finds record id by querying it with criteria.
func (c *Client) FindLunchProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.product was not found with criteria %v and options %v", criteria, options)
}
