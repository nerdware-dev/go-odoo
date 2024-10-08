package odoo

// AccountOnlineLinkWizard represents account.online.link.wizard model.
type AccountOnlineLinkWizard struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	AccountNumber             *String    `xmlrpc:"account_number,omitempty"`
	AccountOnlineWizardId     *Many2One  `xmlrpc:"account_online_wizard_id,omitempty"`
	Action                    *Selection `xmlrpc:"action,omitempty"`
	Balance                   *Float     `xmlrpc:"balance,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	JournalStatementsCreation *Selection `xmlrpc:"journal_statements_creation,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	OnlineAccountId           *Many2One  `xmlrpc:"online_account_id,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountOnlineLinkWizards represents array of account.online.link.wizard model.
type AccountOnlineLinkWizards []AccountOnlineLinkWizard

// AccountOnlineLinkWizardModel is the odoo model name.
const AccountOnlineLinkWizardModel = "account.online.link.wizard"

// Many2One convert AccountOnlineLinkWizard to *Many2One.
func (aolw *AccountOnlineLinkWizard) Many2One() *Many2One {
	return NewMany2One(aolw.Id.Get(), "")
}

// CreateAccountOnlineLinkWizard creates a new account.online.link.wizard model and returns its id.
func (c *Client) CreateAccountOnlineLinkWizard(aolw *AccountOnlineLinkWizard) (int64, error) {
	ids, err := c.CreateAccountOnlineLinkWizards([]*AccountOnlineLinkWizard{aolw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineLinkWizard creates a new account.online.link.wizard model and returns its id.
func (c *Client) CreateAccountOnlineLinkWizards(aolws []*AccountOnlineLinkWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aolws {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineLinkWizardModel, vv, nil)
}

// UpdateAccountOnlineLinkWizard updates an existing account.online.link.wizard record.
func (c *Client) UpdateAccountOnlineLinkWizard(aolw *AccountOnlineLinkWizard) error {
	return c.UpdateAccountOnlineLinkWizards([]int64{aolw.Id.Get()}, aolw)
}

// UpdateAccountOnlineLinkWizards updates existing account.online.link.wizard records.
// All records (represented by ids) will be updated by aolw values.
func (c *Client) UpdateAccountOnlineLinkWizards(ids []int64, aolw *AccountOnlineLinkWizard) error {
	return c.Update(AccountOnlineLinkWizardModel, ids, aolw, nil)
}

// DeleteAccountOnlineLinkWizard deletes an existing account.online.link.wizard record.
func (c *Client) DeleteAccountOnlineLinkWizard(id int64) error {
	return c.DeleteAccountOnlineLinkWizards([]int64{id})
}

// DeleteAccountOnlineLinkWizards deletes existing account.online.link.wizard records.
func (c *Client) DeleteAccountOnlineLinkWizards(ids []int64) error {
	return c.Delete(AccountOnlineLinkWizardModel, ids)
}

// GetAccountOnlineLinkWizard gets account.online.link.wizard existing record.
func (c *Client) GetAccountOnlineLinkWizard(id int64) (*AccountOnlineLinkWizard, error) {
	aolws, err := c.GetAccountOnlineLinkWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aolws)[0]), nil
}

// GetAccountOnlineLinkWizards gets account.online.link.wizard existing records.
func (c *Client) GetAccountOnlineLinkWizards(ids []int64) (*AccountOnlineLinkWizards, error) {
	aolws := &AccountOnlineLinkWizards{}
	if err := c.Read(AccountOnlineLinkWizardModel, ids, nil, aolws); err != nil {
		return nil, err
	}
	return aolws, nil
}

// FindAccountOnlineLinkWizard finds account.online.link.wizard record by querying it with criteria.
func (c *Client) FindAccountOnlineLinkWizard(criteria *Criteria) (*AccountOnlineLinkWizard, error) {
	aolws := &AccountOnlineLinkWizards{}
	if err := c.SearchRead(AccountOnlineLinkWizardModel, criteria, NewOptions().Limit(1), aolws); err != nil {
		return nil, err
	}
	return &((*aolws)[0]), nil
}

// FindAccountOnlineLinkWizards finds account.online.link.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinkWizards(criteria *Criteria, options *Options) (*AccountOnlineLinkWizards, error) {
	aolws := &AccountOnlineLinkWizards{}
	if err := c.SearchRead(AccountOnlineLinkWizardModel, criteria, options, aolws); err != nil {
		return nil, err
	}
	return aolws, nil
}

// FindAccountOnlineLinkWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinkWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOnlineLinkWizardModel, criteria, options)
}

// FindAccountOnlineLinkWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineLinkWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineLinkWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
