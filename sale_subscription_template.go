package odoo

import (
	"fmt"
)

// SaleSubscriptionTemplate represents sale.subscription.template model.
type SaleSubscriptionTemplate struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	Active                    *Bool      `xmlrpc:"active,omptempty"`
	AutoCloseLimit            *Int       `xmlrpc:"auto_close_limit,omptempty"`
	BadHealthDomain           *String    `xmlrpc:"bad_health_domain,omptempty"`
	Code                      *String    `xmlrpc:"code,omptempty"`
	Color                     *Int       `xmlrpc:"color,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description               *String    `xmlrpc:"description,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	GoodHealthDomain          *String    `xmlrpc:"good_health_domain,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	InvoiceMailTemplateId     *Many2One  `xmlrpc:"invoice_mail_template_id,omptempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omptempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	PaymentMode               *Selection `xmlrpc:"payment_mode,omptempty"`
	ProductCount              *Int       `xmlrpc:"product_count,omptempty"`
	ProductIds                *Relation  `xmlrpc:"product_ids,omptempty"`
	RecurringInterval         *Int       `xmlrpc:"recurring_interval,omptempty"`
	RecurringRuleBoundary     *Selection `xmlrpc:"recurring_rule_boundary,omptempty"`
	RecurringRuleCount        *Int       `xmlrpc:"recurring_rule_count,omptempty"`
	RecurringRuleType         *Selection `xmlrpc:"recurring_rule_type,omptempty"`
	RecurringRuleTypeReadonly *Selection `xmlrpc:"recurring_rule_type_readonly,omptempty"`
	SubscriptionCount         *Int       `xmlrpc:"subscription_count,omptempty"`
	TagIds                    *Relation  `xmlrpc:"tag_ids,omptempty"`
	UserClosable              *Bool      `xmlrpc:"user_closable,omptempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionTemplates represents array of sale.subscription.template model.
type SaleSubscriptionTemplates []SaleSubscriptionTemplate

// SaleSubscriptionTemplateModel is the odoo model name.
const SaleSubscriptionTemplateModel = "sale.subscription.template"

// Many2One convert SaleSubscriptionTemplate to *Many2One.
func (sst *SaleSubscriptionTemplate) Many2One() *Many2One {
	return NewMany2One(sst.Id.Get(), "")
}

// CreateSaleSubscriptionTemplate creates a new sale.subscription.template model and returns its id.
func (c *Client) CreateSaleSubscriptionTemplate(sst *SaleSubscriptionTemplate) (int64, error) {
	ids, err := c.CreateSaleSubscriptionTemplates([]*SaleSubscriptionTemplate{sst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionTemplate creates a new sale.subscription.template model and returns its id.
func (c *Client) CreateSaleSubscriptionTemplates(ssts []*SaleSubscriptionTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssts {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionTemplateModel, vv)
}

// UpdateSaleSubscriptionTemplate updates an existing sale.subscription.template record.
func (c *Client) UpdateSaleSubscriptionTemplate(sst *SaleSubscriptionTemplate) error {
	return c.UpdateSaleSubscriptionTemplates([]int64{sst.Id.Get()}, sst)
}

// UpdateSaleSubscriptionTemplates updates existing sale.subscription.template records.
// All records (represented by ids) will be updated by sst values.
func (c *Client) UpdateSaleSubscriptionTemplates(ids []int64, sst *SaleSubscriptionTemplate) error {
	return c.Update(SaleSubscriptionTemplateModel, ids, sst)
}

// DeleteSaleSubscriptionTemplate deletes an existing sale.subscription.template record.
func (c *Client) DeleteSaleSubscriptionTemplate(id int64) error {
	return c.DeleteSaleSubscriptionTemplates([]int64{id})
}

// DeleteSaleSubscriptionTemplates deletes existing sale.subscription.template records.
func (c *Client) DeleteSaleSubscriptionTemplates(ids []int64) error {
	return c.Delete(SaleSubscriptionTemplateModel, ids)
}

// GetSaleSubscriptionTemplate gets sale.subscription.template existing record.
func (c *Client) GetSaleSubscriptionTemplate(id int64) (*SaleSubscriptionTemplate, error) {
	ssts, err := c.GetSaleSubscriptionTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.template not found", id)
}

// GetSaleSubscriptionTemplates gets sale.subscription.template existing records.
func (c *Client) GetSaleSubscriptionTemplates(ids []int64) (*SaleSubscriptionTemplates, error) {
	ssts := &SaleSubscriptionTemplates{}
	if err := c.Read(SaleSubscriptionTemplateModel, ids, nil, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaleSubscriptionTemplate finds sale.subscription.template record by querying it with criteria.
func (c *Client) FindSaleSubscriptionTemplate(criteria *Criteria) (*SaleSubscriptionTemplate, error) {
	ssts := &SaleSubscriptionTemplates{}
	if err := c.SearchRead(SaleSubscriptionTemplateModel, criteria, NewOptions().Limit(1), ssts); err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.template was not found with criteria %v", criteria)
}

// FindSaleSubscriptionTemplates finds sale.subscription.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionTemplates(criteria *Criteria, options *Options) (*SaleSubscriptionTemplates, error) {
	ssts := &SaleSubscriptionTemplates{}
	if err := c.SearchRead(SaleSubscriptionTemplateModel, criteria, options, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaleSubscriptionTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionTemplateId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.template was not found with criteria %v and options %v", criteria, options)
}
