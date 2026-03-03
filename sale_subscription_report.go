package odoo

// SaleSubscriptionReport represents sale.subscription.report model.
type SaleSubscriptionReport struct {
	AnalyticAccountId      *Many2One  `xmlrpc:"analytic_account_id,omitempty" json:"analytic_account_id,omitempty"`
	CampaignId             *Many2One  `xmlrpc:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	CategId                *Many2One  `xmlrpc:"categ_id,omitempty" json:"categ_id,omitempty"`
	ClientOrderRef         *String    `xmlrpc:"client_order_ref,omitempty" json:"client_order_ref,omitempty"`
	CloseReasonId          *Many2One  `xmlrpc:"close_reason_id,omitempty" json:"close_reason_id,omitempty"`
	CommercialPartnerId    *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryId              *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                   *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Discount               *Float     `xmlrpc:"discount,omitempty" json:"discount,omitempty"`
	DiscountAmount         *Float     `xmlrpc:"discount_amount,omitempty" json:"discount_amount,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EndDate                *Time      `xmlrpc:"end_date,omitempty" json:"end_date,omitempty"`
	FirstContractDate      *Time      `xmlrpc:"first_contract_date,omitempty" json:"first_contract_date,omitempty"`
	Health                 *Selection `xmlrpc:"health,omitempty" json:"health,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IndustryId             *Many2One  `xmlrpc:"industry_id,omitempty" json:"industry_id,omitempty"`
	InvoiceStatus          *Selection `xmlrpc:"invoice_status,omitempty" json:"invoice_status,omitempty"`
	IsSubscription         *Bool      `xmlrpc:"is_subscription,omitempty" json:"is_subscription,omitempty"`
	Margin                 *Float     `xmlrpc:"margin,omitempty" json:"margin,omitempty"`
	MediumId               *Many2One  `xmlrpc:"medium_id,omitempty" json:"medium_id,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Nbr                    *Int       `xmlrpc:"nbr,omitempty" json:"nbr,omitempty"`
	NextInvoiceDate        *Time      `xmlrpc:"next_invoice_date,omitempty" json:"next_invoice_date,omitempty"`
	OrderReference         *String    `xmlrpc:"order_reference,omitempty" json:"order_reference,omitempty"`
	OriginOrderId          *Many2One  `xmlrpc:"origin_order_id,omitempty" json:"origin_order_id,omitempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerZip             *String    `xmlrpc:"partner_zip,omitempty" json:"partner_zip,omitempty"`
	PlanId                 *Many2One  `xmlrpc:"plan_id,omitempty" json:"plan_id,omitempty"`
	PriceSubtotal          *Float     `xmlrpc:"price_subtotal,omitempty" json:"price_subtotal,omitempty"`
	PriceTotal             *Float     `xmlrpc:"price_total,omitempty" json:"price_total,omitempty"`
	PricelistId            *Many2One  `xmlrpc:"pricelist_id,omitempty" json:"pricelist_id,omitempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductTmplId          *Many2One  `xmlrpc:"product_tmpl_id,omitempty" json:"product_tmpl_id,omitempty"`
	ProductUom             *Many2One  `xmlrpc:"product_uom,omitempty" json:"product_uom,omitempty"`
	ProductUomQty          *Float     `xmlrpc:"product_uom_qty,omitempty" json:"product_uom_qty,omitempty"`
	QtyDelivered           *Float     `xmlrpc:"qty_delivered,omitempty" json:"qty_delivered,omitempty"`
	QtyInvoiced            *Float     `xmlrpc:"qty_invoiced,omitempty" json:"qty_invoiced,omitempty"`
	QtyToDeliver           *Float     `xmlrpc:"qty_to_deliver,omitempty" json:"qty_to_deliver,omitempty"`
	QtyToInvoice           *Float     `xmlrpc:"qty_to_invoice,omitempty" json:"qty_to_invoice,omitempty"`
	RecurringMonthly       *Float     `xmlrpc:"recurring_monthly,omitempty" json:"recurring_monthly,omitempty"`
	RecurringTotal         *Float     `xmlrpc:"recurring_total,omitempty" json:"recurring_total,omitempty"`
	RecurringYearly        *Float     `xmlrpc:"recurring_yearly,omitempty" json:"recurring_yearly,omitempty"`
	SourceId               *Many2One  `xmlrpc:"source_id,omitempty" json:"source_id,omitempty"`
	State                  *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	StateId                *Many2One  `xmlrpc:"state_id,omitempty" json:"state_id,omitempty"`
	SubscriptionState      *Selection `xmlrpc:"subscription_state,omitempty" json:"subscription_state,omitempty"`
	TeamId                 *Many2One  `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	TemplateId             *Many2One  `xmlrpc:"template_id,omitempty" json:"template_id,omitempty"`
	UntaxedAmountInvoiced  *Float     `xmlrpc:"untaxed_amount_invoiced,omitempty" json:"untaxed_amount_invoiced,omitempty"`
	UntaxedAmountToInvoice *Float     `xmlrpc:"untaxed_amount_to_invoice,omitempty" json:"untaxed_amount_to_invoice,omitempty"`
	UserId                 *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	Volume                 *Float     `xmlrpc:"volume,omitempty" json:"volume,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`
	Weight                 *Float     `xmlrpc:"weight,omitempty" json:"weight,omitempty"`
}

