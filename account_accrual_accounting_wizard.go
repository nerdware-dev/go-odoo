package odoo

// AccountAccrualAccountingWizard represents account.accrual.accounting.wizard model.
type AccountAccrualAccountingWizard struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	AccountType           *Selection `xmlrpc:"account_type,omitempty"`
	ActiveMoveLineIds     *Relation  `xmlrpc:"active_move_line_ids,omitempty"`
	CompanyCurrencyId     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                  *Time      `xmlrpc:"date,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	ExpenseAccrualAccount *Many2One  `xmlrpc:"expense_accrual_account,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	JournalId             *Many2One  `xmlrpc:"journal_id,omitempty"`
	Percentage            *Float     `xmlrpc:"percentage,omitempty"`
	RevenueAccrualAccount *Many2One  `xmlrpc:"revenue_accrual_account,omitempty"`
	TotalAmount           *Float     `xmlrpc:"total_amount,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAccrualAccountingWizards represents array of account.accrual.accounting.wizard model.
type AccountAccrualAccountingWizards []AccountAccrualAccountingWizard

// AccountAccrualAccountingWizardModel is the odoo model name.
const AccountAccrualAccountingWizardModel = "account.accrual.accounting.wizard"

// Many2One convert AccountAccrualAccountingWizard to *Many2One.
func (aaaw *AccountAccrualAccountingWizard) Many2One() *Many2One {
	return NewMany2One(aaaw.Id.Get(), "")
}

// CreateAccountAccrualAccountingWizard creates a new account.accrual.accounting.wizard model and returns its id.
func (c *Client) CreateAccountAccrualAccountingWizard(aaaw *AccountAccrualAccountingWizard) (int64, error) {
	ids, err := c.CreateAccountAccrualAccountingWizards([]*AccountAccrualAccountingWizard{aaaw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAccrualAccountingWizard creates a new account.accrual.accounting.wizard model and returns its id.
func (c *Client) CreateAccountAccrualAccountingWizards(aaaws []*AccountAccrualAccountingWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaaws {
		vv = append(vv, v)
	}
	return c.Create(AccountAccrualAccountingWizardModel, vv, nil)
}

// UpdateAccountAccrualAccountingWizard updates an existing account.accrual.accounting.wizard record.
func (c *Client) UpdateAccountAccrualAccountingWizard(aaaw *AccountAccrualAccountingWizard) error {
	return c.UpdateAccountAccrualAccountingWizards([]int64{aaaw.Id.Get()}, aaaw)
}

// UpdateAccountAccrualAccountingWizards updates existing account.accrual.accounting.wizard records.
// All records (represented by ids) will be updated by aaaw values.
func (c *Client) UpdateAccountAccrualAccountingWizards(ids []int64, aaaw *AccountAccrualAccountingWizard) error {
	return c.Update(AccountAccrualAccountingWizardModel, ids, aaaw, nil)
}

// DeleteAccountAccrualAccountingWizard deletes an existing account.accrual.accounting.wizard record.
func (c *Client) DeleteAccountAccrualAccountingWizard(id int64) error {
	return c.DeleteAccountAccrualAccountingWizards([]int64{id})
}

// DeleteAccountAccrualAccountingWizards deletes existing account.accrual.accounting.wizard records.
func (c *Client) DeleteAccountAccrualAccountingWizards(ids []int64) error {
	return c.Delete(AccountAccrualAccountingWizardModel, ids)
}

// GetAccountAccrualAccountingWizard gets account.accrual.accounting.wizard existing record.
func (c *Client) GetAccountAccrualAccountingWizard(id int64) (*AccountAccrualAccountingWizard, error) {
	aaaws, err := c.GetAccountAccrualAccountingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaaws)[0]), nil
}

// GetAccountAccrualAccountingWizards gets account.accrual.accounting.wizard existing records.
func (c *Client) GetAccountAccrualAccountingWizards(ids []int64) (*AccountAccrualAccountingWizards, error) {
	aaaws := &AccountAccrualAccountingWizards{}
	if err := c.Read(AccountAccrualAccountingWizardModel, ids, nil, aaaws); err != nil {
		return nil, err
	}
	return aaaws, nil
}

// FindAccountAccrualAccountingWizard finds account.accrual.accounting.wizard record by querying it with criteria.
func (c *Client) FindAccountAccrualAccountingWizard(criteria *Criteria) (*AccountAccrualAccountingWizard, error) {
	aaaws := &AccountAccrualAccountingWizards{}
	if err := c.SearchRead(AccountAccrualAccountingWizardModel, criteria, NewOptions().Limit(1), aaaws); err != nil {
		return nil, err
	}
	return &((*aaaws)[0]), nil
}

// FindAccountAccrualAccountingWizards finds account.accrual.accounting.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccrualAccountingWizards(criteria *Criteria, options *Options) (*AccountAccrualAccountingWizards, error) {
	aaaws := &AccountAccrualAccountingWizards{}
	if err := c.SearchRead(AccountAccrualAccountingWizardModel, criteria, options, aaaws); err != nil {
		return nil, err
	}
	return aaaws, nil
}

// FindAccountAccrualAccountingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccrualAccountingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAccrualAccountingWizardModel, criteria, options)
}

// FindAccountAccrualAccountingWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAccrualAccountingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccrualAccountingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
