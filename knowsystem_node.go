package odoo

import (
	"fmt"
)

// KnowsystemNode represents knowsystem.node model.
type KnowsystemNode struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	Active      *Bool   `xmlrpc:"active,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
	Sequence    *Int    `xmlrpc:"sequence,omptempty"`
}

// KnowsystemNodes represents array of knowsystem.node model.
type KnowsystemNodes []KnowsystemNode

// KnowsystemNodeModel is the odoo model name.
const KnowsystemNodeModel = "knowsystem.node"

// Many2One convert KnowsystemNode to *Many2One.
func (kn *KnowsystemNode) Many2One() *Many2One {
	return NewMany2One(kn.Id.Get(), "")
}

// CreateKnowsystemNode creates a new knowsystem.node model and returns its id.
func (c *Client) CreateKnowsystemNode(kn *KnowsystemNode) (int64, error) {
	ids, err := c.CreateKnowsystemNodes([]*KnowsystemNode{kn})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemNode creates a new knowsystem.node model and returns its id.
func (c *Client) CreateKnowsystemNodes(kns []*KnowsystemNode) ([]int64, error) {
	var vv []interface{}
	for _, v := range kns {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemNodeModel, vv)
}

// UpdateKnowsystemNode updates an existing knowsystem.node record.
func (c *Client) UpdateKnowsystemNode(kn *KnowsystemNode) error {
	return c.UpdateKnowsystemNodes([]int64{kn.Id.Get()}, kn)
}

// UpdateKnowsystemNodes updates existing knowsystem.node records.
// All records (represented by ids) will be updated by kn values.
func (c *Client) UpdateKnowsystemNodes(ids []int64, kn *KnowsystemNode) error {
	return c.Update(KnowsystemNodeModel, ids, kn)
}

// DeleteKnowsystemNode deletes an existing knowsystem.node record.
func (c *Client) DeleteKnowsystemNode(id int64) error {
	return c.DeleteKnowsystemNodes([]int64{id})
}

// DeleteKnowsystemNodes deletes existing knowsystem.node records.
func (c *Client) DeleteKnowsystemNodes(ids []int64) error {
	return c.Delete(KnowsystemNodeModel, ids)
}

// GetKnowsystemNode gets knowsystem.node existing record.
func (c *Client) GetKnowsystemNode(id int64) (*KnowsystemNode, error) {
	kns, err := c.GetKnowsystemNodes([]int64{id})
	if err != nil {
		return nil, err
	}
	if kns != nil && len(*kns) > 0 {
		return &((*kns)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.node not found", id)
}

// GetKnowsystemNodes gets knowsystem.node existing records.
func (c *Client) GetKnowsystemNodes(ids []int64) (*KnowsystemNodes, error) {
	kns := &KnowsystemNodes{}
	if err := c.Read(KnowsystemNodeModel, ids, nil, kns); err != nil {
		return nil, err
	}
	return kns, nil
}

// FindKnowsystemNode finds knowsystem.node record by querying it with criteria.
func (c *Client) FindKnowsystemNode(criteria *Criteria) (*KnowsystemNode, error) {
	kns := &KnowsystemNodes{}
	if err := c.SearchRead(KnowsystemNodeModel, criteria, NewOptions().Limit(1), kns); err != nil {
		return nil, err
	}
	if kns != nil && len(*kns) > 0 {
		return &((*kns)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.node was not found with criteria %v", criteria)
}

// FindKnowsystemNodes finds knowsystem.node records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemNodes(criteria *Criteria, options *Options) (*KnowsystemNodes, error) {
	kns := &KnowsystemNodes{}
	if err := c.SearchRead(KnowsystemNodeModel, criteria, options, kns); err != nil {
		return nil, err
	}
	return kns, nil
}

// FindKnowsystemNodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemNodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemNodeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemNodeId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemNodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemNodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.node was not found with criteria %v and options %v", criteria, options)
}
