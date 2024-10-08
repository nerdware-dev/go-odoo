package odoo

// HelpdeskTeam represents helpdesk.team model.
type HelpdeskTeam struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
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
	AllowPortalTicketClosing     *Bool      `xmlrpc:"allow_portal_ticket_closing,omitempty"`
	AssignMethod                 *Selection `xmlrpc:"assign_method,omitempty"`
	CanPublish                   *Bool      `xmlrpc:"can_publish,omitempty"`
	Color                        *Int       `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                  *String    `xmlrpc:"description,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	FeatureFormUrl               *String    `xmlrpc:"feature_form_url,omitempty"`
	ForumId                      *Many2One  `xmlrpc:"forum_id,omitempty"`
	ForumUrl                     *String    `xmlrpc:"forum_url,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IsPublished                  *Bool      `xmlrpc:"is_published,omitempty"`
	MemberIds                    *Relation  `xmlrpc:"member_ids,omitempty"`
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
	PortalRatingUrl              *String    `xmlrpc:"portal_rating_url,omitempty"`
	PortalShowRating             *Bool      `xmlrpc:"portal_show_rating,omitempty"`
	ProjectId                    *Many2One  `xmlrpc:"project_id,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingPercentageSatisfaction *Int       `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	ResourceCalendarId           *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	StageIds                     *Relation  `xmlrpc:"stage_ids,omitempty"`
	TicketIds                    *Relation  `xmlrpc:"ticket_ids,omitempty"`
	UnassignedTickets            *Int       `xmlrpc:"unassigned_tickets,omitempty"`
	UpcomingSlaFailTickets       *Int       `xmlrpc:"upcoming_sla_fail_tickets,omitempty"`
	UseAlias                     *Bool      `xmlrpc:"use_alias,omitempty"`
	UseApi                       *Bool      `xmlrpc:"use_api,omitempty"`
	UseCoupons                   *Bool      `xmlrpc:"use_coupons,omitempty"`
	UseCreditNotes               *Bool      `xmlrpc:"use_credit_notes,omitempty"`
	UseHelpdeskSaleTimesheet     *Bool      `xmlrpc:"use_helpdesk_sale_timesheet,omitempty"`
	UseHelpdeskTimesheet         *Bool      `xmlrpc:"use_helpdesk_timesheet,omitempty"`
	UseProductRepairs            *Bool      `xmlrpc:"use_product_repairs,omitempty"`
	UseProductReturns            *Bool      `xmlrpc:"use_product_returns,omitempty"`
	UseRating                    *Bool      `xmlrpc:"use_rating,omitempty"`
	UseSla                       *Bool      `xmlrpc:"use_sla,omitempty"`
	UseTwitter                   *Bool      `xmlrpc:"use_twitter,omitempty"`
	UseWebsiteHelpdeskForm       *Bool      `xmlrpc:"use_website_helpdesk_form,omitempty"`
	UseWebsiteHelpdeskForum      *Bool      `xmlrpc:"use_website_helpdesk_forum,omitempty"`
	UseWebsiteHelpdeskLivechat   *Bool      `xmlrpc:"use_website_helpdesk_livechat,omitempty"`
	UseWebsiteHelpdeskSlides     *Bool      `xmlrpc:"use_website_helpdesk_slides,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsitePublished             *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                   *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HelpdeskTeams represents array of helpdesk.team model.
type HelpdeskTeams []HelpdeskTeam

// HelpdeskTeamModel is the odoo model name.
const HelpdeskTeamModel = "helpdesk.team"

// Many2One convert HelpdeskTeam to *Many2One.
func (ht *HelpdeskTeam) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeam(ht *HelpdeskTeam) (int64, error) {
	ids, err := c.CreateHelpdeskTeams([]*HelpdeskTeam{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeams(hts []*HelpdeskTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTeamModel, vv, nil)
}

// UpdateHelpdeskTeam updates an existing helpdesk.team record.
func (c *Client) UpdateHelpdeskTeam(ht *HelpdeskTeam) error {
	return c.UpdateHelpdeskTeams([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTeams updates existing helpdesk.team records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTeams(ids []int64, ht *HelpdeskTeam) error {
	return c.Update(HelpdeskTeamModel, ids, ht, nil)
}

// DeleteHelpdeskTeam deletes an existing helpdesk.team record.
func (c *Client) DeleteHelpdeskTeam(id int64) error {
	return c.DeleteHelpdeskTeams([]int64{id})
}

// DeleteHelpdeskTeams deletes existing helpdesk.team records.
func (c *Client) DeleteHelpdeskTeams(ids []int64) error {
	return c.Delete(HelpdeskTeamModel, ids)
}

// GetHelpdeskTeam gets helpdesk.team existing record.
func (c *Client) GetHelpdeskTeam(id int64) (*HelpdeskTeam, error) {
	hts, err := c.GetHelpdeskTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// GetHelpdeskTeams gets helpdesk.team existing records.
func (c *Client) GetHelpdeskTeams(ids []int64) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.Read(HelpdeskTeamModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeam finds helpdesk.team record by querying it with criteria.
func (c *Client) FindHelpdeskTeam(criteria *Criteria) (*HelpdeskTeam, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// FindHelpdeskTeams finds helpdesk.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeams(criteria *Criteria, options *Options) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskTeamModel, criteria, options)
}

// FindHelpdeskTeamId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
