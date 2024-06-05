package odoo

import (
	"fmt"
)

// AccountAssetPause represents account.asset.pause model.
type AccountAssetPause struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AssetId     *Many2One `xmlrpc:"asset_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Date        *Time     `xmlrpc:"date,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountAssetPauses represents array of account.asset.pause model.
type AccountAssetPauses []AccountAssetPause

// AccountAssetPauseModel is the odoo model name.
const AccountAssetPauseModel = "account.asset.pause"

// Many2One convert AccountAssetPause to *Many2One.
func (aap *AccountAssetPause) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAssetPause creates a new account.asset.pause model and returns its id.
func (c *Client) CreateAccountAssetPause(aap *AccountAssetPause) (int64, error) {
	ids, err := c.CreateAccountAssetPauses([]*AccountAssetPause{aap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAssetPause creates a new account.asset.pause model and returns its id.
func (c *Client) CreateAccountAssetPauses(aaps []*AccountAssetPause) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaps {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetPauseModel, vv)
}

// UpdateAccountAssetPause updates an existing account.asset.pause record.
func (c *Client) UpdateAccountAssetPause(aap *AccountAssetPause) error {
	return c.UpdateAccountAssetPauses([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAssetPauses updates existing account.asset.pause records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAssetPauses(ids []int64, aap *AccountAssetPause) error {
	return c.Update(AccountAssetPauseModel, ids, aap)
}

// DeleteAccountAssetPause deletes an existing account.asset.pause record.
func (c *Client) DeleteAccountAssetPause(id int64) error {
	return c.DeleteAccountAssetPauses([]int64{id})
}

// DeleteAccountAssetPauses deletes existing account.asset.pause records.
func (c *Client) DeleteAccountAssetPauses(ids []int64) error {
	return c.Delete(AccountAssetPauseModel, ids)
}

// GetAccountAssetPause gets account.asset.pause existing record.
func (c *Client) GetAccountAssetPause(id int64) (*AccountAssetPause, error) {
	aaps, err := c.GetAccountAssetPauses([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.asset.pause not found", id)
}

// GetAccountAssetPauses gets account.asset.pause existing records.
func (c *Client) GetAccountAssetPauses(ids []int64) (*AccountAssetPauses, error) {
	aaps := &AccountAssetPauses{}
	if err := c.Read(AccountAssetPauseModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAssetPause finds account.asset.pause record by querying it with criteria.
func (c *Client) FindAccountAssetPause(criteria *Criteria) (*AccountAssetPause, error) {
	aaps := &AccountAssetPauses{}
	if err := c.SearchRead(AccountAssetPauseModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("account.asset.pause was not found with criteria %v", criteria)
}

// FindAccountAssetPauses finds account.asset.pause records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetPauses(criteria *Criteria, options *Options) (*AccountAssetPauses, error) {
	aaps := &AccountAssetPauses{}
	if err := c.SearchRead(AccountAssetPauseModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAssetPauseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetPauseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAssetPauseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAssetPauseId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetPauseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetPauseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.asset.pause was not found with criteria %v and options %v", criteria, options)
}
