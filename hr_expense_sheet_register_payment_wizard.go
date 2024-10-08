package odoo

// HrExpenseSheetRegisterPaymentWizard represents hr.expense.sheet.register.payment.wizard model.
type HrExpenseSheetRegisterPaymentWizard struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omitempty"`
	Amount                    *Float    `xmlrpc:"amount,omitempty"`
	Communication             *String   `xmlrpc:"communication,omitempty"`
	CompanyId                 *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	ExpenseSheetId            *Many2One `xmlrpc:"expense_sheet_id,omitempty"`
	HidePaymentMethod         *Bool     `xmlrpc:"hide_payment_method,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	JournalId                 *Many2One `xmlrpc:"journal_id,omitempty"`
	PartnerBankAccountId      *Many2One `xmlrpc:"partner_bank_account_id,omitempty"`
	PartnerId                 *Many2One `xmlrpc:"partner_id,omitempty"`
	PaymentDate               *Time     `xmlrpc:"payment_date,omitempty"`
	PaymentMethodId           *Many2One `xmlrpc:"payment_method_id,omitempty"`
	RequirePartnerBankAccount *Bool     `xmlrpc:"require_partner_bank_account,omitempty"`
	ShowPartnerBankAccount    *Bool     `xmlrpc:"show_partner_bank_account,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrExpenseSheetRegisterPaymentWizards represents array of hr.expense.sheet.register.payment.wizard model.
type HrExpenseSheetRegisterPaymentWizards []HrExpenseSheetRegisterPaymentWizard

// HrExpenseSheetRegisterPaymentWizardModel is the odoo model name.
const HrExpenseSheetRegisterPaymentWizardModel = "hr.expense.sheet.register.payment.wizard"

// Many2One convert HrExpenseSheetRegisterPaymentWizard to *Many2One.
func (hesrpw *HrExpenseSheetRegisterPaymentWizard) Many2One() *Many2One {
	return NewMany2One(hesrpw.Id.Get(), "")
}

// CreateHrExpenseSheetRegisterPaymentWizard creates a new hr.expense.sheet.register.payment.wizard model and returns its id.
func (c *Client) CreateHrExpenseSheetRegisterPaymentWizard(hesrpw *HrExpenseSheetRegisterPaymentWizard) (int64, error) {
	ids, err := c.CreateHrExpenseSheetRegisterPaymentWizards([]*HrExpenseSheetRegisterPaymentWizard{hesrpw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSheetRegisterPaymentWizard creates a new hr.expense.sheet.register.payment.wizard model and returns its id.
func (c *Client) CreateHrExpenseSheetRegisterPaymentWizards(hesrpws []*HrExpenseSheetRegisterPaymentWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hesrpws {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSheetRegisterPaymentWizardModel, vv, nil)
}

// UpdateHrExpenseSheetRegisterPaymentWizard updates an existing hr.expense.sheet.register.payment.wizard record.
func (c *Client) UpdateHrExpenseSheetRegisterPaymentWizard(hesrpw *HrExpenseSheetRegisterPaymentWizard) error {
	return c.UpdateHrExpenseSheetRegisterPaymentWizards([]int64{hesrpw.Id.Get()}, hesrpw)
}

// UpdateHrExpenseSheetRegisterPaymentWizards updates existing hr.expense.sheet.register.payment.wizard records.
// All records (represented by ids) will be updated by hesrpw values.
func (c *Client) UpdateHrExpenseSheetRegisterPaymentWizards(ids []int64, hesrpw *HrExpenseSheetRegisterPaymentWizard) error {
	return c.Update(HrExpenseSheetRegisterPaymentWizardModel, ids, hesrpw, nil)
}

// DeleteHrExpenseSheetRegisterPaymentWizard deletes an existing hr.expense.sheet.register.payment.wizard record.
func (c *Client) DeleteHrExpenseSheetRegisterPaymentWizard(id int64) error {
	return c.DeleteHrExpenseSheetRegisterPaymentWizards([]int64{id})
}

// DeleteHrExpenseSheetRegisterPaymentWizards deletes existing hr.expense.sheet.register.payment.wizard records.
func (c *Client) DeleteHrExpenseSheetRegisterPaymentWizards(ids []int64) error {
	return c.Delete(HrExpenseSheetRegisterPaymentWizardModel, ids)
}

// GetHrExpenseSheetRegisterPaymentWizard gets hr.expense.sheet.register.payment.wizard existing record.
func (c *Client) GetHrExpenseSheetRegisterPaymentWizard(id int64) (*HrExpenseSheetRegisterPaymentWizard, error) {
	hesrpws, err := c.GetHrExpenseSheetRegisterPaymentWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hesrpws)[0]), nil
}

// GetHrExpenseSheetRegisterPaymentWizards gets hr.expense.sheet.register.payment.wizard existing records.
func (c *Client) GetHrExpenseSheetRegisterPaymentWizards(ids []int64) (*HrExpenseSheetRegisterPaymentWizards, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.Read(HrExpenseSheetRegisterPaymentWizardModel, ids, nil, hesrpws); err != nil {
		return nil, err
	}
	return hesrpws, nil
}

// FindHrExpenseSheetRegisterPaymentWizard finds hr.expense.sheet.register.payment.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizard(criteria *Criteria) (*HrExpenseSheetRegisterPaymentWizard, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.SearchRead(HrExpenseSheetRegisterPaymentWizardModel, criteria, NewOptions().Limit(1), hesrpws); err != nil {
		return nil, err
	}
	return &((*hesrpws)[0]), nil
}

// FindHrExpenseSheetRegisterPaymentWizards finds hr.expense.sheet.register.payment.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizards(criteria *Criteria, options *Options) (*HrExpenseSheetRegisterPaymentWizards, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.SearchRead(HrExpenseSheetRegisterPaymentWizardModel, criteria, options, hesrpws); err != nil {
		return nil, err
	}
	return hesrpws, nil
}

// FindHrExpenseSheetRegisterPaymentWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseSheetRegisterPaymentWizardModel, criteria, options)
}

// FindHrExpenseSheetRegisterPaymentWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSheetRegisterPaymentWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
