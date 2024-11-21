package odoo

// ResCompany represents res.company model.
type ResCompany struct {
	AccountCashBasisBaseAccountId               *Many2One   `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One   `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One   `xmlrpc:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One   `xmlrpc:"account_discount_income_allocation_id,omitempty"`
	AccountDisplayRepresentativeField           *Bool       `xmlrpc:"account_display_representative_field,omitempty"`
	AccountEnabledTaxCountryIds                 *Relation   `xmlrpc:"account_enabled_tax_country_ids,omitempty"`
	AccountFiscalCountryId                      *Many2One   `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountFolder                               *Many2One   `xmlrpc:"account_folder,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalPaymentCreditAccountId        *Many2One   `xmlrpc:"account_journal_payment_credit_account_id,omitempty"`
	AccountJournalPaymentDebitAccountId         *Many2One   `xmlrpc:"account_journal_payment_debit_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One   `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountOpeningDate                          *Time       `xmlrpc:"account_opening_date,omitempty"`
	AccountOpeningJournalId                     *Many2One   `xmlrpc:"account_opening_journal_id,omitempty"`
	AccountOpeningMoveId                        *Many2One   `xmlrpc:"account_opening_move_id,omitempty"`
	AccountPurchaseTaxId                        *Many2One   `xmlrpc:"account_purchase_tax_id,omitempty"`
	AccountRepresentativeId                     *Many2One   `xmlrpc:"account_representative_id,omitempty"`
	AccountRevaluationExpenseProvisionAccountId *Many2One   `xmlrpc:"account_revaluation_expense_provision_account_id,omitempty"`
	AccountRevaluationIncomeProvisionAccountId  *Many2One   `xmlrpc:"account_revaluation_income_provision_account_id,omitempty"`
	AccountRevaluationJournalId                 *Many2One   `xmlrpc:"account_revaluation_journal_id,omitempty"`
	AccountSaleTaxId                            *Many2One   `xmlrpc:"account_sale_tax_id,omitempty"`
	AccountStorno                               *Bool       `xmlrpc:"account_storno,omitempty"`
	AccountTaxPeriodicity                       *Selection  `xmlrpc:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId              *Many2One   `xmlrpc:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay            *Int        `xmlrpc:"account_tax_periodicity_reminder_day,omitempty"`
	AccountTaxUnitIds                           *Relation   `xmlrpc:"account_tax_unit_ids,omitempty"`
	AccountUseCreditLimit                       *Bool       `xmlrpc:"account_use_credit_limit,omitempty"`
	Active                                      *Bool       `xmlrpc:"active,omitempty"`
	AliasDomainId                               *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AliasDomainName                             *String     `xmlrpc:"alias_domain_name,omitempty"`
	AllChildIds                                 *Relation   `xmlrpc:"all_child_ids,omitempty"`
	AngloSaxonAccounting                        *Bool       `xmlrpc:"anglo_saxon_accounting,omitempty"`
	AnnualInventoryDay                          *Int        `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection  `xmlrpc:"annual_inventory_month,omitempty"`
	AttendanceBarcodeSource                     *Selection  `xmlrpc:"attendance_barcode_source,omitempty"`
	AttendanceFromSystray                       *Bool       `xmlrpc:"attendance_from_systray,omitempty"`
	AttendanceKioskDelay                        *Int        `xmlrpc:"attendance_kiosk_delay,omitempty"`
	AttendanceKioskKey                          *String     `xmlrpc:"attendance_kiosk_key,omitempty"`
	AttendanceKioskMode                         *Selection  `xmlrpc:"attendance_kiosk_mode,omitempty"`
	AttendanceKioskUrl                          *String     `xmlrpc:"attendance_kiosk_url,omitempty"`
	AttendanceKioskUsePin                       *Bool       `xmlrpc:"attendance_kiosk_use_pin,omitempty"`
	AutomaticEntryDefaultJournalId              *Many2One   `xmlrpc:"automatic_entry_default_journal_id,omitempty"`
	BankAccountCodePrefix                       *String     `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankIds                                     *Relation   `xmlrpc:"bank_ids,omitempty"`
	BankJournalIds                              *Relation   `xmlrpc:"bank_journal_ids,omitempty"`
	BillingRateTarget                           *Int        `xmlrpc:"billing_rate_target,omitempty"`
	BounceEmail                                 *String     `xmlrpc:"bounce_email,omitempty"`
	BounceFormatted                             *String     `xmlrpc:"bounce_formatted,omitempty"`
	CashAccountCodePrefix                       *String     `xmlrpc:"cash_account_code_prefix,omitempty"`
	CatchallEmail                               *String     `xmlrpc:"catchall_email,omitempty"`
	CatchallFormatted                           *String     `xmlrpc:"catchall_formatted,omitempty"`
	ChartTemplate                               *Selection  `xmlrpc:"chart_template,omitempty"`
	CheckAccountAuditTrail                      *Bool       `xmlrpc:"check_account_audit_trail,omitempty"`
	ChildIds                                    *Relation   `xmlrpc:"child_ids,omitempty"`
	City                                        *String     `xmlrpc:"city,omitempty"`
	Color                                       *Int        `xmlrpc:"color,omitempty"`
	CompanyDetails                              *String     `xmlrpc:"company_details,omitempty"`
	CompanyExpenseAllowedPaymentMethodLineIds   *Relation   `xmlrpc:"company_expense_allowed_payment_method_line_ids,omitempty"`
	CompanyRegistry                             *String     `xmlrpc:"company_registry,omitempty"`
	CountryCode                                 *String     `xmlrpc:"country_code,omitempty"`
	CountryId                                   *Many2One   `xmlrpc:"country_id,omitempty"`
	CreateDate                                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId                   *Many2One   `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One   `xmlrpc:"currency_id,omitempty"`
	CurrencyIntervalUnit                        *Selection  `xmlrpc:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                   *Time       `xmlrpc:"currency_next_execution_date,omitempty"`
	CurrencyProvider                            *Selection  `xmlrpc:"currency_provider,omitempty"`
	DaysToPurchase                              *Float      `xmlrpc:"days_to_purchase,omitempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One   `xmlrpc:"default_cash_difference_expense_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One   `xmlrpc:"default_cash_difference_income_account_id,omitempty"`
	DefaultFromEmail                            *String     `xmlrpc:"default_from_email,omitempty"`
	DeferredAmountComputationMethod             *Selection  `xmlrpc:"deferred_amount_computation_method,omitempty"`
	DeferredExpenseAccountId                    *Many2One   `xmlrpc:"deferred_expense_account_id,omitempty"`
	DeferredJournalId                           *Many2One   `xmlrpc:"deferred_journal_id,omitempty"`
	DeferredRevenueAccountId                    *Many2One   `xmlrpc:"deferred_revenue_account_id,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool       `xmlrpc:"display_invoice_amount_total_words,omitempty"`
	DisplayName                                 *String     `xmlrpc:"display_name,omitempty"`
	DocumentsAccountSettings                    *Bool       `xmlrpc:"documents_account_settings,omitempty"`
	DocumentsHrFolder                           *Many2One   `xmlrpc:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                         *Bool       `xmlrpc:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                    *Bool       `xmlrpc:"documents_product_settings,omitempty"`
	DocumentsSpreadsheetFolderId                *Many2One   `xmlrpc:"documents_spreadsheet_folder_id,omitempty"`
	Email                                       *String     `xmlrpc:"email,omitempty"`
	EmailFormatted                              *String     `xmlrpc:"email_formatted,omitempty"`
	EmailPrimaryColor                           *String     `xmlrpc:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String     `xmlrpc:"email_secondary_color,omitempty"`
	EmployeePropertiesDefinition                interface{} `xmlrpc:"employee_properties_definition,omitempty"`
	ExpectsChartOfAccounts                      *Bool       `xmlrpc:"expects_chart_of_accounts,omitempty"`
	ExpenseAccrualAccountId                     *Many2One   `xmlrpc:"expense_accrual_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One   `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExpenseExtractShowOcrOptionSelection        *Selection  `xmlrpc:"expense_extract_show_ocr_option_selection,omitempty"`
	ExpenseJournalId                            *Many2One   `xmlrpc:"expense_journal_id,omitempty"`
	ExpenseProductId                            *Many2One   `xmlrpc:"expense_product_id,omitempty"`
	ExternalReportLayoutId                      *Many2One   `xmlrpc:"external_report_layout_id,omitempty"`
	ExtractInInvoiceDigitalizationMode          *Selection  `xmlrpc:"extract_in_invoice_digitalization_mode,omitempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection  `xmlrpc:"extract_out_invoice_digitalization_mode,omitempty"`
	ExtractSingleLinePerTax                     *Bool       `xmlrpc:"extract_single_line_per_tax,omitempty"`
	FiscalPositionIds                           *Relation   `xmlrpc:"fiscal_position_ids,omitempty"`
	FiscalyearLastDay                           *Int        `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection  `xmlrpc:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                          *Time       `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Font                                        *Selection  `xmlrpc:"font,omitempty"`
	GainAccountId                               *Many2One   `xmlrpc:"gain_account_id,omitempty"`
	GenerateDeferredExpenseEntriesMethod        *Selection  `xmlrpc:"generate_deferred_expense_entries_method,omitempty"`
	GenerateDeferredRevenueEntriesMethod        *Selection  `xmlrpc:"generate_deferred_revenue_entries_method,omitempty"`
	HasMessage                                  *Bool       `xmlrpc:"has_message,omitempty"`
	HasReceivedWarningStockSms                  *Bool       `xmlrpc:"has_received_warning_stock_sms,omitempty"`
	HrAttendanceDisplayOvertime                 *Bool       `xmlrpc:"hr_attendance_display_overtime,omitempty"`
	HrAttendanceOvertime                        *Bool       `xmlrpc:"hr_attendance_overtime,omitempty"`
	HrPresenceControlEmailAmount                *Int        `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIpList                     *String     `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	Id                                          *Int        `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One   `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One   `xmlrpc:"incoterm_id,omitempty"`
	InternalProjectId                           *Many2One   `xmlrpc:"internal_project_id,omitempty"`
	InternalTransitLocationId                   *Many2One   `xmlrpc:"internal_transit_location_id,omitempty"`
	InvoiceIsDownload                           *Bool       `xmlrpc:"invoice_is_download,omitempty"`
	InvoiceIsEmail                              *Bool       `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsUblCii                             *Bool       `xmlrpc:"invoice_is_ubl_cii,omitempty"`
	InvoiceTerms                                *String     `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String     `xmlrpc:"invoice_terms_html,omitempty"`
	InvoicingSwitchThreshold                    *Time       `xmlrpc:"invoicing_switch_threshold,omitempty"`
	IsCompanyDetailsEmpty                       *Bool       `xmlrpc:"is_company_details_empty,omitempty"`
	L10NDeDatevAccountLength                    *Int        `xmlrpc:"l10n_de_datev_account_length,omitempty"`
	L10NDeDatevClientNumber                     *String     `xmlrpc:"l10n_de_datev_client_number,omitempty"`
	L10NDeDatevConsultantNumber                 *String     `xmlrpc:"l10n_de_datev_consultant_number,omitempty"`
	L10NDeStnr                                  *String     `xmlrpc:"l10n_de_stnr,omitempty"`
	L10NDeWidnr                                 *String     `xmlrpc:"l10n_de_widnr,omitempty"`
	LayoutBackground                            *Selection  `xmlrpc:"layout_background,omitempty"`
	LayoutBackgroundImage                       *String     `xmlrpc:"layout_background_image,omitempty"`
	LeaveTimesheetTaskId                        *Many2One   `xmlrpc:"leave_timesheet_task_id,omitempty"`
	Logo                                        *String     `xmlrpc:"logo,omitempty"`
	LogoWeb                                     *String     `xmlrpc:"logo_web,omitempty"`
	LossAccountId                               *Many2One   `xmlrpc:"loss_account_id,omitempty"`
	MaxTaxLockDate                              *Time       `xmlrpc:"max_tax_lock_date,omitempty"`
	MessageAttachmentCount                      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter                    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                                      *String     `xmlrpc:"mobile,omitempty"`
	MultiVatForeignCountryIds                   *Relation   `xmlrpc:"multi_vat_foreign_country_ids,omitempty"`
	Name                                        *String     `xmlrpc:"name,omitempty"`
	NomenclatureId                              *Many2One   `xmlrpc:"nomenclature_id,omitempty"`
	OvertimeCompanyThreshold                    *Int        `xmlrpc:"overtime_company_threshold,omitempty"`
	OvertimeEmployeeThreshold                   *Int        `xmlrpc:"overtime_employee_threshold,omitempty"`
	OvertimeStartDate                           *Time       `xmlrpc:"overtime_start_date,omitempty"`
	PaperformatId                               *Many2One   `xmlrpc:"paperformat_id,omitempty"`
	ParentId                                    *Many2One   `xmlrpc:"parent_id,omitempty"`
	ParentIds                                   *Relation   `xmlrpc:"parent_ids,omitempty"`
	ParentPath                                  *String     `xmlrpc:"parent_path,omitempty"`
	PartnerId                                   *Many2One   `xmlrpc:"partner_id,omitempty"`
	PaymentOnboardingPaymentMethod              *Selection  `xmlrpc:"payment_onboarding_payment_method,omitempty"`
	PeriodLockDate                              *Time       `xmlrpc:"period_lock_date,omitempty"`
	Phone                                       *String     `xmlrpc:"phone,omitempty"`
	PoDoubleValidation                          *Selection  `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float      `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float      `xmlrpc:"po_lead,omitempty"`
	PoLock                                      *Selection  `xmlrpc:"po_lock,omitempty"`
	PortalConfirmationPay                       *Bool       `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool       `xmlrpc:"portal_confirmation_sign,omitempty"`
	PredictBillProduct                          *Bool       `xmlrpc:"predict_bill_product,omitempty"`
	PrepaymentPercent                           *Float      `xmlrpc:"prepayment_percent,omitempty"`
	PrimaryColor                                *String     `xmlrpc:"primary_color,omitempty"`
	ProductFolder                               *Many2One   `xmlrpc:"product_folder,omitempty"`
	ProductTags                                 *Relation   `xmlrpc:"product_tags,omitempty"`
	ProjectTimeModeId                           *Many2One   `xmlrpc:"project_time_mode_id,omitempty"`
	QrCode                                      *Bool       `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection  `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int        `xmlrpc:"quotation_validity_days,omitempty"`
	RatingIds                                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	ReportFooter                                *String     `xmlrpc:"report_footer,omitempty"`
	ReportHeader                                *String     `xmlrpc:"report_header,omitempty"`
	ResourceCalendarId                          *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceCalendarIds                         *Relation   `xmlrpc:"resource_calendar_ids,omitempty"`
	RevenueAccrualAccountId                     *Many2One   `xmlrpc:"revenue_accrual_account_id,omitempty"`
	RootId                                      *Many2One   `xmlrpc:"root_id,omitempty"`
	SaleDiscountProductId                       *Many2One   `xmlrpc:"sale_discount_product_id,omitempty"`
	SaleDownPaymentProductId                    *Many2One   `xmlrpc:"sale_down_payment_product_id,omitempty"`
	SaleFooter                                  *String     `xmlrpc:"sale_footer,omitempty"`
	SaleFooterName                              *String     `xmlrpc:"sale_footer_name,omitempty"`
	SaleHeader                                  *String     `xmlrpc:"sale_header,omitempty"`
	SaleHeaderName                              *String     `xmlrpc:"sale_header_name,omitempty"`
	SaleOnboardingPaymentMethod                 *Selection  `xmlrpc:"sale_onboarding_payment_method,omitempty"`
	SaleOrderTemplateId                         *Many2One   `xmlrpc:"sale_order_template_id,omitempty"`
	SecondaryColor                              *String     `xmlrpc:"secondary_color,omitempty"`
	SecurityLead                                *Float      `xmlrpc:"security_lead,omitempty"`
	Sequence                                    *Int        `xmlrpc:"sequence,omitempty"`
	SignTerms                                   *String     `xmlrpc:"sign_terms,omitempty"`
	SignTermsHtml                               *String     `xmlrpc:"sign_terms_html,omitempty"`
	SignTermsType                               *Selection  `xmlrpc:"sign_terms_type,omitempty"`
	SnailmailColor                              *Bool       `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool       `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                             *Bool       `xmlrpc:"snailmail_duplex,omitempty"`
	StateId                                     *Many2One   `xmlrpc:"state_id,omitempty"`
	StockMailConfirmationTemplateId             *Many2One   `xmlrpc:"stock_mail_confirmation_template_id,omitempty"`
	StockMoveEmailValidation                    *Bool       `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool       `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One   `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	Street                                      *String     `xmlrpc:"street,omitempty"`
	Street2                                     *String     `xmlrpc:"street2,omitempty"`
	TaxCalculationRoundingMethod                *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One   `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool       `xmlrpc:"tax_exigibility,omitempty"`
	TaxLockDate                                 *Time       `xmlrpc:"tax_lock_date,omitempty"`
	TermsType                                   *Selection  `xmlrpc:"terms_type,omitempty"`
	TimesheetEncodeUomId                        *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetMailAllow                          *Bool       `xmlrpc:"timesheet_mail_allow,omitempty"`
	TimesheetMailDelay                          *Int        `xmlrpc:"timesheet_mail_delay,omitempty"`
	TimesheetMailEmployeeAllow                  *Bool       `xmlrpc:"timesheet_mail_employee_allow,omitempty"`
	TimesheetMailEmployeeDelay                  *Int        `xmlrpc:"timesheet_mail_employee_delay,omitempty"`
	TimesheetMailEmployeeInterval               *Selection  `xmlrpc:"timesheet_mail_employee_interval,omitempty"`
	TimesheetMailEmployeeNextdate               *Time       `xmlrpc:"timesheet_mail_employee_nextdate,omitempty"`
	TimesheetMailInterval                       *Selection  `xmlrpc:"timesheet_mail_interval,omitempty"`
	TimesheetMailNextdate                       *Time       `xmlrpc:"timesheet_mail_nextdate,omitempty"`
	TotalsBelowSections                         *Bool       `xmlrpc:"totals_below_sections,omitempty"`
	TransferAccountCodePrefix                   *String     `xmlrpc:"transfer_account_code_prefix,omitempty"`
	TransferAccountId                           *Many2One   `xmlrpc:"transfer_account_id,omitempty"`
	UserIds                                     *Relation   `xmlrpc:"user_ids,omitempty"`
	UsesDefaultLogo                             *Bool       `xmlrpc:"uses_default_logo,omitempty"`
	Vat                                         *String     `xmlrpc:"vat,omitempty"`
	VatCheckVies                                *Bool       `xmlrpc:"vat_check_vies,omitempty"`
	Website                                     *String     `xmlrpc:"website,omitempty"`
	WebsiteMessageIds                           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One   `xmlrpc:"write_uid,omitempty"`
	Zip                                         *String     `xmlrpc:"zip,omitempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	ids, err := c.CreateResCompanys([]*ResCompany{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv, nil)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc, nil)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCompanyModel, criteria, options)
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
