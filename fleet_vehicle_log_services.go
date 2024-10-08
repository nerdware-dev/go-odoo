package odoo

// FleetVehicleLogServices represents fleet.vehicle.log.services model.
type FleetVehicleLogServices struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	Amount        *Float     `xmlrpc:"amount,omitempty"`
	AutoGenerated *Bool      `xmlrpc:"auto_generated,omitempty"`
	CompanyId     *Many2One  `xmlrpc:"company_id,omitempty"`
	ContractId    *Many2One  `xmlrpc:"contract_id,omitempty"`
	CostAmount    *Float     `xmlrpc:"cost_amount,omitempty"`
	CostId        *Many2One  `xmlrpc:"cost_id,omitempty"`
	CostIds       *Relation  `xmlrpc:"cost_ids,omitempty"`
	CostSubtypeId *Many2One  `xmlrpc:"cost_subtype_id,omitempty"`
	CostType      *Selection `xmlrpc:"cost_type,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId    *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date          *Time      `xmlrpc:"date,omitempty"`
	Description   *String    `xmlrpc:"description,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	InvRef        *String    `xmlrpc:"inv_ref,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	Notes         *String    `xmlrpc:"notes,omitempty"`
	Odometer      *Float     `xmlrpc:"odometer,omitempty"`
	OdometerId    *Many2One  `xmlrpc:"odometer_id,omitempty"`
	OdometerUnit  *Selection `xmlrpc:"odometer_unit,omitempty"`
	ParentId      *Many2One  `xmlrpc:"parent_id,omitempty"`
	PurchaserId   *Many2One  `xmlrpc:"purchaser_id,omitempty"`
	VehicleId     *Many2One  `xmlrpc:"vehicle_id,omitempty"`
	VendorId      *Many2One  `xmlrpc:"vendor_id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// FleetVehicleLogServicess represents array of fleet.vehicle.log.services model.
type FleetVehicleLogServicess []FleetVehicleLogServices

// FleetVehicleLogServicesModel is the odoo model name.
const FleetVehicleLogServicesModel = "fleet.vehicle.log.services"

// Many2One convert FleetVehicleLogServices to *Many2One.
func (fvls *FleetVehicleLogServices) Many2One() *Many2One {
	return NewMany2One(fvls.Id.Get(), "")
}

// CreateFleetVehicleLogServices creates a new fleet.vehicle.log.services model and returns its id.
func (c *Client) CreateFleetVehicleLogServices(fvls *FleetVehicleLogServices) (int64, error) {
	ids, err := c.CreateFleetVehicleLogServicess([]*FleetVehicleLogServices{fvls})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleLogServices creates a new fleet.vehicle.log.services model and returns its id.
func (c *Client) CreateFleetVehicleLogServicess(fvlss []*FleetVehicleLogServices) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvlss {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleLogServicesModel, vv, nil)
}

// UpdateFleetVehicleLogServices updates an existing fleet.vehicle.log.services record.
func (c *Client) UpdateFleetVehicleLogServices(fvls *FleetVehicleLogServices) error {
	return c.UpdateFleetVehicleLogServicess([]int64{fvls.Id.Get()}, fvls)
}

// UpdateFleetVehicleLogServicess updates existing fleet.vehicle.log.services records.
// All records (represented by ids) will be updated by fvls values.
func (c *Client) UpdateFleetVehicleLogServicess(ids []int64, fvls *FleetVehicleLogServices) error {
	return c.Update(FleetVehicleLogServicesModel, ids, fvls, nil)
}

// DeleteFleetVehicleLogServices deletes an existing fleet.vehicle.log.services record.
func (c *Client) DeleteFleetVehicleLogServices(id int64) error {
	return c.DeleteFleetVehicleLogServicess([]int64{id})
}

// DeleteFleetVehicleLogServicess deletes existing fleet.vehicle.log.services records.
func (c *Client) DeleteFleetVehicleLogServicess(ids []int64) error {
	return c.Delete(FleetVehicleLogServicesModel, ids)
}

// GetFleetVehicleLogServices gets fleet.vehicle.log.services existing record.
func (c *Client) GetFleetVehicleLogServices(id int64) (*FleetVehicleLogServices, error) {
	fvlss, err := c.GetFleetVehicleLogServicess([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fvlss)[0]), nil
}

// GetFleetVehicleLogServicess gets fleet.vehicle.log.services existing records.
func (c *Client) GetFleetVehicleLogServicess(ids []int64) (*FleetVehicleLogServicess, error) {
	fvlss := &FleetVehicleLogServicess{}
	if err := c.Read(FleetVehicleLogServicesModel, ids, nil, fvlss); err != nil {
		return nil, err
	}
	return fvlss, nil
}

// FindFleetVehicleLogServices finds fleet.vehicle.log.services record by querying it with criteria.
func (c *Client) FindFleetVehicleLogServices(criteria *Criteria) (*FleetVehicleLogServices, error) {
	fvlss := &FleetVehicleLogServicess{}
	if err := c.SearchRead(FleetVehicleLogServicesModel, criteria, NewOptions().Limit(1), fvlss); err != nil {
		return nil, err
	}
	return &((*fvlss)[0]), nil
}

// FindFleetVehicleLogServicess finds fleet.vehicle.log.services records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleLogServicess(criteria *Criteria, options *Options) (*FleetVehicleLogServicess, error) {
	fvlss := &FleetVehicleLogServicess{}
	if err := c.SearchRead(FleetVehicleLogServicesModel, criteria, options, fvlss); err != nil {
		return nil, err
	}
	return fvlss, nil
}

// FindFleetVehicleLogServicesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleLogServicesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FleetVehicleLogServicesModel, criteria, options)
}

// FindFleetVehicleLogServicesId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleLogServicesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleLogServicesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
