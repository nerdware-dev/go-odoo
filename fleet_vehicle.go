package odoo

// FleetVehicle represents fleet.vehicle model.
type FleetVehicle struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AcquisitionDate             *Time      `xmlrpc:"acquisition_date,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BrandId                     *Many2One  `xmlrpc:"brand_id,omitempty"`
	CarValue                    *Float     `xmlrpc:"car_value,omitempty"`
	Co2                         *Float     `xmlrpc:"co2,omitempty"`
	Color                       *String    `xmlrpc:"color,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ContractCount               *Int       `xmlrpc:"contract_count,omitempty"`
	ContractRenewalDueSoon      *Bool      `xmlrpc:"contract_renewal_due_soon,omitempty"`
	ContractRenewalName         *String    `xmlrpc:"contract_renewal_name,omitempty"`
	ContractRenewalOverdue      *Bool      `xmlrpc:"contract_renewal_overdue,omitempty"`
	ContractRenewalTotal        *String    `xmlrpc:"contract_renewal_total,omitempty"`
	CostCount                   *Int       `xmlrpc:"cost_count,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Doors                       *Int       `xmlrpc:"doors,omitempty"`
	DriverId                    *Many2One  `xmlrpc:"driver_id,omitempty"`
	FirstContractDate           *Time      `xmlrpc:"first_contract_date,omitempty"`
	FuelLogsCount               *Int       `xmlrpc:"fuel_logs_count,omitempty"`
	FuelType                    *Selection `xmlrpc:"fuel_type,omitempty"`
	FutureDriverId              *Many2One  `xmlrpc:"future_driver_id,omitempty"`
	HistoryCount                *Int       `xmlrpc:"history_count,omitempty"`
	Horsepower                  *Int       `xmlrpc:"horsepower,omitempty"`
	HorsepowerTax               *Float     `xmlrpc:"horsepower_tax,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	Image128                    *String    `xmlrpc:"image_128,omitempty"`
	LicensePlate                *String    `xmlrpc:"license_plate,omitempty"`
	Location                    *String    `xmlrpc:"location,omitempty"`
	LogContracts                *Relation  `xmlrpc:"log_contracts,omitempty"`
	LogDrivers                  *Relation  `xmlrpc:"log_drivers,omitempty"`
	LogFuel                     *Relation  `xmlrpc:"log_fuel,omitempty"`
	LogServices                 *Relation  `xmlrpc:"log_services,omitempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	ModelId                     *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelYear                   *String    `xmlrpc:"model_year,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NetCarValue                 *Float     `xmlrpc:"net_car_value,omitempty"`
	NextAssignationDate         *Time      `xmlrpc:"next_assignation_date,omitempty"`
	Odometer                    *Float     `xmlrpc:"odometer,omitempty"`
	OdometerCount               *Int       `xmlrpc:"odometer_count,omitempty"`
	OdometerUnit                *Selection `xmlrpc:"odometer_unit,omitempty"`
	PlanToChangeCar             *Bool      `xmlrpc:"plan_to_change_car,omitempty"`
	Power                       *Int       `xmlrpc:"power,omitempty"`
	ResidualValue               *Float     `xmlrpc:"residual_value,omitempty"`
	Seats                       *Int       `xmlrpc:"seats,omitempty"`
	ServiceCount                *Int       `xmlrpc:"service_count,omitempty"`
	StateId                     *Many2One  `xmlrpc:"state_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	Transmission                *Selection `xmlrpc:"transmission,omitempty"`
	VinSn                       *String    `xmlrpc:"vin_sn,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// FleetVehicles represents array of fleet.vehicle model.
type FleetVehicles []FleetVehicle

// FleetVehicleModel is the odoo model name.
const FleetVehicleModel = "fleet.vehicle"

// Many2One convert FleetVehicle to *Many2One.
func (fv *FleetVehicle) Many2One() *Many2One {
	return NewMany2One(fv.Id.Get(), "")
}

// CreateFleetVehicle creates a new fleet.vehicle model and returns its id.
func (c *Client) CreateFleetVehicle(fv *FleetVehicle) (int64, error) {
	ids, err := c.CreateFleetVehicles([]*FleetVehicle{fv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicle creates a new fleet.vehicle model and returns its id.
func (c *Client) CreateFleetVehicles(fvs []*FleetVehicle) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvs {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleModel, vv, nil)
}

// UpdateFleetVehicle updates an existing fleet.vehicle record.
func (c *Client) UpdateFleetVehicle(fv *FleetVehicle) error {
	return c.UpdateFleetVehicles([]int64{fv.Id.Get()}, fv)
}

// UpdateFleetVehicles updates existing fleet.vehicle records.
// All records (represented by ids) will be updated by fv values.
func (c *Client) UpdateFleetVehicles(ids []int64, fv *FleetVehicle) error {
	return c.Update(FleetVehicleModel, ids, fv, nil)
}

// DeleteFleetVehicle deletes an existing fleet.vehicle record.
func (c *Client) DeleteFleetVehicle(id int64) error {
	return c.DeleteFleetVehicles([]int64{id})
}

// DeleteFleetVehicles deletes existing fleet.vehicle records.
func (c *Client) DeleteFleetVehicles(ids []int64) error {
	return c.Delete(FleetVehicleModel, ids)
}

// GetFleetVehicle gets fleet.vehicle existing record.
func (c *Client) GetFleetVehicle(id int64) (*FleetVehicle, error) {
	fvs, err := c.GetFleetVehicles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fvs)[0]), nil
}

// GetFleetVehicles gets fleet.vehicle existing records.
func (c *Client) GetFleetVehicles(ids []int64) (*FleetVehicles, error) {
	fvs := &FleetVehicles{}
	if err := c.Read(FleetVehicleModel, ids, nil, fvs); err != nil {
		return nil, err
	}
	return fvs, nil
}

// FindFleetVehicle finds fleet.vehicle record by querying it with criteria.
func (c *Client) FindFleetVehicle(criteria *Criteria) (*FleetVehicle, error) {
	fvs := &FleetVehicles{}
	if err := c.SearchRead(FleetVehicleModel, criteria, NewOptions().Limit(1), fvs); err != nil {
		return nil, err
	}
	return &((*fvs)[0]), nil
}

// FindFleetVehicles finds fleet.vehicle records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicles(criteria *Criteria, options *Options) (*FleetVehicles, error) {
	fvs := &FleetVehicles{}
	if err := c.SearchRead(FleetVehicleModel, criteria, options, fvs); err != nil {
		return nil, err
	}
	return fvs, nil
}

// FindFleetVehicleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FleetVehicleModel, criteria, options)
}

// FindFleetVehicleId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
