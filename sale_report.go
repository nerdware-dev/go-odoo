package odoo

// SaleReport represents sale.report model.
type SaleReport struct {
	AnalyticAccountId      *Many2One  `xmlrpc:"analytic_account_id,omitempty" json:"analytic_account_id,omitempty"`
	CampaignId             *Many2One  `xmlrpc:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	CategId                *Many2One  `xmlrpc:"categ_id,omitempty" json:"categ_id,omitempty"`
	CommercialPartnerId    *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryId              *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                   *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Discount               *Float     `xmlrpc:"discount,omitempty" json:"discount,omitempty"`
	DiscountAmount         *Float     `xmlrpc:"discount_amount,omitempty" json:"discount_amount,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IndustryId             *Many2One  `xmlrpc:"industry_id,omitempty" json:"industry_id,omitempty"`
	InvoiceStatus          *Selection `xmlrpc:"invoice_status,omitempty" json:"invoice_status,omitempty"`
	Margin                 *Float     `xmlrpc:"margin,omitempty" json:"margin,omitempty"`
	MediumId               *Many2One  `xmlrpc:"medium_id,omitempty" json:"medium_id,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Nbr                    *Int       `xmlrpc:"nbr,omitempty" json:"nbr,omitempty"`
	OrderReference         *String    `xmlrpc:"order_reference,omitempty" json:"order_reference,omitempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerZip             *String    `xmlrpc:"partner_zip,omitempty" json:"partner_zip,omitempty"`
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
	SourceId               *Many2One  `xmlrpc:"source_id,omitempty" json:"source_id,omitempty"`
	State                  *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	StateId                *Many2One  `xmlrpc:"state_id,omitempty" json:"state_id,omitempty"`
	TeamId                 *Many2One  `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	UntaxedAmountInvoiced  *Float     `xmlrpc:"untaxed_amount_invoiced,omitempty" json:"untaxed_amount_invoiced,omitempty"`
	UntaxedAmountToInvoice *Float     `xmlrpc:"untaxed_amount_to_invoice,omitempty" json:"untaxed_amount_to_invoice,omitempty"`
	UserId                 *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	Volume                 *Float     `xmlrpc:"volume,omitempty" json:"volume,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`
	Weight                 *Float     `xmlrpc:"weight,omitempty" json:"weight,omitempty"`
}

// SaleReports represents array of sale.report model.
type SaleReports []SaleReport

// SaleReportModel is the odoo model name.
const SaleReportModel = "sale.report"

// Many2One convert SaleReport to *Many2One.
func (sr *SaleReport) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSaleReport creates a new sale.report model and returns its id.
func (c *Client) CreateSaleReport(sr *SaleReport) (int64, error) {
	ids, err := c.CreateSaleReports([]*SaleReport{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleReport creates a new sale.report model and returns its id.
func (c *Client) CreateSaleReports(srs []*SaleReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(SaleReportModel, vv, nil)
}

// UpdateSaleReport updates an existing sale.report record.
func (c *Client) UpdateSaleReport(sr *SaleReport) error {
	return c.UpdateSaleReports([]int64{sr.Id.Get()}, sr)
}

// UpdateSaleReports updates existing sale.report records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSaleReports(ids []int64, sr *SaleReport) error {
	return c.Update(SaleReportModel, ids, sr, nil)
}

// DeleteSaleReport deletes an existing sale.report record.
func (c *Client) DeleteSaleReport(id int64) error {
	return c.DeleteSaleReports([]int64{id})
}

// DeleteSaleReports deletes existing sale.report records.
func (c *Client) DeleteSaleReports(ids []int64) error {
	return c.Delete(SaleReportModel, ids)
}

// GetSaleReport gets sale.report existing record.
func (c *Client) GetSaleReport(id int64) (*SaleReport, error) {
	srs, err := c.GetSaleReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*srs) == 0 {
		return nil, nil
	}
	return &((*srs)[0]), nil
}

// GetSaleReports gets sale.report existing records.
func (c *Client) GetSaleReports(ids []int64) (*SaleReports, error) {
	srs := &SaleReports{}
	if err := c.Read(SaleReportModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSaleReport finds sale.report record by querying it with criteria.
func (c *Client) FindSaleReport(criteria *Criteria) (*SaleReport, error) {
	srs := &SaleReports{}
	if err := c.SearchRead(SaleReportModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if len(*srs) == 0 {
		return nil, nil
	}
	return &((*srs)[0]), nil
}

// FindSaleReports finds sale.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleReports(criteria *Criteria, options *Options) (*SaleReports, error) {
	srs := &SaleReports{}
	if err := c.SearchRead(SaleReportModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSaleReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleReportModel, criteria, options)
}

// FindSaleReportId finds record id by querying it with criteria.
func (c *Client) FindSaleReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
