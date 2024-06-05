package odoo

import (
	"fmt"
)

// ApprovalRequest represents approval.request model.
type ApprovalRequest struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	Amount                      *Float     `xmlrpc:"amount,omptempty"`
	ApprovalMinimum             *Int       `xmlrpc:"approval_minimum,omptempty"`
	ApproverIds                 *Relation  `xmlrpc:"approver_ids,omptempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omptempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                        *Time      `xmlrpc:"date,omptempty"`
	DateConfirmed               *Time      `xmlrpc:"date_confirmed,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	HasAccessToRequest          *Bool      `xmlrpc:"has_access_to_request,omptempty"`
	HasAmount                   *Selection `xmlrpc:"has_amount,omptempty"`
	HasDate                     *Selection `xmlrpc:"has_date,omptempty"`
	HasItem                     *Selection `xmlrpc:"has_item,omptempty"`
	HasLocation                 *Selection `xmlrpc:"has_location,omptempty"`
	HasPartner                  *Selection `xmlrpc:"has_partner,omptempty"`
	HasPaymentMethod            *Selection `xmlrpc:"has_payment_method,omptempty"`
	HasPeriod                   *Selection `xmlrpc:"has_period,omptempty"`
	HasQuantity                 *Selection `xmlrpc:"has_quantity,omptempty"`
	HasReference                *Selection `xmlrpc:"has_reference,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsManagerApprover           *Bool      `xmlrpc:"is_manager_approver,omptempty"`
	Items                       *String    `xmlrpc:"items,omptempty"`
	Location                    *String    `xmlrpc:"location,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	Quantity                    *Float     `xmlrpc:"quantity,omptempty"`
	Reason                      *String    `xmlrpc:"reason,omptempty"`
	Reference                   *String    `xmlrpc:"reference,omptempty"`
	RequestOwnerId              *Many2One  `xmlrpc:"request_owner_id,omptempty"`
	RequestStatus               *Selection `xmlrpc:"request_status,omptempty"`
	RequirerDocument            *Selection `xmlrpc:"requirer_document,omptempty"`
	UserStatus                  *Selection `xmlrpc:"user_status,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(ApprovalRequestModel, vv)
}

// UpdateApprovalRequest updates an existing approval.request record.
func (c *Client) UpdateApprovalRequest(ar *ApprovalRequest) error {
	return c.UpdateApprovalRequests([]int64{ar.Id.Get()}, ar)
}

// UpdateApprovalRequests updates existing approval.request records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateApprovalRequests(ids []int64, ar *ApprovalRequest) error {
	return c.Update(ApprovalRequestModel, ids, ar)
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
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of approval.request not found", id)
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
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("approval.request was not found with criteria %v", criteria)
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
	ids, err := c.Search(ApprovalRequestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindApprovalRequestId finds record id by querying it with criteria.
func (c *Client) FindApprovalRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApprovalRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("approval.request was not found with criteria %v and options %v", criteria, options)
}
