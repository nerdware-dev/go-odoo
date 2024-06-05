package odoo

import (
	"fmt"
)

// AuditRule represents audit.rule model.
type AuditRule struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId    *Many2One  `xmlrpc:"action_id,omptempty"`
	Active      *Bool      `xmlrpc:"active,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	GroupId     *Many2One  `xmlrpc:"group_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LogCreate   *Bool      `xmlrpc:"log_create,omptempty"`
	LogUnlink   *Bool      `xmlrpc:"log_unlink,omptempty"`
	LogWrite    *Bool      `xmlrpc:"log_write,omptempty"`
	ModelId     *Many2One  `xmlrpc:"model_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AuditRules represents array of audit.rule model.
type AuditRules []AuditRule

// AuditRuleModel is the odoo model name.
const AuditRuleModel = "audit.rule"

// Many2One convert AuditRule to *Many2One.
func (ar *AuditRule) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAuditRule creates a new audit.rule model and returns its id.
func (c *Client) CreateAuditRule(ar *AuditRule) (int64, error) {
	ids, err := c.CreateAuditRules([]*AuditRule{ar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuditRule creates a new audit.rule model and returns its id.
func (c *Client) CreateAuditRules(ars []*AuditRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range ars {
		vv = append(vv, v)
	}
	return c.Create(AuditRuleModel, vv)
}

// UpdateAuditRule updates an existing audit.rule record.
func (c *Client) UpdateAuditRule(ar *AuditRule) error {
	return c.UpdateAuditRules([]int64{ar.Id.Get()}, ar)
}

// UpdateAuditRules updates existing audit.rule records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAuditRules(ids []int64, ar *AuditRule) error {
	return c.Update(AuditRuleModel, ids, ar)
}

// DeleteAuditRule deletes an existing audit.rule record.
func (c *Client) DeleteAuditRule(id int64) error {
	return c.DeleteAuditRules([]int64{id})
}

// DeleteAuditRules deletes existing audit.rule records.
func (c *Client) DeleteAuditRules(ids []int64) error {
	return c.Delete(AuditRuleModel, ids)
}

// GetAuditRule gets audit.rule existing record.
func (c *Client) GetAuditRule(id int64) (*AuditRule, error) {
	ars, err := c.GetAuditRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of audit.rule not found", id)
}

// GetAuditRules gets audit.rule existing records.
func (c *Client) GetAuditRules(ids []int64) (*AuditRules, error) {
	ars := &AuditRules{}
	if err := c.Read(AuditRuleModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAuditRule finds audit.rule record by querying it with criteria.
func (c *Client) FindAuditRule(criteria *Criteria) (*AuditRule, error) {
	ars := &AuditRules{}
	if err := c.SearchRead(AuditRuleModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("audit.rule was not found with criteria %v", criteria)
}

// FindAuditRules finds audit.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuditRules(criteria *Criteria, options *Options) (*AuditRules, error) {
	ars := &AuditRules{}
	if err := c.SearchRead(AuditRuleModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAuditRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuditRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AuditRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAuditRuleId finds record id by querying it with criteria.
func (c *Client) FindAuditRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuditRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("audit.rule was not found with criteria %v and options %v", criteria, options)
}
