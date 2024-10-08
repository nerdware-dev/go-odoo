package odoo

// CreateFromTemplate represents create.from.template model.
type CreateFromTemplate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	TemplateId  *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CreateFromTemplates represents array of create.from.template model.
type CreateFromTemplates []CreateFromTemplate

// CreateFromTemplateModel is the odoo model name.
const CreateFromTemplateModel = "create.from.template"

// Many2One convert CreateFromTemplate to *Many2One.
func (cft *CreateFromTemplate) Many2One() *Many2One {
	return NewMany2One(cft.Id.Get(), "")
}

// CreateCreateFromTemplate creates a new create.from.template model and returns its id.
func (c *Client) CreateCreateFromTemplate(cft *CreateFromTemplate) (int64, error) {
	ids, err := c.CreateCreateFromTemplates([]*CreateFromTemplate{cft})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCreateFromTemplate creates a new create.from.template model and returns its id.
func (c *Client) CreateCreateFromTemplates(cfts []*CreateFromTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range cfts {
		vv = append(vv, v)
	}
	return c.Create(CreateFromTemplateModel, vv, nil)
}

// UpdateCreateFromTemplate updates an existing create.from.template record.
func (c *Client) UpdateCreateFromTemplate(cft *CreateFromTemplate) error {
	return c.UpdateCreateFromTemplates([]int64{cft.Id.Get()}, cft)
}

// UpdateCreateFromTemplates updates existing create.from.template records.
// All records (represented by ids) will be updated by cft values.
func (c *Client) UpdateCreateFromTemplates(ids []int64, cft *CreateFromTemplate) error {
	return c.Update(CreateFromTemplateModel, ids, cft, nil)
}

// DeleteCreateFromTemplate deletes an existing create.from.template record.
func (c *Client) DeleteCreateFromTemplate(id int64) error {
	return c.DeleteCreateFromTemplates([]int64{id})
}

// DeleteCreateFromTemplates deletes existing create.from.template records.
func (c *Client) DeleteCreateFromTemplates(ids []int64) error {
	return c.Delete(CreateFromTemplateModel, ids)
}

// GetCreateFromTemplate gets create.from.template existing record.
func (c *Client) GetCreateFromTemplate(id int64) (*CreateFromTemplate, error) {
	cfts, err := c.GetCreateFromTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cfts)[0]), nil
}

// GetCreateFromTemplates gets create.from.template existing records.
func (c *Client) GetCreateFromTemplates(ids []int64) (*CreateFromTemplates, error) {
	cfts := &CreateFromTemplates{}
	if err := c.Read(CreateFromTemplateModel, ids, nil, cfts); err != nil {
		return nil, err
	}
	return cfts, nil
}

// FindCreateFromTemplate finds create.from.template record by querying it with criteria.
func (c *Client) FindCreateFromTemplate(criteria *Criteria) (*CreateFromTemplate, error) {
	cfts := &CreateFromTemplates{}
	if err := c.SearchRead(CreateFromTemplateModel, criteria, NewOptions().Limit(1), cfts); err != nil {
		return nil, err
	}
	return &((*cfts)[0]), nil
}

// FindCreateFromTemplates finds create.from.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCreateFromTemplates(criteria *Criteria, options *Options) (*CreateFromTemplates, error) {
	cfts := &CreateFromTemplates{}
	if err := c.SearchRead(CreateFromTemplateModel, criteria, options, cfts); err != nil {
		return nil, err
	}
	return cfts, nil
}

// FindCreateFromTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCreateFromTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CreateFromTemplateModel, criteria, options)
}

// FindCreateFromTemplateId finds record id by querying it with criteria.
func (c *Client) FindCreateFromTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CreateFromTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
