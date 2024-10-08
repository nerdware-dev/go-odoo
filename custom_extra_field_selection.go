package odoo

// CustomExtraFieldSelection represents custom.extra.field.selection model.
type CustomExtraFieldSelection struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FieldId     *Int      `xmlrpc:"field_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Key         *String   `xmlrpc:"key,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CustomExtraFieldSelections represents array of custom.extra.field.selection model.
type CustomExtraFieldSelections []CustomExtraFieldSelection

// CustomExtraFieldSelectionModel is the odoo model name.
const CustomExtraFieldSelectionModel = "custom.extra.field.selection"

// Many2One convert CustomExtraFieldSelection to *Many2One.
func (cefs *CustomExtraFieldSelection) Many2One() *Many2One {
	return NewMany2One(cefs.Id.Get(), "")
}

// CreateCustomExtraFieldSelection creates a new custom.extra.field.selection model and returns its id.
func (c *Client) CreateCustomExtraFieldSelection(cefs *CustomExtraFieldSelection) (int64, error) {
	ids, err := c.CreateCustomExtraFieldSelections([]*CustomExtraFieldSelection{cefs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCustomExtraFieldSelection creates a new custom.extra.field.selection model and returns its id.
func (c *Client) CreateCustomExtraFieldSelections(cefss []*CustomExtraFieldSelection) ([]int64, error) {
	var vv []interface{}
	for _, v := range cefss {
		vv = append(vv, v)
	}
	return c.Create(CustomExtraFieldSelectionModel, vv, nil)
}

// UpdateCustomExtraFieldSelection updates an existing custom.extra.field.selection record.
func (c *Client) UpdateCustomExtraFieldSelection(cefs *CustomExtraFieldSelection) error {
	return c.UpdateCustomExtraFieldSelections([]int64{cefs.Id.Get()}, cefs)
}

// UpdateCustomExtraFieldSelections updates existing custom.extra.field.selection records.
// All records (represented by ids) will be updated by cefs values.
func (c *Client) UpdateCustomExtraFieldSelections(ids []int64, cefs *CustomExtraFieldSelection) error {
	return c.Update(CustomExtraFieldSelectionModel, ids, cefs, nil)
}

// DeleteCustomExtraFieldSelection deletes an existing custom.extra.field.selection record.
func (c *Client) DeleteCustomExtraFieldSelection(id int64) error {
	return c.DeleteCustomExtraFieldSelections([]int64{id})
}

// DeleteCustomExtraFieldSelections deletes existing custom.extra.field.selection records.
func (c *Client) DeleteCustomExtraFieldSelections(ids []int64) error {
	return c.Delete(CustomExtraFieldSelectionModel, ids)
}

// GetCustomExtraFieldSelection gets custom.extra.field.selection existing record.
func (c *Client) GetCustomExtraFieldSelection(id int64) (*CustomExtraFieldSelection, error) {
	cefss, err := c.GetCustomExtraFieldSelections([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cefss)[0]), nil
}

// GetCustomExtraFieldSelections gets custom.extra.field.selection existing records.
func (c *Client) GetCustomExtraFieldSelections(ids []int64) (*CustomExtraFieldSelections, error) {
	cefss := &CustomExtraFieldSelections{}
	if err := c.Read(CustomExtraFieldSelectionModel, ids, nil, cefss); err != nil {
		return nil, err
	}
	return cefss, nil
}

// FindCustomExtraFieldSelection finds custom.extra.field.selection record by querying it with criteria.
func (c *Client) FindCustomExtraFieldSelection(criteria *Criteria) (*CustomExtraFieldSelection, error) {
	cefss := &CustomExtraFieldSelections{}
	if err := c.SearchRead(CustomExtraFieldSelectionModel, criteria, NewOptions().Limit(1), cefss); err != nil {
		return nil, err
	}
	return &((*cefss)[0]), nil
}

// FindCustomExtraFieldSelections finds custom.extra.field.selection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomExtraFieldSelections(criteria *Criteria, options *Options) (*CustomExtraFieldSelections, error) {
	cefss := &CustomExtraFieldSelections{}
	if err := c.SearchRead(CustomExtraFieldSelectionModel, criteria, options, cefss); err != nil {
		return nil, err
	}
	return cefss, nil
}

// FindCustomExtraFieldSelectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomExtraFieldSelectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CustomExtraFieldSelectionModel, criteria, options)
}

// FindCustomExtraFieldSelectionId finds record id by querying it with criteria.
func (c *Client) FindCustomExtraFieldSelectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CustomExtraFieldSelectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
