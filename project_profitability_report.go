package odoo

// ProjectProfitabilityReport represents project.profitability.report model.
type ProjectProfitabilityReport struct {
	LastUpdate                    *Time     `xmlrpc:"__last_update,omitempty"`
	AmountUntaxedInvoiced         *Float    `xmlrpc:"amount_untaxed_invoiced,omitempty"`
	AmountUntaxedToInvoice        *Float    `xmlrpc:"amount_untaxed_to_invoice,omitempty"`
	AnalyticAccountId             *Many2One `xmlrpc:"analytic_account_id,omitempty"`
	CompanyId                     *Many2One `xmlrpc:"company_id,omitempty"`
	CurrencyId                    *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName                   *String   `xmlrpc:"display_name,omitempty"`
	ExpenseAmountUntaxedInvoiced  *Float    `xmlrpc:"expense_amount_untaxed_invoiced,omitempty"`
	ExpenseAmountUntaxedToInvoice *Float    `xmlrpc:"expense_amount_untaxed_to_invoice,omitempty"`
	ExpenseCost                   *Float    `xmlrpc:"expense_cost,omitempty"`
	Id                            *Int      `xmlrpc:"id,omitempty"`
	OrderConfirmationDate         *Time     `xmlrpc:"order_confirmation_date,omitempty"`
	PartnerId                     *Many2One `xmlrpc:"partner_id,omitempty"`
	ProductId                     *Many2One `xmlrpc:"product_id,omitempty"`
	ProjectId                     *Many2One `xmlrpc:"project_id,omitempty"`
	SaleLineId                    *Many2One `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId                   *Many2One `xmlrpc:"sale_order_id,omitempty"`
	TimesheetCost                 *Float    `xmlrpc:"timesheet_cost,omitempty"`
	TimesheetUnitAmount           *Float    `xmlrpc:"timesheet_unit_amount,omitempty"`
	UserId                        *Many2One `xmlrpc:"user_id,omitempty"`
}

// ProjectProfitabilityReports represents array of project.profitability.report model.
type ProjectProfitabilityReports []ProjectProfitabilityReport

// ProjectProfitabilityReportModel is the odoo model name.
const ProjectProfitabilityReportModel = "project.profitability.report"

// Many2One convert ProjectProfitabilityReport to *Many2One.
func (ppr *ProjectProfitabilityReport) Many2One() *Many2One {
	return NewMany2One(ppr.Id.Get(), "")
}

// CreateProjectProfitabilityReport creates a new project.profitability.report model and returns its id.
func (c *Client) CreateProjectProfitabilityReport(ppr *ProjectProfitabilityReport) (int64, error) {
	ids, err := c.CreateProjectProfitabilityReports([]*ProjectProfitabilityReport{ppr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProfitabilityReport creates a new project.profitability.report model and returns its id.
func (c *Client) CreateProjectProfitabilityReports(pprs []*ProjectProfitabilityReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range pprs {
		vv = append(vv, v)
	}
	return c.Create(ProjectProfitabilityReportModel, vv, nil)
}

// UpdateProjectProfitabilityReport updates an existing project.profitability.report record.
func (c *Client) UpdateProjectProfitabilityReport(ppr *ProjectProfitabilityReport) error {
	return c.UpdateProjectProfitabilityReports([]int64{ppr.Id.Get()}, ppr)
}

// UpdateProjectProfitabilityReports updates existing project.profitability.report records.
// All records (represented by ids) will be updated by ppr values.
func (c *Client) UpdateProjectProfitabilityReports(ids []int64, ppr *ProjectProfitabilityReport) error {
	return c.Update(ProjectProfitabilityReportModel, ids, ppr, nil)
}

// DeleteProjectProfitabilityReport deletes an existing project.profitability.report record.
func (c *Client) DeleteProjectProfitabilityReport(id int64) error {
	return c.DeleteProjectProfitabilityReports([]int64{id})
}

// DeleteProjectProfitabilityReports deletes existing project.profitability.report records.
func (c *Client) DeleteProjectProfitabilityReports(ids []int64) error {
	return c.Delete(ProjectProfitabilityReportModel, ids)
}

// GetProjectProfitabilityReport gets project.profitability.report existing record.
func (c *Client) GetProjectProfitabilityReport(id int64) (*ProjectProfitabilityReport, error) {
	pprs, err := c.GetProjectProfitabilityReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pprs)[0]), nil
}

// GetProjectProfitabilityReports gets project.profitability.report existing records.
func (c *Client) GetProjectProfitabilityReports(ids []int64) (*ProjectProfitabilityReports, error) {
	pprs := &ProjectProfitabilityReports{}
	if err := c.Read(ProjectProfitabilityReportModel, ids, nil, pprs); err != nil {
		return nil, err
	}
	return pprs, nil
}

// FindProjectProfitabilityReport finds project.profitability.report record by querying it with criteria.
func (c *Client) FindProjectProfitabilityReport(criteria *Criteria) (*ProjectProfitabilityReport, error) {
	pprs := &ProjectProfitabilityReports{}
	if err := c.SearchRead(ProjectProfitabilityReportModel, criteria, NewOptions().Limit(1), pprs); err != nil {
		return nil, err
	}
	return &((*pprs)[0]), nil
}

// FindProjectProfitabilityReports finds project.profitability.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProfitabilityReports(criteria *Criteria, options *Options) (*ProjectProfitabilityReports, error) {
	pprs := &ProjectProfitabilityReports{}
	if err := c.SearchRead(ProjectProfitabilityReportModel, criteria, options, pprs); err != nil {
		return nil, err
	}
	return pprs, nil
}

// FindProjectProfitabilityReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProfitabilityReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectProfitabilityReportModel, criteria, options)
}

// FindProjectProfitabilityReportId finds record id by querying it with criteria.
func (c *Client) FindProjectProfitabilityReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProfitabilityReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
