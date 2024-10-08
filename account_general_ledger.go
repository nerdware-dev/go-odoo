package odoo

// AccountGeneralLedger represents account.general.ledger model.
type AccountGeneralLedger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountGeneralLedgers represents array of account.general.ledger model.
type AccountGeneralLedgers []AccountGeneralLedger

// AccountGeneralLedgerModel is the odoo model name.
const AccountGeneralLedgerModel = "account.general.ledger"

// Many2One convert AccountGeneralLedger to *Many2One.
func (agl *AccountGeneralLedger) Many2One() *Many2One {
	return NewMany2One(agl.Id.Get(), "")
}

// CreateAccountGeneralLedger creates a new account.general.ledger model and returns its id.
func (c *Client) CreateAccountGeneralLedger(agl *AccountGeneralLedger) (int64, error) {
	ids, err := c.CreateAccountGeneralLedgers([]*AccountGeneralLedger{agl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGeneralLedger creates a new account.general.ledger model and returns its id.
func (c *Client) CreateAccountGeneralLedgers(agls []*AccountGeneralLedger) ([]int64, error) {
	var vv []interface{}
	for _, v := range agls {
		vv = append(vv, v)
	}
	return c.Create(AccountGeneralLedgerModel, vv, nil)
}

// UpdateAccountGeneralLedger updates an existing account.general.ledger record.
func (c *Client) UpdateAccountGeneralLedger(agl *AccountGeneralLedger) error {
	return c.UpdateAccountGeneralLedgers([]int64{agl.Id.Get()}, agl)
}

// UpdateAccountGeneralLedgers updates existing account.general.ledger records.
// All records (represented by ids) will be updated by agl values.
func (c *Client) UpdateAccountGeneralLedgers(ids []int64, agl *AccountGeneralLedger) error {
	return c.Update(AccountGeneralLedgerModel, ids, agl, nil)
}

// DeleteAccountGeneralLedger deletes an existing account.general.ledger record.
func (c *Client) DeleteAccountGeneralLedger(id int64) error {
	return c.DeleteAccountGeneralLedgers([]int64{id})
}

// DeleteAccountGeneralLedgers deletes existing account.general.ledger records.
func (c *Client) DeleteAccountGeneralLedgers(ids []int64) error {
	return c.Delete(AccountGeneralLedgerModel, ids)
}

// GetAccountGeneralLedger gets account.general.ledger existing record.
func (c *Client) GetAccountGeneralLedger(id int64) (*AccountGeneralLedger, error) {
	agls, err := c.GetAccountGeneralLedgers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*agls)[0]), nil
}

// GetAccountGeneralLedgers gets account.general.ledger existing records.
func (c *Client) GetAccountGeneralLedgers(ids []int64) (*AccountGeneralLedgers, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.Read(AccountGeneralLedgerModel, ids, nil, agls); err != nil {
		return nil, err
	}
	return agls, nil
}

// FindAccountGeneralLedger finds account.general.ledger record by querying it with criteria.
func (c *Client) FindAccountGeneralLedger(criteria *Criteria) (*AccountGeneralLedger, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.SearchRead(AccountGeneralLedgerModel, criteria, NewOptions().Limit(1), agls); err != nil {
		return nil, err
	}
	return &((*agls)[0]), nil
}

// FindAccountGeneralLedgers finds account.general.ledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgers(criteria *Criteria, options *Options) (*AccountGeneralLedgers, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.SearchRead(AccountGeneralLedgerModel, criteria, options, agls); err != nil {
		return nil, err
	}
	return agls, nil
}

// FindAccountGeneralLedgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGeneralLedgerModel, criteria, options)
}

// FindAccountGeneralLedgerId finds record id by querying it with criteria.
func (c *Client) FindAccountGeneralLedgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGeneralLedgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
