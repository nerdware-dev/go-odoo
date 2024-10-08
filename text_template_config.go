package odoo

// TextTemplateConfig represents text.template.config model.
type TextTemplateConfig struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultText *String   `xmlrpc:"default_text,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Model       *Many2One `xmlrpc:"model,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// TextTemplateConfigs represents array of text.template.config model.
type TextTemplateConfigs []TextTemplateConfig

// TextTemplateConfigModel is the odoo model name.
const TextTemplateConfigModel = "text.template.config"

// Many2One convert TextTemplateConfig to *Many2One.
func (ttc *TextTemplateConfig) Many2One() *Many2One {
	return NewMany2One(ttc.Id.Get(), "")
}

// CreateTextTemplateConfig creates a new text.template.config model and returns its id.
func (c *Client) CreateTextTemplateConfig(ttc *TextTemplateConfig) (int64, error) {
	ids, err := c.CreateTextTemplateConfigs([]*TextTemplateConfig{ttc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTextTemplateConfig creates a new text.template.config model and returns its id.
func (c *Client) CreateTextTemplateConfigs(ttcs []*TextTemplateConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range ttcs {
		vv = append(vv, v)
	}
	return c.Create(TextTemplateConfigModel, vv, nil)
}

// UpdateTextTemplateConfig updates an existing text.template.config record.
func (c *Client) UpdateTextTemplateConfig(ttc *TextTemplateConfig) error {
	return c.UpdateTextTemplateConfigs([]int64{ttc.Id.Get()}, ttc)
}

// UpdateTextTemplateConfigs updates existing text.template.config records.
// All records (represented by ids) will be updated by ttc values.
func (c *Client) UpdateTextTemplateConfigs(ids []int64, ttc *TextTemplateConfig) error {
	return c.Update(TextTemplateConfigModel, ids, ttc, nil)
}

// DeleteTextTemplateConfig deletes an existing text.template.config record.
func (c *Client) DeleteTextTemplateConfig(id int64) error {
	return c.DeleteTextTemplateConfigs([]int64{id})
}

// DeleteTextTemplateConfigs deletes existing text.template.config records.
func (c *Client) DeleteTextTemplateConfigs(ids []int64) error {
	return c.Delete(TextTemplateConfigModel, ids)
}

// GetTextTemplateConfig gets text.template.config existing record.
func (c *Client) GetTextTemplateConfig(id int64) (*TextTemplateConfig, error) {
	ttcs, err := c.GetTextTemplateConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ttcs)[0]), nil
}

// GetTextTemplateConfigs gets text.template.config existing records.
func (c *Client) GetTextTemplateConfigs(ids []int64) (*TextTemplateConfigs, error) {
	ttcs := &TextTemplateConfigs{}
	if err := c.Read(TextTemplateConfigModel, ids, nil, ttcs); err != nil {
		return nil, err
	}
	return ttcs, nil
}

// FindTextTemplateConfig finds text.template.config record by querying it with criteria.
func (c *Client) FindTextTemplateConfig(criteria *Criteria) (*TextTemplateConfig, error) {
	ttcs := &TextTemplateConfigs{}
	if err := c.SearchRead(TextTemplateConfigModel, criteria, NewOptions().Limit(1), ttcs); err != nil {
		return nil, err
	}
	return &((*ttcs)[0]), nil
}

// FindTextTemplateConfigs finds text.template.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTextTemplateConfigs(criteria *Criteria, options *Options) (*TextTemplateConfigs, error) {
	ttcs := &TextTemplateConfigs{}
	if err := c.SearchRead(TextTemplateConfigModel, criteria, options, ttcs); err != nil {
		return nil, err
	}
	return ttcs, nil
}

// FindTextTemplateConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTextTemplateConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TextTemplateConfigModel, criteria, options)
}

// FindTextTemplateConfigId finds record id by querying it with criteria.
func (c *Client) FindTextTemplateConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TextTemplateConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
