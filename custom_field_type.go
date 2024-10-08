package odoo

// CustomFieldType represents custom.field.type model.
type CustomFieldType struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	Active      *Bool   `xmlrpc:"active,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	InputOption *Bool   `xmlrpc:"input_option,omitempty"`
	Name        *String `xmlrpc:"name,omitempty"`
}

// CustomFieldTypes represents array of custom.field.type model.
type CustomFieldTypes []CustomFieldType

// CustomFieldTypeModel is the odoo model name.
const CustomFieldTypeModel = "custom.field.type"

// Many2One convert CustomFieldType to *Many2One.
func (cft *CustomFieldType) Many2One() *Many2One {
	return NewMany2One(cft.Id.Get(), "")
}

// CreateCustomFieldType creates a new custom.field.type model and returns its id.
func (c *Client) CreateCustomFieldType(cft *CustomFieldType) (int64, error) {
	ids, err := c.CreateCustomFieldTypes([]*CustomFieldType{cft})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCustomFieldType creates a new custom.field.type model and returns its id.
func (c *Client) CreateCustomFieldTypes(cfts []*CustomFieldType) ([]int64, error) {
	var vv []interface{}
	for _, v := range cfts {
		vv = append(vv, v)
	}
	return c.Create(CustomFieldTypeModel, vv, nil)
}

// UpdateCustomFieldType updates an existing custom.field.type record.
func (c *Client) UpdateCustomFieldType(cft *CustomFieldType) error {
	return c.UpdateCustomFieldTypes([]int64{cft.Id.Get()}, cft)
}

// UpdateCustomFieldTypes updates existing custom.field.type records.
// All records (represented by ids) will be updated by cft values.
func (c *Client) UpdateCustomFieldTypes(ids []int64, cft *CustomFieldType) error {
	return c.Update(CustomFieldTypeModel, ids, cft, nil)
}

// DeleteCustomFieldType deletes an existing custom.field.type record.
func (c *Client) DeleteCustomFieldType(id int64) error {
	return c.DeleteCustomFieldTypes([]int64{id})
}

// DeleteCustomFieldTypes deletes existing custom.field.type records.
func (c *Client) DeleteCustomFieldTypes(ids []int64) error {
	return c.Delete(CustomFieldTypeModel, ids)
}

// GetCustomFieldType gets custom.field.type existing record.
func (c *Client) GetCustomFieldType(id int64) (*CustomFieldType, error) {
	cfts, err := c.GetCustomFieldTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cfts)[0]), nil
}

// GetCustomFieldTypes gets custom.field.type existing records.
func (c *Client) GetCustomFieldTypes(ids []int64) (*CustomFieldTypes, error) {
	cfts := &CustomFieldTypes{}
	if err := c.Read(CustomFieldTypeModel, ids, nil, cfts); err != nil {
		return nil, err
	}
	return cfts, nil
}

// FindCustomFieldType finds custom.field.type record by querying it with criteria.
func (c *Client) FindCustomFieldType(criteria *Criteria) (*CustomFieldType, error) {
	cfts := &CustomFieldTypes{}
	if err := c.SearchRead(CustomFieldTypeModel, criteria, NewOptions().Limit(1), cfts); err != nil {
		return nil, err
	}
	return &((*cfts)[0]), nil
}

// FindCustomFieldTypes finds custom.field.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomFieldTypes(criteria *Criteria, options *Options) (*CustomFieldTypes, error) {
	cfts := &CustomFieldTypes{}
	if err := c.SearchRead(CustomFieldTypeModel, criteria, options, cfts); err != nil {
		return nil, err
	}
	return cfts, nil
}

// FindCustomFieldTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCustomFieldTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CustomFieldTypeModel, criteria, options)
}

// FindCustomFieldTypeId finds record id by querying it with criteria.
func (c *Client) FindCustomFieldTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CustomFieldTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
