package odoo

// SaleSubscriptionReport represents sale.subscription.report model.
type SaleSubscriptionReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AnalyticAccountId   *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	CategId             *Many2One  `xmlrpc:"categ_id,omitempty"`
	CloseReasonId       *Many2One  `xmlrpc:"close_reason_id,omitempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omitempty"`
	DateEnd             *Time      `xmlrpc:"date_end,omitempty"`
	DateStart           *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Health              *Selection `xmlrpc:"health,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InProgress          *Bool      `xmlrpc:"in_progress,omitempty"`
	IndustryId          *Many2One  `xmlrpc:"industry_id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PricelistId         *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omitempty"`
	Quantity            *Float     `xmlrpc:"quantity,omitempty"`
	RecurringMonthly    *Float     `xmlrpc:"recurring_monthly,omitempty"`
	RecurringTotal      *Float     `xmlrpc:"recurring_total,omitempty"`
	RecurringYearly     *Float     `xmlrpc:"recurring_yearly,omitempty"`
	StageId             *Many2One  `xmlrpc:"stage_id,omitempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omitempty"`
	TemplateId          *Many2One  `xmlrpc:"template_id,omitempty"`
	ToRenew             *Bool      `xmlrpc:"to_renew,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
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
	return ids[0], nil
}
