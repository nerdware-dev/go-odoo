package odoo

import (
	"fmt"
)

// LunchAlert represents lunch.alert model.
type LunchAlert struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Active              *Bool      `xmlrpc:"active,omptempty"`
	AvailableToday      *Bool      `xmlrpc:"available_today,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	LocationIds         *Relation  `xmlrpc:"location_ids,omptempty"`
	Message             *String    `xmlrpc:"message,omptempty"`
	Mode                *Selection `xmlrpc:"mode,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	NotificationMoment  *Selection `xmlrpc:"notification_moment,omptempty"`
	NotificationTime    *Float     `xmlrpc:"notification_time,omptempty"`
	Recipients          *Selection `xmlrpc:"recipients,omptempty"`
	RecurrencyFriday    *Bool      `xmlrpc:"recurrency_friday,omptempty"`
	RecurrencyMonday    *Bool      `xmlrpc:"recurrency_monday,omptempty"`
	RecurrencySaturday  *Bool      `xmlrpc:"recurrency_saturday,omptempty"`
	RecurrencySunday    *Bool      `xmlrpc:"recurrency_sunday,omptempty"`
	RecurrencyThursday  *Bool      `xmlrpc:"recurrency_thursday,omptempty"`
	RecurrencyTuesday   *Bool      `xmlrpc:"recurrency_tuesday,omptempty"`
	RecurrencyWednesday *Bool      `xmlrpc:"recurrency_wednesday,omptempty"`
	Tz                  *Selection `xmlrpc:"tz,omptempty"`
	Until               *Time      `xmlrpc:"until,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LunchAlerts represents array of lunch.alert model.
type LunchAlerts []LunchAlert

// LunchAlertModel is the odoo model name.
const LunchAlertModel = "lunch.alert"

// Many2One convert LunchAlert to *Many2One.
func (la *LunchAlert) Many2One() *Many2One {
	return NewMany2One(la.Id.Get(), "")
}

// CreateLunchAlert creates a new lunch.alert model and returns its id.
func (c *Client) CreateLunchAlert(la *LunchAlert) (int64, error) {
	ids, err := c.CreateLunchAlerts([]*LunchAlert{la})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchAlert creates a new lunch.alert model and returns its id.
func (c *Client) CreateLunchAlerts(las []*LunchAlert) ([]int64, error) {
	var vv []interface{}
	for _, v := range las {
		vv = append(vv, v)
	}
	return c.Create(LunchAlertModel, vv)
}

// UpdateLunchAlert updates an existing lunch.alert record.
func (c *Client) UpdateLunchAlert(la *LunchAlert) error {
	return c.UpdateLunchAlerts([]int64{la.Id.Get()}, la)
}

// UpdateLunchAlerts updates existing lunch.alert records.
// All records (represented by ids) will be updated by la values.
func (c *Client) UpdateLunchAlerts(ids []int64, la *LunchAlert) error {
	return c.Update(LunchAlertModel, ids, la)
}

// DeleteLunchAlert deletes an existing lunch.alert record.
func (c *Client) DeleteLunchAlert(id int64) error {
	return c.DeleteLunchAlerts([]int64{id})
}

// DeleteLunchAlerts deletes existing lunch.alert records.
func (c *Client) DeleteLunchAlerts(ids []int64) error {
	return c.Delete(LunchAlertModel, ids)
}

// GetLunchAlert gets lunch.alert existing record.
func (c *Client) GetLunchAlert(id int64) (*LunchAlert, error) {
	las, err := c.GetLunchAlerts([]int64{id})
	if err != nil {
		return nil, err
	}
	if las != nil && len(*las) > 0 {
		return &((*las)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.alert not found", id)
}

// GetLunchAlerts gets lunch.alert existing records.
func (c *Client) GetLunchAlerts(ids []int64) (*LunchAlerts, error) {
	las := &LunchAlerts{}
	if err := c.Read(LunchAlertModel, ids, nil, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindLunchAlert finds lunch.alert record by querying it with criteria.
func (c *Client) FindLunchAlert(criteria *Criteria) (*LunchAlert, error) {
	las := &LunchAlerts{}
	if err := c.SearchRead(LunchAlertModel, criteria, NewOptions().Limit(1), las); err != nil {
		return nil, err
	}
	if las != nil && len(*las) > 0 {
		return &((*las)[0]), nil
	}
	return nil, fmt.Errorf("lunch.alert was not found with criteria %v", criteria)
}

// FindLunchAlerts finds lunch.alert records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchAlerts(criteria *Criteria, options *Options) (*LunchAlerts, error) {
	las := &LunchAlerts{}
	if err := c.SearchRead(LunchAlertModel, criteria, options, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindLunchAlertIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchAlertIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchAlertModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchAlertId finds record id by querying it with criteria.
func (c *Client) FindLunchAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.alert was not found with criteria %v and options %v", criteria, options)
}
