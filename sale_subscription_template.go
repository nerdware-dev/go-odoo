package odoo

// SaleSubscriptionTemplate represents sale.subscription.template model.
type SaleSubscriptionTemplate struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	AutoCloseLimit            *Int       `xmlrpc:"auto_close_limit,omitempty"`
	BadHealthDomain           *String    `xmlrpc:"bad_health_domain,omitempty"`
	Code                      *String    `xmlrpc:"code,omitempty"`
	Color                     *Int       `xmlrpc:"color,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description               *String    `xmlrpc:"description,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	GoodHealthDomain          *String    `xmlrpc:"good_health_domain,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	InvoiceMailTemplateId     *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PaymentMode               *Selection `xmlrpc:"payment_mode,omitempty"`
	ProductCount              *Int       `xmlrpc:"product_count,omitempty"`
	ProductIds                *Relation  `xmlrpc:"product_ids,omitempty"`
	RecurringInterval         *Int       `xmlrpc:"recurring_interval,omitempty"`
	RecurringRuleBoundary     *Selection `xmlrpc:"recurring_rule_boundary,omitempty"`
	RecurringRuleCount        *Int       `xmlrpc:"recurring_rule_count,omitempty"`
	RecurringRuleType         *Selection `xmlrpc:"recurring_rule_type,omitempty"`
	RecurringRuleTypeReadonly *Selection `xmlrpc:"recurring_rule_type_readonly,omitempty"`
	SubscriptionCount         *Int       `xmlrpc:"subscription_count,omitempty"`
	TagIds                    *Relation  `xmlrpc:"tag_ids,omitempty"`
	UserClosable              *Bool      `xmlrpc:"user_closable,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(SaleSubscriptionTemplateModel, vv, nil)
}

// UpdateSaleSubscriptionTemplate updates an existing sale.subscription.template record.
func (c *Client) UpdateSaleSubscriptionTemplate(sst *SaleSubscriptionTemplate) error {
	return c.UpdateSaleSubscriptionTemplates([]int64{sst.Id.Get()}, sst)
}

// UpdateSaleSubscriptionTemplates updates existing sale.subscription.template records.
// All records (represented by ids) will be updated by sst values.
func (c *Client) UpdateSaleSubscriptionTemplates(ids []int64, sst *SaleSubscriptionTemplate) error {
	return c.Update(SaleSubscriptionTemplateModel, ids, sst, nil)
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
	return &((*ssts)[0]), nil
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
	return &((*ssts)[0]), nil
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
	return c.Search(SaleSubscriptionTemplateModel, criteria, options)
}

// FindSaleSubscriptionTemplateId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
