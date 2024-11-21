package odoo

// GamificationBadge represents gamification.badge model.
type GamificationBadge struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	CanPublish               *Bool      `xmlrpc:"can_publish,omitempty" json:"can_publish,omitempty"`
	ChallengeIds             *Relation  `xmlrpc:"challenge_ids,omitempty" json:"challenge_ids,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	GoalDefinitionIds        *Relation  `xmlrpc:"goal_definition_ids,omitempty" json:"goal_definition_ids,omitempty"`
	GrantedEmployeesCount    *Int       `xmlrpc:"granted_employees_count,omitempty" json:"granted_employees_count,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Image1024                *String    `xmlrpc:"image_1024,omitempty" json:"image_1024,omitempty"`
	Image128                 *String    `xmlrpc:"image_128,omitempty" json:"image_128,omitempty"`
	Image1920                *String    `xmlrpc:"image_1920,omitempty" json:"image_1920,omitempty"`
	Image256                 *String    `xmlrpc:"image_256,omitempty" json:"image_256,omitempty"`
	Image512                 *String    `xmlrpc:"image_512,omitempty" json:"image_512,omitempty"`
	IsPublished              *Bool      `xmlrpc:"is_published,omitempty" json:"is_published,omitempty"`
	Level                    *Selection `xmlrpc:"level,omitempty" json:"level,omitempty"`
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
	OwnerIds                 *Relation  `xmlrpc:"owner_ids,omitempty" json:"owner_ids,omitempty"`
	RemainingSending         *Int       `xmlrpc:"remaining_sending,omitempty" json:"remaining_sending,omitempty"`
	RuleAuth                 *Selection `xmlrpc:"rule_auth,omitempty" json:"rule_auth,omitempty"`
	RuleAuthBadgeIds         *Relation  `xmlrpc:"rule_auth_badge_ids,omitempty" json:"rule_auth_badge_ids,omitempty"`
	RuleAuthUserIds          *Relation  `xmlrpc:"rule_auth_user_ids,omitempty" json:"rule_auth_user_ids,omitempty"`
	RuleMax                  *Bool      `xmlrpc:"rule_max,omitempty" json:"rule_max,omitempty"`
	RuleMaxNumber            *Int       `xmlrpc:"rule_max_number,omitempty" json:"rule_max_number,omitempty"`
	StatCount                *Int       `xmlrpc:"stat_count,omitempty" json:"stat_count,omitempty"`
	StatCountDistinct        *Int       `xmlrpc:"stat_count_distinct,omitempty" json:"stat_count_distinct,omitempty"`
	StatMy                   *Int       `xmlrpc:"stat_my,omitempty" json:"stat_my,omitempty"`
	StatMyMonthlySending     *Int       `xmlrpc:"stat_my_monthly_sending,omitempty" json:"stat_my_monthly_sending,omitempty"`
	StatMyThisMonth          *Int       `xmlrpc:"stat_my_this_month,omitempty" json:"stat_my_this_month,omitempty"`
	StatThisMonth            *Int       `xmlrpc:"stat_this_month,omitempty" json:"stat_this_month,omitempty"`
	UniqueOwnerIds           *Relation  `xmlrpc:"unique_owner_ids,omitempty" json:"unique_owner_ids,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WebsitePublished         *Bool      `xmlrpc:"website_published,omitempty" json:"website_published,omitempty"`
	WebsiteUrl               *String    `xmlrpc:"website_url,omitempty" json:"website_url,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// GamificationBadges represents array of gamification.badge model.
type GamificationBadges []GamificationBadge

// GamificationBadgeModel is the odoo model name.
const GamificationBadgeModel = "gamification.badge"

// Many2One convert GamificationBadge to *Many2One.
func (gb *GamificationBadge) Many2One() *Many2One {
	return NewMany2One(gb.Id.Get(), "")
}

// CreateGamificationBadge creates a new gamification.badge model and returns its id.
func (c *Client) CreateGamificationBadge(gb *GamificationBadge) (int64, error) {
	ids, err := c.CreateGamificationBadges([]*GamificationBadge{gb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationBadge creates a new gamification.badge model and returns its id.
func (c *Client) CreateGamificationBadges(gbs []*GamificationBadge) ([]int64, error) {
	var vv []interface{}
	for _, v := range gbs {
		vv = append(vv, v)
	}
	return c.Create(GamificationBadgeModel, vv, nil)
}

// UpdateGamificationBadge updates an existing gamification.badge record.
func (c *Client) UpdateGamificationBadge(gb *GamificationBadge) error {
	return c.UpdateGamificationBadges([]int64{gb.Id.Get()}, gb)
}

// UpdateGamificationBadges updates existing gamification.badge records.
// All records (represented by ids) will be updated by gb values.
func (c *Client) UpdateGamificationBadges(ids []int64, gb *GamificationBadge) error {
	return c.Update(GamificationBadgeModel, ids, gb, nil)
}

// DeleteGamificationBadge deletes an existing gamification.badge record.
func (c *Client) DeleteGamificationBadge(id int64) error {
	return c.DeleteGamificationBadges([]int64{id})
}

// DeleteGamificationBadges deletes existing gamification.badge records.
func (c *Client) DeleteGamificationBadges(ids []int64) error {
	return c.Delete(GamificationBadgeModel, ids)
}

// GetGamificationBadge gets gamification.badge existing record.
func (c *Client) GetGamificationBadge(id int64) (*GamificationBadge, error) {
	gbs, err := c.GetGamificationBadges([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*gbs)[0]), nil
}

// GetGamificationBadges gets gamification.badge existing records.
func (c *Client) GetGamificationBadges(ids []int64) (*GamificationBadges, error) {
	gbs := &GamificationBadges{}
	if err := c.Read(GamificationBadgeModel, ids, nil, gbs); err != nil {
		return nil, err
	}
	return gbs, nil
}

// FindGamificationBadge finds gamification.badge record by querying it with criteria.
func (c *Client) FindGamificationBadge(criteria *Criteria) (*GamificationBadge, error) {
	gbs := &GamificationBadges{}
	if err := c.SearchRead(GamificationBadgeModel, criteria, NewOptions().Limit(1), gbs); err != nil {
		return nil, err
	}
	return &((*gbs)[0]), nil
}

// FindGamificationBadges finds gamification.badge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadges(criteria *Criteria, options *Options) (*GamificationBadges, error) {
	gbs := &GamificationBadges{}
	if err := c.SearchRead(GamificationBadgeModel, criteria, options, gbs); err != nil {
		return nil, err
	}
	return gbs, nil
}

// FindGamificationBadgeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationBadgeModel, criteria, options)
}

// FindGamificationBadgeId finds record id by querying it with criteria.
func (c *Client) FindGamificationBadgeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationBadgeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
