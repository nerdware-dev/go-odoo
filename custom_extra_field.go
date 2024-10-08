package odoo

// CustomExtraField represents custom.extra.field model.
type CustomExtraField struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	Active          *Bool      `xmlrpc:"active,omitempty"`
	BinaryFieldId   *Many2One  `xmlrpc:"binary_field_id,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	FieldId         *Many2One  `xmlrpc:"field_id,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	InputPlacement  *Selection `xmlrpc:"input_placement,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	Placement       *Selection `xmlrpc:"placement,omitempty"`
	PortalPlacement *Selection `xmlrpc:"portal_placement,omitempty"`
	Required        *Bool      `xmlrpc:"required,omitempty"`
	ResModel        *Selection `xmlrpc:"res_model,omitempty"`
	SelOptionsIds   *Relation  `xmlrpc:"sel_options_ids,omitempty"`
	Sequence        *Int       `xmlrpc:"sequence,omitempty"`
	Ttype           *Selection `xmlrpc:"ttype,omitempty"`
}

// CustomExtraFields represents array of custom.extra.field model.
type CustomExtraFields []CustomExtraField

// CustomExtraFieldModel is the odoo model name.
const CustomExtraFieldModel = "custom.extra.field"

// Many2One convert CustomExtraField to *Many2One.
func (cef *CustomExtraField) Many2One() *Many2One {
	return NewMany2One(cef.Id.Get(), "")
}

// CreateCustomExtraField creates a new custom.extra.field model and returns its id.
func (c *Client) CreateCustomExtraField(cef *CustomExtraField) (int64, error) {
	ids, err := c.CreateCustomExtraFields([]*CustomExtraField{cef})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCustomExtraField creates a new custom.extra.field model and returns its id.
func (c *Client) CreateCustomExtraFields(cefs []*CustomExtraField) ([]int64, error) {
	var vv []interface{}
	for _, v := range cefs {
		vv = append(vv, v)
	}
	return c.Create(CustomExtraFieldModel, vv, nil)
}

// UpdateCustomExtraField updates an existing custom.extra.field record.
func (c *Client) UpdateCustomExtraField(cef *CustomExtraField) error {
	return c.UpdateCustomExtraFields([]int64{cef.Id.Get()}, cef)
}

// UpdateCustomExtraFields updates existing custom.extra.field records.
// All records (represented by ids) will be updated by cef values.
func (c *Client) UpdateCustomExtraFields(ids []int64, cef *CustomExtraField) error {
	return c.Update(CustomExtraFieldModel, ids, cef, nil)
}

// DeleteCustomExtraField deletes an existing custom.extra.field record.
func (c *Client) DeleteCustomExtraField(id int64) error {
	return c.DeleteCustomExtraFields([]int64{id})
}

// DeleteCustomExtraFields deletes existing custom.extra.field records.
func (c *Client) DeleteCustomExtraFields(ids []int64) error {
	return c.Delete(CustomExtraFieldModel, ids)
}

// GetCustomExtraField gets custom.extra.field existing record.
func (c *Client) GetCustomExtraField(id int64) (*CustomExtraField, error) {
	cefs, err := c.GetCustomExtraFields([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cefs)[0]), nil
}

// GetCustomExtraFields gets custom.extra.field existing records.
func (c *Client) GetCustomExtraFields(ids []int64) (*CustomExtraFields, error) {
	cefs := &CustomExtraFields{}
	if err := c.Read(CustomExtraFieldModel, ids, nil, cefs); err != nil {
		return nil, err
	}
	return cefs, nil
}

// FindCustomExtraField finds custom.extra.field record by querying it with criteria.
func (c *Client) FindCustomExtraField(criteria *Criteria) (*CustomExtraField, error) {
	cefs := &CustomExtraFields{}
	if err := c.SearchRead(CustomExtraFieldModel, criteria, NewOptions().Limit(1), cefs); err != nil {
		return nil, err
	}
	return &((*cefs)[0]), nil
}

// FindCustomExtraFields finds custom.extra.field records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomExtraFields(criteria *Criteria, options *Options) (*CustomExtraFields, error) {
	cefs := &CustomExtraFields{}
	if err := c.SearchRead(CustomExtraFieldModel, criteria, options, cefs); err != nil {
		return nil, err
	}
	return cefs, nil
}

// FindCustomExtraFieldIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomExtraFieldIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CustomExtraFieldModel, criteria, options)
}

// FindCustomExtraFieldId finds record id by querying it with criteria.
func (c *Client) FindCustomExtraFieldId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CustomExtraFieldModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
