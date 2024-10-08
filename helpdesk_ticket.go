package odoo

// HelpdeskTicket represents helpdesk.ticket model.
type HelpdeskTicket struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AssignDate                  *Time      `xmlrpc:"assign_date,omitempty"`
	AssignHours                 *Int       `xmlrpc:"assign_hours,omitempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omitempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CloseDate                   *Time      `xmlrpc:"close_date,omitempty"`
	CloseHours                  *Int       `xmlrpc:"close_hours,omitempty"`
	ClosedByPartner             *Bool      `xmlrpc:"closed_by_partner,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	CommercialPartnerId         *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateLastStageUpdate         *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Email                       *String    `xmlrpc:"email,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	ForumPostId                 *Many2One  `xmlrpc:"forum_post_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsClosed                    *Bool      `xmlrpc:"is_closed,omitempty"`
	IsOverdue                   *Bool      `xmlrpc:"is_overdue,omitempty"`
	IsSelfAssigned              *Bool      `xmlrpc:"is_self_assigned,omitempty"`
	IsTaskActive                *Bool      `xmlrpc:"is_task_active,omitempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omitempty"`
	KanbanStateLabel            *String    `xmlrpc:"kanban_state_label,omitempty"`
	LegendBlocked               *String    `xmlrpc:"legend_blocked,omitempty"`
	LegendDone                  *String    `xmlrpc:"legend_done,omitempty"`
	LegendNormal                *String    `xmlrpc:"legend_normal,omitempty"`
	LotId                       *Many2One  `xmlrpc:"lot_id,omitempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omitempty"`
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
	OpenHours                   *Int       `xmlrpc:"open_hours,omitempty"`
	PartnerEmail                *String    `xmlrpc:"partner_email,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerName                 *String    `xmlrpc:"partner_name,omitempty"`
	PartnerTicketCount          *Int       `xmlrpc:"partner_ticket_count,omitempty"`
	PickingIds                  *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingsCount               *Int       `xmlrpc:"pickings_count,omitempty"`
	Priority                    *Selection `xmlrpc:"priority,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omitempty"`
	RatingAvg                   *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingCount                 *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback          *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage             *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastValue             *Float     `xmlrpc:"rating_last_value,omitempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	SlaDeadline                 *Time      `xmlrpc:"sla_deadline,omitempty"`
	SlaFail                     *Bool      `xmlrpc:"sla_fail,omitempty"`
	SlaIds                      *Relation  `xmlrpc:"sla_ids,omitempty"`
	SlaReachedLate              *Bool      `xmlrpc:"sla_reached_late,omitempty"`
	SlaStatusIds                *Relation  `xmlrpc:"sla_status_ids,omitempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omitempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TaskId                      *Many2One  `xmlrpc:"task_id,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	TicketTypeId                *Many2One  `xmlrpc:"ticket_type_id,omitempty"`
	TimesheetIds                *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	Tracking                    *Selection `xmlrpc:"tracking,omitempty"`
	UseCoupons                  *Bool      `xmlrpc:"use_coupons,omitempty"`
	UseCreditNotes              *Bool      `xmlrpc:"use_credit_notes,omitempty"`
	UseHelpdeskSaleTimesheet    *Bool      `xmlrpc:"use_helpdesk_sale_timesheet,omitempty"`
	UseHelpdeskTimesheet        *Bool      `xmlrpc:"use_helpdesk_timesheet,omitempty"`
	UseProductRepairs           *Bool      `xmlrpc:"use_product_repairs,omitempty"`
	UseProductReturns           *Bool      `xmlrpc:"use_product_returns,omitempty"`
	UseWebsiteHelpdeskForum     *Bool      `xmlrpc:"use_website_helpdesk_forum,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HelpdeskTickets represents array of helpdesk.ticket model.
type HelpdeskTickets []HelpdeskTicket

// HelpdeskTicketModel is the odoo model name.
const HelpdeskTicketModel = "helpdesk.ticket"

// Many2One convert HelpdeskTicket to *Many2One.
func (ht *HelpdeskTicket) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTicket creates a new helpdesk.ticket model and returns its id.
func (c *Client) CreateHelpdeskTicket(ht *HelpdeskTicket) (int64, error) {
	ids, err := c.CreateHelpdeskTickets([]*HelpdeskTicket{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicket creates a new helpdesk.ticket model and returns its id.
func (c *Client) CreateHelpdeskTickets(hts []*HelpdeskTicket) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketModel, vv, nil)
}

// UpdateHelpdeskTicket updates an existing helpdesk.ticket record.
func (c *Client) UpdateHelpdeskTicket(ht *HelpdeskTicket) error {
	return c.UpdateHelpdeskTickets([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTickets updates existing helpdesk.ticket records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTickets(ids []int64, ht *HelpdeskTicket) error {
	return c.Update(HelpdeskTicketModel, ids, ht, nil)
}

// DeleteHelpdeskTicket deletes an existing helpdesk.ticket record.
func (c *Client) DeleteHelpdeskTicket(id int64) error {
	return c.DeleteHelpdeskTickets([]int64{id})
}

// DeleteHelpdeskTickets deletes existing helpdesk.ticket records.
func (c *Client) DeleteHelpdeskTickets(ids []int64) error {
	return c.Delete(HelpdeskTicketModel, ids)
}

// GetHelpdeskTicket gets helpdesk.ticket existing record.
func (c *Client) GetHelpdeskTicket(id int64) (*HelpdeskTicket, error) {
	hts, err := c.GetHelpdeskTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// GetHelpdeskTickets gets helpdesk.ticket existing records.
func (c *Client) GetHelpdeskTickets(ids []int64) (*HelpdeskTickets, error) {
	hts := &HelpdeskTickets{}
	if err := c.Read(HelpdeskTicketModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTicket finds helpdesk.ticket record by querying it with criteria.
func (c *Client) FindHelpdeskTicket(criteria *Criteria) (*HelpdeskTicket, error) {
	hts := &HelpdeskTickets{}
	if err := c.SearchRead(HelpdeskTicketModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// FindHelpdeskTickets finds helpdesk.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTickets(criteria *Criteria, options *Options) (*HelpdeskTickets, error) {
	hts := &HelpdeskTickets{}
	if err := c.SearchRead(HelpdeskTicketModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskTicketModel, criteria, options)
}

// FindHelpdeskTicketId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
