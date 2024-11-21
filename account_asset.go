package odoo

// AccountAsset represents account.asset model.
type AccountAsset struct {
	AccountAssetId                 *Many2One   `xmlrpc:"account_asset_id,omitempty" json:"account_asset_id,omitempty"`
	AccountDepreciationExpenseId   *Many2One   `xmlrpc:"account_depreciation_expense_id,omitempty" json:"account_depreciation_expense_id,omitempty"`
	AccountDepreciationId          *Many2One   `xmlrpc:"account_depreciation_id,omitempty" json:"account_depreciation_id,omitempty"`
	AccountType                    *Selection  `xmlrpc:"account_type,omitempty" json:"account_type,omitempty"`
	AcquisitionDate                *Time       `xmlrpc:"acquisition_date,omitempty" json:"acquisition_date,omitempty"`
	Active                         *Bool       `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId        *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline           *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration    *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon          *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                    *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                  *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary                *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon               *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                 *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                 *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AlreadyDepreciatedAmountImport *Float      `xmlrpc:"already_depreciated_amount_import,omitempty" json:"already_depreciated_amount_import,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty" json:"analytic_distribution,omitempty"`
	AnalyticDistributionSearch     interface{} `xmlrpc:"analytic_distribution_search,omitempty" json:"analytic_distribution_search,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty" json:"analytic_precision,omitempty"`
	AssetLifetimeDays              *Float      `xmlrpc:"asset_lifetime_days,omitempty" json:"asset_lifetime_days,omitempty"`
	AssetPausedDays                *Float      `xmlrpc:"asset_paused_days,omitempty" json:"asset_paused_days,omitempty"`
	BookValue                      *Float      `xmlrpc:"book_value,omitempty" json:"book_value,omitempty"`
	ChildrenIds                    *Relation   `xmlrpc:"children_ids,omitempty" json:"children_ids,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                     *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	DepreciationEntriesCount       *Int        `xmlrpc:"depreciation_entries_count,omitempty" json:"depreciation_entries_count,omitempty"`
	DepreciationMoveIds            *Relation   `xmlrpc:"depreciation_move_ids,omitempty" json:"depreciation_move_ids,omitempty"`
	DisplayAccountAssetId          *Bool       `xmlrpc:"display_account_asset_id,omitempty" json:"display_account_asset_id,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisposalDate                   *Time       `xmlrpc:"disposal_date,omitempty" json:"disposal_date,omitempty"`
	GrossIncreaseCount             *Int        `xmlrpc:"gross_increase_count,omitempty" json:"gross_increase_count,omitempty"`
	GrossIncreaseValue             *Float      `xmlrpc:"gross_increase_value,omitempty" json:"gross_increase_value,omitempty"`
	HasMessage                     *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	JournalId                      *Many2One   `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	MessageAttachmentCount         *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds             *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError                *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter         *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError             *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                     *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower              *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction              *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter       *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds              *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	Method                         *Selection  `xmlrpc:"method,omitempty" json:"method,omitempty"`
	MethodNumber                   *Int        `xmlrpc:"method_number,omitempty" json:"method_number,omitempty"`
	MethodPeriod                   *Selection  `xmlrpc:"method_period,omitempty" json:"method_period,omitempty"`
	MethodProgressFactor           *Float      `xmlrpc:"method_progress_factor,omitempty" json:"method_progress_factor,omitempty"`
	ModelId                        *Many2One   `xmlrpc:"model_id,omitempty" json:"model_id,omitempty"`
	MyActivityDateDeadline         *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                           *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NonDeductibleTaxValue          *Float      `xmlrpc:"non_deductible_tax_value,omitempty" json:"non_deductible_tax_value,omitempty"`
	OriginalMoveLineIds            *Relation   `xmlrpc:"original_move_line_ids,omitempty" json:"original_move_line_ids,omitempty"`
	OriginalValue                  *Float      `xmlrpc:"original_value,omitempty" json:"original_value,omitempty"`
	ParentId                       *Many2One   `xmlrpc:"parent_id,omitempty" json:"parent_id,omitempty"`
	PausedProrataDate              *Time       `xmlrpc:"paused_prorata_date,omitempty" json:"paused_prorata_date,omitempty"`
	ProrataComputationType         *Selection  `xmlrpc:"prorata_computation_type,omitempty" json:"prorata_computation_type,omitempty"`
	ProrataDate                    *Time       `xmlrpc:"prorata_date,omitempty" json:"prorata_date,omitempty"`
	RatingIds                      *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RelatedPurchaseValue           *Float      `xmlrpc:"related_purchase_value,omitempty" json:"related_purchase_value,omitempty"`
	SalvageValue                   *Float      `xmlrpc:"salvage_value,omitempty" json:"salvage_value,omitempty"`
	State                          *Selection  `xmlrpc:"state,omitempty" json:"state,omitempty"`
	TotalDepreciableValue          *Float      `xmlrpc:"total_depreciable_value,omitempty" json:"total_depreciable_value,omitempty"`
	TotalDepreciationEntriesCount  *Int        `xmlrpc:"total_depreciation_entries_count,omitempty" json:"total_depreciation_entries_count,omitempty"`
	ValueResidual                  *Float      `xmlrpc:"value_residual,omitempty" json:"value_residual,omitempty"`
	WebsiteMessageIds              *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// AccountAssets represents array of account.asset model.
type AccountAssets []AccountAsset

// AccountAssetModel is the odoo model name.
const AccountAssetModel = "account.asset"

// Many2One convert AccountAsset to *Many2One.
func (aa *AccountAsset) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAsset(aa *AccountAsset) (int64, error) {
	ids, err := c.CreateAccountAssets([]*AccountAsset{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAssets(aas []*AccountAsset) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetModel, vv, nil)
}

// UpdateAccountAsset updates an existing account.asset record.
func (c *Client) UpdateAccountAsset(aa *AccountAsset) error {
	return c.UpdateAccountAssets([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAssets updates existing account.asset records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAssets(ids []int64, aa *AccountAsset) error {
	return c.Update(AccountAssetModel, ids, aa, nil)
}

// DeleteAccountAsset deletes an existing account.asset record.
func (c *Client) DeleteAccountAsset(id int64) error {
	return c.DeleteAccountAssets([]int64{id})
}

// DeleteAccountAssets deletes existing account.asset records.
func (c *Client) DeleteAccountAssets(ids []int64) error {
	return c.Delete(AccountAssetModel, ids)
}

// GetAccountAsset gets account.asset existing record.
func (c *Client) GetAccountAsset(id int64) (*AccountAsset, error) {
	aas, err := c.GetAccountAssets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// GetAccountAssets gets account.asset existing records.
func (c *Client) GetAccountAssets(ids []int64) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.Read(AccountAssetModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAsset finds account.asset record by querying it with criteria.
func (c *Client) FindAccountAsset(criteria *Criteria) (*AccountAsset, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// FindAccountAssets finds account.asset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssets(criteria *Criteria, options *Options) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAssetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAssetModel, criteria, options)
}

// FindAccountAssetId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
