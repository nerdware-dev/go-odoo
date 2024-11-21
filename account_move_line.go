package odoo

// AccountMoveLine represents account.move.line model.
type AccountMoveLine struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	AccountId               *Many2One  `xmlrpc:"account_id,omitempty" json:"account_id,omitempty"`
	AccountInternalType     *Selection `xmlrpc:"account_internal_type,omitempty" json:"account_internal_type,omitempty"`
	AccountRootId           *Many2One  `xmlrpc:"account_root_id,omitempty" json:"account_root_id,omitempty"`
	AlwaysSetCurrencyId     *Many2One  `xmlrpc:"always_set_currency_id,omitempty" json:"always_set_currency_id,omitempty"`
	AmountCurrency          *Float     `xmlrpc:"amount_currency,omitempty" json:"amount_currency,omitempty"`
	AmountResidual          *Float     `xmlrpc:"amount_residual,omitempty" json:"amount_residual,omitempty"`
	AmountResidualCurrency  *Float     `xmlrpc:"amount_residual_currency,omitempty" json:"amount_residual_currency,omitempty"`
	AnalyticAccountId       *Many2One  `xmlrpc:"analytic_account_id,omitempty" json:"analytic_account_id,omitempty"`
	AnalyticLineIds         *Relation  `xmlrpc:"analytic_line_ids,omitempty" json:"analytic_line_ids,omitempty"`
	AnalyticTagIds          *Relation  `xmlrpc:"analytic_tag_ids,omitempty" json:"analytic_tag_ids,omitempty"`
	AssetId                 *Many2One  `xmlrpc:"asset_id,omitempty" json:"asset_id,omitempty"`
	Balance                 *Float     `xmlrpc:"balance,omitempty" json:"balance,omitempty"`
	Blocked                 *Bool      `xmlrpc:"blocked,omitempty" json:"blocked,omitempty"`
	CanBePaid               *Selection `xmlrpc:"can_be_paid,omitempty" json:"can_be_paid,omitempty"`
	CompanyCurrencyId       *Many2One  `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyId               *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryId               *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Credit                  *Float     `xmlrpc:"credit,omitempty" json:"credit,omitempty"`
	CurrencyId              *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                    *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DateMaturity            *Time      `xmlrpc:"date_maturity,omitempty" json:"date_maturity,omitempty"`
	Debit                   *Float     `xmlrpc:"debit,omitempty" json:"debit,omitempty"`
	Discount                *Float     `xmlrpc:"discount,omitempty" json:"discount,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisplayType             *Selection `xmlrpc:"display_type,omitempty" json:"display_type,omitempty"`
	ExcludeFromInvoiceTab   *Bool      `xmlrpc:"exclude_from_invoice_tab,omitempty" json:"exclude_from_invoice_tab,omitempty"`
	ExpectedPayDate         *Time      `xmlrpc:"expected_pay_date,omitempty" json:"expected_pay_date,omitempty"`
	ExpenseId               *Many2One  `xmlrpc:"expense_id,omitempty" json:"expense_id,omitempty"`
	FollowupDate            *Time      `xmlrpc:"followup_date,omitempty" json:"followup_date,omitempty"`
	FollowupLineId          *Many2One  `xmlrpc:"followup_line_id,omitempty" json:"followup_line_id,omitempty"`
	FullReconcileId         *Many2One  `xmlrpc:"full_reconcile_id,omitempty" json:"full_reconcile_id,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InternalNote            *String    `xmlrpc:"internal_note,omitempty" json:"internal_note,omitempty"`
	IntrastatTransactionId  *Many2One  `xmlrpc:"intrastat_transaction_id,omitempty" json:"intrastat_transaction_id,omitempty"`
	IsAngloSaxonLine        *Bool      `xmlrpc:"is_anglo_saxon_line,omitempty" json:"is_anglo_saxon_line,omitempty"`
	IsRoundingLine          *Bool      `xmlrpc:"is_rounding_line,omitempty" json:"is_rounding_line,omitempty"`
	JournalId               *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	MatchedCreditIds        *Relation  `xmlrpc:"matched_credit_ids,omitempty" json:"matched_credit_ids,omitempty"`
	MatchedDebitIds         *Relation  `xmlrpc:"matched_debit_ids,omitempty" json:"matched_debit_ids,omitempty"`
	MoveAttachmentIds       *Relation  `xmlrpc:"move_attachment_ids,omitempty" json:"move_attachment_ids,omitempty"`
	MoveId                  *Many2One  `xmlrpc:"move_id,omitempty" json:"move_id,omitempty"`
	MoveName                *String    `xmlrpc:"move_name,omitempty" json:"move_name,omitempty"`
	Name                    *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NextActionDate          *Time      `xmlrpc:"next_action_date,omitempty" json:"next_action_date,omitempty"`
	ParentState             *Selection `xmlrpc:"parent_state,omitempty" json:"parent_state,omitempty"`
	PartnerId               *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PaymentId               *Many2One  `xmlrpc:"payment_id,omitempty" json:"payment_id,omitempty"`
	PriceSubtotal           *Float     `xmlrpc:"price_subtotal,omitempty" json:"price_subtotal,omitempty"`
	PriceTotal              *Float     `xmlrpc:"price_total,omitempty" json:"price_total,omitempty"`
	PriceUnit               *Float     `xmlrpc:"price_unit,omitempty" json:"price_unit,omitempty"`
	ProductId               *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductUomId            *Many2One  `xmlrpc:"product_uom_id,omitempty" json:"product_uom_id,omitempty"`
	PurchaseLineId          *Many2One  `xmlrpc:"purchase_line_id,omitempty" json:"purchase_line_id,omitempty"`
	Quantity                *Float     `xmlrpc:"quantity,omitempty" json:"quantity,omitempty"`
	RecomputeTaxLine        *Bool      `xmlrpc:"recompute_tax_line,omitempty" json:"recompute_tax_line,omitempty"`
	ReconcileModelId        *Many2One  `xmlrpc:"reconcile_model_id,omitempty" json:"reconcile_model_id,omitempty"`
	Reconciled              *Bool      `xmlrpc:"reconciled,omitempty" json:"reconciled,omitempty"`
	ReconciliationInvoiceId *Relation  `xmlrpc:"reconciliation_invoice_id,omitempty" json:"reconciliation_invoice_id,omitempty"`
	Ref                     *String    `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	ReportComputeDate       *Time      `xmlrpc:"report_compute_date,omitempty" json:"report_compute_date,omitempty"`
	ReportLineIndex         *Int       `xmlrpc:"report_line_index,omitempty" json:"report_line_index,omitempty"`
	SaleLineIds             *Relation  `xmlrpc:"sale_line_ids,omitempty" json:"sale_line_ids,omitempty"`
	Sequence                *Int       `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	StatementId             *Many2One  `xmlrpc:"statement_id,omitempty" json:"statement_id,omitempty"`
	StatementLineId         *Many2One  `xmlrpc:"statement_line_id,omitempty" json:"statement_line_id,omitempty"`
	SubscriptionEndDate     *Time      `xmlrpc:"subscription_end_date,omitempty" json:"subscription_end_date,omitempty"`
	SubscriptionId          *Many2One  `xmlrpc:"subscription_id,omitempty" json:"subscription_id,omitempty"`
	SubscriptionMrr         *Float     `xmlrpc:"subscription_mrr,omitempty" json:"subscription_mrr,omitempty"`
	SubscriptionStartDate   *Time      `xmlrpc:"subscription_start_date,omitempty" json:"subscription_start_date,omitempty"`
	TagIds                  *Relation  `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TaxAudit                *String    `xmlrpc:"tax_audit,omitempty" json:"tax_audit,omitempty"`
	TaxBaseAmount           *Float     `xmlrpc:"tax_base_amount,omitempty" json:"tax_base_amount,omitempty"`
	TaxExigible             *Bool      `xmlrpc:"tax_exigible,omitempty" json:"tax_exigible,omitempty"`
	TaxGroupId              *Many2One  `xmlrpc:"tax_group_id,omitempty" json:"tax_group_id,omitempty"`
	TaxIds                  *Relation  `xmlrpc:"tax_ids,omitempty" json:"tax_ids,omitempty"`
	TaxLineId               *Many2One  `xmlrpc:"tax_line_id,omitempty" json:"tax_line_id,omitempty"`
	TaxRepartitionLineId    *Many2One  `xmlrpc:"tax_repartition_line_id,omitempty" json:"tax_repartition_line_id,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// AccountMoveLines represents array of account.move.line model.
type AccountMoveLines []AccountMoveLine

// AccountMoveLineModel is the odoo model name.
const AccountMoveLineModel = "account.move.line"

// Many2One convert AccountMoveLine to *Many2One.
func (aml *AccountMoveLine) Many2One() *Many2One {
	return NewMany2One(aml.Id.Get(), "")
}

// CreateAccountMoveLine creates a new account.move.line model and returns its id.
func (c *Client) CreateAccountMoveLine(aml *AccountMoveLine) (int64, error) {
	ids, err := c.CreateAccountMoveLines([]*AccountMoveLine{aml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveLine creates a new account.move.line model and returns its id.
func (c *Client) CreateAccountMoveLines(amls []*AccountMoveLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range amls {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveLineModel, vv, nil)
}

// UpdateAccountMoveLine updates an existing account.move.line record.
func (c *Client) UpdateAccountMoveLine(aml *AccountMoveLine) error {
	return c.UpdateAccountMoveLines([]int64{aml.Id.Get()}, aml)
}

// UpdateAccountMoveLines updates existing account.move.line records.
// All records (represented by ids) will be updated by aml values.
func (c *Client) UpdateAccountMoveLines(ids []int64, aml *AccountMoveLine) error {
	return c.Update(AccountMoveLineModel, ids, aml, nil)
}

// DeleteAccountMoveLine deletes an existing account.move.line record.
func (c *Client) DeleteAccountMoveLine(id int64) error {
	return c.DeleteAccountMoveLines([]int64{id})
}

// DeleteAccountMoveLines deletes existing account.move.line records.
func (c *Client) DeleteAccountMoveLines(ids []int64) error {
	return c.Delete(AccountMoveLineModel, ids)
}

// GetAccountMoveLine gets account.move.line existing record.
func (c *Client) GetAccountMoveLine(id int64) (*AccountMoveLine, error) {
	amls, err := c.GetAccountMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// GetAccountMoveLines gets account.move.line existing records.
func (c *Client) GetAccountMoveLines(ids []int64) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.Read(AccountMoveLineModel, ids, nil, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLine finds account.move.line record by querying it with criteria.
func (c *Client) FindAccountMoveLine(criteria *Criteria) (*AccountMoveLine, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, NewOptions().Limit(1), amls); err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// FindAccountMoveLines finds account.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLines(criteria *Criteria, options *Options) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, options, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveLineModel, criteria, options)
}

// FindAccountMoveLineId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
