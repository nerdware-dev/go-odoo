package odoo

import (
	"fmt"
)

// AccountAsset represents account.asset model.
type AccountAsset struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId             *Many2One  `xmlrpc:"account_analytic_id,omptempty"`
	AccountAssetId                *Many2One  `xmlrpc:"account_asset_id,omptempty"`
	AccountDepreciationExpenseId  *Many2One  `xmlrpc:"account_depreciation_expense_id,omptempty"`
	AccountDepreciationId         *Many2One  `xmlrpc:"account_depreciation_id,omptempty"`
	AcquisitionDate               *Time      `xmlrpc:"acquisition_date,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AnalyticTagIds                *Relation  `xmlrpc:"analytic_tag_ids,omptempty"`
	AssetType                     *Selection `xmlrpc:"asset_type,omptempty"`
	BookValue                     *Float     `xmlrpc:"book_value,omptempty"`
	ChildrenIds                   *Relation  `xmlrpc:"children_ids,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	DepreciationEntriesCount      *Int       `xmlrpc:"depreciation_entries_count,omptempty"`
	DepreciationMoveIds           *Relation  `xmlrpc:"depreciation_move_ids,omptempty"`
	DisplayAccountAssetId         *Bool      `xmlrpc:"display_account_asset_id,omptempty"`
	DisplayModelChoice            *Bool      `xmlrpc:"display_model_choice,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	DisposalDate                  *Time      `xmlrpc:"disposal_date,omptempty"`
	DoNotShowInAssetReport        *Bool      `xmlrpc:"do_not_show_in_asset_report,omptempty"`
	FirstDepreciationDate         *Time      `xmlrpc:"first_depreciation_date,omptempty"`
	GrossIncreaseCount            *Int       `xmlrpc:"gross_increase_count,omptempty"`
	GrossIncreaseValue            *Float     `xmlrpc:"gross_increase_value,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	JournalId                     *Many2One  `xmlrpc:"journal_id,omptempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Method                        *Selection `xmlrpc:"method,omptempty"`
	MethodNumber                  *Int       `xmlrpc:"method_number,omptempty"`
	MethodPeriod                  *Selection `xmlrpc:"method_period,omptempty"`
	MethodProgressFactor          *Float     `xmlrpc:"method_progress_factor,omptempty"`
	ModelId                       *Many2One  `xmlrpc:"model_id,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OriginalMoveLineIds           *Relation  `xmlrpc:"original_move_line_ids,omptempty"`
	OriginalValue                 *Float     `xmlrpc:"original_value,omptempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omptempty"`
	Prorata                       *Bool      `xmlrpc:"prorata,omptempty"`
	ProrataDate                   *Time      `xmlrpc:"prorata_date,omptempty"`
	SalvageValue                  *Float     `xmlrpc:"salvage_value,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	TotalDepreciationEntriesCount *Int       `xmlrpc:"total_depreciation_entries_count,omptempty"`
	UserTypeId                    *Many2One  `xmlrpc:"user_type_id,omptempty"`
	ValueResidual                 *Float     `xmlrpc:"value_residual,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAssets represents array of account.asset model.
type AccountAssets []AccountAsset

// AccountAssetModel is the odoo model name.
const AccountAssetModel = "account.asset"

// Many2One convert AccountAsset to *Many2One.
func (aa *AccountAsset) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAsset(aa *AccountAsset) (int64, error) {
	ids, err := c.CreateAccountAssets([]*AccountAsset{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAssets(aas []*AccountAsset) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetModel, vv)
}

// UpdateAccountAsset updates an existing account.asset record.
func (c *Client) UpdateAccountAsset(aa *AccountAsset) error {
	return c.UpdateAccountAssets([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAssets updates existing account.asset records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAssets(ids []int64, aa *AccountAsset) error {
	return c.Update(AccountAssetModel, ids, aa)
}

// DeleteAccountAsset deletes an existing account.asset record.
func (c *Client) DeleteAccountAsset(id int64) error {
	return c.DeleteAccountAssets([]int64{id})
}

// DeleteAccountAssets deletes existing account.asset records.
func (c *Client) DeleteAccountAssets(ids []int64) error {
	return c.Delete(AccountAssetModel, ids)
}

// GetAccountAsset gets account.asset existing record.
func (c *Client) GetAccountAsset(id int64) (*AccountAsset, error) {
	aas, err := c.GetAccountAssets([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.asset not found", id)
}

// GetAccountAssets gets account.asset existing records.
func (c *Client) GetAccountAssets(ids []int64) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.Read(AccountAssetModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAsset finds account.asset record by querying it with criteria.
func (c *Client) FindAccountAsset(criteria *Criteria) (*AccountAsset, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("account.asset was not found with criteria %v", criteria)
}

// FindAccountAssets finds account.asset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssets(criteria *Criteria, options *Options) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAssetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAssetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAssetId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.asset was not found with criteria %v and options %v", criteria, options)
}
