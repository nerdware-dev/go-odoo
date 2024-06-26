package odoo

import (
	"fmt"
)

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omptempty"`
	AccountInvoice                        *String    `xmlrpc:"account_invoice,omptempty"`
	AccountRefund                         *String    `xmlrpc:"account_refund,omptempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountByGroup                         *String    `xmlrpc:"amount_by_group,omptempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omptempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omptempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omptempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omptempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omptempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omptempty"`
	AssetAssetType                        *Selection `xmlrpc:"asset_asset_type,omptempty"`
	AssetDepreciatedValue                 *Float     `xmlrpc:"asset_depreciated_value,omptempty"`
	AssetId                               *Many2One  `xmlrpc:"asset_id,omptempty"`
	AssetIdDisplayName                    *String    `xmlrpc:"asset_id_display_name,omptempty"`
	AssetIds                              *Relation  `xmlrpc:"asset_ids,omptempty"`
	AssetIdsDisplayName                   *String    `xmlrpc:"asset_ids_display_name,omptempty"`
	AssetManuallyModified                 *Bool      `xmlrpc:"asset_manually_modified,omptempty"`
	AssetRemainingValue                   *Float     `xmlrpc:"asset_remaining_value,omptempty"`
	AssetValueChange                      *Bool      `xmlrpc:"asset_value_change,omptempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omptempty"`
	AutoPost                              *Bool      `xmlrpc:"auto_post,omptempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omptempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                                  *Time      `xmlrpc:"date,omptempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omptempty"`
	DocumentRequestLineId                 *Many2One  `xmlrpc:"document_request_line_id,omptempty"`
	DraftAssetIds                         *Bool      `xmlrpc:"draft_asset_ids,omptempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	ForceReleaseToPay                     *Bool      `xmlrpc:"force_release_to_pay,omptempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omptempty"`
	Id                                    *Int       `xmlrpc:"id,omptempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omptempty"`
	IntrastatCountryId                    *Many2One  `xmlrpc:"intrastat_country_id,omptempty"`
	IntrastatTransportModeId              *Many2One  `xmlrpc:"intrastat_transport_mode_id,omptempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omptempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omptempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omptempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omptempty"`
	InvoiceHasMatchingSuspenseAmount      *Bool      `xmlrpc:"invoice_has_matching_suspense_amount,omptempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omptempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omptempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omptempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omptempty"`
	InvoicePartnerBankId                  *Many2One  `xmlrpc:"invoice_partner_bank_id,omptempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omptempty"`
	InvoicePartnerIcon                    *String    `xmlrpc:"invoice_partner_icon,omptempty"`
	InvoicePaymentRef                     *String    `xmlrpc:"invoice_payment_ref,omptempty"`
	InvoicePaymentState                   *Selection `xmlrpc:"invoice_payment_state,omptempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omptempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omptempty"`
	InvoiceSent                           *Bool      `xmlrpc:"invoice_sent,omptempty"`
	InvoiceSequenceNumberNext             *String    `xmlrpc:"invoice_sequence_number_next,omptempty"`
	InvoiceSequenceNumberNextPrefix       *String    `xmlrpc:"invoice_sequence_number_next_prefix,omptempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omptempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omptempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omptempty"`
	IsTaxClosing                          *Bool      `xmlrpc:"is_tax_closing,omptempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omptempty"`
	L10NDeDatevMainAccountId              *Many2One  `xmlrpc:"l10n_de_datev_main_account_id,omptempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omptempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds                     *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                         *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter                  *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                                  *String    `xmlrpc:"name,omptempty"`
	Narration                             *String    `xmlrpc:"narration,omptempty"`
	NumberAssetIds                        *Int       `xmlrpc:"number_asset_ids,omptempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PurchaseId                            *Many2One  `xmlrpc:"purchase_id,omptempty"`
	PurchaseVendorBillId                  *Many2One  `xmlrpc:"purchase_vendor_bill_id,omptempty"`
	Ref                                   *String    `xmlrpc:"ref,omptempty"`
	ReleaseToPay                          *Selection `xmlrpc:"release_to_pay,omptempty"`
	ReleaseToPayManual                    *Selection `xmlrpc:"release_to_pay_manual,omptempty"`
	ReportComputeDate                     *Time      `xmlrpc:"report_compute_date,omptempty"`
	ReportLineIndex                       *Int       `xmlrpc:"report_line_index,omptempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omptempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omptempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omptempty"`
	SddPayingMandateId                    *Many2One  `xmlrpc:"sdd_paying_mandate_id,omptempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omptempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omptempty"`
	State                                 *Selection `xmlrpc:"state,omptempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omptempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omptempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omptempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omptempty"`
	TaxReportControlError                 *Bool      `xmlrpc:"tax_report_control_error,omptempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omptempty"`
	TimesheetCount                        *Int       `xmlrpc:"timesheet_count,omptempty"`
	TimesheetIds                          *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omptempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omptempty"`
	TransferModelId                       *Many2One  `xmlrpc:"transfer_model_id,omptempty"`
	TungstenFileIndex                     *Int       `xmlrpc:"tungsten_file_index,omptempty"`
	Type                                  *Selection `xmlrpc:"type,omptempty"`
	TypeName                              *String    `xmlrpc:"type_name,omptempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(AccountMoveModel, vv)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am)
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
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move not found", id)
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
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("account.move was not found with criteria %v", criteria)
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
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move was not found with criteria %v and options %v", criteria, options)
}
