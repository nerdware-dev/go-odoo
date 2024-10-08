package odoo

// FleetVehicleTag represents fleet.vehicle.tag model.
type FleetVehicleTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(FleetVehicleTagModel, vv, nil)
}

// UpdateFleetVehicleTag updates an existing fleet.vehicle.tag record.
func (c *Client) UpdateFleetVehicleTag(fvt *FleetVehicleTag) error {
	return c.UpdateFleetVehicleTags([]int64{fvt.Id.Get()}, fvt)
}

// UpdateFleetVehicleTags updates existing fleet.vehicle.tag records.
// All records (represented by ids) will be updated by fvt values.
func (c *Client) UpdateFleetVehicleTags(ids []int64, fvt *FleetVehicleTag) error {
	return c.Update(FleetVehicleTagModel, ids, fvt, nil)
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
	return &((*fvts)[0]), nil
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
	return &((*fvts)[0]), nil
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
	return c.Search(FleetVehicleTagModel, criteria, options)
}

// FindFleetVehicleTagId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
