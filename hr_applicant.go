package odoo

// HrApplicant represents hr.applicant model.
type HrApplicant struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	ApplicationCount            *Int       `xmlrpc:"application_count,omitempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omitempty"`
	Availability                *Time      `xmlrpc:"availability,omitempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CategIds                    *Relation  `xmlrpc:"categ_ids,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateClosed                  *Time      `xmlrpc:"date_closed,omitempty"`
	DateLastStageUpdate         *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	DateOpen                    *Time      `xmlrpc:"date_open,omitempty"`
	DayClose                    *Float     `xmlrpc:"day_close,omitempty"`
	DayOpen                     *Float     `xmlrpc:"day_open,omitempty"`
	DelayClose                  *Float     `xmlrpc:"delay_close,omitempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omitempty"`
	EmpId                       *Many2One  `xmlrpc:"emp_id,omitempty"`
	EmployeeName                *String    `xmlrpc:"employee_name,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omitempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omitempty"`
	LastStageId                 *Many2One  `xmlrpc:"last_stage_id,omitempty"`
	LegendBlocked               *String    `xmlrpc:"legend_blocked,omitempty"`
	LegendDone                  *String    `xmlrpc:"legend_done,omitempty"`
	LegendNormal                *String    `xmlrpc:"legend_normal,omitempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omitempty"`
	MeetingCount                *Int       `xmlrpc:"meeting_count,omitempty"`
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
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerMobile               *String    `xmlrpc:"partner_mobile,omitempty"`
	PartnerName                 *String    `xmlrpc:"partner_name,omitempty"`
	PartnerPhone                *String    `xmlrpc:"partner_phone,omitempty"`
	Priority                    *Selection `xmlrpc:"priority,omitempty"`
	Probability                 *Float     `xmlrpc:"probability,omitempty"`
	SalaryExpected              *Float     `xmlrpc:"salary_expected,omitempty"`
	SalaryExpectedExtra         *String    `xmlrpc:"salary_expected_extra,omitempty"`
	SalaryProposed              *Float     `xmlrpc:"salary_proposed,omitempty"`
	SalaryProposedExtra         *String    `xmlrpc:"salary_proposed_extra,omitempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omitempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omitempty"`
	TypeId                      *Many2One  `xmlrpc:"type_id,omitempty"`
	UserEmail                   *String    `xmlrpc:"user_email,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrApplicants represents array of hr.applicant model.
type HrApplicants []HrApplicant

// HrApplicantModel is the odoo model name.
const HrApplicantModel = "hr.applicant"

// Many2One convert HrApplicant to *Many2One.
func (ha *HrApplicant) Many2One() *Many2One {
	return NewMany2One(ha.Id.Get(), "")
}

// CreateHrApplicant creates a new hr.applicant model and returns its id.
func (c *Client) CreateHrApplicant(ha *HrApplicant) (int64, error) {
	ids, err := c.CreateHrApplicants([]*HrApplicant{ha})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrApplicant creates a new hr.applicant model and returns its id.
func (c *Client) CreateHrApplicants(has []*HrApplicant) ([]int64, error) {
	var vv []interface{}
	for _, v := range has {
		vv = append(vv, v)
	}
	return c.Create(HrApplicantModel, vv, nil)
}

// UpdateHrApplicant updates an existing hr.applicant record.
func (c *Client) UpdateHrApplicant(ha *HrApplicant) error {
	return c.UpdateHrApplicants([]int64{ha.Id.Get()}, ha)
}

// UpdateHrApplicants updates existing hr.applicant records.
// All records (represented by ids) will be updated by ha values.
func (c *Client) UpdateHrApplicants(ids []int64, ha *HrApplicant) error {
	return c.Update(HrApplicantModel, ids, ha, nil)
}

// DeleteHrApplicant deletes an existing hr.applicant record.
func (c *Client) DeleteHrApplicant(id int64) error {
	return c.DeleteHrApplicants([]int64{id})
}

// DeleteHrApplicants deletes existing hr.applicant records.
func (c *Client) DeleteHrApplicants(ids []int64) error {
	return c.Delete(HrApplicantModel, ids)
}

// GetHrApplicant gets hr.applicant existing record.
func (c *Client) GetHrApplicant(id int64) (*HrApplicant, error) {
	has, err := c.GetHrApplicants([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*has)[0]), nil
}

// GetHrApplicants gets hr.applicant existing records.
func (c *Client) GetHrApplicants(ids []int64) (*HrApplicants, error) {
	has := &HrApplicants{}
	if err := c.Read(HrApplicantModel, ids, nil, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrApplicant finds hr.applicant record by querying it with criteria.
func (c *Client) FindHrApplicant(criteria *Criteria) (*HrApplicant, error) {
	has := &HrApplicants{}
	if err := c.SearchRead(HrApplicantModel, criteria, NewOptions().Limit(1), has); err != nil {
		return nil, err
	}
	return &((*has)[0]), nil
}

// FindHrApplicants finds hr.applicant records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicants(criteria *Criteria, options *Options) (*HrApplicants, error) {
	has := &HrApplicants{}
	if err := c.SearchRead(HrApplicantModel, criteria, options, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrApplicantIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrApplicantModel, criteria, options)
}

// FindHrApplicantId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
