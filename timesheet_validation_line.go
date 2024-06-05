package odoo

import (
	"fmt"
)

// TimesheetValidationLine represents timesheet.validation.line model.
type TimesheetValidationLine struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId   *Many2One `xmlrpc:"employee_id,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Validate     *Bool     `xmlrpc:"validate,omptempty"`
	ValidationId *Many2One `xmlrpc:"validation_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// TimesheetValidationLines represents array of timesheet.validation.line model.
type TimesheetValidationLines []TimesheetValidationLine

// TimesheetValidationLineModel is the odoo model name.
const TimesheetValidationLineModel = "timesheet.validation.line"

// Many2One convert TimesheetValidationLine to *Many2One.
func (tvl *TimesheetValidationLine) Many2One() *Many2One {
	return NewMany2One(tvl.Id.Get(), "")
}

// CreateTimesheetValidationLine creates a new timesheet.validation.line model and returns its id.
func (c *Client) CreateTimesheetValidationLine(tvl *TimesheetValidationLine) (int64, error) {
	ids, err := c.CreateTimesheetValidationLines([]*TimesheetValidationLine{tvl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimesheetValidationLine creates a new timesheet.validation.line model and returns its id.
func (c *Client) CreateTimesheetValidationLines(tvls []*TimesheetValidationLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range tvls {
		vv = append(vv, v)
	}
	return c.Create(TimesheetValidationLineModel, vv)
}

// UpdateTimesheetValidationLine updates an existing timesheet.validation.line record.
func (c *Client) UpdateTimesheetValidationLine(tvl *TimesheetValidationLine) error {
	return c.UpdateTimesheetValidationLines([]int64{tvl.Id.Get()}, tvl)
}

// UpdateTimesheetValidationLines updates existing timesheet.validation.line records.
// All records (represented by ids) will be updated by tvl values.
func (c *Client) UpdateTimesheetValidationLines(ids []int64, tvl *TimesheetValidationLine) error {
	return c.Update(TimesheetValidationLineModel, ids, tvl)
}

// DeleteTimesheetValidationLine deletes an existing timesheet.validation.line record.
func (c *Client) DeleteTimesheetValidationLine(id int64) error {
	return c.DeleteTimesheetValidationLines([]int64{id})
}

// DeleteTimesheetValidationLines deletes existing timesheet.validation.line records.
func (c *Client) DeleteTimesheetValidationLines(ids []int64) error {
	return c.Delete(TimesheetValidationLineModel, ids)
}

// GetTimesheetValidationLine gets timesheet.validation.line existing record.
func (c *Client) GetTimesheetValidationLine(id int64) (*TimesheetValidationLine, error) {
	tvls, err := c.GetTimesheetValidationLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if tvls != nil && len(*tvls) > 0 {
		return &((*tvls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of timesheet.validation.line not found", id)
}

// GetTimesheetValidationLines gets timesheet.validation.line existing records.
func (c *Client) GetTimesheetValidationLines(ids []int64) (*TimesheetValidationLines, error) {
	tvls := &TimesheetValidationLines{}
	if err := c.Read(TimesheetValidationLineModel, ids, nil, tvls); err != nil {
		return nil, err
	}
	return tvls, nil
}

// FindTimesheetValidationLine finds timesheet.validation.line record by querying it with criteria.
func (c *Client) FindTimesheetValidationLine(criteria *Criteria) (*TimesheetValidationLine, error) {
	tvls := &TimesheetValidationLines{}
	if err := c.SearchRead(TimesheetValidationLineModel, criteria, NewOptions().Limit(1), tvls); err != nil {
		return nil, err
	}
	if tvls != nil && len(*tvls) > 0 {
		return &((*tvls)[0]), nil
	}
	return nil, fmt.Errorf("timesheet.validation.line was not found with criteria %v", criteria)
}

// FindTimesheetValidationLines finds timesheet.validation.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetValidationLines(criteria *Criteria, options *Options) (*TimesheetValidationLines, error) {
	tvls := &TimesheetValidationLines{}
	if err := c.SearchRead(TimesheetValidationLineModel, criteria, options, tvls); err != nil {
		return nil, err
	}
	return tvls, nil
}

// FindTimesheetValidationLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetValidationLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TimesheetValidationLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTimesheetValidationLineId finds record id by querying it with criteria.
func (c *Client) FindTimesheetValidationLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimesheetValidationLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("timesheet.validation.line was not found with criteria %v and options %v", criteria, options)
}
