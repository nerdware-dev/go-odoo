package odoo

// LunchProductReport represents lunch.product.report model.
type LunchProductReport struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	Active        *Bool     `xmlrpc:"active,omitempty"`
	CategoryId    *Many2One `xmlrpc:"category_id,omitempty"`
	CompanyId     *Many2One `xmlrpc:"company_id,omitempty"`
	CurrencyId    *Many2One `xmlrpc:"currency_id,omitempty"`
	Description   *String   `xmlrpc:"description,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Image128      *String   `xmlrpc:"image_128,omitempty"`
	IsAvailableAt *Many2One `xmlrpc:"is_available_at,omitempty"`
	IsFavorite    *Bool     `xmlrpc:"is_favorite,omitempty"`
	IsNew         *Bool     `xmlrpc:"is_new,omitempty"`
	LastOrderDate *Time     `xmlrpc:"last_order_date,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	Price         *Float    `xmlrpc:"price,omitempty"`
	ProductId     *Many2One `xmlrpc:"product_id,omitempty"`
	SupplierId    *Many2One `xmlrpc:"supplier_id,omitempty"`
	UserId        *Many2One `xmlrpc:"user_id,omitempty"`
}

// LunchProductReports represents array of lunch.product.report model.
type LunchProductReports []LunchProductReport

// LunchProductReportModel is the odoo model name.
const LunchProductReportModel = "lunch.product.report"

// Many2One convert LunchProductReport to *Many2One.
func (lpr *LunchProductReport) Many2One() *Many2One {
	return NewMany2One(lpr.Id.Get(), "")
}

// CreateLunchProductReport creates a new lunch.product.report model and returns its id.
func (c *Client) CreateLunchProductReport(lpr *LunchProductReport) (int64, error) {
	ids, err := c.CreateLunchProductReports([]*LunchProductReport{lpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchProductReport creates a new lunch.product.report model and returns its id.
func (c *Client) CreateLunchProductReports(lprs []*LunchProductReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range lprs {
		vv = append(vv, v)
	}
	return c.Create(LunchProductReportModel, vv, nil)
}

// UpdateLunchProductReport updates an existing lunch.product.report record.
func (c *Client) UpdateLunchProductReport(lpr *LunchProductReport) error {
	return c.UpdateLunchProductReports([]int64{lpr.Id.Get()}, lpr)
}

// UpdateLunchProductReports updates existing lunch.product.report records.
// All records (represented by ids) will be updated by lpr values.
func (c *Client) UpdateLunchProductReports(ids []int64, lpr *LunchProductReport) error {
	return c.Update(LunchProductReportModel, ids, lpr, nil)
}

// DeleteLunchProductReport deletes an existing lunch.product.report record.
func (c *Client) DeleteLunchProductReport(id int64) error {
	return c.DeleteLunchProductReports([]int64{id})
}

// DeleteLunchProductReports deletes existing lunch.product.report records.
func (c *Client) DeleteLunchProductReports(ids []int64) error {
	return c.Delete(LunchProductReportModel, ids)
}

// GetLunchProductReport gets lunch.product.report existing record.
func (c *Client) GetLunchProductReport(id int64) (*LunchProductReport, error) {
	lprs, err := c.GetLunchProductReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lprs)[0]), nil
}

// GetLunchProductReports gets lunch.product.report existing records.
func (c *Client) GetLunchProductReports(ids []int64) (*LunchProductReports, error) {
	lprs := &LunchProductReports{}
	if err := c.Read(LunchProductReportModel, ids, nil, lprs); err != nil {
		return nil, err
	}
	return lprs, nil
}

// FindLunchProductReport finds lunch.product.report record by querying it with criteria.
func (c *Client) FindLunchProductReport(criteria *Criteria) (*LunchProductReport, error) {
	lprs := &LunchProductReports{}
	if err := c.SearchRead(LunchProductReportModel, criteria, NewOptions().Limit(1), lprs); err != nil {
		return nil, err
	}
	return &((*lprs)[0]), nil
}

// FindLunchProductReports finds lunch.product.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductReports(criteria *Criteria, options *Options) (*LunchProductReports, error) {
	lprs := &LunchProductReports{}
	if err := c.SearchRead(LunchProductReportModel, criteria, options, lprs); err != nil {
		return nil, err
	}
	return lprs, nil
}

// FindLunchProductReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchProductReportModel, criteria, options)
}

// FindLunchProductReportId finds record id by querying it with criteria.
func (c *Client) FindLunchProductReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchProductReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
