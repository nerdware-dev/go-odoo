package odoo

import (
	"fmt"
)

// AccountBatchDownloadWizard represents account.batch.download.wizard model.
type AccountBatchDownloadWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	BatchPaymentId *Many2One `xmlrpc:"batch_payment_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	ExportFile     *String   `xmlrpc:"export_file,omptempty"`
	ExportFilename *String   `xmlrpc:"export_filename,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	WarningMessage *String   `xmlrpc:"warning_message,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBatchDownloadWizards represents array of account.batch.download.wizard model.
type AccountBatchDownloadWizards []AccountBatchDownloadWizard

// AccountBatchDownloadWizardModel is the odoo model name.
const AccountBatchDownloadWizardModel = "account.batch.download.wizard"

// Many2One convert AccountBatchDownloadWizard to *Many2One.
func (abdw *AccountBatchDownloadWizard) Many2One() *Many2One {
	return NewMany2One(abdw.Id.Get(), "")
}

// CreateAccountBatchDownloadWizard creates a new account.batch.download.wizard model and returns its id.
func (c *Client) CreateAccountBatchDownloadWizard(abdw *AccountBatchDownloadWizard) (int64, error) {
	ids, err := c.CreateAccountBatchDownloadWizards([]*AccountBatchDownloadWizard{abdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchDownloadWizard creates a new account.batch.download.wizard model and returns its id.
func (c *Client) CreateAccountBatchDownloadWizards(abdws []*AccountBatchDownloadWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range abdws {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchDownloadWizardModel, vv)
}

// UpdateAccountBatchDownloadWizard updates an existing account.batch.download.wizard record.
func (c *Client) UpdateAccountBatchDownloadWizard(abdw *AccountBatchDownloadWizard) error {
	return c.UpdateAccountBatchDownloadWizards([]int64{abdw.Id.Get()}, abdw)
}

// UpdateAccountBatchDownloadWizards updates existing account.batch.download.wizard records.
// All records (represented by ids) will be updated by abdw values.
func (c *Client) UpdateAccountBatchDownloadWizards(ids []int64, abdw *AccountBatchDownloadWizard) error {
	return c.Update(AccountBatchDownloadWizardModel, ids, abdw)
}

// DeleteAccountBatchDownloadWizard deletes an existing account.batch.download.wizard record.
func (c *Client) DeleteAccountBatchDownloadWizard(id int64) error {
	return c.DeleteAccountBatchDownloadWizards([]int64{id})
}

// DeleteAccountBatchDownloadWizards deletes existing account.batch.download.wizard records.
func (c *Client) DeleteAccountBatchDownloadWizards(ids []int64) error {
	return c.Delete(AccountBatchDownloadWizardModel, ids)
}

// GetAccountBatchDownloadWizard gets account.batch.download.wizard existing record.
func (c *Client) GetAccountBatchDownloadWizard(id int64) (*AccountBatchDownloadWizard, error) {
	abdws, err := c.GetAccountBatchDownloadWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if abdws != nil && len(*abdws) > 0 {
		return &((*abdws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.batch.download.wizard not found", id)
}

// GetAccountBatchDownloadWizards gets account.batch.download.wizard existing records.
func (c *Client) GetAccountBatchDownloadWizards(ids []int64) (*AccountBatchDownloadWizards, error) {
	abdws := &AccountBatchDownloadWizards{}
	if err := c.Read(AccountBatchDownloadWizardModel, ids, nil, abdws); err != nil {
		return nil, err
	}
	return abdws, nil
}

// FindAccountBatchDownloadWizard finds account.batch.download.wizard record by querying it with criteria.
func (c *Client) FindAccountBatchDownloadWizard(criteria *Criteria) (*AccountBatchDownloadWizard, error) {
	abdws := &AccountBatchDownloadWizards{}
	if err := c.SearchRead(AccountBatchDownloadWizardModel, criteria, NewOptions().Limit(1), abdws); err != nil {
		return nil, err
	}
	if abdws != nil && len(*abdws) > 0 {
		return &((*abdws)[0]), nil
	}
	return nil, fmt.Errorf("account.batch.download.wizard was not found with criteria %v", criteria)
}

// FindAccountBatchDownloadWizards finds account.batch.download.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchDownloadWizards(criteria *Criteria, options *Options) (*AccountBatchDownloadWizards, error) {
	abdws := &AccountBatchDownloadWizards{}
	if err := c.SearchRead(AccountBatchDownloadWizardModel, criteria, options, abdws); err != nil {
		return nil, err
	}
	return abdws, nil
}

// FindAccountBatchDownloadWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchDownloadWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBatchDownloadWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBatchDownloadWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchDownloadWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchDownloadWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.batch.download.wizard was not found with criteria %v and options %v", criteria, options)
}
