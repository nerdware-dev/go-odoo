package odoo

// AccountAnalyticGroup represents account.analytic.group model.
type AccountAnalyticGroup struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omitempty"`
	ChildrenIds  *Relation `xmlrpc:"children_ids,omitempty"`
	CompanyId    *Many2One `xmlrpc:"company_id,omitempty"`
	CompleteName *String   `xmlrpc:"complete_name,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	Description  *String   `xmlrpc:"description,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	ParentId     *Many2One `xmlrpc:"parent_id,omitempty"`
	ParentPath   *String   `xmlrpc:"parent_path,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticGroups represents array of account.analytic.group model.
type AccountAnalyticGroups []AccountAnalyticGroup

// AccountAnalyticGroupModel is the odoo model name.
const AccountAnalyticGroupModel = "account.analytic.group"

// Many2One convert AccountAnalyticGroup to *Many2One.
func (aag *AccountAnalyticGroup) Many2One() *Many2One {
	return NewMany2One(aag.Id.Get(), "")
}

// CreateAccountAnalyticGroup creates a new account.analytic.group model and returns its id.
func (c *Client) CreateAccountAnalyticGroup(aag *AccountAnalyticGroup) (int64, error) {
	ids, err := c.CreateAccountAnalyticGroups([]*AccountAnalyticGroup{aag})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticGroup creates a new account.analytic.group model and returns its id.
func (c *Client) CreateAccountAnalyticGroups(aags []*AccountAnalyticGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range aags {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticGroupModel, vv, nil)
}

// UpdateAccountAnalyticGroup updates an existing account.analytic.group record.
func (c *Client) UpdateAccountAnalyticGroup(aag *AccountAnalyticGroup) error {
	return c.UpdateAccountAnalyticGroups([]int64{aag.Id.Get()}, aag)
}

// UpdateAccountAnalyticGroups updates existing account.analytic.group records.
// All records (represented by ids) will be updated by aag values.
func (c *Client) UpdateAccountAnalyticGroups(ids []int64, aag *AccountAnalyticGroup) error {
	return c.Update(AccountAnalyticGroupModel, ids, aag, nil)
}

// DeleteAccountAnalyticGroup deletes an existing account.analytic.group record.
func (c *Client) DeleteAccountAnalyticGroup(id int64) error {
	return c.DeleteAccountAnalyticGroups([]int64{id})
}

// DeleteAccountAnalyticGroups deletes existing account.analytic.group records.
func (c *Client) DeleteAccountAnalyticGroups(ids []int64) error {
	return c.Delete(AccountAnalyticGroupModel, ids)
}

// GetAccountAnalyticGroup gets account.analytic.group existing record.
func (c *Client) GetAccountAnalyticGroup(id int64) (*AccountAnalyticGroup, error) {
	aags, err := c.GetAccountAnalyticGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aags)[0]), nil
}

// GetAccountAnalyticGroups gets account.analytic.group existing records.
func (c *Client) GetAccountAnalyticGroups(ids []int64) (*AccountAnalyticGroups, error) {
	aags := &AccountAnalyticGroups{}
	if err := c.Read(AccountAnalyticGroupModel, ids, nil, aags); err != nil {
		return nil, err
	}
	return aags, nil
}

// FindAccountAnalyticGroup finds account.analytic.group record by querying it with criteria.
func (c *Client) FindAccountAnalyticGroup(criteria *Criteria) (*AccountAnalyticGroup, error) {
	aags := &AccountAnalyticGroups{}
	if err := c.SearchRead(AccountAnalyticGroupModel, criteria, NewOptions().Limit(1), aags); err != nil {
		return nil, err
	}
	return &((*aags)[0]), nil
}

// FindAccountAnalyticGroups finds account.analytic.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticGroups(criteria *Criteria, options *Options) (*AccountAnalyticGroups, error) {
	aags := &AccountAnalyticGroups{}
	if err := c.SearchRead(AccountAnalyticGroupModel, criteria, options, aags); err != nil {
		return nil, err
	}
	return aags, nil
}

// FindAccountAnalyticGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticGroupModel, criteria, options)
}

// FindAccountAnalyticGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
