package odoo

// AuditLog represents audit.log model.
type AuditLog struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Data        *String   `xmlrpc:"data,omitempty"`
	DataHtml    *String   `xmlrpc:"data_html,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Method      *String   `xmlrpc:"method,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ResId       *Int      `xmlrpc:"res_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(AuditLogModel, vv, nil)
}

// UpdateAuditLog updates an existing audit.log record.
func (c *Client) UpdateAuditLog(al *AuditLog) error {
	return c.UpdateAuditLogs([]int64{al.Id.Get()}, al)
}

// UpdateAuditLogs updates existing audit.log records.
// All records (represented by ids) will be updated by al values.
func (c *Client) UpdateAuditLogs(ids []int64, al *AuditLog) error {
	return c.Update(AuditLogModel, ids, al, nil)
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
	return &((*als)[0]), nil
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
	return &((*als)[0]), nil
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
	return c.Search(AuditLogModel, criteria, options)
}

// FindAuditLogId finds record id by querying it with criteria.
func (c *Client) FindAuditLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuditLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
