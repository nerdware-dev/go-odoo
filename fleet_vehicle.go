package odoo

import (
	"fmt"
)

// FleetVehicle represents fleet.vehicle model.
type FleetVehicle struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	AcquisitionDate             *Time      `xmlrpc:"acquisition_date,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	BrandId                     *Many2One  `xmlrpc:"brand_id,omptempty"`
	CarValue                    *Float     `xmlrpc:"car_value,omptempty"`
	Co2                         *Float     `xmlrpc:"co2,omptempty"`
	Color                       *String    `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractCount               *Int       `xmlrpc:"contract_count,omptempty"`
	ContractRenewalDueSoon      *Bool      `xmlrpc:"contract_renewal_due_soon,omptempty"`
	ContractRenewalName         *String    `xmlrpc:"contract_renewal_name,omptempty"`
	ContractRenewalOverdue      *Bool      `xmlrpc:"contract_renewal_overdue,omptempty"`
	ContractRenewalTotal        *String    `xmlrpc:"contract_renewal_total,omptempty"`
	CostCount                   *Int       `xmlrpc:"cost_count,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Doors                       *Int       `xmlrpc:"doors,omptempty"`
	DriverId                    *Many2One  `xmlrpc:"driver_id,omptempty"`
	FirstContractDate           *Time      `xmlrpc:"first_contract_date,omptempty"`
	FuelLogsCount               *Int       `xmlrpc:"fuel_logs_count,omptempty"`
	FuelType                    *Selection `xmlrpc:"fuel_type,omptempty"`
	FutureDriverId              *Many2One  `xmlrpc:"future_driver_id,omptempty"`
	HistoryCount                *Int       `xmlrpc:"history_count,omptempty"`
	Horsepower                  *Int       `xmlrpc:"horsepower,omptempty"`
	HorsepowerTax               *Float     `xmlrpc:"horsepower_tax,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Image128                    *String    `xmlrpc:"image_128,omptempty"`
	LicensePlate                *String    `xmlrpc:"license_plate,omptempty"`
	Location                    *String    `xmlrpc:"location,omptempty"`
	LogContracts                *Relation  `xmlrpc:"log_contracts,omptempty"`
	LogDrivers                  *Relation  `xmlrpc:"log_drivers,omptempty"`
	LogFuel                     *Relation  `xmlrpc:"log_fuel,omptempty"`
	LogServices                 *Relation  `xmlrpc:"log_services,omptempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	ModelId                     *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelYear                   *String    `xmlrpc:"model_year,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	NetCarValue                 *Float     `xmlrpc:"net_car_value,omptempty"`
	NextAssignationDate         *Time      `xmlrpc:"next_assignation_date,omptempty"`
	Odometer                    *Float     `xmlrpc:"odometer,omptempty"`
	OdometerCount               *Int       `xmlrpc:"odometer_count,omptempty"`
	OdometerUnit                *Selection `xmlrpc:"odometer_unit,omptempty"`
	PlanToChangeCar             *Bool      `xmlrpc:"plan_to_change_car,omptempty"`
	Power                       *Int       `xmlrpc:"power,omptempty"`
	ResidualValue               *Float     `xmlrpc:"residual_value,omptempty"`
	Seats                       *Int       `xmlrpc:"seats,omptempty"`
	ServiceCount                *Int       `xmlrpc:"service_count,omptempty"`
	StateId                     *Many2One  `xmlrpc:"state_id,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	Transmission                *Selection `xmlrpc:"transmission,omptempty"`
	VinSn                       *String    `xmlrpc:"vin_sn,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(FleetVehicleModel, vv)
}

// UpdateFleetVehicle updates an existing fleet.vehicle record.
func (c *Client) UpdateFleetVehicle(fv *FleetVehicle) error {
	return c.UpdateFleetVehicles([]int64{fv.Id.Get()}, fv)
}

// UpdateFleetVehicles updates existing fleet.vehicle records.
// All records (represented by ids) will be updated by fv values.
func (c *Client) UpdateFleetVehicles(ids []int64, fv *FleetVehicle) error {
	return c.Update(FleetVehicleModel, ids, fv)
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
	if fvs != nil && len(*fvs) > 0 {
		return &((*fvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fleet.vehicle not found", id)
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
	if fvs != nil && len(*fvs) > 0 {
		return &((*fvs)[0]), nil
	}
	return nil, fmt.Errorf("fleet.vehicle was not found with criteria %v", criteria)
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
	ids, err := c.Search(FleetVehicleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFleetVehicleId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fleet.vehicle was not found with criteria %v and options %v", criteria, options)
}
