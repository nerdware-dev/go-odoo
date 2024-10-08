package odoo

// AccountReconcileModel represents account.reconcile.model model.
type AccountReconcileModel struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId                  *Many2One  `xmlrpc:"account_id,omitempty"`
	ActivityTypeId             *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	Amount                     *Float     `xmlrpc:"amount,omitempty"`
	AmountFromLabelRegex       *String    `xmlrpc:"amount_from_label_regex,omitempty"`
	AmountType                 *Selection `xmlrpc:"amount_type,omitempty"`
	AnalyticAccountId          *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AnalyticTagIds             *Relation  `xmlrpc:"analytic_tag_ids,omitempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	ForceSecondTaxIncluded     *Bool      `xmlrpc:"force_second_tax_included,omitempty"`
	ForceTaxIncluded           *Bool      `xmlrpc:"force_tax_included,omitempty"`
	HasSecondLine              *Bool      `xmlrpc:"has_second_line,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	JournalId                  *Many2One  `xmlrpc:"journal_id,omitempty"`
	Label                      *String    `xmlrpc:"label,omitempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omitempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omitempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omitempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omitempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omitempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omitempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omitempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omitempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omitempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omitempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omitempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omitempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omitempty"`
	MatchTotalAmount           *Bool      `xmlrpc:"match_total_amount,omitempty"`
	MatchTotalAmountParam      *Float     `xmlrpc:"match_total_amount_param,omitempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omitempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NumberEntries              *Int       `xmlrpc:"number_entries,omitempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omitempty"`
	SecondAccountId            *Many2One  `xmlrpc:"second_account_id,omitempty"`
	SecondAmount               *Float     `xmlrpc:"second_amount,omitempty"`
	SecondAmountFromLabelRegex *String    `xmlrpc:"second_amount_from_label_regex,omitempty"`
	SecondAmountType           *Selection `xmlrpc:"second_amount_type,omitempty"`
	SecondAnalyticAccountId    *Many2One  `xmlrpc:"second_analytic_account_id,omitempty"`
	SecondAnalyticTagIds       *Relation  `xmlrpc:"second_analytic_tag_ids,omitempty"`
	SecondJournalId            *Many2One  `xmlrpc:"second_journal_id,omitempty"`
	SecondLabel                *String    `xmlrpc:"second_label,omitempty"`
	SecondTaxIds               *Relation  `xmlrpc:"second_tax_ids,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	ShowForceTaxIncluded       *Bool      `xmlrpc:"show_force_tax_included,omitempty"`
	ShowSecondForceTaxIncluded *Bool      `xmlrpc:"show_second_force_tax_included,omitempty"`
	TaxIds                     *Relation  `xmlrpc:"tax_ids,omitempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModels represents array of account.reconcile.model model.
type AccountReconcileModels []AccountReconcileModel

// AccountReconcileModelModel is the odoo model name.
const AccountReconcileModelModel = "account.reconcile.model"

// Many2One convert AccountReconcileModel to *Many2One.
func (arm *AccountReconcileModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModel(arm *AccountReconcileModel) (int64, error) {
	ids, err := c.CreateAccountReconcileModels([]*AccountReconcileModel{arm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModels(arms []*AccountReconcileModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range arms {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelModel, vv, nil)
}

// UpdateAccountReconcileModel updates an existing account.reconcile.model record.
func (c *Client) UpdateAccountReconcileModel(arm *AccountReconcileModel) error {
	return c.UpdateAccountReconcileModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReconcileModels updates existing account.reconcile.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReconcileModels(ids []int64, arm *AccountReconcileModel) error {
	return c.Update(AccountReconcileModelModel, ids, arm, nil)
}

// DeleteAccountReconcileModel deletes an existing account.reconcile.model record.
func (c *Client) DeleteAccountReconcileModel(id int64) error {
	return c.DeleteAccountReconcileModels([]int64{id})
}

// DeleteAccountReconcileModels deletes existing account.reconcile.model records.
func (c *Client) DeleteAccountReconcileModels(ids []int64) error {
	return c.Delete(AccountReconcileModelModel, ids)
}

// GetAccountReconcileModel gets account.reconcile.model existing record.
func (c *Client) GetAccountReconcileModel(id int64) (*AccountReconcileModel, error) {
	arms, err := c.GetAccountReconcileModels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// GetAccountReconcileModels gets account.reconcile.model existing records.
func (c *Client) GetAccountReconcileModels(ids []int64) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.Read(AccountReconcileModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModel finds account.reconcile.model record by querying it with criteria.
func (c *Client) FindAccountReconcileModel(criteria *Criteria) (*AccountReconcileModel, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// FindAccountReconcileModels finds account.reconcile.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModels(criteria *Criteria, options *Options) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelModel, criteria, options)
}

// FindAccountReconcileModelId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
