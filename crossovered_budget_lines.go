package odoo

// CrossoveredBudgetLines represents crossovered.budget.lines model.
type CrossoveredBudgetLines struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	AnalyticAccountId      *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AnalyticGroupId        *Many2One  `xmlrpc:"analytic_group_id,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrossoveredBudgetId    *Many2One  `xmlrpc:"crossovered_budget_id,omitempty"`
	CrossoveredBudgetState *Selection `xmlrpc:"crossovered_budget_state,omitempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateFrom               *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                 *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	GeneralBudgetId        *Many2One  `xmlrpc:"general_budget_id,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	IsAboveBudget          *Bool      `xmlrpc:"is_above_budget,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	PaidDate               *Time      `xmlrpc:"paid_date,omitempty"`
	Percentage             *Float     `xmlrpc:"percentage,omitempty"`
	PlannedAmount          *Float     `xmlrpc:"planned_amount,omitempty"`
	PracticalAmount        *Float     `xmlrpc:"practical_amount,omitempty"`
	TheoriticalAmount      *Float     `xmlrpc:"theoritical_amount,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CrossoveredBudgetLiness represents array of crossovered.budget.lines model.
type CrossoveredBudgetLiness []CrossoveredBudgetLines

// CrossoveredBudgetLinesModel is the odoo model name.
const CrossoveredBudgetLinesModel = "crossovered.budget.lines"

// Many2One convert CrossoveredBudgetLines to *Many2One.
func (cbl *CrossoveredBudgetLines) Many2One() *Many2One {
	return NewMany2One(cbl.Id.Get(), "")
}

// CreateCrossoveredBudgetLines creates a new crossovered.budget.lines model and returns its id.
func (c *Client) CreateCrossoveredBudgetLines(cbl *CrossoveredBudgetLines) (int64, error) {
	ids, err := c.CreateCrossoveredBudgetLiness([]*CrossoveredBudgetLines{cbl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrossoveredBudgetLines creates a new crossovered.budget.lines model and returns its id.
func (c *Client) CreateCrossoveredBudgetLiness(cbls []*CrossoveredBudgetLines) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbls {
		vv = append(vv, v)
	}
	return c.Create(CrossoveredBudgetLinesModel, vv, nil)
}

// UpdateCrossoveredBudgetLines updates an existing crossovered.budget.lines record.
func (c *Client) UpdateCrossoveredBudgetLines(cbl *CrossoveredBudgetLines) error {
	return c.UpdateCrossoveredBudgetLiness([]int64{cbl.Id.Get()}, cbl)
}

// UpdateCrossoveredBudgetLiness updates existing crossovered.budget.lines records.
// All records (represented by ids) will be updated by cbl values.
func (c *Client) UpdateCrossoveredBudgetLiness(ids []int64, cbl *CrossoveredBudgetLines) error {
	return c.Update(CrossoveredBudgetLinesModel, ids, cbl, nil)
}

// DeleteCrossoveredBudgetLines deletes an existing crossovered.budget.lines record.
func (c *Client) DeleteCrossoveredBudgetLines(id int64) error {
	return c.DeleteCrossoveredBudgetLiness([]int64{id})
}

// DeleteCrossoveredBudgetLiness deletes existing crossovered.budget.lines records.
func (c *Client) DeleteCrossoveredBudgetLiness(ids []int64) error {
	return c.Delete(CrossoveredBudgetLinesModel, ids)
}

// GetCrossoveredBudgetLines gets crossovered.budget.lines existing record.
func (c *Client) GetCrossoveredBudgetLines(id int64) (*CrossoveredBudgetLines, error) {
	cbls, err := c.GetCrossoveredBudgetLiness([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cbls)[0]), nil
}

// GetCrossoveredBudgetLiness gets crossovered.budget.lines existing records.
func (c *Client) GetCrossoveredBudgetLiness(ids []int64) (*CrossoveredBudgetLiness, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.Read(CrossoveredBudgetLinesModel, ids, nil, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCrossoveredBudgetLines finds crossovered.budget.lines record by querying it with criteria.
func (c *Client) FindCrossoveredBudgetLines(criteria *Criteria) (*CrossoveredBudgetLines, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.SearchRead(CrossoveredBudgetLinesModel, criteria, NewOptions().Limit(1), cbls); err != nil {
		return nil, err
	}
	return &((*cbls)[0]), nil
}

// FindCrossoveredBudgetLiness finds crossovered.budget.lines records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetLiness(criteria *Criteria, options *Options) (*CrossoveredBudgetLiness, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.SearchRead(CrossoveredBudgetLinesModel, criteria, options, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCrossoveredBudgetLinesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetLinesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrossoveredBudgetLinesModel, criteria, options)
}

// FindCrossoveredBudgetLinesId finds record id by querying it with criteria.
func (c *Client) FindCrossoveredBudgetLinesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrossoveredBudgetLinesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
