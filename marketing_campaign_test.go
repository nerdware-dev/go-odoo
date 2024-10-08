package odoo

// MarketingCampaignTest represents marketing.campaign.test model.
type MarketingCampaignTest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CampaignId  *Many2One `xmlrpc:"campaign_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omitempty"`
	ModelName   *String   `xmlrpc:"model_name,omitempty"`
	ResId       *Int      `xmlrpc:"res_id,omitempty"`
	ResourceRef *String   `xmlrpc:"resource_ref,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MarketingCampaignTests represents array of marketing.campaign.test model.
type MarketingCampaignTests []MarketingCampaignTest

// MarketingCampaignTestModel is the odoo model name.
const MarketingCampaignTestModel = "marketing.campaign.test"

// Many2One convert MarketingCampaignTest to *Many2One.
func (mct *MarketingCampaignTest) Many2One() *Many2One {
	return NewMany2One(mct.Id.Get(), "")
}

// CreateMarketingCampaignTest creates a new marketing.campaign.test model and returns its id.
func (c *Client) CreateMarketingCampaignTest(mct *MarketingCampaignTest) (int64, error) {
	ids, err := c.CreateMarketingCampaignTests([]*MarketingCampaignTest{mct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingCampaignTest creates a new marketing.campaign.test model and returns its id.
func (c *Client) CreateMarketingCampaignTests(mcts []*MarketingCampaignTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcts {
		vv = append(vv, v)
	}
	return c.Create(MarketingCampaignTestModel, vv, nil)
}

// UpdateMarketingCampaignTest updates an existing marketing.campaign.test record.
func (c *Client) UpdateMarketingCampaignTest(mct *MarketingCampaignTest) error {
	return c.UpdateMarketingCampaignTests([]int64{mct.Id.Get()}, mct)
}

// UpdateMarketingCampaignTests updates existing marketing.campaign.test records.
// All records (represented by ids) will be updated by mct values.
func (c *Client) UpdateMarketingCampaignTests(ids []int64, mct *MarketingCampaignTest) error {
	return c.Update(MarketingCampaignTestModel, ids, mct, nil)
}

// DeleteMarketingCampaignTest deletes an existing marketing.campaign.test record.
func (c *Client) DeleteMarketingCampaignTest(id int64) error {
	return c.DeleteMarketingCampaignTests([]int64{id})
}

// DeleteMarketingCampaignTests deletes existing marketing.campaign.test records.
func (c *Client) DeleteMarketingCampaignTests(ids []int64) error {
	return c.Delete(MarketingCampaignTestModel, ids)
}

// GetMarketingCampaignTest gets marketing.campaign.test existing record.
func (c *Client) GetMarketingCampaignTest(id int64) (*MarketingCampaignTest, error) {
	mcts, err := c.GetMarketingCampaignTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcts)[0]), nil
}

// GetMarketingCampaignTests gets marketing.campaign.test existing records.
func (c *Client) GetMarketingCampaignTests(ids []int64) (*MarketingCampaignTests, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.Read(MarketingCampaignTestModel, ids, nil, mcts); err != nil {
		return nil, err
	}
	return mcts, nil
}

// FindMarketingCampaignTest finds marketing.campaign.test record by querying it with criteria.
func (c *Client) FindMarketingCampaignTest(criteria *Criteria) (*MarketingCampaignTest, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.SearchRead(MarketingCampaignTestModel, criteria, NewOptions().Limit(1), mcts); err != nil {
		return nil, err
	}
	return &((*mcts)[0]), nil
}

// FindMarketingCampaignTests finds marketing.campaign.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignTests(criteria *Criteria, options *Options) (*MarketingCampaignTests, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.SearchRead(MarketingCampaignTestModel, criteria, options, mcts); err != nil {
		return nil, err
	}
	return mcts, nil
}

// FindMarketingCampaignTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MarketingCampaignTestModel, criteria, options)
}

// FindMarketingCampaignTestId finds record id by querying it with criteria.
func (c *Client) FindMarketingCampaignTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingCampaignTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
