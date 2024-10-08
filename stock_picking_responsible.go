package odoo

// StockPickingResponsible represents stock.picking.responsible model.
type StockPickingResponsible struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockPickingResponsibles represents array of stock.picking.responsible model.
type StockPickingResponsibles []StockPickingResponsible

// StockPickingResponsibleModel is the odoo model name.
const StockPickingResponsibleModel = "stock.picking.responsible"

// Many2One convert StockPickingResponsible to *Many2One.
func (spr *StockPickingResponsible) Many2One() *Many2One {
	return NewMany2One(spr.Id.Get(), "")
}

// CreateStockPickingResponsible creates a new stock.picking.responsible model and returns its id.
func (c *Client) CreateStockPickingResponsible(spr *StockPickingResponsible) (int64, error) {
	ids, err := c.CreateStockPickingResponsibles([]*StockPickingResponsible{spr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPickingResponsible creates a new stock.picking.responsible model and returns its id.
func (c *Client) CreateStockPickingResponsibles(sprs []*StockPickingResponsible) ([]int64, error) {
	var vv []interface{}
	for _, v := range sprs {
		vv = append(vv, v)
	}
	return c.Create(StockPickingResponsibleModel, vv, nil)
}

// UpdateStockPickingResponsible updates an existing stock.picking.responsible record.
func (c *Client) UpdateStockPickingResponsible(spr *StockPickingResponsible) error {
	return c.UpdateStockPickingResponsibles([]int64{spr.Id.Get()}, spr)
}

// UpdateStockPickingResponsibles updates existing stock.picking.responsible records.
// All records (represented by ids) will be updated by spr values.
func (c *Client) UpdateStockPickingResponsibles(ids []int64, spr *StockPickingResponsible) error {
	return c.Update(StockPickingResponsibleModel, ids, spr, nil)
}

// DeleteStockPickingResponsible deletes an existing stock.picking.responsible record.
func (c *Client) DeleteStockPickingResponsible(id int64) error {
	return c.DeleteStockPickingResponsibles([]int64{id})
}

// DeleteStockPickingResponsibles deletes existing stock.picking.responsible records.
func (c *Client) DeleteStockPickingResponsibles(ids []int64) error {
	return c.Delete(StockPickingResponsibleModel, ids)
}

// GetStockPickingResponsible gets stock.picking.responsible existing record.
func (c *Client) GetStockPickingResponsible(id int64) (*StockPickingResponsible, error) {
	sprs, err := c.GetStockPickingResponsibles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sprs)[0]), nil
}

// GetStockPickingResponsibles gets stock.picking.responsible existing records.
func (c *Client) GetStockPickingResponsibles(ids []int64) (*StockPickingResponsibles, error) {
	sprs := &StockPickingResponsibles{}
	if err := c.Read(StockPickingResponsibleModel, ids, nil, sprs); err != nil {
		return nil, err
	}
	return sprs, nil
}

// FindStockPickingResponsible finds stock.picking.responsible record by querying it with criteria.
func (c *Client) FindStockPickingResponsible(criteria *Criteria) (*StockPickingResponsible, error) {
	sprs := &StockPickingResponsibles{}
	if err := c.SearchRead(StockPickingResponsibleModel, criteria, NewOptions().Limit(1), sprs); err != nil {
		return nil, err
	}
	return &((*sprs)[0]), nil
}

// FindStockPickingResponsibles finds stock.picking.responsible records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingResponsibles(criteria *Criteria, options *Options) (*StockPickingResponsibles, error) {
	sprs := &StockPickingResponsibles{}
	if err := c.SearchRead(StockPickingResponsibleModel, criteria, options, sprs); err != nil {
		return nil, err
	}
	return sprs, nil
}

// FindStockPickingResponsibleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingResponsibleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingResponsibleModel, criteria, options)
}

// FindStockPickingResponsibleId finds record id by querying it with criteria.
func (c *Client) FindStockPickingResponsibleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingResponsibleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
