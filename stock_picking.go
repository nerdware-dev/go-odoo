package odoo

// StockPicking represents stock.picking model.
type StockPicking struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	BarcodeScanned              *String    `xmlrpc:"_barcode_scanned,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BackorderId                 *Many2One  `xmlrpc:"backorder_id,omitempty"`
	BackorderIds                *Relation  `xmlrpc:"backorder_ids,omitempty"`
	CarrierId                   *Many2One  `xmlrpc:"carrier_id,omitempty"`
	CarrierPrice                *Float     `xmlrpc:"carrier_price,omitempty"`
	CarrierTrackingRef          *String    `xmlrpc:"carrier_tracking_ref,omitempty"`
	CarrierTrackingUrl          *String    `xmlrpc:"carrier_tracking_url,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	DateDone                    *Time      `xmlrpc:"date_done,omitempty"`
	DeliveryPackagingIds        *Relation  `xmlrpc:"delivery_packaging_ids,omitempty"`
	DeliveryType                *Selection `xmlrpc:"delivery_type,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	GroupId                     *Many2One  `xmlrpc:"group_id,omitempty"`
	HasPackages                 *Bool      `xmlrpc:"has_packages,omitempty"`
	HasScrapMove                *Bool      `xmlrpc:"has_scrap_move,omitempty"`
	HasTracking                 *Bool      `xmlrpc:"has_tracking,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	ImmediateTransfer           *Bool      `xmlrpc:"immediate_transfer,omitempty"`
	IsLocked                    *Bool      `xmlrpc:"is_locked,omitempty"`
	IsReturnPicking             *Bool      `xmlrpc:"is_return_picking,omitempty"`
	LocationDestId              *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId                  *Many2One  `xmlrpc:"location_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveIdsWithoutPackage       *Relation  `xmlrpc:"move_ids_without_package,omitempty"`
	MoveLineExist               *Bool      `xmlrpc:"move_line_exist,omitempty"`
	MoveLineIds                 *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveLineIdsWithoutPackage   *Relation  `xmlrpc:"move_line_ids_without_package,omitempty"`
	MoveLineNosuggestIds        *Relation  `xmlrpc:"move_line_nosuggest_ids,omitempty"`
	MoveLines                   *Relation  `xmlrpc:"move_lines,omitempty"`
	MoveType                    *Selection `xmlrpc:"move_type,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	Origin                      *String    `xmlrpc:"origin,omitempty"`
	OwnerId                     *Many2One  `xmlrpc:"owner_id,omitempty"`
	PackageIds                  *Relation  `xmlrpc:"package_ids,omitempty"`
	PackageLevelIds             *Relation  `xmlrpc:"package_level_ids,omitempty"`
	PackageLevelIdsDetails      *Relation  `xmlrpc:"package_level_ids_details,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingTypeCode             *Selection `xmlrpc:"picking_type_code,omitempty"`
	PickingTypeEntirePacks      *Bool      `xmlrpc:"picking_type_entire_packs,omitempty"`
	PickingTypeId               *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	Printed                     *Bool      `xmlrpc:"printed,omitempty"`
	Priority                    *Selection `xmlrpc:"priority,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	PurchaseId                  *Many2One  `xmlrpc:"purchase_id,omitempty"`
	ReportComputeDate           *Time      `xmlrpc:"report_compute_date,omitempty"`
	ReportLineIndex             *Int       `xmlrpc:"report_line_index,omitempty"`
	ReturnLabelIds              *Relation  `xmlrpc:"return_label_ids,omitempty"`
	SaleId                      *Many2One  `xmlrpc:"sale_id,omitempty"`
	ScheduledDate               *Time      `xmlrpc:"scheduled_date,omitempty"`
	ShippingWeight              *Float     `xmlrpc:"shipping_weight,omitempty"`
	ShowCheckAvailability       *Bool      `xmlrpc:"show_check_availability,omitempty"`
	ShowLotsText                *Bool      `xmlrpc:"show_lots_text,omitempty"`
	ShowMarkAsTodo              *Bool      `xmlrpc:"show_mark_as_todo,omitempty"`
	ShowOperations              *Bool      `xmlrpc:"show_operations,omitempty"`
	ShowReserved                *Bool      `xmlrpc:"show_reserved,omitempty"`
	ShowValidate                *Bool      `xmlrpc:"show_validate,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	UseCreateLots               *Bool      `xmlrpc:"use_create_lots,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	Volume                      *Float     `xmlrpc:"volume,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	Weight                      *Float     `xmlrpc:"weight,omitempty"`
	WeightBulk                  *Float     `xmlrpc:"weight_bulk,omitempty"`
	WeightUomName               *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockPickings represents array of stock.picking model.
type StockPickings []StockPicking

// StockPickingModel is the odoo model name.
const StockPickingModel = "stock.picking"

// Many2One convert StockPicking to *Many2One.
func (sp *StockPicking) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPicking creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPicking(sp *StockPicking) (int64, error) {
	ids, err := c.CreateStockPickings([]*StockPicking{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPicking creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPickings(sps []*StockPicking) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(StockPickingModel, vv, nil)
}

// UpdateStockPicking updates an existing stock.picking record.
func (c *Client) UpdateStockPicking(sp *StockPicking) error {
	return c.UpdateStockPickings([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPickings updates existing stock.picking records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPickings(ids []int64, sp *StockPicking) error {
	return c.Update(StockPickingModel, ids, sp, nil)
}

// DeleteStockPicking deletes an existing stock.picking record.
func (c *Client) DeleteStockPicking(id int64) error {
	return c.DeleteStockPickings([]int64{id})
}

// DeleteStockPickings deletes existing stock.picking records.
func (c *Client) DeleteStockPickings(ids []int64) error {
	return c.Delete(StockPickingModel, ids)
}

// GetStockPicking gets stock.picking existing record.
func (c *Client) GetStockPicking(id int64) (*StockPicking, error) {
	sps, err := c.GetStockPickings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// GetStockPickings gets stock.picking existing records.
func (c *Client) GetStockPickings(ids []int64) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.Read(StockPickingModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPicking finds stock.picking record by querying it with criteria.
func (c *Client) FindStockPicking(criteria *Criteria) (*StockPicking, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// FindStockPickings finds stock.picking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickings(criteria *Criteria, options *Options) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPickingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingModel, criteria, options)
}

// FindStockPickingId finds record id by querying it with criteria.
func (c *Client) FindStockPickingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
