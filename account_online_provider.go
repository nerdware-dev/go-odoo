package odoo

// AccountOnlineProvider represents account.online.provider model.
type AccountOnlineProvider struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	AccountOnlineJournalIds   *Relation  `xmlrpc:"account_online_journal_ids,omitempty"`
	ActionRequired            *Bool      `xmlrpc:"action_required,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	LastRefresh               *Time      `xmlrpc:"last_refresh,omitempty"`
	Message                   *String    `xmlrpc:"message,omitempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	NextRefresh               *Time      `xmlrpc:"next_refresh,omitempty"`
	PlaidErrorType            *String    `xmlrpc:"plaid_error_type,omitempty"`
	PlaidItemId               *String    `xmlrpc:"plaid_item_id,omitempty"`
	PontoToken                *String    `xmlrpc:"ponto_token,omitempty"`
	ProviderAccountIdentifier *String    `xmlrpc:"provider_account_identifier,omitempty"`
	ProviderIdentifier        *String    `xmlrpc:"provider_identifier,omitempty"`
	ProviderType              *Selection `xmlrpc:"provider_type,omitempty"`
	Status                    *String    `xmlrpc:"status,omitempty"`
	StatusCode                *String    `xmlrpc:"status_code,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountOnlineProviders represents array of account.online.provider model.
type AccountOnlineProviders []AccountOnlineProvider

// AccountOnlineProviderModel is the odoo model name.
const AccountOnlineProviderModel = "account.online.provider"

// Many2One convert AccountOnlineProvider to *Many2One.
func (aop *AccountOnlineProvider) Many2One() *Many2One {
	return NewMany2One(aop.Id.Get(), "")
}

// CreateAccountOnlineProvider creates a new account.online.provider model and returns its id.
func (c *Client) CreateAccountOnlineProvider(aop *AccountOnlineProvider) (int64, error) {
	ids, err := c.CreateAccountOnlineProviders([]*AccountOnlineProvider{aop})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineProvider creates a new account.online.provider model and returns its id.
func (c *Client) CreateAccountOnlineProviders(aops []*AccountOnlineProvider) ([]int64, error) {
	var vv []interface{}
	for _, v := range aops {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineProviderModel, vv, nil)
}

// UpdateAccountOnlineProvider updates an existing account.online.provider record.
func (c *Client) UpdateAccountOnlineProvider(aop *AccountOnlineProvider) error {
	return c.UpdateAccountOnlineProviders([]int64{aop.Id.Get()}, aop)
}

// UpdateAccountOnlineProviders updates existing account.online.provider records.
// All records (represented by ids) will be updated by aop values.
func (c *Client) UpdateAccountOnlineProviders(ids []int64, aop *AccountOnlineProvider) error {
	return c.Update(AccountOnlineProviderModel, ids, aop, nil)
}

// DeleteAccountOnlineProvider deletes an existing account.online.provider record.
func (c *Client) DeleteAccountOnlineProvider(id int64) error {
	return c.DeleteAccountOnlineProviders([]int64{id})
}

// DeleteAccountOnlineProviders deletes existing account.online.provider records.
func (c *Client) DeleteAccountOnlineProviders(ids []int64) error {
	return c.Delete(AccountOnlineProviderModel, ids)
}

// GetAccountOnlineProvider gets account.online.provider existing record.
func (c *Client) GetAccountOnlineProvider(id int64) (*AccountOnlineProvider, error) {
	aops, err := c.GetAccountOnlineProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aops)[0]), nil
}

// GetAccountOnlineProviders gets account.online.provider existing records.
func (c *Client) GetAccountOnlineProviders(ids []int64) (*AccountOnlineProviders, error) {
	aops := &AccountOnlineProviders{}
	if err := c.Read(AccountOnlineProviderModel, ids, nil, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAccountOnlineProvider finds account.online.provider record by querying it with criteria.
func (c *Client) FindAccountOnlineProvider(criteria *Criteria) (*AccountOnlineProvider, error) {
	aops := &AccountOnlineProviders{}
	if err := c.SearchRead(AccountOnlineProviderModel, criteria, NewOptions().Limit(1), aops); err != nil {
		return nil, err
	}
	return &((*aops)[0]), nil
}

// FindAccountOnlineProviders finds account.online.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineProviders(criteria *Criteria, options *Options) (*AccountOnlineProviders, error) {
	aops := &AccountOnlineProviders{}
	if err := c.SearchRead(AccountOnlineProviderModel, criteria, options, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAccountOnlineProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOnlineProviderModel, criteria, options)
}

// FindAccountOnlineProviderId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
