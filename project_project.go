package odoo

// ProjectProject represents project.project model.
type ProjectProject struct {
	AccessInstructionMessage     *String     `xmlrpc:"access_instruction_message,omitempty" json:"access_instruction_message,omitempty"`
	AccessToken                  *String     `xmlrpc:"access_token,omitempty" json:"access_token,omitempty"`
	AccessUrl                    *String     `xmlrpc:"access_url,omitempty" json:"access_url,omitempty"`
	AccessWarning                *String     `xmlrpc:"access_warning,omitempty" json:"access_warning,omitempty"`
	Active                       *Bool       `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId      *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline         *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary              *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon             *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AliasBouncedContent          *String     `xmlrpc:"alias_bounced_content,omitempty" json:"alias_bounced_content,omitempty"`
	AliasContact                 *Selection  `xmlrpc:"alias_contact,omitempty" json:"alias_contact,omitempty"`
	AliasDefaults                *String     `xmlrpc:"alias_defaults,omitempty" json:"alias_defaults,omitempty"`
	AliasDomain                  *String     `xmlrpc:"alias_domain,omitempty" json:"alias_domain,omitempty"`
	AliasDomainId                *Many2One   `xmlrpc:"alias_domain_id,omitempty" json:"alias_domain_id,omitempty"`
	AliasEmail                   *String     `xmlrpc:"alias_email,omitempty" json:"alias_email,omitempty"`
	AliasForceThreadId           *Int        `xmlrpc:"alias_force_thread_id,omitempty" json:"alias_force_thread_id,omitempty"`
	AliasFullName                *String     `xmlrpc:"alias_full_name,omitempty" json:"alias_full_name,omitempty"`
	AliasId                      *Many2One   `xmlrpc:"alias_id,omitempty" json:"alias_id,omitempty"`
	AliasIncomingLocal           *Bool       `xmlrpc:"alias_incoming_local,omitempty" json:"alias_incoming_local,omitempty"`
	AliasModelId                 *Many2One   `xmlrpc:"alias_model_id,omitempty" json:"alias_model_id,omitempty"`
	AliasName                    *String     `xmlrpc:"alias_name,omitempty" json:"alias_name,omitempty"`
	AliasParentModelId           *Many2One   `xmlrpc:"alias_parent_model_id,omitempty" json:"alias_parent_model_id,omitempty"`
	AliasParentThreadId          *Int        `xmlrpc:"alias_parent_thread_id,omitempty" json:"alias_parent_thread_id,omitempty"`
	AliasStatus                  *Selection  `xmlrpc:"alias_status,omitempty" json:"alias_status,omitempty"`
	AllocatedHours               *Float      `xmlrpc:"allocated_hours,omitempty" json:"allocated_hours,omitempty"`
	AllowBillable                *Bool       `xmlrpc:"allow_billable,omitempty" json:"allow_billable,omitempty"`
	AllowMilestones              *Bool       `xmlrpc:"allow_milestones,omitempty" json:"allow_milestones,omitempty"`
	AllowRating                  *Bool       `xmlrpc:"allow_rating,omitempty" json:"allow_rating,omitempty"`
	AllowTaskDependencies        *Bool       `xmlrpc:"allow_task_dependencies,omitempty" json:"allow_task_dependencies,omitempty"`
	AllowTimesheets              *Bool       `xmlrpc:"allow_timesheets,omitempty" json:"allow_timesheets,omitempty"`
	AnalyticAccountBalance       *Float      `xmlrpc:"analytic_account_balance,omitempty" json:"analytic_account_balance,omitempty"`
	AnalyticAccountId            *Many2One   `xmlrpc:"analytic_account_id,omitempty" json:"analytic_account_id,omitempty"`
	AssetsCount                  *Int        `xmlrpc:"assets_count,omitempty" json:"assets_count,omitempty"`
	BillablePercentage           *Int        `xmlrpc:"billable_percentage,omitempty" json:"billable_percentage,omitempty"`
	BillingType                  *Selection  `xmlrpc:"billing_type,omitempty" json:"billing_type,omitempty"`
	Budget                       *Int        `xmlrpc:"budget,omitempty" json:"budget,omitempty"`
	ClosedTaskCount              *Int        `xmlrpc:"closed_task_count,omitempty" json:"closed_task_count,omitempty"`
	CollaboratorCount            *Int        `xmlrpc:"collaborator_count,omitempty" json:"collaborator_count,omitempty"`
	CollaboratorIds              *Relation   `xmlrpc:"collaborator_ids,omitempty" json:"collaborator_ids,omitempty"`
	Color                        *Int        `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CompanyId                    *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                   *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                    *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                   *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	Date                         *Time       `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DateStart                    *Time       `xmlrpc:"date_start,omitempty" json:"date_start,omitempty"`
	Description                  *String     `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                  *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisplaySalesStatButtons      *Bool       `xmlrpc:"display_sales_stat_buttons,omitempty" json:"display_sales_stat_buttons,omitempty"`
	DocCount                     *Int        `xmlrpc:"doc_count,omitempty" json:"doc_count,omitempty"`
	DocumentCount                *Int        `xmlrpc:"document_count,omitempty" json:"document_count,omitempty"`
	DocumentsFolderId            *Many2One   `xmlrpc:"documents_folder_id,omitempty" json:"documents_folder_id,omitempty"`
	DocumentsTagIds              *Relation   `xmlrpc:"documents_tag_ids,omitempty" json:"documents_tag_ids,omitempty"`
	EncodeUomInDays              *Bool       `xmlrpc:"encode_uom_in_days,omitempty" json:"encode_uom_in_days,omitempty"`
	ExpensesCount                *Int        `xmlrpc:"expenses_count,omitempty" json:"expenses_count,omitempty"`
	FavoriteUserIds              *Relation   `xmlrpc:"favorite_user_ids,omitempty" json:"favorite_user_ids,omitempty"`
	HasAnySoToInvoice            *Bool       `xmlrpc:"has_any_so_to_invoice,omitempty" json:"has_any_so_to_invoice,omitempty"`
	HasAnySoWithNothingToInvoice *Bool       `xmlrpc:"has_any_so_with_nothing_to_invoice,omitempty" json:"has_any_so_with_nothing_to_invoice,omitempty"`
	HasMessage                   *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                           *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InvoiceCount                 *Int        `xmlrpc:"invoice_count,omitempty" json:"invoice_count,omitempty"`
	IsFavorite                   *Bool       `xmlrpc:"is_favorite,omitempty" json:"is_favorite,omitempty"`
	IsInternalProject            *Bool       `xmlrpc:"is_internal_project,omitempty" json:"is_internal_project,omitempty"`
	IsMilestoneExceeded          *Bool       `xmlrpc:"is_milestone_exceeded,omitempty" json:"is_milestone_exceeded,omitempty"`
	IsProjectOvertime            *Bool       `xmlrpc:"is_project_overtime,omitempty" json:"is_project_overtime,omitempty"`
	LabelTasks                   *String     `xmlrpc:"label_tasks,omitempty" json:"label_tasks,omitempty"`
	LastUpdateColor              *Int        `xmlrpc:"last_update_color,omitempty" json:"last_update_color,omitempty"`
	LastUpdateId                 *Many2One   `xmlrpc:"last_update_id,omitempty" json:"last_update_id,omitempty"`
	LastUpdateStatus             *Selection  `xmlrpc:"last_update_status,omitempty" json:"last_update_status,omitempty"`
	MessageAttachmentCount       *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError              *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower            *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MilestoneCount               *Int        `xmlrpc:"milestone_count,omitempty" json:"milestone_count,omitempty"`
	MilestoneCountReached        *Int        `xmlrpc:"milestone_count_reached,omitempty" json:"milestone_count_reached,omitempty"`
	MilestoneIds                 *Relation   `xmlrpc:"milestone_ids,omitempty" json:"milestone_ids,omitempty"`
	MyActivityDateDeadline       *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                         *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OpenTaskCount                *Int        `xmlrpc:"open_task_count,omitempty" json:"open_task_count,omitempty"`
	PartnerId                    *Many2One   `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PricingType                  *Selection  `xmlrpc:"pricing_type,omitempty" json:"pricing_type,omitempty"`
	PrivacyVisibility            *Selection  `xmlrpc:"privacy_visibility,omitempty" json:"privacy_visibility,omitempty"`
	PrivacyVisibilityWarning     *String     `xmlrpc:"privacy_visibility_warning,omitempty" json:"privacy_visibility_warning,omitempty"`
	PurchaseOrdersCount          *Int        `xmlrpc:"purchase_orders_count,omitempty" json:"purchase_orders_count,omitempty"`
	RatingActive                 *Bool       `xmlrpc:"rating_active,omitempty" json:"rating_active,omitempty"`
	RatingAvg                    *Float      `xmlrpc:"rating_avg,omitempty" json:"rating_avg,omitempty"`
	RatingAvgPercentage          *Float      `xmlrpc:"rating_avg_percentage,omitempty" json:"rating_avg_percentage,omitempty"`
	RatingCount                  *Int        `xmlrpc:"rating_count,omitempty" json:"rating_count,omitempty"`
	RatingIds                    *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RatingPercentageSatisfaction *Int        `xmlrpc:"rating_percentage_satisfaction,omitempty" json:"rating_percentage_satisfaction,omitempty"`
	RatingRequestDeadline        *Time       `xmlrpc:"rating_request_deadline,omitempty" json:"rating_request_deadline,omitempty"`
	RatingStatus                 *Selection  `xmlrpc:"rating_status,omitempty" json:"rating_status,omitempty"`
	RatingStatusPeriod           *Selection  `xmlrpc:"rating_status_period,omitempty" json:"rating_status_period,omitempty"`
	RemainingHours               *Float      `xmlrpc:"remaining_hours,omitempty" json:"remaining_hours,omitempty"`
	ResourceCalendarId           *Many2One   `xmlrpc:"resource_calendar_id,omitempty" json:"resource_calendar_id,omitempty"`
	SaleLineEmployeeIds          *Relation   `xmlrpc:"sale_line_employee_ids,omitempty" json:"sale_line_employee_ids,omitempty"`
	SaleLineId                   *Many2One   `xmlrpc:"sale_line_id,omitempty" json:"sale_line_id,omitempty"`
	SaleOrderCount               *Int        `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SaleOrderId                  *Many2One   `xmlrpc:"sale_order_id,omitempty" json:"sale_order_id,omitempty"`
	SaleOrderLineCount           *Int        `xmlrpc:"sale_order_line_count,omitempty" json:"sale_order_line_count,omitempty"`
	Sequence                     *Int        `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	SharedDocumentCount          *Int        `xmlrpc:"shared_document_count,omitempty" json:"shared_document_count,omitempty"`
	SharedDocumentIds            *Relation   `xmlrpc:"shared_document_ids,omitempty" json:"shared_document_ids,omitempty"`
	StageId                      *Many2One   `xmlrpc:"stage_id,omitempty" json:"stage_id,omitempty"`
	TagIds                       *Relation   `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TaskCount                    *Int        `xmlrpc:"task_count,omitempty" json:"task_count,omitempty"`
	TaskIds                      *Relation   `xmlrpc:"task_ids,omitempty" json:"task_ids,omitempty"`
	TaskPropertiesDefinition     interface{} `xmlrpc:"task_properties_definition,omitempty" json:"task_properties_definition,omitempty"`
	Tasks                        *Relation   `xmlrpc:"tasks,omitempty" json:"tasks,omitempty"`
	TimesheetEncodeUomId         *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty" json:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                 *Relation   `xmlrpc:"timesheet_ids,omitempty" json:"timesheet_ids,omitempty"`
	TimesheetProductId           *Many2One   `xmlrpc:"timesheet_product_id,omitempty" json:"timesheet_product_id,omitempty"`
	TotalBudgetProgress          *Float      `xmlrpc:"total_budget_progress,omitempty" json:"total_budget_progress,omitempty"`
	TotalPlannedAmount           *Float      `xmlrpc:"total_planned_amount,omitempty" json:"total_planned_amount,omitempty"`
	TotalPracticalAmount         *Float      `xmlrpc:"total_practical_amount,omitempty" json:"total_practical_amount,omitempty"`
	TotalTimesheetTime           *Int        `xmlrpc:"total_timesheet_time,omitempty" json:"total_timesheet_time,omitempty"`
	TypeIds                      *Relation   `xmlrpc:"type_ids,omitempty" json:"type_ids,omitempty"`
	UpdateIds                    *Relation   `xmlrpc:"update_ids,omitempty" json:"update_ids,omitempty"`
	UseDocuments                 *Bool       `xmlrpc:"use_documents,omitempty" json:"use_documents,omitempty"`
	UserId                       *Many2One   `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	VendorBillCount              *Int        `xmlrpc:"vendor_bill_count,omitempty" json:"vendor_bill_count,omitempty"`
	WarningEmployeeRate          *Bool       `xmlrpc:"warning_employee_rate,omitempty" json:"warning_employee_rate,omitempty"`
	WebsiteMessageIds            *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                    *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                     *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// ProjectProjects represents array of project.project model.
type ProjectProjects []ProjectProject

// ProjectProjectModel is the odoo model name.
const ProjectProjectModel = "project.project"

// Many2One convert ProjectProject to *Many2One.
func (pp *ProjectProject) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProject(pp *ProjectProject) (int64, error) {
	ids, err := c.CreateProjectProjects([]*ProjectProject{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProjects(pps []*ProjectProject) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectModel, vv, nil)
}

// UpdateProjectProject updates an existing project.project record.
func (c *Client) UpdateProjectProject(pp *ProjectProject) error {
	return c.UpdateProjectProjects([]int64{pp.Id.Get()}, pp)
}

// UpdateProjectProjects updates existing project.project records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProjectProjects(ids []int64, pp *ProjectProject) error {
	return c.Update(ProjectProjectModel, ids, pp, nil)
}

// DeleteProjectProject deletes an existing project.project record.
func (c *Client) DeleteProjectProject(id int64) error {
	return c.DeleteProjectProjects([]int64{id})
}

// DeleteProjectProjects deletes existing project.project records.
func (c *Client) DeleteProjectProjects(ids []int64) error {
	return c.Delete(ProjectProjectModel, ids)
}

// GetProjectProject gets project.project existing record.
func (c *Client) GetProjectProject(id int64) (*ProjectProject, error) {
	pps, err := c.GetProjectProjects([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetProjectProjects gets project.project existing records.
func (c *Client) GetProjectProjects(ids []int64) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.Read(ProjectProjectModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProject finds project.project record by querying it with criteria.
func (c *Client) FindProjectProject(criteria *Criteria) (*ProjectProject, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindProjectProjects finds project.project records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjects(criteria *Criteria, options *Options) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProjectIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectProjectModel, criteria, options)
}

// FindProjectProjectId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
