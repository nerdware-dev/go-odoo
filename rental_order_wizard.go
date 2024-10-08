package odoo

// RentalOrderWizard represents rental.order.wizard model.
type RentalOrderWizard struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	HasLateLines         *Bool      `xmlrpc:"has_late_lines,omitempty"`
	HasLinesMissingStock *Bool      `xmlrpc:"has_lines_missing_stock,omitempty"`
	HasTrackedLines      *Bool      `xmlrpc:"has_tracked_lines,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	OrderId              *Many2One  `xmlrpc:"order_id,omitempty"`
	RentalWizardLineIds  *Relation  `xmlrpc:"rental_wizard_line_ids,omitempty"`
	Status               *Selection `xmlrpc:"status,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// RentalOrderWizards represents array of rental.order.wizard model.
type RentalOrderWizards []RentalOrderWizard

// RentalOrderWizardModel is the odoo model name.
const RentalOrderWizardModel = "rental.order.wizard"

// Many2One convert RentalOrderWizard to *Many2One.
func (row *RentalOrderWizard) Many2One() *Many2One {
	return NewMany2One(row.Id.Get(), "")
}

// CreateRentalOrderWizard creates a new rental.order.wizard model and returns its id.
func (c *Client) CreateRentalOrderWizard(row *RentalOrderWizard) (int64, error) {
	ids, err := c.CreateRentalOrderWizards([]*RentalOrderWizard{row})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRentalOrderWizard creates a new rental.order.wizard model and returns its id.
func (c *Client) CreateRentalOrderWizards(rows []*RentalOrderWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range rows {
		vv = append(vv, v)
	}
	return c.Create(RentalOrderWizardModel, vv, nil)
}

// UpdateRentalOrderWizard updates an existing rental.order.wizard record.
func (c *Client) UpdateRentalOrderWizard(row *RentalOrderWizard) error {
	return c.UpdateRentalOrderWizards([]int64{row.Id.Get()}, row)
}

// UpdateRentalOrderWizards updates existing rental.order.wizard records.
// All records (represented by ids) will be updated by row values.
func (c *Client) UpdateRentalOrderWizards(ids []int64, row *RentalOrderWizard) error {
	return c.Update(RentalOrderWizardModel, ids, row, nil)
}

// DeleteRentalOrderWizard deletes an existing rental.order.wizard record.
func (c *Client) DeleteRentalOrderWizard(id int64) error {
	return c.DeleteRentalOrderWizards([]int64{id})
}

// DeleteRentalOrderWizards deletes existing rental.order.wizard records.
func (c *Client) DeleteRentalOrderWizards(ids []int64) error {
	return c.Delete(RentalOrderWizardModel, ids)
}

// GetRentalOrderWizard gets rental.order.wizard existing record.
func (c *Client) GetRentalOrderWizard(id int64) (*RentalOrderWizard, error) {
	rows, err := c.GetRentalOrderWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rows)[0]), nil
}

// GetRentalOrderWizards gets rental.order.wizard existing records.
func (c *Client) GetRentalOrderWizards(ids []int64) (*RentalOrderWizards, error) {
	rows := &RentalOrderWizards{}
	if err := c.Read(RentalOrderWizardModel, ids, nil, rows); err != nil {
		return nil, err
	}
	return rows, nil
}

// FindRentalOrderWizard finds rental.order.wizard record by querying it with criteria.
func (c *Client) FindRentalOrderWizard(criteria *Criteria) (*RentalOrderWizard, error) {
	rows := &RentalOrderWizards{}
	if err := c.SearchRead(RentalOrderWizardModel, criteria, NewOptions().Limit(1), rows); err != nil {
		return nil, err
	}
	return &((*rows)[0]), nil
}

// FindRentalOrderWizards finds rental.order.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalOrderWizards(criteria *Criteria, options *Options) (*RentalOrderWizards, error) {
	rows := &RentalOrderWizards{}
	if err := c.SearchRead(RentalOrderWizardModel, criteria, options, rows); err != nil {
		return nil, err
	}
	return rows, nil
}

// FindRentalOrderWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalOrderWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RentalOrderWizardModel, criteria, options)
}

// FindRentalOrderWizardId finds record id by querying it with criteria.
func (c *Client) FindRentalOrderWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RentalOrderWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
