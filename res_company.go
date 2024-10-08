package odoo

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                               *Time      `xmlrpc:"__last_update,omitempty"`
	AccountBankReconciliationStart           *Time      `xmlrpc:"account_bank_reconciliation_start,omitempty"`
	AccountDashboardOnboardingState          *Selection `xmlrpc:"account_dashboard_onboarding_state,omitempty"`
	AccountDefaultPosReceivableAccountId     *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountFolder                            *Many2One  `xmlrpc:"account_folder,omitempty"`
	AccountInvoiceOnboardingState            *Selection `xmlrpc:"account_invoice_onboarding_state,omitempty"`
	AccountNo                                *String    `xmlrpc:"account_no,omitempty"`
	AccountOnboardingInvoiceLayoutState      *Selection `xmlrpc:"account_onboarding_invoice_layout_state,omitempty"`
	AccountOnboardingSaleTaxState            *Selection `xmlrpc:"account_onboarding_sale_tax_state,omitempty"`
	AccountOnboardingSampleInvoiceState      *Selection `xmlrpc:"account_onboarding_sample_invoice_state,omitempty"`
	AccountOpeningDate                       *Time      `xmlrpc:"account_opening_date,omitempty"`
	AccountOpeningJournalId                  *Many2One  `xmlrpc:"account_opening_journal_id,omitempty"`
	AccountOpeningMoveId                     *Many2One  `xmlrpc:"account_opening_move_id,omitempty"`
	AccountPurchaseTaxId                     *Many2One  `xmlrpc:"account_purchase_tax_id,omitempty"`
	AccountSaleTaxId                         *Many2One  `xmlrpc:"account_sale_tax_id,omitempty"`
	AccountSetupBankDataState                *Selection `xmlrpc:"account_setup_bank_data_state,omitempty"`
	AccountSetupCoaState                     *Selection `xmlrpc:"account_setup_coa_state,omitempty"`
	AccountSetupFyDataState                  *Selection `xmlrpc:"account_setup_fy_data_state,omitempty"`
	AccountTaxNextActivityType               *Many2One  `xmlrpc:"account_tax_next_activity_type,omitempty"`
	AccountTaxOriginalPeriodicityReminderDay *Int       `xmlrpc:"account_tax_original_periodicity_reminder_day,omitempty"`
	AccountTaxPeriodicity                    *Selection `xmlrpc:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId           *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay         *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omitempty"`
	AccrualDefaultJournalId                  *Many2One  `xmlrpc:"accrual_default_journal_id,omitempty"`
	AngloSaxonAccounting                     *Bool      `xmlrpc:"anglo_saxon_accounting,omitempty"`
	BankAccountCodePrefix                    *String    `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankIds                                  *Relation  `xmlrpc:"bank_ids,omitempty"`
	BankJournalIds                           *Relation  `xmlrpc:"bank_journal_ids,omitempty"`
	BaseOnboardingCompanyState               *Selection `xmlrpc:"base_onboarding_company_state,omitempty"`
	CashAccountCodePrefix                    *String    `xmlrpc:"cash_account_code_prefix,omitempty"`
	Catchall                                 *String    `xmlrpc:"catchall,omitempty"`
	ChartTemplateId                          *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	ChiefExecutiveOfficer                    *String    `xmlrpc:"chief_executive_officer,omitempty"`
	ChildIds                                 *Relation  `xmlrpc:"child_ids,omitempty"`
	City                                     *String    `xmlrpc:"city,omitempty"`
	CompanyRegistry                          *String    `xmlrpc:"company_registry,omitempty"`
	CountryId                                *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId                *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                               *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyIntervalUnit                     *Selection `xmlrpc:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                *Time      `xmlrpc:"currency_next_execution_date,omitempty"`
	CurrencyProvider                         *Selection `xmlrpc:"currency_provider,omitempty"`
	DefaultCashDifferenceExpenseAccountId    *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId     *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omitempty"`
	DisplayName                              *String    `xmlrpc:"display_name,omitempty"`
	DocumentsAccountSettings                 *Bool      `xmlrpc:"documents_account_settings,omitempty"`
	DocumentsHrContractsTags                 *Relation  `xmlrpc:"documents_hr_contracts_tags,omitempty"`
	DocumentsHrFolder                        *Many2One  `xmlrpc:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                      *Bool      `xmlrpc:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                 *Bool      `xmlrpc:"documents_product_settings,omitempty"`
	DocumentsProjectSettings                 *Bool      `xmlrpc:"documents_project_settings,omitempty"`
	DocumentsRecruitmentSettings             *Bool      `xmlrpc:"documents_recruitment_settings,omitempty"`
	Email                                    *String    `xmlrpc:"email,omitempty"`
	ExpectsChartOfAccounts                   *Bool      `xmlrpc:"expects_chart_of_accounts,omitempty"`
	ExpenseAccrualAccountId                  *Many2One  `xmlrpc:"expense_accrual_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId         *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalReportLayoutId                   *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	ExtraDay                                 *Float     `xmlrpc:"extra_day,omitempty"`
	ExtraHour                                *Float     `xmlrpc:"extra_hour,omitempty"`
	ExtraProduct                             *Many2One  `xmlrpc:"extra_product,omitempty"`
	Favicon                                  *String    `xmlrpc:"favicon,omitempty"`
	FiscalyearLastDay                        *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                      *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                       *Time      `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Font                                     *Selection `xmlrpc:"font,omitempty"`
	GainAccountId                            *Many2One  `xmlrpc:"gain_account_id,omitempty"`
	HasReceivedWarningStockSms               *Bool      `xmlrpc:"has_received_warning_stock_sms,omitempty"`
	HrPresenceControlEmailAmount             *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIpList                  *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	Id                                       *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId          *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                               *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InternalTransitLocationId                *Many2One  `xmlrpc:"internal_transit_location_id,omitempty"`
	IntrastatRegionId                        *Many2One  `xmlrpc:"intrastat_region_id,omitempty"`
	IntrastatTransportModeId                 *Many2One  `xmlrpc:"intrastat_transport_mode_id,omitempty"`
	InvoiceIsEmail                           *Bool      `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsPrint                           *Bool      `xmlrpc:"invoice_is_print,omitempty"`
	InvoiceTerms                             *String    `xmlrpc:"invoice_terms,omitempty"`
	L10NDeDatevClientNumber                  *String    `xmlrpc:"l10n_de_datev_client_number,omitempty"`
	L10NDeDatevConsultantNumber              *String    `xmlrpc:"l10n_de_datev_consultant_number,omitempty"`
	LeaveTimesheetProjectId                  *Many2One  `xmlrpc:"leave_timesheet_project_id,omitempty"`
	LeaveTimesheetTaskId                     *Many2One  `xmlrpc:"leave_timesheet_task_id,omitempty"`
	Logo                                     *String    `xmlrpc:"logo,omitempty"`
	LogoWeb                                  *String    `xmlrpc:"logo_web,omitempty"`
	LossAccountId                            *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	LunchMinimumThreshold                    *Float     `xmlrpc:"lunch_minimum_threshold,omitempty"`
	MinExtraHour                             *Int       `xmlrpc:"min_extra_hour,omitempty"`
	Name                                     *String    `xmlrpc:"name,omitempty"`
	NomenclatureId                           *Many2One  `xmlrpc:"nomenclature_id,omitempty"`
	PaddingTime                              *Float     `xmlrpc:"padding_time,omitempty"`
	PaperformatId                            *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	ParentId                                 *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerId                                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentAcquirerOnboardingState           *Selection `xmlrpc:"payment_acquirer_onboarding_state,omitempty"`
	PaymentOnboardingPaymentMethod           *Selection `xmlrpc:"payment_onboarding_payment_method,omitempty"`
	PeriodLockDate                           *Time      `xmlrpc:"period_lock_date,omitempty"`
	Phone                                    *String    `xmlrpc:"phone,omitempty"`
	PlanningAllowSelfUnassign                *Bool      `xmlrpc:"planning_allow_self_unassign,omitempty"`
	PlanningGenerationInterval               *Int       `xmlrpc:"planning_generation_interval,omitempty"`
	PoDoubleValidation                       *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                 *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                   *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                                   *Selection `xmlrpc:"po_lock,omitempty"`
	PortalConfirmationPay                    *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                   *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PrimaryColor                             *String    `xmlrpc:"primary_color,omitempty"`
	ProductFolder                            *Many2One  `xmlrpc:"product_folder,omitempty"`
	ProductTags                              *Relation  `xmlrpc:"product_tags,omitempty"`
	ProjectFolder                            *Many2One  `xmlrpc:"project_folder,omitempty"`
	ProjectTags                              *Relation  `xmlrpc:"project_tags,omitempty"`
	ProjectTimeModeId                        *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PropertyStockAccountInputCategId         *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId        *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockValuationAccountId          *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	QrCode                                   *Bool      `xmlrpc:"qr_code,omitempty"`
	QuotationValidityDays                    *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	RecruitmentFolderId                      *Many2One  `xmlrpc:"recruitment_folder_id,omitempty"`
	RecruitmentTagIds                        *Relation  `xmlrpc:"recruitment_tag_ids,omitempty"`
	RentalLocId                              *Many2One  `xmlrpc:"rental_loc_id,omitempty"`
	ReportFooter                             *String    `xmlrpc:"report_footer,omitempty"`
	ReportHeader                             *String    `xmlrpc:"report_header,omitempty"`
	ReportTablePosition                      *Bool      `xmlrpc:"report_table_position,omitempty"`
	ReportTablePositionContinuous            *Bool      `xmlrpc:"report_table_position_continuous,omitempty"`
	ResourceCalendarId                       *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceCalendarIds                      *Relation  `xmlrpc:"resource_calendar_ids,omitempty"`
	RevenueAccrualAccountId                  *Many2One  `xmlrpc:"revenue_accrual_account_id,omitempty"`
	SaleOnboardingOrderConfirmationState     *Selection `xmlrpc:"sale_onboarding_order_confirmation_state,omitempty"`
	SaleOnboardingPaymentMethod              *Selection `xmlrpc:"sale_onboarding_payment_method,omitempty"`
	SaleOnboardingSampleQuotationState       *Selection `xmlrpc:"sale_onboarding_sample_quotation_state,omitempty"`
	SaleQuotationOnboardingState             *Selection `xmlrpc:"sale_quotation_onboarding_state,omitempty"`
	SddCreditorIdentifier                    *String    `xmlrpc:"sdd_creditor_identifier,omitempty"`
	SecondaryColor                           *String    `xmlrpc:"secondary_color,omitempty"`
	SecurityLead                             *Float     `xmlrpc:"security_lead,omitempty"`
	SepaInitiatingPartyName                  *String    `xmlrpc:"sepa_initiating_party_name,omitempty"`
	SepaOrgidId                              *String    `xmlrpc:"sepa_orgid_id,omitempty"`
	SepaOrgidIssr                            *String    `xmlrpc:"sepa_orgid_issr,omitempty"`
	SepaPainVersion                          *Selection `xmlrpc:"sepa_pain_version,omitempty"`
	Sequence                                 *Int       `xmlrpc:"sequence,omitempty"`
	SnailmailColor                           *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                           *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                          *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	SocialFacebook                           *String    `xmlrpc:"social_facebook,omitempty"`
	SocialGithub                             *String    `xmlrpc:"social_github,omitempty"`
	SocialInstagram                          *String    `xmlrpc:"social_instagram,omitempty"`
	SocialLinkedin                           *String    `xmlrpc:"social_linkedin,omitempty"`
	SocialTwitter                            *String    `xmlrpc:"social_twitter,omitempty"`
	SocialYoutube                            *String    `xmlrpc:"social_youtube,omitempty"`
	StateId                                  *Many2One  `xmlrpc:"state_id,omitempty"`
	StockMailConfirmationTemplateId          *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omitempty"`
	StockMoveEmailValidation                 *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                   *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId           *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	Street                                   *String    `xmlrpc:"street,omitempty"`
	Street2                                  *String    `xmlrpc:"street2,omitempty"`
	TaxCalculationRoundingMethod             *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                    *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                           *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TaxLockDate                              *Time      `xmlrpc:"tax_lock_date,omitempty"`
	TimesheetEncodeUomId                     *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetMailEmployeeAllow               *Bool      `xmlrpc:"timesheet_mail_employee_allow,omitempty"`
	TimesheetMailEmployeeDelay               *Int       `xmlrpc:"timesheet_mail_employee_delay,omitempty"`
	TimesheetMailEmployeeInterval            *Selection `xmlrpc:"timesheet_mail_employee_interval,omitempty"`
	TimesheetMailEmployeeNextdate            *Time      `xmlrpc:"timesheet_mail_employee_nextdate,omitempty"`
	TimesheetMailManagerAllow                *Bool      `xmlrpc:"timesheet_mail_manager_allow,omitempty"`
	TimesheetMailManagerDelay                *Int       `xmlrpc:"timesheet_mail_manager_delay,omitempty"`
	TimesheetMailManagerInterval             *Selection `xmlrpc:"timesheet_mail_manager_interval,omitempty"`
	TimesheetMailManagerNextdate             *Time      `xmlrpc:"timesheet_mail_manager_nextdate,omitempty"`
	TotalsBelowSections                      *Bool      `xmlrpc:"totals_below_sections,omitempty"`
	TransferAccountCodePrefix                *String    `xmlrpc:"transfer_account_code_prefix,omitempty"`
	TransferAccountId                        *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	UserIds                                  *Relation  `xmlrpc:"user_ids,omitempty"`
	Vat                                      *String    `xmlrpc:"vat,omitempty"`
	VatCheckVies                             *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	Website                                  *String    `xmlrpc:"website,omitempty"`
	WebsiteThemeOnboardingDone               *Bool      `xmlrpc:"website_theme_onboarding_done,omitempty"`
	WriteDate                                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                 *Many2One  `xmlrpc:"write_uid,omitempty"`
	YodleeAccessToken                        *String    `xmlrpc:"yodlee_access_token,omitempty"`
	YodleeUserAccessToken                    *String    `xmlrpc:"yodlee_user_access_token,omitempty"`
	YodleeUserLogin                          *String    `xmlrpc:"yodlee_user_login,omitempty"`
	YodleeUserPassword                       *String    `xmlrpc:"yodlee_user_password,omitempty"`
	Zip                                      *String    `xmlrpc:"zip,omitempty"`
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
