package odoo

// AccountPayment represents account.payment model.
type AccountPayment struct {
	ActivityDateDeadline            *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration     *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon           *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                     *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                   *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary                 *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon                *String    `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                  *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                  *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	Amount                          *Float     `xmlrpc:"amount,omitempty" json:"amount,omitempty"`
	AmountAvailableForRefund        *Float     `xmlrpc:"amount_available_for_refund,omitempty" json:"amount_available_for_refund,omitempty"`
	AmountCompanyCurrencySigned     *Float     `xmlrpc:"amount_company_currency_signed,omitempty" json:"amount_company_currency_signed,omitempty"`
	AmountSigned                    *Float     `xmlrpc:"amount_signed,omitempty" json:"amount_signed,omitempty"`
	AttachmentIds                   *Relation  `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	AvailableJournalIds             *Relation  `xmlrpc:"available_journal_ids,omitempty" json:"available_journal_ids,omitempty"`
	AvailablePartnerBankIds         *Relation  `xmlrpc:"available_partner_bank_ids,omitempty" json:"available_partner_bank_ids,omitempty"`
	AvailablePaymentMethodLineIds   *Relation  `xmlrpc:"available_payment_method_line_ids,omitempty" json:"available_payment_method_line_ids,omitempty"`
	CompanyCurrencyId               *Many2One  `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyId                       *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryCode                     *String    `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CreateDate                      *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                       *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                      *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                            *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DestinationAccountId            *Many2One  `xmlrpc:"destination_account_id,omitempty" json:"destination_account_id,omitempty"`
	DisplayName                     *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DuplicatePaymentIds             *Relation  `xmlrpc:"duplicate_payment_ids,omitempty" json:"duplicate_payment_ids,omitempty"`
	ExpenseSheetId                  *Many2One  `xmlrpc:"expense_sheet_id,omitempty" json:"expense_sheet_id,omitempty"`
	HasMessage                      *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                              *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InvoiceIds                      *Relation  `xmlrpc:"invoice_ids,omitempty" json:"invoice_ids,omitempty"`
	IsMatched                       *Bool      `xmlrpc:"is_matched,omitempty" json:"is_matched,omitempty"`
	IsReconciled                    *Bool      `xmlrpc:"is_reconciled,omitempty" json:"is_reconciled,omitempty"`
	IsSent                          *Bool      `xmlrpc:"is_sent,omitempty" json:"is_sent,omitempty"`
	JournalId                       *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	Memo                            *String    `xmlrpc:"memo,omitempty" json:"memo,omitempty"`
	MessageAttachmentCount          *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds              *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError                 *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter          *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError              *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                      *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower               *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentId         *Many2One  `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction               *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter        *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds               *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MoveId                          *Many2One  `xmlrpc:"move_id,omitempty" json:"move_id,omitempty"`
	MyActivityDateDeadline          *Time      `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                            *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NeedCancelRequest               *Bool      `xmlrpc:"need_cancel_request,omitempty" json:"need_cancel_request,omitempty"`
	OutstandingAccountId            *Many2One  `xmlrpc:"outstanding_account_id,omitempty" json:"outstanding_account_id,omitempty"`
	PairedInternalTransferPaymentId *Many2One  `xmlrpc:"paired_internal_transfer_payment_id,omitempty" json:"paired_internal_transfer_payment_id,omitempty"`
	PartnerBankId                   *Many2One  `xmlrpc:"partner_bank_id,omitempty" json:"partner_bank_id,omitempty"`
	PartnerId                       *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerType                     *Selection `xmlrpc:"partner_type,omitempty" json:"partner_type,omitempty"`
	PaymentMethodCode               *String    `xmlrpc:"payment_method_code,omitempty" json:"payment_method_code,omitempty"`
	PaymentMethodId                 *Many2One  `xmlrpc:"payment_method_id,omitempty" json:"payment_method_id,omitempty"`
	PaymentMethodLineId             *Many2One  `xmlrpc:"payment_method_line_id,omitempty" json:"payment_method_line_id,omitempty"`
	PaymentReceiptTitle             *String    `xmlrpc:"payment_receipt_title,omitempty" json:"payment_receipt_title,omitempty"`
	PaymentReference                *String    `xmlrpc:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	PaymentTokenId                  *Many2One  `xmlrpc:"payment_token_id,omitempty" json:"payment_token_id,omitempty"`
	PaymentTransactionId            *Many2One  `xmlrpc:"payment_transaction_id,omitempty" json:"payment_transaction_id,omitempty"`
	PaymentType                     *Selection `xmlrpc:"payment_type,omitempty" json:"payment_type,omitempty"`
	QrCode                          *String    `xmlrpc:"qr_code,omitempty" json:"qr_code,omitempty"`
	RatingIds                       *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	ReconciledBillIds               *Relation  `xmlrpc:"reconciled_bill_ids,omitempty" json:"reconciled_bill_ids,omitempty"`
	ReconciledBillsCount            *Int       `xmlrpc:"reconciled_bills_count,omitempty" json:"reconciled_bills_count,omitempty"`
	ReconciledInvoiceIds            *Relation  `xmlrpc:"reconciled_invoice_ids,omitempty" json:"reconciled_invoice_ids,omitempty"`
	ReconciledInvoicesCount         *Int       `xmlrpc:"reconciled_invoices_count,omitempty" json:"reconciled_invoices_count,omitempty"`
	ReconciledInvoicesType          *Selection `xmlrpc:"reconciled_invoices_type,omitempty" json:"reconciled_invoices_type,omitempty"`
	ReconciledStatementLineIds      *Relation  `xmlrpc:"reconciled_statement_line_ids,omitempty" json:"reconciled_statement_line_ids,omitempty"`
	ReconciledStatementLinesCount   *Int       `xmlrpc:"reconciled_statement_lines_count,omitempty" json:"reconciled_statement_lines_count,omitempty"`
	RefundsCount                    *Int       `xmlrpc:"refunds_count,omitempty" json:"refunds_count,omitempty"`
	RequirePartnerBankAccount       *Bool      `xmlrpc:"require_partner_bank_account,omitempty" json:"require_partner_bank_account,omitempty"`
	ShowPartnerBankAccount          *Bool      `xmlrpc:"show_partner_bank_account,omitempty" json:"show_partner_bank_account,omitempty"`
	SourcePaymentId                 *Many2One  `xmlrpc:"source_payment_id,omitempty" json:"source_payment_id,omitempty"`
	State                           *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	SuitablePaymentTokenIds         *Relation  `xmlrpc:"suitable_payment_token_ids,omitempty" json:"suitable_payment_token_ids,omitempty"`
	UseElectronicPaymentMethod      *Bool      `xmlrpc:"use_electronic_payment_method,omitempty" json:"use_electronic_payment_method,omitempty"`
	WebsiteMessageIds               *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// AccountPayments represents array of account.payment model.
type AccountPayments []AccountPayment

// AccountPaymentModel is the odoo model name.
const AccountPaymentModel = "account.payment"

// Many2One convert AccountPayment to *Many2One.
func (ap *AccountPayment) Many2One() *Many2One {
	return NewMany2One(ap.Id.Get(), "")
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayment(ap *AccountPayment) (int64, error) {
	ids, err := c.CreateAccountPayments([]*AccountPayment{ap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayments(aps []*AccountPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range aps {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentModel, vv, nil)
}

// UpdateAccountPayment updates an existing account.payment record.
func (c *Client) UpdateAccountPayment(ap *AccountPayment) error {
	return c.UpdateAccountPayments([]int64{ap.Id.Get()}, ap)
}

// UpdateAccountPayments updates existing account.payment records.
// All records (represented by ids) will be updated by ap values.
func (c *Client) UpdateAccountPayments(ids []int64, ap *AccountPayment) error {
	return c.Update(AccountPaymentModel, ids, ap, nil)
}

// DeleteAccountPayment deletes an existing account.payment record.
func (c *Client) DeleteAccountPayment(id int64) error {
	return c.DeleteAccountPayments([]int64{id})
}

// DeleteAccountPayments deletes existing account.payment records.
func (c *Client) DeleteAccountPayments(ids []int64) error {
	return c.Delete(AccountPaymentModel, ids)
}

// GetAccountPayment gets account.payment existing record.
func (c *Client) GetAccountPayment(id int64) (*AccountPayment, error) {
	aps, err := c.GetAccountPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aps)[0]), nil
}

// GetAccountPayments gets account.payment existing records.
func (c *Client) GetAccountPayments(ids []int64) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.Read(AccountPaymentModel, ids, nil, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPayment finds account.payment record by querying it with criteria.
func (c *Client) FindAccountPayment(criteria *Criteria) (*AccountPayment, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, NewOptions().Limit(1), aps); err != nil {
		return nil, err
	}
	return &((*aps)[0]), nil
}

// FindAccountPayments finds account.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPayments(criteria *Criteria, options *Options) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, options, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentModel, criteria, options)
}

// FindAccountPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
