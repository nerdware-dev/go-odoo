package odoo

import (
	"fmt"
)

// AddToTour represents add.to.tour model.
type AddToTour struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Articles    *String   `xmlrpc:"articles,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	TourId      *Many2One `xmlrpc:"tour_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AddToTours represents array of add.to.tour model.
type AddToTours []AddToTour

// AddToTourModel is the odoo model name.
const AddToTourModel = "add.to.tour"

// Many2One convert AddToTour to *Many2One.
func (att *AddToTour) Many2One() *Many2One {
	return NewMany2One(att.Id.Get(), "")
}

// CreateAddToTour creates a new add.to.tour model and returns its id.
func (c *Client) CreateAddToTour(att *AddToTour) (int64, error) {
	ids, err := c.CreateAddToTours([]*AddToTour{att})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAddToTour creates a new add.to.tour model and returns its id.
func (c *Client) CreateAddToTours(atts []*AddToTour) ([]int64, error) {
	var vv []interface{}
	for _, v := range atts {
		vv = append(vv, v)
	}
	return c.Create(AddToTourModel, vv)
}

// UpdateAddToTour updates an existing add.to.tour record.
func (c *Client) UpdateAddToTour(att *AddToTour) error {
	return c.UpdateAddToTours([]int64{att.Id.Get()}, att)
}

// UpdateAddToTours updates existing add.to.tour records.
// All records (represented by ids) will be updated by att values.
func (c *Client) UpdateAddToTours(ids []int64, att *AddToTour) error {
	return c.Update(AddToTourModel, ids, att)
}

// DeleteAddToTour deletes an existing add.to.tour record.
func (c *Client) DeleteAddToTour(id int64) error {
	return c.DeleteAddToTours([]int64{id})
}

// DeleteAddToTours deletes existing add.to.tour records.
func (c *Client) DeleteAddToTours(ids []int64) error {
	return c.Delete(AddToTourModel, ids)
}

// GetAddToTour gets add.to.tour existing record.
func (c *Client) GetAddToTour(id int64) (*AddToTour, error) {
	atts, err := c.GetAddToTours([]int64{id})
	if err != nil {
		return nil, err
	}
	if atts != nil && len(*atts) > 0 {
		return &((*atts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of add.to.tour not found", id)
}

// GetAddToTours gets add.to.tour existing records.
func (c *Client) GetAddToTours(ids []int64) (*AddToTours, error) {
	atts := &AddToTours{}
	if err := c.Read(AddToTourModel, ids, nil, atts); err != nil {
		return nil, err
	}
	return atts, nil
}

// FindAddToTour finds add.to.tour record by querying it with criteria.
func (c *Client) FindAddToTour(criteria *Criteria) (*AddToTour, error) {
	atts := &AddToTours{}
	if err := c.SearchRead(AddToTourModel, criteria, NewOptions().Limit(1), atts); err != nil {
		return nil, err
	}
	if atts != nil && len(*atts) > 0 {
		return &((*atts)[0]), nil
	}
	return nil, fmt.Errorf("add.to.tour was not found with criteria %v", criteria)
}

// FindAddToTours finds add.to.tour records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAddToTours(criteria *Criteria, options *Options) (*AddToTours, error) {
	atts := &AddToTours{}
	if err := c.SearchRead(AddToTourModel, criteria, options, atts); err != nil {
		return nil, err
	}
	return atts, nil
}

// FindAddToTourIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAddToTourIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AddToTourModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAddToTourId finds record id by querying it with criteria.
func (c *Client) FindAddToTourId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AddToTourModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("add.to.tour was not found with criteria %v and options %v", criteria, options)
}
