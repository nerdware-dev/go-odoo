package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty" json:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omitempty" json:"account_default_credit_limit,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One  `xmlrpc:"account_discount_expense_allocation_id,omitempty" json:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One  `xmlrpc:"account_discount_income_allocation_id,omitempty" json:"account_discount_income_allocation_id,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty" json:"account_fiscal_country_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty" json:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty" json:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty" json:"account_journal_suspense_account_id,omitempty"`
	AccountPeppolContactEmail                   *String    `xmlrpc:"account_peppol_contact_email,omitempty" json:"account_peppol_contact_email,omitempty"`
	AccountPeppolEas                            *Selection `xmlrpc:"account_peppol_eas,omitempty" json:"account_peppol_eas,omitempty"`
	AccountPeppolEdiIdentification              *String    `xmlrpc:"account_peppol_edi_identification,omitempty" json:"account_peppol_edi_identification,omitempty"`
	AccountPeppolEdiMode                        *Selection `xmlrpc:"account_peppol_edi_mode,omitempty" json:"account_peppol_edi_mode,omitempty"`
	AccountPeppolEdiUser                        *Many2One  `xmlrpc:"account_peppol_edi_user,omitempty" json:"account_peppol_edi_user,omitempty"`
	AccountPeppolEndpoint                       *String    `xmlrpc:"account_peppol_endpoint,omitempty" json:"account_peppol_endpoint,omitempty"`
	AccountPeppolMigrationKey                   *String    `xmlrpc:"account_peppol_migration_key,omitempty" json:"account_peppol_migration_key,omitempty"`
	AccountPeppolPhoneNumber                    *String    `xmlrpc:"account_peppol_phone_number,omitempty" json:"account_peppol_phone_number,omitempty"`
	AccountPeppolProxyState                     *Selection `xmlrpc:"account_peppol_proxy_state,omitempty" json:"account_peppol_proxy_state,omitempty"`
	AccountPeppolPurchaseJournalId              *Many2One  `xmlrpc:"account_peppol_purchase_journal_id,omitempty" json:"account_peppol_purchase_journal_id,omitempty"`
	AccountPriceInclude                         *Selection `xmlrpc:"account_price_include,omitempty" json:"account_price_include,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty" json:"account_storno,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty" json:"account_use_credit_limit,omitempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omitempty" json:"active_user_count,omitempty"`
	AliasDomainId                               *Many2One  `xmlrpc:"alias_domain_id,omitempty" json:"alias_domain_id,omitempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omitempty" json:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty" json:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omitempty" json:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omitempty" json:"automatic_invoice,omitempty"`
	AutopostBills                               *Bool      `xmlrpc:"autopost_bills,omitempty" json:"autopost_bills,omitempty"`
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
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty" json:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omitempty" json:"default_invoice_policy,omitempty"`
	DefaultPurchaseMethod                       *Selection `xmlrpc:"default_purchase_method,omitempty" json:"default_purchase_method,omitempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omitempty" json:"digest_emails,omitempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omitempty" json:"digest_id,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool      `xmlrpc:"display_invoice_amount_total_words,omitempty" json:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            *Bool      `xmlrpc:"display_invoice_tax_company_currency,omitempty" json:"display_invoice_tax_company_currency,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmailPrimaryColor                           *String    `xmlrpc:"email_primary_color,omitempty" json:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String    `xmlrpc:"email_secondary_color,omitempty" json:"email_secondary_color,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty" json:"expense_currency_exchange_account_id,omitempty"`
	ExpenseJournalId                            *Many2One  `xmlrpc:"expense_journal_id,omitempty" json:"expense_journal_id,omitempty"`
	ExpenseOutstandingAccountId                 *Many2One  `xmlrpc:"expense_outstanding_account_id,omitempty" json:"expense_outstanding_account_id,omitempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omitempty" json:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty" json:"external_report_layout_id,omitempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omitempty" json:"fail_counter,omitempty"`
	GoogleGmailClientIdentifier                 *String    `xmlrpc:"google_gmail_client_identifier,omitempty" json:"google_gmail_client_identifier,omitempty"`
	GoogleGmailClientSecret                     *String    `xmlrpc:"google_gmail_client_secret,omitempty" json:"google_gmail_client_secret,omitempty"`
	GoogleTranslateApiKey                       *String    `xmlrpc:"google_translate_api_key,omitempty" json:"google_translate_api_key,omitempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omitempty" json:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omitempty" json:"group_auto_done_setting,omitempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omitempty" json:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omitempty" json:"group_discount_per_so_line,omitempty"`
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
	GroupSendReminder                           *Bool      `xmlrpc:"group_send_reminder,omitempty" json:"group_send_reminder,omitempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty" json:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omitempty" json:"group_show_sale_receipts,omitempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omitempty" json:"group_stock_packaging,omitempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omitempty" json:"group_uom,omitempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omitempty" json:"group_warning_account,omitempty"`
	GroupWarningPurchase                        *Bool      `xmlrpc:"group_warning_purchase,omitempty" json:"group_warning_purchase,omitempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omitempty" json:"group_warning_sale,omitempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omitempty" json:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omitempty" json:"has_chart_of_accounts,omitempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omitempty" json:"hr_employee_self_edit,omitempty"`
	HrExpenseAliasDomainId                      *Many2One  `xmlrpc:"hr_expense_alias_domain_id,omitempty" json:"hr_expense_alias_domain_id,omitempty"`
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
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty" json:"invoice_mail_template_id,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty" json:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty" json:"invoice_terms_html,omitempty"`
	IsAccountPeppolEligible                     *Bool      `xmlrpc:"is_account_peppol_eligible,omitempty" json:"is_account_peppol_eligible,omitempty"`
	IsRootCompany                               *Bool      `xmlrpc:"is_root_company,omitempty" json:"is_root_company,omitempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omitempty" json:"language_count,omitempty"`
	LockConfirmedPo                             *Bool      `xmlrpc:"lock_confirmed_po,omitempty" json:"lock_confirmed_po,omitempty"`
	ModuleAccount3WayMatch                      *Bool      `xmlrpc:"module_account_3way_match,omitempty" json:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omitempty" json:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementExtract           *Bool      `xmlrpc:"module_account_bank_statement_extract,omitempty" json:"module_account_bank_statement_extract,omitempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty" json:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty" json:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty" json:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty" json:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omitempty" json:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omitempty" json:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omitempty" json:"module_account_check_printing,omitempty"`
	ModuleAccountExtract                        *Bool      `xmlrpc:"module_account_extract,omitempty" json:"module_account_extract,omitempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty" json:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omitempty" json:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omitempty" json:"module_account_invoice_extract,omitempty"`
	ModuleAccountIso20022                       *Bool      `xmlrpc:"module_account_iso20022,omitempty" json:"module_account_iso20022,omitempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omitempty" json:"module_account_payment,omitempty"`
	ModuleAccountPeppol                         *Bool      `xmlrpc:"module_account_peppol,omitempty" json:"module_account_peppol,omitempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omitempty" json:"module_account_reports,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty" json:"module_account_sepa_direct_debit,omitempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omitempty" json:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omitempty" json:"module_auth_oauth,omitempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omitempty" json:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omitempty" json:"module_base_import,omitempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omitempty" json:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omitempty" json:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omitempty" json:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omitempty" json:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omitempty" json:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omitempty" json:"module_delivery_fedex,omitempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omitempty" json:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryShiprocket                    *Bool      `xmlrpc:"module_delivery_shiprocket,omitempty" json:"module_delivery_shiprocket,omitempty"`
	ModuleDeliveryStarshipit                    *Bool      `xmlrpc:"module_delivery_starshipit,omitempty" json:"module_delivery_starshipit,omitempty"`
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
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omitempty" json:"module_product_images,omitempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omitempty" json:"module_product_margin,omitempty"`
	ModulePurchaseProductMatrix                 *Bool      `xmlrpc:"module_purchase_product_matrix,omitempty" json:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition                   *Bool      `xmlrpc:"module_purchase_requisition,omitempty" json:"module_purchase_requisition,omitempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omitempty" json:"module_sale_amazon,omitempty"`
	ModuleSaleCommission                        *Bool      `xmlrpc:"module_sale_commission,omitempty" json:"module_sale_commission,omitempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omitempty" json:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omitempty" json:"module_sale_margin,omitempty"`
	ModuleSalePdfQuoteBuilder                   *Bool      `xmlrpc:"module_sale_pdf_quote_builder,omitempty" json:"module_sale_pdf_quote_builder,omitempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omitempty" json:"module_sale_product_matrix,omitempty"`
	ModuleSms                                   *Bool      `xmlrpc:"module_sms,omitempty" json:"module_sms,omitempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omitempty" json:"module_snailmail_account,omitempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omitempty" json:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omitempty" json:"module_web_unsplash,omitempty"`
	ModuleWebsiteCfTurnstile                    *Bool      `xmlrpc:"module_website_cf_turnstile,omitempty" json:"module_website_cf_turnstile,omitempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty" json:"partner_autocomplete_insufficient_credit,omitempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omitempty" json:"pay_invoices_online,omitempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omitempty" json:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omitempty" json:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omitempty" json:"po_lead,omitempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omitempty" json:"po_lock,omitempty"`
	PoOrderApproval                             *Bool      `xmlrpc:"po_order_approval,omitempty" json:"po_order_approval,omitempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omitempty" json:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty" json:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty" json:"portal_confirmation_sign,omitempty"`
	PrepaymentPercent                           *Float     `xmlrpc:"prepayment_percent,omitempty" json:"prepayment_percent,omitempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omitempty" json:"preview_ready,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty" json:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omitempty" json:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omitempty" json:"profiling_enabled_until,omitempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omitempty" json:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty" json:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty" json:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty" json:"quotation_validity_days,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty" json:"report_footer,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty" json:"resource_calendar_id,omitempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omitempty" json:"restrict_template_rendering,omitempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omitempty" json:"sale_tax_id,omitempty"`
	SfuServerKey                                *String    `xmlrpc:"sfu_server_key,omitempty" json:"sfu_server_key,omitempty"`
	SfuServerUrl                                *String    `xmlrpc:"sfu_server_url,omitempty" json:"sfu_server_url,omitempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omitempty" json:"show_effect,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty" json:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty" json:"snailmail_cover,omitempty"`
	SnailmailCoverReadonly                      *Bool      `xmlrpc:"snailmail_cover_readonly,omitempty" json:"snailmail_cover_readonly,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty" json:"snailmail_duplex,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty" json:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty" json:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty" json:"tax_exigibility,omitempty"`
	TenorApiKey                                 *String    `xmlrpc:"tenor_api_key,omitempty" json:"tenor_api_key,omitempty"`
	TenorContentFilter                          *Selection `xmlrpc:"tenor_content_filter,omitempty" json:"tenor_content_filter,omitempty"`
	TenorGifLimit                               *Int       `xmlrpc:"tenor_gif_limit,omitempty" json:"tenor_gif_limit,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty" json:"terms_type,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty" json:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omitempty" json:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omitempty" json:"twilio_account_token,omitempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omitempty" json:"unsplash_access_key,omitempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omitempty" json:"unsplash_app_id,omitempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omitempty" json:"use_invoice_terms,omitempty"`
	UsePoLead                                   *Bool      `xmlrpc:"use_po_lead,omitempty" json:"use_po_lead,omitempty"`
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
