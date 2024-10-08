package odoo

// HrEmployeeBase represents hr.employee.base model.
type HrEmployeeBase struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	Active                *Bool      `xmlrpc:"active,omitempty"`
	AddressId             *Many2One  `xmlrpc:"address_id,omitempty"`
	AllocationCount       *Float     `xmlrpc:"allocation_count,omitempty"`
	AllocationDisplay     *String    `xmlrpc:"allocation_display,omitempty"`
	AllocationUsedCount   *Float     `xmlrpc:"allocation_used_count,omitempty"`
	AllocationUsedDisplay *String    `xmlrpc:"allocation_used_display,omitempty"`
	AttendanceIds         *Relation  `xmlrpc:"attendance_ids,omitempty"`
	AttendanceState       *Selection `xmlrpc:"attendance_state,omitempty"`
	BadgeIds              *Relation  `xmlrpc:"badge_ids,omitempty"`
	ChildAllCount         *Int       `xmlrpc:"child_all_count,omitempty"`
	CoachId               *Many2One  `xmlrpc:"coach_id,omitempty"`
	Color                 *Int       `xmlrpc:"color,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CurrentLeaveId        *Many2One  `xmlrpc:"current_leave_id,omitempty"`
	CurrentLeaveState     *Selection `xmlrpc:"current_leave_state,omitempty"`
	DepartmentId          *Many2One  `xmlrpc:"department_id,omitempty"`
	DirectBadgeIds        *Relation  `xmlrpc:"direct_badge_ids,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	GoalIds               *Relation  `xmlrpc:"goal_ids,omitempty"`
	HasBadges             *Bool      `xmlrpc:"has_badges,omitempty"`
	HoursLastMonth        *Float     `xmlrpc:"hours_last_month,omitempty"`
	HoursLastMonthDisplay *String    `xmlrpc:"hours_last_month_display,omitempty"`
	HoursToday            *Float     `xmlrpc:"hours_today,omitempty"`
	HrPresenceState       *Selection `xmlrpc:"hr_presence_state,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	IsAbsent              *Bool      `xmlrpc:"is_absent,omitempty"`
	JobId                 *Many2One  `xmlrpc:"job_id,omitempty"`
	JobTitle              *String    `xmlrpc:"job_title,omitempty"`
	LastActivity          *Time      `xmlrpc:"last_activity,omitempty"`
	LastActivityTime      *String    `xmlrpc:"last_activity_time,omitempty"`
	LastAttendanceId      *Many2One  `xmlrpc:"last_attendance_id,omitempty"`
	LastCheckIn           *Time      `xmlrpc:"last_check_in,omitempty"`
	LastCheckOut          *Time      `xmlrpc:"last_check_out,omitempty"`
	LeaveDateFrom         *Time      `xmlrpc:"leave_date_from,omitempty"`
	LeaveDateTo           *Time      `xmlrpc:"leave_date_to,omitempty"`
	LeaveManagerId        *Many2One  `xmlrpc:"leave_manager_id,omitempty"`
	LeavesCount           *Float     `xmlrpc:"leaves_count,omitempty"`
	MobilePhone           *String    `xmlrpc:"mobile_phone,omitempty"`
	Name                  *String    `xmlrpc:"name,omitempty"`
	ParentId              *Many2One  `xmlrpc:"parent_id,omitempty"`
	RemainingLeaves       *Float     `xmlrpc:"remaining_leaves,omitempty"`
	ResourceCalendarId    *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId            *Many2One  `xmlrpc:"resource_id,omitempty"`
	ShowLeaves            *Bool      `xmlrpc:"show_leaves,omitempty"`
	Tz                    *Selection `xmlrpc:"tz,omitempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omitempty"`
	WorkEmail             *String    `xmlrpc:"work_email,omitempty"`
	WorkLocation          *String    `xmlrpc:"work_location,omitempty"`
	WorkPhone             *String    `xmlrpc:"work_phone,omitempty"`
}

// HrEmployeeBases represents array of hr.employee.base model.
type HrEmployeeBases []HrEmployeeBase

// HrEmployeeBaseModel is the odoo model name.
const HrEmployeeBaseModel = "hr.employee.base"

// Many2One convert HrEmployeeBase to *Many2One.
func (heb *HrEmployeeBase) Many2One() *Many2One {
	return NewMany2One(heb.Id.Get(), "")
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBase(heb *HrEmployeeBase) (int64, error) {
	ids, err := c.CreateHrEmployeeBases([]*HrEmployeeBase{heb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBases(hebs []*HrEmployeeBase) ([]int64, error) {
	var vv []interface{}
	for _, v := range hebs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeBaseModel, vv, nil)
}

// UpdateHrEmployeeBase updates an existing hr.employee.base record.
func (c *Client) UpdateHrEmployeeBase(heb *HrEmployeeBase) error {
	return c.UpdateHrEmployeeBases([]int64{heb.Id.Get()}, heb)
}

// UpdateHrEmployeeBases updates existing hr.employee.base records.
// All records (represented by ids) will be updated by heb values.
func (c *Client) UpdateHrEmployeeBases(ids []int64, heb *HrEmployeeBase) error {
	return c.Update(HrEmployeeBaseModel, ids, heb, nil)
}

// DeleteHrEmployeeBase deletes an existing hr.employee.base record.
func (c *Client) DeleteHrEmployeeBase(id int64) error {
	return c.DeleteHrEmployeeBases([]int64{id})
}

// DeleteHrEmployeeBases deletes existing hr.employee.base records.
func (c *Client) DeleteHrEmployeeBases(ids []int64) error {
	return c.Delete(HrEmployeeBaseModel, ids)
}

// GetHrEmployeeBase gets hr.employee.base existing record.
func (c *Client) GetHrEmployeeBase(id int64) (*HrEmployeeBase, error) {
	hebs, err := c.GetHrEmployeeBases([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hebs)[0]), nil
}

// GetHrEmployeeBases gets hr.employee.base existing records.
func (c *Client) GetHrEmployeeBases(ids []int64) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.Read(HrEmployeeBaseModel, ids, nil, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBase finds hr.employee.base record by querying it with criteria.
func (c *Client) FindHrEmployeeBase(criteria *Criteria) (*HrEmployeeBase, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, NewOptions().Limit(1), hebs); err != nil {
		return nil, err
	}
	return &((*hebs)[0]), nil
}

// FindHrEmployeeBases finds hr.employee.base records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBases(criteria *Criteria, options *Options) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, options, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeBaseModel, criteria, options)
}

// FindHrEmployeeBaseId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeBaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeBaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
