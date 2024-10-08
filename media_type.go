package odoo

// MediaType represents media.type model.
type MediaType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MediaTypes represents array of media.type model.
type MediaTypes []MediaType

// MediaTypeModel is the odoo model name.
const MediaTypeModel = "media.type"

// Many2One convert MediaType to *Many2One.
func (mt *MediaType) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMediaType creates a new media.type model and returns its id.
func (c *Client) CreateMediaType(mt *MediaType) (int64, error) {
	ids, err := c.CreateMediaTypes([]*MediaType{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMediaType creates a new media.type model and returns its id.
func (c *Client) CreateMediaTypes(mts []*MediaType) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MediaTypeModel, vv, nil)
}

// UpdateMediaType updates an existing media.type record.
func (c *Client) UpdateMediaType(mt *MediaType) error {
	return c.UpdateMediaTypes([]int64{mt.Id.Get()}, mt)
}

// UpdateMediaTypes updates existing media.type records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMediaTypes(ids []int64, mt *MediaType) error {
	return c.Update(MediaTypeModel, ids, mt, nil)
}

// DeleteMediaType deletes an existing media.type record.
func (c *Client) DeleteMediaType(id int64) error {
	return c.DeleteMediaTypes([]int64{id})
}

// DeleteMediaTypes deletes existing media.type records.
func (c *Client) DeleteMediaTypes(ids []int64) error {
	return c.Delete(MediaTypeModel, ids)
}

// GetMediaType gets media.type existing record.
func (c *Client) GetMediaType(id int64) (*MediaType, error) {
	mts, err := c.GetMediaTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// GetMediaTypes gets media.type existing records.
func (c *Client) GetMediaTypes(ids []int64) (*MediaTypes, error) {
	mts := &MediaTypes{}
	if err := c.Read(MediaTypeModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMediaType finds media.type record by querying it with criteria.
func (c *Client) FindMediaType(criteria *Criteria) (*MediaType, error) {
	mts := &MediaTypes{}
	if err := c.SearchRead(MediaTypeModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// FindMediaTypes finds media.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMediaTypes(criteria *Criteria, options *Options) (*MediaTypes, error) {
	mts := &MediaTypes{}
	if err := c.SearchRead(MediaTypeModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMediaTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMediaTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MediaTypeModel, criteria, options)
}

// FindMediaTypeId finds record id by querying it with criteria.
func (c *Client) FindMediaTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MediaTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
