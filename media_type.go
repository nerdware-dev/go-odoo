package odoo

import (
	"fmt"
)

// MediaType represents media.type model.
type MediaType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(MediaTypeModel, vv)
}

// UpdateMediaType updates an existing media.type record.
func (c *Client) UpdateMediaType(mt *MediaType) error {
	return c.UpdateMediaTypes([]int64{mt.Id.Get()}, mt)
}

// UpdateMediaTypes updates existing media.type records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMediaTypes(ids []int64, mt *MediaType) error {
	return c.Update(MediaTypeModel, ids, mt)
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
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of media.type not found", id)
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
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("media.type was not found with criteria %v", criteria)
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
	ids, err := c.Search(MediaTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMediaTypeId finds record id by querying it with criteria.
func (c *Client) FindMediaTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MediaTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("media.type was not found with criteria %v and options %v", criteria, options)
}
