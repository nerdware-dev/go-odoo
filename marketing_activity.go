package odoo

// MarketingActivity represents marketing.activity model.
type MarketingActivity struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDomain           *String    `xmlrpc:"activity_domain,omitempty"`
	ActivityType             *Selection `xmlrpc:"activity_type,omitempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omitempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Domain                   *String    `xmlrpc:"domain,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IntervalNumber           *Int       `xmlrpc:"interval_number,omitempty"`
	IntervalStandardized     *Int       `xmlrpc:"interval_standardized,omitempty"`
	IntervalType             *Selection `xmlrpc:"interval_type,omitempty"`
	MassMailingId            *Many2One  `xmlrpc:"mass_mailing_id,omitempty"`
	MassMailingIdMailingType *Selection `xmlrpc:"mass_mailing_id_mailing_type,omitempty"`
	ModelId                  *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelName                *String    `xmlrpc:"model_name,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omitempty"`
	Processed                *Int       `xmlrpc:"processed,omitempty"`
	Rejected                 *Int       `xmlrpc:"rejected,omitempty"`
	RequireSync              *Bool      `xmlrpc:"require_sync,omitempty"`
	ServerActionId           *Many2One  `xmlrpc:"server_action_id,omitempty"`
	StatisticsGraphData      *String    `xmlrpc:"statistics_graph_data,omitempty"`
	TotalBounce              *Int       `xmlrpc:"total_bounce,omitempty"`
	TotalClick               *Int       `xmlrpc:"total_click,omitempty"`
	TotalOpen                *Int       `xmlrpc:"total_open,omitempty"`
	TotalReply               *Int       `xmlrpc:"total_reply,omitempty"`
	TotalSent                *Int       `xmlrpc:"total_sent,omitempty"`
	TraceIds                 *Relation  `xmlrpc:"trace_ids,omitempty"`
	TriggerType              *Selection `xmlrpc:"trigger_type,omitempty"`
	UtmCampaignId            *Many2One  `xmlrpc:"utm_campaign_id,omitempty"`
	UtmSourceId              *Many2One  `xmlrpc:"utm_source_id,omitempty"`
	ValidityDuration         *Bool      `xmlrpc:"validity_duration,omitempty"`
	ValidityDurationNumber   *Int       `xmlrpc:"validity_duration_number,omitempty"`
	ValidityDurationType     *Selection `xmlrpc:"validity_duration_type,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MarketingActivitys represents array of marketing.activity model.
type MarketingActivitys []MarketingActivity

// MarketingActivityModel is the odoo model name.
const MarketingActivityModel = "marketing.activity"

// Many2One convert MarketingActivity to *Many2One.
func (ma *MarketingActivity) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMarketingActivity creates a new marketing.activity model and returns its id.
func (c *Client) CreateMarketingActivity(ma *MarketingActivity) (int64, error) {
	ids, err := c.CreateMarketingActivitys([]*MarketingActivity{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingActivity creates a new marketing.activity model and returns its id.
func (c *Client) CreateMarketingActivitys(mas []*MarketingActivity) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MarketingActivityModel, vv, nil)
}

// UpdateMarketingActivity updates an existing marketing.activity record.
func (c *Client) UpdateMarketingActivity(ma *MarketingActivity) error {
	return c.UpdateMarketingActivitys([]int64{ma.Id.Get()}, ma)
}

// UpdateMarketingActivitys updates existing marketing.activity records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMarketingActivitys(ids []int64, ma *MarketingActivity) error {
	return c.Update(MarketingActivityModel, ids, ma, nil)
}

// DeleteMarketingActivity deletes an existing marketing.activity record.
func (c *Client) DeleteMarketingActivity(id int64) error {
	return c.DeleteMarketingActivitys([]int64{id})
}

// DeleteMarketingActivitys deletes existing marketing.activity records.
func (c *Client) DeleteMarketingActivitys(ids []int64) error {
	return c.Delete(MarketingActivityModel, ids)
}

// GetMarketingActivity gets marketing.activity existing record.
func (c *Client) GetMarketingActivity(id int64) (*MarketingActivity, error) {
	mas, err := c.GetMarketingActivitys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mas)[0]), nil
}

// GetMarketingActivitys gets marketing.activity existing records.
func (c *Client) GetMarketingActivitys(ids []int64) (*MarketingActivitys, error) {
	mas := &MarketingActivitys{}
	if err := c.Read(MarketingActivityModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMarketingActivity finds marketing.activity record by querying it with criteria.
func (c *Client) FindMarketingActivity(criteria *Criteria) (*MarketingActivity, error) {
	mas := &MarketingActivitys{}
	if err := c.SearchRead(MarketingActivityModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	return &((*mas)[0]), nil
}

// FindMarketingActivitys finds marketing.activity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingActivitys(criteria *Criteria, options *Options) (*MarketingActivitys, error) {
	mas := &MarketingActivitys{}
	if err := c.SearchRead(MarketingActivityModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMarketingActivityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingActivityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MarketingActivityModel, criteria, options)
}

// FindMarketingActivityId finds record id by querying it with criteria.
func (c *Client) FindMarketingActivityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingActivityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
