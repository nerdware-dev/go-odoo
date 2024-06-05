package odoo

import (
	"fmt"
)

// KnowsystemTourUser represents knowsystem.tour.user model.
type KnowsystemTourUser struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrentArticleId *Many2One `xmlrpc:"current_article_id,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	LeftArticleIds   *Relation `xmlrpc:"left_article_ids,omptempty"`
	Progress         *Float    `xmlrpc:"progress,omptempty"`
	TourId           *Many2One `xmlrpc:"tour_id,omptempty"`
	UserId           *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemTourUsers represents array of knowsystem.tour.user model.
type KnowsystemTourUsers []KnowsystemTourUser

// KnowsystemTourUserModel is the odoo model name.
const KnowsystemTourUserModel = "knowsystem.tour.user"

// Many2One convert KnowsystemTourUser to *Many2One.
func (ktu *KnowsystemTourUser) Many2One() *Many2One {
	return NewMany2One(ktu.Id.Get(), "")
}

// CreateKnowsystemTourUser creates a new knowsystem.tour.user model and returns its id.
func (c *Client) CreateKnowsystemTourUser(ktu *KnowsystemTourUser) (int64, error) {
	ids, err := c.CreateKnowsystemTourUsers([]*KnowsystemTourUser{ktu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemTourUser creates a new knowsystem.tour.user model and returns its id.
func (c *Client) CreateKnowsystemTourUsers(ktus []*KnowsystemTourUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range ktus {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemTourUserModel, vv)
}

// UpdateKnowsystemTourUser updates an existing knowsystem.tour.user record.
func (c *Client) UpdateKnowsystemTourUser(ktu *KnowsystemTourUser) error {
	return c.UpdateKnowsystemTourUsers([]int64{ktu.Id.Get()}, ktu)
}

// UpdateKnowsystemTourUsers updates existing knowsystem.tour.user records.
// All records (represented by ids) will be updated by ktu values.
func (c *Client) UpdateKnowsystemTourUsers(ids []int64, ktu *KnowsystemTourUser) error {
	return c.Update(KnowsystemTourUserModel, ids, ktu)
}

// DeleteKnowsystemTourUser deletes an existing knowsystem.tour.user record.
func (c *Client) DeleteKnowsystemTourUser(id int64) error {
	return c.DeleteKnowsystemTourUsers([]int64{id})
}

// DeleteKnowsystemTourUsers deletes existing knowsystem.tour.user records.
func (c *Client) DeleteKnowsystemTourUsers(ids []int64) error {
	return c.Delete(KnowsystemTourUserModel, ids)
}

// GetKnowsystemTourUser gets knowsystem.tour.user existing record.
func (c *Client) GetKnowsystemTourUser(id int64) (*KnowsystemTourUser, error) {
	ktus, err := c.GetKnowsystemTourUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	if ktus != nil && len(*ktus) > 0 {
		return &((*ktus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.tour.user not found", id)
}

// GetKnowsystemTourUsers gets knowsystem.tour.user existing records.
func (c *Client) GetKnowsystemTourUsers(ids []int64) (*KnowsystemTourUsers, error) {
	ktus := &KnowsystemTourUsers{}
	if err := c.Read(KnowsystemTourUserModel, ids, nil, ktus); err != nil {
		return nil, err
	}
	return ktus, nil
}

// FindKnowsystemTourUser finds knowsystem.tour.user record by querying it with criteria.
func (c *Client) FindKnowsystemTourUser(criteria *Criteria) (*KnowsystemTourUser, error) {
	ktus := &KnowsystemTourUsers{}
	if err := c.SearchRead(KnowsystemTourUserModel, criteria, NewOptions().Limit(1), ktus); err != nil {
		return nil, err
	}
	if ktus != nil && len(*ktus) > 0 {
		return &((*ktus)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.tour.user was not found with criteria %v", criteria)
}

// FindKnowsystemTourUsers finds knowsystem.tour.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTourUsers(criteria *Criteria, options *Options) (*KnowsystemTourUsers, error) {
	ktus := &KnowsystemTourUsers{}
	if err := c.SearchRead(KnowsystemTourUserModel, criteria, options, ktus); err != nil {
		return nil, err
	}
	return ktus, nil
}

// FindKnowsystemTourUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTourUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemTourUserModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemTourUserId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTourUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTourUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.tour.user was not found with criteria %v and options %v", criteria, options)
}