// SaleSubscriptionReports represents array of sale.subscription.report model.
type SaleSubscriptionReports []SaleSubscriptionReport

// SaleSubscriptionReportModel is the odoo model name.
const SaleSubscriptionReportModel = "sale.subscription.report"

// Many2One convert SaleSubscriptionReport to *Many2One.
func (ssr *SaleSubscriptionReport) Many2One() *Many2One {
	return NewMany2One(ssr.Id.Get(), "")
}

// CreateSaleSubscriptionReport creates a new sale.subscription.report model and returns its id.
func (c *Client) CreateSaleSubscriptionReport(ssr *SaleSubscriptionReport) (int64, error) {
	ids, err := c.CreateSaleSubscriptionReports([]*SaleSubscriptionReport{ssr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionReport creates a new sale.subscription.report model and returns its id.
func (c *Client) CreateSaleSubscriptionReports(ssrs []*SaleSubscriptionReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrs {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionReportModel, vv, nil)
}

// UpdateSaleSubscriptionReport updates an existing sale.subscription.report record.
func (c *Client) UpdateSaleSubscriptionReport(ssr *SaleSubscriptionReport) error {
	return c.UpdateSaleSubscriptionReports([]int64{ssr.Id.Get()}, ssr)
}

// UpdateSaleSubscriptionReports updates existing sale.subscription.report records.
// All records (represented by ids) will be updated by ssr values.
func (c *Client) UpdateSaleSubscriptionReports(ids []int64, ssr *SaleSubscriptionReport) error {
	return c.Update(SaleSubscriptionReportModel, ids, ssr, nil)
}

// DeleteSaleSubscriptionReport deletes an existing sale.subscription.report record.
func (c *Client) DeleteSaleSubscriptionReport(id int64) error {
	return c.DeleteSaleSubscriptionReports([]int64{id})
}

// DeleteSaleSubscriptionReports deletes existing sale.subscription.report records.
func (c *Client) DeleteSaleSubscriptionReports(ids []int64) error {
	return c.Delete(SaleSubscriptionReportModel, ids)
}

// GetSaleSubscriptionReport gets sale.subscription.report existing record.
func (c *Client) GetSaleSubscriptionReport(id int64) (*SaleSubscriptionReport, error) {
	ssrs, err := c.GetSaleSubscriptionReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*ssrs) == 0 {
		return nil, nil
	}
	return &((*ssrs)[0]), nil
}

// GetSaleSubscriptionReports gets sale.subscription.report existing records.
func (c *Client) GetSaleSubscriptionReports(ids []int64) (*SaleSubscriptionReports, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.Read(SaleSubscriptionReportModel, ids, nil, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSaleSubscriptionReport finds sale.subscription.report record by querying it with criteria.
func (c *Client) FindSaleSubscriptionReport(criteria *Criteria) (*SaleSubscriptionReport, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.SearchRead(SaleSubscriptionReportModel, criteria, NewOptions().Limit(1), ssrs); err != nil {
		return nil, err
	}
	if len(*ssrs) == 0 {
		return nil, nil
	}
	return &((*ssrs)[0]), nil
}

// FindSaleSubscriptionReports finds sale.subscription.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionReports(criteria *Criteria, options *Options) (*SaleSubscriptionReports, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.SearchRead(SaleSubscriptionReportModel, criteria, options, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSaleSubscriptionReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionReportModel, criteria, options)
}

// FindSaleSubscriptionReportId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
