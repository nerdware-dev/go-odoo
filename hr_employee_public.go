package odoo

// HrEmployeePublic represents hr.employee.public model.
type HrEmployeePublic struct {
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
	ChildIds              *Relation  `xmlrpc:"child_ids,omitempty"`
	CoachId               *Many2One  `xmlrpc:"coach_id,omitempty"`
	Color                 *Int       `xmlrpc:"color,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentLeaveId        *Many2One  `xmlrpc:"current_leave_id,omitempty"`
	CurrentLeaveState     *Selection `xmlrpc:"current_leave_state,omitempty"`
	DepartmentId          *Many2One  `xmlrpc:"department_id,omitempty"`
	DirectBadgeIds        *Relation  `xmlrpc:"direct_badge_ids,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	EmployeeSkillIds      *Relation  `xmlrpc:"employee_skill_ids,omitempty"`
	ExpenseManagerId      *Many2One  `xmlrpc:"expense_manager_id,omitempty"`
	GoalIds               *Relation  `xmlrpc:"goal_ids,omitempty"`
	HasBadges             *Bool      `xmlrpc:"has_badges,omitempty"`
	HoursLastMonth        *Float     `xmlrpc:"hours_last_month,omitempty"`
	HoursLastMonthDisplay *String    `xmlrpc:"hours_last_month_display,omitempty"`
	HoursToday            *Float     `xmlrpc:"hours_today,omitempty"`
	HrPresenceState       *Selection `xmlrpc:"hr_presence_state,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	Image1024             *String    `xmlrpc:"image_1024,omitempty"`
	Image128              *String    `xmlrpc:"image_128,omitempty"`
	Image1920             *String    `xmlrpc:"image_1920,omitempty"`
	Image256              *String    `xmlrpc:"image_256,omitempty"`
	Image512              *String    `xmlrpc:"image_512,omitempty"`
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
	ResumeLineIds         *Relation  `xmlrpc:"resume_line_ids,omitempty"`
	ShowLeaves            *Bool      `xmlrpc:"show_leaves,omitempty"`
	SubordinateIds        *Relation  `xmlrpc:"subordinate_ids,omitempty"`
	TimesheetManagerId    *Many2One  `xmlrpc:"timesheet_manager_id,omitempty"`
	Tz                    *Selection `xmlrpc:"tz,omitempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omitempty"`
	WorkEmail             *String    `xmlrpc:"work_email,omitempty"`
	WorkLocation          *String    `xmlrpc:"work_location,omitempty"`
	WorkPhone             *String    `xmlrpc:"work_phone,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeePublics represents array of hr.employee.public model.
type HrEmployeePublics []HrEmployeePublic

// HrEmployeePublicModel is the odoo model name.
const HrEmployeePublicModel = "hr.employee.public"

// Many2One convert HrEmployeePublic to *Many2One.
func (hep *HrEmployeePublic) Many2One() *Many2One {
	return NewMany2One(hep.Id.Get(), "")
}

// CreateHrEmployeePublic creates a new hr.employee.public model and returns its id.
func (c *Client) CreateHrEmployeePublic(hep *HrEmployeePublic) (int64, error) {
	ids, err := c.CreateHrEmployeePublics([]*HrEmployeePublic{hep})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeePublic creates a new hr.employee.public model and returns its id.
func (c *Client) CreateHrEmployeePublics(heps []*HrEmployeePublic) ([]int64, error) {
	var vv []interface{}
	for _, v := range heps {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeePublicModel, vv, nil)
}

// UpdateHrEmployeePublic updates an existing hr.employee.public record.
func (c *Client) UpdateHrEmployeePublic(hep *HrEmployeePublic) error {
	return c.UpdateHrEmployeePublics([]int64{hep.Id.Get()}, hep)
}

// UpdateHrEmployeePublics updates existing hr.employee.public records.
// All records (represented by ids) will be updated by hep values.
func (c *Client) UpdateHrEmployeePublics(ids []int64, hep *HrEmployeePublic) error {
	return c.Update(HrEmployeePublicModel, ids, hep, nil)
}

// DeleteHrEmployeePublic deletes an existing hr.employee.public record.
func (c *Client) DeleteHrEmployeePublic(id int64) error {
	return c.DeleteHrEmployeePublics([]int64{id})
}

// DeleteHrEmployeePublics deletes existing hr.employee.public records.
func (c *Client) DeleteHrEmployeePublics(ids []int64) error {
	return c.Delete(HrEmployeePublicModel, ids)
}

// GetHrEmployeePublic gets hr.employee.public existing record.
func (c *Client) GetHrEmployeePublic(id int64) (*HrEmployeePublic, error) {
	heps, err := c.GetHrEmployeePublics([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*heps)[0]), nil
}

// GetHrEmployeePublics gets hr.employee.public existing records.
func (c *Client) GetHrEmployeePublics(ids []int64) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.Read(HrEmployeePublicModel, ids, nil, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublic finds hr.employee.public record by querying it with criteria.
func (c *Client) FindHrEmployeePublic(criteria *Criteria) (*HrEmployeePublic, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, NewOptions().Limit(1), heps); err != nil {
		return nil, err
	}
	return &((*heps)[0]), nil
}

// FindHrEmployeePublics finds hr.employee.public records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublics(criteria *Criteria, options *Options) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, options, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublicIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublicIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeePublicModel, criteria, options)
}

// FindHrEmployeePublicId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeePublicId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeePublicModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
