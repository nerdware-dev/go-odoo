package odoo

import (
	"fmt"
)

// StockBarcodeLot represents stock_barcode.lot model.
type StockBarcodeLot struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	BarcodeScanned         *String   `xmlrpc:"_barcode_scanned,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DefaultMoveId          *Many2One `xmlrpc:"default_move_id,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	PickingId              *Many2One `xmlrpc:"picking_id,omptempty"`
	ProductId              *Many2One `xmlrpc:"product_id,omptempty"`
	QtyDone                *Float    `xmlrpc:"qty_done,omptempty"`
	QtyReserved            *Float    `xmlrpc:"qty_reserved,omptempty"`
	StockBarcodeLotLineIds *Relation `xmlrpc:"stock_barcode_lot_line_ids,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(StockBarcodeLotModel, vv)
}

// UpdateStockBarcodeLot updates an existing stock_barcode.lot record.
func (c *Client) UpdateStockBarcodeLot(sl *StockBarcodeLot) error {
	return c.UpdateStockBarcodeLots([]int64{sl.Id.Get()}, sl)
}

// UpdateStockBarcodeLots updates existing stock_barcode.lot records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockBarcodeLots(ids []int64, sl *StockBarcodeLot) error {
	return c.Update(StockBarcodeLotModel, ids, sl)
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
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock_barcode.lot not found", id)
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
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("stock_barcode.lot was not found with criteria %v", criteria)
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
	ids, err := c.Search(StockBarcodeLotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockBarcodeLotId finds record id by querying it with criteria.
func (c *Client) FindStockBarcodeLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBarcodeLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock_barcode.lot was not found with criteria %v and options %v", criteria, options)
}
