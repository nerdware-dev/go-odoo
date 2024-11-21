package odoo

// HrJob represents hr.job model.
type HrJob struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	AddressId                *Many2One  `xmlrpc:"address_id,omitempty" json:"address_id,omitempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omitempty" json:"alias_contact,omitempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omitempty" json:"alias_defaults,omitempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omitempty" json:"alias_domain,omitempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omitempty" json:"alias_force_thread_id,omitempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omitempty" json:"alias_id,omitempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omitempty" json:"alias_model_id,omitempty"`
	AliasName                *String    `xmlrpc:"alias_name,omitempty" json:"alias_name,omitempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omitempty" json:"alias_parent_model_id,omitempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omitempty" json:"alias_parent_thread_id,omitempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omitempty" json:"alias_user_id,omitempty"`
	ApplicationCount         *Int       `xmlrpc:"application_count,omitempty" json:"application_count,omitempty"`
	ApplicationIds           *Relation  `xmlrpc:"application_ids,omitempty" json:"application_ids,omitempty"`
	CanPublish               *Bool      `xmlrpc:"can_publish,omitempty" json:"can_publish,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty" json:"department_id,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentIds              *Relation  `xmlrpc:"document_ids,omitempty" json:"document_ids,omitempty"`
	DocumentsCount           *Int       `xmlrpc:"documents_count,omitempty" json:"documents_count,omitempty"`
	EmployeeIds              *Relation  `xmlrpc:"employee_ids,omitempty" json:"employee_ids,omitempty"`
	ExpectedEmployees        *Int       `xmlrpc:"expected_employees,omitempty" json:"expected_employees,omitempty"`
	FavoriteUserIds          *Relation  `xmlrpc:"favorite_user_ids,omitempty" json:"favorite_user_ids,omitempty"`
	HrResponsibleId          *Many2One  `xmlrpc:"hr_responsible_id,omitempty" json:"hr_responsible_id,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsFavorite               *Bool      `xmlrpc:"is_favorite,omitempty" json:"is_favorite,omitempty"`
	IsPublished              *Bool      `xmlrpc:"is_published,omitempty" json:"is_published,omitempty"`
	IsSeoOptimized           *Bool      `xmlrpc:"is_seo_optimized,omitempty" json:"is_seo_optimized,omitempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omitempty" json:"manager_id,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty" json:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty" json:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty" json:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NewApplicationCount      *Int       `xmlrpc:"new_application_count,omitempty" json:"new_application_count,omitempty"`
	NoOfEmployee             *Int       `xmlrpc:"no_of_employee,omitempty" json:"no_of_employee,omitempty"`
	NoOfHiredEmployee        *Int       `xmlrpc:"no_of_hired_employee,omitempty" json:"no_of_hired_employee,omitempty"`
	NoOfRecruitment          *Int       `xmlrpc:"no_of_recruitment,omitempty" json:"no_of_recruitment,omitempty"`
	Requirements             *String    `xmlrpc:"requirements,omitempty" json:"requirements,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	WebsiteDescription       *String    `xmlrpc:"website_description,omitempty" json:"website_description,omitempty"`
	WebsiteId                *Many2One  `xmlrpc:"website_id,omitempty" json:"website_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WebsiteMetaDescription   *String    `xmlrpc:"website_meta_description,omitempty" json:"website_meta_description,omitempty"`
	WebsiteMetaKeywords      *String    `xmlrpc:"website_meta_keywords,omitempty" json:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg         *String    `xmlrpc:"website_meta_og_img,omitempty" json:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle         *String    `xmlrpc:"website_meta_title,omitempty" json:"website_meta_title,omitempty"`
	WebsitePublished         *Bool      `xmlrpc:"website_published,omitempty" json:"website_published,omitempty"`
	WebsiteUrl               *String    `xmlrpc:"website_url,omitempty" json:"website_url,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// HrJobs represents array of hr.job model.
type HrJobs []HrJob

// HrJobModel is the odoo model name.
const HrJobModel = "hr.job"

// Many2One convert HrJob to *Many2One.
func (hj *HrJob) Many2One() *Many2One {
	return NewMany2One(hj.Id.Get(), "")
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJob(hj *HrJob) (int64, error) {
	ids, err := c.CreateHrJobs([]*HrJob{hj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJobs(hjs []*HrJob) ([]int64, error) {
	var vv []interface{}
	for _, v := range hjs {
		vv = append(vv, v)
	}
	return c.Create(HrJobModel, vv, nil)
}

// UpdateHrJob updates an existing hr.job record.
func (c *Client) UpdateHrJob(hj *HrJob) error {
	return c.UpdateHrJobs([]int64{hj.Id.Get()}, hj)
}

// UpdateHrJobs updates existing hr.job records.
// All records (represented by ids) will be updated by hj values.
func (c *Client) UpdateHrJobs(ids []int64, hj *HrJob) error {
	return c.Update(HrJobModel, ids, hj, nil)
}

// DeleteHrJob deletes an existing hr.job record.
func (c *Client) DeleteHrJob(id int64) error {
	return c.DeleteHrJobs([]int64{id})
}

// DeleteHrJobs deletes existing hr.job records.
func (c *Client) DeleteHrJobs(ids []int64) error {
	return c.Delete(HrJobModel, ids)
}

// GetHrJob gets hr.job existing record.
func (c *Client) GetHrJob(id int64) (*HrJob, error) {
	hjs, err := c.GetHrJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hjs)[0]), nil
}

// GetHrJobs gets hr.job existing records.
func (c *Client) GetHrJobs(ids []int64) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.Read(HrJobModel, ids, nil, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJob finds hr.job record by querying it with criteria.
func (c *Client) FindHrJob(criteria *Criteria) (*HrJob, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, NewOptions().Limit(1), hjs); err != nil {
		return nil, err
	}
	return &((*hjs)[0]), nil
}

// FindHrJobs finds hr.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobs(criteria *Criteria, options *Options) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, options, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrJobModel, criteria, options)
}

// FindHrJobId finds record id by querying it with criteria.
func (c *Client) FindHrJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
