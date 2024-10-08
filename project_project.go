package odoo

// ProjectProject represents project.project model.
type ProjectProject struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                  *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omitempty"`
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	AliasContact                 *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId           *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId                 *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId           *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId          *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId                  *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	AllowBillable                *Bool      `xmlrpc:"allow_billable,omitempty"`
	AllowTimesheetTimer          *Bool      `xmlrpc:"allow_timesheet_timer,omitempty"`
	AllowTimesheets              *Bool      `xmlrpc:"allow_timesheets,omitempty"`
	AnalyticAccountId            *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	BillableType                 *Selection `xmlrpc:"billable_type,omitempty"`
	Color                        *Int       `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                         *Time      `xmlrpc:"date,omitempty"`
	DateStart                    *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	DocCount                     *Int       `xmlrpc:"doc_count,omitempty"`
	FavoriteUserIds              *Relation  `xmlrpc:"favorite_user_ids,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IsFavorite                   *Bool      `xmlrpc:"is_favorite,omitempty"`
	LabelTasks                   *String    `xmlrpc:"label_tasks,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds            *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter         *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omitempty"`
	PortalShowRating             *Bool      `xmlrpc:"portal_show_rating,omitempty"`
	PrivacyVisibility            *Selection `xmlrpc:"privacy_visibility,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingPercentageSatisfaction *Int       `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	RatingRequestDeadline        *Time      `xmlrpc:"rating_request_deadline,omitempty"`
	RatingStatus                 *Selection `xmlrpc:"rating_status,omitempty"`
	RatingStatusPeriod           *Selection `xmlrpc:"rating_status_period,omitempty"`
	ResourceCalendarId           *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SaleLineEmployeeIds          *Relation  `xmlrpc:"sale_line_employee_ids,omitempty"`
	SaleLineId                   *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId                  *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	SubtaskProjectId             *Many2One  `xmlrpc:"subtask_project_id,omitempty"`
	TaskCount                    *Int       `xmlrpc:"task_count,omitempty"`
	TaskIds                      *Relation  `xmlrpc:"task_ids,omitempty"`
	Tasks                        *Relation  `xmlrpc:"tasks,omitempty"`
	TicketCount                  *Int       `xmlrpc:"ticket_count,omitempty"`
	TicketIds                    *Relation  `xmlrpc:"ticket_ids,omitempty"`
	TypeIds                      *Relation  `xmlrpc:"type_ids,omitempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
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
