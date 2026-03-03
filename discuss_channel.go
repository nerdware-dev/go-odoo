package odoo

// DiscussChannel represents discuss.channel model.
type DiscussChannel struct {
	Active                    *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	AllowPublicUpload         *Bool      `xmlrpc:"allow_public_upload,omitempty" json:"allow_public_upload,omitempty"`
	Avatar128                 *String    `xmlrpc:"avatar_128,omitempty" json:"avatar_128,omitempty"`
	ChannelMemberIds          *Relation  `xmlrpc:"channel_member_ids,omitempty" json:"channel_member_ids,omitempty"`
	ChannelPartnerIds         *Relation  `xmlrpc:"channel_partner_ids,omitempty" json:"channel_partner_ids,omitempty"`
	ChannelType               *Selection `xmlrpc:"channel_type,omitempty" json:"channel_type,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DefaultDisplayMode        *Selection `xmlrpc:"default_display_mode,omitempty" json:"default_display_mode,omitempty"`
	Description               *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	GroupIds                  *Relation  `xmlrpc:"group_ids,omitempty" json:"group_ids,omitempty"`
	GroupPublicId             *Many2One  `xmlrpc:"group_public_id,omitempty" json:"group_public_id,omitempty"`
	HasMessage                *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Image128                  *String    `xmlrpc:"image_128,omitempty" json:"image_128,omitempty"`
	InvitationUrl             *String    `xmlrpc:"invitation_url,omitempty" json:"invitation_url,omitempty"`
	IsChat                    *Bool      `xmlrpc:"is_chat,omitempty" json:"is_chat,omitempty"`
	IsEditable                *Bool      `xmlrpc:"is_editable,omitempty" json:"is_editable,omitempty"`
	IsMember                  *Bool      `xmlrpc:"is_member,omitempty" json:"is_member,omitempty"`
	MemberCount               *Int       `xmlrpc:"member_count,omitempty" json:"member_count,omitempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	PinnedMessageIds          *Relation  `xmlrpc:"pinned_message_ids,omitempty" json:"pinned_message_ids,omitempty"`
	RatingIds                 *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RtcSessionIds             *Relation  `xmlrpc:"rtc_session_ids,omitempty" json:"rtc_session_ids,omitempty"`
	SfuChannelUuid            *String    `xmlrpc:"sfu_channel_uuid,omitempty" json:"sfu_channel_uuid,omitempty"`
	SfuServerUrl              *String    `xmlrpc:"sfu_server_url,omitempty" json:"sfu_server_url,omitempty"`
	SubscriptionDepartmentIds *Relation  `xmlrpc:"subscription_department_ids,omitempty" json:"subscription_department_ids,omitempty"`
	Uuid                      *String    `xmlrpc:"uuid,omitempty" json:"uuid,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// DiscussChannels represents array of discuss.channel model.
type DiscussChannels []DiscussChannel

// DiscussChannelModel is the odoo model name.
const DiscussChannelModel = "discuss.channel"

// Many2One convert DiscussChannel to *Many2One.
func (dc *DiscussChannel) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannel(dc *DiscussChannel) (int64, error) {
	ids, err := c.CreateDiscussChannels([]*DiscussChannel{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannels(dcs []*DiscussChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DiscussChannelModel, vv, nil)
}

// UpdateDiscussChannel updates an existing discuss.channel record.
func (c *Client) UpdateDiscussChannel(dc *DiscussChannel) error {
	return c.UpdateDiscussChannels([]int64{dc.Id.Get()}, dc)
}

// UpdateDiscussChannels updates existing discuss.channel records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDiscussChannels(ids []int64, dc *DiscussChannel) error {
	return c.Update(DiscussChannelModel, ids, dc, nil)
}

// DeleteDiscussChannel deletes an existing discuss.channel record.
func (c *Client) DeleteDiscussChannel(id int64) error {
	return c.DeleteDiscussChannels([]int64{id})
}

// DeleteDiscussChannels deletes existing discuss.channel records.
func (c *Client) DeleteDiscussChannels(ids []int64) error {
	return c.Delete(DiscussChannelModel, ids)
}

// GetDiscussChannel gets discuss.channel existing record.
func (c *Client) GetDiscussChannel(id int64) (*DiscussChannel, error) {
	dcs, err := c.GetDiscussChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*dcs) == 0 {
		return nil, nil
	}
	return &((*dcs)[0]), nil
}

// GetDiscussChannels gets discuss.channel existing records.
func (c *Client) GetDiscussChannels(ids []int64) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.Read(DiscussChannelModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannel finds discuss.channel record by querying it with criteria.
func (c *Client) FindDiscussChannel(criteria *Criteria) (*DiscussChannel, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	if len(*dcs) == 0 {
		return nil, nil
	}
	return &((*dcs)[0]), nil
}

// FindDiscussChannels finds discuss.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannels(criteria *Criteria, options *Options) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussChannelModel, criteria, options)
}

// FindDiscussChannelId finds record id by querying it with criteria.
func (c *Client) FindDiscussChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
