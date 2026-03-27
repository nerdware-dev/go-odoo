package odoo

// HrExpense represents hr.expense model.
type HrExpense struct {
	AccountId                     *Many2One   `xmlrpc:"account_id,omitempty" json:"account_id,omitempty"`
	AccountingDate                *Time       `xmlrpc:"accounting_date,omitempty" json:"accounting_date,omitempty"`
	ActivityCalendarEventId       *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline          *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration   *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon         *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                   *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                 *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary               *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon              *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AnalyticDistribution          interface{} `xmlrpc:"analytic_distribution,omitempty" json:"analytic_distribution,omitempty"`
	AnalyticDistributionSearch    interface{} `xmlrpc:"analytic_distribution_search,omitempty" json:"analytic_distribution_search,omitempty"`
	AnalyticPrecision             *Int        `xmlrpc:"analytic_precision,omitempty" json:"analytic_precision,omitempty"`
	ApprovedBy                    *Many2One   `xmlrpc:"approved_by,omitempty" json:"approved_by,omitempty"`
	ApprovedOn                    *Time       `xmlrpc:"approved_on,omitempty" json:"approved_on,omitempty"`
	AttachmentIds                 *Relation   `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	CanBeReinvoiced               *Bool       `xmlrpc:"can_be_reinvoiced,omitempty" json:"can_be_reinvoiced,omitempty"`
	CompanyCurrencyId             *Many2One   `xmlrpc:"company_currency_id,omitempty" json:"company_currency_id,omitempty"`
	CompanyId                     *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                    *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                     *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                    *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	CurrencyRate                  *Float      `xmlrpc:"currency_rate,omitempty" json:"currency_rate,omitempty"`
	Date                          *Time       `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Description                   *String     `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                   *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentCount                 *Int        `xmlrpc:"document_count,omitempty" json:"document_count,omitempty"`
	DuplicateExpenseIds           *Relation   `xmlrpc:"duplicate_expense_ids,omitempty" json:"duplicate_expense_ids,omitempty"`
	EmployeeId                    *Many2One   `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	ExtractCanShowSendButton      *Bool       `xmlrpc:"extract_can_show_send_button,omitempty" json:"extract_can_show_send_button,omitempty"`
	ExtractDocumentUuid           *String     `xmlrpc:"extract_document_uuid,omitempty" json:"extract_document_uuid,omitempty"`
	ExtractErrorMessage           *String     `xmlrpc:"extract_error_message,omitempty" json:"extract_error_message,omitempty"`
	ExtractState                  *Selection  `xmlrpc:"extract_state,omitempty" json:"extract_state,omitempty"`
	ExtractStateProcessed         *Bool       `xmlrpc:"extract_state_processed,omitempty" json:"extract_state_processed,omitempty"`
	ExtractStatus                 *String     `xmlrpc:"extract_status,omitempty" json:"extract_status,omitempty"`
	HasMessage                    *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                            *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsEditable                    *Bool       `xmlrpc:"is_editable,omitempty" json:"is_editable,omitempty"`
	IsInExtractableState          *Bool       `xmlrpc:"is_in_extractable_state,omitempty" json:"is_in_extractable_state,omitempty"`
	IsMultipleCurrency            *Bool       `xmlrpc:"is_multiple_currency,omitempty" json:"is_multiple_currency,omitempty"`
	LabelCurrencyRate             *String     `xmlrpc:"label_currency_rate,omitempty" json:"label_currency_rate,omitempty"`
	MessageAttachmentCount        *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds            *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError               *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower             *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentChecksum *String     `xmlrpc:"message_main_attachment_checksum,omitempty" json:"message_main_attachment_checksum,omitempty"`
	MessageMainAttachmentId       *Many2One   `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction             *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MyActivityDateDeadline        *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                          *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NbAttachment                  *Int        `xmlrpc:"nb_attachment,omitempty" json:"nb_attachment,omitempty"`
	PaymentMode                   *Selection  `xmlrpc:"payment_mode,omitempty" json:"payment_mode,omitempty"`
	PredictedCategory             *Selection  `xmlrpc:"predicted_category,omitempty" json:"predicted_category,omitempty"`
	PriceUnit                     *Float      `xmlrpc:"price_unit,omitempty" json:"price_unit,omitempty"`
	ProductDescription            *String     `xmlrpc:"product_description,omitempty" json:"product_description,omitempty"`
	ProductHasCost                *Bool       `xmlrpc:"product_has_cost,omitempty" json:"product_has_cost,omitempty"`
	ProductHasTax                 *Bool       `xmlrpc:"product_has_tax,omitempty" json:"product_has_tax,omitempty"`
	ProductId                     *Many2One   `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductUomCategoryId          *Many2One   `xmlrpc:"product_uom_category_id,omitempty" json:"product_uom_category_id,omitempty"`
	ProductUomId                  *Many2One   `xmlrpc:"product_uom_id,omitempty" json:"product_uom_id,omitempty"`
	Quantity                      *Float      `xmlrpc:"quantity,omitempty" json:"quantity,omitempty"`
	RatingIds                     *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	SaleOrderId                   *Many2One   `xmlrpc:"sale_order_id,omitempty" json:"sale_order_id,omitempty"`
	Sample                        *Bool       `xmlrpc:"sample,omitempty" json:"sample,omitempty"`
	SheetId                       *Many2One   `xmlrpc:"sheet_id,omitempty" json:"sheet_id,omitempty"`
	State                         *Selection  `xmlrpc:"state,omitempty" json:"state,omitempty"`
	TaxAmount                     *Float      `xmlrpc:"tax_amount,omitempty" json:"tax_amount,omitempty"`
	TaxAmountCurrency             *Float      `xmlrpc:"tax_amount_currency,omitempty" json:"tax_amount_currency,omitempty"`
	TaxIds                        *Relation   `xmlrpc:"tax_ids,omitempty" json:"tax_ids,omitempty"`
	TotalAmount                   *Float      `xmlrpc:"total_amount,omitempty" json:"total_amount,omitempty"`
	TotalAmountCurrency           *Float      `xmlrpc:"total_amount_currency,omitempty" json:"total_amount_currency,omitempty"`
	UntaxedAmountCurrency         *Float      `xmlrpc:"untaxed_amount_currency,omitempty" json:"untaxed_amount_currency,omitempty"`
	WebsiteMessageIds             *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                     *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                      *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
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
	if len(*hes) == 0 {
		return nil, nil
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
	if len(*hes) == 0 {
		return nil, nil
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
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
