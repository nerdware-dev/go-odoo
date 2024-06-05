package odoo

import (
	"fmt"
)

// LunchLocation represents lunch.location model.
type LunchLocation struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Address     *String   `xmlrpc:"address,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LunchLocations represents array of lunch.location model.
type LunchLocations []LunchLocation

// LunchLocationModel is the odoo model name.
const LunchLocationModel = "lunch.location"

// Many2One convert LunchLocation to *Many2One.
func (ll *LunchLocation) Many2One() *Many2One {
	return NewMany2One(ll.Id.Get(), "")
}

// CreateLunchLocation creates a new lunch.location model and returns its id.
func (c *Client) CreateLunchLocation(ll *LunchLocation) (int64, error) {
	ids, err := c.CreateLunchLocations([]*LunchLocation{ll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchLocation creates a new lunch.location model and returns its id.
func (c *Client) CreateLunchLocations(lls []*LunchLocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range lls {
		vv = append(vv, v)
	}
	return c.Create(LunchLocationModel, vv)
}

// UpdateLunchLocation updates an existing lunch.location record.
func (c *Client) UpdateLunchLocation(ll *LunchLocation) error {
	return c.UpdateLunchLocations([]int64{ll.Id.Get()}, ll)
}

// UpdateLunchLocations updates existing lunch.location records.
// All records (represented by ids) will be updated by ll values.
func (c *Client) UpdateLunchLocations(ids []int64, ll *LunchLocation) error {
	return c.Update(LunchLocationModel, ids, ll)
}

// DeleteLunchLocation deletes an existing lunch.location record.
func (c *Client) DeleteLunchLocation(id int64) error {
	return c.DeleteLunchLocations([]int64{id})
}

// DeleteLunchLocations deletes existing lunch.location records.
func (c *Client) DeleteLunchLocations(ids []int64) error {
	return c.Delete(LunchLocationModel, ids)
}

// GetLunchLocation gets lunch.location existing record.
func (c *Client) GetLunchLocation(id int64) (*LunchLocation, error) {
	lls, err := c.GetLunchLocations([]int64{id})
	if err != nil {
		return nil, err
	}
	if lls != nil && len(*lls) > 0 {
		return &((*lls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.location not found", id)
}

// GetLunchLocations gets lunch.location existing records.
func (c *Client) GetLunchLocations(ids []int64) (*LunchLocations, error) {
	lls := &LunchLocations{}
	if err := c.Read(LunchLocationModel, ids, nil, lls); err != nil {
		return nil, err
	}
	return lls, nil
}

// FindLunchLocation finds lunch.location record by querying it with criteria.
func (c *Client) FindLunchLocation(criteria *Criteria) (*LunchLocation, error) {
	lls := &LunchLocations{}
	if err := c.SearchRead(LunchLocationModel, criteria, NewOptions().Limit(1), lls); err != nil {
		return nil, err
	}
	if lls != nil && len(*lls) > 0 {
		return &((*lls)[0]), nil
	}
	return nil, fmt.Errorf("lunch.location was not found with criteria %v", criteria)
}

// FindLunchLocations finds lunch.location records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchLocations(criteria *Criteria, options *Options) (*LunchLocations, error) {
	lls := &LunchLocations{}
	if err := c.SearchRead(LunchLocationModel, criteria, options, lls); err != nil {
		return nil, err
	}
	return lls, nil
}

// FindLunchLocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchLocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchLocationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchLocationId finds record id by querying it with criteria.
func (c *Client) FindLunchLocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchLocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.location was not found with criteria %v and options %v", criteria, options)
}
