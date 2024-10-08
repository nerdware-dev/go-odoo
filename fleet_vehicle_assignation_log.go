package odoo

// FleetVehicleAssignationLog represents fleet.vehicle.assignation.log model.
type FleetVehicleAssignationLog struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	AttachmentNumber *Int      `xmlrpc:"attachment_number,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DateEnd          *Time     `xmlrpc:"date_end,omitempty"`
	DateStart        *Time     `xmlrpc:"date_start,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	DriverId         *Many2One `xmlrpc:"driver_id,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	VehicleId        *Many2One `xmlrpc:"vehicle_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// FleetVehicleAssignationLogs represents array of fleet.vehicle.assignation.log model.
type FleetVehicleAssignationLogs []FleetVehicleAssignationLog

// FleetVehicleAssignationLogModel is the odoo model name.
const FleetVehicleAssignationLogModel = "fleet.vehicle.assignation.log"

// Many2One convert FleetVehicleAssignationLog to *Many2One.
func (fval *FleetVehicleAssignationLog) Many2One() *Many2One {
	return NewMany2One(fval.Id.Get(), "")
}

// CreateFleetVehicleAssignationLog creates a new fleet.vehicle.assignation.log model and returns its id.
func (c *Client) CreateFleetVehicleAssignationLog(fval *FleetVehicleAssignationLog) (int64, error) {
	ids, err := c.CreateFleetVehicleAssignationLogs([]*FleetVehicleAssignationLog{fval})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleAssignationLog creates a new fleet.vehicle.assignation.log model and returns its id.
func (c *Client) CreateFleetVehicleAssignationLogs(fvals []*FleetVehicleAssignationLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvals {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleAssignationLogModel, vv, nil)
}

// UpdateFleetVehicleAssignationLog updates an existing fleet.vehicle.assignation.log record.
func (c *Client) UpdateFleetVehicleAssignationLog(fval *FleetVehicleAssignationLog) error {
	return c.UpdateFleetVehicleAssignationLogs([]int64{fval.Id.Get()}, fval)
}

// UpdateFleetVehicleAssignationLogs updates existing fleet.vehicle.assignation.log records.
// All records (represented by ids) will be updated by fval values.
func (c *Client) UpdateFleetVehicleAssignationLogs(ids []int64, fval *FleetVehicleAssignationLog) error {
	return c.Update(FleetVehicleAssignationLogModel, ids, fval, nil)
}

// DeleteFleetVehicleAssignationLog deletes an existing fleet.vehicle.assignation.log record.
func (c *Client) DeleteFleetVehicleAssignationLog(id int64) error {
	return c.DeleteFleetVehicleAssignationLogs([]int64{id})
}

// DeleteFleetVehicleAssignationLogs deletes existing fleet.vehicle.assignation.log records.
func (c *Client) DeleteFleetVehicleAssignationLogs(ids []int64) error {
	return c.Delete(FleetVehicleAssignationLogModel, ids)
}

// GetFleetVehicleAssignationLog gets fleet.vehicle.assignation.log existing record.
func (c *Client) GetFleetVehicleAssignationLog(id int64) (*FleetVehicleAssignationLog, error) {
	fvals, err := c.GetFleetVehicleAssignationLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fvals)[0]), nil
}

// GetFleetVehicleAssignationLogs gets fleet.vehicle.assignation.log existing records.
func (c *Client) GetFleetVehicleAssignationLogs(ids []int64) (*FleetVehicleAssignationLogs, error) {
	fvals := &FleetVehicleAssignationLogs{}
	if err := c.Read(FleetVehicleAssignationLogModel, ids, nil, fvals); err != nil {
		return nil, err
	}
	return fvals, nil
}

// FindFleetVehicleAssignationLog finds fleet.vehicle.assignation.log record by querying it with criteria.
func (c *Client) FindFleetVehicleAssignationLog(criteria *Criteria) (*FleetVehicleAssignationLog, error) {
	fvals := &FleetVehicleAssignationLogs{}
	if err := c.SearchRead(FleetVehicleAssignationLogModel, criteria, NewOptions().Limit(1), fvals); err != nil {
		return nil, err
	}
	return &((*fvals)[0]), nil
}

// FindFleetVehicleAssignationLogs finds fleet.vehicle.assignation.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleAssignationLogs(criteria *Criteria, options *Options) (*FleetVehicleAssignationLogs, error) {
	fvals := &FleetVehicleAssignationLogs{}
	if err := c.SearchRead(FleetVehicleAssignationLogModel, criteria, options, fvals); err != nil {
		return nil, err
	}
	return fvals, nil
}

// FindFleetVehicleAssignationLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleAssignationLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FleetVehicleAssignationLogModel, criteria, options)
}

// FindFleetVehicleAssignationLogId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleAssignationLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleAssignationLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
