package odoo

import (
	"fmt"
)

// LunchSupplier represents lunch.supplier model.
type LunchSupplier struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AutomaticEmailTime          *Float     `xmlrpc:"automatic_email_time,omptempty"`
	AvailableLocationIds        *Relation  `xmlrpc:"available_location_ids,omptempty"`
	AvailableToday              *Bool      `xmlrpc:"available_today,omptempty"`
	City                        *String    `xmlrpc:"city,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Delivery                    *Selection `xmlrpc:"delivery,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Email                       *String    `xmlrpc:"email,omptempty"`
	EmailFormatted              *String    `xmlrpc:"email_formatted,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Moment                      *Selection `xmlrpc:"moment,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	Phone                       *String    `xmlrpc:"phone,omptempty"`
	RecurrencyEndDate           *Time      `xmlrpc:"recurrency_end_date,omptempty"`
	RecurrencyFriday            *Bool      `xmlrpc:"recurrency_friday,omptempty"`
	RecurrencyMonday            *Bool      `xmlrpc:"recurrency_monday,omptempty"`
	RecurrencySaturday          *Bool      `xmlrpc:"recurrency_saturday,omptempty"`
	RecurrencySunday            *Bool      `xmlrpc:"recurrency_sunday,omptempty"`
	RecurrencyThursday          *Bool      `xmlrpc:"recurrency_thursday,omptempty"`
	RecurrencyTuesday           *Bool      `xmlrpc:"recurrency_tuesday,omptempty"`
	RecurrencyWednesday         *Bool      `xmlrpc:"recurrency_wednesday,omptempty"`
	ResponsibleId               *Many2One  `xmlrpc:"responsible_id,omptempty"`
	SendBy                      *Selection `xmlrpc:"send_by,omptempty"`
	StateId                     *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                      *String    `xmlrpc:"street,omptempty"`
	Street2                     *String    `xmlrpc:"street2,omptempty"`
	Tz                          *Selection `xmlrpc:"tz,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	ZipCode                     *String    `xmlrpc:"zip_code,omptempty"`
}

// LunchSuppliers represents array of lunch.supplier model.
type LunchSuppliers []LunchSupplier

// LunchSupplierModel is the odoo model name.
const LunchSupplierModel = "lunch.supplier"

// Many2One convert LunchSupplier to *Many2One.
func (ls *LunchSupplier) Many2One() *Many2One {
	return NewMany2One(ls.Id.Get(), "")
}

// CreateLunchSupplier creates a new lunch.supplier model and returns its id.
func (c *Client) CreateLunchSupplier(ls *LunchSupplier) (int64, error) {
	ids, err := c.CreateLunchSuppliers([]*LunchSupplier{ls})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchSupplier creates a new lunch.supplier model and returns its id.
func (c *Client) CreateLunchSuppliers(lss []*LunchSupplier) ([]int64, error) {
	var vv []interface{}
	for _, v := range lss {
		vv = append(vv, v)
	}
	return c.Create(LunchSupplierModel, vv)
}

// UpdateLunchSupplier updates an existing lunch.supplier record.
func (c *Client) UpdateLunchSupplier(ls *LunchSupplier) error {
	return c.UpdateLunchSuppliers([]int64{ls.Id.Get()}, ls)
}

// UpdateLunchSuppliers updates existing lunch.supplier records.
// All records (represented by ids) will be updated by ls values.
func (c *Client) UpdateLunchSuppliers(ids []int64, ls *LunchSupplier) error {
	return c.Update(LunchSupplierModel, ids, ls)
}

// DeleteLunchSupplier deletes an existing lunch.supplier record.
func (c *Client) DeleteLunchSupplier(id int64) error {
	return c.DeleteLunchSuppliers([]int64{id})
}

// DeleteLunchSuppliers deletes existing lunch.supplier records.
func (c *Client) DeleteLunchSuppliers(ids []int64) error {
	return c.Delete(LunchSupplierModel, ids)
}

// GetLunchSupplier gets lunch.supplier existing record.
func (c *Client) GetLunchSupplier(id int64) (*LunchSupplier, error) {
	lss, err := c.GetLunchSuppliers([]int64{id})
	if err != nil {
		return nil, err
	}
	if lss != nil && len(*lss) > 0 {
		return &((*lss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lunch.supplier not found", id)
}

// GetLunchSuppliers gets lunch.supplier existing records.
func (c *Client) GetLunchSuppliers(ids []int64) (*LunchSuppliers, error) {
	lss := &LunchSuppliers{}
	if err := c.Read(LunchSupplierModel, ids, nil, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindLunchSupplier finds lunch.supplier record by querying it with criteria.
func (c *Client) FindLunchSupplier(criteria *Criteria) (*LunchSupplier, error) {
	lss := &LunchSuppliers{}
	if err := c.SearchRead(LunchSupplierModel, criteria, NewOptions().Limit(1), lss); err != nil {
		return nil, err
	}
	if lss != nil && len(*lss) > 0 {
		return &((*lss)[0]), nil
	}
	return nil, fmt.Errorf("lunch.supplier was not found with criteria %v", criteria)
}

// FindLunchSuppliers finds lunch.supplier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchSuppliers(criteria *Criteria, options *Options) (*LunchSuppliers, error) {
	lss := &LunchSuppliers{}
	if err := c.SearchRead(LunchSupplierModel, criteria, options, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindLunchSupplierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchSupplierIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LunchSupplierModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLunchSupplierId finds record id by querying it with criteria.
func (c *Client) FindLunchSupplierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchSupplierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lunch.supplier was not found with criteria %v and options %v", criteria, options)
}
