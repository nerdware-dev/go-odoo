package odoo

// WebsitePublishedMixin represents website.published.mixin model.
type WebsitePublishedMixin struct {
	LastUpdate       *Time   `xmlrpc:"__last_update,omitempty"`
	CanPublish       *Bool   `xmlrpc:"can_publish,omitempty"`
	DisplayName      *String `xmlrpc:"display_name,omitempty"`
	Id               *Int    `xmlrpc:"id,omitempty"`
	IsPublished      *Bool   `xmlrpc:"is_published,omitempty"`
	WebsitePublished *Bool   `xmlrpc:"website_published,omitempty"`
	WebsiteUrl       *String `xmlrpc:"website_url,omitempty"`
}

// WebsitePublishedMixins represents array of website.published.mixin model.
type WebsitePublishedMixins []WebsitePublishedMixin

// WebsitePublishedMixinModel is the odoo model name.
const WebsitePublishedMixinModel = "website.published.mixin"

// Many2One convert WebsitePublishedMixin to *Many2One.
func (wpm *WebsitePublishedMixin) Many2One() *Many2One {
	return NewMany2One(wpm.Id.Get(), "")
}

// CreateWebsitePublishedMixin creates a new website.published.mixin model and returns its id.
func (c *Client) CreateWebsitePublishedMixin(wpm *WebsitePublishedMixin) (int64, error) {
	ids, err := c.CreateWebsitePublishedMixins([]*WebsitePublishedMixin{wpm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsitePublishedMixin creates a new website.published.mixin model and returns its id.
func (c *Client) CreateWebsitePublishedMixins(wpms []*WebsitePublishedMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range wpms {
		vv = append(vv, v)
	}
	return c.Create(WebsitePublishedMixinModel, vv, nil)
}

// UpdateWebsitePublishedMixin updates an existing website.published.mixin record.
func (c *Client) UpdateWebsitePublishedMixin(wpm *WebsitePublishedMixin) error {
	return c.UpdateWebsitePublishedMixins([]int64{wpm.Id.Get()}, wpm)
}

// UpdateWebsitePublishedMixins updates existing website.published.mixin records.
// All records (represented by ids) will be updated by wpm values.
func (c *Client) UpdateWebsitePublishedMixins(ids []int64, wpm *WebsitePublishedMixin) error {
	return c.Update(WebsitePublishedMixinModel, ids, wpm, nil)
}

// DeleteWebsitePublishedMixin deletes an existing website.published.mixin record.
func (c *Client) DeleteWebsitePublishedMixin(id int64) error {
	return c.DeleteWebsitePublishedMixins([]int64{id})
}

// DeleteWebsitePublishedMixins deletes existing website.published.mixin records.
func (c *Client) DeleteWebsitePublishedMixins(ids []int64) error {
	return c.Delete(WebsitePublishedMixinModel, ids)
}

// GetWebsitePublishedMixin gets website.published.mixin existing record.
func (c *Client) GetWebsitePublishedMixin(id int64) (*WebsitePublishedMixin, error) {
	wpms, err := c.GetWebsitePublishedMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wpms)[0]), nil
}

// GetWebsitePublishedMixins gets website.published.mixin existing records.
func (c *Client) GetWebsitePublishedMixins(ids []int64) (*WebsitePublishedMixins, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.Read(WebsitePublishedMixinModel, ids, nil, wpms); err != nil {
		return nil, err
	}
	return wpms, nil
}

// FindWebsitePublishedMixin finds website.published.mixin record by querying it with criteria.
func (c *Client) FindWebsitePublishedMixin(criteria *Criteria) (*WebsitePublishedMixin, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.SearchRead(WebsitePublishedMixinModel, criteria, NewOptions().Limit(1), wpms); err != nil {
		return nil, err
	}
	return &((*wpms)[0]), nil
}

// FindWebsitePublishedMixins finds website.published.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMixins(criteria *Criteria, options *Options) (*WebsitePublishedMixins, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.SearchRead(WebsitePublishedMixinModel, criteria, options, wpms); err != nil {
		return nil, err
	}
	return wpms, nil
}

// FindWebsitePublishedMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsitePublishedMixinModel, criteria, options)
}

// FindWebsitePublishedMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsitePublishedMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePublishedMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
