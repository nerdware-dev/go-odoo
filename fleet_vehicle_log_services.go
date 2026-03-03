package odoo

// FleetVehicleLogServices represents fleet.vehicle.log.services model.
type FleetVehicleLogServices struct {
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
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InvRef                      *String    `xmlrpc:"inv_ref,omitempty" json:"inv_ref,omitempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omitempty" json:"manager_id,omitempty"`
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
	Notes                       *String    `xmlrpc:"notes,omitempty" json:"notes,omitempty"`
	Odometer                    *Float     `xmlrpc:"odometer,omitempty" json:"odometer,omitempty"`
	OdometerId                  *Many2One  `xmlrpc:"odometer_id,omitempty" json:"odometer_id,omitempty"`
	OdometerUnit                *Selection `xmlrpc:"odometer_unit,omitempty" json:"odometer_unit,omitempty"`
	PurchaserEmployeeId         *Many2One  `xmlrpc:"purchaser_employee_id,omitempty" json:"purchaser_employee_id,omitempty"`
	PurchaserId                 *Many2One  `xmlrpc:"purchaser_id,omitempty" json:"purchaser_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	ServiceTypeId               *Many2One  `xmlrpc:"service_type_id,omitempty" json:"service_type_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	VehicleId                   *Many2One  `xmlrpc:"vehicle_id,omitempty" json:"vehicle_id,omitempty"`
	VendorId                    *Many2One  `xmlrpc:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
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
	if len(*fvlss) == 0 {
		return nil, nil
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
	if len(*fvlss) == 0 {
		return nil, nil
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
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
