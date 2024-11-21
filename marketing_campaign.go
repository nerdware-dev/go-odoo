package odoo

// MarketingCampaign represents marketing.campaign model.
type MarketingCampaign struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Active                    *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	Bounced                   *Int       `xmlrpc:"bounced,omitempty" json:"bounced,omitempty"`
	BouncedRatio              *Int       `xmlrpc:"bounced_ratio,omitempty" json:"bounced_ratio,omitempty"`
	ClickCount                *Int       `xmlrpc:"click_count,omitempty" json:"click_count,omitempty"`
	Color                     *Int       `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CompletedParticipantCount *Int       `xmlrpc:"completed_participant_count,omitempty" json:"completed_participant_count,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CrmLeadActivated          *Bool      `xmlrpc:"crm_lead_activated,omitempty" json:"crm_lead_activated,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Delivered                 *Int       `xmlrpc:"delivered,omitempty" json:"delivered,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	Domain                    *String    `xmlrpc:"domain,omitempty" json:"domain,omitempty"`
	Failed                    *Int       `xmlrpc:"failed,omitempty" json:"failed,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Ignored                   *Int       `xmlrpc:"ignored,omitempty" json:"ignored,omitempty"`
	InvoicedAmount            *Int       `xmlrpc:"invoiced_amount,omitempty" json:"invoiced_amount,omitempty"`
	IsWebsite                 *Bool      `xmlrpc:"is_website,omitempty" json:"is_website,omitempty"`
	LastSyncDate              *Time      `xmlrpc:"last_sync_date,omitempty" json:"last_sync_date,omitempty"`
	LeadCount                 *Int       `xmlrpc:"lead_count,omitempty" json:"lead_count,omitempty"`
	LinkTrackerClickCount     *Int       `xmlrpc:"link_tracker_click_count,omitempty" json:"link_tracker_click_count,omitempty"`
	MailingClicked            *Int       `xmlrpc:"mailing_clicked,omitempty" json:"mailing_clicked,omitempty"`
	MailingClicksRatio        *Int       `xmlrpc:"mailing_clicks_ratio,omitempty" json:"mailing_clicks_ratio,omitempty"`
	MailingItems              *Int       `xmlrpc:"mailing_items,omitempty" json:"mailing_items,omitempty"`
	MailingMailCount          *Int       `xmlrpc:"mailing_mail_count,omitempty" json:"mailing_mail_count,omitempty"`
	MailingMailIds            *Relation  `xmlrpc:"mailing_mail_ids,omitempty" json:"mailing_mail_ids,omitempty"`
	MarketingActivityIds      *Relation  `xmlrpc:"marketing_activity_ids,omitempty" json:"marketing_activity_ids,omitempty"`
	MassMailingCount          *Int       `xmlrpc:"mass_mailing_count,omitempty" json:"mass_mailing_count,omitempty"`
	ModelId                   *Many2One  `xmlrpc:"model_id,omitempty" json:"model_id,omitempty"`
	ModelName                 *String    `xmlrpc:"model_name,omitempty" json:"model_name,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Opened                    *Int       `xmlrpc:"opened,omitempty" json:"opened,omitempty"`
	OpenedRatio               *Int       `xmlrpc:"opened_ratio,omitempty" json:"opened_ratio,omitempty"`
	OpportunityCount          *Int       `xmlrpc:"opportunity_count,omitempty" json:"opportunity_count,omitempty"`
	ParticipantIds            *Relation  `xmlrpc:"participant_ids,omitempty" json:"participant_ids,omitempty"`
	QuotationCount            *Int       `xmlrpc:"quotation_count,omitempty" json:"quotation_count,omitempty"`
	ReceivedRatio             *Int       `xmlrpc:"received_ratio,omitempty" json:"received_ratio,omitempty"`
	Replied                   *Int       `xmlrpc:"replied,omitempty" json:"replied,omitempty"`
	RepliedRatio              *Int       `xmlrpc:"replied_ratio,omitempty" json:"replied_ratio,omitempty"`
	RequireSync               *Bool      `xmlrpc:"require_sync,omitempty" json:"require_sync,omitempty"`
	RunningParticipantCount   *Int       `xmlrpc:"running_participant_count,omitempty" json:"running_participant_count,omitempty"`
	Scheduled                 *Int       `xmlrpc:"scheduled,omitempty" json:"scheduled,omitempty"`
	Sent                      *Int       `xmlrpc:"sent,omitempty" json:"sent,omitempty"`
	StageId                   *Many2One  `xmlrpc:"stage_id,omitempty" json:"stage_id,omitempty"`
	State                     *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	TagIds                    *Relation  `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TestParticipantCount      *Int       `xmlrpc:"test_participant_count,omitempty" json:"test_participant_count,omitempty"`
	Total                     *Int       `xmlrpc:"total,omitempty" json:"total,omitempty"`
	TotalParticipantCount     *Int       `xmlrpc:"total_participant_count,omitempty" json:"total_participant_count,omitempty"`
	UniqueFieldId             *Many2One  `xmlrpc:"unique_field_id,omitempty" json:"unique_field_id,omitempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	UtmCampaignId             *Many2One  `xmlrpc:"utm_campaign_id,omitempty" json:"utm_campaign_id,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// MarketingCampaigns represents array of marketing.campaign model.
