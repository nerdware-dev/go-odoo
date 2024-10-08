package odoo

// RentalOrderWizardLine represents rental.order.wizard.line model.
type RentalOrderWizardLine struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IsLate              *Bool      `xmlrpc:"is_late,omitempty"`
	IsProductStorable   *Bool      `xmlrpc:"is_product_storable,omitempty"`
	OrderLineId         *Many2One  `xmlrpc:"order_line_id,omitempty"`
	PickeableLotIds     *Relation  `xmlrpc:"pickeable_lot_ids,omitempty"`
	PickedupLotIds      *Relation  `xmlrpc:"pickedup_lot_ids,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty"`
	QtyAvailable        *Float     `xmlrpc:"qty_available,omitempty"`
	QtyDelivered        *Float     `xmlrpc:"qty_delivered,omitempty"`
	QtyReserved         *Float     `xmlrpc:"qty_reserved,omitempty"`
	QtyReturned         *Float     `xmlrpc:"qty_returned,omitempty"`
	RentalOrderWizardId *Many2One  `xmlrpc:"rental_order_wizard_id,omitempty"`
	ReturnableLotIds    *Relation  `xmlrpc:"returnable_lot_ids,omitempty"`
	ReturnedLotIds      *Relation  `xmlrpc:"returned_lot_ids,omitempty"`
	Status              *Selection `xmlrpc:"status,omitempty"`
	Tracking            *Selection `xmlrpc:"tracking,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// RentalOrderWizardLines represents array of rental.order.wizard.line model.
type RentalOrderWizardLines []RentalOrderWizardLine

// RentalOrderWizardLineModel is the odoo model name.
const RentalOrderWizardLineModel = "rental.order.wizard.line"

// Many2One convert RentalOrderWizardLine to *Many2One.
func (rowl *RentalOrderWizardLine) Many2One() *Many2One {
	return NewMany2One(rowl.Id.Get(), "")
}

// CreateRentalOrderWizardLine creates a new rental.order.wizard.line model and returns its id.
func (c *Client) CreateRentalOrderWizardLine(rowl *RentalOrderWizardLine) (int64, error) {
	ids, err := c.CreateRentalOrderWizardLines([]*RentalOrderWizardLine{rowl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRentalOrderWizardLine creates a new rental.order.wizard.line model and returns its id.
func (c *Client) CreateRentalOrderWizardLines(rowls []*RentalOrderWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range rowls {
		vv = append(vv, v)
	}
	return c.Create(RentalOrderWizardLineModel, vv, nil)
}

// UpdateRentalOrderWizardLine updates an existing rental.order.wizard.line record.
func (c *Client) UpdateRentalOrderWizardLine(rowl *RentalOrderWizardLine) error {
	return c.UpdateRentalOrderWizardLines([]int64{rowl.Id.Get()}, rowl)
}

// UpdateRentalOrderWizardLines updates existing rental.order.wizard.line records.
// All records (represented by ids) will be updated by rowl values.
func (c *Client) UpdateRentalOrderWizardLines(ids []int64, rowl *RentalOrderWizardLine) error {
	return c.Update(RentalOrderWizardLineModel, ids, rowl, nil)
}

// DeleteRentalOrderWizardLine deletes an existing rental.order.wizard.line record.
func (c *Client) DeleteRentalOrderWizardLine(id int64) error {
	return c.DeleteRentalOrderWizardLines([]int64{id})
}

// DeleteRentalOrderWizardLines deletes existing rental.order.wizard.line records.
func (c *Client) DeleteRentalOrderWizardLines(ids []int64) error {
	return c.Delete(RentalOrderWizardLineModel, ids)
}

// GetRentalOrderWizardLine gets rental.order.wizard.line existing record.
func (c *Client) GetRentalOrderWizardLine(id int64) (*RentalOrderWizardLine, error) {
	rowls, err := c.GetRentalOrderWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rowls)[0]), nil
}

// GetRentalOrderWizardLines gets rental.order.wizard.line existing records.
func (c *Client) GetRentalOrderWizardLines(ids []int64) (*RentalOrderWizardLines, error) {
	rowls := &RentalOrderWizardLines{}
	if err := c.Read(RentalOrderWizardLineModel, ids, nil, rowls); err != nil {
		return nil, err
	}
	return rowls, nil
}

// FindRentalOrderWizardLine finds rental.order.wizard.line record by querying it with criteria.
func (c *Client) FindRentalOrderWizardLine(criteria *Criteria) (*RentalOrderWizardLine, error) {
	rowls := &RentalOrderWizardLines{}
	if err := c.SearchRead(RentalOrderWizardLineModel, criteria, NewOptions().Limit(1), rowls); err != nil {
		return nil, err
	}
	return &((*rowls)[0]), nil
}

// FindRentalOrderWizardLines finds rental.order.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalOrderWizardLines(criteria *Criteria, options *Options) (*RentalOrderWizardLines, error) {
	rowls := &RentalOrderWizardLines{}
	if err := c.SearchRead(RentalOrderWizardLineModel, criteria, options, rowls); err != nil {
		return nil, err
	}
	return rowls, nil
}

// FindRentalOrderWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalOrderWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RentalOrderWizardLineModel, criteria, options)
}

// FindRentalOrderWizardLineId finds record id by querying it with criteria.
func (c *Client) FindRentalOrderWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RentalOrderWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
