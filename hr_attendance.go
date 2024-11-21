package odoo

// HrAttendance represents hr.attendance model.
type HrAttendance struct {
	CheckIn                  *Time      `xmlrpc:"check_in,omitempty" json:"check_in,omitempty"`
	CheckOut                 *Time      `xmlrpc:"check_out,omitempty" json:"check_out,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty" json:"department_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmployeeId               *Many2One  `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InBrowser                *String    `xmlrpc:"in_browser,omitempty" json:"in_browser,omitempty"`
	InCity                   *String    `xmlrpc:"in_city,omitempty" json:"in_city,omitempty"`
	InCountryName            *String    `xmlrpc:"in_country_name,omitempty" json:"in_country_name,omitempty"`
	InIpAddress              *String    `xmlrpc:"in_ip_address,omitempty" json:"in_ip_address,omitempty"`
	InLatitude               *Float     `xmlrpc:"in_latitude,omitempty" json:"in_latitude,omitempty"`
	InLongitude              *Float     `xmlrpc:"in_longitude,omitempty" json:"in_longitude,omitempty"`
	InMode                   *Selection `xmlrpc:"in_mode,omitempty" json:"in_mode,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	OutBrowser               *String    `xmlrpc:"out_browser,omitempty" json:"out_browser,omitempty"`
	OutCity                  *String    `xmlrpc:"out_city,omitempty" json:"out_city,omitempty"`
	OutCountryName           *String    `xmlrpc:"out_country_name,omitempty" json:"out_country_name,omitempty"`
	OutIpAddress             *String    `xmlrpc:"out_ip_address,omitempty" json:"out_ip_address,omitempty"`
	OutLatitude              *Float     `xmlrpc:"out_latitude,omitempty" json:"out_latitude,omitempty"`
	OutLongitude             *Float     `xmlrpc:"out_longitude,omitempty" json:"out_longitude,omitempty"`
	OutMode                  *Selection `xmlrpc:"out_mode,omitempty" json:"out_mode,omitempty"`
	OvertimeHours            *Float     `xmlrpc:"overtime_hours,omitempty" json:"overtime_hours,omitempty"`
	OvertimeProgress         *Float     `xmlrpc:"overtime_progress,omitempty" json:"overtime_progress,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WorkedHours              *Float     `xmlrpc:"worked_hours,omitempty" json:"worked_hours,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// HrAttendances represents array of hr.attendance model.
type HrAttendances []HrAttendance

// HrAttendanceModel is the odoo model name.
const HrAttendanceModel = "hr.attendance"

// Many2One convert HrAttendance to *Many2One.
func (ha *HrAttendance) Many2One() *Many2One {
	return NewMany2One(ha.Id.Get(), "")
}

// CreateHrAttendance creates a new hr.attendance model and returns its id.
func (c *Client) CreateHrAttendance(ha *HrAttendance) (int64, error) {
	ids, err := c.CreateHrAttendances([]*HrAttendance{ha})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAttendance creates a new hr.attendance model and returns its id.
func (c *Client) CreateHrAttendances(has []*HrAttendance) ([]int64, error) {
	var vv []interface{}
	for _, v := range has {
		vv = append(vv, v)
	}
	return c.Create(HrAttendanceModel, vv, nil)
}

// UpdateHrAttendance updates an existing hr.attendance record.
func (c *Client) UpdateHrAttendance(ha *HrAttendance) error {
	return c.UpdateHrAttendances([]int64{ha.Id.Get()}, ha)
}

// UpdateHrAttendances updates existing hr.attendance records.
// All records (represented by ids) will be updated by ha values.
func (c *Client) UpdateHrAttendances(ids []int64, ha *HrAttendance) error {
	return c.Update(HrAttendanceModel, ids, ha, nil)
}

// DeleteHrAttendance deletes an existing hr.attendance record.
func (c *Client) DeleteHrAttendance(id int64) error {
	return c.DeleteHrAttendances([]int64{id})
}

// DeleteHrAttendances deletes existing hr.attendance records.
func (c *Client) DeleteHrAttendances(ids []int64) error {
	return c.Delete(HrAttendanceModel, ids)
}

// GetHrAttendance gets hr.attendance existing record.
func (c *Client) GetHrAttendance(id int64) (*HrAttendance, error) {
	has, err := c.GetHrAttendances([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*has)[0]), nil
}

// GetHrAttendances gets hr.attendance existing records.
func (c *Client) GetHrAttendances(ids []int64) (*HrAttendances, error) {
	has := &HrAttendances{}
	if err := c.Read(HrAttendanceModel, ids, nil, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAttendance finds hr.attendance record by querying it with criteria.
func (c *Client) FindHrAttendance(criteria *Criteria) (*HrAttendance, error) {
	has := &HrAttendances{}
	if err := c.SearchRead(HrAttendanceModel, criteria, NewOptions().Limit(1), has); err != nil {
		return nil, err
	}
	return &((*has)[0]), nil
}

// FindHrAttendances finds hr.attendance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendances(criteria *Criteria, options *Options) (*HrAttendances, error) {
	has := &HrAttendances{}
	if err := c.SearchRead(HrAttendanceModel, criteria, options, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAttendanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrAttendanceModel, criteria, options)
}

// FindHrAttendanceId finds record id by querying it with criteria.
func (c *Client) FindHrAttendanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAttendanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
