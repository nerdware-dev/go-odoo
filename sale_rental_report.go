package odoo

import (
	"fmt"
)

// SaleRentalReport represents sale.rental.report model.
type SaleRentalReport struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CategId       *Many2One  `xmlrpc:"categ_id,omptempty"`
	CompanyId     *Many2One  `xmlrpc:"company_id,omptempty"`
	CurrencyId    *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date          *Time      `xmlrpc:"date,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	OrderId       *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId     *Many2One  `xmlrpc:"partner_id,omptempty"`
	Price         *Float     `xmlrpc:"price,omptempty"`
	ProductId     *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductTmplId *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUom    *Many2One  `xmlrpc:"product_uom,omptempty"`
	QtyDelivered  *Float     `xmlrpc:"qty_delivered,omptempty"`
	QtyReturned   *Float     `xmlrpc:"qty_returned,omptempty"`
	Quantity      *Float     `xmlrpc:"quantity,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	UserId        *Many2One  `xmlrpc:"user_id,omptempty"`
}

// SaleRentalReports represents array of sale.rental.report model.
type SaleRentalReports []SaleRentalReport

// SaleRentalReportModel is the odoo model name.
const SaleRentalReportModel = "sale.rental.report"

// Many2One convert SaleRentalReport to *Many2One.
func (srr *SaleRentalReport) Many2One() *Many2One {
	return NewMany2One(srr.Id.Get(), "")
}

// CreateSaleRentalReport creates a new sale.rental.report model and returns its id.
func (c *Client) CreateSaleRentalReport(srr *SaleRentalReport) (int64, error) {
	ids, err := c.CreateSaleRentalReports([]*SaleRentalReport{srr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleRentalReport creates a new sale.rental.report model and returns its id.
func (c *Client) CreateSaleRentalReports(srrs []*SaleRentalReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range srrs {
		vv = append(vv, v)
	}
	return c.Create(SaleRentalReportModel, vv)
}

// UpdateSaleRentalReport updates an existing sale.rental.report record.
func (c *Client) UpdateSaleRentalReport(srr *SaleRentalReport) error {
	return c.UpdateSaleRentalReports([]int64{srr.Id.Get()}, srr)
}

// UpdateSaleRentalReports updates existing sale.rental.report records.
// All records (represented by ids) will be updated by srr values.
func (c *Client) UpdateSaleRentalReports(ids []int64, srr *SaleRentalReport) error {
	return c.Update(SaleRentalReportModel, ids, srr)
}

// DeleteSaleRentalReport deletes an existing sale.rental.report record.
func (c *Client) DeleteSaleRentalReport(id int64) error {
	return c.DeleteSaleRentalReports([]int64{id})
}

// DeleteSaleRentalReports deletes existing sale.rental.report records.
func (c *Client) DeleteSaleRentalReports(ids []int64) error {
	return c.Delete(SaleRentalReportModel, ids)
}

// GetSaleRentalReport gets sale.rental.report existing record.
func (c *Client) GetSaleRentalReport(id int64) (*SaleRentalReport, error) {
	srrs, err := c.GetSaleRentalReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.rental.report not found", id)
}

// GetSaleRentalReports gets sale.rental.report existing records.
func (c *Client) GetSaleRentalReports(ids []int64) (*SaleRentalReports, error) {
	srrs := &SaleRentalReports{}
	if err := c.Read(SaleRentalReportModel, ids, nil, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindSaleRentalReport finds sale.rental.report record by querying it with criteria.
func (c *Client) FindSaleRentalReport(criteria *Criteria) (*SaleRentalReport, error) {
	srrs := &SaleRentalReports{}
	if err := c.SearchRead(SaleRentalReportModel, criteria, NewOptions().Limit(1), srrs); err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("sale.rental.report was not found with criteria %v", criteria)
}

// FindSaleRentalReports finds sale.rental.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleRentalReports(criteria *Criteria, options *Options) (*SaleRentalReports, error) {
	srrs := &SaleRentalReports{}
	if err := c.SearchRead(SaleRentalReportModel, criteria, options, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindSaleRentalReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleRentalReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleRentalReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleRentalReportId finds record id by querying it with criteria.
func (c *Client) FindSaleRentalReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleRentalReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.rental.report was not found with criteria %v and options %v", criteria, options)
}
