package odoo

// PurchaseReport represents purchase.report model.
type PurchaseReport struct {
	CategoryId          *Many2One  `xmlrpc:"category_id,omitempty" json:"category_id,omitempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	DateApprove         *Time      `xmlrpc:"date_approve,omitempty" json:"date_approve,omitempty"`
	DateOrder           *Time      `xmlrpc:"date_order,omitempty" json:"date_order,omitempty"`
	Delay               *Float     `xmlrpc:"delay,omitempty" json:"delay,omitempty"`
	DelayPass           *Float     `xmlrpc:"delay_pass,omitempty" json:"delay_pass,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	FiscalPositionId    *Many2One  `xmlrpc:"fiscal_position_id,omitempty" json:"fiscal_position_id,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	NbrLines            *Int       `xmlrpc:"nbr_lines,omitempty" json:"nbr_lines,omitempty"`
	OrderId             *Many2One  `xmlrpc:"order_id,omitempty" json:"order_id,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PriceAverage        *Float     `xmlrpc:"price_average,omitempty" json:"price_average,omitempty"`
	PriceTotal          *Float     `xmlrpc:"price_total,omitempty" json:"price_total,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omitempty" json:"product_tmpl_id,omitempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omitempty" json:"product_uom,omitempty"`
	QtyBilled           *Float     `xmlrpc:"qty_billed,omitempty" json:"qty_billed,omitempty"`
	QtyOrdered          *Float     `xmlrpc:"qty_ordered,omitempty" json:"qty_ordered,omitempty"`
	QtyReceived         *Float     `xmlrpc:"qty_received,omitempty" json:"qty_received,omitempty"`
	QtyToBeBilled       *Float     `xmlrpc:"qty_to_be_billed,omitempty" json:"qty_to_be_billed,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	UntaxedTotal        *Float     `xmlrpc:"untaxed_total,omitempty" json:"untaxed_total,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	Volume              *Float     `xmlrpc:"volume,omitempty" json:"volume,omitempty"`
	Weight              *Float     `xmlrpc:"weight,omitempty" json:"weight,omitempty"`
}

// PurchaseReports represents array of purchase.report model.
type PurchaseReports []PurchaseReport

// PurchaseReportModel is the odoo model name.
const PurchaseReportModel = "purchase.report"

// Many2One convert PurchaseReport to *Many2One.
func (pr *PurchaseReport) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePurchaseReport creates a new purchase.report model and returns its id.
func (c *Client) CreatePurchaseReport(pr *PurchaseReport) (int64, error) {
	ids, err := c.CreatePurchaseReports([]*PurchaseReport{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseReport creates a new purchase.report model and returns its id.
func (c *Client) CreatePurchaseReports(prs []*PurchaseReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(PurchaseReportModel, vv, nil)
}

// UpdatePurchaseReport updates an existing purchase.report record.
func (c *Client) UpdatePurchaseReport(pr *PurchaseReport) error {
	return c.UpdatePurchaseReports([]int64{pr.Id.Get()}, pr)
}

// UpdatePurchaseReports updates existing purchase.report records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePurchaseReports(ids []int64, pr *PurchaseReport) error {
	return c.Update(PurchaseReportModel, ids, pr, nil)
}

// DeletePurchaseReport deletes an existing purchase.report record.
func (c *Client) DeletePurchaseReport(id int64) error {
	return c.DeletePurchaseReports([]int64{id})
}

// DeletePurchaseReports deletes existing purchase.report records.
func (c *Client) DeletePurchaseReports(ids []int64) error {
	return c.Delete(PurchaseReportModel, ids)
}

// GetPurchaseReport gets purchase.report existing record.
func (c *Client) GetPurchaseReport(id int64) (*PurchaseReport, error) {
	prs, err := c.GetPurchaseReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetPurchaseReports gets purchase.report existing records.
func (c *Client) GetPurchaseReports(ids []int64) (*PurchaseReports, error) {
	prs := &PurchaseReports{}
	if err := c.Read(PurchaseReportModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseReport finds purchase.report record by querying it with criteria.
func (c *Client) FindPurchaseReport(criteria *Criteria) (*PurchaseReport, error) {
	prs := &PurchaseReports{}
	if err := c.SearchRead(PurchaseReportModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindPurchaseReports finds purchase.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseReports(criteria *Criteria, options *Options) (*PurchaseReports, error) {
	prs := &PurchaseReports{}
	if err := c.SearchRead(PurchaseReportModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseReportModel, criteria, options)
}

// FindPurchaseReportId finds record id by querying it with criteria.
func (c *Client) FindPurchaseReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
