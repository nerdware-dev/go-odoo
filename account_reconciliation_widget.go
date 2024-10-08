package odoo

// AccountReconciliationWidget represents account.reconciliation.widget model.
type AccountReconciliationWidget struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountReconciliationWidgets represents array of account.reconciliation.widget model.
type AccountReconciliationWidgets []AccountReconciliationWidget

// AccountReconciliationWidgetModel is the odoo model name.
const AccountReconciliationWidgetModel = "account.reconciliation.widget"

// Many2One convert AccountReconciliationWidget to *Many2One.
func (arw *AccountReconciliationWidget) Many2One() *Many2One {
	return NewMany2One(arw.Id.Get(), "")
}

// CreateAccountReconciliationWidget creates a new account.reconciliation.widget model and returns its id.
func (c *Client) CreateAccountReconciliationWidget(arw *AccountReconciliationWidget) (int64, error) {
	ids, err := c.CreateAccountReconciliationWidgets([]*AccountReconciliationWidget{arw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconciliationWidget creates a new account.reconciliation.widget model and returns its id.
func (c *Client) CreateAccountReconciliationWidgets(arws []*AccountReconciliationWidget) ([]int64, error) {
	var vv []interface{}
	for _, v := range arws {
		vv = append(vv, v)
	}
	return c.Create(AccountReconciliationWidgetModel, vv, nil)
}

// UpdateAccountReconciliationWidget updates an existing account.reconciliation.widget record.
func (c *Client) UpdateAccountReconciliationWidget(arw *AccountReconciliationWidget) error {
	return c.UpdateAccountReconciliationWidgets([]int64{arw.Id.Get()}, arw)
}

// UpdateAccountReconciliationWidgets updates existing account.reconciliation.widget records.
// All records (represented by ids) will be updated by arw values.
func (c *Client) UpdateAccountReconciliationWidgets(ids []int64, arw *AccountReconciliationWidget) error {
	return c.Update(AccountReconciliationWidgetModel, ids, arw, nil)
}

// DeleteAccountReconciliationWidget deletes an existing account.reconciliation.widget record.
func (c *Client) DeleteAccountReconciliationWidget(id int64) error {
	return c.DeleteAccountReconciliationWidgets([]int64{id})
}

// DeleteAccountReconciliationWidgets deletes existing account.reconciliation.widget records.
func (c *Client) DeleteAccountReconciliationWidgets(ids []int64) error {
	return c.Delete(AccountReconciliationWidgetModel, ids)
}

// GetAccountReconciliationWidget gets account.reconciliation.widget existing record.
func (c *Client) GetAccountReconciliationWidget(id int64) (*AccountReconciliationWidget, error) {
	arws, err := c.GetAccountReconciliationWidgets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arws)[0]), nil
}

// GetAccountReconciliationWidgets gets account.reconciliation.widget existing records.
func (c *Client) GetAccountReconciliationWidgets(ids []int64) (*AccountReconciliationWidgets, error) {
	arws := &AccountReconciliationWidgets{}
	if err := c.Read(AccountReconciliationWidgetModel, ids, nil, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountReconciliationWidget finds account.reconciliation.widget record by querying it with criteria.
func (c *Client) FindAccountReconciliationWidget(criteria *Criteria) (*AccountReconciliationWidget, error) {
	arws := &AccountReconciliationWidgets{}
	if err := c.SearchRead(AccountReconciliationWidgetModel, criteria, NewOptions().Limit(1), arws); err != nil {
		return nil, err
	}
	return &((*arws)[0]), nil
}

// FindAccountReconciliationWidgets finds account.reconciliation.widget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconciliationWidgets(criteria *Criteria, options *Options) (*AccountReconciliationWidgets, error) {
	arws := &AccountReconciliationWidgets{}
	if err := c.SearchRead(AccountReconciliationWidgetModel, criteria, options, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountReconciliationWidgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconciliationWidgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconciliationWidgetModel, criteria, options)
}

// FindAccountReconciliationWidgetId finds record id by querying it with criteria.
func (c *Client) FindAccountReconciliationWidgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconciliationWidgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
