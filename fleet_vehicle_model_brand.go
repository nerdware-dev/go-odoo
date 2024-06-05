package odoo

import (
	"fmt"
)

// FleetVehicleModelBrand represents fleet.vehicle.model.brand model.
type FleetVehicleModelBrand struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Image128    *String   `xmlrpc:"image_128,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// FleetVehicleModelBrands represents array of fleet.vehicle.model.brand model.
type FleetVehicleModelBrands []FleetVehicleModelBrand

// FleetVehicleModelBrandModel is the odoo model name.
const FleetVehicleModelBrandModel = "fleet.vehicle.model.brand"

// Many2One convert FleetVehicleModelBrand to *Many2One.
func (fvmb *FleetVehicleModelBrand) Many2One() *Many2One {
	return NewMany2One(fvmb.Id.Get(), "")
}

// CreateFleetVehicleModelBrand creates a new fleet.vehicle.model.brand model and returns its id.
func (c *Client) CreateFleetVehicleModelBrand(fvmb *FleetVehicleModelBrand) (int64, error) {
	ids, err := c.CreateFleetVehicleModelBrands([]*FleetVehicleModelBrand{fvmb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleModelBrand creates a new fleet.vehicle.model.brand model and returns its id.
func (c *Client) CreateFleetVehicleModelBrands(fvmbs []*FleetVehicleModelBrand) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvmbs {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleModelBrandModel, vv)
}

// UpdateFleetVehicleModelBrand updates an existing fleet.vehicle.model.brand record.
func (c *Client) UpdateFleetVehicleModelBrand(fvmb *FleetVehicleModelBrand) error {
	return c.UpdateFleetVehicleModelBrands([]int64{fvmb.Id.Get()}, fvmb)
}

// UpdateFleetVehicleModelBrands updates existing fleet.vehicle.model.brand records.
// All records (represented by ids) will be updated by fvmb values.
func (c *Client) UpdateFleetVehicleModelBrands(ids []int64, fvmb *FleetVehicleModelBrand) error {
	return c.Update(FleetVehicleModelBrandModel, ids, fvmb)
}

// DeleteFleetVehicleModelBrand deletes an existing fleet.vehicle.model.brand record.
func (c *Client) DeleteFleetVehicleModelBrand(id int64) error {
	return c.DeleteFleetVehicleModelBrands([]int64{id})
}

// DeleteFleetVehicleModelBrands deletes existing fleet.vehicle.model.brand records.
func (c *Client) DeleteFleetVehicleModelBrands(ids []int64) error {
	return c.Delete(FleetVehicleModelBrandModel, ids)
}

// GetFleetVehicleModelBrand gets fleet.vehicle.model.brand existing record.
func (c *Client) GetFleetVehicleModelBrand(id int64) (*FleetVehicleModelBrand, error) {
	fvmbs, err := c.GetFleetVehicleModelBrands([]int64{id})
	if err != nil {
		return nil, err
	}
	if fvmbs != nil && len(*fvmbs) > 0 {
		return &((*fvmbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.vehicle.model.brand not found", id)
}

// GetFleetVehicleModelBrands gets fleet.vehicle.model.brand existing records.
func (c *Client) GetFleetVehicleModelBrands(ids []int64) (*FleetVehicleModelBrands, error) {
	fvmbs := &FleetVehicleModelBrands{}
	if err := c.Read(FleetVehicleModelBrandModel, ids, nil, fvmbs); err != nil {
		return nil, err
	}
	return fvmbs, nil
}

// FindFleetVehicleModelBrand finds fleet.vehicle.model.brand record by querying it with criteria.
func (c *Client) FindFleetVehicleModelBrand(criteria *Criteria) (*FleetVehicleModelBrand, error) {
	fvmbs := &FleetVehicleModelBrands{}
	if err := c.SearchRead(FleetVehicleModelBrandModel, criteria, NewOptions().Limit(1), fvmbs); err != nil {
		return nil, err
	}
	if fvmbs != nil && len(*fvmbs) > 0 {
		return &((*fvmbs)[0]), nil
	}
	return nil, fmt.Errorf("fleet.vehicle.model.brand was not found with criteria %v", criteria)
}

// FindFleetVehicleModelBrands finds fleet.vehicle.model.brand records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleModelBrands(criteria *Criteria, options *Options) (*FleetVehicleModelBrands, error) {
	fvmbs := &FleetVehicleModelBrands{}
	if err := c.SearchRead(FleetVehicleModelBrandModel, criteria, options, fvmbs); err != nil {
		return nil, err
	}
	return fvmbs, nil
}

// FindFleetVehicleModelBrandIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleModelBrandIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FleetVehicleModelBrandModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetVehicleModelBrandId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleModelBrandId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleModelBrandModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.vehicle.model.brand was not found with criteria %v and options %v", criteria, options)
}
