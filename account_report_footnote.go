package odoo

// AccountReportFootnote represents account.report.footnote model.
type AccountReportFootnote struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Line        *String   `xmlrpc:"line,omitempty"`
	ManagerId   *Many2One `xmlrpc:"manager_id,omitempty"`
	Text        *String   `xmlrpc:"text,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReportFootnotes represents array of account.report.footnote model.
type AccountReportFootnotes []AccountReportFootnote

// AccountReportFootnoteModel is the odoo model name.
const AccountReportFootnoteModel = "account.report.footnote"

// Many2One convert AccountReportFootnote to *Many2One.
func (arf *AccountReportFootnote) Many2One() *Many2One {
	return NewMany2One(arf.Id.Get(), "")
}

// CreateAccountReportFootnote creates a new account.report.footnote model and returns its id.
func (c *Client) CreateAccountReportFootnote(arf *AccountReportFootnote) (int64, error) {
	ids, err := c.CreateAccountReportFootnotes([]*AccountReportFootnote{arf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportFootnote creates a new account.report.footnote model and returns its id.
func (c *Client) CreateAccountReportFootnotes(arfs []*AccountReportFootnote) ([]int64, error) {
	var vv []interface{}
	for _, v := range arfs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportFootnoteModel, vv, nil)
}

// UpdateAccountReportFootnote updates an existing account.report.footnote record.
func (c *Client) UpdateAccountReportFootnote(arf *AccountReportFootnote) error {
	return c.UpdateAccountReportFootnotes([]int64{arf.Id.Get()}, arf)
}

// UpdateAccountReportFootnotes updates existing account.report.footnote records.
// All records (represented by ids) will be updated by arf values.
func (c *Client) UpdateAccountReportFootnotes(ids []int64, arf *AccountReportFootnote) error {
	return c.Update(AccountReportFootnoteModel, ids, arf, nil)
}

// DeleteAccountReportFootnote deletes an existing account.report.footnote record.
func (c *Client) DeleteAccountReportFootnote(id int64) error {
	return c.DeleteAccountReportFootnotes([]int64{id})
}

// DeleteAccountReportFootnotes deletes existing account.report.footnote records.
func (c *Client) DeleteAccountReportFootnotes(ids []int64) error {
	return c.Delete(AccountReportFootnoteModel, ids)
}

// GetAccountReportFootnote gets account.report.footnote existing record.
func (c *Client) GetAccountReportFootnote(id int64) (*AccountReportFootnote, error) {
	arfs, err := c.GetAccountReportFootnotes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arfs)[0]), nil
}

// GetAccountReportFootnotes gets account.report.footnote existing records.
func (c *Client) GetAccountReportFootnotes(ids []int64) (*AccountReportFootnotes, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.Read(AccountReportFootnoteModel, ids, nil, arfs); err != nil {
		return nil, err
	}
	return arfs, nil
}

// FindAccountReportFootnote finds account.report.footnote record by querying it with criteria.
func (c *Client) FindAccountReportFootnote(criteria *Criteria) (*AccountReportFootnote, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.SearchRead(AccountReportFootnoteModel, criteria, NewOptions().Limit(1), arfs); err != nil {
		return nil, err
	}
	return &((*arfs)[0]), nil
}

// FindAccountReportFootnotes finds account.report.footnote records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFootnotes(criteria *Criteria, options *Options) (*AccountReportFootnotes, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.SearchRead(AccountReportFootnoteModel, criteria, options, arfs); err != nil {
		return nil, err
	}
	return arfs, nil
}

// FindAccountReportFootnoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFootnoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportFootnoteModel, criteria, options)
}

// FindAccountReportFootnoteId finds record id by querying it with criteria.
func (c *Client) FindAccountReportFootnoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportFootnoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
