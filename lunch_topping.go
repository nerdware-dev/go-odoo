package odoo

import (
	"fmt"
)

// LunchTopping represents lunch.topping model.
type LunchTopping struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omptempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Price           *Float    `xmlrpc:"price,omptempty"`
	ToppingCategory *Int      `xmlrpc:"topping_category,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LunchToppings represents array of lunch.topping model.
type LunchToppings []LunchTopping

// LunchToppingModel is the odoo model name.
const LunchToppingModel = "lunch.topping"

// Many2One convert LunchTopping to *Many2One.
func (lt *LunchTopping) Many2One() *Many2One {
	return NewMany2One(lt.Id.Get(), "")
}

// CreateLunchTopping creates a new lunch.topping model and returns its id.
func (c *Client) CreateLunchTopping(lt *LunchTopping) (int64, error) {
	ids, err := c.CreateLunchToppings([]*LunchTopping{lt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchTopping creates a new lunch.topping model and returns its id.
func (c *Client) CreateLunchToppings(lts []*LunchTopping) ([]int64, error) {
	var vv []interface{}
	for _, v := range lts {
		vv = append(vv, v)
	}
	return c.Create(LunchToppingModel, vv)
}

// UpdateLunchTopping updates an existing lunch.topping record.
func (c *Client) UpdateLunchTopping(lt *LunchTopping) error {
	return c.UpdateLunchToppings([]int64{lt.Id.Get()}, lt)
}

// UpdateLunchToppings updates existing lunch.topping records.
// All records (represented by ids) will be updated by lt values.
func (c *Client) UpdateLunchToppings(ids []int64, lt *LunchTopping) error {
	return c.Update(LunchToppingModel, ids, lt)
}

// DeleteLunchTopping deletes an existing lunch.topping record.
func (c *Client) DeleteLunchTopping(id int64) error {
	return c.DeleteLunchToppings([]int64{id})
}

// DeleteLunchToppings deletes existing lunch.topping records.
func (c *Client) DeleteLunchToppings(ids []int64) error {
	return c.Delete(LunchToppingModel, ids)
}

// GetLunchTopping gets lunch.topping existing record.
func (c *Client) GetLunchTopping(id int64) (*LunchTopping, error) {
	lts, err := c.GetLunchToppings([]int64{id})
	if err != nil {
		return nil, err
	}
	if lts != nil && len(*lts) > 0 {
		return &((*lts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.topping not found", id)
}

// GetLunchToppings gets lunch.topping existing records.
func (c *Client) GetLunchToppings(ids []int64) (*LunchToppings, error) {
	lts := &LunchToppings{}
	if err := c.Read(LunchToppingModel, ids, nil, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLunchTopping finds lunch.topping record by querying it with criteria.
func (c *Client) FindLunchTopping(criteria *Criteria) (*LunchTopping, error) {
	lts := &LunchToppings{}
	if err := c.SearchRead(LunchToppingModel, criteria, NewOptions().Limit(1), lts); err != nil {
		return nil, err
	}
	if lts != nil && len(*lts) > 0 {
		return &((*lts)[0]), nil
	}
	return nil, fmt.Errorf("lunch.topping was not found with criteria %v", criteria)
}

// FindLunchToppings finds lunch.topping records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchToppings(criteria *Criteria, options *Options) (*LunchToppings, error) {
	lts := &LunchToppings{}
	if err := c.SearchRead(LunchToppingModel, criteria, options, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLunchToppingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchToppingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchToppingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchToppingId finds record id by querying it with criteria.
func (c *Client) FindLunchToppingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchToppingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.topping was not found with criteria %v and options %v", criteria, options)
}
