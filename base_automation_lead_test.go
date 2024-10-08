package odoo

// BaseAutomationLeadTest represents base.automation.lead.test model.
type BaseAutomationLeadTest struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	Active            *Bool      `xmlrpc:"active,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateActionLast    *Time      `xmlrpc:"date_action_last,omitempty"`
	Deadline          *Bool      `xmlrpc:"deadline,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Employee          *Bool      `xmlrpc:"employee,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IsAssignedToAdmin *Bool      `xmlrpc:"is_assigned_to_admin,omitempty"`
	LineIds           *Relation  `xmlrpc:"line_ids,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	Priority          *Bool      `xmlrpc:"priority,omitempty"`
	State             *Selection `xmlrpc:"state,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BaseAutomationLeadTests represents array of base.automation.lead.test model.
type BaseAutomationLeadTests []BaseAutomationLeadTest

// BaseAutomationLeadTestModel is the odoo model name.
const BaseAutomationLeadTestModel = "base.automation.lead.test"

// Many2One convert BaseAutomationLeadTest to *Many2One.
func (balt *BaseAutomationLeadTest) Many2One() *Many2One {
	return NewMany2One(balt.Id.Get(), "")
}

// CreateBaseAutomationLeadTest creates a new base.automation.lead.test model and returns its id.
func (c *Client) CreateBaseAutomationLeadTest(balt *BaseAutomationLeadTest) (int64, error) {
	ids, err := c.CreateBaseAutomationLeadTests([]*BaseAutomationLeadTest{balt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseAutomationLeadTest creates a new base.automation.lead.test model and returns its id.
func (c *Client) CreateBaseAutomationLeadTests(balts []*BaseAutomationLeadTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range balts {
		vv = append(vv, v)
	}
	return c.Create(BaseAutomationLeadTestModel, vv, nil)
}

// UpdateBaseAutomationLeadTest updates an existing base.automation.lead.test record.
func (c *Client) UpdateBaseAutomationLeadTest(balt *BaseAutomationLeadTest) error {
	return c.UpdateBaseAutomationLeadTests([]int64{balt.Id.Get()}, balt)
}

// UpdateBaseAutomationLeadTests updates existing base.automation.lead.test records.
// All records (represented by ids) will be updated by balt values.
func (c *Client) UpdateBaseAutomationLeadTests(ids []int64, balt *BaseAutomationLeadTest) error {
	return c.Update(BaseAutomationLeadTestModel, ids, balt, nil)
}

// DeleteBaseAutomationLeadTest deletes an existing base.automation.lead.test record.
func (c *Client) DeleteBaseAutomationLeadTest(id int64) error {
	return c.DeleteBaseAutomationLeadTests([]int64{id})
}

// DeleteBaseAutomationLeadTests deletes existing base.automation.lead.test records.
func (c *Client) DeleteBaseAutomationLeadTests(ids []int64) error {
	return c.Delete(BaseAutomationLeadTestModel, ids)
}

// GetBaseAutomationLeadTest gets base.automation.lead.test existing record.
func (c *Client) GetBaseAutomationLeadTest(id int64) (*BaseAutomationLeadTest, error) {
	balts, err := c.GetBaseAutomationLeadTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*balts)[0]), nil
}

// GetBaseAutomationLeadTests gets base.automation.lead.test existing records.
func (c *Client) GetBaseAutomationLeadTests(ids []int64) (*BaseAutomationLeadTests, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.Read(BaseAutomationLeadTestModel, ids, nil, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLeadTest finds base.automation.lead.test record by querying it with criteria.
func (c *Client) FindBaseAutomationLeadTest(criteria *Criteria) (*BaseAutomationLeadTest, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.SearchRead(BaseAutomationLeadTestModel, criteria, NewOptions().Limit(1), balts); err != nil {
		return nil, err
	}
	return &((*balts)[0]), nil
}

// FindBaseAutomationLeadTests finds base.automation.lead.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLeadTests(criteria *Criteria, options *Options) (*BaseAutomationLeadTests, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.SearchRead(BaseAutomationLeadTestModel, criteria, options, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLeadTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLeadTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseAutomationLeadTestModel, criteria, options)
}

// FindBaseAutomationLeadTestId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationLeadTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationLeadTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
