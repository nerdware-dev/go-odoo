package odoo

// SaleSubscriptionStage represents sale.subscription.stage model.
type SaleSubscriptionStage struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Description      *String   `xmlrpc:"description,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Fold             *Bool     `xmlrpc:"fold,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	InProgress       *Bool     `xmlrpc:"in_progress,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	RatingTemplateId *Many2One `xmlrpc:"rating_template_id,omitempty"`
	Sequence         *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleSubscriptionStages represents array of sale.subscription.stage model.
type SaleSubscriptionStages []SaleSubscriptionStage

// SaleSubscriptionStageModel is the odoo model name.
const SaleSubscriptionStageModel = "sale.subscription.stage"

// Many2One convert SaleSubscriptionStage to *Many2One.
func (sss *SaleSubscriptionStage) Many2One() *Many2One {
	return NewMany2One(sss.Id.Get(), "")
}

// CreateSaleSubscriptionStage creates a new sale.subscription.stage model and returns its id.
func (c *Client) CreateSaleSubscriptionStage(sss *SaleSubscriptionStage) (int64, error) {
	ids, err := c.CreateSaleSubscriptionStages([]*SaleSubscriptionStage{sss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionStage creates a new sale.subscription.stage model and returns its id.
func (c *Client) CreateSaleSubscriptionStages(ssss []*SaleSubscriptionStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssss {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionStageModel, vv, nil)
}

// UpdateSaleSubscriptionStage updates an existing sale.subscription.stage record.
func (c *Client) UpdateSaleSubscriptionStage(sss *SaleSubscriptionStage) error {
	return c.UpdateSaleSubscriptionStages([]int64{sss.Id.Get()}, sss)
}

// UpdateSaleSubscriptionStages updates existing sale.subscription.stage records.
// All records (represented by ids) will be updated by sss values.
func (c *Client) UpdateSaleSubscriptionStages(ids []int64, sss *SaleSubscriptionStage) error {
	return c.Update(SaleSubscriptionStageModel, ids, sss, nil)
}

// DeleteSaleSubscriptionStage deletes an existing sale.subscription.stage record.
func (c *Client) DeleteSaleSubscriptionStage(id int64) error {
	return c.DeleteSaleSubscriptionStages([]int64{id})
}

// DeleteSaleSubscriptionStages deletes existing sale.subscription.stage records.
func (c *Client) DeleteSaleSubscriptionStages(ids []int64) error {
	return c.Delete(SaleSubscriptionStageModel, ids)
}

// GetSaleSubscriptionStage gets sale.subscription.stage existing record.
func (c *Client) GetSaleSubscriptionStage(id int64) (*SaleSubscriptionStage, error) {
	ssss, err := c.GetSaleSubscriptionStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssss)[0]), nil
}

// GetSaleSubscriptionStages gets sale.subscription.stage existing records.
func (c *Client) GetSaleSubscriptionStages(ids []int64) (*SaleSubscriptionStages, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.Read(SaleSubscriptionStageModel, ids, nil, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionStage finds sale.subscription.stage record by querying it with criteria.
func (c *Client) FindSaleSubscriptionStage(criteria *Criteria) (*SaleSubscriptionStage, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.SearchRead(SaleSubscriptionStageModel, criteria, NewOptions().Limit(1), ssss); err != nil {
		return nil, err
	}
	return &((*ssss)[0]), nil
}

// FindSaleSubscriptionStages finds sale.subscription.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionStages(criteria *Criteria, options *Options) (*SaleSubscriptionStages, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.SearchRead(SaleSubscriptionStageModel, criteria, options, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionStageModel, criteria, options)
}

// FindSaleSubscriptionStageId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
