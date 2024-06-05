package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountBankReconciliationStart           *Time      `xmlrpc:"account_bank_reconciliation_start,omptempty"`
	AccountDashboardOnboardingState          *Selection `xmlrpc:"account_dashboard_onboarding_state,omptempty"`
	AccountDefaultPosReceivableAccountId     *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omptempty"`
	AccountFolder                            *Many2One  `xmlrpc:"account_folder,omptempty"`
	AccountInvoiceOnboardingState            *Selection `xmlrpc:"account_invoice_onboarding_state,omptempty"`
	AccountNo                                *String    `xmlrpc:"account_no,omptempty"`
	AccountOnboardingInvoiceLayoutState      *Selection `xmlrpc:"account_onboarding_invoice_layout_state,omptempty"`
	AccountOnboardingSaleTaxState            *Selection `xmlrpc:"account_onboarding_sale_tax_state,omptempty"`
	AccountOnboardingSampleInvoiceState      *Selection `xmlrpc:"account_onboarding_sample_invoice_state,omptempty"`
	AccountOpeningDate                       *Time      `xmlrpc:"account_opening_date,omptempty"`
	AccountOpeningJournalId                  *Many2One  `xmlrpc:"account_opening_journal_id,omptempty"`
	AccountOpeningMoveId                     *Many2One  `xmlrpc:"account_opening_move_id,omptempty"`
	AccountPurchaseTaxId                     *Many2One  `xmlrpc:"account_purchase_tax_id,omptempty"`
	AccountSaleTaxId                         *Many2One  `xmlrpc:"account_sale_tax_id,omptempty"`
	AccountSetupBankDataState                *Selection `xmlrpc:"account_setup_bank_data_state,omptempty"`
	AccountSetupCoaState                     *Selection `xmlrpc:"account_setup_coa_state,omptempty"`
	AccountSetupFyDataState                  *Selection `xmlrpc:"account_setup_fy_data_state,omptempty"`
	AccountTaxNextActivityType               *Many2One  `xmlrpc:"account_tax_next_activity_type,omptempty"`
	AccountTaxOriginalPeriodicityReminderDay *Int       `xmlrpc:"account_tax_original_periodicity_reminder_day,omptempty"`
	AccountTaxPeriodicity                    *Selection `xmlrpc:"account_tax_periodicity,omptempty"`
	AccountTaxPeriodicityJournalId           *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omptempty"`
	AccountTaxPeriodicityReminderDay         *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omptempty"`
	AccrualDefaultJournalId                  *Many2One  `xmlrpc:"accrual_default_journal_id,omptempty"`
	AngloSaxonAccounting                     *Bool      `xmlrpc:"anglo_saxon_accounting,omptempty"`
	BankAccountCodePrefix                    *String    `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankIds                                  *Relation  `xmlrpc:"bank_ids,omptempty"`
	BankJournalIds                           *Relation  `xmlrpc:"bank_journal_ids,omptempty"`
	BaseOnboardingCompanyState               *Selection `xmlrpc:"base_onboarding_company_state,omptempty"`
	CashAccountCodePrefix                    *String    `xmlrpc:"cash_account_code_prefix,omptempty"`
	Catchall                                 *String    `xmlrpc:"catchall,omptempty"`
	ChartTemplateId                          *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	ChiefExecutiveOfficer                    *String    `xmlrpc:"chief_executive_officer,omptempty"`
	ChildIds                                 *Relation  `xmlrpc:"child_ids,omptempty"`
	City                                     *String    `xmlrpc:"city,omptempty"`
	CompanyRegistry                          *String    `xmlrpc:"company_registry,omptempty"`
	CountryId                                *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId                *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                               *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyIntervalUnit                     *Selection `xmlrpc:"currency_interval_unit,omptempty"`
	CurrencyNextExecutionDate                *Time      `xmlrpc:"currency_next_execution_date,omptempty"`
	CurrencyProvider                         *Selection `xmlrpc:"currency_provider,omptempty"`
	DefaultCashDifferenceExpenseAccountId    *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omptempty"`
	DefaultCashDifferenceIncomeAccountId     *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omptempty"`
	DisplayName                              *String    `xmlrpc:"display_name,omptempty"`
	DocumentsAccountSettings                 *Bool      `xmlrpc:"documents_account_settings,omptempty"`
	DocumentsHrContractsTags                 *Relation  `xmlrpc:"documents_hr_contracts_tags,omptempty"`
	DocumentsHrFolder                        *Many2One  `xmlrpc:"documents_hr_folder,omptempty"`
	DocumentsHrSettings                      *Bool      `xmlrpc:"documents_hr_settings,omptempty"`
	DocumentsProductSettings                 *Bool      `xmlrpc:"documents_product_settings,omptempty"`
	DocumentsProjectSettings                 *Bool      `xmlrpc:"documents_project_settings,omptempty"`
	DocumentsRecruitmentSettings             *Bool      `xmlrpc:"documents_recruitment_settings,omptempty"`
	Email                                    *String    `xmlrpc:"email,omptempty"`
	ExpectsChartOfAccounts                   *Bool      `xmlrpc:"expects_chart_of_accounts,omptempty"`
	ExpenseAccrualAccountId                  *Many2One  `xmlrpc:"expense_accrual_account_id,omptempty"`
	ExpenseCurrencyExchangeAccountId         *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExternalReportLayoutId                   *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	ExtraDay                                 *Float     `xmlrpc:"extra_day,omptempty"`
	ExtraHour                                *Float     `xmlrpc:"extra_hour,omptempty"`
	ExtraProduct                             *Many2One  `xmlrpc:"extra_product,omptempty"`
	Favicon                                  *String    `xmlrpc:"favicon,omptempty"`
	FiscalyearLastDay                        *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth                      *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                       *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Font                                     *Selection `xmlrpc:"font,omptempty"`
	GainAccountId                            *Many2One  `xmlrpc:"gain_account_id,omptempty"`
	HasReceivedWarningStockSms               *Bool      `xmlrpc:"has_received_warning_stock_sms,omptempty"`
	HrPresenceControlEmailAmount             *Int       `xmlrpc:"hr_presence_control_email_amount,omptempty"`
	HrPresenceControlIpList                  *String    `xmlrpc:"hr_presence_control_ip_list,omptempty"`
	Id                                       *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId          *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                               *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InternalTransitLocationId                *Many2One  `xmlrpc:"internal_transit_location_id,omptempty"`
	IntrastatRegionId                        *Many2One  `xmlrpc:"intrastat_region_id,omptempty"`
	IntrastatTransportModeId                 *Many2One  `xmlrpc:"intrastat_transport_mode_id,omptempty"`
	InvoiceIsEmail                           *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                           *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceTerms                             *String    `xmlrpc:"invoice_terms,omptempty"`
	L10NDeDatevClientNumber                  *String    `xmlrpc:"l10n_de_datev_client_number,omptempty"`
	L10NDeDatevConsultantNumber              *String    `xmlrpc:"l10n_de_datev_consultant_number,omptempty"`
	LeaveTimesheetProjectId                  *Many2One  `xmlrpc:"leave_timesheet_project_id,omptempty"`
	LeaveTimesheetTaskId                     *Many2One  `xmlrpc:"leave_timesheet_task_id,omptempty"`
	Logo                                     *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                                  *String    `xmlrpc:"logo_web,omptempty"`
	LossAccountId                            *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	LunchMinimumThreshold                    *Float     `xmlrpc:"lunch_minimum_threshold,omptempty"`
	MinExtraHour                             *Int       `xmlrpc:"min_extra_hour,omptempty"`
	Name                                     *String    `xmlrpc:"name,omptempty"`
	NomenclatureId                           *Many2One  `xmlrpc:"nomenclature_id,omptempty"`
	PaddingTime                              *Float     `xmlrpc:"padding_time,omptempty"`
	PaperformatId                            *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerId                                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentAcquirerOnboardingState           *Selection `xmlrpc:"payment_acquirer_onboarding_state,omptempty"`
	PaymentOnboardingPaymentMethod           *Selection `xmlrpc:"payment_onboarding_payment_method,omptempty"`
	PeriodLockDate                           *Time      `xmlrpc:"period_lock_date,omptempty"`
	Phone                                    *String    `xmlrpc:"phone,omptempty"`
	PlanningAllowSelfUnassign                *Bool      `xmlrpc:"planning_allow_self_unassign,omptempty"`
	PlanningGenerationInterval               *Int       `xmlrpc:"planning_generation_interval,omptempty"`
	PoDoubleValidation                       *Selection `xmlrpc:"po_double_validation,omptempty"`
	PoDoubleValidationAmount                 *Float     `xmlrpc:"po_double_validation_amount,omptempty"`
	PoLead                                   *Float     `xmlrpc:"po_lead,omptempty"`
	PoLock                                   *Selection `xmlrpc:"po_lock,omptempty"`
	PortalConfirmationPay                    *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                   *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PrimaryColor                             *String    `xmlrpc:"primary_color,omptempty"`
	ProductFolder                            *Many2One  `xmlrpc:"product_folder,omptempty"`
	ProductTags                              *Relation  `xmlrpc:"product_tags,omptempty"`
	ProjectFolder                            *Many2One  `xmlrpc:"project_folder,omptempty"`
	ProjectTags                              *Relation  `xmlrpc:"project_tags,omptempty"`
	ProjectTimeModeId                        *Many2One  `xmlrpc:"project_time_mode_id,omptempty"`
	PropertyStockAccountInputCategId         *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId        *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId          *Many2One  `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	QrCode                                   *Bool      `xmlrpc:"qr_code,omptempty"`
	QuotationValidityDays                    *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	RecruitmentFolderId                      *Many2One  `xmlrpc:"recruitment_folder_id,omptempty"`
	RecruitmentTagIds                        *Relation  `xmlrpc:"recruitment_tag_ids,omptempty"`
	RentalLocId                              *Many2One  `xmlrpc:"rental_loc_id,omptempty"`
	ReportFooter                             *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader                             *String    `xmlrpc:"report_header,omptempty"`
	ReportTablePosition                      *Bool      `xmlrpc:"report_table_position,omptempty"`
	ReportTablePositionContinuous            *Bool      `xmlrpc:"report_table_position_continuous,omptempty"`
	ResourceCalendarId                       *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceCalendarIds                      *Relation  `xmlrpc:"resource_calendar_ids,omptempty"`
	RevenueAccrualAccountId                  *Many2One  `xmlrpc:"revenue_accrual_account_id,omptempty"`
	SaleOnboardingOrderConfirmationState     *Selection `xmlrpc:"sale_onboarding_order_confirmation_state,omptempty"`
	SaleOnboardingPaymentMethod              *Selection `xmlrpc:"sale_onboarding_payment_method,omptempty"`
	SaleOnboardingSampleQuotationState       *Selection `xmlrpc:"sale_onboarding_sample_quotation_state,omptempty"`
	SaleQuotationOnboardingState             *Selection `xmlrpc:"sale_quotation_onboarding_state,omptempty"`
	SddCreditorIdentifier                    *String    `xmlrpc:"sdd_creditor_identifier,omptempty"`
	SecondaryColor                           *String    `xmlrpc:"secondary_color,omptempty"`
	SecurityLead                             *Float     `xmlrpc:"security_lead,omptempty"`
	SepaInitiatingPartyName                  *String    `xmlrpc:"sepa_initiating_party_name,omptempty"`
	SepaOrgidId                              *String    `xmlrpc:"sepa_orgid_id,omptempty"`
	SepaOrgidIssr                            *String    `xmlrpc:"sepa_orgid_issr,omptempty"`
	SepaPainVersion                          *Selection `xmlrpc:"sepa_pain_version,omptempty"`
	Sequence                                 *Int       `xmlrpc:"sequence,omptempty"`
	SnailmailColor                           *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                           *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                          *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialFacebook                           *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                             *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                          *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                           *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                            *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                            *String    `xmlrpc:"social_youtube,omptempty"`
	StateId                                  *Many2One  `xmlrpc:"state_id,omptempty"`
	StockMailConfirmationTemplateId          *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omptempty"`
	StockMoveEmailValidation                 *Bool      `xmlrpc:"stock_move_email_validation,omptempty"`
	StockMoveSmsValidation                   *Bool      `xmlrpc:"stock_move_sms_validation,omptempty"`
	StockSmsConfirmationTemplateId           *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omptempty"`
	Street                                   *String    `xmlrpc:"street,omptempty"`
	Street2                                  *String    `xmlrpc:"street2,omptempty"`
	TaxCalculationRoundingMethod             *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                    *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                           *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TaxLockDate                              *Time      `xmlrpc:"tax_lock_date,omptempty"`
	TimesheetEncodeUomId                     *Many2One  `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetMailEmployeeAllow               *Bool      `xmlrpc:"timesheet_mail_employee_allow,omptempty"`
	TimesheetMailEmployeeDelay               *Int       `xmlrpc:"timesheet_mail_employee_delay,omptempty"`
	TimesheetMailEmployeeInterval            *Selection `xmlrpc:"timesheet_mail_employee_interval,omptempty"`
	TimesheetMailEmployeeNextdate            *Time      `xmlrpc:"timesheet_mail_employee_nextdate,omptempty"`
	TimesheetMailManagerAllow                *Bool      `xmlrpc:"timesheet_mail_manager_allow,omptempty"`
	TimesheetMailManagerDelay                *Int       `xmlrpc:"timesheet_mail_manager_delay,omptempty"`
	TimesheetMailManagerInterval             *Selection `xmlrpc:"timesheet_mail_manager_interval,omptempty"`
	TimesheetMailManagerNextdate             *Time      `xmlrpc:"timesheet_mail_manager_nextdate,omptempty"`
	TotalsBelowSections                      *Bool      `xmlrpc:"totals_below_sections,omptempty"`
	TransferAccountCodePrefix                *String    `xmlrpc:"transfer_account_code_prefix,omptempty"`
	TransferAccountId                        *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	UserIds                                  *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                                      *String    `xmlrpc:"vat,omptempty"`
	VatCheckVies                             *Bool      `xmlrpc:"vat_check_vies,omptempty"`
	Website                                  *String    `xmlrpc:"website,omptempty"`
	WebsiteThemeOnboardingDone               *Bool      `xmlrpc:"website_theme_onboarding_done,omptempty"`
	WriteDate                                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                 *Many2One  `xmlrpc:"write_uid,omptempty"`
	YodleeAccessToken                        *String    `xmlrpc:"yodlee_access_token,omptempty"`
	YodleeUserAccessToken                    *String    `xmlrpc:"yodlee_user_access_token,omptempty"`
	YodleeUserLogin                          *String    `xmlrpc:"yodlee_user_login,omptempty"`
	YodleeUserPassword                       *String    `xmlrpc:"yodlee_user_password,omptempty"`
	Zip                                      *String    `xmlrpc:"zip,omptempty"`
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
	return c.Create(ResCompanyModel, vv)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
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
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
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
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.company was not found with criteria %v", criteria)
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
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.company was not found with criteria %v and options %v", criteria, options)
}
