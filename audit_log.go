package odoo

import (
	"fmt"
)

// AuditLog represents audit.log model.
type AuditLog struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Data        *String   `xmlrpc:"data,omptempty"`
	DataHtml    *String   `xmlrpc:"data_html,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Method      *String   `xmlrpc:"method,omptempty"`
	Model       *String   `xmlrpc:"model,omptempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ResId       *Int      `xmlrpc:"res_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AuditLogs represents array of audit.log model.
type AuditLogs []AuditLog

// AuditLogModel is the odoo model name.
const AuditLogModel = "audit.log"

// Many2One convert AuditLog to *Many2One.
func (al *AuditLog) Many2One() *Many2One {
	return NewMany2One(al.Id.Get(), "")
}

// CreateAuditLog creates a new audit.log model and returns its id.
func (c *Client) CreateAuditLog(al *AuditLog) (int64, error) {
	ids, err := c.CreateAuditLogs([]*AuditLog{al})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuditLog creates a new audit.log model and returns its id.
func (c *Client) CreateAuditLogs(als []*AuditLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range als {
		vv = append(vv, v)
	}
	return c.Create(AuditLogModel, vv)
}

// UpdateAuditLog updates an existing audit.log record.
func (c *Client) UpdateAuditLog(al *AuditLog) error {
	return c.UpdateAuditLogs([]int64{al.Id.Get()}, al)
}

// UpdateAuditLogs updates existing audit.log records.
// All records (represented by ids) will be updated by al values.
func (c *Client) UpdateAuditLogs(ids []int64, al *AuditLog) error {
	return c.Update(AuditLogModel, ids, al)
}

// DeleteAuditLog deletes an existing audit.log record.
func (c *Client) DeleteAuditLog(id int64) error {
	return c.DeleteAuditLogs([]int64{id})
}

// DeleteAuditLogs deletes existing audit.log records.
func (c *Client) DeleteAuditLogs(ids []int64) error {
	return c.Delete(AuditLogModel, ids)
}

// GetAuditLog gets audit.log existing record.
func (c *Client) GetAuditLog(id int64) (*AuditLog, error) {
	als, err := c.GetAuditLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	if als != nil && len(*als) > 0 {
		return &((*als)[0]), nil
	}
	return nil, fmt.Errorf("id %v of audit.log not found", id)
}

// GetAuditLogs gets audit.log existing records.
func (c *Client) GetAuditLogs(ids []int64) (*AuditLogs, error) {
	als := &AuditLogs{}
	if err := c.Read(AuditLogModel, ids, nil, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAuditLog finds audit.log record by querying it with criteria.
func (c *Client) FindAuditLog(criteria *Criteria) (*AuditLog, error) {
	als := &AuditLogs{}
	if err := c.SearchRead(AuditLogModel, criteria, NewOptions().Limit(1), als); err != nil {
		return nil, err
	}
	if als != nil && len(*als) > 0 {
		return &((*als)[0]), nil
	}
	return nil, fmt.Errorf("audit.log was not found with criteria %v", criteria)
}

// FindAuditLogs finds audit.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuditLogs(criteria *Criteria, options *Options) (*AuditLogs, error) {
	als := &AuditLogs{}
	if err := c.SearchRead(AuditLogModel, criteria, options, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAuditLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuditLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AuditLogModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAuditLogId finds record id by querying it with criteria.
func (c *Client) FindAuditLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuditLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("audit.log was not found with criteria %v and options %v", criteria, options)
}
