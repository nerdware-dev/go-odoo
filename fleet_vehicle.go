package odoo

// FleetVehicle represents fleet.vehicle model.
type FleetVehicle struct {
	AccountMoveIds              *Relation   `xmlrpc:"account_move_ids,omitempty" json:"account_move_ids,omitempty"`
	AcquisitionDate             *Time       `xmlrpc:"acquisition_date,omitempty" json:"acquisition_date,omitempty"`
	Active                      *Bool       `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	Avatar1024                  *String     `xmlrpc:"avatar_1024,omitempty" json:"avatar_1024,omitempty"`
	Avatar128                   *String     `xmlrpc:"avatar_128,omitempty" json:"avatar_128,omitempty"`
	Avatar1920                  *String     `xmlrpc:"avatar_1920,omitempty" json:"avatar_1920,omitempty"`
	Avatar256                   *String     `xmlrpc:"avatar_256,omitempty" json:"avatar_256,omitempty"`
	Avatar512                   *String     `xmlrpc:"avatar_512,omitempty" json:"avatar_512,omitempty"`
	BillCount                   *Int        `xmlrpc:"bill_count,omitempty" json:"bill_count,omitempty"`
	BrandId                     *Many2One   `xmlrpc:"brand_id,omitempty" json:"brand_id,omitempty"`
	CarValue                    *Float      `xmlrpc:"car_value,omitempty" json:"car_value,omitempty"`
	CategoryId                  *Many2One   `xmlrpc:"category_id,omitempty" json:"category_id,omitempty"`
	Co2                         *Float      `xmlrpc:"co2,omitempty" json:"co2,omitempty"`
	Co2Standard                 *String     `xmlrpc:"co2_standard,omitempty" json:"co2_standard,omitempty"`
	Color                       *String     `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	ContractCount               *Int        `xmlrpc:"contract_count,omitempty" json:"contract_count,omitempty"`
	ContractRenewalDueSoon      *Bool       `xmlrpc:"contract_renewal_due_soon,omitempty" json:"contract_renewal_due_soon,omitempty"`
	ContractRenewalName         *String     `xmlrpc:"contract_renewal_name,omitempty" json:"contract_renewal_name,omitempty"`
	ContractRenewalOverdue      *Bool       `xmlrpc:"contract_renewal_overdue,omitempty" json:"contract_renewal_overdue,omitempty"`
	ContractRenewalTotal        *String     `xmlrpc:"contract_renewal_total,omitempty" json:"contract_renewal_total,omitempty"`
	ContractState               *Selection  `xmlrpc:"contract_state,omitempty" json:"contract_state,omitempty"`
	CountryCode                 *String     `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CountryId                   *Many2One   `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                  *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Description                 *String     `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentCount               *Int        `xmlrpc:"document_count,omitempty" json:"document_count,omitempty"`
	DocumentsFleetSettings      *Bool       `xmlrpc:"documents_fleet_settings,omitempty" json:"documents_fleet_settings,omitempty"`
	Doors                       *Int        `xmlrpc:"doors,omitempty" json:"doors,omitempty"`
	DriverEmployeeId            *Many2One   `xmlrpc:"driver_employee_id,omitempty" json:"driver_employee_id,omitempty"`
	DriverEmployeeName          *String     `xmlrpc:"driver_employee_name,omitempty" json:"driver_employee_name,omitempty"`
	DriverId                    *Many2One   `xmlrpc:"driver_id,omitempty" json:"driver_id,omitempty"`
	ElectricAssistance          *Bool       `xmlrpc:"electric_assistance,omitempty" json:"electric_assistance,omitempty"`
	FirstContractDate           *Time       `xmlrpc:"first_contract_date,omitempty" json:"first_contract_date,omitempty"`
	FrameSize                   *Float      `xmlrpc:"frame_size,omitempty" json:"frame_size,omitempty"`
	FrameType                   *Selection  `xmlrpc:"frame_type,omitempty" json:"frame_type,omitempty"`
	FuelType                    *Selection  `xmlrpc:"fuel_type,omitempty" json:"fuel_type,omitempty"`
	FutureDriverEmployeeId      *Many2One   `xmlrpc:"future_driver_employee_id,omitempty" json:"future_driver_employee_id,omitempty"`
	FutureDriverId              *Many2One   `xmlrpc:"future_driver_id,omitempty" json:"future_driver_id,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	HistoryCount                *Int        `xmlrpc:"history_count,omitempty" json:"history_count,omitempty"`
	Horsepower                  *Int        `xmlrpc:"horsepower,omitempty" json:"horsepower,omitempty"`
	HorsepowerTax               *Float      `xmlrpc:"horsepower_tax,omitempty" json:"horsepower_tax,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Image1024                   *String     `xmlrpc:"image_1024,omitempty" json:"image_1024,omitempty"`
	Image128                    *String     `xmlrpc:"image_128,omitempty" json:"image_128,omitempty"`
	Image1920                   *String     `xmlrpc:"image_1920,omitempty" json:"image_1920,omitempty"`
	Image256                    *String     `xmlrpc:"image_256,omitempty" json:"image_256,omitempty"`
	Image512                    *String     `xmlrpc:"image_512,omitempty" json:"image_512,omitempty"`
	LicensePlate                *String     `xmlrpc:"license_plate,omitempty" json:"license_plate,omitempty"`
	Location                    *String     `xmlrpc:"location,omitempty" json:"location,omitempty"`
	LogContracts                *Relation   `xmlrpc:"log_contracts,omitempty" json:"log_contracts,omitempty"`
	LogDrivers                  *Relation   `xmlrpc:"log_drivers,omitempty" json:"log_drivers,omitempty"`
	LogServices                 *Relation   `xmlrpc:"log_services,omitempty" json:"log_services,omitempty"`
	ManagerId                   *Many2One   `xmlrpc:"manager_id,omitempty" json:"manager_id,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MobilityCard                *String     `xmlrpc:"mobility_card,omitempty" json:"mobility_card,omitempty"`
	ModelId                     *Many2One   `xmlrpc:"model_id,omitempty" json:"model_id,omitempty"`
	ModelYear                   *String     `xmlrpc:"model_year,omitempty" json:"model_year,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NetCarValue                 *Float      `xmlrpc:"net_car_value,omitempty" json:"net_car_value,omitempty"`
	NextAssignationDate         *Time       `xmlrpc:"next_assignation_date,omitempty" json:"next_assignation_date,omitempty"`
	Odometer                    *Float      `xmlrpc:"odometer,omitempty" json:"odometer,omitempty"`
	OdometerCount               *Int        `xmlrpc:"odometer_count,omitempty" json:"odometer_count,omitempty"`
	OdometerUnit                *Selection  `xmlrpc:"odometer_unit,omitempty" json:"odometer_unit,omitempty"`
	OrderDate                   *Time       `xmlrpc:"order_date,omitempty" json:"order_date,omitempty"`
	PlanToChangeBike            *Bool       `xmlrpc:"plan_to_change_bike,omitempty" json:"plan_to_change_bike,omitempty"`
	PlanToChangeCar             *Bool       `xmlrpc:"plan_to_change_car,omitempty" json:"plan_to_change_car,omitempty"`
	Power                       *Int        `xmlrpc:"power,omitempty" json:"power,omitempty"`
	RateIds                     *Relation   `xmlrpc:"rate_ids,omitempty" json:"rate_ids,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	ResidualValue               *Float      `xmlrpc:"residual_value,omitempty" json:"residual_value,omitempty"`
	Seats                       *Int        `xmlrpc:"seats,omitempty" json:"seats,omitempty"`
	ServiceActivity             *Selection  `xmlrpc:"service_activity,omitempty" json:"service_activity,omitempty"`
	ServiceCount                *Int        `xmlrpc:"service_count,omitempty" json:"service_count,omitempty"`
	StateId                     *Many2One   `xmlrpc:"state_id,omitempty" json:"state_id,omitempty"`
	TagIds                      *Relation   `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TrailerHook                 *Bool       `xmlrpc:"trailer_hook,omitempty" json:"trailer_hook,omitempty"`
	Transmission                *Selection  `xmlrpc:"transmission,omitempty" json:"transmission,omitempty"`
	VehicleProperties           interface{} `xmlrpc:"vehicle_properties,omitempty" json:"vehicle_properties,omitempty"`
	VehicleType                 *Selection  `xmlrpc:"vehicle_type,omitempty" json:"vehicle_type,omitempty"`
	VinSn                       *String     `xmlrpc:"vin_sn,omitempty" json:"vin_sn,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteOffDate                *Time       `xmlrpc:"write_off_date,omitempty" json:"write_off_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
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
	if len(*fvs) == 0 {
		return nil, nil
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
	if len(*fvs) == 0 {
		return nil, nil
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
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
