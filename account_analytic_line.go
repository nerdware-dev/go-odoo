package odoo

// AccountAnalyticLine represents account.analytic.line model.
type AccountAnalyticLine struct {
	AccountId                *Many2One  `xmlrpc:"account_id,omitempty" json:"account_id,omitempty"`
	AllowBillable            *Bool      `xmlrpc:"allow_billable,omitempty" json:"allow_billable,omitempty"`
	AllowedSoLineIds         *Relation  `xmlrpc:"allowed_so_line_ids,omitempty" json:"allowed_so_line_ids,omitempty"`
	Amount                   *Float     `xmlrpc:"amount,omitempty" json:"amount,omitempty"`
	AutoAccountId            *Many2One  `xmlrpc:"auto_account_id,omitempty" json:"auto_account_id,omitempty"`
	Category                 *Selection `xmlrpc:"category,omitempty" json:"category,omitempty"`
	Code                     *String    `xmlrpc:"code,omitempty" json:"code,omitempty"`
	CommercialPartnerId      *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty" json:"department_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisplayTimer             *Bool      `xmlrpc:"display_timer,omitempty" json:"display_timer,omitempty"`
	DisplayTimerPause        *Bool      `xmlrpc:"display_timer_pause,omitempty" json:"display_timer_pause,omitempty"`
	DisplayTimerResume       *Bool      `xmlrpc:"display_timer_resume,omitempty" json:"display_timer_resume,omitempty"`
	DisplayTimerStartPrimary *Bool      `xmlrpc:"display_timer_start_primary,omitempty" json:"display_timer_start_primary,omitempty"`
	DisplayTimerStop         *Bool      `xmlrpc:"display_timer_stop,omitempty" json:"display_timer_stop,omitempty"`
	DurationUnitAmount       *Float     `xmlrpc:"duration_unit_amount,omitempty" json:"duration_unit_amount,omitempty"`
	EmployeeId               *Many2One  `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	EncodingUomId            *Many2One  `xmlrpc:"encoding_uom_id,omitempty" json:"encoding_uom_id,omitempty"`
	GeneralAccountId         *Many2One  `xmlrpc:"general_account_id,omitempty" json:"general_account_id,omitempty"`
	GlobalLeaveId            *Many2One  `xmlrpc:"global_leave_id,omitempty" json:"global_leave_id,omitempty"`
	HolidayId                *Many2One  `xmlrpc:"holiday_id,omitempty" json:"holiday_id,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsSoLineEdited           *Bool      `xmlrpc:"is_so_line_edited,omitempty" json:"is_so_line_edited,omitempty"`
	IsTimerRunning           *Bool      `xmlrpc:"is_timer_running,omitempty" json:"is_timer_running,omitempty"`
	IsTimesheet              *Bool      `xmlrpc:"is_timesheet,omitempty" json:"is_timesheet,omitempty"`
	JobTitle                 *String    `xmlrpc:"job_title,omitempty" json:"job_title,omitempty"`
	JournalId                *Many2One  `xmlrpc:"journal_id,omitempty" json:"journal_id,omitempty"`
	L10NDin5008DocumentTitle *String    `xmlrpc:"l10n_din5008_document_title,omitempty" json:"l10n_din5008_document_title,omitempty"`
	L10NDin5008TemplateData  *String    `xmlrpc:"l10n_din5008_template_data,omitempty" json:"l10n_din5008_template_data,omitempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omitempty" json:"manager_id,omitempty"`
	MoveLineId               *Many2One  `xmlrpc:"move_line_id,omitempty" json:"move_line_id,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OrderId                  *Many2One  `xmlrpc:"order_id,omitempty" json:"order_id,omitempty"`
	ParentTaskId             *Many2One  `xmlrpc:"parent_task_id,omitempty" json:"parent_task_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductUomCategoryId     *Many2One  `xmlrpc:"product_uom_category_id,omitempty" json:"product_uom_category_id,omitempty"`
	ProductUomId             *Many2One  `xmlrpc:"product_uom_id,omitempty" json:"product_uom_id,omitempty"`
	ProjectId                *Many2One  `xmlrpc:"project_id,omitempty" json:"project_id,omitempty"`
	ReadonlyTimesheet        *Bool      `xmlrpc:"readonly_timesheet,omitempty" json:"readonly_timesheet,omitempty"`
	Ref                      *String    `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	SoLine                   *Many2One  `xmlrpc:"so_line,omitempty" json:"so_line,omitempty"`
	TaskId                   *Many2One  `xmlrpc:"task_id,omitempty" json:"task_id,omitempty"`
	TimerPause               *Time      `xmlrpc:"timer_pause,omitempty" json:"timer_pause,omitempty"`
	TimerStart               *Time      `xmlrpc:"timer_start,omitempty" json:"timer_start,omitempty"`
	TimesheetInvoiceId       *Many2One  `xmlrpc:"timesheet_invoice_id,omitempty" json:"timesheet_invoice_id,omitempty"`
	TimesheetInvoiceType     *Selection `xmlrpc:"timesheet_invoice_type,omitempty" json:"timesheet_invoice_type,omitempty"`
	UnitAmount               *Float     `xmlrpc:"unit_amount,omitempty" json:"unit_amount,omitempty"`
	UnitAmountValidate       *Float     `xmlrpc:"unit_amount_validate,omitempty" json:"unit_amount_validate,omitempty"`
	UserCanValidate          *Bool      `xmlrpc:"user_can_validate,omitempty" json:"user_can_validate,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	UserTimerId              *Relation  `xmlrpc:"user_timer_id,omitempty" json:"user_timer_id,omitempty"`
	Validated                *Bool      `xmlrpc:"validated,omitempty" json:"validated,omitempty"`
	ValidatedStatus          *Selection `xmlrpc:"validated_status,omitempty" json:"validated_status,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
	XPlan2Id                 *Many2One  `xmlrpc:"x_plan2_id,omitempty" json:"x_plan2_id,omitempty"`
}

// AccountAnalyticLines represents array of account.analytic.line model.
type AccountAnalyticLines []AccountAnalyticLine

// AccountAnalyticLineModel is the odoo model name.
const AccountAnalyticLineModel = "account.analytic.line"

// Many2One convert AccountAnalyticLine to *Many2One.
func (aal *AccountAnalyticLine) Many2One() *Many2One {
	return NewMany2One(aal.Id.Get(), "")
}

// CreateAccountAnalyticLine creates a new account.analytic.line model and returns its id.
func (c *Client) CreateAccountAnalyticLine(aal *AccountAnalyticLine) (int64, error) {
	ids, err := c.CreateAccountAnalyticLines([]*AccountAnalyticLine{aal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticLine creates a new account.analytic.line model and returns its id.
func (c *Client) CreateAccountAnalyticLines(aals []*AccountAnalyticLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range aals {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticLineModel, vv, nil)
}

// UpdateAccountAnalyticLine updates an existing account.analytic.line record.
func (c *Client) UpdateAccountAnalyticLine(aal *AccountAnalyticLine) error {
	return c.UpdateAccountAnalyticLines([]int64{aal.Id.Get()}, aal)
}

// UpdateAccountAnalyticLines updates existing account.analytic.line records.
// All records (represented by ids) will be updated by aal values.
func (c *Client) UpdateAccountAnalyticLines(ids []int64, aal *AccountAnalyticLine) error {
	return c.Update(AccountAnalyticLineModel, ids, aal, nil)
}

// DeleteAccountAnalyticLine deletes an existing account.analytic.line record.
func (c *Client) DeleteAccountAnalyticLine(id int64) error {
	return c.DeleteAccountAnalyticLines([]int64{id})
}

// DeleteAccountAnalyticLines deletes existing account.analytic.line records.
func (c *Client) DeleteAccountAnalyticLines(ids []int64) error {
	return c.Delete(AccountAnalyticLineModel, ids)
}

// GetAccountAnalyticLine gets account.analytic.line existing record.
func (c *Client) GetAccountAnalyticLine(id int64) (*AccountAnalyticLine, error) {
	aals, err := c.GetAccountAnalyticLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*aals) == 0 {
		return nil, nil
	}
	return &((*aals)[0]), nil
}

// GetAccountAnalyticLines gets account.analytic.line existing records.
func (c *Client) GetAccountAnalyticLines(ids []int64) (*AccountAnalyticLines, error) {
	aals := &AccountAnalyticLines{}
	if err := c.Read(AccountAnalyticLineModel, ids, nil, aals); err != nil {
		return nil, err
	}
	return aals, nil
}

// FindAccountAnalyticLine finds account.analytic.line record by querying it with criteria.
func (c *Client) FindAccountAnalyticLine(criteria *Criteria) (*AccountAnalyticLine, error) {
	aals := &AccountAnalyticLines{}
	if err := c.SearchRead(AccountAnalyticLineModel, criteria, NewOptions().Limit(1), aals); err != nil {
		return nil, err
	}
	if len(*aals) == 0 {
		return nil, nil
	}
	return &((*aals)[0]), nil
}

// FindAccountAnalyticLines finds account.analytic.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLines(criteria *Criteria, options *Options) (*AccountAnalyticLines, error) {
	aals := &AccountAnalyticLines{}
	if err := c.SearchRead(AccountAnalyticLineModel, criteria, options, aals); err != nil {
		return nil, err
	}
	return aals, nil
}

// FindAccountAnalyticLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticLineModel, criteria, options)
}

// FindAccountAnalyticLineId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
