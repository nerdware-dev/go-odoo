package odoo

// AccountPaymentRegister represents account.payment.register model.
type AccountPaymentRegister struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	GroupPayment    *Bool     `xmlrpc:"group_payment,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	InvoiceIds      *Relation `xmlrpc:"invoice_ids,omitempty"`
	JournalId       *Many2One `xmlrpc:"journal_id,omitempty"`
	PaymentDate     *Time     `xmlrpc:"payment_date,omitempty"`
	PaymentMethodId *Many2One `xmlrpc:"payment_method_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountPaymentRegisters represents array of account.payment.register model.
type AccountPaymentRegisters []AccountPaymentRegister

// AccountPaymentRegisterModel is the odoo model name.
const AccountPaymentRegisterModel = "account.payment.register"

// Many2One convert AccountPaymentRegister to *Many2One.
func (apr *AccountPaymentRegister) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegister(apr *AccountPaymentRegister) (int64, error) {
	ids, err := c.CreateAccountPaymentRegisters([]*AccountPaymentRegister{apr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegisters(aprs []*AccountPaymentRegister) ([]int64, error) {
	var vv []interface{}
	for _, v := range aprs {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentRegisterModel, vv, nil)
}

// UpdateAccountPaymentRegister updates an existing account.payment.register record.
func (c *Client) UpdateAccountPaymentRegister(apr *AccountPaymentRegister) error {
	return c.UpdateAccountPaymentRegisters([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPaymentRegisters updates existing account.payment.register records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPaymentRegisters(ids []int64, apr *AccountPaymentRegister) error {
	return c.Update(AccountPaymentRegisterModel, ids, apr, nil)
}

// DeleteAccountPaymentRegister deletes an existing account.payment.register record.
func (c *Client) DeleteAccountPaymentRegister(id int64) error {
	return c.DeleteAccountPaymentRegisters([]int64{id})
}

// DeleteAccountPaymentRegisters deletes existing account.payment.register records.
func (c *Client) DeleteAccountPaymentRegisters(ids []int64) error {
	return c.Delete(AccountPaymentRegisterModel, ids)
}

// GetAccountPaymentRegister gets account.payment.register existing record.
func (c *Client) GetAccountPaymentRegister(id int64) (*AccountPaymentRegister, error) {
	aprs, err := c.GetAccountPaymentRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// GetAccountPaymentRegisters gets account.payment.register existing records.
func (c *Client) GetAccountPaymentRegisters(ids []int64) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.Read(AccountPaymentRegisterModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegister finds account.payment.register record by querying it with criteria.
func (c *Client) FindAccountPaymentRegister(criteria *Criteria) (*AccountPaymentRegister, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// FindAccountPaymentRegisters finds account.payment.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisters(criteria *Criteria, options *Options) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentRegisterModel, criteria, options)
}

// FindAccountPaymentRegisterId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
