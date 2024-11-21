package odoo

// StockPicking represents stock.picking model.
type StockPicking struct {
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	BackorderId                 *Many2One   `xmlrpc:"backorder_id,omitempty" json:"backorder_id,omitempty"`
	BackorderIds                *Relation   `xmlrpc:"backorder_ids,omitempty" json:"backorder_ids,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CountryCode                 *String     `xmlrpc:"country_code,omitempty" json:"country_code,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Date                        *Time       `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DateDeadline                *Time       `xmlrpc:"date_deadline,omitempty" json:"date_deadline,omitempty"`
	DateDone                    *Time       `xmlrpc:"date_done,omitempty" json:"date_done,omitempty"`
	DelayAlertDate              *Time       `xmlrpc:"delay_alert_date,omitempty" json:"delay_alert_date,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	GroupId                     *Many2One   `xmlrpc:"group_id,omitempty" json:"group_id,omitempty"`
	HasDeadlineIssue            *Bool       `xmlrpc:"has_deadline_issue,omitempty" json:"has_deadline_issue,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	HasPackages                 *Bool       `xmlrpc:"has_packages,omitempty" json:"has_packages,omitempty"`
	HasScrapMove                *Bool       `xmlrpc:"has_scrap_move,omitempty" json:"has_scrap_move,omitempty"`
	HasTracking                 *Bool       `xmlrpc:"has_tracking,omitempty" json:"has_tracking,omitempty"`
	HidePickingType             *Bool       `xmlrpc:"hide_picking_type,omitempty" json:"hide_picking_type,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsLocked                    *Bool       `xmlrpc:"is_locked,omitempty" json:"is_locked,omitempty"`
	IsSigned                    *Bool       `xmlrpc:"is_signed,omitempty" json:"is_signed,omitempty"`
	JsonPopover                 *String     `xmlrpc:"json_popover,omitempty" json:"json_popover,omitempty"`
	L10NDin5008Addresses        *String     `xmlrpc:"l10n_din5008_addresses,omitempty" json:"l10n_din5008_addresses,omitempty"`
	LocationDestId              *Many2One   `xmlrpc:"location_dest_id,omitempty" json:"location_dest_id,omitempty"`
	LocationId                  *Many2One   `xmlrpc:"location_id,omitempty" json:"location_id,omitempty"`
	LotId                       *Many2One   `xmlrpc:"lot_id,omitempty" json:"lot_id,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MoveIds                     *Relation   `xmlrpc:"move_ids,omitempty" json:"move_ids,omitempty"`
	MoveIdsWithoutPackage       *Relation   `xmlrpc:"move_ids_without_package,omitempty" json:"move_ids_without_package,omitempty"`
	MoveLineExist               *Bool       `xmlrpc:"move_line_exist,omitempty" json:"move_line_exist,omitempty"`
	MoveLineIds                 *Relation   `xmlrpc:"move_line_ids,omitempty" json:"move_line_ids,omitempty"`
	MoveLineIdsWithoutPackage   *Relation   `xmlrpc:"move_line_ids_without_package,omitempty" json:"move_line_ids_without_package,omitempty"`
	MoveType                    *Selection  `xmlrpc:"move_type,omitempty" json:"move_type,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Note                        *String     `xmlrpc:"note,omitempty" json:"note,omitempty"`
	Origin                      *String     `xmlrpc:"origin,omitempty" json:"origin,omitempty"`
	OwnerId                     *Many2One   `xmlrpc:"owner_id,omitempty" json:"owner_id,omitempty"`
	PackageLevelIds             *Relation   `xmlrpc:"package_level_ids,omitempty" json:"package_level_ids,omitempty"`
	PackageLevelIdsDetails      *Relation   `xmlrpc:"package_level_ids_details,omitempty" json:"package_level_ids_details,omitempty"`
	PartnerId                   *Many2One   `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PickingProperties           interface{} `xmlrpc:"picking_properties,omitempty" json:"picking_properties,omitempty"`
	PickingTypeCode             *Selection  `xmlrpc:"picking_type_code,omitempty" json:"picking_type_code,omitempty"`
	PickingTypeEntirePacks      *Bool       `xmlrpc:"picking_type_entire_packs,omitempty" json:"picking_type_entire_packs,omitempty"`
	PickingTypeId               *Many2One   `xmlrpc:"picking_type_id,omitempty" json:"picking_type_id,omitempty"`
	Printed                     *Bool       `xmlrpc:"printed,omitempty" json:"printed,omitempty"`
	Priority                    *Selection  `xmlrpc:"priority,omitempty" json:"priority,omitempty"`
	ProductId                   *Many2One   `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductsAvailability        *String     `xmlrpc:"products_availability,omitempty" json:"products_availability,omitempty"`
	ProductsAvailabilityState   *Selection  `xmlrpc:"products_availability_state,omitempty" json:"products_availability_state,omitempty"`
	PurchaseId                  *Many2One   `xmlrpc:"purchase_id,omitempty" json:"purchase_id,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	ReturnCount                 *Int        `xmlrpc:"return_count,omitempty" json:"return_count,omitempty"`
	ReturnId                    *Many2One   `xmlrpc:"return_id,omitempty" json:"return_id,omitempty"`
	ReturnIds                   *Relation   `xmlrpc:"return_ids,omitempty" json:"return_ids,omitempty"`
	SaleId                      *Many2One   `xmlrpc:"sale_id,omitempty" json:"sale_id,omitempty"`
	ScheduledDate               *Time       `xmlrpc:"scheduled_date,omitempty" json:"scheduled_date,omitempty"`
	ShowAllocation              *Bool       `xmlrpc:"show_allocation,omitempty" json:"show_allocation,omitempty"`
	ShowCheckAvailability       *Bool       `xmlrpc:"show_check_availability,omitempty" json:"show_check_availability,omitempty"`
	ShowClearQtyButton          *Bool       `xmlrpc:"show_clear_qty_button,omitempty" json:"show_clear_qty_button,omitempty"`
	ShowLotsText                *Bool       `xmlrpc:"show_lots_text,omitempty" json:"show_lots_text,omitempty"`
	ShowOperations              *Bool       `xmlrpc:"show_operations,omitempty" json:"show_operations,omitempty"`
	ShowReserved                *Bool       `xmlrpc:"show_reserved,omitempty" json:"show_reserved,omitempty"`
	ShowSetQtyButton            *Bool       `xmlrpc:"show_set_qty_button,omitempty" json:"show_set_qty_button,omitempty"`
	Signature                   *String     `xmlrpc:"signature,omitempty" json:"signature,omitempty"`
	State                       *Selection  `xmlrpc:"state,omitempty" json:"state,omitempty"`
	UseCreateLots               *Bool       `xmlrpc:"use_create_lots,omitempty" json:"use_create_lots,omitempty"`
	UseExistingLots             *Bool       `xmlrpc:"use_existing_lots,omitempty" json:"use_existing_lots,omitempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
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
