package odoo

// MailAddressMixin represents mail.address.mixin model.
type MailAddressMixin struct {
	LastUpdate      *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName     *String `xmlrpc:"display_name,omitempty"`
	EmailNormalized *String `xmlrpc:"email_normalized,omitempty"`
	Id              *Int    `xmlrpc:"id,omitempty"`
}

// MailAddressMixins represents array of mail.address.mixin model.
type MailAddressMixins []MailAddressMixin

// MailAddressMixinModel is the odoo model name.
const MailAddressMixinModel = "mail.address.mixin"

// Many2One convert MailAddressMixin to *Many2One.
func (mam *MailAddressMixin) Many2One() *Many2One {
	return NewMany2One(mam.Id.Get(), "")
}

// CreateMailAddressMixin creates a new mail.address.mixin model and returns its id.
func (c *Client) CreateMailAddressMixin(mam *MailAddressMixin) (int64, error) {
	ids, err := c.CreateMailAddressMixins([]*MailAddressMixin{mam})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailAddressMixin creates a new mail.address.mixin model and returns its id.
func (c *Client) CreateMailAddressMixins(mams []*MailAddressMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mams {
		vv = append(vv, v)
	}
	return c.Create(MailAddressMixinModel, vv, nil)
}

// UpdateMailAddressMixin updates an existing mail.address.mixin record.
func (c *Client) UpdateMailAddressMixin(mam *MailAddressMixin) error {
	return c.UpdateMailAddressMixins([]int64{mam.Id.Get()}, mam)
}

// UpdateMailAddressMixins updates existing mail.address.mixin records.
// All records (represented by ids) will be updated by mam values.
func (c *Client) UpdateMailAddressMixins(ids []int64, mam *MailAddressMixin) error {
	return c.Update(MailAddressMixinModel, ids, mam, nil)
}

// DeleteMailAddressMixin deletes an existing mail.address.mixin record.
func (c *Client) DeleteMailAddressMixin(id int64) error {
	return c.DeleteMailAddressMixins([]int64{id})
}

// DeleteMailAddressMixins deletes existing mail.address.mixin records.
func (c *Client) DeleteMailAddressMixins(ids []int64) error {
	return c.Delete(MailAddressMixinModel, ids)
}

// GetMailAddressMixin gets mail.address.mixin existing record.
func (c *Client) GetMailAddressMixin(id int64) (*MailAddressMixin, error) {
	mams, err := c.GetMailAddressMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mams)[0]), nil
}

// GetMailAddressMixins gets mail.address.mixin existing records.
func (c *Client) GetMailAddressMixins(ids []int64) (*MailAddressMixins, error) {
	mams := &MailAddressMixins{}
	if err := c.Read(MailAddressMixinModel, ids, nil, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAddressMixin finds mail.address.mixin record by querying it with criteria.
func (c *Client) FindMailAddressMixin(criteria *Criteria) (*MailAddressMixin, error) {
	mams := &MailAddressMixins{}
	if err := c.SearchRead(MailAddressMixinModel, criteria, NewOptions().Limit(1), mams); err != nil {
		return nil, err
	}
	return &((*mams)[0]), nil
}

// FindMailAddressMixins finds mail.address.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAddressMixins(criteria *Criteria, options *Options) (*MailAddressMixins, error) {
	mams := &MailAddressMixins{}
	if err := c.SearchRead(MailAddressMixinModel, criteria, options, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAddressMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAddressMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailAddressMixinModel, criteria, options)
}

// FindMailAddressMixinId finds record id by querying it with criteria.
func (c *Client) FindMailAddressMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAddressMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
