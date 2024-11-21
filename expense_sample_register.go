package odoo

// ExpenseSampleRegister represents expense.sample.register model.
type ExpenseSampleRegister struct {
	Amount                        *Float     `xmlrpc:"amount,omitempty" json:"amount,omitempty"`
	AvailablePaymentMethodLineIds *Relation  `xmlrpc:"available_payment_method_line_ids,omitempty" json:"available_payment_method_line_ids,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                          *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	HidePartial                   *Bool      `xmlrpc:"hide_partial,omitempty" json:"hide_partial,omitempty"`
	HidePaymentMethodLine         *Bool      `xmlrpc:"hide_payment_method_line,omitempty" json:"hide_payment_method_line,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	JournalId                     *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	Memo                          *String    `xmlrpc:"memo,omitempty" json:"memo,omitempty"`
	PartialMode                   *Selection `xmlrpc:"partial_mode,omitempty" json:"partial_mode,omitempty"`
	PaymentMethodLineId           *Many2One  `xmlrpc:"payment_method_line_id,omitempty" json:"payment_method_line_id,omitempty"`
	SheetId                       *Many2One  `xmlrpc:"sheet_id,omitempty" json:"sheet_id,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// ExpenseSampleRegisters represents array of expense.sample.register model.
type ExpenseSampleRegisters []ExpenseSampleRegister

// ExpenseSampleRegisterModel is the odoo model name.
const ExpenseSampleRegisterModel = "expense.sample.register"

// Many2One convert ExpenseSampleRegister to *Many2One.
func (esr *ExpenseSampleRegister) Many2One() *Many2One {
	return NewMany2One(esr.Id.Get(), "")
}

// CreateExpenseSampleRegister creates a new expense.sample.register model and returns its id.
func (c *Client) CreateExpenseSampleRegister(esr *ExpenseSampleRegister) (int64, error) {
	ids, err := c.CreateExpenseSampleRegisters([]*ExpenseSampleRegister{esr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateExpenseSampleRegister creates a new expense.sample.register model and returns its id.
func (c *Client) CreateExpenseSampleRegisters(esrs []*ExpenseSampleRegister) ([]int64, error) {
	var vv []interface{}
	for _, v := range esrs {
		vv = append(vv, v)
	}
	return c.Create(ExpenseSampleRegisterModel, vv, nil)
}

// UpdateExpenseSampleRegister updates an existing expense.sample.register record.
func (c *Client) UpdateExpenseSampleRegister(esr *ExpenseSampleRegister) error {
	return c.UpdateExpenseSampleRegisters([]int64{esr.Id.Get()}, esr)
}

// UpdateExpenseSampleRegisters updates existing expense.sample.register records.
// All records (represented by ids) will be updated by esr values.
func (c *Client) UpdateExpenseSampleRegisters(ids []int64, esr *ExpenseSampleRegister) error {
	return c.Update(ExpenseSampleRegisterModel, ids, esr, nil)
}

// DeleteExpenseSampleRegister deletes an existing expense.sample.register record.
func (c *Client) DeleteExpenseSampleRegister(id int64) error {
	return c.DeleteExpenseSampleRegisters([]int64{id})
}

// DeleteExpenseSampleRegisters deletes existing expense.sample.register records.
func (c *Client) DeleteExpenseSampleRegisters(ids []int64) error {
	return c.Delete(ExpenseSampleRegisterModel, ids)
}

// GetExpenseSampleRegister gets expense.sample.register existing record.
func (c *Client) GetExpenseSampleRegister(id int64) (*ExpenseSampleRegister, error) {
	esrs, err := c.GetExpenseSampleRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*esrs)[0]), nil
}

// GetExpenseSampleRegisters gets expense.sample.register existing records.
func (c *Client) GetExpenseSampleRegisters(ids []int64) (*ExpenseSampleRegisters, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.Read(ExpenseSampleRegisterModel, ids, nil, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleRegister finds expense.sample.register record by querying it with criteria.
func (c *Client) FindExpenseSampleRegister(criteria *Criteria) (*ExpenseSampleRegister, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.SearchRead(ExpenseSampleRegisterModel, criteria, NewOptions().Limit(1), esrs); err != nil {
		return nil, err
	}
	return &((*esrs)[0]), nil
}

// FindExpenseSampleRegisters finds expense.sample.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleRegisters(criteria *Criteria, options *Options) (*ExpenseSampleRegisters, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.SearchRead(ExpenseSampleRegisterModel, criteria, options, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ExpenseSampleRegisterModel, criteria, options)
}

// FindExpenseSampleRegisterId finds record id by querying it with criteria.
func (c *Client) FindExpenseSampleRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ExpenseSampleRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
