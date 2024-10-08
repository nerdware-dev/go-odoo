package odoo

// AccountAgedReceivable represents account.aged.receivable model.
type AccountAgedReceivable struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAgedReceivables represents array of account.aged.receivable model.
type AccountAgedReceivables []AccountAgedReceivable

// AccountAgedReceivableModel is the odoo model name.
const AccountAgedReceivableModel = "account.aged.receivable"

// Many2One convert AccountAgedReceivable to *Many2One.
func (aar *AccountAgedReceivable) Many2One() *Many2One {
	return NewMany2One(aar.Id.Get(), "")
}

// CreateAccountAgedReceivable creates a new account.aged.receivable model and returns its id.
func (c *Client) CreateAccountAgedReceivable(aar *AccountAgedReceivable) (int64, error) {
	ids, err := c.CreateAccountAgedReceivables([]*AccountAgedReceivable{aar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAgedReceivable creates a new account.aged.receivable model and returns its id.
func (c *Client) CreateAccountAgedReceivables(aars []*AccountAgedReceivable) ([]int64, error) {
	var vv []interface{}
	for _, v := range aars {
		vv = append(vv, v)
	}
	return c.Create(AccountAgedReceivableModel, vv, nil)
}

// UpdateAccountAgedReceivable updates an existing account.aged.receivable record.
func (c *Client) UpdateAccountAgedReceivable(aar *AccountAgedReceivable) error {
	return c.UpdateAccountAgedReceivables([]int64{aar.Id.Get()}, aar)
}

// UpdateAccountAgedReceivables updates existing account.aged.receivable records.
// All records (represented by ids) will be updated by aar values.
func (c *Client) UpdateAccountAgedReceivables(ids []int64, aar *AccountAgedReceivable) error {
	return c.Update(AccountAgedReceivableModel, ids, aar, nil)
}

// DeleteAccountAgedReceivable deletes an existing account.aged.receivable record.
func (c *Client) DeleteAccountAgedReceivable(id int64) error {
	return c.DeleteAccountAgedReceivables([]int64{id})
}

// DeleteAccountAgedReceivables deletes existing account.aged.receivable records.
func (c *Client) DeleteAccountAgedReceivables(ids []int64) error {
	return c.Delete(AccountAgedReceivableModel, ids)
}

// GetAccountAgedReceivable gets account.aged.receivable existing record.
func (c *Client) GetAccountAgedReceivable(id int64) (*AccountAgedReceivable, error) {
	aars, err := c.GetAccountAgedReceivables([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aars)[0]), nil
}

// GetAccountAgedReceivables gets account.aged.receivable existing records.
func (c *Client) GetAccountAgedReceivables(ids []int64) (*AccountAgedReceivables, error) {
	aars := &AccountAgedReceivables{}
	if err := c.Read(AccountAgedReceivableModel, ids, nil, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAgedReceivable finds account.aged.receivable record by querying it with criteria.
func (c *Client) FindAccountAgedReceivable(criteria *Criteria) (*AccountAgedReceivable, error) {
	aars := &AccountAgedReceivables{}
	if err := c.SearchRead(AccountAgedReceivableModel, criteria, NewOptions().Limit(1), aars); err != nil {
		return nil, err
	}
	return &((*aars)[0]), nil
}

// FindAccountAgedReceivables finds account.aged.receivable records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedReceivables(criteria *Criteria, options *Options) (*AccountAgedReceivables, error) {
	aars := &AccountAgedReceivables{}
	if err := c.SearchRead(AccountAgedReceivableModel, criteria, options, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAgedReceivableIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedReceivableIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAgedReceivableModel, criteria, options)
}

// FindAccountAgedReceivableId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedReceivableId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedReceivableModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
