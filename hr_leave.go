package odoo

// HrLeave represents hr.leave model.
type HrLeave struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CanApprove                  *Bool      `xmlrpc:"can_approve,omitempty"`
	CanReset                    *Bool      `xmlrpc:"can_reset,omitempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omitempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	DurationDisplay             *String    `xmlrpc:"duration_display,omitempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omitempty"`
	FirstApproverId             *Many2One  `xmlrpc:"first_approver_id,omitempty"`
	HolidayStatusId             *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	HolidayType                 *Selection `xmlrpc:"holiday_type,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	LeaveTypeRequestUnit        *Selection `xmlrpc:"leave_type_request_unit,omitempty"`
	LinkedRequestIds            *Relation  `xmlrpc:"linked_request_ids,omitempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omitempty"`
	MeetingId                   *Many2One  `xmlrpc:"meeting_id,omitempty"`
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
	ModeCompanyId               *Many2One  `xmlrpc:"mode_company_id,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	Notes                       *String    `xmlrpc:"notes,omitempty"`
	NumberOfDays                *Float     `xmlrpc:"number_of_days,omitempty"`
	NumberOfDaysDisplay         *Float     `xmlrpc:"number_of_days_display,omitempty"`
	NumberOfHoursDisplay        *Float     `xmlrpc:"number_of_hours_display,omitempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	PayslipStatus               *Bool      `xmlrpc:"payslip_status,omitempty"`
	ReportNote                  *String    `xmlrpc:"report_note,omitempty"`
	RequestDateFrom             *Time      `xmlrpc:"request_date_from,omitempty"`
	RequestDateFromPeriod       *Selection `xmlrpc:"request_date_from_period,omitempty"`
	RequestDateTo               *Time      `xmlrpc:"request_date_to,omitempty"`
	RequestHourFrom             *Selection `xmlrpc:"request_hour_from,omitempty"`
	RequestHourTo               *Selection `xmlrpc:"request_hour_to,omitempty"`
	RequestUnitCustom           *Bool      `xmlrpc:"request_unit_custom,omitempty"`
	RequestUnitHalf             *Bool      `xmlrpc:"request_unit_half,omitempty"`
	RequestUnitHours            *Bool      `xmlrpc:"request_unit_hours,omitempty"`
	SecondApproverId            *Many2One  `xmlrpc:"second_approver_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TimesheetIds                *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	Tz                          *Selection `xmlrpc:"tz,omitempty"`
	TzMismatch                  *Bool      `xmlrpc:"tz_mismatch,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	ValidationType              *Selection `xmlrpc:"validation_type,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaves represents array of hr.leave model.
type HrLeaves []HrLeave

// HrLeaveModel is the odoo model name.
const HrLeaveModel = "hr.leave"

// Many2One convert HrLeave to *Many2One.
func (hl *HrLeave) Many2One() *Many2One {
	return NewMany2One(hl.Id.Get(), "")
}

// CreateHrLeave creates a new hr.leave model and returns its id.
func (c *Client) CreateHrLeave(hl *HrLeave) (int64, error) {
	ids, err := c.CreateHrLeaves([]*HrLeave{hl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeave creates a new hr.leave model and returns its id.
func (c *Client) CreateHrLeaves(hls []*HrLeave) ([]int64, error) {
	var vv []interface{}
	for _, v := range hls {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveModel, vv, nil)
}

// UpdateHrLeave updates an existing hr.leave record.
func (c *Client) UpdateHrLeave(hl *HrLeave) error {
	return c.UpdateHrLeaves([]int64{hl.Id.Get()}, hl)
}

// UpdateHrLeaves updates existing hr.leave records.
// All records (represented by ids) will be updated by hl values.
func (c *Client) UpdateHrLeaves(ids []int64, hl *HrLeave) error {
	return c.Update(HrLeaveModel, ids, hl, nil)
}

// DeleteHrLeave deletes an existing hr.leave record.
func (c *Client) DeleteHrLeave(id int64) error {
	return c.DeleteHrLeaves([]int64{id})
}

// DeleteHrLeaves deletes existing hr.leave records.
func (c *Client) DeleteHrLeaves(ids []int64) error {
	return c.Delete(HrLeaveModel, ids)
}

// GetHrLeave gets hr.leave existing record.
func (c *Client) GetHrLeave(id int64) (*HrLeave, error) {
	hls, err := c.GetHrLeaves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hls)[0]), nil
}

// GetHrLeaves gets hr.leave existing records.
func (c *Client) GetHrLeaves(ids []int64) (*HrLeaves, error) {
	hls := &HrLeaves{}
	if err := c.Read(HrLeaveModel, ids, nil, hls); err != nil {
		return nil, err
	}
	return hls, nil
}

// FindHrLeave finds hr.leave record by querying it with criteria.
func (c *Client) FindHrLeave(criteria *Criteria) (*HrLeave, error) {
	hls := &HrLeaves{}
	if err := c.SearchRead(HrLeaveModel, criteria, NewOptions().Limit(1), hls); err != nil {
		return nil, err
	}
	return &((*hls)[0]), nil
}

// FindHrLeaves finds hr.leave records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaves(criteria *Criteria, options *Options) (*HrLeaves, error) {
	hls := &HrLeaves{}
	if err := c.SearchRead(HrLeaveModel, criteria, options, hls); err != nil {
		return nil, err
	}
	return hls, nil
}

// FindHrLeaveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveModel, criteria, options)
}

// FindHrLeaveId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
