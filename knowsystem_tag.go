package odoo

// KnowsystemTag represents knowsystem.tag model.
type KnowsystemTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	ApplyToAll  *Bool     `xmlrpc:"apply_to_all,omitempty"`
	ArticleIds  *Relation `xmlrpc:"article_ids,omitempty"`
	CanPublish  *Bool     `xmlrpc:"can_publish,omitempty"`
	ChildIds    *Relation `xmlrpc:"child_ids,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FilterIds   *Relation `xmlrpc:"filter_ids,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	IsPublished *Bool     `xmlrpc:"is_published,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WebsiteId   *Many2One `xmlrpc:"website_id,omitempty"`
	WebsiteUrl  *String   `xmlrpc:"website_url,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowsystemTags represents array of knowsystem.tag model.
type KnowsystemTags []KnowsystemTag

// KnowsystemTagModel is the odoo model name.
const KnowsystemTagModel = "knowsystem.tag"

// Many2One convert KnowsystemTag to *Many2One.
func (kt *KnowsystemTag) Many2One() *Many2One {
	return NewMany2One(kt.Id.Get(), "")
}

// CreateKnowsystemTag creates a new knowsystem.tag model and returns its id.
func (c *Client) CreateKnowsystemTag(kt *KnowsystemTag) (int64, error) {
	ids, err := c.CreateKnowsystemTags([]*KnowsystemTag{kt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemTag creates a new knowsystem.tag model and returns its id.
func (c *Client) CreateKnowsystemTags(kts []*KnowsystemTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range kts {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemTagModel, vv, nil)
}

// UpdateKnowsystemTag updates an existing knowsystem.tag record.
func (c *Client) UpdateKnowsystemTag(kt *KnowsystemTag) error {
	return c.UpdateKnowsystemTags([]int64{kt.Id.Get()}, kt)
}

// UpdateKnowsystemTags updates existing knowsystem.tag records.
// All records (represented by ids) will be updated by kt values.
func (c *Client) UpdateKnowsystemTags(ids []int64, kt *KnowsystemTag) error {
	return c.Update(KnowsystemTagModel, ids, kt, nil)
}

// DeleteKnowsystemTag deletes an existing knowsystem.tag record.
func (c *Client) DeleteKnowsystemTag(id int64) error {
	return c.DeleteKnowsystemTags([]int64{id})
}

// DeleteKnowsystemTags deletes existing knowsystem.tag records.
func (c *Client) DeleteKnowsystemTags(ids []int64) error {
	return c.Delete(KnowsystemTagModel, ids)
}

// GetKnowsystemTag gets knowsystem.tag existing record.
func (c *Client) GetKnowsystemTag(id int64) (*KnowsystemTag, error) {
	kts, err := c.GetKnowsystemTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kts)[0]), nil
}

// GetKnowsystemTags gets knowsystem.tag existing records.
func (c *Client) GetKnowsystemTags(ids []int64) (*KnowsystemTags, error) {
	kts := &KnowsystemTags{}
	if err := c.Read(KnowsystemTagModel, ids, nil, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTag finds knowsystem.tag record by querying it with criteria.
func (c *Client) FindKnowsystemTag(criteria *Criteria) (*KnowsystemTag, error) {
	kts := &KnowsystemTags{}
	if err := c.SearchRead(KnowsystemTagModel, criteria, NewOptions().Limit(1), kts); err != nil {
		return nil, err
	}
	return &((*kts)[0]), nil
}

// FindKnowsystemTags finds knowsystem.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTags(criteria *Criteria, options *Options) (*KnowsystemTags, error) {
	kts := &KnowsystemTags{}
	if err := c.SearchRead(KnowsystemTagModel, criteria, options, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowsystemTagModel, criteria, options)
}

// FindKnowsystemTagId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
