package odoo

// FleetVehicleOdometer represents fleet.vehicle.odometer model.
type FleetVehicleOdometer struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date        *Time      `xmlrpc:"date,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	DriverId    *Many2One  `xmlrpc:"driver_id,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Unit        *Selection `xmlrpc:"unit,omitempty"`
	Value       *Float     `xmlrpc:"value,omitempty"`
	VehicleId   *Many2One  `xmlrpc:"vehicle_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// FleetVehicleOdometers represents array of fleet.vehicle.odometer model.
type FleetVehicleOdometers []FleetVehicleOdometer

// FleetVehicleOdometerModel is the odoo model name.
const FleetVehicleOdometerModel = "fleet.vehicle.odometer"

// Many2One convert FleetVehicleOdometer to *Many2One.
func (fvo *FleetVehicleOdometer) Many2One() *Many2One {
	return NewMany2One(fvo.Id.Get(), "")
}

// CreateFleetVehicleOdometer creates a new fleet.vehicle.odometer model and returns its id.
func (c *Client) CreateFleetVehicleOdometer(fvo *FleetVehicleOdometer) (int64, error) {
	ids, err := c.CreateFleetVehicleOdometers([]*FleetVehicleOdometer{fvo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleOdometer creates a new fleet.vehicle.odometer model and returns its id.
func (c *Client) CreateFleetVehicleOdometers(fvos []*FleetVehicleOdometer) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvos {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleOdometerModel, vv, nil)
}

// UpdateFleetVehicleOdometer updates an existing fleet.vehicle.odometer record.
func (c *Client) UpdateFleetVehicleOdometer(fvo *FleetVehicleOdometer) error {
	return c.UpdateFleetVehicleOdometers([]int64{fvo.Id.Get()}, fvo)
}

// UpdateFleetVehicleOdometers updates existing fleet.vehicle.odometer records.
// All records (represented by ids) will be updated by fvo values.
func (c *Client) UpdateFleetVehicleOdometers(ids []int64, fvo *FleetVehicleOdometer) error {
	return c.Update(FleetVehicleOdometerModel, ids, fvo, nil)
}

// DeleteFleetVehicleOdometer deletes an existing fleet.vehicle.odometer record.
func (c *Client) DeleteFleetVehicleOdometer(id int64) error {
	return c.DeleteFleetVehicleOdometers([]int64{id})
}

// DeleteFleetVehicleOdometers deletes existing fleet.vehicle.odometer records.
func (c *Client) DeleteFleetVehicleOdometers(ids []int64) error {
	return c.Delete(FleetVehicleOdometerModel, ids)
}

// GetFleetVehicleOdometer gets fleet.vehicle.odometer existing record.
func (c *Client) GetFleetVehicleOdometer(id int64) (*FleetVehicleOdometer, error) {
	fvos, err := c.GetFleetVehicleOdometers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fvos)[0]), nil
}

// GetFleetVehicleOdometers gets fleet.vehicle.odometer existing records.
func (c *Client) GetFleetVehicleOdometers(ids []int64) (*FleetVehicleOdometers, error) {
	fvos := &FleetVehicleOdometers{}
	if err := c.Read(FleetVehicleOdometerModel, ids, nil, fvos); err != nil {
		return nil, err
	}
	return fvos, nil
}

// FindFleetVehicleOdometer finds fleet.vehicle.odometer record by querying it with criteria.
func (c *Client) FindFleetVehicleOdometer(criteria *Criteria) (*FleetVehicleOdometer, error) {
	fvos := &FleetVehicleOdometers{}
	if err := c.SearchRead(FleetVehicleOdometerModel, criteria, NewOptions().Limit(1), fvos); err != nil {
		return nil, err
	}
	return &((*fvos)[0]), nil
}

// FindFleetVehicleOdometers finds fleet.vehicle.odometer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleOdometers(criteria *Criteria, options *Options) (*FleetVehicleOdometers, error) {
	fvos := &FleetVehicleOdometers{}
	if err := c.SearchRead(FleetVehicleOdometerModel, criteria, options, fvos); err != nil {
		return nil, err
	}
	return fvos, nil
}

// FindFleetVehicleOdometerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleOdometerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FleetVehicleOdometerModel, criteria, options)
}

// FindFleetVehicleOdometerId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleOdometerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleOdometerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
