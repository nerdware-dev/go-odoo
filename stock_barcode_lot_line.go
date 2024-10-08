package odoo

// StockBarcodeLotLine represents stock_barcode.lot.line model.
type StockBarcodeLotLine struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	LotName           *String   `xmlrpc:"lot_name,omitempty"`
	MoveLineId        *Many2One `xmlrpc:"move_line_id,omitempty"`
	ProductBarcode    *String   `xmlrpc:"product_barcode,omitempty"`
	QtyDone           *Float    `xmlrpc:"qty_done,omitempty"`
	QtyReserved       *Float    `xmlrpc:"qty_reserved,omitempty"`
	StockBarcodeLotId *Many2One `xmlrpc:"stock_barcode_lot_id,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockBarcodeLotLines represents array of stock_barcode.lot.line model.
type StockBarcodeLotLines []StockBarcodeLotLine

// StockBarcodeLotLineModel is the odoo model name.
const StockBarcodeLotLineModel = "stock_barcode.lot.line"

// Many2One convert StockBarcodeLotLine to *Many2One.
func (sll *StockBarcodeLotLine) Many2One() *Many2One {
	return NewMany2One(sll.Id.Get(), "")
}

// CreateStockBarcodeLotLine creates a new stock_barcode.lot.line model and returns its id.
func (c *Client) CreateStockBarcodeLotLine(sll *StockBarcodeLotLine) (int64, error) {
	ids, err := c.CreateStockBarcodeLotLines([]*StockBarcodeLotLine{sll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockBarcodeLotLine creates a new stock_barcode.lot.line model and returns its id.
func (c *Client) CreateStockBarcodeLotLines(slls []*StockBarcodeLotLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range slls {
		vv = append(vv, v)
	}
	return c.Create(StockBarcodeLotLineModel, vv, nil)
}

// UpdateStockBarcodeLotLine updates an existing stock_barcode.lot.line record.
func (c *Client) UpdateStockBarcodeLotLine(sll *StockBarcodeLotLine) error {
	return c.UpdateStockBarcodeLotLines([]int64{sll.Id.Get()}, sll)
}

// UpdateStockBarcodeLotLines updates existing stock_barcode.lot.line records.
// All records (represented by ids) will be updated by sll values.
func (c *Client) UpdateStockBarcodeLotLines(ids []int64, sll *StockBarcodeLotLine) error {
	return c.Update(StockBarcodeLotLineModel, ids, sll, nil)
}

// DeleteStockBarcodeLotLine deletes an existing stock_barcode.lot.line record.
func (c *Client) DeleteStockBarcodeLotLine(id int64) error {
	return c.DeleteStockBarcodeLotLines([]int64{id})
}

// DeleteStockBarcodeLotLines deletes existing stock_barcode.lot.line records.
func (c *Client) DeleteStockBarcodeLotLines(ids []int64) error {
	return c.Delete(StockBarcodeLotLineModel, ids)
}

// GetStockBarcodeLotLine gets stock_barcode.lot.line existing record.
func (c *Client) GetStockBarcodeLotLine(id int64) (*StockBarcodeLotLine, error) {
	slls, err := c.GetStockBarcodeLotLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slls)[0]), nil
}

// GetStockBarcodeLotLines gets stock_barcode.lot.line existing records.
func (c *Client) GetStockBarcodeLotLines(ids []int64) (*StockBarcodeLotLines, error) {
	slls := &StockBarcodeLotLines{}
	if err := c.Read(StockBarcodeLotLineModel, ids, nil, slls); err != nil {
		return nil, err
	}
	return slls, nil
}

// FindStockBarcodeLotLine finds stock_barcode.lot.line record by querying it with criteria.
func (c *Client) FindStockBarcodeLotLine(criteria *Criteria) (*StockBarcodeLotLine, error) {
	slls := &StockBarcodeLotLines{}
	if err := c.SearchRead(StockBarcodeLotLineModel, criteria, NewOptions().Limit(1), slls); err != nil {
		return nil, err
	}
	return &((*slls)[0]), nil
}

// FindStockBarcodeLotLines finds stock_barcode.lot.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeLotLines(criteria *Criteria, options *Options) (*StockBarcodeLotLines, error) {
	slls := &StockBarcodeLotLines{}
	if err := c.SearchRead(StockBarcodeLotLineModel, criteria, options, slls); err != nil {
		return nil, err
	}
	return slls, nil
}

// FindStockBarcodeLotLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeLotLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockBarcodeLotLineModel, criteria, options)
}

// FindStockBarcodeLotLineId finds record id by querying it with criteria.
func (c *Client) FindStockBarcodeLotLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBarcodeLotLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
