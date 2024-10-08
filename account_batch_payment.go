package odoo

// AccountBatchPayment represents account.batch.payment model.
type AccountBatchPayment struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	Amount                    *Float     `xmlrpc:"amount,omitempty"`
	AvailablePaymentMethodIds *Relation  `xmlrpc:"available_payment_method_ids,omitempty"`
	BatchType                 *Selection `xmlrpc:"batch_type,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                      *Time      `xmlrpc:"date,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	ExportFile                *String    `xmlrpc:"export_file,omitempty"`
	ExportFileCreateDate      *Time      `xmlrpc:"export_file_create_date,omitempty"`
	ExportFilename            *String    `xmlrpc:"export_filename,omitempty"`
	FileGenerationEnabled     *Bool      `xmlrpc:"file_generation_enabled,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PaymentIds                *Relation  `xmlrpc:"payment_ids,omitempty"`
	PaymentMethodCode         *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId           *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	SctGeneric                *Bool      `xmlrpc:"sct_generic,omitempty"`
	SctWarning                *String    `xmlrpc:"sct_warning,omitempty"`
	SddBatchBooking           *Bool      `xmlrpc:"sdd_batch_booking,omitempty"`
	SddRequiredCollectionDate *Time      `xmlrpc:"sdd_required_collection_date,omitempty"`
	State                     *Selection `xmlrpc:"state,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountBatchPayments represents array of account.batch.payment model.
type AccountBatchPayments []AccountBatchPayment

// AccountBatchPaymentModel is the odoo model name.
const AccountBatchPaymentModel = "account.batch.payment"

// Many2One convert AccountBatchPayment to *Many2One.
func (abp *AccountBatchPayment) Many2One() *Many2One {
	return NewMany2One(abp.Id.Get(), "")
}

// CreateAccountBatchPayment creates a new account.batch.payment model and returns its id.
func (c *Client) CreateAccountBatchPayment(abp *AccountBatchPayment) (int64, error) {
	ids, err := c.CreateAccountBatchPayments([]*AccountBatchPayment{abp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchPayment creates a new account.batch.payment model and returns its id.
func (c *Client) CreateAccountBatchPayments(abps []*AccountBatchPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range abps {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchPaymentModel, vv, nil)
}

// UpdateAccountBatchPayment updates an existing account.batch.payment record.
func (c *Client) UpdateAccountBatchPayment(abp *AccountBatchPayment) error {
	return c.UpdateAccountBatchPayments([]int64{abp.Id.Get()}, abp)
}

// UpdateAccountBatchPayments updates existing account.batch.payment records.
// All records (represented by ids) will be updated by abp values.
func (c *Client) UpdateAccountBatchPayments(ids []int64, abp *AccountBatchPayment) error {
	return c.Update(AccountBatchPaymentModel, ids, abp, nil)
}

// DeleteAccountBatchPayment deletes an existing account.batch.payment record.
func (c *Client) DeleteAccountBatchPayment(id int64) error {
	return c.DeleteAccountBatchPayments([]int64{id})
}

// DeleteAccountBatchPayments deletes existing account.batch.payment records.
func (c *Client) DeleteAccountBatchPayments(ids []int64) error {
	return c.Delete(AccountBatchPaymentModel, ids)
}

// GetAccountBatchPayment gets account.batch.payment existing record.
func (c *Client) GetAccountBatchPayment(id int64) (*AccountBatchPayment, error) {
	abps, err := c.GetAccountBatchPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abps)[0]), nil
}

// GetAccountBatchPayments gets account.batch.payment existing records.
func (c *Client) GetAccountBatchPayments(ids []int64) (*AccountBatchPayments, error) {
	abps := &AccountBatchPayments{}
	if err := c.Read(AccountBatchPaymentModel, ids, nil, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBatchPayment finds account.batch.payment record by querying it with criteria.
func (c *Client) FindAccountBatchPayment(criteria *Criteria) (*AccountBatchPayment, error) {
	abps := &AccountBatchPayments{}
	if err := c.SearchRead(AccountBatchPaymentModel, criteria, NewOptions().Limit(1), abps); err != nil {
		return nil, err
	}
	return &((*abps)[0]), nil
}

// FindAccountBatchPayments finds account.batch.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPayments(criteria *Criteria, options *Options) (*AccountBatchPayments, error) {
	abps := &AccountBatchPayments{}
	if err := c.SearchRead(AccountBatchPaymentModel, criteria, options, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBatchPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBatchPaymentModel, criteria, options)
}

// FindAccountBatchPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
