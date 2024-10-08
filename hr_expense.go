package odoo

// HrExpense represents hr.expense model.
type HrExpense struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId                   *Many2One  `xmlrpc:"account_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AnalyticAccountId           *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AnalyticTagIds              *Relation  `xmlrpc:"analytic_tag_ids,omitempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omitempty"`
	CanBeReinvoiced             *Bool      `xmlrpc:"can_be_reinvoiced,omitempty"`
	CompanyCurrencyId           *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsEditable                  *Bool      `xmlrpc:"is_editable,omitempty"`
	IsRefEditable               *Bool      `xmlrpc:"is_ref_editable,omitempty"`
	IsRefused                   *Bool      `xmlrpc:"is_refused,omitempty"`
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
	PaymentMode                 *Selection `xmlrpc:"payment_mode,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomCategoryId        *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	Quantity                    *Float     `xmlrpc:"quantity,omitempty"`
	Reference                   *String    `xmlrpc:"reference,omitempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	SheetId                     *Many2One  `xmlrpc:"sheet_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TaxIds                      *Relation  `xmlrpc:"tax_ids,omitempty"`
	TotalAmount                 *Float     `xmlrpc:"total_amount,omitempty"`
	TotalAmountCompany          *Float     `xmlrpc:"total_amount_company,omitempty"`
	UnitAmount                  *Float     `xmlrpc:"unit_amount,omitempty"`
	UntaxedAmount               *Float     `xmlrpc:"untaxed_amount,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrExpenses represents array of hr.expense model.
type HrExpenses []HrExpense

// HrExpenseModel is the odoo model name.
const HrExpenseModel = "hr.expense"

// Many2One convert HrExpense to *Many2One.
func (he *HrExpense) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrExpense creates a new hr.expense model and returns its id.
func (c *Client) CreateHrExpense(he *HrExpense) (int64, error) {
	ids, err := c.CreateHrExpenses([]*HrExpense{he})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpense creates a new hr.expense model and returns its id.
func (c *Client) CreateHrExpenses(hes []*HrExpense) ([]int64, error) {
	var vv []interface{}
	for _, v := range hes {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseModel, vv, nil)
}

// UpdateHrExpense updates an existing hr.expense record.
func (c *Client) UpdateHrExpense(he *HrExpense) error {
	return c.UpdateHrExpenses([]int64{he.Id.Get()}, he)
}

// UpdateHrExpenses updates existing hr.expense records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrExpenses(ids []int64, he *HrExpense) error {
	return c.Update(HrExpenseModel, ids, he, nil)
}

// DeleteHrExpense deletes an existing hr.expense record.
func (c *Client) DeleteHrExpense(id int64) error {
	return c.DeleteHrExpenses([]int64{id})
}

// DeleteHrExpenses deletes existing hr.expense records.
func (c *Client) DeleteHrExpenses(ids []int64) error {
	return c.Delete(HrExpenseModel, ids)
}

// GetHrExpense gets hr.expense existing record.
func (c *Client) GetHrExpense(id int64) (*HrExpense, error) {
	hes, err := c.GetHrExpenses([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hes)[0]), nil
}

// GetHrExpenses gets hr.expense existing records.
func (c *Client) GetHrExpenses(ids []int64) (*HrExpenses, error) {
	hes := &HrExpenses{}
	if err := c.Read(HrExpenseModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrExpense finds hr.expense record by querying it with criteria.
func (c *Client) FindHrExpense(criteria *Criteria) (*HrExpense, error) {
	hes := &HrExpenses{}
	if err := c.SearchRead(HrExpenseModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	return &((*hes)[0]), nil
}

// FindHrExpenses finds hr.expense records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenses(criteria *Criteria, options *Options) (*HrExpenses, error) {
	hes := &HrExpenses{}
	if err := c.SearchRead(HrExpenseModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrExpenseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseModel, criteria, options)
}

// FindHrExpenseId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
