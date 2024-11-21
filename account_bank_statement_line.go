package odoo

// AccountBankStatementLine represents account.bank.statement.line model.
type AccountBankStatementLine struct {
	AccessToken                           *String     `xmlrpc:"access_token,omitempty" json:"access_token,omitempty"`
	AccessUrl                             *String     `xmlrpc:"access_url,omitempty" json:"access_url,omitempty"`
	AccessWarning                         *String     `xmlrpc:"access_warning,omitempty" json:"access_warning,omitempty"`
	AccountNumber                         *String     `xmlrpc:"account_number,omitempty" json:"account_number,omitempty"`
	ActivityCalendarEventId               *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                         *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary                       *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AlwaysTaxExigible                     *Bool       `xmlrpc:"always_tax_exigible,omitempty" json:"always_tax_exigible,omitempty"`
	Amount                                *Float      `xmlrpc:"amount,omitempty" json:"amount,omitempty"`
	AmountCurrency                        *Float      `xmlrpc:"amount_currency,omitempty" json:"amount_currency,omitempty"`
	AmountPaid                            *Float      `xmlrpc:"amount_paid,omitempty" json:"amount_paid,omitempty"`
	AmountResidual                        *Float      `xmlrpc:"amount_residual,omitempty" json:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float      `xmlrpc:"amount_residual_signed,omitempty" json:"amount_residual_signed,omitempty"`
	AmountTax                             *Float      `xmlrpc:"amount_tax,omitempty" json:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float      `xmlrpc:"amount_tax_signed,omitempty" json:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float      `xmlrpc:"amount_total,omitempty" json:"amount_total,omitempty"`
	AmountTotalInCurrencySigned           *Float      `xmlrpc:"amount_total_in_currency_signed,omitempty" json:"amount_total_in_currency_signed,omitempty"`
	AmountTotalSigned                     *Float      `xmlrpc:"amount_total_signed,omitempty" json:"amount_total_signed,omitempty"`
	AmountTotalWords                      *String     `xmlrpc:"amount_total_words,omitempty" json:"amount_total_words,omitempty"`
	AmountUntaxed                         *Float      `xmlrpc:"amount_untaxed,omitempty" json:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float      `xmlrpc:"amount_untaxed_signed,omitempty" json:"amount_untaxed_signed,omitempty"`
	AssetDepreciatedValue                 *Float      `xmlrpc:"asset_depreciated_value,omitempty" json:"asset_depreciated_value,omitempty"`
	AssetDepreciationBeginningDate        *Time       `xmlrpc:"asset_depreciation_beginning_date,omitempty" json:"asset_depreciation_beginning_date,omitempty"`
	AssetId                               *Many2One   `xmlrpc:"asset_id,omitempty" json:"asset_id,omitempty"`
	AssetIdDisplayName                    *String     `xmlrpc:"asset_id_display_name,omitempty" json:"asset_id_display_name,omitempty"`
	AssetIds                              *Relation   `xmlrpc:"asset_ids,omitempty" json:"asset_ids,omitempty"`
	AssetNumberDays                       *Int        `xmlrpc:"asset_number_days,omitempty" json:"asset_number_days,omitempty"`
	AssetRemainingValue                   *Float      `xmlrpc:"asset_remaining_value,omitempty" json:"asset_remaining_value,omitempty"`
	AssetValueChange                      *Bool       `xmlrpc:"asset_value_change,omitempty" json:"asset_value_change,omitempty"`
	AttachmentIds                         *Relation   `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	AuthorizedTransactionIds              *Relation   `xmlrpc:"authorized_transaction_ids,omitempty" json:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Selection  `xmlrpc:"auto_post,omitempty" json:"auto_post,omitempty"`
	AutoPostOriginId                      *Many2One   `xmlrpc:"auto_post_origin_id,omitempty" json:"auto_post_origin_id,omitempty"`
	AutoPostUntil                         *Time       `xmlrpc:"auto_post_until,omitempty" json:"auto_post_until,omitempty"`
	BankPartnerId                         *Many2One   `xmlrpc:"bank_partner_id,omitempty" json:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One   `xmlrpc:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One   `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One   `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyId                             *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountAsset                            *Int        `xmlrpc:"count_asset,omitempty" json:"count_asset,omitempty"`
	CountryCode                           *String     `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CronLastCheck                         *Time       `xmlrpc:"cron_last_check,omitempty" json:"cron_last_check,omitempty"`
	CurrencyId                            *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                                  *Time       `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DeferredEntryType                     *Selection  `xmlrpc:"deferred_entry_type,omitempty" json:"deferred_entry_type,omitempty"`
	DeferredMoveIds                       *Relation   `xmlrpc:"deferred_move_ids,omitempty" json:"deferred_move_ids,omitempty"`
	DeferredOriginalMoveIds               *Relation   `xmlrpc:"deferred_original_move_ids,omitempty" json:"deferred_original_move_ids,omitempty"`
	DeliveryDate                          *Time       `xmlrpc:"delivery_date,omitempty" json:"delivery_date,omitempty"`
	DepreciationValue                     *Float      `xmlrpc:"depreciation_value,omitempty" json:"depreciation_value,omitempty"`
	DirectionSign                         *Int        `xmlrpc:"direction_sign,omitempty" json:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool       `xmlrpc:"display_inactive_currency_warning,omitempty" json:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisplayQrCode                         *Bool       `xmlrpc:"display_qr_code,omitempty" json:"display_qr_code,omitempty"`
	DraftAssetExists                      *Bool       `xmlrpc:"draft_asset_exists,omitempty" json:"draft_asset_exists,omitempty"`
	DuplicatedRefIds                      *Relation   `xmlrpc:"duplicated_ref_ids,omitempty" json:"duplicated_ref_ids,omitempty"`
	ExpenseSheetId                        *Many2One   `xmlrpc:"expense_sheet_id,omitempty" json:"expense_sheet_id,omitempty"`
	ExtractAttachmentId                   *Many2One   `xmlrpc:"extract_attachment_id,omitempty" json:"extract_attachment_id,omitempty"`
	ExtractCanShowBanners                 *Bool       `xmlrpc:"extract_can_show_banners,omitempty" json:"extract_can_show_banners,omitempty"`
	ExtractCanShowSendButton              *Bool       `xmlrpc:"extract_can_show_send_button,omitempty" json:"extract_can_show_send_button,omitempty"`
	ExtractDetectedLayout                 *Int        `xmlrpc:"extract_detected_layout,omitempty" json:"extract_detected_layout,omitempty"`
	ExtractDocumentUuid                   *String     `xmlrpc:"extract_document_uuid,omitempty" json:"extract_document_uuid,omitempty"`
	ExtractErrorMessage                   *String     `xmlrpc:"extract_error_message,omitempty" json:"extract_error_message,omitempty"`
	ExtractPartnerName                    *String     `xmlrpc:"extract_partner_name,omitempty" json:"extract_partner_name,omitempty"`
	ExtractState                          *Selection  `xmlrpc:"extract_state,omitempty" json:"extract_state,omitempty"`
	ExtractStateProcessed                 *Bool       `xmlrpc:"extract_state_processed,omitempty" json:"extract_state_processed,omitempty"`
	ExtractStatus                         *String     `xmlrpc:"extract_status,omitempty" json:"extract_status,omitempty"`
	ExtractWordIds                        *Relation   `xmlrpc:"extract_word_ids,omitempty" json:"extract_word_ids,omitempty"`
	FiscalPositionId                      *Many2One   `xmlrpc:"fiscal_position_id,omitempty" json:"fiscal_position_id,omitempty"`
	ForeignCurrencyId                     *Many2One   `xmlrpc:"foreign_currency_id,omitempty" json:"foreign_currency_id,omitempty"`
	HasMessage                            *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	HasReconciledEntries                  *Bool       `xmlrpc:"has_reconciled_entries,omitempty" json:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool       `xmlrpc:"hide_post_button,omitempty" json:"hide_post_button,omitempty"`
	HighestName                           *String     `xmlrpc:"highest_name,omitempty" json:"highest_name,omitempty"`
	Id                                    *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InalterableHash                       *String     `xmlrpc:"inalterable_hash,omitempty" json:"inalterable_hash,omitempty"`
	IncotermLocation                      *String     `xmlrpc:"incoterm_location,omitempty" json:"incoterm_location,omitempty"`
	InternalIndex                         *String     `xmlrpc:"internal_index,omitempty" json:"internal_index,omitempty"`
	InvoiceCashRoundingId                 *Many2One   `xmlrpc:"invoice_cash_rounding_id,omitempty" json:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time       `xmlrpc:"invoice_date,omitempty" json:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time       `xmlrpc:"invoice_date_due,omitempty" json:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String     `xmlrpc:"invoice_filter_type_domain,omitempty" json:"invoice_filter_type_domain,omitempty"`
	InvoiceHasOutstanding                 *Bool       `xmlrpc:"invoice_has_outstanding,omitempty" json:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One   `xmlrpc:"invoice_incoterm_id,omitempty" json:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation   `xmlrpc:"invoice_line_ids,omitempty" json:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String     `xmlrpc:"invoice_origin,omitempty" json:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String     `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty" json:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String     `xmlrpc:"invoice_partner_display_name,omitempty" json:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One   `xmlrpc:"invoice_payment_term_id,omitempty" json:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String     `xmlrpc:"invoice_payments_widget,omitempty" json:"invoice_payments_widget,omitempty"`
	InvoicePdfReportFile                  *String     `xmlrpc:"invoice_pdf_report_file,omitempty" json:"invoice_pdf_report_file,omitempty"`
	InvoicePdfReportId                    *Many2One   `xmlrpc:"invoice_pdf_report_id,omitempty" json:"invoice_pdf_report_id,omitempty"`
	InvoiceSourceEmail                    *String     `xmlrpc:"invoice_source_email,omitempty" json:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One   `xmlrpc:"invoice_user_id,omitempty" json:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One   `xmlrpc:"invoice_vendor_bill_id,omitempty" json:"invoice_vendor_bill_id,omitempty"`
	IsBeingSent                           *Bool       `xmlrpc:"is_being_sent,omitempty" json:"is_being_sent,omitempty"`
	IsInExtractableState                  *Bool       `xmlrpc:"is_in_extractable_state,omitempty" json:"is_in_extractable_state,omitempty"`
	IsMoveSent                            *Bool       `xmlrpc:"is_move_sent,omitempty" json:"is_move_sent,omitempty"`
	IsReconciled                          *Bool       `xmlrpc:"is_reconciled,omitempty" json:"is_reconciled,omitempty"`
	IsStorno                              *Bool       `xmlrpc:"is_storno,omitempty" json:"is_storno,omitempty"`
	JournalId                             *Many2One   `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	L10NDeDatevMainAccountId              *Many2One   `xmlrpc:"l10n_de_datev_main_account_id,omitempty" json:"l10n_de_datev_main_account_id,omitempty"`
	L10NDin5008Addresses                  *String     `xmlrpc:"l10n_din5008_addresses,omitempty" json:"l10n_din5008_addresses,omitempty"`
	L10NDin5008DocumentTitle              *String     `xmlrpc:"l10n_din5008_document_title,omitempty" json:"l10n_din5008_document_title,omitempty"`
	L10NDin5008TemplateData               *String     `xmlrpc:"l10n_din5008_template_data,omitempty" json:"l10n_din5008_template_data,omitempty"`
	LineIds                               *Relation   `xmlrpc:"line_ids,omitempty" json:"line_ids,omitempty"`
	MadeSequenceHole                      *Bool       `xmlrpc:"made_sequence_hole,omitempty" json:"made_sequence_hole,omitempty"`
	MediumId                              *Many2One   `xmlrpc:"medium_id,omitempty" json:"medium_id,omitempty"`
	MessageAttachmentCount                *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower                     *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One   `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MoveId                                *Many2One   `xmlrpc:"move_id,omitempty" json:"move_id,omitempty"`
	MoveType                              *Selection  `xmlrpc:"move_type,omitempty" json:"move_type,omitempty"`
	MyActivityDateDeadline                *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                                  *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Narration                             *String     `xmlrpc:"narration,omitempty" json:"narration,omitempty"`
	NeedCancelRequest                     *Bool       `xmlrpc:"need_cancel_request,omitempty" json:"need_cancel_request,omitempty"`
	NeededTerms                           *String     `xmlrpc:"needed_terms,omitempty" json:"needed_terms,omitempty"`
	NeededTermsDirty                      *Bool       `xmlrpc:"needed_terms_dirty,omitempty" json:"needed_terms_dirty,omitempty"`
	OnlineAccountId                       *Many2One   `xmlrpc:"online_account_id,omitempty" json:"online_account_id,omitempty"`
	OnlineLinkId                          *Many2One   `xmlrpc:"online_link_id,omitempty" json:"online_link_id,omitempty"`
	OnlinePartnerInformation              *String     `xmlrpc:"online_partner_information,omitempty" json:"online_partner_information,omitempty"`
	OnlineTransactionIdentifier           *String     `xmlrpc:"online_transaction_identifier,omitempty" json:"online_transaction_identifier,omitempty"`
	PartnerBankId                         *Many2One   `xmlrpc:"partner_bank_id,omitempty" json:"partner_bank_id,omitempty"`
	PartnerCredit                         *Float      `xmlrpc:"partner_credit,omitempty" json:"partner_credit,omitempty"`
	PartnerCreditWarning                  *String     `xmlrpc:"partner_credit_warning,omitempty" json:"partner_credit_warning,omitempty"`
	PartnerId                             *Many2One   `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerName                           *String     `xmlrpc:"partner_name,omitempty" json:"partner_name,omitempty"`
	PartnerShippingId                     *Many2One   `xmlrpc:"partner_shipping_id,omitempty" json:"partner_shipping_id,omitempty"`
	PaymentId                             *Many2One   `xmlrpc:"payment_id,omitempty" json:"payment_id,omitempty"`
	PaymentIds                            *Relation   `xmlrpc:"payment_ids,omitempty" json:"payment_ids,omitempty"`
	PaymentRef                            *String     `xmlrpc:"payment_ref,omitempty" json:"payment_ref,omitempty"`
	PaymentReference                      *String     `xmlrpc:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	PaymentState                          *Selection  `xmlrpc:"payment_state,omitempty" json:"payment_state,omitempty"`
	PaymentStateBeforeSwitch              *String     `xmlrpc:"payment_state_before_switch,omitempty" json:"payment_state_before_switch,omitempty"`
	PaymentTermDetails                    *String     `xmlrpc:"payment_term_details,omitempty" json:"payment_term_details,omitempty"`
	PostedBefore                          *Bool       `xmlrpc:"posted_before,omitempty" json:"posted_before,omitempty"`
	PurchaseId                            *Many2One   `xmlrpc:"purchase_id,omitempty" json:"purchase_id,omitempty"`
	PurchaseOrderCount                    *Int        `xmlrpc:"purchase_order_count,omitempty" json:"purchase_order_count,omitempty"`
	PurchaseVendorBillId                  *Many2One   `xmlrpc:"purchase_vendor_bill_id,omitempty" json:"purchase_vendor_bill_id,omitempty"`
	QrCodeMethod                          *Selection  `xmlrpc:"qr_code_method,omitempty" json:"qr_code_method,omitempty"`
	QuickEditMode                         *Bool       `xmlrpc:"quick_edit_mode,omitempty" json:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float      `xmlrpc:"quick_edit_total_amount,omitempty" json:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     *String     `xmlrpc:"quick_encoding_vals,omitempty" json:"quick_encoding_vals,omitempty"`
	RatingIds                             *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	Ref                                   *String     `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	RestrictModeHashTable                 *Bool       `xmlrpc:"restrict_mode_hash_table,omitempty" json:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation   `xmlrpc:"reversal_move_id,omitempty" json:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One   `xmlrpc:"reversed_entry_id,omitempty" json:"reversed_entry_id,omitempty"`
	RunningBalance                        *Float      `xmlrpc:"running_balance,omitempty" json:"running_balance,omitempty"`
	SaleOrderCount                        *Int        `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int        `xmlrpc:"secure_sequence_number,omitempty" json:"secure_sequence_number,omitempty"`
	SendAndPrintValues                    interface{} `xmlrpc:"send_and_print_values,omitempty" json:"send_and_print_values,omitempty"`
	Sequence                              *Int        `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	SequenceNumber                        *Int        `xmlrpc:"sequence_number,omitempty" json:"sequence_number,omitempty"`
	SequencePrefix                        *String     `xmlrpc:"sequence_prefix,omitempty" json:"sequence_prefix,omitempty"`
	ShowCommercialPartnerWarning          *Bool       `xmlrpc:"show_commercial_partner_warning,omitempty" json:"show_commercial_partner_warning,omitempty"`
	ShowDeliveryDate                      *Bool       `xmlrpc:"show_delivery_date,omitempty" json:"show_delivery_date,omitempty"`
	ShowDiscountDetails                   *Bool       `xmlrpc:"show_discount_details,omitempty" json:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool       `xmlrpc:"show_name_warning,omitempty" json:"show_name_warning,omitempty"`
	ShowPaymentTermDetails                *Bool       `xmlrpc:"show_payment_term_details,omitempty" json:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool       `xmlrpc:"show_reset_to_draft_button,omitempty" json:"show_reset_to_draft_button,omitempty"`
	ShowUpdateFpos                        *Bool       `xmlrpc:"show_update_fpos,omitempty" json:"show_update_fpos,omitempty"`
	SourceId                              *Many2One   `xmlrpc:"source_id,omitempty" json:"source_id,omitempty"`
	State                                 *Selection  `xmlrpc:"state,omitempty" json:"state,omitempty"`
	StatementBalanceEndReal               *Float      `xmlrpc:"statement_balance_end_real,omitempty" json:"statement_balance_end_real,omitempty"`
	StatementComplete                     *Bool       `xmlrpc:"statement_complete,omitempty" json:"statement_complete,omitempty"`
	StatementId                           *Many2One   `xmlrpc:"statement_id,omitempty" json:"statement_id,omitempty"`
	StatementLineId                       *Many2One   `xmlrpc:"statement_line_id,omitempty" json:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation   `xmlrpc:"statement_line_ids,omitempty" json:"statement_line_ids,omitempty"`
	StatementName                         *String     `xmlrpc:"statement_name,omitempty" json:"statement_name,omitempty"`
	StatementValid                        *Bool       `xmlrpc:"statement_valid,omitempty" json:"statement_valid,omitempty"`
	StockMoveId                           *Many2One   `xmlrpc:"stock_move_id,omitempty" json:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty" json:"stock_valuation_layer_ids,omitempty"`
	StringToHash                          *String     `xmlrpc:"string_to_hash,omitempty" json:"string_to_hash,omitempty"`
	SuitableJournalIds                    *Relation   `xmlrpc:"suitable_journal_ids,omitempty" json:"suitable_journal_ids,omitempty"`
	SuspenseStatementLineId               *Many2One   `xmlrpc:"suspense_statement_line_id,omitempty" json:"suspense_statement_line_id,omitempty"`
	TaxCalculationRoundingMethod          *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty" json:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation   `xmlrpc:"tax_cash_basis_created_move_ids,omitempty" json:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One   `xmlrpc:"tax_cash_basis_origin_move_id,omitempty" json:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One   `xmlrpc:"tax_cash_basis_rec_id,omitempty" json:"tax_cash_basis_rec_id,omitempty"`
	TaxClosingAlert                       *Bool       `xmlrpc:"tax_closing_alert,omitempty" json:"tax_closing_alert,omitempty"`
	TaxClosingEndDate                     *Time       `xmlrpc:"tax_closing_end_date,omitempty" json:"tax_closing_end_date,omitempty"`
	TaxClosingShowMultiClosingWarning     *Bool       `xmlrpc:"tax_closing_show_multi_closing_warning,omitempty" json:"tax_closing_show_multi_closing_warning,omitempty"`
	TaxCountryCode                        *String     `xmlrpc:"tax_country_code,omitempty" json:"tax_country_code,omitempty"`
	TaxCountryId                          *Many2One   `xmlrpc:"tax_country_id,omitempty" json:"tax_country_id,omitempty"`
	TaxLockDateMessage                    *String     `xmlrpc:"tax_lock_date_message,omitempty" json:"tax_lock_date_message,omitempty"`
	TaxReportControlError                 *Bool       `xmlrpc:"tax_report_control_error,omitempty" json:"tax_report_control_error,omitempty"`
	TaxTotals                             *String     `xmlrpc:"tax_totals,omitempty" json:"tax_totals,omitempty"`
	TeamId                                *Many2One   `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	TimesheetCount                        *Int        `xmlrpc:"timesheet_count,omitempty" json:"timesheet_count,omitempty"`
	TimesheetEncodeUomId                  *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty" json:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                          *Relation   `xmlrpc:"timesheet_ids,omitempty" json:"timesheet_ids,omitempty"`
	TimesheetTotalDuration                *Int        `xmlrpc:"timesheet_total_duration,omitempty" json:"timesheet_total_duration,omitempty"`
	ToCheck                               *Bool       `xmlrpc:"to_check,omitempty" json:"to_check,omitempty"`
	TransactionDetails                    interface{} `xmlrpc:"transaction_details,omitempty" json:"transaction_details,omitempty"`
	TransactionIds                        *Relation   `xmlrpc:"transaction_ids,omitempty" json:"transaction_ids,omitempty"`
	TransactionType                       *String     `xmlrpc:"transaction_type,omitempty" json:"transaction_type,omitempty"`
	TransferModelId                       *Many2One   `xmlrpc:"transfer_model_id,omitempty" json:"transfer_model_id,omitempty"`
	TypeName                              *String     `xmlrpc:"type_name,omitempty" json:"type_name,omitempty"`
	UblCiiXmlFile                         *String     `xmlrpc:"ubl_cii_xml_file,omitempty" json:"ubl_cii_xml_file,omitempty"`
	UblCiiXmlId                           *Many2One   `xmlrpc:"ubl_cii_xml_id,omitempty" json:"ubl_cii_xml_id,omitempty"`
	UniqueImportId                        *String     `xmlrpc:"unique_import_id,omitempty" json:"unique_import_id,omitempty"`
	UserId                                *Many2One   `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// AccountBankStatementLines represents array of account.bank.statement.line model.
type AccountBankStatementLines []AccountBankStatementLine

// AccountBankStatementLineModel is the odoo model name.
const AccountBankStatementLineModel = "account.bank.statement.line"

// Many2One convert AccountBankStatementLine to *Many2One.
func (absl *AccountBankStatementLine) Many2One() *Many2One {
	return NewMany2One(absl.Id.Get(), "")
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLine(absl *AccountBankStatementLine) (int64, error) {
	ids, err := c.CreateAccountBankStatementLines([]*AccountBankStatementLine{absl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLines(absls []*AccountBankStatementLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range absls {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementLineModel, vv, nil)
}

// UpdateAccountBankStatementLine updates an existing account.bank.statement.line record.
func (c *Client) UpdateAccountBankStatementLine(absl *AccountBankStatementLine) error {
	return c.UpdateAccountBankStatementLines([]int64{absl.Id.Get()}, absl)
}

// UpdateAccountBankStatementLines updates existing account.bank.statement.line records.
// All records (represented by ids) will be updated by absl values.
func (c *Client) UpdateAccountBankStatementLines(ids []int64, absl *AccountBankStatementLine) error {
	return c.Update(AccountBankStatementLineModel, ids, absl, nil)
}

// DeleteAccountBankStatementLine deletes an existing account.bank.statement.line record.
func (c *Client) DeleteAccountBankStatementLine(id int64) error {
	return c.DeleteAccountBankStatementLines([]int64{id})
}

// DeleteAccountBankStatementLines deletes existing account.bank.statement.line records.
func (c *Client) DeleteAccountBankStatementLines(ids []int64) error {
	return c.Delete(AccountBankStatementLineModel, ids)
}

// GetAccountBankStatementLine gets account.bank.statement.line existing record.
func (c *Client) GetAccountBankStatementLine(id int64) (*AccountBankStatementLine, error) {
	absls, err := c.GetAccountBankStatementLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*absls)[0]), nil
}

// GetAccountBankStatementLines gets account.bank.statement.line existing records.
func (c *Client) GetAccountBankStatementLines(ids []int64) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.Read(AccountBankStatementLineModel, ids, nil, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLine finds account.bank.statement.line record by querying it with criteria.
func (c *Client) FindAccountBankStatementLine(criteria *Criteria) (*AccountBankStatementLine, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, NewOptions().Limit(1), absls); err != nil {
		return nil, err
	}
	return &((*absls)[0]), nil
}

// FindAccountBankStatementLines finds account.bank.statement.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLines(criteria *Criteria, options *Options) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, options, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankStatementLineModel, criteria, options)
}

// FindAccountBankStatementLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
