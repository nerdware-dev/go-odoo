package odoo

import (
	"fmt"
)

// FollowupSend represents followup.send model.
type FollowupSend struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	InvalidAddresses  *Int      `xmlrpc:"invalid_addresses,omptempty"`
	InvalidPartnerIds *Relation `xmlrpc:"invalid_partner_ids,omptempty"`
	LettersQty        *Int      `xmlrpc:"letters_qty,omptempty"`
	PartnerIds        *Relation `xmlrpc:"partner_ids,omptempty"`
	SnailmailCost     *Float    `xmlrpc:"snailmail_cost,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// FollowupSends represents array of followup.send model.
type FollowupSends []FollowupSend

// FollowupSendModel is the odoo model name.
const FollowupSendModel = "followup.send"

// Many2One convert FollowupSend to *Many2One.
func (fs *FollowupSend) Many2One() *Many2One {
	return NewMany2One(fs.Id.Get(), "")
}

// CreateFollowupSend creates a new followup.send model and returns its id.
func (c *Client) CreateFollowupSend(fs *FollowupSend) (int64, error) {
	ids, err := c.CreateFollowupSends([]*FollowupSend{fs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFollowupSend creates a new followup.send model and returns its id.
func (c *Client) CreateFollowupSends(fss []*FollowupSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range fss {
		vv = append(vv, v)
	}
	return c.Create(FollowupSendModel, vv)
}

// UpdateFollowupSend updates an existing followup.send record.
func (c *Client) UpdateFollowupSend(fs *FollowupSend) error {
	return c.UpdateFollowupSends([]int64{fs.Id.Get()}, fs)
}

// UpdateFollowupSends updates existing followup.send records.
// All records (represented by ids) will be updated by fs values.
func (c *Client) UpdateFollowupSends(ids []int64, fs *FollowupSend) error {
	return c.Update(FollowupSendModel, ids, fs)
}

// DeleteFollowupSend deletes an existing followup.send record.
func (c *Client) DeleteFollowupSend(id int64) error {
	return c.DeleteFollowupSends([]int64{id})
}

// DeleteFollowupSends deletes existing followup.send records.
func (c *Client) DeleteFollowupSends(ids []int64) error {
	return c.Delete(FollowupSendModel, ids)
}

// GetFollowupSend gets followup.send existing record.
func (c *Client) GetFollowupSend(id int64) (*FollowupSend, error) {
	fss, err := c.GetFollowupSends([]int64{id})
	if err != nil {
		return nil, err
	}
	if fss != nil && len(*fss) > 0 {
		return &((*fss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of followup.send not found", id)
}

// GetFollowupSends gets followup.send existing records.
func (c *Client) GetFollowupSends(ids []int64) (*FollowupSends, error) {
	fss := &FollowupSends{}
	if err := c.Read(FollowupSendModel, ids, nil, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFollowupSend finds followup.send record by querying it with criteria.
func (c *Client) FindFollowupSend(criteria *Criteria) (*FollowupSend, error) {
	fss := &FollowupSends{}
	if err := c.SearchRead(FollowupSendModel, criteria, NewOptions().Limit(1), fss); err != nil {
		return nil, err
	}
	if fss != nil && len(*fss) > 0 {
		return &((*fss)[0]), nil
	}
	return nil, fmt.Errorf("followup.send was not found with criteria %v", criteria)
}

// FindFollowupSends finds followup.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFollowupSends(criteria *Criteria, options *Options) (*FollowupSends, error) {
	fss := &FollowupSends{}
	if err := c.SearchRead(FollowupSendModel, criteria, options, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFollowupSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFollowupSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FollowupSendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFollowupSendId finds record id by querying it with criteria.
func (c *Client) FindFollowupSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FollowupSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("followup.send was not found with criteria %v and options %v", criteria, options)
}
