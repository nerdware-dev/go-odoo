package odoo

// AccountBatchDownloadWizard represents account.batch.download.wizard model.
type AccountBatchDownloadWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	BatchPaymentId *Many2One `xmlrpc:"batch_payment_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	ExportFile     *String   `xmlrpc:"export_file,omitempty"`
	ExportFilename *String   `xmlrpc:"export_filename,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	WarningMessage *String   `xmlrpc:"warning_message,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(AccountBatchDownloadWizardModel, vv, nil)
}

// UpdateAccountBatchDownloadWizard updates an existing account.batch.download.wizard record.
func (c *Client) UpdateAccountBatchDownloadWizard(abdw *AccountBatchDownloadWizard) error {
	return c.UpdateAccountBatchDownloadWizards([]int64{abdw.Id.Get()}, abdw)
}

// UpdateAccountBatchDownloadWizards updates existing account.batch.download.wizard records.
// All records (represented by ids) will be updated by abdw values.
func (c *Client) UpdateAccountBatchDownloadWizards(ids []int64, abdw *AccountBatchDownloadWizard) error {
	return c.Update(AccountBatchDownloadWizardModel, ids, abdw, nil)
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
	return &((*abdws)[0]), nil
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
	return &((*abdws)[0]), nil
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
	return c.Search(AccountBatchDownloadWizardModel, criteria, options)
}

// FindAccountBatchDownloadWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchDownloadWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchDownloadWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
