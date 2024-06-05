package odoo

import (
	"fmt"
)

// LunchCashmove represents lunch.cashmove model.
type LunchCashmove struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Amount      *Float    `xmlrpc:"amount,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omptempty"`
	Date        *Time     `xmlrpc:"date,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LunchCashmoves represents array of lunch.cashmove model.
type LunchCashmoves []LunchCashmove

// LunchCashmoveModel is the odoo model name.
const LunchCashmoveModel = "lunch.cashmove"

// Many2One convert LunchCashmove to *Many2One.
func (lc *LunchCashmove) Many2One() *Many2One {
	return NewMany2One(lc.Id.Get(), "")
}

// CreateLunchCashmove creates a new lunch.cashmove model and returns its id.
func (c *Client) CreateLunchCashmove(lc *LunchCashmove) (int64, error) {
	ids, err := c.CreateLunchCashmoves([]*LunchCashmove{lc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchCashmove creates a new lunch.cashmove model and returns its id.
func (c *Client) CreateLunchCashmoves(lcs []*LunchCashmove) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcs {
		vv = append(vv, v)
	}
	return c.Create(LunchCashmoveModel, vv)
}

// UpdateLunchCashmove updates an existing lunch.cashmove record.
func (c *Client) UpdateLunchCashmove(lc *LunchCashmove) error {
	return c.UpdateLunchCashmoves([]int64{lc.Id.Get()}, lc)
}

// UpdateLunchCashmoves updates existing lunch.cashmove records.
// All records (represented by ids) will be updated by lc values.
func (c *Client) UpdateLunchCashmoves(ids []int64, lc *LunchCashmove) error {
	return c.Update(LunchCashmoveModel, ids, lc)
}

// DeleteLunchCashmove deletes an existing lunch.cashmove record.
func (c *Client) DeleteLunchCashmove(id int64) error {
	return c.DeleteLunchCashmoves([]int64{id})
}

// DeleteLunchCashmoves deletes existing lunch.cashmove records.
func (c *Client) DeleteLunchCashmoves(ids []int64) error {
	return c.Delete(LunchCashmoveModel, ids)
}

// GetLunchCashmove gets lunch.cashmove existing record.
func (c *Client) GetLunchCashmove(id int64) (*LunchCashmove, error) {
	lcs, err := c.GetLunchCashmoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if lcs != nil && len(*lcs) > 0 {
		return &((*lcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.cashmove not found", id)
}

// GetLunchCashmoves gets lunch.cashmove existing records.
func (c *Client) GetLunchCashmoves(ids []int64) (*LunchCashmoves, error) {
	lcs := &LunchCashmoves{}
	if err := c.Read(LunchCashmoveModel, ids, nil, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLunchCashmove finds lunch.cashmove record by querying it with criteria.
func (c *Client) FindLunchCashmove(criteria *Criteria) (*LunchCashmove, error) {
	lcs := &LunchCashmoves{}
	if err := c.SearchRead(LunchCashmoveModel, criteria, NewOptions().Limit(1), lcs); err != nil {
		return nil, err
	}
	if lcs != nil && len(*lcs) > 0 {
		return &((*lcs)[0]), nil
	}
	return nil, fmt.Errorf("lunch.cashmove was not found with criteria %v", criteria)
}

// FindLunchCashmoves finds lunch.cashmove records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoves(criteria *Criteria, options *Options) (*LunchCashmoves, error) {
	lcs := &LunchCashmoves{}
	if err := c.SearchRead(LunchCashmoveModel, criteria, options, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLunchCashmoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchCashmoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchCashmoveId finds record id by querying it with criteria.
func (c *Client) FindLunchCashmoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchCashmoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.cashmove was not found with criteria %v and options %v", criteria, options)
}
