package odoo

// LunchOrderTemp represents lunch.order.temp model.
type LunchOrderTemp struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	AvailableToppings1 *Bool      `xmlrpc:"available_toppings_1,omitempty"`
	AvailableToppings2 *Bool      `xmlrpc:"available_toppings_2,omitempty"`
	AvailableToppings3 *Bool      `xmlrpc:"available_toppings_3,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Edit               *Bool      `xmlrpc:"edit,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	Image128           *String    `xmlrpc:"image_128,omitempty"`
	Image1920          *String    `xmlrpc:"image_1920,omitempty"`
	Note               *String    `xmlrpc:"note,omitempty"`
	PriceTotal         *Float     `xmlrpc:"price_total,omitempty"`
	ProductCategory    *Many2One  `xmlrpc:"product_category,omitempty"`
	ProductDescription *String    `xmlrpc:"product_description,omitempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductName        *String    `xmlrpc:"product_name,omitempty"`
	Quantity           *Float     `xmlrpc:"quantity,omitempty"`
	ToppingIds1        *Relation  `xmlrpc:"topping_ids_1,omitempty"`
	ToppingIds2        *Relation  `xmlrpc:"topping_ids_2,omitempty"`
	ToppingIds3        *Relation  `xmlrpc:"topping_ids_3,omitempty"`
	ToppingLabel1      *String    `xmlrpc:"topping_label_1,omitempty"`
	ToppingLabel2      *String    `xmlrpc:"topping_label_2,omitempty"`
	ToppingLabel3      *String    `xmlrpc:"topping_label_3,omitempty"`
	ToppingQuantity1   *Selection `xmlrpc:"topping_quantity_1,omitempty"`
	ToppingQuantity2   *Selection `xmlrpc:"topping_quantity_2,omitempty"`
	ToppingQuantity3   *Selection `xmlrpc:"topping_quantity_3,omitempty"`
	UserId             *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LunchOrderTemps represents array of lunch.order.temp model.
type LunchOrderTemps []LunchOrderTemp

// LunchOrderTempModel is the odoo model name.
const LunchOrderTempModel = "lunch.order.temp"

// Many2One convert LunchOrderTemp to *Many2One.
func (lot *LunchOrderTemp) Many2One() *Many2One {
	return NewMany2One(lot.Id.Get(), "")
}

// CreateLunchOrderTemp creates a new lunch.order.temp model and returns its id.
func (c *Client) CreateLunchOrderTemp(lot *LunchOrderTemp) (int64, error) {
	ids, err := c.CreateLunchOrderTemps([]*LunchOrderTemp{lot})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchOrderTemp creates a new lunch.order.temp model and returns its id.
func (c *Client) CreateLunchOrderTemps(lots []*LunchOrderTemp) ([]int64, error) {
	var vv []interface{}
	for _, v := range lots {
		vv = append(vv, v)
	}
	return c.Create(LunchOrderTempModel, vv, nil)
}

// UpdateLunchOrderTemp updates an existing lunch.order.temp record.
func (c *Client) UpdateLunchOrderTemp(lot *LunchOrderTemp) error {
	return c.UpdateLunchOrderTemps([]int64{lot.Id.Get()}, lot)
}

// UpdateLunchOrderTemps updates existing lunch.order.temp records.
// All records (represented by ids) will be updated by lot values.
func (c *Client) UpdateLunchOrderTemps(ids []int64, lot *LunchOrderTemp) error {
	return c.Update(LunchOrderTempModel, ids, lot, nil)
}

// DeleteLunchOrderTemp deletes an existing lunch.order.temp record.
func (c *Client) DeleteLunchOrderTemp(id int64) error {
	return c.DeleteLunchOrderTemps([]int64{id})
}

// DeleteLunchOrderTemps deletes existing lunch.order.temp records.
func (c *Client) DeleteLunchOrderTemps(ids []int64) error {
	return c.Delete(LunchOrderTempModel, ids)
}

// GetLunchOrderTemp gets lunch.order.temp existing record.
func (c *Client) GetLunchOrderTemp(id int64) (*LunchOrderTemp, error) {
	lots, err := c.GetLunchOrderTemps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lots)[0]), nil
}

// GetLunchOrderTemps gets lunch.order.temp existing records.
func (c *Client) GetLunchOrderTemps(ids []int64) (*LunchOrderTemps, error) {
	lots := &LunchOrderTemps{}
	if err := c.Read(LunchOrderTempModel, ids, nil, lots); err != nil {
		return nil, err
	}
	return lots, nil
}

// FindLunchOrderTemp finds lunch.order.temp record by querying it with criteria.
func (c *Client) FindLunchOrderTemp(criteria *Criteria) (*LunchOrderTemp, error) {
	lots := &LunchOrderTemps{}
	if err := c.SearchRead(LunchOrderTempModel, criteria, NewOptions().Limit(1), lots); err != nil {
		return nil, err
	}
	return &((*lots)[0]), nil
}

// FindLunchOrderTemps finds lunch.order.temp records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrderTemps(criteria *Criteria, options *Options) (*LunchOrderTemps, error) {
	lots := &LunchOrderTemps{}
	if err := c.SearchRead(LunchOrderTempModel, criteria, options, lots); err != nil {
		return nil, err
	}
	return lots, nil
}

// FindLunchOrderTempIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchOrderTempIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchOrderTempModel, criteria, options)
}

// FindLunchOrderTempId finds record id by querying it with criteria.
func (c *Client) FindLunchOrderTempId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchOrderTempModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
