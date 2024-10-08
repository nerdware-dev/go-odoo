package odoo

// AccountOnlineJournal represents account.online.journal model.
type AccountOnlineJournal struct {
	LastUpdate                         *Time     `xmlrpc:"__last_update,omitempty"`
	AccountNumber                      *String   `xmlrpc:"account_number,omitempty"`
	AccountOnlineProviderId            *Many2One `xmlrpc:"account_online_provider_id,omitempty"`
	Balance                            *Float    `xmlrpc:"balance,omitempty"`
	CreateDate                         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                        *String   `xmlrpc:"display_name,omitempty"`
	Id                                 *Int      `xmlrpc:"id,omitempty"`
	JournalIds                         *Relation `xmlrpc:"journal_ids,omitempty"`
	LastSync                           *Time     `xmlrpc:"last_sync,omitempty"`
	Name                               *String   `xmlrpc:"name,omitempty"`
	OnlineIdentifier                   *String   `xmlrpc:"online_identifier,omitempty"`
	PontoLastSynchronizationIdentifier *String   `xmlrpc:"ponto_last_synchronization_identifier,omitempty"`
	ProviderName                       *String   `xmlrpc:"provider_name,omitempty"`
	WriteDate                          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                           *Many2One `xmlrpc:"write_uid,omitempty"`
	YodleeAccountStatus                *String   `xmlrpc:"yodlee_account_status,omitempty"`
	YodleeStatusCode                   *Int      `xmlrpc:"yodlee_status_code,omitempty"`
}

// AccountOnlineJournals represents array of account.online.journal model.
type AccountOnlineJournals []AccountOnlineJournal

// AccountOnlineJournalModel is the odoo model name.
const AccountOnlineJournalModel = "account.online.journal"

// Many2One convert AccountOnlineJournal to *Many2One.
func (aoj *AccountOnlineJournal) Many2One() *Many2One {
	return NewMany2One(aoj.Id.Get(), "")
}

// CreateAccountOnlineJournal creates a new account.online.journal model and returns its id.
func (c *Client) CreateAccountOnlineJournal(aoj *AccountOnlineJournal) (int64, error) {
	ids, err := c.CreateAccountOnlineJournals([]*AccountOnlineJournal{aoj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineJournal creates a new account.online.journal model and returns its id.
func (c *Client) CreateAccountOnlineJournals(aojs []*AccountOnlineJournal) ([]int64, error) {
	var vv []interface{}
	for _, v := range aojs {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineJournalModel, vv, nil)
}

// UpdateAccountOnlineJournal updates an existing account.online.journal record.
func (c *Client) UpdateAccountOnlineJournal(aoj *AccountOnlineJournal) error {
	return c.UpdateAccountOnlineJournals([]int64{aoj.Id.Get()}, aoj)
}

// UpdateAccountOnlineJournals updates existing account.online.journal records.
// All records (represented by ids) will be updated by aoj values.
func (c *Client) UpdateAccountOnlineJournals(ids []int64, aoj *AccountOnlineJournal) error {
	return c.Update(AccountOnlineJournalModel, ids, aoj, nil)
}

// DeleteAccountOnlineJournal deletes an existing account.online.journal record.
func (c *Client) DeleteAccountOnlineJournal(id int64) error {
	return c.DeleteAccountOnlineJournals([]int64{id})
}

// DeleteAccountOnlineJournals deletes existing account.online.journal records.
func (c *Client) DeleteAccountOnlineJournals(ids []int64) error {
	return c.Delete(AccountOnlineJournalModel, ids)
}

// GetAccountOnlineJournal gets account.online.journal existing record.
func (c *Client) GetAccountOnlineJournal(id int64) (*AccountOnlineJournal, error) {
	aojs, err := c.GetAccountOnlineJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aojs)[0]), nil
}

// GetAccountOnlineJournals gets account.online.journal existing records.
func (c *Client) GetAccountOnlineJournals(ids []int64) (*AccountOnlineJournals, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.Read(AccountOnlineJournalModel, ids, nil, aojs); err != nil {
		return nil, err
	}
	return aojs, nil
}

// FindAccountOnlineJournal finds account.online.journal record by querying it with criteria.
func (c *Client) FindAccountOnlineJournal(criteria *Criteria) (*AccountOnlineJournal, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.SearchRead(AccountOnlineJournalModel, criteria, NewOptions().Limit(1), aojs); err != nil {
		return nil, err
	}
	return &((*aojs)[0]), nil
}

// FindAccountOnlineJournals finds account.online.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineJournals(criteria *Criteria, options *Options) (*AccountOnlineJournals, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.SearchRead(AccountOnlineJournalModel, criteria, options, aojs); err != nil {
		return nil, err
	}
	return aojs, nil
}

// FindAccountOnlineJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOnlineJournalModel, criteria, options)
}

// FindAccountOnlineJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
