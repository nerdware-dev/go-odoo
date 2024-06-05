package odoo

import (
	"fmt"
)

// KnowsystemTour represents knowsystem.tour model.
type KnowsystemTour struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	AccessUserIds             *Relation `xmlrpc:"access_user_ids,omptempty"`
	Active                    *Bool     `xmlrpc:"active,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrentArticleDescription *String   `xmlrpc:"current_article_description,omptempty"`
	CurrentArticleTitle       *String   `xmlrpc:"current_article_title,omptempty"`
	Description               *String   `xmlrpc:"description,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	Name                      *String   `xmlrpc:"name,omptempty"`
	Progress                  *Float    `xmlrpc:"progress,omptempty"`
	Sequence                  *Int      `xmlrpc:"sequence,omptempty"`
	ThisUserProgressId        *Many2One `xmlrpc:"this_user_progress_id,omptempty"`
	TourArticleIds            *Relation `xmlrpc:"tour_article_ids,omptempty"`
	UserGroupIds              *Relation `xmlrpc:"user_group_ids,omptempty"`
	UserIds                   *Relation `xmlrpc:"user_ids,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemTours represents array of knowsystem.tour model.
type KnowsystemTours []KnowsystemTour

// KnowsystemTourModel is the odoo model name.
const KnowsystemTourModel = "knowsystem.tour"

// Many2One convert KnowsystemTour to *Many2One.
func (kt *KnowsystemTour) Many2One() *Many2One {
	return NewMany2One(kt.Id.Get(), "")
}

// CreateKnowsystemTour creates a new knowsystem.tour model and returns its id.
func (c *Client) CreateKnowsystemTour(kt *KnowsystemTour) (int64, error) {
	ids, err := c.CreateKnowsystemTours([]*KnowsystemTour{kt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemTour creates a new knowsystem.tour model and returns its id.
func (c *Client) CreateKnowsystemTours(kts []*KnowsystemTour) ([]int64, error) {
	var vv []interface{}
	for _, v := range kts {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemTourModel, vv)
}

// UpdateKnowsystemTour updates an existing knowsystem.tour record.
func (c *Client) UpdateKnowsystemTour(kt *KnowsystemTour) error {
	return c.UpdateKnowsystemTours([]int64{kt.Id.Get()}, kt)
}

// UpdateKnowsystemTours updates existing knowsystem.tour records.
// All records (represented by ids) will be updated by kt values.
func (c *Client) UpdateKnowsystemTours(ids []int64, kt *KnowsystemTour) error {
	return c.Update(KnowsystemTourModel, ids, kt)
}

// DeleteKnowsystemTour deletes an existing knowsystem.tour record.
func (c *Client) DeleteKnowsystemTour(id int64) error {
	return c.DeleteKnowsystemTours([]int64{id})
}

// DeleteKnowsystemTours deletes existing knowsystem.tour records.
func (c *Client) DeleteKnowsystemTours(ids []int64) error {
	return c.Delete(KnowsystemTourModel, ids)
}

// GetKnowsystemTour gets knowsystem.tour existing record.
func (c *Client) GetKnowsystemTour(id int64) (*KnowsystemTour, error) {
	kts, err := c.GetKnowsystemTours([]int64{id})
	if err != nil {
		return nil, err
	}
	if kts != nil && len(*kts) > 0 {
		return &((*kts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.tour not found", id)
}

// GetKnowsystemTours gets knowsystem.tour existing records.
func (c *Client) GetKnowsystemTours(ids []int64) (*KnowsystemTours, error) {
	kts := &KnowsystemTours{}
	if err := c.Read(KnowsystemTourModel, ids, nil, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTour finds knowsystem.tour record by querying it with criteria.
func (c *Client) FindKnowsystemTour(criteria *Criteria) (*KnowsystemTour, error) {
	kts := &KnowsystemTours{}
	if err := c.SearchRead(KnowsystemTourModel, criteria, NewOptions().Limit(1), kts); err != nil {
		return nil, err
	}
	if kts != nil && len(*kts) > 0 {
		return &((*kts)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.tour was not found with criteria %v", criteria)
}

// FindKnowsystemTours finds knowsystem.tour records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTours(criteria *Criteria, options *Options) (*KnowsystemTours, error) {
	kts := &KnowsystemTours{}
	if err := c.SearchRead(KnowsystemTourModel, criteria, options, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTourIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTourIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemTourModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemTourId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTourId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTourModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.tour was not found with criteria %v and options %v", criteria, options)
}
