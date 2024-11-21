package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty" json:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omitempty" json:"account_default_credit_limit,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One  `xmlrpc:"account_discount_expense_allocation_id,omitempty" json:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One  `xmlrpc:"account_discount_income_allocation_id,omitempty" json:"account_discount_income_allocation_id,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty" json:"account_fiscal_country_id,omitempty"`
	AccountFolder                               *Many2One  `xmlrpc:"account_folder,omitempty" json:"account_folder,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty" json:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty" json:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omitempty" json:"account_journal_payment_credit_account_id,omitempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omitempty" json:"account_journal_payment_debit_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty" json:"account_journal_suspense_account_id,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty" json:"account_storno,omitempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omitempty" json:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omitempty" json:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omitempty" json:"account_tax_periodicity_reminder_day,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty" json:"account_use_credit_limit,omitempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omitempty" json:"active_user_count,omitempty"`
	AliasDomainId                               *Many2One  `xmlrpc:"alias_domain_id,omitempty" json:"alias_domain_id,omitempty"`
	AnalyticPlanId                              *Many2One  `xmlrpc:"analytic_plan_id,omitempty" json:"analytic_plan_id,omitempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omitempty" json:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omitempty" json:"annual_inventory_month,omitempty"`
	AttendanceBarcodeSource                     *Selection `xmlrpc:"attendance_barcode_source,omitempty" json:"attendance_barcode_source,omitempty"`
	AttendanceFromSystray                       *Bool      `xmlrpc:"attendance_from_systray,omitempty" json:"attendance_from_systray,omitempty"`
	AttendanceKioskDelay                        *Int       `xmlrpc:"attendance_kiosk_delay,omitempty" json:"attendance_kiosk_delay,omitempty"`
	AttendanceKioskMode                         *Selection `xmlrpc:"attendance_kiosk_mode,omitempty" json:"attendance_kiosk_mode,omitempty"`
	AttendanceKioskUrl                          *String    `xmlrpc:"attendance_kiosk_url,omitempty" json:"attendance_kiosk_url,omitempty"`
	AttendanceKioskUsePin                       *Bool      `xmlrpc:"attendance_kiosk_use_pin,omitempty" json:"attendance_kiosk_use_pin,omitempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omitempty" json:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty" json:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omitempty" json:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omitempty" json:"automatic_invoice,omitempty"`
	BillingRateTarget                           *Int       `xmlrpc:"billing_rate_target,omitempty" json:"billing_rate_target,omitempty"`
	ChartTemplate                               *Selection `xmlrpc:"chart_template,omitempty" json:"chart_template,omitempty"`
	CheckAccountAuditTrail                      *Bool      `xmlrpc:"check_account_audit_trail,omitempty" json:"check_account_audit_trail,omitempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omitempty" json:"company_count,omitempty"`
	CompanyCountryCode                          *String    `xmlrpc:"company_country_code,omitempty" json:"company_country_code,omitempty"`
	CompanyCurrencyId                           *Many2One  `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyExpenseAllowedPaymentMethodLineIds   *Relation  `xmlrpc:"company_expense_allowed_payment_method_line_ids,omitempty" json:"company_expense_allowed_payment_method_line_ids,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omitempty" json:"company_informations,omitempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omitempty" json:"company_name,omitempty"`
	CompanySoTemplateId                         *Many2One  `xmlrpc:"company_so_template_id,omitempty" json:"company_so_template_id,omitempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CrmAutoAssignmentAction                     *Selection `xmlrpc:"crm_auto_assignment_action,omitempty" json:"crm_auto_assignment_action,omitempty"`
	CrmAutoAssignmentIntervalNumber             *Int       `xmlrpc:"crm_auto_assignment_interval_number,omitempty" json:"crm_auto_assignment_interval_number,omitempty"`
	CrmAutoAssignmentIntervalType               *Selection `xmlrpc:"crm_auto_assignment_interval_type,omitempty" json:"crm_auto_assignment_interval_type,omitempty"`
	CrmAutoAssignmentRunDatetime                *Time      `xmlrpc:"crm_auto_assignment_run_datetime,omitempty" json:"crm_auto_assignment_run_datetime,omitempty"`
	CrmUseAutoAssignment                        *Bool      `xmlrpc:"crm_use_auto_assignment,omitempty" json:"crm_use_auto_assignment,omitempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty" json:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omitempty" json:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omitempty" json:"currency_next_execution_date,omitempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omitempty" json:"currency_provider,omitempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omitempty" json:"days_to_purchase,omitempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omitempty" json:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                        *Selection `xmlrpc:"default_picking_policy,omitempty" json:"default_picking_policy,omitempty"`
	DefaultPurchaseMethod                       *Selection `xmlrpc:"default_purchase_method,omitempty" json:"default_purchase_method,omitempty"`
	DeferredAmountComputationMethod             *Selection `xmlrpc:"deferred_amount_computation_method,omitempty" json:"deferred_amount_computation_method,omitempty"`
	DeferredExpenseAccountId                    *Many2One  `xmlrpc:"deferred_expense_account_id,omitempty" json:"deferred_expense_account_id,omitempty"`
	DeferredJournalId                           *Many2One  `xmlrpc:"deferred_journal_id,omitempty" json:"deferred_journal_id,omitempty"`
	DeferredRevenueAccountId                    *Many2One  `xmlrpc:"deferred_revenue_account_id,omitempty" json:"deferred_revenue_account_id,omitempty"`
	DeletionDelay                               *Int       `xmlrpc:"deletion_delay,omitempty" json:"deletion_delay,omitempty"`
	DepositDefaultProductId                     *Many2One  `xmlrpc:"deposit_default_product_id,omitempty" json:"deposit_default_product_id,omitempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omitempty" json:"digest_emails,omitempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omitempty" json:"digest_id,omitempty"`
	DisableRedirectFirebaseDynamicLink          *Bool      `xmlrpc:"disable_redirect_firebase_dynamic_link,omitempty" json:"disable_redirect_firebase_dynamic_link,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool      `xmlrpc:"display_invoice_amount_total_words,omitempty" json:"display_invoice_amount_total_words,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentsAccountSettings                    *Bool      `xmlrpc:"documents_account_settings,omitempty" json:"documents_account_settings,omitempty"`
	DocumentsHrFolder                           *Many2One  `xmlrpc:"documents_hr_folder,omitempty" json:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                         *Bool      `xmlrpc:"documents_hr_settings,omitempty" json:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                    *Bool      `xmlrpc:"documents_product_settings,omitempty" json:"documents_product_settings,omitempty"`
	DocumentsSpreadsheetFolderId                *Many2One  `xmlrpc:"documents_spreadsheet_folder_id,omitempty" json:"documents_spreadsheet_folder_id,omitempty"`
	EmailPrimaryColor                           *String    `xmlrpc:"email_primary_color,omitempty" json:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String    `xmlrpc:"email_secondary_color,omitempty" json:"email_secondary_color,omitempty"`
	EnableOcn                                   *Bool      `xmlrpc:"enable_ocn,omitempty" json:"enable_ocn,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty" json:"expense_currency_exchange_account_id,omitempty"`
	ExpenseExtractShowOcrOptionSelection        *Selection `xmlrpc:"expense_extract_show_ocr_option_selection,omitempty" json:"expense_extract_show_ocr_option_selection,omitempty"`
	ExpenseJournalId                            *Many2One  `xmlrpc:"expense_journal_id,omitempty" json:"expense_journal_id,omitempty"`
	ExpenseProductId                            *Many2One  `xmlrpc:"expense_product_id,omitempty" json:"expense_product_id,omitempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omitempty" json:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty" json:"external_report_layout_id,omitempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omitempty" json:"extract_in_invoice_digitalization_mode,omitempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omitempty" json:"extract_out_invoice_digitalization_mode,omitempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omitempty" json:"extract_single_line_per_tax,omitempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omitempty" json:"fail_counter,omitempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omitempty" json:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omitempty" json:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omitempty" json:"fiscalyear_lock_date,omitempty"`
	GenerateDeferredExpenseEntriesMethod        *Selection `xmlrpc:"generate_deferred_expense_entries_method,omitempty" json:"generate_deferred_expense_entries_method,omitempty"`
	GenerateDeferredRevenueEntriesMethod        *Selection `xmlrpc:"generate_deferred_revenue_entries_method,omitempty" json:"generate_deferred_revenue_entries_method,omitempty"`
	GoogleGmailClientIdentifier                 *String    `xmlrpc:"google_gmail_client_identifier,omitempty" json:"google_gmail_client_identifier,omitempty"`
	GoogleGmailClientSecret                     *String    `xmlrpc:"google_gmail_client_secret,omitempty" json:"google_gmail_client_secret,omitempty"`
	GoogleTranslateApiKey                       *String    `xmlrpc:"google_translate_api_key,omitempty" json:"google_translate_api_key,omitempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omitempty" json:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omitempty" json:"group_auto_done_setting,omitempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omitempty" json:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omitempty" json:"group_discount_per_so_line,omitempty"`
	GroupFiscalYear                             *Bool      `xmlrpc:"group_fiscal_year,omitempty" json:"group_fiscal_year,omitempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omitempty" json:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omitempty" json:"group_lot_on_invoice,omitempty"`
	GroupManageTemplateAccess                   *Bool      `xmlrpc:"group_manage_template_access,omitempty" json:"group_manage_template_access,omitempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omitempty" json:"group_multi_currency,omitempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omitempty" json:"group_product_pricelist,omitempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omitempty" json:"group_product_variant,omitempty"`
	GroupProformaSales                          *Bool      `xmlrpc:"group_proforma_sales,omitempty" json:"group_proforma_sales,omitempty"`
	GroupProjectMilestone                       *Bool      `xmlrpc:"group_project_milestone,omitempty" json:"group_project_milestone,omitempty"`
	GroupProjectRating                          *Bool      `xmlrpc:"group_project_rating,omitempty" json:"group_project_rating,omitempty"`
	GroupProjectRecurringTasks                  *Bool      `xmlrpc:"group_project_recurring_tasks,omitempty" json:"group_project_recurring_tasks,omitempty"`
	GroupProjectStages                          *Bool      `xmlrpc:"group_project_stages,omitempty" json:"group_project_stages,omitempty"`
	GroupProjectTaskDependencies                *Bool      `xmlrpc:"group_project_task_dependencies,omitempty" json:"group_project_task_dependencies,omitempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omitempty" json:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                      *Bool      `xmlrpc:"group_sale_order_template,omitempty" json:"group_sale_order_template,omitempty"`
	GroupSalePricelist                          *Bool      `xmlrpc:"group_sale_pricelist,omitempty" json:"group_sale_pricelist,omitempty"`
	GroupSendReminder                           *Bool      `xmlrpc:"group_send_reminder,omitempty" json:"group_send_reminder,omitempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty" json:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omitempty" json:"group_show_sale_receipts,omitempty"`
	GroupStockAccountingAutomatic               *Bool      `xmlrpc:"group_stock_accounting_automatic,omitempty" json:"group_stock_accounting_automatic,omitempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omitempty" json:"group_stock_adv_location,omitempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omitempty" json:"group_stock_lot_print_gs1,omitempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omitempty" json:"group_stock_multi_locations,omitempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omitempty" json:"group_stock_packaging,omitempty"`
	GroupStockPickingWave                       *Bool      `xmlrpc:"group_stock_picking_wave,omitempty" json:"group_stock_picking_wave,omitempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omitempty" json:"group_stock_production_lot,omitempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omitempty" json:"group_stock_reception_report,omitempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omitempty" json:"group_stock_sign_delivery,omitempty"`
	GroupStockStorageCategories                 *Bool      `xmlrpc:"group_stock_storage_categories,omitempty" json:"group_stock_storage_categories,omitempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty" json:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty" json:"group_stock_tracking_owner,omitempty"`
	GroupTimesheetLeaderboardShowRates          *Bool      `xmlrpc:"group_timesheet_leaderboard_show_rates,omitempty" json:"group_timesheet_leaderboard_show_rates,omitempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omitempty" json:"group_uom,omitempty"`
	GroupUseLead                                *Bool      `xmlrpc:"group_use_lead,omitempty" json:"group_use_lead,omitempty"`
	GroupUseRecurringRevenues                   *Bool      `xmlrpc:"group_use_recurring_revenues,omitempty" json:"group_use_recurring_revenues,omitempty"`
	GroupUseTimesheetLeaderboard                *Bool      `xmlrpc:"group_use_timesheet_leaderboard,omitempty" json:"group_use_timesheet_leaderboard,omitempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omitempty" json:"group_warning_account,omitempty"`
	GroupWarningPurchase                        *Bool      `xmlrpc:"group_warning_purchase,omitempty" json:"group_warning_purchase,omitempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omitempty" json:"group_warning_sale,omitempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omitempty" json:"group_warning_stock,omitempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omitempty" json:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omitempty" json:"has_chart_of_accounts,omitempty"`
	HrAttendanceDisplayOvertime                 *Bool      `xmlrpc:"hr_attendance_display_overtime,omitempty" json:"hr_attendance_display_overtime,omitempty"`
	HrAttendanceOvertime                        *Bool      `xmlrpc:"hr_attendance_overtime,omitempty" json:"hr_attendance_overtime,omitempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omitempty" json:"hr_employee_self_edit,omitempty"`
	HrExpenseAliasPrefix                        *String    `xmlrpc:"hr_expense_alias_prefix,omitempty" json:"hr_expense_alias_prefix,omitempty"`
	HrExpenseUseMailgateway                     *Bool      `xmlrpc:"hr_expense_use_mailgateway,omitempty" json:"hr_expense_use_mailgateway,omitempty"`
	HrPresenceControlEmail                      *Bool      `xmlrpc:"hr_presence_control_email,omitempty" json:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty" json:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                         *Bool      `xmlrpc:"hr_presence_control_ip,omitempty" json:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omitempty" json:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                      *Bool      `xmlrpc:"hr_presence_control_login,omitempty" json:"hr_presence_control_login,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty" json:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omitempty" json:"incoterm_id,omitempty"`
	InternalProjectId                           *Many2One  `xmlrpc:"internal_project_id,omitempty" json:"internal_project_id,omitempty"`
	InvoiceIsDownload                           *Bool      `xmlrpc:"invoice_is_download,omitempty" json:"invoice_is_download,omitempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omitempty" json:"invoice_is_email,omitempty"`
	InvoiceIsUblCii                             *Bool      `xmlrpc:"invoice_is_ubl_cii,omitempty" json:"invoice_is_ubl_cii,omitempty"`
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty" json:"invoice_mail_template_id,omitempty"`
	InvoicePolicy                               *Bool      `xmlrpc:"invoice_policy,omitempty" json:"invoice_policy,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty" json:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty" json:"invoice_terms_html,omitempty"`
	InvoicedTimesheet                           *Selection `xmlrpc:"invoiced_timesheet,omitempty" json:"invoiced_timesheet,omitempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omitempty" json:"invoicing_switch_threshold,omitempty"`
	IsAccountPeppolEligible                     *Bool      `xmlrpc:"is_account_peppol_eligible,omitempty" json:"is_account_peppol_eligible,omitempty"`
	IsEncodeUomDays                             *Bool      `xmlrpc:"is_encode_uom_days,omitempty" json:"is_encode_uom_days,omitempty"`
	IsInstalledSale                             *Bool      `xmlrpc:"is_installed_sale,omitempty" json:"is_installed_sale,omitempty"`
	IsMembershipMulti                           *Bool      `xmlrpc:"is_membership_multi,omitempty" json:"is_membership_multi,omitempty"`
	IsRootCompany                               *Bool      `xmlrpc:"is_root_company,omitempty" json:"is_root_company,omitempty"`
	L10NDeDatevAccountLength                    *Int       `xmlrpc:"l10n_de_datev_account_length,omitempty" json:"l10n_de_datev_account_length,omitempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omitempty" json:"language_count,omitempty"`
	LeadEnrichAuto                              *Selection `xmlrpc:"lead_enrich_auto,omitempty" json:"lead_enrich_auto,omitempty"`
	LeadMiningInPipeline                        *Bool      `xmlrpc:"lead_mining_in_pipeline,omitempty" json:"lead_mining_in_pipeline,omitempty"`
	LeaveTimesheetTaskId                        *Many2One  `xmlrpc:"leave_timesheet_task_id,omitempty" json:"leave_timesheet_task_id,omitempty"`
	LockConfirmedPo                             *Bool      `xmlrpc:"lock_confirmed_po,omitempty" json:"lock_confirmed_po,omitempty"`
	MapBoxToken                                 *String    `xmlrpc:"map_box_token,omitempty" json:"map_box_token,omitempty"`
	ModuleAccount3WayMatch                      *Bool      `xmlrpc:"module_account_3way_match,omitempty" json:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omitempty" json:"module_account_accountant,omitempty"`
	ModuleAccountAvatax                         *Bool      `xmlrpc:"module_account_avatax,omitempty" json:"module_account_avatax,omitempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty" json:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty" json:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty" json:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty" json:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omitempty" json:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omitempty" json:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omitempty" json:"module_account_check_printing,omitempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty" json:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omitempty" json:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omitempty" json:"module_account_invoice_extract,omitempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omitempty" json:"module_account_payment,omitempty"`
	ModuleAccountPeppol                         *Bool      `xmlrpc:"module_account_peppol,omitempty" json:"module_account_peppol,omitempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omitempty" json:"module_account_reports,omitempty"`
	ModuleAccountSepa                           *Bool      `xmlrpc:"module_account_sepa,omitempty" json:"module_account_sepa,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty" json:"module_account_sepa_direct_debit,omitempty"`
	ModuleAccountTaxcloud                       *Bool      `xmlrpc:"module_account_taxcloud,omitempty" json:"module_account_taxcloud,omitempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omitempty" json:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omitempty" json:"module_auth_oauth,omitempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omitempty" json:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omitempty" json:"module_base_import,omitempty"`
	ModuleCrmIapEnrich                          *Bool      `xmlrpc:"module_crm_iap_enrich,omitempty" json:"module_crm_iap_enrich,omitempty"`
	ModuleCrmIapMine                            *Bool      `xmlrpc:"module_crm_iap_mine,omitempty" json:"module_crm_iap_mine,omitempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omitempty" json:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omitempty" json:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omitempty" json:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omitempty" json:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omitempty" json:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omitempty" json:"module_delivery_fedex,omitempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omitempty" json:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryShiprocket                    *Bool      `xmlrpc:"module_delivery_shiprocket,omitempty" json:"module_delivery_shiprocket,omitempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omitempty" json:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omitempty" json:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omitempty" json:"module_google_calendar,omitempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omitempty" json:"module_google_gmail,omitempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omitempty" json:"module_google_recaptcha,omitempty"`
	ModuleHrAttendance                          *Bool      `xmlrpc:"module_hr_attendance,omitempty" json:"module_hr_attendance,omitempty"`
	ModuleHrExpenseExtract                      *Bool      `xmlrpc:"module_hr_expense_extract,omitempty" json:"module_hr_expense_extract,omitempty"`
	ModuleHrHomeworking                         *Bool      `xmlrpc:"module_hr_homeworking,omitempty" json:"module_hr_homeworking,omitempty"`
	ModuleHrPayrollExpense                      *Bool      `xmlrpc:"module_hr_payroll_expense,omitempty" json:"module_hr_payroll_expense,omitempty"`
	ModuleHrPresence                            *Bool      `xmlrpc:"module_hr_presence,omitempty" json:"module_hr_presence,omitempty"`
	ModuleHrSkills                              *Bool      `xmlrpc:"module_hr_skills,omitempty" json:"module_hr_skills,omitempty"`
	ModuleHrTimesheet                           *Bool      `xmlrpc:"module_hr_timesheet,omitempty" json:"module_hr_timesheet,omitempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omitempty" json:"module_l10n_eu_oss,omitempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omitempty" json:"module_loyalty,omitempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omitempty" json:"module_mail_plugin,omitempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omitempty" json:"module_microsoft_calendar,omitempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omitempty" json:"module_microsoft_outlook,omitempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omitempty" json:"module_partner_autocomplete,omitempty"`
	ModuleProductEmailTemplate                  *Bool      `xmlrpc:"module_product_email_template,omitempty" json:"module_product_email_template,omitempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omitempty" json:"module_product_expiry,omitempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omitempty" json:"module_product_images,omitempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omitempty" json:"module_product_margin,omitempty"`
	ModuleProjectTimesheetHolidays              *Bool      `xmlrpc:"module_project_timesheet_holidays,omitempty" json:"module_project_timesheet_holidays,omitempty"`
	ModulePurchaseProductMatrix                 *Bool      `xmlrpc:"module_purchase_product_matrix,omitempty" json:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition                   *Bool      `xmlrpc:"module_purchase_requisition,omitempty" json:"module_purchase_requisition,omitempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omitempty" json:"module_quality_control,omitempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omitempty" json:"module_quality_control_worksheet,omitempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omitempty" json:"module_sale_amazon,omitempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omitempty" json:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omitempty" json:"module_sale_margin,omitempty"`
	ModuleSalePdfQuoteBuilder                   *Bool      `xmlrpc:"module_sale_pdf_quote_builder,omitempty" json:"module_sale_pdf_quote_builder,omitempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omitempty" json:"module_sale_product_matrix,omitempty"`
	ModuleSignItsme                             *Bool      `xmlrpc:"module_sign_itsme,omitempty" json:"module_sign_itsme,omitempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omitempty" json:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omitempty" json:"module_stock_barcode,omitempty"`
	ModuleStockDropshipping                     *Bool      `xmlrpc:"module_stock_dropshipping,omitempty" json:"module_stock_dropshipping,omitempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omitempty" json:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omitempty" json:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omitempty" json:"module_stock_sms,omitempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omitempty" json:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omitempty" json:"module_web_unsplash,omitempty"`
	ModuleWebsiteCfTurnstile                    *Bool      `xmlrpc:"module_website_cf_turnstile,omitempty" json:"module_website_cf_turnstile,omitempty"`
	ModuleWebsiteCrmIapReveal                   *Bool      `xmlrpc:"module_website_crm_iap_reveal,omitempty" json:"module_website_crm_iap_reveal,omitempty"`
	OvertimeCompanyThreshold                    *Int       `xmlrpc:"overtime_company_threshold,omitempty" json:"overtime_company_threshold,omitempty"`
	OvertimeEmployeeThreshold                   *Int       `xmlrpc:"overtime_employee_threshold,omitempty" json:"overtime_employee_threshold,omitempty"`
	OvertimeStartDate                           *Time      `xmlrpc:"overtime_start_date,omitempty" json:"overtime_start_date,omitempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omitempty" json:"pay_invoices_online,omitempty"`
	PeriodLockDate                              *Time      `xmlrpc:"period_lock_date,omitempty" json:"period_lock_date,omitempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omitempty" json:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omitempty" json:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omitempty" json:"po_lead,omitempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omitempty" json:"po_lock,omitempty"`
	PoOrderApproval                             *Bool      `xmlrpc:"po_order_approval,omitempty" json:"po_order_approval,omitempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omitempty" json:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty" json:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty" json:"portal_confirmation_sign,omitempty"`
	PredictBillProduct                          *Bool      `xmlrpc:"predict_bill_product,omitempty" json:"predict_bill_product,omitempty"`
	PredictiveLeadScoringFieldLabels            *String    `xmlrpc:"predictive_lead_scoring_field_labels,omitempty" json:"predictive_lead_scoring_field_labels,omitempty"`
	PredictiveLeadScoringFields                 *Relation  `xmlrpc:"predictive_lead_scoring_fields,omitempty" json:"predictive_lead_scoring_fields,omitempty"`
	PredictiveLeadScoringFieldsStr              *String    `xmlrpc:"predictive_lead_scoring_fields_str,omitempty" json:"predictive_lead_scoring_fields_str,omitempty"`
	PredictiveLeadScoringStartDate              *Time      `xmlrpc:"predictive_lead_scoring_start_date,omitempty" json:"predictive_lead_scoring_start_date,omitempty"`
	PredictiveLeadScoringStartDateStr           *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omitempty" json:"predictive_lead_scoring_start_date_str,omitempty"`
	PrepaymentPercent                           *Float     `xmlrpc:"prepayment_percent,omitempty" json:"prepayment_percent,omitempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omitempty" json:"preview_ready,omitempty"`
	ProductFolder                               *Many2One  `xmlrpc:"product_folder,omitempty" json:"product_folder,omitempty"`
	ProductPricelistSetting                     *Selection `xmlrpc:"product_pricelist_setting,omitempty" json:"product_pricelist_setting,omitempty"`
	ProductTags                                 *Relation  `xmlrpc:"product_tags,omitempty" json:"product_tags,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty" json:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omitempty" json:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omitempty" json:"profiling_enabled_until,omitempty"`
	ProjectTimeModeId                           *Many2One  `xmlrpc:"project_time_mode_id,omitempty" json:"project_time_mode_id,omitempty"`
	PropertyAccountExpenseCategId               *Many2One  `xmlrpc:"property_account_expense_categ_id,omitempty" json:"property_account_expense_categ_id,omitempty"`
	PropertyAccountIncomeCategId                *Many2One  `xmlrpc:"property_account_income_categ_id,omitempty" json:"property_account_income_categ_id,omitempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty" json:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty" json:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockJournal                        *Many2One  `xmlrpc:"property_stock_journal,omitempty" json:"property_stock_journal,omitempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty" json:"property_stock_valuation_account_id,omitempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omitempty" json:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty" json:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty" json:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty" json:"quotation_validity_days,omitempty"`
	ReminderAllow                               *Bool      `xmlrpc:"reminder_allow,omitempty" json:"reminder_allow,omitempty"`
	ReminderDelay                               *Int       `xmlrpc:"reminder_delay,omitempty" json:"reminder_delay,omitempty"`
	ReminderInterval                            *Selection `xmlrpc:"reminder_interval,omitempty" json:"reminder_interval,omitempty"`
	ReminderUserAllow                           *Bool      `xmlrpc:"reminder_user_allow,omitempty" json:"reminder_user_allow,omitempty"`
	ReminderUserDelay                           *Int       `xmlrpc:"reminder_user_delay,omitempty" json:"reminder_user_delay,omitempty"`
	ReminderUserInterval                        *Selection `xmlrpc:"reminder_user_interval,omitempty" json:"reminder_user_interval,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty" json:"report_footer,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty" json:"resource_calendar_id,omitempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omitempty" json:"restrict_template_rendering,omitempty"`
	SaleFooter                                  *String    `xmlrpc:"sale_footer,omitempty" json:"sale_footer,omitempty"`
	SaleFooterName                              *String    `xmlrpc:"sale_footer_name,omitempty" json:"sale_footer_name,omitempty"`
	SaleHeader                                  *String    `xmlrpc:"sale_header,omitempty" json:"sale_header,omitempty"`
	SaleHeaderName                              *String    `xmlrpc:"sale_header_name,omitempty" json:"sale_header_name,omitempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omitempty" json:"sale_tax_id,omitempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omitempty" json:"security_lead,omitempty"`
	SfuServerKey                                *String    `xmlrpc:"sfu_server_key,omitempty" json:"sfu_server_key,omitempty"`
	SfuServerUrl                                *String    `xmlrpc:"sfu_server_url,omitempty" json:"sfu_server_url,omitempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omitempty" json:"show_effect,omitempty"`
	SignPreviewReady                            *Bool      `xmlrpc:"sign_preview_ready,omitempty" json:"sign_preview_ready,omitempty"`
	SignTerms                                   *String    `xmlrpc:"sign_terms,omitempty" json:"sign_terms,omitempty"`
	SignTermsHtml                               *String    `xmlrpc:"sign_terms_html,omitempty" json:"sign_terms_html,omitempty"`
	SignTermsType                               *Selection `xmlrpc:"sign_terms_type,omitempty" json:"sign_terms_type,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty" json:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty" json:"snailmail_cover,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty" json:"snailmail_duplex,omitempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omitempty" json:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omitempty" json:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty" json:"stock_sms_confirmation_template_id,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty" json:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty" json:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty" json:"tax_exigibility,omitempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omitempty" json:"tax_lock_date,omitempty"`
	TenorApiKey                                 *String    `xmlrpc:"tenor_api_key,omitempty" json:"tenor_api_key,omitempty"`
	TenorContentFilter                          *Selection `xmlrpc:"tenor_content_filter,omitempty" json:"tenor_content_filter,omitempty"`
	TenorGifLimit                               *Int       `xmlrpc:"tenor_gif_limit,omitempty" json:"tenor_gif_limit,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty" json:"terms_type,omitempty"`
	TimesheetEncodeMethod                       *Selection `xmlrpc:"timesheet_encode_method,omitempty" json:"timesheet_encode_method,omitempty"`
	TimesheetMinDuration                        *Int       `xmlrpc:"timesheet_min_duration,omitempty" json:"timesheet_min_duration,omitempty"`
	TimesheetRounding                           *Int       `xmlrpc:"timesheet_rounding,omitempty" json:"timesheet_rounding,omitempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omitempty" json:"totals_below_sections,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty" json:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omitempty" json:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omitempty" json:"twilio_account_token,omitempty"`
	UseAngloSaxon                               *Bool      `xmlrpc:"use_anglo_saxon,omitempty" json:"use_anglo_saxon,omitempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omitempty" json:"use_invoice_terms,omitempty"`
	UsePoLead                                   *Bool      `xmlrpc:"use_po_lead,omitempty" json:"use_po_lead,omitempty"`
	UseSecurityLead                             *Bool      `xmlrpc:"use_security_lead,omitempty" json:"use_security_lead,omitempty"`
	UseSignTerms                                *Bool      `xmlrpc:"use_sign_terms,omitempty" json:"use_sign_terms,omitempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omitempty" json:"use_twilio_rtc_servers,omitempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omitempty" json:"user_default_rights,omitempty"`
	VatCheckVies                                *Bool      `xmlrpc:"vat_check_vies,omitempty" json:"vat_check_vies,omitempty"`
	WebAppName                                  *String    `xmlrpc:"web_app_name,omitempty" json:"web_app_name,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
