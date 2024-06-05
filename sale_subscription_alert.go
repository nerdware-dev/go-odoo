package odoo

import (
	"fmt"
)

// SaleSubscriptionAlert represents sale.subscription.alert model.
type SaleSubscriptionAlert struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	Action                        *Selection `xmlrpc:"action,omptempty"`
	ActionServerId                *Many2One  `xmlrpc:"action_server_id,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUser                  *Selection `xmlrpc:"activity_user,omptempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ActivityUserIds               *Relation  `xmlrpc:"activity_user_ids,omptempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omptempty"`
	AutomationId                  *Many2One  `xmlrpc:"automation_id,omptempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omptempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	Code                          *String    `xmlrpc:"code,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomerIds                   *Relation  `xmlrpc:"customer_ids,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	FieldsLines                   *Relation  `xmlrpc:"fields_lines,omptempty"`
	FilterDomain                  *String    `xmlrpc:"filter_domain,omptempty"`
	FilterPreDomain               *String    `xmlrpc:"filter_pre_domain,omptempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omptempty"`
	Help                          *String    `xmlrpc:"help,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	LastRun                       *Time      `xmlrpc:"last_run,omptempty"`
	LinkFieldId                   *Many2One  `xmlrpc:"link_field_id,omptempty"`
	ModelId                       *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName                     *String    `xmlrpc:"model_name,omptempty"`
	MrrChangeAmount               *Float     `xmlrpc:"mrr_change_amount,omptempty"`
	MrrChangePeriod               *Selection `xmlrpc:"mrr_change_period,omptempty"`
	MrrChangeUnit                 *Selection `xmlrpc:"mrr_change_unit,omptempty"`
	MrrMax                        *Float     `xmlrpc:"mrr_max,omptempty"`
	MrrMin                        *Float     `xmlrpc:"mrr_min,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OnChangeFields                *String    `xmlrpc:"on_change_fields,omptempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omptempty"`
	ProductIds                    *Relation  `xmlrpc:"product_ids,omptempty"`
	RatingOperator                *Selection `xmlrpc:"rating_operator,omptempty"`
	RatingPercentage              *Int       `xmlrpc:"rating_percentage,omptempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omptempty"`
	SmsMassKeepLog                *Bool      `xmlrpc:"sms_mass_keep_log,omptempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	StageFromId                   *Many2One  `xmlrpc:"stage_from_id,omptempty"`
	StageId                       *Many2One  `xmlrpc:"stage_id,omptempty"`
	StageToId                     *Many2One  `xmlrpc:"stage_to_id,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omptempty"`
	SubscriptionTemplateIds       *Relation  `xmlrpc:"subscription_template_ids,omptempty"`
	TagId                         *Many2One  `xmlrpc:"tag_id,omptempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omptempty"`
	TrgDateCalendarId             *Many2One  `xmlrpc:"trg_date_calendar_id,omptempty"`
	TrgDateId                     *Many2One  `xmlrpc:"trg_date_id,omptempty"`
	TrgDateRange                  *Int       `xmlrpc:"trg_date_range,omptempty"`
	TrgDateRangeType              *Selection `xmlrpc:"trg_date_range_type,omptempty"`
	TrgDateResourceFieldId        *Many2One  `xmlrpc:"trg_date_resource_field_id,omptempty"`
	Trigger                       *Selection `xmlrpc:"trigger,omptempty"`
	TriggerCondition              *Selection `xmlrpc:"trigger_condition,omptempty"`
	TriggerFieldIds               *Relation  `xmlrpc:"trigger_field_ids,omptempty"`
	Type                          *String    `xmlrpc:"type,omptempty"`
	Usage                         *Selection `xmlrpc:"usage,omptempty"`
	WebsitePath                   *String    `xmlrpc:"website_path,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId                         *String    `xmlrpc:"xml_id,omptempty"`
}

// SaleSubscriptionAlerts represents array of sale.subscription.alert model.
type SaleSubscriptionAlerts []SaleSubscriptionAlert

