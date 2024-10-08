package odoo

// AccountConsolidatedJournal represents account.consolidated.journal model.
type AccountConsolidatedJournal struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountConsolidatedJournals represents array of account.consolidated.journal model.
type AccountConsolidatedJournals []AccountConsolidatedJournal

// AccountConsolidatedJournalModel is the odoo model name.
const AccountConsolidatedJournalModel = "account.consolidated.journal"

// Many2One convert AccountConsolidatedJournal to *Many2One.
func (acj *AccountConsolidatedJournal) Many2One() *Many2One {
	return NewMany2One(acj.Id.Get(), "")
}

// CreateAccountConsolidatedJournal creates a new account.consolidated.journal model and returns its id.
func (c *Client) CreateAccountConsolidatedJournal(acj *AccountConsolidatedJournal) (int64, error) {
	ids, err := c.CreateAccountConsolidatedJournals([]*AccountConsolidatedJournal{acj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountConsolidatedJournal creates a new account.consolidated.journal model and returns its id.
func (c *Client) CreateAccountConsolidatedJournals(acjs []*AccountConsolidatedJournal) ([]int64, error) {
	var vv []interface{}
	for _, v := range acjs {
		vv = append(vv, v)
	}
	return c.Create(AccountConsolidatedJournalModel, vv, nil)
}

// UpdateAccountConsolidatedJournal updates an existing account.consolidated.journal record.
func (c *Client) UpdateAccountConsolidatedJournal(acj *AccountConsolidatedJournal) error {
	return c.UpdateAccountConsolidatedJournals([]int64{acj.Id.Get()}, acj)
}

// UpdateAccountConsolidatedJournals updates existing account.consolidated.journal records.
// All records (represented by ids) will be updated by acj values.
func (c *Client) UpdateAccountConsolidatedJournals(ids []int64, acj *AccountConsolidatedJournal) error {
	return c.Update(AccountConsolidatedJournalModel, ids, acj, nil)
}

// DeleteAccountConsolidatedJournal deletes an existing account.consolidated.journal record.
func (c *Client) DeleteAccountConsolidatedJournal(id int64) error {
	return c.DeleteAccountConsolidatedJournals([]int64{id})
}

// DeleteAccountConsolidatedJournals deletes existing account.consolidated.journal records.
func (c *Client) DeleteAccountConsolidatedJournals(ids []int64) error {
	return c.Delete(AccountConsolidatedJournalModel, ids)
}

// GetAccountConsolidatedJournal gets account.consolidated.journal existing record.
func (c *Client) GetAccountConsolidatedJournal(id int64) (*AccountConsolidatedJournal, error) {
	acjs, err := c.GetAccountConsolidatedJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acjs)[0]), nil
}

// GetAccountConsolidatedJournals gets account.consolidated.journal existing records.
func (c *Client) GetAccountConsolidatedJournals(ids []int64) (*AccountConsolidatedJournals, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.Read(AccountConsolidatedJournalModel, ids, nil, acjs); err != nil {
		return nil, err
	}
	return acjs, nil
}

// FindAccountConsolidatedJournal finds account.consolidated.journal record by querying it with criteria.
func (c *Client) FindAccountConsolidatedJournal(criteria *Criteria) (*AccountConsolidatedJournal, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.SearchRead(AccountConsolidatedJournalModel, criteria, NewOptions().Limit(1), acjs); err != nil {
		return nil, err
	}
	return &((*acjs)[0]), nil
}

// FindAccountConsolidatedJournals finds account.consolidated.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountConsolidatedJournals(criteria *Criteria, options *Options) (*AccountConsolidatedJournals, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.SearchRead(AccountConsolidatedJournalModel, criteria, options, acjs); err != nil {
		return nil, err
	}
	return acjs, nil
}

// FindAccountConsolidatedJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountConsolidatedJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountConsolidatedJournalModel, criteria, options)
}

// FindAccountConsolidatedJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountConsolidatedJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountConsolidatedJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
