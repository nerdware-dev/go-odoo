package odoo

import (
	"fmt"
)

// ApprovalApprover represents approval.approver model.
type ApprovalApprover struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	RequestId   *Many2One  `xmlrpc:"request_id,omptempty"`
	Status      *Selection `xmlrpc:"status,omptempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ApprovalApprovers represents array of approval.approver model.
type ApprovalApprovers []ApprovalApprover

// ApprovalApproverModel is the odoo model name.
const ApprovalApproverModel = "approval.approver"

// Many2One convert ApprovalApprover to *Many2One.
func (aa *ApprovalApprover) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateApprovalApprover creates a new approval.approver model and returns its id.
func (c *Client) CreateApprovalApprover(aa *ApprovalApprover) (int64, error) {
	ids, err := c.CreateApprovalApprovers([]*ApprovalApprover{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApprovalApprover creates a new approval.approver model and returns its id.
func (c *Client) CreateApprovalApprovers(aas []*ApprovalApprover) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(ApprovalApproverModel, vv)
}

// UpdateApprovalApprover updates an existing approval.approver record.
func (c *Client) UpdateApprovalApprover(aa *ApprovalApprover) error {
	return c.UpdateApprovalApprovers([]int64{aa.Id.Get()}, aa)
}

// UpdateApprovalApprovers updates existing approval.approver records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateApprovalApprovers(ids []int64, aa *ApprovalApprover) error {
	return c.Update(ApprovalApproverModel, ids, aa)
}

// DeleteApprovalApprover deletes an existing approval.approver record.
func (c *Client) DeleteApprovalApprover(id int64) error {
	return c.DeleteApprovalApprovers([]int64{id})
}

// DeleteApprovalApprovers deletes existing approval.approver records.
func (c *Client) DeleteApprovalApprovers(ids []int64) error {
	return c.Delete(ApprovalApproverModel, ids)
}

// GetApprovalApprover gets approval.approver existing record.
func (c *Client) GetApprovalApprover(id int64) (*ApprovalApprover, error) {
	aas, err := c.GetApprovalApprovers([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of approval.approver not found", id)
}

// GetApprovalApprovers gets approval.approver existing records.
func (c *Client) GetApprovalApprovers(ids []int64) (*ApprovalApprovers, error) {
	aas := &ApprovalApprovers{}
	if err := c.Read(ApprovalApproverModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindApprovalApprover finds approval.approver record by querying it with criteria.
func (c *Client) FindApprovalApprover(criteria *Criteria) (*ApprovalApprover, error) {
	aas := &ApprovalApprovers{}
	if err := c.SearchRead(ApprovalApproverModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("approval.approver was not found with criteria %v", criteria)
}

// FindApprovalApprovers finds approval.approver records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalApprovers(criteria *Criteria, options *Options) (*ApprovalApprovers, error) {
	aas := &ApprovalApprovers{}
	if err := c.SearchRead(ApprovalApproverModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindApprovalApproverIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalApproverIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ApprovalApproverModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindApprovalApproverId finds record id by querying it with criteria.
func (c *Client) FindApprovalApproverId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApprovalApproverModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("approval.approver was not found with criteria %v and options %v", criteria, options)
}
