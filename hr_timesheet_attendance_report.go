package odoo

// HrTimesheetAttendanceReport represents hr.timesheet.attendance.report model.
type HrTimesheetAttendanceReport struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	Date            *Time     `xmlrpc:"date,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	TotalAttendance *Float    `xmlrpc:"total_attendance,omitempty"`
	TotalDifference *Float    `xmlrpc:"total_difference,omitempty"`
	TotalTimesheet  *Float    `xmlrpc:"total_timesheet,omitempty"`
	UserId          *Many2One `xmlrpc:"user_id,omitempty"`
}

// HrTimesheetAttendanceReports represents array of hr.timesheet.attendance.report model.
type HrTimesheetAttendanceReports []HrTimesheetAttendanceReport

// HrTimesheetAttendanceReportModel is the odoo model name.
const HrTimesheetAttendanceReportModel = "hr.timesheet.attendance.report"

// Many2One convert HrTimesheetAttendanceReport to *Many2One.
func (htar *HrTimesheetAttendanceReport) Many2One() *Many2One {
	return NewMany2One(htar.Id.Get(), "")
}

// CreateHrTimesheetAttendanceReport creates a new hr.timesheet.attendance.report model and returns its id.
func (c *Client) CreateHrTimesheetAttendanceReport(htar *HrTimesheetAttendanceReport) (int64, error) {
	ids, err := c.CreateHrTimesheetAttendanceReports([]*HrTimesheetAttendanceReport{htar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrTimesheetAttendanceReport creates a new hr.timesheet.attendance.report model and returns its id.
func (c *Client) CreateHrTimesheetAttendanceReports(htars []*HrTimesheetAttendanceReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range htars {
		vv = append(vv, v)
	}
	return c.Create(HrTimesheetAttendanceReportModel, vv, nil)
}

// UpdateHrTimesheetAttendanceReport updates an existing hr.timesheet.attendance.report record.
func (c *Client) UpdateHrTimesheetAttendanceReport(htar *HrTimesheetAttendanceReport) error {
	return c.UpdateHrTimesheetAttendanceReports([]int64{htar.Id.Get()}, htar)
}

// UpdateHrTimesheetAttendanceReports updates existing hr.timesheet.attendance.report records.
// All records (represented by ids) will be updated by htar values.
func (c *Client) UpdateHrTimesheetAttendanceReports(ids []int64, htar *HrTimesheetAttendanceReport) error {
	return c.Update(HrTimesheetAttendanceReportModel, ids, htar, nil)
}

// DeleteHrTimesheetAttendanceReport deletes an existing hr.timesheet.attendance.report record.
func (c *Client) DeleteHrTimesheetAttendanceReport(id int64) error {
	return c.DeleteHrTimesheetAttendanceReports([]int64{id})
}

// DeleteHrTimesheetAttendanceReports deletes existing hr.timesheet.attendance.report records.
func (c *Client) DeleteHrTimesheetAttendanceReports(ids []int64) error {
	return c.Delete(HrTimesheetAttendanceReportModel, ids)
}

// GetHrTimesheetAttendanceReport gets hr.timesheet.attendance.report existing record.
func (c *Client) GetHrTimesheetAttendanceReport(id int64) (*HrTimesheetAttendanceReport, error) {
	htars, err := c.GetHrTimesheetAttendanceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*htars)[0]), nil
}

// GetHrTimesheetAttendanceReports gets hr.timesheet.attendance.report existing records.
func (c *Client) GetHrTimesheetAttendanceReports(ids []int64) (*HrTimesheetAttendanceReports, error) {
	htars := &HrTimesheetAttendanceReports{}
	if err := c.Read(HrTimesheetAttendanceReportModel, ids, nil, htars); err != nil {
		return nil, err
	}
	return htars, nil
}

// FindHrTimesheetAttendanceReport finds hr.timesheet.attendance.report record by querying it with criteria.
func (c *Client) FindHrTimesheetAttendanceReport(criteria *Criteria) (*HrTimesheetAttendanceReport, error) {
	htars := &HrTimesheetAttendanceReports{}
	if err := c.SearchRead(HrTimesheetAttendanceReportModel, criteria, NewOptions().Limit(1), htars); err != nil {
		return nil, err
	}
	return &((*htars)[0]), nil
}

// FindHrTimesheetAttendanceReports finds hr.timesheet.attendance.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetAttendanceReports(criteria *Criteria, options *Options) (*HrTimesheetAttendanceReports, error) {
	htars := &HrTimesheetAttendanceReports{}
	if err := c.SearchRead(HrTimesheetAttendanceReportModel, criteria, options, htars); err != nil {
		return nil, err
	}
	return htars, nil
}

// FindHrTimesheetAttendanceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetAttendanceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrTimesheetAttendanceReportModel, criteria, options)
}

// FindHrTimesheetAttendanceReportId finds record id by querying it with criteria.
func (c *Client) FindHrTimesheetAttendanceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrTimesheetAttendanceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
