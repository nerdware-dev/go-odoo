package odoo

// CalendarEvent represents calendar.event model.
type CalendarEvent struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	AlarmIds                 *Relation  `xmlrpc:"alarm_ids,omitempty" json:"alarm_ids,omitempty"`
	Allday                   *Bool      `xmlrpc:"allday,omitempty" json:"allday,omitempty"`
	ApplicantId              *Many2One  `xmlrpc:"applicant_id,omitempty" json:"applicant_id,omitempty"`
	AttendeeIds              *Relation  `xmlrpc:"attendee_ids,omitempty" json:"attendee_ids,omitempty"`
	AttendeeStatus           *Selection `xmlrpc:"attendee_status,omitempty" json:"attendee_status,omitempty"`
	Byday                    *Selection `xmlrpc:"byday,omitempty" json:"byday,omitempty"`
	CategIds                 *Relation  `xmlrpc:"categ_ids,omitempty" json:"categ_ids,omitempty"`
	Count                    *Int       `xmlrpc:"count,omitempty" json:"count,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Day                      *Int       `xmlrpc:"day,omitempty" json:"day,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DisplayStart             *String    `xmlrpc:"display_start,omitempty" json:"display_start,omitempty"`
	DisplayTime              *String    `xmlrpc:"display_time,omitempty" json:"display_time,omitempty"`
	Duration                 *Float     `xmlrpc:"duration,omitempty" json:"duration,omitempty"`
	EndType                  *Selection `xmlrpc:"end_type,omitempty" json:"end_type,omitempty"`
	EventTz                  *Selection `xmlrpc:"event_tz,omitempty" json:"event_tz,omitempty"`
	FinalDate                *Time      `xmlrpc:"final_date,omitempty" json:"final_date,omitempty"`
	Fr                       *Bool      `xmlrpc:"fr,omitempty" json:"fr,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Interval                 *Int       `xmlrpc:"interval,omitempty" json:"interval,omitempty"`
	IsAttendee               *Bool      `xmlrpc:"is_attendee,omitempty" json:"is_attendee,omitempty"`
	IsHighlighted            *Bool      `xmlrpc:"is_highlighted,omitempty" json:"is_highlighted,omitempty"`
	Location                 *String    `xmlrpc:"location,omitempty" json:"location,omitempty"`
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
	Mo                       *Bool      `xmlrpc:"mo,omitempty" json:"mo,omitempty"`
	MonthBy                  *Selection `xmlrpc:"month_by,omitempty" json:"month_by,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omitempty" json:"opportunity_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omitempty" json:"partner_ids,omitempty"`
	Privacy                  *Selection `xmlrpc:"privacy,omitempty" json:"privacy,omitempty"`
	Recurrency               *Bool      `xmlrpc:"recurrency,omitempty" json:"recurrency,omitempty"`
	RecurrentId              *Int       `xmlrpc:"recurrent_id,omitempty" json:"recurrent_id,omitempty"`
	RecurrentIdDate          *Time      `xmlrpc:"recurrent_id_date,omitempty" json:"recurrent_id_date,omitempty"`
	ResId                    *Int       `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	ResModel                 *String    `xmlrpc:"res_model,omitempty" json:"res_model,omitempty"`
	ResModelId               *Many2One  `xmlrpc:"res_model_id,omitempty" json:"res_model_id,omitempty"`
	Rrule                    *String    `xmlrpc:"rrule,omitempty" json:"rrule,omitempty"`
	RruleType                *Selection `xmlrpc:"rrule_type,omitempty" json:"rrule_type,omitempty"`
	Sa                       *Bool      `xmlrpc:"sa,omitempty" json:"sa,omitempty"`
	ShowAs                   *Selection `xmlrpc:"show_as,omitempty" json:"show_as,omitempty"`
	Start                    *Time      `xmlrpc:"start,omitempty" json:"start,omitempty"`
	StartDate                *Time      `xmlrpc:"start_date,omitempty" json:"start_date,omitempty"`
	StartDatetime            *Time      `xmlrpc:"start_datetime,omitempty" json:"start_datetime,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	Stop                     *Time      `xmlrpc:"stop,omitempty" json:"stop,omitempty"`
	StopDate                 *Time      `xmlrpc:"stop_date,omitempty" json:"stop_date,omitempty"`
	StopDatetime             *Time      `xmlrpc:"stop_datetime,omitempty" json:"stop_datetime,omitempty"`
	Su                       *Bool      `xmlrpc:"su,omitempty" json:"su,omitempty"`
	Th                       *Bool      `xmlrpc:"th,omitempty" json:"th,omitempty"`
	Tu                       *Bool      `xmlrpc:"tu,omitempty" json:"tu,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	We                       *Bool      `xmlrpc:"we,omitempty" json:"we,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WeekList                 *Selection `xmlrpc:"week_list,omitempty" json:"week_list,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// CalendarEvents represents array of calendar.event model.
type CalendarEvents []CalendarEvent

// CalendarEventModel is the odoo model name.
const CalendarEventModel = "calendar.event"

// Many2One convert CalendarEvent to *Many2One.
func (ce *CalendarEvent) Many2One() *Many2One {
	return NewMany2One(ce.Id.Get(), "")
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvent(ce *CalendarEvent) (int64, error) {
	ids, err := c.CreateCalendarEvents([]*CalendarEvent{ce})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvents(ces []*CalendarEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ces {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventModel, vv, nil)
}

// UpdateCalendarEvent updates an existing calendar.event record.
func (c *Client) UpdateCalendarEvent(ce *CalendarEvent) error {
	return c.UpdateCalendarEvents([]int64{ce.Id.Get()}, ce)
}

// UpdateCalendarEvents updates existing calendar.event records.
// All records (represented by ids) will be updated by ce values.
func (c *Client) UpdateCalendarEvents(ids []int64, ce *CalendarEvent) error {
	return c.Update(CalendarEventModel, ids, ce, nil)
}

// DeleteCalendarEvent deletes an existing calendar.event record.
func (c *Client) DeleteCalendarEvent(id int64) error {
	return c.DeleteCalendarEvents([]int64{id})
}

// DeleteCalendarEvents deletes existing calendar.event records.
func (c *Client) DeleteCalendarEvents(ids []int64) error {
	return c.Delete(CalendarEventModel, ids)
}

// GetCalendarEvent gets calendar.event existing record.
func (c *Client) GetCalendarEvent(id int64) (*CalendarEvent, error) {
	ces, err := c.GetCalendarEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ces)[0]), nil
}

// GetCalendarEvents gets calendar.event existing records.
func (c *Client) GetCalendarEvents(ids []int64) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.Read(CalendarEventModel, ids, nil, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEvent finds calendar.event record by querying it with criteria.
func (c *Client) FindCalendarEvent(criteria *Criteria) (*CalendarEvent, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, NewOptions().Limit(1), ces); err != nil {
		return nil, err
	}
	return &((*ces)[0]), nil
}

// FindCalendarEvents finds calendar.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEvents(criteria *Criteria, options *Options) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, options, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarEventModel, criteria, options)
}

// FindCalendarEventId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
