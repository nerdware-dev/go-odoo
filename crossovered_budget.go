package odoo

// CrossoveredBudget represents crossovered.budget model.
type CrossoveredBudget struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrossoveredBudgetLine       *Relation  `xmlrpc:"crossovered_budget_line,omitempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CrossoveredBudgets represents array of crossovered.budget model.
type CrossoveredBudgets []CrossoveredBudget

// CrossoveredBudgetModel is the odoo model name.
const CrossoveredBudgetModel = "crossovered.budget"

// Many2One convert CrossoveredBudget to *Many2One.
func (cb *CrossoveredBudget) Many2One() *Many2One {
	return NewMany2One(cb.Id.Get(), "")
}

// CreateCrossoveredBudget creates a new crossovered.budget model and returns its id.
func (c *Client) CreateCrossoveredBudget(cb *CrossoveredBudget) (int64, error) {
	ids, err := c.CreateCrossoveredBudgets([]*CrossoveredBudget{cb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrossoveredBudget creates a new crossovered.budget model and returns its id.
func (c *Client) CreateCrossoveredBudgets(cbs []*CrossoveredBudget) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbs {
		vv = append(vv, v)
	}
	return c.Create(CrossoveredBudgetModel, vv, nil)
}

// UpdateCrossoveredBudget updates an existing crossovered.budget record.
func (c *Client) UpdateCrossoveredBudget(cb *CrossoveredBudget) error {
	return c.UpdateCrossoveredBudgets([]int64{cb.Id.Get()}, cb)
}

// UpdateCrossoveredBudgets updates existing crossovered.budget records.
// All records (represented by ids) will be updated by cb values.
func (c *Client) UpdateCrossoveredBudgets(ids []int64, cb *CrossoveredBudget) error {
	return c.Update(CrossoveredBudgetModel, ids, cb, nil)
}

// DeleteCrossoveredBudget deletes an existing crossovered.budget record.
func (c *Client) DeleteCrossoveredBudget(id int64) error {
	return c.DeleteCrossoveredBudgets([]int64{id})
}

// DeleteCrossoveredBudgets deletes existing crossovered.budget records.
func (c *Client) DeleteCrossoveredBudgets(ids []int64) error {
	return c.Delete(CrossoveredBudgetModel, ids)
}

// GetCrossoveredBudget gets crossovered.budget existing record.
func (c *Client) GetCrossoveredBudget(id int64) (*CrossoveredBudget, error) {
	cbs, err := c.GetCrossoveredBudgets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cbs)[0]), nil
}

// GetCrossoveredBudgets gets crossovered.budget existing records.
func (c *Client) GetCrossoveredBudgets(ids []int64) (*CrossoveredBudgets, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.Read(CrossoveredBudgetModel, ids, nil, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCrossoveredBudget finds crossovered.budget record by querying it with criteria.
func (c *Client) FindCrossoveredBudget(criteria *Criteria) (*CrossoveredBudget, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.SearchRead(CrossoveredBudgetModel, criteria, NewOptions().Limit(1), cbs); err != nil {
		return nil, err
	}
	return &((*cbs)[0]), nil
}

// FindCrossoveredBudgets finds crossovered.budget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgets(criteria *Criteria, options *Options) (*CrossoveredBudgets, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.SearchRead(CrossoveredBudgetModel, criteria, options, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCrossoveredBudgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrossoveredBudgetModel, criteria, options)
}

// FindCrossoveredBudgetId finds record id by querying it with criteria.
func (c *Client) FindCrossoveredBudgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrossoveredBudgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
