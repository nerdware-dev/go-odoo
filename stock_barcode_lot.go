package odoo

// StockBarcodeLot represents stock_barcode.lot model.
type StockBarcodeLot struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omitempty"`
	BarcodeScanned         *String   `xmlrpc:"_barcode_scanned,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultMoveId          *Many2One `xmlrpc:"default_move_id,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	PickingId              *Many2One `xmlrpc:"picking_id,omitempty"`
	ProductId              *Many2One `xmlrpc:"product_id,omitempty"`
	QtyDone                *Float    `xmlrpc:"qty_done,omitempty"`
	QtyReserved            *Float    `xmlrpc:"qty_reserved,omitempty"`
	StockBarcodeLotLineIds *Relation `xmlrpc:"stock_barcode_lot_line_ids,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockBarcodeLots represents array of stock_barcode.lot model.
type StockBarcodeLots []StockBarcodeLot

// StockBarcodeLotModel is the odoo model name.
const StockBarcodeLotModel = "stock_barcode.lot"

// Many2One convert StockBarcodeLot to *Many2One.
func (sl *StockBarcodeLot) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateStockBarcodeLot creates a new stock_barcode.lot model and returns its id.
func (c *Client) CreateStockBarcodeLot(sl *StockBarcodeLot) (int64, error) {
	ids, err := c.CreateStockBarcodeLots([]*StockBarcodeLot{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockBarcodeLot creates a new stock_barcode.lot model and returns its id.
func (c *Client) CreateStockBarcodeLots(sls []*StockBarcodeLot) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(StockBarcodeLotModel, vv, nil)
}

// UpdateStockBarcodeLot updates an existing stock_barcode.lot record.
func (c *Client) UpdateStockBarcodeLot(sl *StockBarcodeLot) error {
	return c.UpdateStockBarcodeLots([]int64{sl.Id.Get()}, sl)
}

// UpdateStockBarcodeLots updates existing stock_barcode.lot records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockBarcodeLots(ids []int64, sl *StockBarcodeLot) error {
	return c.Update(StockBarcodeLotModel, ids, sl, nil)
}

// DeleteStockBarcodeLot deletes an existing stock_barcode.lot record.
func (c *Client) DeleteStockBarcodeLot(id int64) error {
	return c.DeleteStockBarcodeLots([]int64{id})
}

// DeleteStockBarcodeLots deletes existing stock_barcode.lot records.
func (c *Client) DeleteStockBarcodeLots(ids []int64) error {
	return c.Delete(StockBarcodeLotModel, ids)
}

// GetStockBarcodeLot gets stock_barcode.lot existing record.
func (c *Client) GetStockBarcodeLot(id int64) (*StockBarcodeLot, error) {
	sls, err := c.GetStockBarcodeLots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// GetStockBarcodeLots gets stock_barcode.lot existing records.
func (c *Client) GetStockBarcodeLots(ids []int64) (*StockBarcodeLots, error) {
	sls := &StockBarcodeLots{}
	if err := c.Read(StockBarcodeLotModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockBarcodeLot finds stock_barcode.lot record by querying it with criteria.
func (c *Client) FindStockBarcodeLot(criteria *Criteria) (*StockBarcodeLot, error) {
	sls := &StockBarcodeLots{}
	if err := c.SearchRead(StockBarcodeLotModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// FindStockBarcodeLots finds stock_barcode.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeLots(criteria *Criteria, options *Options) (*StockBarcodeLots, error) {
	sls := &StockBarcodeLots{}
	if err := c.SearchRead(StockBarcodeLotModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockBarcodeLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockBarcodeLotModel, criteria, options)
}

// FindStockBarcodeLotId finds record id by querying it with criteria.
func (c *Client) FindStockBarcodeLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBarcodeLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
