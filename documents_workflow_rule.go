package odoo

// DocumentsWorkflowRule represents documents.workflow.rule model.
type DocumentsWorkflowRule struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omitempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omitempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omitempty"`
	ActivityOption                *Bool      `xmlrpc:"activity_option,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	ConditionType                 *Selection `xmlrpc:"condition_type,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty"`
	CreateModel                   *Selection `xmlrpc:"create_model,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty"`
	CriteriaOwnerId               *Many2One  `xmlrpc:"criteria_owner_id,omitempty"`
	CriteriaPartnerId             *Many2One  `xmlrpc:"criteria_partner_id,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty"`
	Domain                        *String    `xmlrpc:"domain,omitempty"`
	DomainFolderId                *Many2One  `xmlrpc:"domain_folder_id,omitempty"`
	ExcludedTagIds                *Relation  `xmlrpc:"excluded_tag_ids,omitempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omitempty"`
	HasBusinessOption             *Bool      `xmlrpc:"has_business_option,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty"`
	LimitedToSingleRecord         *Bool      `xmlrpc:"limited_to_single_record,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	Note                          *String    `xmlrpc:"note,omitempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omitempty"`
	RemoveActivities              *Bool      `xmlrpc:"remove_activities,omitempty"`
	RequiredTagIds                *Relation  `xmlrpc:"required_tag_ids,omitempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omitempty"`
	TagActionIds                  *Relation  `xmlrpc:"tag_action_ids,omitempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DocumentsWorkflowRules represents array of documents.workflow.rule model.
type DocumentsWorkflowRules []DocumentsWorkflowRule

// DocumentsWorkflowRuleModel is the odoo model name.
const DocumentsWorkflowRuleModel = "documents.workflow.rule"

// Many2One convert DocumentsWorkflowRule to *Many2One.
func (dwr *DocumentsWorkflowRule) Many2One() *Many2One {
	return NewMany2One(dwr.Id.Get(), "")
}

// CreateDocumentsWorkflowRule creates a new documents.workflow.rule model and returns its id.
func (c *Client) CreateDocumentsWorkflowRule(dwr *DocumentsWorkflowRule) (int64, error) {
	ids, err := c.CreateDocumentsWorkflowRules([]*DocumentsWorkflowRule{dwr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsWorkflowRule creates a new documents.workflow.rule model and returns its id.
func (c *Client) CreateDocumentsWorkflowRules(dwrs []*DocumentsWorkflowRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range dwrs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsWorkflowRuleModel, vv, nil)
}

// UpdateDocumentsWorkflowRule updates an existing documents.workflow.rule record.
func (c *Client) UpdateDocumentsWorkflowRule(dwr *DocumentsWorkflowRule) error {
	return c.UpdateDocumentsWorkflowRules([]int64{dwr.Id.Get()}, dwr)
}

// UpdateDocumentsWorkflowRules updates existing documents.workflow.rule records.
// All records (represented by ids) will be updated by dwr values.
func (c *Client) UpdateDocumentsWorkflowRules(ids []int64, dwr *DocumentsWorkflowRule) error {
	return c.Update(DocumentsWorkflowRuleModel, ids, dwr, nil)
}

// DeleteDocumentsWorkflowRule deletes an existing documents.workflow.rule record.
func (c *Client) DeleteDocumentsWorkflowRule(id int64) error {
	return c.DeleteDocumentsWorkflowRules([]int64{id})
}

// DeleteDocumentsWorkflowRules deletes existing documents.workflow.rule records.
func (c *Client) DeleteDocumentsWorkflowRules(ids []int64) error {
	return c.Delete(DocumentsWorkflowRuleModel, ids)
}

// GetDocumentsWorkflowRule gets documents.workflow.rule existing record.
func (c *Client) GetDocumentsWorkflowRule(id int64) (*DocumentsWorkflowRule, error) {
	dwrs, err := c.GetDocumentsWorkflowRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dwrs)[0]), nil
}

// GetDocumentsWorkflowRules gets documents.workflow.rule existing records.
func (c *Client) GetDocumentsWorkflowRules(ids []int64) (*DocumentsWorkflowRules, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.Read(DocumentsWorkflowRuleModel, ids, nil, dwrs); err != nil {
		return nil, err
	}
	return dwrs, nil
}

// FindDocumentsWorkflowRule finds documents.workflow.rule record by querying it with criteria.
func (c *Client) FindDocumentsWorkflowRule(criteria *Criteria) (*DocumentsWorkflowRule, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.SearchRead(DocumentsWorkflowRuleModel, criteria, NewOptions().Limit(1), dwrs); err != nil {
		return nil, err
	}
	return &((*dwrs)[0]), nil
}

// FindDocumentsWorkflowRules finds documents.workflow.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowRules(criteria *Criteria, options *Options) (*DocumentsWorkflowRules, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.SearchRead(DocumentsWorkflowRuleModel, criteria, options, dwrs); err != nil {
		return nil, err
	}
	return dwrs, nil
}

// FindDocumentsWorkflowRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsWorkflowRuleModel, criteria, options)
}

// FindDocumentsWorkflowRuleId finds record id by querying it with criteria.
func (c *Client) FindDocumentsWorkflowRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsWorkflowRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