// SaleSubscriptionAlertModel is the odoo model name.
const SaleSubscriptionAlertModel = "sale.subscription.alert"

// Many2One convert SaleSubscriptionAlert to *Many2One.
func (ssa *SaleSubscriptionAlert) Many2One() *Many2One {
	return NewMany2One(ssa.Id.Get(), "")
}

// CreateSaleSubscriptionAlert creates a new sale.subscription.alert model and returns its id.
func (c *Client) CreateSaleSubscriptionAlert(ssa *SaleSubscriptionAlert) (int64, error) {
	ids, err := c.CreateSaleSubscriptionAlerts([]*SaleSubscriptionAlert{ssa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionAlert creates a new sale.subscription.alert model and returns its id.
func (c *Client) CreateSaleSubscriptionAlerts(ssas []*SaleSubscriptionAlert) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssas {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionAlertModel, vv)
}

// UpdateSaleSubscriptionAlert updates an existing sale.subscription.alert record.
func (c *Client) UpdateSaleSubscriptionAlert(ssa *SaleSubscriptionAlert) error {
	return c.UpdateSaleSubscriptionAlerts([]int64{ssa.Id.Get()}, ssa)
}

// UpdateSaleSubscriptionAlerts updates existing sale.subscription.alert records.
// All records (represented by ids) will be updated by ssa values.
func (c *Client) UpdateSaleSubscriptionAlerts(ids []int64, ssa *SaleSubscriptionAlert) error {
	return c.Update(SaleSubscriptionAlertModel, ids, ssa)
}

// DeleteSaleSubscriptionAlert deletes an existing sale.subscription.alert record.
func (c *Client) DeleteSaleSubscriptionAlert(id int64) error {
	return c.DeleteSaleSubscriptionAlerts([]int64{id})
}

// DeleteSaleSubscriptionAlerts deletes existing sale.subscription.alert records.
func (c *Client) DeleteSaleSubscriptionAlerts(ids []int64) error {
	return c.Delete(SaleSubscriptionAlertModel, ids)
}

// GetSaleSubscriptionAlert gets sale.subscription.alert existing record.
func (c *Client) GetSaleSubscriptionAlert(id int64) (*SaleSubscriptionAlert, error) {
	ssas, err := c.GetSaleSubscriptionAlerts([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssas != nil && len(*ssas) > 0 {
		return &((*ssas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.alert not found", id)
}

// GetSaleSubscriptionAlerts gets sale.subscription.alert existing records.
func (c *Client) GetSaleSubscriptionAlerts(ids []int64) (*SaleSubscriptionAlerts, error) {
	ssas := &SaleSubscriptionAlerts{}
	if err := c.Read(SaleSubscriptionAlertModel, ids, nil, ssas); err != nil {
		return nil, err
	}
	return ssas, nil
}

// FindSaleSubscriptionAlert finds sale.subscription.alert record by querying it with criteria.
func (c *Client) FindSaleSubscriptionAlert(criteria *Criteria) (*SaleSubscriptionAlert, error) {
	ssas := &SaleSubscriptionAlerts{}
	if err := c.SearchRead(SaleSubscriptionAlertModel, criteria, NewOptions().Limit(1), ssas); err != nil {
		return nil, err
	}
	if ssas != nil && len(*ssas) > 0 {
		return &((*ssas)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.alert was not found with criteria %v", criteria)
}

// FindSaleSubscriptionAlerts finds sale.subscription.alert records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionAlerts(criteria *Criteria, options *Options) (*SaleSubscriptionAlerts, error) {
	ssas := &SaleSubscriptionAlerts{}
	if err := c.SearchRead(SaleSubscriptionAlertModel, criteria, options, ssas); err != nil {
		return nil, err
	}
	return ssas, nil
}

// FindSaleSubscriptionAlertIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionAlertIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionAlertModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionAlertId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.alert was not found with criteria %v and options %v", criteria, options)
}
