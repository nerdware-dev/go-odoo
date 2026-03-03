package odoo

// HrExpenseSheet represents hr.expense.sheet model.
type HrExpenseSheet struct {
	AccountMoveIds                 *Relation  `xmlrpc:"account_move_ids,omitempty" json:"account_move_ids,omitempty"`
	AccountingDate                 *Time      `xmlrpc:"accounting_date,omitempty" json:"accounting_date,omitempty"`
	ActivityCalendarEventId        *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration    *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon          *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon               *String    `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AmountResidual                 *Float     `xmlrpc:"amount_residual,omitempty" json:"amount_residual,omitempty"`
	ApprovalDate                   *Time      `xmlrpc:"approval_date,omitempty" json:"approval_date,omitempty"`
	ApprovalState                  *Selection `xmlrpc:"approval_state,omitempty" json:"approval_state,omitempty"`
	AttachmentIds                  *Relation  `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	CanApprove                     *Bool      `xmlrpc:"can_approve,omitempty" json:"can_approve,omitempty"`
	CanReset                       *Bool      `xmlrpc:"can_reset,omitempty" json:"can_reset,omitempty"`
	CannotApproveReason            *String    `xmlrpc:"cannot_approve_reason,omitempty" json:"cannot_approve_reason,omitempty"`
	CompanyCurrencyId              *Many2One  `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	DepartmentId                   *Many2One  `xmlrpc:"department_id,omitempty" json:"department_id,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmployeeId                     *Many2One  `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	EmployeeJournalId              *Many2One  `xmlrpc:"employee_journal_id,omitempty" json:"employee_journal_id,omitempty"`
	ExpenseLineIds                 *Relation  `xmlrpc:"expense_line_ids,omitempty" json:"expense_line_ids,omitempty"`
	HasMessage                     *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsEditable                     *Bool      `xmlrpc:"is_editable,omitempty" json:"is_editable,omitempty"`
	IsMultipleCurrency             *Bool      `xmlrpc:"is_multiple_currency,omitempty" json:"is_multiple_currency,omitempty"`
	JournalId                      *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	MessageAttachmentCount         *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError                *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter         *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError             *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentId        *Many2One  `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MyActivityDateDeadline         *Time      `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                           *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NbAccountMove                  *Int       `xmlrpc:"nb_account_move,omitempty" json:"nb_account_move,omitempty"`
	NbExpense                      *Int       `xmlrpc:"nb_expense,omitempty" json:"nb_expense,omitempty"`
	PaymentMethodLineId            *Many2One  `xmlrpc:"payment_method_line_id,omitempty" json:"payment_method_line_id,omitempty"`
	PaymentMode                    *Selection `xmlrpc:"payment_mode,omitempty" json:"payment_mode,omitempty"`
	PaymentState                   *Selection `xmlrpc:"payment_state,omitempty" json:"payment_state,omitempty"`
	ProductIds                     *Relation  `xmlrpc:"product_ids,omitempty" json:"product_ids,omitempty"`
	RatingIds                      *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	SaleOrderCount                 *Int       `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SelectablePaymentMethodLineIds *Relation  `xmlrpc:"selectable_payment_method_line_ids,omitempty" json:"selectable_payment_method_line_ids,omitempty"`
	State                          *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	TotalAmount                    *Float     `xmlrpc:"total_amount,omitempty" json:"total_amount,omitempty"`
	TotalTaxAmount                 *Float     `xmlrpc:"total_tax_amount,omitempty" json:"total_tax_amount,omitempty"`
	UntaxedAmount                  *Float     `xmlrpc:"untaxed_amount,omitempty" json:"untaxed_amount,omitempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// HrExpenseSheets represents array of hr.expense.sheet model.
type HrExpenseSheets []HrExpenseSheet

// HrExpenseSheetModel is the odoo model name.
const HrExpenseSheetModel = "hr.expense.sheet"

// Many2One convert HrExpenseSheet to *Many2One.
func (hes *HrExpenseSheet) Many2One() *Many2One {
	return NewMany2One(hes.Id.Get(), "")
}

// CreateHrExpenseSheet creates a new hr.expense.sheet model and returns its id.
func (c *Client) CreateHrExpenseSheet(hes *HrExpenseSheet) (int64, error) {
	ids, err := c.CreateHrExpenseSheets([]*HrExpenseSheet{hes})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSheet creates a new hr.expense.sheet model and returns its id.
func (c *Client) CreateHrExpenseSheets(hess []*HrExpenseSheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range hess {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSheetModel, vv, nil)
}

// UpdateHrExpenseSheet updates an existing hr.expense.sheet record.
func (c *Client) UpdateHrExpenseSheet(hes *HrExpenseSheet) error {
	return c.UpdateHrExpenseSheets([]int64{hes.Id.Get()}, hes)
}

// UpdateHrExpenseSheets updates existing hr.expense.sheet records.
// All records (represented by ids) will be updated by hes values.
func (c *Client) UpdateHrExpenseSheets(ids []int64, hes *HrExpenseSheet) error {
	return c.Update(HrExpenseSheetModel, ids, hes, nil)
}

// DeleteHrExpenseSheet deletes an existing hr.expense.sheet record.
func (c *Client) DeleteHrExpenseSheet(id int64) error {
	return c.DeleteHrExpenseSheets([]int64{id})
}

// DeleteHrExpenseSheets deletes existing hr.expense.sheet records.
func (c *Client) DeleteHrExpenseSheets(ids []int64) error {
	return c.Delete(HrExpenseSheetModel, ids)
}

// GetHrExpenseSheet gets hr.expense.sheet existing record.
func (c *Client) GetHrExpenseSheet(id int64) (*HrExpenseSheet, error) {
	hess, err := c.GetHrExpenseSheets([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*hess) == 0 {
		return nil, nil
	}
	return &((*hess)[0]), nil
}

// GetHrExpenseSheets gets hr.expense.sheet existing records.
func (c *Client) GetHrExpenseSheets(ids []int64) (*HrExpenseSheets, error) {
	hess := &HrExpenseSheets{}
	if err := c.Read(HrExpenseSheetModel, ids, nil, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSheet finds hr.expense.sheet record by querying it with criteria.
func (c *Client) FindHrExpenseSheet(criteria *Criteria) (*HrExpenseSheet, error) {
	hess := &HrExpenseSheets{}
	if err := c.SearchRead(HrExpenseSheetModel, criteria, NewOptions().Limit(1), hess); err != nil {
		return nil, err
	}
	if len(*hess) == 0 {
		return nil, nil
	}
	return &((*hess)[0]), nil
}

// FindHrExpenseSheets finds hr.expense.sheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheets(criteria *Criteria, options *Options) (*HrExpenseSheets, error) {
	hess := &HrExpenseSheets{}
	if err := c.SearchRead(HrExpenseSheetModel, criteria, options, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseSheetModel, criteria, options)
}

// FindHrExpenseSheetId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
