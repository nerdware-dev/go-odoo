package odoo

// StockLot represents stock.lot model.
type StockLot struct {
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
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DeliveryCount               *Int        `xmlrpc:"delivery_count,omitempty" json:"delivery_count,omitempty"`
	DeliveryIds                 *Relation   `xmlrpc:"delivery_ids,omitempty" json:"delivery_ids,omitempty"`
	DisplayComplete             *Bool       `xmlrpc:"display_complete,omitempty" json:"display_complete,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	LastDeliveryPartnerId       *Many2One   `xmlrpc:"last_delivery_partner_id,omitempty" json:"last_delivery_partner_id,omitempty"`
	LocationId                  *Many2One   `xmlrpc:"location_id,omitempty" json:"location_id,omitempty"`
	LotProperties               interface{} `xmlrpc:"lot_properties,omitempty" json:"lot_properties,omitempty"`
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
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	Note                        *String     `xmlrpc:"note,omitempty" json:"note,omitempty"`
	ProductId                   *Many2One   `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductQty                  *Float      `xmlrpc:"product_qty,omitempty" json:"product_qty,omitempty"`
	ProductUomId                *Many2One   `xmlrpc:"product_uom_id,omitempty" json:"product_uom_id,omitempty"`
	PurchaseOrderCount          *Int        `xmlrpc:"purchase_order_count,omitempty" json:"purchase_order_count,omitempty"`
	PurchaseOrderIds            *Relation   `xmlrpc:"purchase_order_ids,omitempty" json:"purchase_order_ids,omitempty"`
	QuantIds                    *Relation   `xmlrpc:"quant_ids,omitempty" json:"quant_ids,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	Ref                         *String     `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	SaleOrderCount              *Int        `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SaleOrderIds                *Relation   `xmlrpc:"sale_order_ids,omitempty" json:"sale_order_ids,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// StockLots represents array of stock.lot model.
type StockLots []StockLot

// StockLotModel is the odoo model name.
const StockLotModel = "stock.lot"

// Many2One convert StockLot to *Many2One.
func (sl *StockLot) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateStockLot creates a new stock.lot model and returns its id.
func (c *Client) CreateStockLot(sl *StockLot) (int64, error) {
	ids, err := c.CreateStockLots([]*StockLot{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLot creates a new stock.lot model and returns its id.
func (c *Client) CreateStockLots(sls []*StockLot) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(StockLotModel, vv, nil)
}

// UpdateStockLot updates an existing stock.lot record.
func (c *Client) UpdateStockLot(sl *StockLot) error {
	return c.UpdateStockLots([]int64{sl.Id.Get()}, sl)
}

// UpdateStockLots updates existing stock.lot records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockLots(ids []int64, sl *StockLot) error {
	return c.Update(StockLotModel, ids, sl, nil)
}

// DeleteStockLot deletes an existing stock.lot record.
func (c *Client) DeleteStockLot(id int64) error {
	return c.DeleteStockLots([]int64{id})
}

// DeleteStockLots deletes existing stock.lot records.
func (c *Client) DeleteStockLots(ids []int64) error {
	return c.Delete(StockLotModel, ids)
}

// GetStockLot gets stock.lot existing record.
func (c *Client) GetStockLot(id int64) (*StockLot, error) {
	sls, err := c.GetStockLots([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*sls) == 0 {
		return nil, nil
	}
	return &((*sls)[0]), nil
}

// GetStockLots gets stock.lot existing records.
func (c *Client) GetStockLots(ids []int64) (*StockLots, error) {
	sls := &StockLots{}
	if err := c.Read(StockLotModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLot finds stock.lot record by querying it with criteria.
func (c *Client) FindStockLot(criteria *Criteria) (*StockLot, error) {
	sls := &StockLots{}
	if err := c.SearchRead(StockLotModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	if len(*sls) == 0 {
		return nil, nil
	}
	return &((*sls)[0]), nil
}

// FindStockLots finds stock.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLots(criteria *Criteria, options *Options) (*StockLots, error) {
	sls := &StockLots{}
	if err := c.SearchRead(StockLotModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockLotModel, criteria, options)
}

// FindStockLotId finds record id by querying it with criteria.
func (c *Client) FindStockLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
