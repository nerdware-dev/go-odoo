package odoo

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omitempty"`
	AccountInvoice                        *String    `xmlrpc:"account_invoice,omitempty"`
	AccountRefund                         *String    `xmlrpc:"account_refund,omitempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountByGroup                         *String    `xmlrpc:"amount_by_group,omitempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omitempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omitempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omitempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omitempty"`
	AssetAssetType                        *Selection `xmlrpc:"asset_asset_type,omitempty"`
	AssetDepreciatedValue                 *Float     `xmlrpc:"asset_depreciated_value,omitempty"`
	AssetId                               *Many2One  `xmlrpc:"asset_id,omitempty"`
	AssetIdDisplayName                    *String    `xmlrpc:"asset_id_display_name,omitempty"`
	AssetIds                              *Relation  `xmlrpc:"asset_ids,omitempty"`
	AssetIdsDisplayName                   *String    `xmlrpc:"asset_ids_display_name,omitempty"`
	AssetManuallyModified                 *Bool      `xmlrpc:"asset_manually_modified,omitempty"`
	AssetRemainingValue                   *Float     `xmlrpc:"asset_remaining_value,omitempty"`
	AssetValueChange                      *Bool      `xmlrpc:"asset_value_change,omitempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Bool      `xmlrpc:"auto_post,omitempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time      `xmlrpc:"date,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	DocumentRequestLineId                 *Many2One  `xmlrpc:"document_request_line_id,omitempty"`
	DraftAssetIds                         *Bool      `xmlrpc:"draft_asset_ids,omitempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	ForceReleaseToPay                     *Bool      `xmlrpc:"force_release_to_pay,omitempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omitempty"`
	IntrastatCountryId                    *Many2One  `xmlrpc:"intrastat_country_id,omitempty"`
	IntrastatTransportModeId              *Many2One  `xmlrpc:"intrastat_transport_mode_id,omitempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasMatchingSuspenseAmount      *Bool      `xmlrpc:"invoice_has_matching_suspense_amount,omitempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerBankId                  *Many2One  `xmlrpc:"invoice_partner_bank_id,omitempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePartnerIcon                    *String    `xmlrpc:"invoice_partner_icon,omitempty"`
	InvoicePaymentRef                     *String    `xmlrpc:"invoice_payment_ref,omitempty"`
	InvoicePaymentState                   *Selection `xmlrpc:"invoice_payment_state,omitempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoiceSent                           *Bool      `xmlrpc:"invoice_sent,omitempty"`
	InvoiceSequenceNumberNext             *String    `xmlrpc:"invoice_sequence_number_next,omitempty"`
	InvoiceSequenceNumberNextPrefix       *String    `xmlrpc:"invoice_sequence_number_next_prefix,omitempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsTaxClosing                          *Bool      `xmlrpc:"is_tax_closing,omitempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omitempty"`
	L10NDeDatevMainAccountId              *Many2One  `xmlrpc:"l10n_de_datev_main_account_id,omitempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omitempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds                     *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                         *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter                  *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	Narration                             *String    `xmlrpc:"narration,omitempty"`
	NumberAssetIds                        *Int       `xmlrpc:"number_asset_ids,omitempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PurchaseId                            *Many2One  `xmlrpc:"purchase_id,omitempty"`
	PurchaseVendorBillId                  *Many2One  `xmlrpc:"purchase_vendor_bill_id,omitempty"`
	Ref                                   *String    `xmlrpc:"ref,omitempty"`
	ReleaseToPay                          *Selection `xmlrpc:"release_to_pay,omitempty"`
	ReleaseToPayManual                    *Selection `xmlrpc:"release_to_pay_manual,omitempty"`
	ReportComputeDate                     *Time      `xmlrpc:"report_compute_date,omitempty"`
	ReportLineIndex                       *Int       `xmlrpc:"report_line_index,omitempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omitempty"`
	SddPayingMandateId                    *Many2One  `xmlrpc:"sdd_paying_mandate_id,omitempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omitempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omitempty"`
	State                                 *Selection `xmlrpc:"state,omitempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omitempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omitempty"`
	TaxReportControlError                 *Bool      `xmlrpc:"tax_report_control_error,omitempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omitempty"`
	TimesheetCount                        *Int       `xmlrpc:"timesheet_count,omitempty"`
	TimesheetIds                          *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omitempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TransferModelId                       *Many2One  `xmlrpc:"transfer_model_id,omitempty"`
	TungstenFileIndex                     *Int       `xmlrpc:"tungsten_file_index,omitempty"`
	Type                                  *Selection `xmlrpc:"type,omitempty"`
	TypeName                              *String    `xmlrpc:"type_name,omitempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountMoves represents array of account.move model.
type AccountMoves []AccountMove

// AccountMoveModel is the odoo model name.
const AccountMoveModel = "account.move"

// Many2One convert AccountMove to *Many2One.
func (am *AccountMove) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMove(am *AccountMove) (int64, error) {
	ids, err := c.CreateAccountMoves([]*AccountMove{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMoves(ams []*AccountMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveModel, vv, nil)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am, nil)
}

// DeleteAccountMove deletes an existing account.move record.
func (c *Client) DeleteAccountMove(id int64) error {
	return c.DeleteAccountMoves([]int64{id})
}

// DeleteAccountMoves deletes existing account.move records.
func (c *Client) DeleteAccountMoves(ids []int64) error {
	return c.Delete(AccountMoveModel, ids)
}

// GetAccountMove gets account.move existing record.
func (c *Client) GetAccountMove(id int64) (*AccountMove, error) {
	ams, err := c.GetAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAccountMoves gets account.move existing records.
func (c *Client) GetAccountMoves(ids []int64) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.Read(AccountMoveModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMove finds account.move record by querying it with criteria.
func (c *Client) FindAccountMove(criteria *Criteria) (*AccountMove, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAccountMoves finds account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoves(criteria *Criteria, options *Options) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveModel, criteria, options)
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
