package odoo

import (
	"fmt"
)

// FleetVehicleTag represents fleet.vehicle.tag model.
type FleetVehicleTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// FleetVehicleTags represents array of fleet.vehicle.tag model.
type FleetVehicleTags []FleetVehicleTag

// FleetVehicleTagModel is the odoo model name.
const FleetVehicleTagModel = "fleet.vehicle.tag"

// Many2One convert FleetVehicleTag to *Many2One.
func (fvt *FleetVehicleTag) Many2One() *Many2One {
	return NewMany2One(fvt.Id.Get(), "")
}

// CreateFleetVehicleTag creates a new fleet.vehicle.tag model and returns its id.
func (c *Client) CreateFleetVehicleTag(fvt *FleetVehicleTag) (int64, error) {
	ids, err := c.CreateFleetVehicleTags([]*FleetVehicleTag{fvt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleTag creates a new fleet.vehicle.tag model and returns its id.
func (c *Client) CreateFleetVehicleTags(fvts []*FleetVehicleTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvts {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleTagModel, vv)
}

// UpdateFleetVehicleTag updates an existing fleet.vehicle.tag record.
func (c *Client) UpdateFleetVehicleTag(fvt *FleetVehicleTag) error {
	return c.UpdateFleetVehicleTags([]int64{fvt.Id.Get()}, fvt)
}

// UpdateFleetVehicleTags updates existing fleet.vehicle.tag records.
// All records (represented by ids) will be updated by fvt values.
func (c *Client) UpdateFleetVehicleTags(ids []int64, fvt *FleetVehicleTag) error {
	return c.Update(FleetVehicleTagModel, ids, fvt)
}

// DeleteFleetVehicleTag deletes an existing fleet.vehicle.tag record.
func (c *Client) DeleteFleetVehicleTag(id int64) error {
	return c.DeleteFleetVehicleTags([]int64{id})
}

// DeleteFleetVehicleTags deletes existing fleet.vehicle.tag records.
func (c *Client) DeleteFleetVehicleTags(ids []int64) error {
	return c.Delete(FleetVehicleTagModel, ids)
}

// GetFleetVehicleTag gets fleet.vehicle.tag existing record.
func (c *Client) GetFleetVehicleTag(id int64) (*FleetVehicleTag, error) {
	fvts, err := c.GetFleetVehicleTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if fvts != nil && len(*fvts) > 0 {
		return &((*fvts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.vehicle.tag not found", id)
}

// GetFleetVehicleTags gets fleet.vehicle.tag existing records.
func (c *Client) GetFleetVehicleTags(ids []int64) (*FleetVehicleTags, error) {
	fvts := &FleetVehicleTags{}
	if err := c.Read(FleetVehicleTagModel, ids, nil, fvts); err != nil {
		return nil, err
	}
	return fvts, nil
}

// FindFleetVehicleTag finds fleet.vehicle.tag record by querying it with criteria.
func (c *Client) FindFleetVehicleTag(criteria *Criteria) (*FleetVehicleTag, error) {
	fvts := &FleetVehicleTags{}
	if err := c.SearchRead(FleetVehicleTagModel, criteria, NewOptions().Limit(1), fvts); err != nil {
		return nil, err
	}
	if fvts != nil && len(*fvts) > 0 {
		return &((*fvts)[0]), nil
	}
	return nil, fmt.Errorf("fleet.vehicle.tag was not found with criteria %v", criteria)
}

// FindFleetVehicleTags finds fleet.vehicle.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleTags(criteria *Criteria, options *Options) (*FleetVehicleTags, error) {
	fvts := &FleetVehicleTags{}
	if err := c.SearchRead(FleetVehicleTagModel, criteria, options, fvts); err != nil {
		return nil, err
	}
	return fvts, nil
}

// FindFleetVehicleTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FleetVehicleTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetVehicleTagId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.vehicle.tag was not found with criteria %v and options %v", criteria, options)
}
