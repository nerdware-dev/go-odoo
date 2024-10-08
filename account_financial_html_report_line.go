package odoo

// AccountFinancialHtmlReportLine represents account.financial.html.report.line model.
type AccountFinancialHtmlReportLine struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	ActionId           *Many2One  `xmlrpc:"action_id,omitempty"`
	ChildrenIds        *Relation  `xmlrpc:"children_ids,omitempty"`
	Code               *String    `xmlrpc:"code,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Domain             *String    `xmlrpc:"domain,omitempty"`
	FigureType         *Selection `xmlrpc:"figure_type,omitempty"`
	FinancialReportId  *Many2One  `xmlrpc:"financial_report_id,omitempty"`
	Formulas           *String    `xmlrpc:"formulas,omitempty"`
	GreenOnPositive    *Bool      `xmlrpc:"green_on_positive,omitempty"`
	Groupby            *String    `xmlrpc:"groupby,omitempty"`
	HideIfEmpty        *Bool      `xmlrpc:"hide_if_empty,omitempty"`
	HideIfZero         *Bool      `xmlrpc:"hide_if_zero,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	Level              *Int       `xmlrpc:"level,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentPath         *String    `xmlrpc:"parent_path,omitempty"`
	PrintOnNewPage     *Bool      `xmlrpc:"print_on_new_page,omitempty"`
	Sequence           *Int       `xmlrpc:"sequence,omitempty"`
	ShowDomain         *Selection `xmlrpc:"show_domain,omitempty"`
	SpecialDateChanger *Selection `xmlrpc:"special_date_changer,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFinancialHtmlReportLines represents array of account.financial.html.report.line model.
type AccountFinancialHtmlReportLines []AccountFinancialHtmlReportLine

// AccountFinancialHtmlReportLineModel is the odoo model name.
const AccountFinancialHtmlReportLineModel = "account.financial.html.report.line"

// Many2One convert AccountFinancialHtmlReportLine to *Many2One.
func (afhrl *AccountFinancialHtmlReportLine) Many2One() *Many2One {
	return NewMany2One(afhrl.Id.Get(), "")
}

// CreateAccountFinancialHtmlReportLine creates a new account.financial.html.report.line model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReportLine(afhrl *AccountFinancialHtmlReportLine) (int64, error) {
	ids, err := c.CreateAccountFinancialHtmlReportLines([]*AccountFinancialHtmlReportLine{afhrl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFinancialHtmlReportLine creates a new account.financial.html.report.line model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReportLines(afhrls []*AccountFinancialHtmlReportLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range afhrls {
		vv = append(vv, v)
	}
	return c.Create(AccountFinancialHtmlReportLineModel, vv, nil)
}

// UpdateAccountFinancialHtmlReportLine updates an existing account.financial.html.report.line record.
func (c *Client) UpdateAccountFinancialHtmlReportLine(afhrl *AccountFinancialHtmlReportLine) error {
	return c.UpdateAccountFinancialHtmlReportLines([]int64{afhrl.Id.Get()}, afhrl)
}

// UpdateAccountFinancialHtmlReportLines updates existing account.financial.html.report.line records.
// All records (represented by ids) will be updated by afhrl values.
func (c *Client) UpdateAccountFinancialHtmlReportLines(ids []int64, afhrl *AccountFinancialHtmlReportLine) error {
	return c.Update(AccountFinancialHtmlReportLineModel, ids, afhrl, nil)
}

// DeleteAccountFinancialHtmlReportLine deletes an existing account.financial.html.report.line record.
func (c *Client) DeleteAccountFinancialHtmlReportLine(id int64) error {
	return c.DeleteAccountFinancialHtmlReportLines([]int64{id})
}

// DeleteAccountFinancialHtmlReportLines deletes existing account.financial.html.report.line records.
func (c *Client) DeleteAccountFinancialHtmlReportLines(ids []int64) error {
	return c.Delete(AccountFinancialHtmlReportLineModel, ids)
}

// GetAccountFinancialHtmlReportLine gets account.financial.html.report.line existing record.
func (c *Client) GetAccountFinancialHtmlReportLine(id int64) (*AccountFinancialHtmlReportLine, error) {
	afhrls, err := c.GetAccountFinancialHtmlReportLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afhrls)[0]), nil
}

// GetAccountFinancialHtmlReportLines gets account.financial.html.report.line existing records.
func (c *Client) GetAccountFinancialHtmlReportLines(ids []int64) (*AccountFinancialHtmlReportLines, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.Read(AccountFinancialHtmlReportLineModel, ids, nil, afhrls); err != nil {
		return nil, err
	}
	return afhrls, nil
}

// FindAccountFinancialHtmlReportLine finds account.financial.html.report.line record by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportLine(criteria *Criteria) (*AccountFinancialHtmlReportLine, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.SearchRead(AccountFinancialHtmlReportLineModel, criteria, NewOptions().Limit(1), afhrls); err != nil {
		return nil, err
	}
	return &((*afhrls)[0]), nil
}

// FindAccountFinancialHtmlReportLines finds account.financial.html.report.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportLines(criteria *Criteria, options *Options) (*AccountFinancialHtmlReportLines, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.SearchRead(AccountFinancialHtmlReportLineModel, criteria, options, afhrls); err != nil {
		return nil, err
	}
	return afhrls, nil
}

// FindAccountFinancialHtmlReportLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFinancialHtmlReportLineModel, criteria, options)
}

// FindAccountFinancialHtmlReportLineId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
