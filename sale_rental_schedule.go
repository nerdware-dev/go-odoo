package odoo

import (
	"fmt"
)

// SaleRentalSchedule represents sale.rental.schedule model.
type SaleRentalSchedule struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AnalyticAccountId   *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CardName            *String    `xmlrpc:"card_name,omptempty"`
	CategId             *Many2One  `xmlrpc:"categ_id,omptempty"`
	Color               *Int       `xmlrpc:"color,omptempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omptempty"`
	Description         *String    `xmlrpc:"description,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	Late                *Bool      `xmlrpc:"late,omptempty"`
	LotId               *Many2One  `xmlrpc:"lot_id,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	OrderDate           *Time      `xmlrpc:"order_date,omptempty"`
	OrderId             *Many2One  `xmlrpc:"order_id,omptempty"`
	OrderLineId         *Many2One  `xmlrpc:"order_line_id,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PickupDate          *Time      `xmlrpc:"pickup_date,omptempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductName         *String    `xmlrpc:"product_name,omptempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omptempty"`
	ProductUomQty       *Float     `xmlrpc:"product_uom_qty,omptempty"`
	QtyDelivered        *Float     `xmlrpc:"qty_delivered,omptempty"`
	QtyReturned         *Float     `xmlrpc:"qty_returned,omptempty"`
	RentalStatus        *Selection `xmlrpc:"rental_status,omptempty"`
	ReportLineStatus    *Selection `xmlrpc:"report_line_status,omptempty"`
	ReturnDate          *Time      `xmlrpc:"return_date,omptempty"`
	State               *Selection `xmlrpc:"state,omptempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omptempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omptempty"`
	WarehouseId         *Many2One  `xmlrpc:"warehouse_id,omptempty"`
}

// SaleRentalSchedules represents array of sale.rental.schedule model.
type SaleRentalSchedules []SaleRentalSchedule

// SaleRentalScheduleModel is the odoo model name.
const SaleRentalScheduleModel = "sale.rental.schedule"

// Many2One convert SaleRentalSchedule to *Many2One.
func (srs *SaleRentalSchedule) Many2One() *Many2One {
	return NewMany2One(srs.Id.Get(), "")
}

// CreateSaleRentalSchedule creates a new sale.rental.schedule model and returns its id.
func (c *Client) CreateSaleRentalSchedule(srs *SaleRentalSchedule) (int64, error) {
	ids, err := c.CreateSaleRentalSchedules([]*SaleRentalSchedule{srs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleRentalSchedule creates a new sale.rental.schedule model and returns its id.
func (c *Client) CreateSaleRentalSchedules(srss []*SaleRentalSchedule) ([]int64, error) {
	var vv []interface{}
	for _, v := range srss {
		vv = append(vv, v)
	}
	return c.Create(SaleRentalScheduleModel, vv)
}

// UpdateSaleRentalSchedule updates an existing sale.rental.schedule record.
func (c *Client) UpdateSaleRentalSchedule(srs *SaleRentalSchedule) error {
	return c.UpdateSaleRentalSchedules([]int64{srs.Id.Get()}, srs)
}

// UpdateSaleRentalSchedules updates existing sale.rental.schedule records.
// All records (represented by ids) will be updated by srs values.
func (c *Client) UpdateSaleRentalSchedules(ids []int64, srs *SaleRentalSchedule) error {
	return c.Update(SaleRentalScheduleModel, ids, srs)
}

// DeleteSaleRentalSchedule deletes an existing sale.rental.schedule record.
func (c *Client) DeleteSaleRentalSchedule(id int64) error {
	return c.DeleteSaleRentalSchedules([]int64{id})
}

// DeleteSaleRentalSchedules deletes existing sale.rental.schedule records.
func (c *Client) DeleteSaleRentalSchedules(ids []int64) error {
	return c.Delete(SaleRentalScheduleModel, ids)
}

// GetSaleRentalSchedule gets sale.rental.schedule existing record.
func (c *Client) GetSaleRentalSchedule(id int64) (*SaleRentalSchedule, error) {
	srss, err := c.GetSaleRentalSchedules([]int64{id})
	if err != nil {
		return nil, err
	}
	if srss != nil && len(*srss) > 0 {
		return &((*srss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.rental.schedule not found", id)
}

// GetSaleRentalSchedules gets sale.rental.schedule existing records.
func (c *Client) GetSaleRentalSchedules(ids []int64) (*SaleRentalSchedules, error) {
	srss := &SaleRentalSchedules{}
	if err := c.Read(SaleRentalScheduleModel, ids, nil, srss); err != nil {
		return nil, err
	}
	return srss, nil
}

// FindSaleRentalSchedule finds sale.rental.schedule record by querying it with criteria.
func (c *Client) FindSaleRentalSchedule(criteria *Criteria) (*SaleRentalSchedule, error) {
	srss := &SaleRentalSchedules{}
	if err := c.SearchRead(SaleRentalScheduleModel, criteria, NewOptions().Limit(1), srss); err != nil {
		return nil, err
	}
	if srss != nil && len(*srss) > 0 {
		return &((*srss)[0]), nil
	}
	return nil, fmt.Errorf("sale.rental.schedule was not found with criteria %v", criteria)
}

// FindSaleRentalSchedules finds sale.rental.schedule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleRentalSchedules(criteria *Criteria, options *Options) (*SaleRentalSchedules, error) {
	srss := &SaleRentalSchedules{}
	if err := c.SearchRead(SaleRentalScheduleModel, criteria, options, srss); err != nil {
		return nil, err
	}
	return srss, nil
}

// FindSaleRentalScheduleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleRentalScheduleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleRentalScheduleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleRentalScheduleId finds record id by querying it with criteria.
func (c *Client) FindSaleRentalScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleRentalScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.rental.schedule was not found with criteria %v and options %v", criteria, options)
}
