package odoo

import (
	"fmt"
)

// SaleSubscription represents sale.subscription model.
type SaleSubscription struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AnalyticAccountId           *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CloseReasonId               *Many2One  `xmlrpc:"close_reason_id,omptempty"`
	Code                        *String    `xmlrpc:"code,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                        *Time      `xmlrpc:"date,omptempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Health                      *Selection `xmlrpc:"health,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InProgress                  *Bool      `xmlrpc:"in_progress,omptempty"`
	IndustryId                  *Many2One  `xmlrpc:"industry_id,omptempty"`
	InvoiceCount                *Int       `xmlrpc:"invoice_count,omptempty"`
	Kpi1MonthMrrDelta           *Float     `xmlrpc:"kpi_1month_mrr_delta,omptempty"`
	Kpi1MonthMrrPercentage      *Float     `xmlrpc:"kpi_1month_mrr_percentage,omptempty"`
	Kpi3MonthsMrrDelta          *Float     `xmlrpc:"kpi_3months_mrr_delta,omptempty"`
	Kpi3MonthsMrrPercentage     *Float     `xmlrpc:"kpi_3months_mrr_percentage,omptempty"`
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
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentMode                 *Selection `xmlrpc:"payment_mode,omptempty"`
	PaymentTokenId              *Many2One  `xmlrpc:"payment_token_id,omptempty"`
	PercentageSatisfaction      *Int       `xmlrpc:"percentage_satisfaction,omptempty"`
	PricelistId                 *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	RatingAvg                   *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingCount                 *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback          *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage             *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastValue             *Float     `xmlrpc:"rating_last_value,omptempty"`
	RecurringAmountTax          *Float     `xmlrpc:"recurring_amount_tax,omptempty"`
	RecurringAmountTotal        *Float     `xmlrpc:"recurring_amount_total,omptempty"`
	RecurringInterval           *Int       `xmlrpc:"recurring_interval,omptempty"`
	RecurringInvoiceDay         *Int       `xmlrpc:"recurring_invoice_day,omptempty"`
	RecurringInvoiceLineIds     *Relation  `xmlrpc:"recurring_invoice_line_ids,omptempty"`
	RecurringMonthly            *Float     `xmlrpc:"recurring_monthly,omptempty"`
	RecurringNextDate           *Time      `xmlrpc:"recurring_next_date,omptempty"`
	RecurringRuleBoundary       *Selection `xmlrpc:"recurring_rule_boundary,omptempty"`
	RecurringRuleType           *Selection `xmlrpc:"recurring_rule_type,omptempty"`
	RecurringTotal              *Float     `xmlrpc:"recurring_total,omptempty"`
	SaleOrderCount              *Int       `xmlrpc:"sale_order_count,omptempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omptempty"`
	Starred                     *Bool      `xmlrpc:"starred,omptempty"`
	StarredUserIds              *Relation  `xmlrpc:"starred_user_ids,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omptempty"`
	TeamUserId                  *Many2One  `xmlrpc:"team_user_id,omptempty"`
	TemplateId                  *Many2One  `xmlrpc:"template_id,omptempty"`
	ToRenew                     *Bool      `xmlrpc:"to_renew,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	Uuid                        *String    `xmlrpc:"uuid,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteUrl                  *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(SaleSubscriptionModel, vv)
}

// UpdateSaleSubscription updates an existing sale.subscription record.
func (c *Client) UpdateSaleSubscription(ss *SaleSubscription) error {
	return c.UpdateSaleSubscriptions([]int64{ss.Id.Get()}, ss)
}

// UpdateSaleSubscriptions updates existing sale.subscription records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSaleSubscriptions(ids []int64, ss *SaleSubscription) error {
	return c.Update(SaleSubscriptionModel, ids, ss)
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
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription not found", id)
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
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription was not found with criteria %v", criteria)
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
	ids, err := c.Search(SaleSubscriptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription was not found with criteria %v and options %v", criteria, options)
}
