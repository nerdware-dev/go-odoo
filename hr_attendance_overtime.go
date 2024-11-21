package odoo

// HrAttendanceOvertime represents hr.attendance.overtime model.
type HrAttendanceOvertime struct {
	Adjustment   *Bool     `xmlrpc:"adjustment,omitempty" json:"adjustment,omitempty"`
	CompanyId    *Many2One `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Date         *Time     `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	Duration     *Float    `xmlrpc:"duration,omitempty" json:"duration,omitempty"`
	DurationReal *Float    `xmlrpc:"duration_real,omitempty" json:"duration_real,omitempty"`
	EmployeeId   *Many2One `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty" json:"id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// HrAttendanceOvertimes represents array of hr.attendance.overtime model.
type HrAttendanceOvertimes []HrAttendanceOvertime

// HrAttendanceOvertimeModel is the odoo model name.
const HrAttendanceOvertimeModel = "hr.attendance.overtime"

// Many2One convert HrAttendanceOvertime to *Many2One.
func (hao *HrAttendanceOvertime) Many2One() *Many2One {
	return NewMany2One(hao.Id.Get(), "")
}

// CreateHrAttendanceOvertime creates a new hr.attendance.overtime model and returns its id.
func (c *Client) CreateHrAttendanceOvertime(hao *HrAttendanceOvertime) (int64, error) {
	ids, err := c.CreateHrAttendanceOvertimes([]*HrAttendanceOvertime{hao})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAttendanceOvertime creates a new hr.attendance.overtime model and returns its id.
func (c *Client) CreateHrAttendanceOvertimes(haos []*HrAttendanceOvertime) ([]int64, error) {
	var vv []interface{}
	for _, v := range haos {
		vv = append(vv, v)
	}
	return c.Create(HrAttendanceOvertimeModel, vv, nil)
}

// UpdateHrAttendanceOvertime updates an existing hr.attendance.overtime record.
func (c *Client) UpdateHrAttendanceOvertime(hao *HrAttendanceOvertime) error {
	return c.UpdateHrAttendanceOvertimes([]int64{hao.Id.Get()}, hao)
}

// UpdateHrAttendanceOvertimes updates existing hr.attendance.overtime records.
// All records (represented by ids) will be updated by hao values.
func (c *Client) UpdateHrAttendanceOvertimes(ids []int64, hao *HrAttendanceOvertime) error {
	return c.Update(HrAttendanceOvertimeModel, ids, hao, nil)
}

// DeleteHrAttendanceOvertime deletes an existing hr.attendance.overtime record.
func (c *Client) DeleteHrAttendanceOvertime(id int64) error {
	return c.DeleteHrAttendanceOvertimes([]int64{id})
}

// DeleteHrAttendanceOvertimes deletes existing hr.attendance.overtime records.
func (c *Client) DeleteHrAttendanceOvertimes(ids []int64) error {
	return c.Delete(HrAttendanceOvertimeModel, ids)
}

// GetHrAttendanceOvertime gets hr.attendance.overtime existing record.
func (c *Client) GetHrAttendanceOvertime(id int64) (*HrAttendanceOvertime, error) {
	haos, err := c.GetHrAttendanceOvertimes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*haos)[0]), nil
}

// GetHrAttendanceOvertimes gets hr.attendance.overtime existing records.
func (c *Client) GetHrAttendanceOvertimes(ids []int64) (*HrAttendanceOvertimes, error) {
	haos := &HrAttendanceOvertimes{}
	if err := c.Read(HrAttendanceOvertimeModel, ids, nil, haos); err != nil {
		return nil, err
	}
	return haos, nil
}

// FindHrAttendanceOvertime finds hr.attendance.overtime record by querying it with criteria.
func (c *Client) FindHrAttendanceOvertime(criteria *Criteria) (*HrAttendanceOvertime, error) {
	haos := &HrAttendanceOvertimes{}
	if err := c.SearchRead(HrAttendanceOvertimeModel, criteria, NewOptions().Limit(1), haos); err != nil {
		return nil, err
	}
	return &((*haos)[0]), nil
}

// FindHrAttendanceOvertimes finds hr.attendance.overtime records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceOvertimes(criteria *Criteria, options *Options) (*HrAttendanceOvertimes, error) {
	haos := &HrAttendanceOvertimes{}
	if err := c.SearchRead(HrAttendanceOvertimeModel, criteria, options, haos); err != nil {
		return nil, err
	}
	return haos, nil
}

// FindHrAttendanceOvertimeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceOvertimeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrAttendanceOvertimeModel, criteria, options)
}

// FindHrAttendanceOvertimeId finds record id by querying it with criteria.
func (c *Client) FindHrAttendanceOvertimeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAttendanceOvertimeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
