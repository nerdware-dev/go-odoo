package odoo

// SaleSubscriptionSnapshot represents sale.subscription.snapshot model.
type SaleSubscriptionSnapshot struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Date             *Time     `xmlrpc:"date,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	RecurringMonthly *Float    `xmlrpc:"recurring_monthly,omitempty"`
	SubscriptionId   *Many2One `xmlrpc:"subscription_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleSubscriptionSnapshots represents array of sale.subscription.snapshot model.
type SaleSubscriptionSnapshots []SaleSubscriptionSnapshot

// SaleSubscriptionSnapshotModel is the odoo model name.
const SaleSubscriptionSnapshotModel = "sale.subscription.snapshot"

// Many2One convert SaleSubscriptionSnapshot to *Many2One.
func (sss *SaleSubscriptionSnapshot) Many2One() *Many2One {
	return NewMany2One(sss.Id.Get(), "")
}

// CreateSaleSubscriptionSnapshot creates a new sale.subscription.snapshot model and returns its id.
func (c *Client) CreateSaleSubscriptionSnapshot(sss *SaleSubscriptionSnapshot) (int64, error) {
	ids, err := c.CreateSaleSubscriptionSnapshots([]*SaleSubscriptionSnapshot{sss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionSnapshot creates a new sale.subscription.snapshot model and returns its id.
func (c *Client) CreateSaleSubscriptionSnapshots(ssss []*SaleSubscriptionSnapshot) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssss {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionSnapshotModel, vv, nil)
}

// UpdateSaleSubscriptionSnapshot updates an existing sale.subscription.snapshot record.
func (c *Client) UpdateSaleSubscriptionSnapshot(sss *SaleSubscriptionSnapshot) error {
	return c.UpdateSaleSubscriptionSnapshots([]int64{sss.Id.Get()}, sss)
}

// UpdateSaleSubscriptionSnapshots updates existing sale.subscription.snapshot records.
// All records (represented by ids) will be updated by sss values.
func (c *Client) UpdateSaleSubscriptionSnapshots(ids []int64, sss *SaleSubscriptionSnapshot) error {
	return c.Update(SaleSubscriptionSnapshotModel, ids, sss, nil)
}

// DeleteSaleSubscriptionSnapshot deletes an existing sale.subscription.snapshot record.
func (c *Client) DeleteSaleSubscriptionSnapshot(id int64) error {
	return c.DeleteSaleSubscriptionSnapshots([]int64{id})
}

// DeleteSaleSubscriptionSnapshots deletes existing sale.subscription.snapshot records.
func (c *Client) DeleteSaleSubscriptionSnapshots(ids []int64) error {
	return c.Delete(SaleSubscriptionSnapshotModel, ids)
}

// GetSaleSubscriptionSnapshot gets sale.subscription.snapshot existing record.
func (c *Client) GetSaleSubscriptionSnapshot(id int64) (*SaleSubscriptionSnapshot, error) {
	ssss, err := c.GetSaleSubscriptionSnapshots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssss)[0]), nil
}

// GetSaleSubscriptionSnapshots gets sale.subscription.snapshot existing records.
func (c *Client) GetSaleSubscriptionSnapshots(ids []int64) (*SaleSubscriptionSnapshots, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.Read(SaleSubscriptionSnapshotModel, ids, nil, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionSnapshot finds sale.subscription.snapshot record by querying it with criteria.
func (c *Client) FindSaleSubscriptionSnapshot(criteria *Criteria) (*SaleSubscriptionSnapshot, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.SearchRead(SaleSubscriptionSnapshotModel, criteria, NewOptions().Limit(1), ssss); err != nil {
		return nil, err
	}
	return &((*ssss)[0]), nil
}

// FindSaleSubscriptionSnapshots finds sale.subscription.snapshot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionSnapshots(criteria *Criteria, options *Options) (*SaleSubscriptionSnapshots, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.SearchRead(SaleSubscriptionSnapshotModel, criteria, options, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionSnapshotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionSnapshotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionSnapshotModel, criteria, options)
}

// FindSaleSubscriptionSnapshotId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionSnapshotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionSnapshotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