type MarketingCampaigns []MarketingCampaign

// MarketingCampaignModel is the odoo model name.
const MarketingCampaignModel = "marketing.campaign"

// Many2One convert MarketingCampaign to *Many2One.
func (mc *MarketingCampaign) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMarketingCampaign creates a new marketing.campaign model and returns its id.
func (c *Client) CreateMarketingCampaign(mc *MarketingCampaign) (int64, error) {
	ids, err := c.CreateMarketingCampaigns([]*MarketingCampaign{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingCampaign creates a new marketing.campaign model and returns its id.
func (c *Client) CreateMarketingCampaigns(mcs []*MarketingCampaign) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MarketingCampaignModel, vv, nil)
}

// UpdateMarketingCampaign updates an existing marketing.campaign record.
func (c *Client) UpdateMarketingCampaign(mc *MarketingCampaign) error {
	return c.UpdateMarketingCampaigns([]int64{mc.Id.Get()}, mc)
}

// UpdateMarketingCampaigns updates existing marketing.campaign records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMarketingCampaigns(ids []int64, mc *MarketingCampaign) error {
	return c.Update(MarketingCampaignModel, ids, mc, nil)
}

// DeleteMarketingCampaign deletes an existing marketing.campaign record.
func (c *Client) DeleteMarketingCampaign(id int64) error {
	return c.DeleteMarketingCampaigns([]int64{id})
}

// DeleteMarketingCampaigns deletes existing marketing.campaign records.
func (c *Client) DeleteMarketingCampaigns(ids []int64) error {
	return c.Delete(MarketingCampaignModel, ids)
}

// GetMarketingCampaign gets marketing.campaign existing record.
func (c *Client) GetMarketingCampaign(id int64) (*MarketingCampaign, error) {
	mcs, err := c.GetMarketingCampaigns([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// GetMarketingCampaigns gets marketing.campaign existing records.
func (c *Client) GetMarketingCampaigns(ids []int64) (*MarketingCampaigns, error) {
	mcs := &MarketingCampaigns{}
	if err := c.Read(MarketingCampaignModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMarketingCampaign finds marketing.campaign record by querying it with criteria.
func (c *Client) FindMarketingCampaign(criteria *Criteria) (*MarketingCampaign, error) {
	mcs := &MarketingCampaigns{}
	if err := c.SearchRead(MarketingCampaignModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// FindMarketingCampaigns finds marketing.campaign records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaigns(criteria *Criteria, options *Options) (*MarketingCampaigns, error) {
	mcs := &MarketingCampaigns{}
	if err := c.SearchRead(MarketingCampaignModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMarketingCampaignIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MarketingCampaignModel, criteria, options)
}

// FindMarketingCampaignId finds record id by querying it with criteria.
func (c *Client) FindMarketingCampaignId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingCampaignModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
