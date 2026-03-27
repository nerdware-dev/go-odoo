package odoo

// SaleOrder represents sale.order model.
type SaleOrder struct {
	AccessToken                  *String    `xmlrpc:"access_token,omitempty" json:"access_token,omitempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omitempty" json:"access_url,omitempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omitempty" json:"access_warning,omitempty"`
	ActivityCalendarEventId      *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AmountInvoiced               *Float     `xmlrpc:"amount_invoiced,omitempty" json:"amount_invoiced,omitempty"`
	AmountPaid                   *Float     `xmlrpc:"amount_paid,omitempty" json:"amount_paid,omitempty"`
	AmountTax                    *Float     `xmlrpc:"amount_tax,omitempty" json:"amount_tax,omitempty"`
	AmountToInvoice              *Float     `xmlrpc:"amount_to_invoice,omitempty" json:"amount_to_invoice,omitempty"`
	AmountTotal                  *Float     `xmlrpc:"amount_total,omitempty" json:"amount_total,omitempty"`
	AmountUndiscounted           *Float     `xmlrpc:"amount_undiscounted,omitempty" json:"amount_undiscounted,omitempty"`
	AmountUntaxed                *Float     `xmlrpc:"amount_untaxed,omitempty" json:"amount_untaxed,omitempty"`
	AnalyticAccountId            *Many2One  `xmlrpc:"analytic_account_id,omitempty" json:"analytic_account_id,omitempty"`
	ArchivedProductCount         *Int       `xmlrpc:"archived_product_count,omitempty" json:"archived_product_count,omitempty"`
	ArchivedProductIds           *Relation  `xmlrpc:"archived_product_ids,omitempty" json:"archived_product_ids,omitempty"`
	AuthorizedTransactionIds     *Relation  `xmlrpc:"authorized_transaction_ids,omitempty" json:"authorized_transaction_ids,omitempty"`
	CampaignId                   *Many2One  `xmlrpc:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	ClientOrderRef               *String    `xmlrpc:"client_order_ref,omitempty" json:"client_order_ref,omitempty"`
	CloseReasonId                *Many2One  `xmlrpc:"close_reason_id,omitempty" json:"close_reason_id,omitempty"`
	CommercialPartnerId          *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CommitmentDate               *Time      `xmlrpc:"commitment_date,omitempty" json:"commitment_date,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	CurrencyRate                 *Float     `xmlrpc:"currency_rate,omitempty" json:"currency_rate,omitempty"`
	DateOrder                    *Time      `xmlrpc:"date_order,omitempty" json:"date_order,omitempty"`
	DeliveryCount                *Int       `xmlrpc:"delivery_count,omitempty" json:"delivery_count,omitempty"`
	DeliveryStatus               *Selection `xmlrpc:"delivery_status,omitempty" json:"delivery_status,omitempty"`
	DisplayLate                  *Bool      `xmlrpc:"display_late,omitempty" json:"display_late,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EffectiveDate                *Time      `xmlrpc:"effective_date,omitempty" json:"effective_date,omitempty"`
	EndDate                      *Time      `xmlrpc:"end_date,omitempty" json:"end_date,omitempty"`
	ExpectedDate                 *Time      `xmlrpc:"expected_date,omitempty" json:"expected_date,omitempty"`
	ExpenseCount                 *Int       `xmlrpc:"expense_count,omitempty" json:"expense_count,omitempty"`
	ExpenseIds                   *Relation  `xmlrpc:"expense_ids,omitempty" json:"expense_ids,omitempty"`
	FirstContractDate            *Time      `xmlrpc:"first_contract_date,omitempty" json:"first_contract_date,omitempty"`
	FiscalPositionId             *Many2One  `xmlrpc:"fiscal_position_id,omitempty" json:"fiscal_position_id,omitempty"`
	HasActivePricelist           *Bool      `xmlrpc:"has_active_pricelist,omitempty" json:"has_active_pricelist,omitempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	HasRecurringLine             *Bool      `xmlrpc:"has_recurring_line,omitempty" json:"has_recurring_line,omitempty"`
	Health                       *Selection `xmlrpc:"health,omitempty" json:"health,omitempty"`
	HistoryCount                 *Int       `xmlrpc:"history_count,omitempty" json:"history_count,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Incoterm                     *Many2One  `xmlrpc:"incoterm,omitempty" json:"incoterm,omitempty"`
	IncotermLocation             *String    `xmlrpc:"incoterm_location,omitempty" json:"incoterm_location,omitempty"`
	InternalNote                 *String    `xmlrpc:"internal_note,omitempty" json:"internal_note,omitempty"`
	InternalNoteDisplay          *String    `xmlrpc:"internal_note_display,omitempty" json:"internal_note_display,omitempty"`
	InvoiceCount                 *Int       `xmlrpc:"invoice_count,omitempty" json:"invoice_count,omitempty"`
	InvoiceIds                   *Relation  `xmlrpc:"invoice_ids,omitempty" json:"invoice_ids,omitempty"`
	InvoiceStatus                *Selection `xmlrpc:"invoice_status,omitempty" json:"invoice_status,omitempty"`
	IsBatch                      *Bool      `xmlrpc:"is_batch,omitempty" json:"is_batch,omitempty"`
	IsExpired                    *Bool      `xmlrpc:"is_expired,omitempty" json:"is_expired,omitempty"`
	IsInvoiceCron                *Bool      `xmlrpc:"is_invoice_cron,omitempty" json:"is_invoice_cron,omitempty"`
	IsProductMilestone           *Bool      `xmlrpc:"is_product_milestone,omitempty" json:"is_product_milestone,omitempty"`
	IsRenewing                   *Bool      `xmlrpc:"is_renewing,omitempty" json:"is_renewing,omitempty"`
	IsSubscription               *Bool      `xmlrpc:"is_subscription,omitempty" json:"is_subscription,omitempty"`
	IsUpselling                  *Bool      `xmlrpc:"is_upselling,omitempty" json:"is_upselling,omitempty"`
	JournalId                    *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	JsonPopover                  *String    `xmlrpc:"json_popover,omitempty" json:"json_popover,omitempty"`
	Kpi1MonthMrrDelta            *Float     `xmlrpc:"kpi_1month_mrr_delta,omitempty" json:"kpi_1month_mrr_delta,omitempty"`
	Kpi1MonthMrrPercentage       *Float     `xmlrpc:"kpi_1month_mrr_percentage,omitempty" json:"kpi_1month_mrr_percentage,omitempty"`
	Kpi3MonthsMrrDelta           *Float     `xmlrpc:"kpi_3months_mrr_delta,omitempty" json:"kpi_3months_mrr_delta,omitempty"`
	Kpi3MonthsMrrPercentage      *Float     `xmlrpc:"kpi_3months_mrr_percentage,omitempty" json:"kpi_3months_mrr_percentage,omitempty"`
	L10NDin5008Addresses         *String    `xmlrpc:"l10n_din5008_addresses,omitempty" json:"l10n_din5008_addresses,omitempty"`
	L10NDin5008DocumentTitle     *String    `xmlrpc:"l10n_din5008_document_title,omitempty" json:"l10n_din5008_document_title,omitempty"`
	L10NDin5008TemplateData      *String    `xmlrpc:"l10n_din5008_template_data,omitempty" json:"l10n_din5008_template_data,omitempty"`
	LastInvoiceDate              *Time      `xmlrpc:"last_invoice_date,omitempty" json:"last_invoice_date,omitempty"`
	Locked                       *Bool      `xmlrpc:"locked,omitempty" json:"locked,omitempty"`
	Margin                       *Float     `xmlrpc:"margin,omitempty" json:"margin,omitempty"`
	MarginPercent                *Float     `xmlrpc:"margin_percent,omitempty" json:"margin_percent,omitempty"`
	MediumId                     *Many2One  `xmlrpc:"medium_id,omitempty" json:"medium_id,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MilestoneCount               *Int       `xmlrpc:"milestone_count,omitempty" json:"milestone_count,omitempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NextInvoiceDate              *Time      `xmlrpc:"next_invoice_date,omitempty" json:"next_invoice_date,omitempty"`
	NonRecurringTotal            *Float     `xmlrpc:"non_recurring_total,omitempty" json:"non_recurring_total,omitempty"`
	Note                         *String    `xmlrpc:"note,omitempty" json:"note,omitempty"`
	NoteOrder                    *Many2One  `xmlrpc:"note_order,omitempty" json:"note_order,omitempty"`
	OpportunityId                *Many2One  `xmlrpc:"opportunity_id,omitempty" json:"opportunity_id,omitempty"`
	OrderLine                    *Relation  `xmlrpc:"order_line,omitempty" json:"order_line,omitempty"`
	OrderLogIds                  *Relation  `xmlrpc:"order_log_ids,omitempty" json:"order_log_ids,omitempty"`
	Origin                       *String    `xmlrpc:"origin,omitempty" json:"origin,omitempty"`
	OriginOrderId                *Many2One  `xmlrpc:"origin_order_id,omitempty" json:"origin_order_id,omitempty"`
	PartnerCreditWarning         *String    `xmlrpc:"partner_credit_warning,omitempty" json:"partner_credit_warning,omitempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerInvoiceId             *Many2One  `xmlrpc:"partner_invoice_id,omitempty" json:"partner_invoice_id,omitempty"`
	PartnerShippingId            *Many2One  `xmlrpc:"partner_shipping_id,omitempty" json:"partner_shipping_id,omitempty"`
	PaymentException             *Bool      `xmlrpc:"payment_exception,omitempty" json:"payment_exception,omitempty"`
	PaymentTermId                *Many2One  `xmlrpc:"payment_term_id,omitempty" json:"payment_term_id,omitempty"`
	PaymentTokenId               *Many2One  `xmlrpc:"payment_token_id,omitempty" json:"payment_token_id,omitempty"`
	PendingEmailTemplateId       *Many2One  `xmlrpc:"pending_email_template_id,omitempty" json:"pending_email_template_id,omitempty"`
	PendingTransaction           *Bool      `xmlrpc:"pending_transaction,omitempty" json:"pending_transaction,omitempty"`
	PercentageSatisfaction       *Int       `xmlrpc:"percentage_satisfaction,omitempty" json:"percentage_satisfaction,omitempty"`
	PickingIds                   *Relation  `xmlrpc:"picking_ids,omitempty" json:"picking_ids,omitempty"`
	PickingPolicy                *Selection `xmlrpc:"picking_policy,omitempty" json:"picking_policy,omitempty"`
	PlanId                       *Many2One  `xmlrpc:"plan_id,omitempty" json:"plan_id,omitempty"`
	PrepaymentPercent            *Float     `xmlrpc:"prepayment_percent,omitempty" json:"prepayment_percent,omitempty"`
	PricelistId                  *Many2One  `xmlrpc:"pricelist_id,omitempty" json:"pricelist_id,omitempty"`
	ProcurementGroupId           *Many2One  `xmlrpc:"procurement_group_id,omitempty" json:"procurement_group_id,omitempty"`
	ProjectCount                 *Int       `xmlrpc:"project_count,omitempty" json:"project_count,omitempty"`
	ProjectId                    *Many2One  `xmlrpc:"project_id,omitempty" json:"project_id,omitempty"`
	ProjectIds                   *Relation  `xmlrpc:"project_ids,omitempty" json:"project_ids,omitempty"`
	PurchaseOrderCount           *Int       `xmlrpc:"purchase_order_count,omitempty" json:"purchase_order_count,omitempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omitempty" json:"rating_avg,omitempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omitempty" json:"rating_avg_text,omitempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omitempty" json:"rating_count,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omitempty" json:"rating_last_feedback,omitempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omitempty" json:"rating_last_image,omitempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omitempty" json:"rating_last_text,omitempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omitempty" json:"rating_last_value,omitempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omitempty" json:"rating_percentage_satisfaction,omitempty"`
	RecurringDetails             *String    `xmlrpc:"recurring_details,omitempty" json:"recurring_details,omitempty"`
	RecurringMonthly             *Float     `xmlrpc:"recurring_monthly,omitempty" json:"recurring_monthly,omitempty"`
	RecurringTotal               *Float     `xmlrpc:"recurring_total,omitempty" json:"recurring_total,omitempty"`
	Reference                    *String    `xmlrpc:"reference,omitempty" json:"reference,omitempty"`
	RenewalCount                 *Int       `xmlrpc:"renewal_count,omitempty" json:"renewal_count,omitempty"`
	RequirePayment               *Bool      `xmlrpc:"require_payment,omitempty" json:"require_payment,omitempty"`
	RequireSignature             *Bool      `xmlrpc:"require_signature,omitempty" json:"require_signature,omitempty"`
	SaleOrderOptionIds           *Relation  `xmlrpc:"sale_order_option_ids,omitempty" json:"sale_order_option_ids,omitempty"`
	SaleOrderTemplateId          *Many2One  `xmlrpc:"sale_order_template_id,omitempty" json:"sale_order_template_id,omitempty"`
	ShowCreateProjectButton      *Bool      `xmlrpc:"show_create_project_button,omitempty" json:"show_create_project_button,omitempty"`
	ShowHoursRecordedButton      *Bool      `xmlrpc:"show_hours_recorded_button,omitempty" json:"show_hours_recorded_button,omitempty"`
	ShowJsonPopover              *Bool      `xmlrpc:"show_json_popover,omitempty" json:"show_json_popover,omitempty"`
	ShowProjectButton            *Bool      `xmlrpc:"show_project_button,omitempty" json:"show_project_button,omitempty"`
	ShowTaskButton               *Bool      `xmlrpc:"show_task_button,omitempty" json:"show_task_button,omitempty"`
	ShowUpdateFpos               *Bool      `xmlrpc:"show_update_fpos,omitempty" json:"show_update_fpos,omitempty"`
	ShowUpdatePricelist          *Bool      `xmlrpc:"show_update_pricelist,omitempty" json:"show_update_pricelist,omitempty"`
	Signature                    *String    `xmlrpc:"signature,omitempty" json:"signature,omitempty"`
	SignedBy                     *String    `xmlrpc:"signed_by,omitempty" json:"signed_by,omitempty"`
	SignedOn                     *Time      `xmlrpc:"signed_on,omitempty" json:"signed_on,omitempty"`
	SourceId                     *Many2One  `xmlrpc:"source_id,omitempty" json:"source_id,omitempty"`
	Starred                      *Bool      `xmlrpc:"starred,omitempty" json:"starred,omitempty"`
	StarredUserIds               *Relation  `xmlrpc:"starred_user_ids,omitempty" json:"starred_user_ids,omitempty"`
	StartDate                    *Time      `xmlrpc:"start_date,omitempty" json:"start_date,omitempty"`
	State                        *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	SubscriptionChildIds         *Relation  `xmlrpc:"subscription_child_ids,omitempty" json:"subscription_child_ids,omitempty"`
	SubscriptionId               *Many2One  `xmlrpc:"subscription_id,omitempty" json:"subscription_id,omitempty"`
	SubscriptionState            *Selection `xmlrpc:"subscription_state,omitempty" json:"subscription_state,omitempty"`
	TagIds                       *Relation  `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TasksCount                   *Int       `xmlrpc:"tasks_count,omitempty" json:"tasks_count,omitempty"`
	TasksIds                     *Relation  `xmlrpc:"tasks_ids,omitempty" json:"tasks_ids,omitempty"`
	TaxCalculationRoundingMethod *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty" json:"tax_calculation_rounding_method,omitempty"`
	TaxCountryId                 *Many2One  `xmlrpc:"tax_country_id,omitempty" json:"tax_country_id,omitempty"`
	TaxTotals                    *String    `xmlrpc:"tax_totals,omitempty" json:"tax_totals,omitempty"`
	TeamId                       *Many2One  `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	TeamUserId                   *Many2One  `xmlrpc:"team_user_id,omitempty" json:"team_user_id,omitempty"`
	TermsType                    *Selection `xmlrpc:"terms_type,omitempty" json:"terms_type,omitempty"`
	TimesheetCount               *Float     `xmlrpc:"timesheet_count,omitempty" json:"timesheet_count,omitempty"`
	TimesheetEncodeUomId         *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty" json:"timesheet_encode_uom_id,omitempty"`
	TimesheetTotalDuration       *Int       `xmlrpc:"timesheet_total_duration,omitempty" json:"timesheet_total_duration,omitempty"`
	TransactionIds               *Relation  `xmlrpc:"transaction_ids,omitempty" json:"transaction_ids,omitempty"`
	TypeName                     *String    `xmlrpc:"type_name,omitempty" json:"type_name,omitempty"`
	UpsellCount                  *Int       `xmlrpc:"upsell_count,omitempty" json:"upsell_count,omitempty"`
	UserClosable                 *Bool      `xmlrpc:"user_closable,omitempty" json:"user_closable,omitempty"`
	UserExtend                   *Bool      `xmlrpc:"user_extend,omitempty" json:"user_extend,omitempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	UserQuantity                 *Bool      `xmlrpc:"user_quantity,omitempty" json:"user_quantity,omitempty"`
	ValidityDate                 *Time      `xmlrpc:"validity_date,omitempty" json:"validity_date,omitempty"`
	VisibleProject               *Bool      `xmlrpc:"visible_project,omitempty" json:"visible_project,omitempty"`
	WarehouseId                  *Many2One  `xmlrpc:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// SaleOrders represents array of sale.order model.
type SaleOrders []SaleOrder

// SaleOrderModel is the odoo model name.
const SaleOrderModel = "sale.order"

// Many2One convert SaleOrder to *Many2One.
func (so *SaleOrder) Many2One() *Many2One {
	return NewMany2One(so.Id.Get(), "")
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrder(so *SaleOrder) (int64, error) {
	ids, err := c.CreateSaleOrders([]*SaleOrder{so})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv, nil)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so, nil)
}

// DeleteSaleOrder deletes an existing sale.order record.
func (c *Client) DeleteSaleOrder(id int64) error {
	return c.DeleteSaleOrders([]int64{id})
}

// DeleteSaleOrders deletes existing sale.order records.
func (c *Client) DeleteSaleOrders(ids []int64) error {
	return c.Delete(SaleOrderModel, ids)
}

// GetSaleOrder gets sale.order existing record.
func (c *Client) GetSaleOrder(id int64) (*SaleOrder, error) {
	sos, err := c.GetSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*sos) == 0 {
		return nil, nil
	}
	return &((*sos)[0]), nil
}

// GetSaleOrders gets sale.order existing records.
func (c *Client) GetSaleOrders(ids []int64) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.Read(SaleOrderModel, ids, nil, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrder finds sale.order record by querying it with criteria.
func (c *Client) FindSaleOrder(criteria *Criteria) (*SaleOrder, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, NewOptions().Limit(1), sos); err != nil {
		return nil, err
	}
	if len(*sos) == 0 {
		return nil, nil
	}
	return &((*sos)[0]), nil
}

// FindSaleOrders finds sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrders(criteria *Criteria, options *Options) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, options, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderModel, criteria, options)
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
