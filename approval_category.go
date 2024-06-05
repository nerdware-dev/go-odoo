package odoo

import (
	"fmt"
)

// ApprovalCategory represents approval.category model.
type ApprovalCategory struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	ApprovalMinimum        *Int       `xmlrpc:"approval_minimum,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description            *String    `xmlrpc:"description,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	HasAmount              *Selection `xmlrpc:"has_amount,omptempty"`
	HasDate                *Selection `xmlrpc:"has_date,omptempty"`
	HasItem                *Selection `xmlrpc:"has_item,omptempty"`
	HasLocation            *Selection `xmlrpc:"has_location,omptempty"`
	HasPartner             *Selection `xmlrpc:"has_partner,omptempty"`
	HasPaymentMethod       *Selection `xmlrpc:"has_payment_method,omptempty"`
	HasPeriod              *Selection `xmlrpc:"has_period,omptempty"`
	HasQuantity            *Selection `xmlrpc:"has_quantity,omptempty"`
	HasReference           *Selection `xmlrpc:"has_reference,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	Image                  *String    `xmlrpc:"image,omptempty"`
	IsManagerApprover      *Bool      `xmlrpc:"is_manager_approver,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	RequestToValidateCount *Int       `xmlrpc:"request_to_validate_count,omptempty"`
	RequirerDocument       *Selection `xmlrpc:"requirer_document,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	UserIds                *Relation  `xmlrpc:"user_ids,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ApprovalCategorys represents array of approval.category model.
type ApprovalCategorys []ApprovalCategory

// ApprovalCategoryModel is the odoo model name.
const ApprovalCategoryModel = "approval.category"

// Many2One convert ApprovalCategory to *Many2One.
func (ac *ApprovalCategory) Many2One() *Many2One {
	return NewMany2One(ac.Id.Get(), "")
}

// CreateApprovalCategory creates a new approval.category model and returns its id.
func (c *Client) CreateApprovalCategory(ac *ApprovalCategory) (int64, error) {
	ids, err := c.CreateApprovalCategorys([]*ApprovalCategory{ac})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApprovalCategory creates a new approval.category model and returns its id.
func (c *Client) CreateApprovalCategorys(acs []*ApprovalCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range acs {
		vv = append(vv, v)
	}
	return c.Create(ApprovalCategoryModel, vv)
}

// UpdateApprovalCategory updates an existing approval.category record.
func (c *Client) UpdateApprovalCategory(ac *ApprovalCategory) error {
	return c.UpdateApprovalCategorys([]int64{ac.Id.Get()}, ac)
}

// UpdateApprovalCategorys updates existing approval.category records.
// All records (represented by ids) will be updated by ac values.
func (c *Client) UpdateApprovalCategorys(ids []int64, ac *ApprovalCategory) error {
	return c.Update(ApprovalCategoryModel, ids, ac)
}

// DeleteApprovalCategory deletes an existing approval.category record.
func (c *Client) DeleteApprovalCategory(id int64) error {
	return c.DeleteApprovalCategorys([]int64{id})
}

// DeleteApprovalCategorys deletes existing approval.category records.
func (c *Client) DeleteApprovalCategorys(ids []int64) error {
	return c.Delete(ApprovalCategoryModel, ids)
}

// GetApprovalCategory gets approval.category existing record.
func (c *Client) GetApprovalCategory(id int64) (*ApprovalCategory, error) {
	acs, err := c.GetApprovalCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if acs != nil && len(*acs) > 0 {
		return &((*acs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of approval.category not found", id)
}

// GetApprovalCategorys gets approval.category existing records.
func (c *Client) GetApprovalCategorys(ids []int64) (*ApprovalCategorys, error) {
	acs := &ApprovalCategorys{}
	if err := c.Read(ApprovalCategoryModel, ids, nil, acs); err != nil {
		return nil, err
	}
	return acs, nil
}

// FindApprovalCategory finds approval.category record by querying it with criteria.
func (c *Client) FindApprovalCategory(criteria *Criteria) (*ApprovalCategory, error) {
	acs := &ApprovalCategorys{}
	if err := c.SearchRead(ApprovalCategoryModel, criteria, NewOptions().Limit(1), acs); err != nil {
		return nil, err
	}
	if acs != nil && len(*acs) > 0 {
		return &((*acs)[0]), nil
	}
	return nil, fmt.Errorf("approval.category was not found with criteria %v", criteria)
}

// FindApprovalCategorys finds approval.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalCategorys(criteria *Criteria, options *Options) (*ApprovalCategorys, error) {
	acs := &ApprovalCategorys{}
	if err := c.SearchRead(ApprovalCategoryModel, criteria, options, acs); err != nil {
		return nil, err
	}
	return acs, nil
}

// FindApprovalCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ApprovalCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindApprovalCategoryId finds record id by querying it with criteria.
func (c *Client) FindApprovalCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApprovalCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("approval.category was not found with criteria %v and options %v", criteria, options)
}
