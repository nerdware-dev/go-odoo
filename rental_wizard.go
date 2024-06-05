package odoo

import (
	"fmt"
)

// RentalWizard represents rental.wizard model.
type RentalWizard struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	Duration                 *Int       `xmlrpc:"duration,omptempty"`
	DurationUnit             *Selection `xmlrpc:"duration_unit,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsProductStorable        *Bool      `xmlrpc:"is_product_storable,omptempty"`
	LotIds                   *Relation  `xmlrpc:"lot_ids,omptempty"`
	PickupDate               *Time      `xmlrpc:"pickup_date,omptempty"`
	PricelistId              *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	PricingExplanation       *String    `xmlrpc:"pricing_explanation,omptempty"`
	PricingId                *Many2One  `xmlrpc:"pricing_id,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomId             *String    `xmlrpc:"product_uom_id,omptempty"`
	QtyAvailableDuringPeriod *Float     `xmlrpc:"qty_available_during_period,omptempty"`
	Quantity                 *Float     `xmlrpc:"quantity,omptempty"`
	RentableLotIds           *Relation  `xmlrpc:"rentable_lot_ids,omptempty"`
	RentableQty              *Float     `xmlrpc:"rentable_qty,omptempty"`
	RentalOrderLineId        *Many2One  `xmlrpc:"rental_order_line_id,omptempty"`
	RentedLotIds             *Relation  `xmlrpc:"rented_lot_ids,omptempty"`
	RentedQtyDuringPeriod    *Float     `xmlrpc:"rented_qty_during_period,omptempty"`
	ReturnDate               *Time      `xmlrpc:"return_date,omptempty"`
	Tracking                 *Selection `xmlrpc:"tracking,omptempty"`
	UnitPrice                *Float     `xmlrpc:"unit_price,omptempty"`
	UomId                    *Many2One  `xmlrpc:"uom_id,omptempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// RentalWizards represents array of rental.wizard model.
type RentalWizards []RentalWizard

// RentalWizardModel is the odoo model name.
const RentalWizardModel = "rental.wizard"

// Many2One convert RentalWizard to *Many2One.
func (rw *RentalWizard) Many2One() *Many2One {
	return NewMany2One(rw.Id.Get(), "")
}

// CreateRentalWizard creates a new rental.wizard model and returns its id.
func (c *Client) CreateRentalWizard(rw *RentalWizard) (int64, error) {
	ids, err := c.CreateRentalWizards([]*RentalWizard{rw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRentalWizard creates a new rental.wizard model and returns its id.
func (c *Client) CreateRentalWizards(rws []*RentalWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range rws {
		vv = append(vv, v)
	}
	return c.Create(RentalWizardModel, vv)
}

// UpdateRentalWizard updates an existing rental.wizard record.
func (c *Client) UpdateRentalWizard(rw *RentalWizard) error {
	return c.UpdateRentalWizards([]int64{rw.Id.Get()}, rw)
}

// UpdateRentalWizards updates existing rental.wizard records.
// All records (represented by ids) will be updated by rw values.
func (c *Client) UpdateRentalWizards(ids []int64, rw *RentalWizard) error {
	return c.Update(RentalWizardModel, ids, rw)
}

// DeleteRentalWizard deletes an existing rental.wizard record.
func (c *Client) DeleteRentalWizard(id int64) error {
	return c.DeleteRentalWizards([]int64{id})
}

// DeleteRentalWizards deletes existing rental.wizard records.
func (c *Client) DeleteRentalWizards(ids []int64) error {
	return c.Delete(RentalWizardModel, ids)
}

// GetRentalWizard gets rental.wizard existing record.
func (c *Client) GetRentalWizard(id int64) (*RentalWizard, error) {
	rws, err := c.GetRentalWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if rws != nil && len(*rws) > 0 {
		return &((*rws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of rental.wizard not found", id)
}

// GetRentalWizards gets rental.wizard existing records.
func (c *Client) GetRentalWizards(ids []int64) (*RentalWizards, error) {
	rws := &RentalWizards{}
	if err := c.Read(RentalWizardModel, ids, nil, rws); err != nil {
		return nil, err
	}
	return rws, nil
}

// FindRentalWizard finds rental.wizard record by querying it with criteria.
func (c *Client) FindRentalWizard(criteria *Criteria) (*RentalWizard, error) {
	rws := &RentalWizards{}
	if err := c.SearchRead(RentalWizardModel, criteria, NewOptions().Limit(1), rws); err != nil {
		return nil, err
	}
	if rws != nil && len(*rws) > 0 {
		return &((*rws)[0]), nil
	}
	return nil, fmt.Errorf("rental.wizard was not found with criteria %v", criteria)
}

// FindRentalWizards finds rental.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalWizards(criteria *Criteria, options *Options) (*RentalWizards, error) {
	rws := &RentalWizards{}
	if err := c.SearchRead(RentalWizardModel, criteria, options, rws); err != nil {
		return nil, err
	}
	return rws, nil
}

// FindRentalWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RentalWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRentalWizardId finds record id by querying it with criteria.
func (c *Client) FindRentalWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RentalWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("rental.wizard was not found with criteria %v and options %v", criteria, options)
}
