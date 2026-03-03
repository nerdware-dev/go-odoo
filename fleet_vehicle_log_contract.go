package odoo

// FleetVehicleLogContract represents fleet.vehicle.log.contract model.
type FleetVehicleLogContract struct {
	Active                      *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	Amount                      *Float     `xmlrpc:"amount,omitempty" json:"amount,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CostFrequency               *Selection `xmlrpc:"cost_frequency,omitempty" json:"cost_frequency,omitempty"`
	CostGenerated               *Float     `xmlrpc:"cost_generated,omitempty" json:"cost_generated,omitempty"`
	CostSubtypeId               *Many2One  `xmlrpc:"cost_subtype_id,omitempty" json:"cost_subtype_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DaysLeft                    *Int       `xmlrpc:"days_left,omitempty" json:"days_left,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	ExpirationDate              *Time      `xmlrpc:"expiration_date,omitempty" json:"expiration_date,omitempty"`
	ExpiresToday                *Bool      `xmlrpc:"expires_today,omitempty" json:"expires_today,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InsRef                      *String    `xmlrpc:"ins_ref,omitempty" json:"ins_ref,omitempty"`
	InsurerId                   *Many2One  `xmlrpc:"insurer_id,omitempty" json:"insurer_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Notes                       *String    `xmlrpc:"notes,omitempty" json:"notes,omitempty"`
	PurchaserEmployeeId         *Many2One  `xmlrpc:"purchaser_employee_id,omitempty" json:"purchaser_employee_id,omitempty"`
	PurchaserId                 *Many2One  `xmlrpc:"purchaser_id,omitempty" json:"purchaser_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	ServiceIds                  *Relation  `xmlrpc:"service_ids,omitempty" json:"service_ids,omitempty"`
	StartDate                   *Time      `xmlrpc:"start_date,omitempty" json:"start_date,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	VehicleId                   *Many2One  `xmlrpc:"vehicle_id,omitempty" json:"vehicle_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// FleetVehicleLogContracts represents array of fleet.vehicle.log.contract model.
type FleetVehicleLogContracts []FleetVehicleLogContract

// FleetVehicleLogContractModel is the odoo model name.
const FleetVehicleLogContractModel = "fleet.vehicle.log.contract"

// Many2One convert FleetVehicleLogContract to *Many2One.
func (fvlc *FleetVehicleLogContract) Many2One() *Many2One {
	return NewMany2One(fvlc.Id.Get(), "")
}

// CreateFleetVehicleLogContract creates a new fleet.vehicle.log.contract model and returns its id.
func (c *Client) CreateFleetVehicleLogContract(fvlc *FleetVehicleLogContract) (int64, error) {
	ids, err := c.CreateFleetVehicleLogContracts([]*FleetVehicleLogContract{fvlc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFleetVehicleLogContract creates a new fleet.vehicle.log.contract model and returns its id.
func (c *Client) CreateFleetVehicleLogContracts(fvlcs []*FleetVehicleLogContract) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvlcs {
		vv = append(vv, v)
	}
	return c.Create(FleetVehicleLogContractModel, vv, nil)
}

// UpdateFleetVehicleLogContract updates an existing fleet.vehicle.log.contract record.
func (c *Client) UpdateFleetVehicleLogContract(fvlc *FleetVehicleLogContract) error {
	return c.UpdateFleetVehicleLogContracts([]int64{fvlc.Id.Get()}, fvlc)
}

// UpdateFleetVehicleLogContracts updates existing fleet.vehicle.log.contract records.
// All records (represented by ids) will be updated by fvlc values.
func (c *Client) UpdateFleetVehicleLogContracts(ids []int64, fvlc *FleetVehicleLogContract) error {
	return c.Update(FleetVehicleLogContractModel, ids, fvlc, nil)
}

// DeleteFleetVehicleLogContract deletes an existing fleet.vehicle.log.contract record.
func (c *Client) DeleteFleetVehicleLogContract(id int64) error {
	return c.DeleteFleetVehicleLogContracts([]int64{id})
}

// DeleteFleetVehicleLogContracts deletes existing fleet.vehicle.log.contract records.
func (c *Client) DeleteFleetVehicleLogContracts(ids []int64) error {
	return c.Delete(FleetVehicleLogContractModel, ids)
}

// GetFleetVehicleLogContract gets fleet.vehicle.log.contract existing record.
func (c *Client) GetFleetVehicleLogContract(id int64) (*FleetVehicleLogContract, error) {
	fvlcs, err := c.GetFleetVehicleLogContracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*fvlcs) == 0 {
		return nil, nil
	}
	return &((*fvlcs)[0]), nil
}

// GetFleetVehicleLogContracts gets fleet.vehicle.log.contract existing records.
func (c *Client) GetFleetVehicleLogContracts(ids []int64) (*FleetVehicleLogContracts, error) {
	fvlcs := &FleetVehicleLogContracts{}
	if err := c.Read(FleetVehicleLogContractModel, ids, nil, fvlcs); err != nil {
		return nil, err
	}
	return fvlcs, nil
}

// FindFleetVehicleLogContract finds fleet.vehicle.log.contract record by querying it with criteria.
func (c *Client) FindFleetVehicleLogContract(criteria *Criteria) (*FleetVehicleLogContract, error) {
	fvlcs := &FleetVehicleLogContracts{}
	if err := c.SearchRead(FleetVehicleLogContractModel, criteria, NewOptions().Limit(1), fvlcs); err != nil {
		return nil, err
	}
	if len(*fvlcs) == 0 {
		return nil, nil
	}
	return &((*fvlcs)[0]), nil
}

// FindFleetVehicleLogContracts finds fleet.vehicle.log.contract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleLogContracts(criteria *Criteria, options *Options) (*FleetVehicleLogContracts, error) {
	fvlcs := &FleetVehicleLogContracts{}
	if err := c.SearchRead(FleetVehicleLogContractModel, criteria, options, fvlcs); err != nil {
		return nil, err
	}
	return fvlcs, nil
}

// FindFleetVehicleLogContractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFleetVehicleLogContractIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FleetVehicleLogContractModel, criteria, options)
}

// FindFleetVehicleLogContractId finds record id by querying it with criteria.
func (c *Client) FindFleetVehicleLogContractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FleetVehicleLogContractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
