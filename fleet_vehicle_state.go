package odoo

import (
	"fmt"
)

// FleetVehicleState represents fleet.vehicle.state model.
type FleetVehicleState struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// FleetVehicleStates represents array of fleet.vehicle.state model.
type FleetVehicleStates []FleetVehicleState

// FleetVehicleStateModel is the odoo model name.
const FleetVehicleStateModel = "fleet.vehicle.state"

// Many2One convert FleetVehicleState to *Many2One.
func (fvs *FleetVehicleState) Many2One() *Many2One {
	return NewMany2One(fvs.Id.Get(), "")
}

// CreateFleetVehicleState creates a new fleet.vehicle.state model and returns its id.
func (c *Client) CreateFleetVehicleState(fvs *FleetVehicleState) (int64, error) {
	ids, err := c.CreateFleetVehicleStates([]*FleetVehicleState{fvs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleState creates a new fleet.vehicle.state model and returns its id.
func (c *Client) CreateFleetVehicleStates(fvss []*FleetVehicleState) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvss {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleStateModel, vv)
}

// UpdateFleetVehicleState updates an existing fleet.vehicle.state record.
func (c *Client) UpdateFleetVehicleState(fvs *FleetVehicleState) error {
	return c.UpdateFleetVehicleStates([]int64{fvs.Id.Get()}, fvs)
}

// UpdateFleetVehicleStates updates existing fleet.vehicle.state records.
// All records (represented by ids) will be updated by fvs values.
func (c *Client) UpdateFleetVehicleStates(ids []int64, fvs *FleetVehicleState) error {
	return c.Update(FleetVehicleStateModel, ids, fvs)
}

// DeleteFleetVehicleState deletes an existing fleet.vehicle.state record.
func (c *Client) DeleteFleetVehicleState(id int64) error {
	return c.DeleteFleetVehicleStates([]int64{id})
}

// DeleteFleetVehicleStates deletes existing fleet.vehicle.state records.
func (c *Client) DeleteFleetVehicleStates(ids []int64) error {
	return c.Delete(FleetVehicleStateModel, ids)
}

// GetFleetVehicleState gets fleet.vehicle.state existing record.
func (c *Client) GetFleetVehicleState(id int64) (*FleetVehicleState, error) {
	fvss, err := c.GetFleetVehicleStates([]int64{id})
	if err != nil {
		return nil, err
	}
	if fvss != nil && len(*fvss) > 0 {
		return &((*fvss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.vehicle.state not found", id)
}

// GetFleetVehicleStates gets fleet.vehicle.state existing records.
func (c *Client) GetFleetVehicleStates(ids []int64) (*FleetVehicleStates, error) {
	fvss := &FleetVehicleStates{}
	if err := c.Read(FleetVehicleStateModel, ids, nil, fvss); err != nil {
		return nil, err
	}
	return fvss, nil
}

// FindFleetVehicleState finds fleet.vehicle.state record by querying it with criteria.
func (c *Client) FindFleetVehicleState(criteria *Criteria) (*FleetVehicleState, error) {
	fvss := &FleetVehicleStates{}
	if err := c.SearchRead(FleetVehicleStateModel, criteria, NewOptions().Limit(1), fvss); err != nil {
		return nil, err
	}
	if fvss != nil && len(*fvss) > 0 {
		return &((*fvss)[0]), nil
	}
	return nil, fmt.Errorf("fleet.vehicle.state was not found with criteria %v", criteria)
}

// FindFleetVehicleStates finds fleet.vehicle.state records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleStates(criteria *Criteria, options *Options) (*FleetVehicleStates, error) {
	fvss := &FleetVehicleStates{}
	if err := c.SearchRead(FleetVehicleStateModel, criteria, options, fvss); err != nil {
		return nil, err
	}
	return fvss, nil
}

// FindFleetVehicleStateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleStateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FleetVehicleStateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetVehicleStateId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleStateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleStateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.vehicle.state was not found with criteria %v and options %v", criteria, options)
}
