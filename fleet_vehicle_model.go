package odoo

import (
	"fmt"
)

// FleetVehicleModel represents fleet.vehicle.model model.
type FleetVehicleModel struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	BrandId     *Many2One `xmlrpc:"brand_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Image128    *String   `xmlrpc:"image_128,omptempty"`
	ManagerId   *Many2One `xmlrpc:"manager_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Vendors     *Relation `xmlrpc:"vendors,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// FleetVehicleModels represents array of fleet.vehicle.model model.
type FleetVehicleModels []FleetVehicleModel

// FleetVehicleModelModel is the odoo model name.
const FleetVehicleModelModel = "fleet.vehicle.model"

// Many2One convert FleetVehicleModel to *Many2One.
func (fvm *FleetVehicleModel) Many2One() *Many2One {
	return NewMany2One(fvm.Id.Get(), "")
}

// CreateFleetVehicleModel creates a new fleet.vehicle.model model and returns its id.
func (c *Client) CreateFleetVehicleModel(fvm *FleetVehicleModel) (int64, error) {
	ids, err := c.CreateFleetVehicleModels([]*FleetVehicleModel{fvm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleModel creates a new fleet.vehicle.model model and returns its id.
func (c *Client) CreateFleetVehicleModels(fvms []*FleetVehicleModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvms {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleModelModel, vv)
}

// UpdateFleetVehicleModel updates an existing fleet.vehicle.model record.
func (c *Client) UpdateFleetVehicleModel(fvm *FleetVehicleModel) error {
	return c.UpdateFleetVehicleModels([]int64{fvm.Id.Get()}, fvm)
}

// UpdateFleetVehicleModels updates existing fleet.vehicle.model records.
// All records (represented by ids) will be updated by fvm values.
func (c *Client) UpdateFleetVehicleModels(ids []int64, fvm *FleetVehicleModel) error {
	return c.Update(FleetVehicleModelModel, ids, fvm)
}

// DeleteFleetVehicleModel deletes an existing fleet.vehicle.model record.
func (c *Client) DeleteFleetVehicleModel(id int64) error {
	return c.DeleteFleetVehicleModels([]int64{id})
}

// DeleteFleetVehicleModels deletes existing fleet.vehicle.model records.
func (c *Client) DeleteFleetVehicleModels(ids []int64) error {
	return c.Delete(FleetVehicleModelModel, ids)
}

// GetFleetVehicleModel gets fleet.vehicle.model existing record.
func (c *Client) GetFleetVehicleModel(id int64) (*FleetVehicleModel, error) {
	fvms, err := c.GetFleetVehicleModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if fvms != nil && len(*fvms) > 0 {
		return &((*fvms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.vehicle.model not found", id)
}

// GetFleetVehicleModels gets fleet.vehicle.model existing records.
func (c *Client) GetFleetVehicleModels(ids []int64) (*FleetVehicleModels, error) {
	fvms := &FleetVehicleModels{}
	if err := c.Read(FleetVehicleModelModel, ids, nil, fvms); err != nil {
		return nil, err
	}
	return fvms, nil
}

// FindFleetVehicleModel finds fleet.vehicle.model record by querying it with criteria.
func (c *Client) FindFleetVehicleModel(criteria *Criteria) (*FleetVehicleModel, error) {
	fvms := &FleetVehicleModels{}
	if err := c.SearchRead(FleetVehicleModelModel, criteria, NewOptions().Limit(1), fvms); err != nil {
		return nil, err
	}
	if fvms != nil && len(*fvms) > 0 {
		return &((*fvms)[0]), nil
	}
	return nil, fmt.Errorf("fleet.vehicle.model was not found with criteria %v", criteria)
}

// FindFleetVehicleModels finds fleet.vehicle.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleModels(criteria *Criteria, options *Options) (*FleetVehicleModels, error) {
	fvms := &FleetVehicleModels{}
	if err := c.SearchRead(FleetVehicleModelModel, criteria, options, fvms); err != nil {
		return nil, err
	}
	return fvms, nil
}

// FindFleetVehicleModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FleetVehicleModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetVehicleModelId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.vehicle.model was not found with criteria %v and options %v", criteria, options)
}
