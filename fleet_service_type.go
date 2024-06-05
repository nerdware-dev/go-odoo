package odoo

import (
	"fmt"
)

// FleetServiceType represents fleet.service.type model.
type FleetServiceType struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Category    *Selection `xmlrpc:"category,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// FleetServiceTypes represents array of fleet.service.type model.
type FleetServiceTypes []FleetServiceType

// FleetServiceTypeModel is the odoo model name.
const FleetServiceTypeModel = "fleet.service.type"

// Many2One convert FleetServiceType to *Many2One.
func (fst *FleetServiceType) Many2One() *Many2One {
	return NewMany2One(fst.Id.Get(), "")
}

// CreateFleetServiceType creates a new fleet.service.type model and returns its id.
func (c *Client) CreateFleetServiceType(fst *FleetServiceType) (int64, error) {
	ids, err := c.CreateFleetServiceTypes([]*FleetServiceType{fst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetServiceType creates a new fleet.service.type model and returns its id.
func (c *Client) CreateFleetServiceTypes(fsts []*FleetServiceType) ([]int64, error) {
	var vv []interface{}
	for _, v := range fsts {
		vv = append(vv, v)
	}
	return c.Create(FleetServiceTypeModel, vv)
}

// UpdateFleetServiceType updates an existing fleet.service.type record.
func (c *Client) UpdateFleetServiceType(fst *FleetServiceType) error {
	return c.UpdateFleetServiceTypes([]int64{fst.Id.Get()}, fst)
}

// UpdateFleetServiceTypes updates existing fleet.service.type records.
// All records (represented by ids) will be updated by fst values.
func (c *Client) UpdateFleetServiceTypes(ids []int64, fst *FleetServiceType) error {
	return c.Update(FleetServiceTypeModel, ids, fst)
}

// DeleteFleetServiceType deletes an existing fleet.service.type record.
func (c *Client) DeleteFleetServiceType(id int64) error {
	return c.DeleteFleetServiceTypes([]int64{id})
}

// DeleteFleetServiceTypes deletes existing fleet.service.type records.
func (c *Client) DeleteFleetServiceTypes(ids []int64) error {
	return c.Delete(FleetServiceTypeModel, ids)
}

// GetFleetServiceType gets fleet.service.type existing record.
func (c *Client) GetFleetServiceType(id int64) (*FleetServiceType, error) {
	fsts, err := c.GetFleetServiceTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if fsts != nil && len(*fsts) > 0 {
		return &((*fsts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.service.type not found", id)
}

// GetFleetServiceTypes gets fleet.service.type existing records.
func (c *Client) GetFleetServiceTypes(ids []int64) (*FleetServiceTypes, error) {
	fsts := &FleetServiceTypes{}
	if err := c.Read(FleetServiceTypeModel, ids, nil, fsts); err != nil {
		return nil, err
	}
	return fsts, nil
}

// FindFleetServiceType finds fleet.service.type record by querying it with criteria.
func (c *Client) FindFleetServiceType(criteria *Criteria) (*FleetServiceType, error) {
	fsts := &FleetServiceTypes{}
	if err := c.SearchRead(FleetServiceTypeModel, criteria, NewOptions().Limit(1), fsts); err != nil {
		return nil, err
	}
	if fsts != nil && len(*fsts) > 0 {
		return &((*fsts)[0]), nil
	}
	return nil, fmt.Errorf("fleet.service.type was not found with criteria %v", criteria)
}

// FindFleetServiceTypes finds fleet.service.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetServiceTypes(criteria *Criteria, options *Options) (*FleetServiceTypes, error) {
	fsts := &FleetServiceTypes{}
	if err := c.SearchRead(FleetServiceTypeModel, criteria, options, fsts); err != nil {
		return nil, err
	}
	return fsts, nil
}

// FindFleetServiceTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetServiceTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FleetServiceTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetServiceTypeId finds record id by querying it with criteria.
func (c *Client) FindFleetServiceTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetServiceTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.service.type was not found with criteria %v and options %v", criteria, options)
}
