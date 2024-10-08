package odoo

// SaleSubscription represents sale.subscription model.
type SaleSubscription struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AnalyticAccountId           *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	CloseReasonId               *Many2One  `xmlrpc:"close_reason_id,omitempty"`
	Code                        *String    `xmlrpc:"code,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Health                      *Selection `xmlrpc:"health,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InProgress                  *Bool      `xmlrpc:"in_progress,omitempty"`
	IndustryId                  *Many2One  `xmlrpc:"industry_id,omitempty"`
	InvoiceCount                *Int       `xmlrpc:"invoice_count,omitempty"`
	Kpi1MonthMrrDelta           *Float     `xmlrpc:"kpi_1month_mrr_delta,omitempty"`
	Kpi1MonthMrrPercentage      *Float     `xmlrpc:"kpi_1month_mrr_percentage,omitempty"`
	Kpi3MonthsMrrDelta          *Float     `xmlrpc:"kpi_3months_mrr_delta,omitempty"`
	Kpi3MonthsMrrPercentage     *Float     `xmlrpc:"kpi_3months_mrr_percentage,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentMode                 *Selection `xmlrpc:"payment_mode,omitempty"`
	PaymentTokenId              *Many2One  `xmlrpc:"payment_token_id,omitempty"`
	PercentageSatisfaction      *Int       `xmlrpc:"percentage_satisfaction,omitempty"`
	PricelistId                 *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	RatingAvg                   *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingCount                 *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback          *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage             *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastValue             *Float     `xmlrpc:"rating_last_value,omitempty"`
	RecurringAmountTax          *Float     `xmlrpc:"recurring_amount_tax,omitempty"`
	RecurringAmountTotal        *Float     `xmlrpc:"recurring_amount_total,omitempty"`
	RecurringInterval           *Int       `xmlrpc:"recurring_interval,omitempty"`
	RecurringInvoiceDay         *Int       `xmlrpc:"recurring_invoice_day,omitempty"`
	RecurringInvoiceLineIds     *Relation  `xmlrpc:"recurring_invoice_line_ids,omitempty"`
	RecurringMonthly            *Float     `xmlrpc:"recurring_monthly,omitempty"`
	RecurringNextDate           *Time      `xmlrpc:"recurring_next_date,omitempty"`
	RecurringRuleBoundary       *Selection `xmlrpc:"recurring_rule_boundary,omitempty"`
	RecurringRuleType           *Selection `xmlrpc:"recurring_rule_type,omitempty"`
	RecurringTotal              *Float     `xmlrpc:"recurring_total,omitempty"`
	SaleOrderCount              *Int       `xmlrpc:"sale_order_count,omitempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omitempty"`
	Starred                     *Bool      `xmlrpc:"starred,omitempty"`
	StarredUserIds              *Relation  `xmlrpc:"starred_user_ids,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	TeamUserId                  *Many2One  `xmlrpc:"team_user_id,omitempty"`
	TemplateId                  *Many2One  `xmlrpc:"template_id,omitempty"`
	ToRenew                     *Bool      `xmlrpc:"to_renew,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	Uuid                        *String    `xmlrpc:"uuid,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteUrl                  *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SaleSubscriptions represents array of sale.subscription model.
type SaleSubscriptions []SaleSubscription

// SaleSubscriptionModel is the odoo model name.
const SaleSubscriptionModel = "sale.subscription"

// Many2One convert SaleSubscription to *Many2One.
func (ss *SaleSubscription) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSaleSubscription creates a new sale.subscription model and returns its id.
func (c *Client) CreateSaleSubscription(ss *SaleSubscription) (int64, error) {
	ids, err := c.CreateSaleSubscriptions([]*SaleSubscription{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscription creates a new sale.subscription model and returns its id.
func (c *Client) CreateSaleSubscriptions(sss []*SaleSubscription) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionModel, vv, nil)
}

// UpdateSaleSubscription updates an existing sale.subscription record.
func (c *Client) UpdateSaleSubscription(ss *SaleSubscription) error {
	return c.UpdateSaleSubscriptions([]int64{ss.Id.Get()}, ss)
}

// UpdateSaleSubscriptions updates existing sale.subscription records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSaleSubscriptions(ids []int64, ss *SaleSubscription) error {
	return c.Update(SaleSubscriptionModel, ids, ss, nil)
}

// DeleteSaleSubscription deletes an existing sale.subscription record.
func (c *Client) DeleteSaleSubscription(id int64) error {
	return c.DeleteSaleSubscriptions([]int64{id})
}

// DeleteSaleSubscriptions deletes existing sale.subscription records.
func (c *Client) DeleteSaleSubscriptions(ids []int64) error {
	return c.Delete(SaleSubscriptionModel, ids)
}

// GetSaleSubscription gets sale.subscription existing record.
func (c *Client) GetSaleSubscription(id int64) (*SaleSubscription, error) {
	sss, err := c.GetSaleSubscriptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// GetSaleSubscriptions gets sale.subscription existing records.
func (c *Client) GetSaleSubscriptions(ids []int64) (*SaleSubscriptions, error) {
	sss := &SaleSubscriptions{}
	if err := c.Read(SaleSubscriptionModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSaleSubscription finds sale.subscription record by querying it with criteria.
func (c *Client) FindSaleSubscription(criteria *Criteria) (*SaleSubscription, error) {
	sss := &SaleSubscriptions{}
	if err := c.SearchRead(SaleSubscriptionModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// FindSaleSubscriptions finds sale.subscription records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptions(criteria *Criteria, options *Options) (*SaleSubscriptions, error) {
	sss := &SaleSubscriptions{}
	if err := c.SearchRead(SaleSubscriptionModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSaleSubscriptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionModel, criteria, options)
}

// FindSaleSubscriptionId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
