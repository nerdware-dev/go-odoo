package odoo

// ApprovalApprover represents approval.approver model.
type ApprovalApprover struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	RequestId   *Many2One  `xmlrpc:"request_id,omitempty"`
	Status      *Selection `xmlrpc:"status,omitempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(ApprovalApproverModel, vv, nil)
}

// UpdateApprovalApprover updates an existing approval.approver record.
func (c *Client) UpdateApprovalApprover(aa *ApprovalApprover) error {
	return c.UpdateApprovalApprovers([]int64{aa.Id.Get()}, aa)
}

// UpdateApprovalApprovers updates existing approval.approver records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateApprovalApprovers(ids []int64, aa *ApprovalApprover) error {
	return c.Update(ApprovalApproverModel, ids, aa, nil)
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
	return &((*aas)[0]), nil
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
	return &((*aas)[0]), nil
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
	return c.Search(ApprovalApproverModel, criteria, options)
}

// FindApprovalApproverId finds record id by querying it with criteria.
func (c *Client) FindApprovalApproverId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApprovalApproverModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
