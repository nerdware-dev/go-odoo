package odoo

// ApprovalRequest represents approval.request model.
type ApprovalRequest struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	Amount                      *Float     `xmlrpc:"amount,omitempty"`
	ApprovalMinimum             *Int       `xmlrpc:"approval_minimum,omitempty"`
	ApproverIds                 *Relation  `xmlrpc:"approver_ids,omitempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omitempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	DateConfirmed               *Time      `xmlrpc:"date_confirmed,omitempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omitempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasAccessToRequest          *Bool      `xmlrpc:"has_access_to_request,omitempty"`
	HasAmount                   *Selection `xmlrpc:"has_amount,omitempty"`
	HasDate                     *Selection `xmlrpc:"has_date,omitempty"`
	HasItem                     *Selection `xmlrpc:"has_item,omitempty"`
	HasLocation                 *Selection `xmlrpc:"has_location,omitempty"`
	HasPartner                  *Selection `xmlrpc:"has_partner,omitempty"`
	HasPaymentMethod            *Selection `xmlrpc:"has_payment_method,omitempty"`
	HasPeriod                   *Selection `xmlrpc:"has_period,omitempty"`
	HasQuantity                 *Selection `xmlrpc:"has_quantity,omitempty"`
	HasReference                *Selection `xmlrpc:"has_reference,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsManagerApprover           *Bool      `xmlrpc:"is_manager_approver,omitempty"`
	Items                       *String    `xmlrpc:"items,omitempty"`
	Location                    *String    `xmlrpc:"location,omitempty"`
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
	Quantity                    *Float     `xmlrpc:"quantity,omitempty"`
	Reason                      *String    `xmlrpc:"reason,omitempty"`
	Reference                   *String    `xmlrpc:"reference,omitempty"`
	RequestOwnerId              *Many2One  `xmlrpc:"request_owner_id,omitempty"`
	RequestStatus               *Selection `xmlrpc:"request_status,omitempty"`
	RequirerDocument            *Selection `xmlrpc:"requirer_document,omitempty"`
	UserStatus                  *Selection `xmlrpc:"user_status,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ApprovalRequests represents array of approval.request model.
type ApprovalRequests []ApprovalRequest

// ApprovalRequestModel is the odoo model name.
const ApprovalRequestModel = "approval.request"

// Many2One convert ApprovalRequest to *Many2One.
func (ar *ApprovalRequest) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateApprovalRequest creates a new approval.request model and returns its id.
func (c *Client) CreateApprovalRequest(ar *ApprovalRequest) (int64, error) {
	ids, err := c.CreateApprovalRequests([]*ApprovalRequest{ar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApprovalRequest creates a new approval.request model and returns its id.
func (c *Client) CreateApprovalRequests(ars []*ApprovalRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range ars {
		vv = append(vv, v)
	}
	return c.Create(ApprovalRequestModel, vv, nil)
}

// UpdateApprovalRequest updates an existing approval.request record.
func (c *Client) UpdateApprovalRequest(ar *ApprovalRequest) error {
	return c.UpdateApprovalRequests([]int64{ar.Id.Get()}, ar)
}

// UpdateApprovalRequests updates existing approval.request records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateApprovalRequests(ids []int64, ar *ApprovalRequest) error {
	return c.Update(ApprovalRequestModel, ids, ar, nil)
}

// DeleteApprovalRequest deletes an existing approval.request record.
func (c *Client) DeleteApprovalRequest(id int64) error {
	return c.DeleteApprovalRequests([]int64{id})
}

// DeleteApprovalRequests deletes existing approval.request records.
func (c *Client) DeleteApprovalRequests(ids []int64) error {
	return c.Delete(ApprovalRequestModel, ids)
}

// GetApprovalRequest gets approval.request existing record.
func (c *Client) GetApprovalRequest(id int64) (*ApprovalRequest, error) {
	ars, err := c.GetApprovalRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// GetApprovalRequests gets approval.request existing records.
func (c *Client) GetApprovalRequests(ids []int64) (*ApprovalRequests, error) {
	ars := &ApprovalRequests{}
	if err := c.Read(ApprovalRequestModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindApprovalRequest finds approval.request record by querying it with criteria.
func (c *Client) FindApprovalRequest(criteria *Criteria) (*ApprovalRequest, error) {
	ars := &ApprovalRequests{}
	if err := c.SearchRead(ApprovalRequestModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// FindApprovalRequests finds approval.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalRequests(criteria *Criteria, options *Options) (*ApprovalRequests, error) {
	ars := &ApprovalRequests{}
	if err := c.SearchRead(ApprovalRequestModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindApprovalRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApprovalRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ApprovalRequestModel, criteria, options)
}

// FindApprovalRequestId finds record id by querying it with criteria.
func (c *Client) FindApprovalRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApprovalRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
