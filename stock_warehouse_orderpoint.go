package odoo

// StockWarehouseOrderpoint represents stock.warehouse.orderpoint model.
type StockWarehouseOrderpoint struct {
	Active                 *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	AllowedLocationIds     *Relation  `xmlrpc:"allowed_location_ids,omitempty" json:"allowed_location_ids,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DaysToOrder            *Float     `xmlrpc:"days_to_order,omitempty" json:"days_to_order,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omitempty" json:"group_id,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	LeadDaysDate           *Time      `xmlrpc:"lead_days_date,omitempty" json:"lead_days_date,omitempty"`
	LocationId             *Many2One  `xmlrpc:"location_id,omitempty" json:"location_id,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	ProductCategoryId      *Many2One  `xmlrpc:"product_category_id,omitempty" json:"product_category_id,omitempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductMaxQty          *Float     `xmlrpc:"product_max_qty,omitempty" json:"product_max_qty,omitempty"`
	ProductMinQty          *Float     `xmlrpc:"product_min_qty,omitempty" json:"product_min_qty,omitempty"`
	ProductTmplId          *Many2One  `xmlrpc:"product_tmpl_id,omitempty" json:"product_tmpl_id,omitempty"`
	ProductUom             *Many2One  `xmlrpc:"product_uom,omitempty" json:"product_uom,omitempty"`
	ProductUomName         *String    `xmlrpc:"product_uom_name,omitempty" json:"product_uom_name,omitempty"`
	PurchaseVisibilityDays *Float     `xmlrpc:"purchase_visibility_days,omitempty" json:"purchase_visibility_days,omitempty"`
	QtyForecast            *Float     `xmlrpc:"qty_forecast,omitempty" json:"qty_forecast,omitempty"`
	QtyMultiple            *Float     `xmlrpc:"qty_multiple,omitempty" json:"qty_multiple,omitempty"`
	QtyOnHand              *Float     `xmlrpc:"qty_on_hand,omitempty" json:"qty_on_hand,omitempty"`
	QtyToOrder             *Float     `xmlrpc:"qty_to_order,omitempty" json:"qty_to_order,omitempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omitempty" json:"route_id,omitempty"`
	RuleIds                *Relation  `xmlrpc:"rule_ids,omitempty" json:"rule_ids,omitempty"`
	ShowSupplier           *Bool      `xmlrpc:"show_supplier,omitempty" json:"show_supplier,omitempty"`
	SnoozedUntil           *Time      `xmlrpc:"snoozed_until,omitempty" json:"snoozed_until,omitempty"`
	SupplierId             *Many2One  `xmlrpc:"supplier_id,omitempty" json:"supplier_id,omitempty"`
	Trigger                *Selection `xmlrpc:"trigger,omitempty" json:"trigger,omitempty"`
	UnwantedReplenish      *Bool      `xmlrpc:"unwanted_replenish,omitempty" json:"unwanted_replenish,omitempty"`
	VendorId               *Many2One  `xmlrpc:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	VisibilityDays         *Float     `xmlrpc:"visibility_days,omitempty" json:"visibility_days,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// StockWarehouseOrderpoints represents array of stock.warehouse.orderpoint model.
type StockWarehouseOrderpoints []StockWarehouseOrderpoint

// StockWarehouseOrderpointModel is the odoo model name.
const StockWarehouseOrderpointModel = "stock.warehouse.orderpoint"

// Many2One convert StockWarehouseOrderpoint to *Many2One.
func (swo *StockWarehouseOrderpoint) Many2One() *Many2One {
	return NewMany2One(swo.Id.Get(), "")
}

// CreateStockWarehouseOrderpoint creates a new stock.warehouse.orderpoint model and returns its id.
func (c *Client) CreateStockWarehouseOrderpoint(swo *StockWarehouseOrderpoint) (int64, error) {
	ids, err := c.CreateStockWarehouseOrderpoints([]*StockWarehouseOrderpoint{swo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockWarehouseOrderpoint creates a new stock.warehouse.orderpoint model and returns its id.
func (c *Client) CreateStockWarehouseOrderpoints(swos []*StockWarehouseOrderpoint) ([]int64, error) {
	var vv []interface{}
	for _, v := range swos {
		vv = append(vv, v)
	}
	return c.Create(StockWarehouseOrderpointModel, vv, nil)
}

// UpdateStockWarehouseOrderpoint updates an existing stock.warehouse.orderpoint record.
func (c *Client) UpdateStockWarehouseOrderpoint(swo *StockWarehouseOrderpoint) error {
	return c.UpdateStockWarehouseOrderpoints([]int64{swo.Id.Get()}, swo)
}

// UpdateStockWarehouseOrderpoints updates existing stock.warehouse.orderpoint records.
// All records (represented by ids) will be updated by swo values.
func (c *Client) UpdateStockWarehouseOrderpoints(ids []int64, swo *StockWarehouseOrderpoint) error {
	return c.Update(StockWarehouseOrderpointModel, ids, swo, nil)
}

// DeleteStockWarehouseOrderpoint deletes an existing stock.warehouse.orderpoint record.
func (c *Client) DeleteStockWarehouseOrderpoint(id int64) error {
	return c.DeleteStockWarehouseOrderpoints([]int64{id})
}

// DeleteStockWarehouseOrderpoints deletes existing stock.warehouse.orderpoint records.
func (c *Client) DeleteStockWarehouseOrderpoints(ids []int64) error {
	return c.Delete(StockWarehouseOrderpointModel, ids)
}

// GetStockWarehouseOrderpoint gets stock.warehouse.orderpoint existing record.
func (c *Client) GetStockWarehouseOrderpoint(id int64) (*StockWarehouseOrderpoint, error) {
	swos, err := c.GetStockWarehouseOrderpoints([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*swos)[0]), nil
}

// GetStockWarehouseOrderpoints gets stock.warehouse.orderpoint existing records.
func (c *Client) GetStockWarehouseOrderpoints(ids []int64) (*StockWarehouseOrderpoints, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.Read(StockWarehouseOrderpointModel, ids, nil, swos); err != nil {
		return nil, err
	}
	return swos, nil
}

// FindStockWarehouseOrderpoint finds stock.warehouse.orderpoint record by querying it with criteria.
func (c *Client) FindStockWarehouseOrderpoint(criteria *Criteria) (*StockWarehouseOrderpoint, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.SearchRead(StockWarehouseOrderpointModel, criteria, NewOptions().Limit(1), swos); err != nil {
		return nil, err
	}
	return &((*swos)[0]), nil
}

// FindStockWarehouseOrderpoints finds stock.warehouse.orderpoint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouseOrderpoints(criteria *Criteria, options *Options) (*StockWarehouseOrderpoints, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.SearchRead(StockWarehouseOrderpointModel, criteria, options, swos); err != nil {
		return nil, err
	}
	return swos, nil
}

// FindStockWarehouseOrderpointIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouseOrderpointIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockWarehouseOrderpointModel, criteria, options)
}

// FindStockWarehouseOrderpointId finds record id by querying it with criteria.
func (c *Client) FindStockWarehouseOrderpointId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarehouseOrderpointModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
