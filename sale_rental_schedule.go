package odoo

// SaleRentalSchedule represents sale.rental.schedule model.
type SaleRentalSchedule struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AnalyticAccountId   *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	CardName            *String    `xmlrpc:"card_name,omitempty"`
	CategId             *Many2One  `xmlrpc:"categ_id,omitempty"`
	Color               *Int       `xmlrpc:"color,omitempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omitempty"`
	Description         *String    `xmlrpc:"description,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Late                *Bool      `xmlrpc:"late,omitempty"`
	LotId               *Many2One  `xmlrpc:"lot_id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	OrderDate           *Time      `xmlrpc:"order_date,omitempty"`
	OrderId             *Many2One  `xmlrpc:"order_id,omitempty"`
	OrderLineId         *Many2One  `xmlrpc:"order_line_id,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickupDate          *Time      `xmlrpc:"pickup_date,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductName         *String    `xmlrpc:"product_name,omitempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omitempty"`
	ProductUomQty       *Float     `xmlrpc:"product_uom_qty,omitempty"`
	QtyDelivered        *Float     `xmlrpc:"qty_delivered,omitempty"`
	QtyReturned         *Float     `xmlrpc:"qty_returned,omitempty"`
	RentalStatus        *Selection `xmlrpc:"rental_status,omitempty"`
	ReportLineStatus    *Selection `xmlrpc:"report_line_status,omitempty"`
	ReturnDate          *Time      `xmlrpc:"return_date,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
	WarehouseId         *Many2One  `xmlrpc:"warehouse_id,omitempty"`
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
	return c.Create(SaleRentalScheduleModel, vv, nil)
}

// UpdateSaleRentalSchedule updates an existing sale.rental.schedule record.
func (c *Client) UpdateSaleRentalSchedule(srs *SaleRentalSchedule) error {
	return c.UpdateSaleRentalSchedules([]int64{srs.Id.Get()}, srs)
}

// UpdateSaleRentalSchedules updates existing sale.rental.schedule records.
// All records (represented by ids) will be updated by srs values.
func (c *Client) UpdateSaleRentalSchedules(ids []int64, srs *SaleRentalSchedule) error {
	return c.Update(SaleRentalScheduleModel, ids, srs, nil)
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
	return &((*srss)[0]), nil
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
	return &((*srss)[0]), nil
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
	return c.Search(SaleRentalScheduleModel, criteria, options)
}

// FindSaleRentalScheduleId finds record id by querying it with criteria.
func (c *Client) FindSaleRentalScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleRentalScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
