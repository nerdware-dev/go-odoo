package odoo

// AccountOnlineWizard represents account.online.wizard model.
type AccountOnlineWizard struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omitempty"`
	AccountIds   *Relation  `xmlrpc:"account_ids,omitempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	HideTable    *Bool      `xmlrpc:"hide_table,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	Message      *String    `xmlrpc:"message,omitempty"`
	Method       *Selection `xmlrpc:"method,omitempty"`
	NumberAdded  *Int       `xmlrpc:"number_added,omitempty"`
	Status       *Selection `xmlrpc:"status,omitempty"`
	SyncDate     *Time      `xmlrpc:"sync_date,omitempty"`
	Transactions *String    `xmlrpc:"transactions,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountOnlineWizards represents array of account.online.wizard model.
type AccountOnlineWizards []AccountOnlineWizard

// AccountOnlineWizardModel is the odoo model name.
const AccountOnlineWizardModel = "account.online.wizard"

// Many2One convert AccountOnlineWizard to *Many2One.
func (aow *AccountOnlineWizard) Many2One() *Many2One {
	return NewMany2One(aow.Id.Get(), "")
}

// CreateAccountOnlineWizard creates a new account.online.wizard model and returns its id.
func (c *Client) CreateAccountOnlineWizard(aow *AccountOnlineWizard) (int64, error) {
	ids, err := c.CreateAccountOnlineWizards([]*AccountOnlineWizard{aow})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineWizard creates a new account.online.wizard model and returns its id.
func (c *Client) CreateAccountOnlineWizards(aows []*AccountOnlineWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aows {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineWizardModel, vv, nil)
}

// UpdateAccountOnlineWizard updates an existing account.online.wizard record.
func (c *Client) UpdateAccountOnlineWizard(aow *AccountOnlineWizard) error {
	return c.UpdateAccountOnlineWizards([]int64{aow.Id.Get()}, aow)
}

// UpdateAccountOnlineWizards updates existing account.online.wizard records.
// All records (represented by ids) will be updated by aow values.
func (c *Client) UpdateAccountOnlineWizards(ids []int64, aow *AccountOnlineWizard) error {
	return c.Update(AccountOnlineWizardModel, ids, aow, nil)
}

// DeleteAccountOnlineWizard deletes an existing account.online.wizard record.
func (c *Client) DeleteAccountOnlineWizard(id int64) error {
	return c.DeleteAccountOnlineWizards([]int64{id})
}

// DeleteAccountOnlineWizards deletes existing account.online.wizard records.
func (c *Client) DeleteAccountOnlineWizards(ids []int64) error {
	return c.Delete(AccountOnlineWizardModel, ids)
}

// GetAccountOnlineWizard gets account.online.wizard existing record.
func (c *Client) GetAccountOnlineWizard(id int64) (*AccountOnlineWizard, error) {
	aows, err := c.GetAccountOnlineWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aows)[0]), nil
}

// GetAccountOnlineWizards gets account.online.wizard existing records.
func (c *Client) GetAccountOnlineWizards(ids []int64) (*AccountOnlineWizards, error) {
	aows := &AccountOnlineWizards{}
	if err := c.Read(AccountOnlineWizardModel, ids, nil, aows); err != nil {
		return nil, err
	}
	return aows, nil
}

// FindAccountOnlineWizard finds account.online.wizard record by querying it with criteria.
func (c *Client) FindAccountOnlineWizard(criteria *Criteria) (*AccountOnlineWizard, error) {
	aows := &AccountOnlineWizards{}
	if err := c.SearchRead(AccountOnlineWizardModel, criteria, NewOptions().Limit(1), aows); err != nil {
		return nil, err
	}
	return &((*aows)[0]), nil
}

// FindAccountOnlineWizards finds account.online.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineWizards(criteria *Criteria, options *Options) (*AccountOnlineWizards, error) {
	aows := &AccountOnlineWizards{}
	if err := c.SearchRead(AccountOnlineWizardModel, criteria, options, aows); err != nil {
		return nil, err
	}
	return aows, nil
}

// FindAccountOnlineWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOnlineWizardModel, criteria, options)
}

// FindAccountOnlineWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
