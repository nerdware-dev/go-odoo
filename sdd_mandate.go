package odoo

// SddMandate represents sdd.mandate model.
type SddMandate struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DebtorIdCode                *String    `xmlrpc:"debtor_id_code,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EndDate                     *Time      `xmlrpc:"end_date,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
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
	OneOff                      *Bool      `xmlrpc:"one_off,omitempty"`
	OriginalDoc                 *String    `xmlrpc:"original_doc,omitempty"`
	OriginalDocFilename         *String    `xmlrpc:"original_doc_filename,omitempty"`
	PaidInvoiceIds              *Relation  `xmlrpc:"paid_invoice_ids,omitempty"`
	PaidInvoicesNber            *Int       `xmlrpc:"paid_invoices_nber,omitempty"`
	PartnerBankId               *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentIds                  *Relation  `xmlrpc:"payment_ids,omitempty"`
	PaymentJournalId            *Many2One  `xmlrpc:"payment_journal_id,omitempty"`
	PaymentsToCollectNber       *Int       `xmlrpc:"payments_to_collect_nber,omitempty"`
	StartDate                   *Time      `xmlrpc:"start_date,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SddMandates represents array of sdd.mandate model.
type SddMandates []SddMandate

// SddMandateModel is the odoo model name.
const SddMandateModel = "sdd.mandate"

// Many2One convert SddMandate to *Many2One.
func (sm *SddMandate) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateSddMandate creates a new sdd.mandate model and returns its id.
func (c *Client) CreateSddMandate(sm *SddMandate) (int64, error) {
	ids, err := c.CreateSddMandates([]*SddMandate{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSddMandate creates a new sdd.mandate model and returns its id.
func (c *Client) CreateSddMandates(sms []*SddMandate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(SddMandateModel, vv, nil)
}

// UpdateSddMandate updates an existing sdd.mandate record.
func (c *Client) UpdateSddMandate(sm *SddMandate) error {
	return c.UpdateSddMandates([]int64{sm.Id.Get()}, sm)
}

// UpdateSddMandates updates existing sdd.mandate records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateSddMandates(ids []int64, sm *SddMandate) error {
	return c.Update(SddMandateModel, ids, sm, nil)
}

// DeleteSddMandate deletes an existing sdd.mandate record.
func (c *Client) DeleteSddMandate(id int64) error {
	return c.DeleteSddMandates([]int64{id})
}

// DeleteSddMandates deletes existing sdd.mandate records.
func (c *Client) DeleteSddMandates(ids []int64) error {
	return c.Delete(SddMandateModel, ids)
}

// GetSddMandate gets sdd.mandate existing record.
func (c *Client) GetSddMandate(id int64) (*SddMandate, error) {
	sms, err := c.GetSddMandates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// GetSddMandates gets sdd.mandate existing records.
func (c *Client) GetSddMandates(ids []int64) (*SddMandates, error) {
	sms := &SddMandates{}
	if err := c.Read(SddMandateModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSddMandate finds sdd.mandate record by querying it with criteria.
func (c *Client) FindSddMandate(criteria *Criteria) (*SddMandate, error) {
	sms := &SddMandates{}
	if err := c.SearchRead(SddMandateModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// FindSddMandates finds sdd.mandate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSddMandates(criteria *Criteria, options *Options) (*SddMandates, error) {
	sms := &SddMandates{}
	if err := c.SearchRead(SddMandateModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSddMandateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSddMandateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SddMandateModel, criteria, options)
}

// FindSddMandateId finds record id by querying it with criteria.
func (c *Client) FindSddMandateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SddMandateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
