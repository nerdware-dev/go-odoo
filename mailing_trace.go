package odoo

// MailingTrace represents mailing.trace model.
type MailingTrace struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Bounced          *Time      `xmlrpc:"bounced,omitempty" json:"bounced,omitempty"`
	CampaignId       *Many2One  `xmlrpc:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	Clicked          *Time      `xmlrpc:"clicked,omitempty" json:"clicked,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	Email            *String    `xmlrpc:"email,omitempty" json:"email,omitempty"`
	Exception        *Time      `xmlrpc:"exception,omitempty" json:"exception,omitempty"`
	FailureType      *Selection `xmlrpc:"failure_type,omitempty" json:"failure_type,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Ignored          *Time      `xmlrpc:"ignored,omitempty" json:"ignored,omitempty"`
	LinksClickIds    *Relation  `xmlrpc:"links_click_ids,omitempty" json:"links_click_ids,omitempty"`
	MailMailId       *Many2One  `xmlrpc:"mail_mail_id,omitempty" json:"mail_mail_id,omitempty"`
	MailMailIdInt    *Int       `xmlrpc:"mail_mail_id_int,omitempty" json:"mail_mail_id_int,omitempty"`
	MarketingTraceId *Many2One  `xmlrpc:"marketing_trace_id,omitempty" json:"marketing_trace_id,omitempty"`
	MassMailingId    *Many2One  `xmlrpc:"mass_mailing_id,omitempty" json:"mass_mailing_id,omitempty"`
	MessageId        *String    `xmlrpc:"message_id,omitempty" json:"message_id,omitempty"`
	Model            *String    `xmlrpc:"model,omitempty" json:"model,omitempty"`
	Opened           *Time      `xmlrpc:"opened,omitempty" json:"opened,omitempty"`
	Replied          *Time      `xmlrpc:"replied,omitempty" json:"replied,omitempty"`
	ResId            *Int       `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	Scheduled        *Time      `xmlrpc:"scheduled,omitempty" json:"scheduled,omitempty"`
	Sent             *Time      `xmlrpc:"sent,omitempty" json:"sent,omitempty"`
	State            *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	StateUpdate      *Time      `xmlrpc:"state_update,omitempty" json:"state_update,omitempty"`
	TraceType        *Selection `xmlrpc:"trace_type,omitempty" json:"trace_type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// MailingTraces represents array of mailing.trace model.
type MailingTraces []MailingTrace

// MailingTraceModel is the odoo model name.
const MailingTraceModel = "mailing.trace"

// Many2One convert MailingTrace to *Many2One.
func (mt *MailingTrace) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMailingTrace creates a new mailing.trace model and returns its id.
func (c *Client) CreateMailingTrace(mt *MailingTrace) (int64, error) {
	ids, err := c.CreateMailingTraces([]*MailingTrace{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingTrace creates a new mailing.trace model and returns its id.
func (c *Client) CreateMailingTraces(mts []*MailingTrace) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MailingTraceModel, vv, nil)
}

// UpdateMailingTrace updates an existing mailing.trace record.
func (c *Client) UpdateMailingTrace(mt *MailingTrace) error {
	return c.UpdateMailingTraces([]int64{mt.Id.Get()}, mt)
}

// UpdateMailingTraces updates existing mailing.trace records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailingTraces(ids []int64, mt *MailingTrace) error {
	return c.Update(MailingTraceModel, ids, mt, nil)
}

// DeleteMailingTrace deletes an existing mailing.trace record.
func (c *Client) DeleteMailingTrace(id int64) error {
	return c.DeleteMailingTraces([]int64{id})
}

// DeleteMailingTraces deletes existing mailing.trace records.
func (c *Client) DeleteMailingTraces(ids []int64) error {
	return c.Delete(MailingTraceModel, ids)
}

// GetMailingTrace gets mailing.trace existing record.
func (c *Client) GetMailingTrace(id int64) (*MailingTrace, error) {
	mts, err := c.GetMailingTraces([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// GetMailingTraces gets mailing.trace existing records.
func (c *Client) GetMailingTraces(ids []int64) (*MailingTraces, error) {
	mts := &MailingTraces{}
	if err := c.Read(MailingTraceModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailingTrace finds mailing.trace record by querying it with criteria.
func (c *Client) FindMailingTrace(criteria *Criteria) (*MailingTrace, error) {
	mts := &MailingTraces{}
	if err := c.SearchRead(MailingTraceModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// FindMailingTraces finds mailing.trace records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraces(criteria *Criteria, options *Options) (*MailingTraces, error) {
	mts := &MailingTraces{}
	if err := c.SearchRead(MailingTraceModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailingTraceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingTraceModel, criteria, options)
}

// FindMailingTraceId finds record id by querying it with criteria.
func (c *Client) FindMailingTraceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingTraceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
