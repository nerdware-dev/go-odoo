package odoo

// HrLeaveAllocation represents hr.leave.allocation model.
type HrLeaveAllocation struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccrualLimit                *Int       `xmlrpc:"accrual_limit,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllocationType              *Selection `xmlrpc:"allocation_type,omitempty"`
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
	IntervalNumber              *Int       `xmlrpc:"interval_number,omitempty"`
	IntervalUnit                *Selection `xmlrpc:"interval_unit,omitempty"`
	LeavesTaken                 *Float     `xmlrpc:"leaves_taken,omitempty"`
	LinkedRequestIds            *Relation  `xmlrpc:"linked_request_ids,omitempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omitempty"`
	MaxLeaves                   *Float     `xmlrpc:"max_leaves,omitempty"`
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
	Nextcall                    *Time      `xmlrpc:"nextcall,omitempty"`
	Notes                       *String    `xmlrpc:"notes,omitempty"`
	NumberOfDays                *Float     `xmlrpc:"number_of_days,omitempty"`
	NumberOfDaysDisplay         *Float     `xmlrpc:"number_of_days_display,omitempty"`
	NumberOfHoursDisplay        *Float     `xmlrpc:"number_of_hours_display,omitempty"`
	NumberPerInterval           *Float     `xmlrpc:"number_per_interval,omitempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	SecondApproverId            *Many2One  `xmlrpc:"second_approver_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TypeRequestUnit             *Selection `xmlrpc:"type_request_unit,omitempty"`
	UnitPerInterval             *Selection `xmlrpc:"unit_per_interval,omitempty"`
	ValidationType              *Selection `xmlrpc:"validation_type,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveAllocations represents array of hr.leave.allocation model.
type HrLeaveAllocations []HrLeaveAllocation

// HrLeaveAllocationModel is the odoo model name.
const HrLeaveAllocationModel = "hr.leave.allocation"

// Many2One convert HrLeaveAllocation to *Many2One.
func (hla *HrLeaveAllocation) Many2One() *Many2One {
	return NewMany2One(hla.Id.Get(), "")
}

// CreateHrLeaveAllocation creates a new hr.leave.allocation model and returns its id.
func (c *Client) CreateHrLeaveAllocation(hla *HrLeaveAllocation) (int64, error) {
	ids, err := c.CreateHrLeaveAllocations([]*HrLeaveAllocation{hla})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAllocation creates a new hr.leave.allocation model and returns its id.
func (c *Client) CreateHrLeaveAllocations(hlas []*HrLeaveAllocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlas {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAllocationModel, vv, nil)
}

// UpdateHrLeaveAllocation updates an existing hr.leave.allocation record.
func (c *Client) UpdateHrLeaveAllocation(hla *HrLeaveAllocation) error {
	return c.UpdateHrLeaveAllocations([]int64{hla.Id.Get()}, hla)
}

// UpdateHrLeaveAllocations updates existing hr.leave.allocation records.
// All records (represented by ids) will be updated by hla values.
func (c *Client) UpdateHrLeaveAllocations(ids []int64, hla *HrLeaveAllocation) error {
	return c.Update(HrLeaveAllocationModel, ids, hla, nil)
}

// DeleteHrLeaveAllocation deletes an existing hr.leave.allocation record.
func (c *Client) DeleteHrLeaveAllocation(id int64) error {
	return c.DeleteHrLeaveAllocations([]int64{id})
}

// DeleteHrLeaveAllocations deletes existing hr.leave.allocation records.
func (c *Client) DeleteHrLeaveAllocations(ids []int64) error {
	return c.Delete(HrLeaveAllocationModel, ids)
}

// GetHrLeaveAllocation gets hr.leave.allocation existing record.
func (c *Client) GetHrLeaveAllocation(id int64) (*HrLeaveAllocation, error) {
	hlas, err := c.GetHrLeaveAllocations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlas)[0]), nil
}

// GetHrLeaveAllocations gets hr.leave.allocation existing records.
func (c *Client) GetHrLeaveAllocations(ids []int64) (*HrLeaveAllocations, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.Read(HrLeaveAllocationModel, ids, nil, hlas); err != nil {
		return nil, err
	}
	return hlas, nil
}

// FindHrLeaveAllocation finds hr.leave.allocation record by querying it with criteria.
func (c *Client) FindHrLeaveAllocation(criteria *Criteria) (*HrLeaveAllocation, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.SearchRead(HrLeaveAllocationModel, criteria, NewOptions().Limit(1), hlas); err != nil {
		return nil, err
	}
	return &((*hlas)[0]), nil
}

// FindHrLeaveAllocations finds hr.leave.allocation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocations(criteria *Criteria, options *Options) (*HrLeaveAllocations, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.SearchRead(HrLeaveAllocationModel, criteria, options, hlas); err != nil {
		return nil, err
	}
	return hlas, nil
}

// FindHrLeaveAllocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveAllocationModel, criteria, options)
}

// FindHrLeaveAllocationId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAllocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAllocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
