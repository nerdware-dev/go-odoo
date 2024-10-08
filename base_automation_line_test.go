package odoo

// BaseAutomationLineTest represents base.automation.line.test model.
type BaseAutomationLineTest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LeadId      *Many2One `xmlrpc:"lead_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseAutomationLineTests represents array of base.automation.line.test model.
type BaseAutomationLineTests []BaseAutomationLineTest

// BaseAutomationLineTestModel is the odoo model name.
const BaseAutomationLineTestModel = "base.automation.line.test"

// Many2One convert BaseAutomationLineTest to *Many2One.
func (balt *BaseAutomationLineTest) Many2One() *Many2One {
	return NewMany2One(balt.Id.Get(), "")
}

// CreateBaseAutomationLineTest creates a new base.automation.line.test model and returns its id.
func (c *Client) CreateBaseAutomationLineTest(balt *BaseAutomationLineTest) (int64, error) {
	ids, err := c.CreateBaseAutomationLineTests([]*BaseAutomationLineTest{balt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseAutomationLineTest creates a new base.automation.line.test model and returns its id.
func (c *Client) CreateBaseAutomationLineTests(balts []*BaseAutomationLineTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range balts {
		vv = append(vv, v)
	}
	return c.Create(BaseAutomationLineTestModel, vv, nil)
}

// UpdateBaseAutomationLineTest updates an existing base.automation.line.test record.
func (c *Client) UpdateBaseAutomationLineTest(balt *BaseAutomationLineTest) error {
	return c.UpdateBaseAutomationLineTests([]int64{balt.Id.Get()}, balt)
}

// UpdateBaseAutomationLineTests updates existing base.automation.line.test records.
// All records (represented by ids) will be updated by balt values.
func (c *Client) UpdateBaseAutomationLineTests(ids []int64, balt *BaseAutomationLineTest) error {
	return c.Update(BaseAutomationLineTestModel, ids, balt, nil)
}

// DeleteBaseAutomationLineTest deletes an existing base.automation.line.test record.
func (c *Client) DeleteBaseAutomationLineTest(id int64) error {
	return c.DeleteBaseAutomationLineTests([]int64{id})
}

// DeleteBaseAutomationLineTests deletes existing base.automation.line.test records.
func (c *Client) DeleteBaseAutomationLineTests(ids []int64) error {
	return c.Delete(BaseAutomationLineTestModel, ids)
}

// GetBaseAutomationLineTest gets base.automation.line.test existing record.
func (c *Client) GetBaseAutomationLineTest(id int64) (*BaseAutomationLineTest, error) {
	balts, err := c.GetBaseAutomationLineTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*balts)[0]), nil
}

// GetBaseAutomationLineTests gets base.automation.line.test existing records.
func (c *Client) GetBaseAutomationLineTests(ids []int64) (*BaseAutomationLineTests, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.Read(BaseAutomationLineTestModel, ids, nil, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLineTest finds base.automation.line.test record by querying it with criteria.
func (c *Client) FindBaseAutomationLineTest(criteria *Criteria) (*BaseAutomationLineTest, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.SearchRead(BaseAutomationLineTestModel, criteria, NewOptions().Limit(1), balts); err != nil {
		return nil, err
	}
	return &((*balts)[0]), nil
}

// FindBaseAutomationLineTests finds base.automation.line.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLineTests(criteria *Criteria, options *Options) (*BaseAutomationLineTests, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.SearchRead(BaseAutomationLineTestModel, criteria, options, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLineTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLineTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseAutomationLineTestModel, criteria, options)
}

// FindBaseAutomationLineTestId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationLineTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationLineTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
