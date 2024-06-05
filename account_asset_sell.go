package odoo

import (
	"fmt"
)

// AccountAssetSell represents account.asset.sell model.
type AccountAssetSell struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Action              *Selection `xmlrpc:"action,omptempty"`
	AssetId             *Many2One  `xmlrpc:"asset_id,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	GainAccountId       *Many2One  `xmlrpc:"gain_account_id,omptempty"`
	GainOrLoss          *Selection `xmlrpc:"gain_or_loss,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InvoiceId           *Many2One  `xmlrpc:"invoice_id,omptempty"`
	InvoiceLineId       *Many2One  `xmlrpc:"invoice_line_id,omptempty"`
	LossAccountId       *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	SelectInvoiceLineId *Bool      `xmlrpc:"select_invoice_line_id,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(AccountAssetSellModel, vv)
}

// UpdateAccountAssetSell updates an existing account.asset.sell record.
func (c *Client) UpdateAccountAssetSell(aas *AccountAssetSell) error {
	return c.UpdateAccountAssetSells([]int64{aas.Id.Get()}, aas)
}

// UpdateAccountAssetSells updates existing account.asset.sell records.
// All records (represented by ids) will be updated by aas values.
func (c *Client) UpdateAccountAssetSells(ids []int64, aas *AccountAssetSell) error {
	return c.Update(AccountAssetSellModel, ids, aas)
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
	if aass != nil && len(*aass) > 0 {
		return &((*aass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.asset.sell not found", id)
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
	if aass != nil && len(*aass) > 0 {
		return &((*aass)[0]), nil
	}
	return nil, fmt.Errorf("account.asset.sell was not found with criteria %v", criteria)
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
	ids, err := c.Search(AccountAssetSellModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAssetSellId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetSellId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetSellModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.asset.sell was not found with criteria %v and options %v", criteria, options)
}
