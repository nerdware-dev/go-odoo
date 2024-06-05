package odoo

import (
	"fmt"
)

// TimesheetValidation represents timesheet.validation model.
type TimesheetValidation struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	ValidationDate    *Time     `xmlrpc:"validation_date,omptempty"`
	ValidationLineIds *Relation `xmlrpc:"validation_line_ids,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// TimesheetValidations represents array of timesheet.validation model.
type TimesheetValidations []TimesheetValidation

// TimesheetValidationModel is the odoo model name.
const TimesheetValidationModel = "timesheet.validation"

// Many2One convert TimesheetValidation to *Many2One.
func (tv *TimesheetValidation) Many2One() *Many2One {
	return NewMany2One(tv.Id.Get(), "")
}

// CreateTimesheetValidation creates a new timesheet.validation model and returns its id.
func (c *Client) CreateTimesheetValidation(tv *TimesheetValidation) (int64, error) {
	ids, err := c.CreateTimesheetValidations([]*TimesheetValidation{tv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimesheetValidation creates a new timesheet.validation model and returns its id.
func (c *Client) CreateTimesheetValidations(tvs []*TimesheetValidation) ([]int64, error) {
	var vv []interface{}
	for _, v := range tvs {
		vv = append(vv, v)
	}
	return c.Create(TimesheetValidationModel, vv)
}

// UpdateTimesheetValidation updates an existing timesheet.validation record.
func (c *Client) UpdateTimesheetValidation(tv *TimesheetValidation) error {
	return c.UpdateTimesheetValidations([]int64{tv.Id.Get()}, tv)
}

// UpdateTimesheetValidations updates existing timesheet.validation records.
// All records (represented by ids) will be updated by tv values.
func (c *Client) UpdateTimesheetValidations(ids []int64, tv *TimesheetValidation) error {
	return c.Update(TimesheetValidationModel, ids, tv)
}

// DeleteTimesheetValidation deletes an existing timesheet.validation record.
func (c *Client) DeleteTimesheetValidation(id int64) error {
	return c.DeleteTimesheetValidations([]int64{id})
}

// DeleteTimesheetValidations deletes existing timesheet.validation records.
func (c *Client) DeleteTimesheetValidations(ids []int64) error {
	return c.Delete(TimesheetValidationModel, ids)
}

// GetTimesheetValidation gets timesheet.validation existing record.
func (c *Client) GetTimesheetValidation(id int64) (*TimesheetValidation, error) {
	tvs, err := c.GetTimesheetValidations([]int64{id})
	if err != nil {
		return nil, err
	}
	if tvs != nil && len(*tvs) > 0 {
		return &((*tvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of timesheet.validation not found", id)
}

// GetTimesheetValidations gets timesheet.validation existing records.
func (c *Client) GetTimesheetValidations(ids []int64) (*TimesheetValidations, error) {
	tvs := &TimesheetValidations{}
	if err := c.Read(TimesheetValidationModel, ids, nil, tvs); err != nil {
		return nil, err
	}
	return tvs, nil
}

// FindTimesheetValidation finds timesheet.validation record by querying it with criteria.
func (c *Client) FindTimesheetValidation(criteria *Criteria) (*TimesheetValidation, error) {
	tvs := &TimesheetValidations{}
	if err := c.SearchRead(TimesheetValidationModel, criteria, NewOptions().Limit(1), tvs); err != nil {
		return nil, err
	}
	if tvs != nil && len(*tvs) > 0 {
		return &((*tvs)[0]), nil
	}
	return nil, fmt.Errorf("timesheet.validation was not found with criteria %v", criteria)
}

// FindTimesheetValidations finds timesheet.validation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetValidations(criteria *Criteria, options *Options) (*TimesheetValidations, error) {
	tvs := &TimesheetValidations{}
	if err := c.SearchRead(TimesheetValidationModel, criteria, options, tvs); err != nil {
		return nil, err
	}
	return tvs, nil
}

// FindTimesheetValidationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetValidationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TimesheetValidationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTimesheetValidationId finds record id by querying it with criteria.
func (c *Client) FindTimesheetValidationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimesheetValidationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("timesheet.validation was not found with criteria %v and options %v", criteria, options)
}
