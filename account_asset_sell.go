package odoo

// AccountAssetSell represents account.asset.sell model.
type AccountAssetSell struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	Action              *Selection `xmlrpc:"action,omitempty"`
	AssetId             *Many2One  `xmlrpc:"asset_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	GainAccountId       *Many2One  `xmlrpc:"gain_account_id,omitempty"`
	GainOrLoss          *Selection `xmlrpc:"gain_or_loss,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InvoiceId           *Many2One  `xmlrpc:"invoice_id,omitempty"`
	InvoiceLineId       *Many2One  `xmlrpc:"invoice_line_id,omitempty"`
	LossAccountId       *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	SelectInvoiceLineId *Bool      `xmlrpc:"select_invoice_line_id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAssetSells represents array of account.asset.sell model.
type AccountAssetSells []AccountAssetSell

// AccountAssetSellModel is the odoo model name.
const AccountAssetSellModel = "account.asset.sell"

// Many2One convert AccountAssetSell to *Many2One.
func (aas *AccountAssetSell) Many2One() *Many2One {
	return NewMany2One(aas.Id.Get(), "")
}

// CreateAccountAssetSell creates a new account.asset.sell model and returns its id.
func (c *Client) CreateAccountAssetSell(aas *AccountAssetSell) (int64, error) {
	ids, err := c.CreateAccountAssetSells([]*AccountAssetSell{aas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAssetSell creates a new account.asset.sell model and returns its id.
func (c *Client) CreateAccountAssetSells(aass []*AccountAssetSell) ([]int64, error) {
	var vv []interface{}
	for _, v := range aass {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetSellModel, vv, nil)
}

// UpdateAccountAssetSell updates an existing account.asset.sell record.
func (c *Client) UpdateAccountAssetSell(aas *AccountAssetSell) error {
	return c.UpdateAccountAssetSells([]int64{aas.Id.Get()}, aas)
}

// UpdateAccountAssetSells updates existing account.asset.sell records.
// All records (represented by ids) will be updated by aas values.
func (c *Client) UpdateAccountAssetSells(ids []int64, aas *AccountAssetSell) error {
	return c.Update(AccountAssetSellModel, ids, aas, nil)
}

// DeleteAccountAssetSell deletes an existing account.asset.sell record.
func (c *Client) DeleteAccountAssetSell(id int64) error {
	return c.DeleteAccountAssetSells([]int64{id})
}

// DeleteAccountAssetSells deletes existing account.asset.sell records.
func (c *Client) DeleteAccountAssetSells(ids []int64) error {
	return c.Delete(AccountAssetSellModel, ids)
}

// GetAccountAssetSell gets account.asset.sell existing record.
func (c *Client) GetAccountAssetSell(id int64) (*AccountAssetSell, error) {
	aass, err := c.GetAccountAssetSells([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aass)[0]), nil
}

// GetAccountAssetSells gets account.asset.sell existing records.
func (c *Client) GetAccountAssetSells(ids []int64) (*AccountAssetSells, error) {
	aass := &AccountAssetSells{}
	if err := c.Read(AccountAssetSellModel, ids, nil, aass); err != nil {
		return nil, err
	}
	return aass, nil
}

// FindAccountAssetSell finds account.asset.sell record by querying it with criteria.
func (c *Client) FindAccountAssetSell(criteria *Criteria) (*AccountAssetSell, error) {
	aass := &AccountAssetSells{}
	if err := c.SearchRead(AccountAssetSellModel, criteria, NewOptions().Limit(1), aass); err != nil {
		return nil, err
	}
	return &((*aass)[0]), nil
}

// FindAccountAssetSells finds account.asset.sell records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetSells(criteria *Criteria, options *Options) (*AccountAssetSells, error) {
	aass := &AccountAssetSells{}
	if err := c.SearchRead(AccountAssetSellModel, criteria, options, aass); err != nil {
		return nil, err
	}
	return aass, nil
}

// FindAccountAssetSellIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetSellIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAssetSellModel, criteria, options)
}

// FindAccountAssetSellId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetSellId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetSellModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
