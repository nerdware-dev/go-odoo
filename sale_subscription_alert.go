package odoo

// SaleSubscriptionAlert represents sale.subscription.alert model.
type SaleSubscriptionAlert struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Action                        *Selection `xmlrpc:"action,omitempty" json:"action,omitempty"`
	ActionServerId                *Many2One  `xmlrpc:"action_server_id,omitempty" json:"action_server_id,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omitempty" json:"activity_date_deadline_range,omitempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omitempty" json:"activity_date_deadline_range_type,omitempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omitempty" json:"activity_note,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUser                  *Selection `xmlrpc:"activity_user,omitempty" json:"activity_user,omitempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omitempty" json:"activity_user_field_name,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	ActivityUserIds               *Relation  `xmlrpc:"activity_user_ids,omitempty" json:"activity_user_ids,omitempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omitempty" json:"activity_user_type,omitempty"`
	AutomationId                  *Many2One  `xmlrpc:"automation_id,omitempty" json:"automation_id,omitempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omitempty" json:"binding_model_id,omitempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omitempty" json:"binding_type,omitempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omitempty" json:"binding_view_types,omitempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omitempty" json:"channel_ids,omitempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	Code                          *String    `xmlrpc:"code,omitempty" json:"code,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omitempty" json:"crud_model_id,omitempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omitempty" json:"crud_model_name,omitempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	CustomerIds                   *Relation  `xmlrpc:"customer_ids,omitempty" json:"customer_ids,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	FieldsLines                   *Relation  `xmlrpc:"fields_lines,omitempty" json:"fields_lines,omitempty"`
	FilterDomain                  *String    `xmlrpc:"filter_domain,omitempty" json:"filter_domain,omitempty"`
	FilterPreDomain               *String    `xmlrpc:"filter_pre_domain,omitempty" json:"filter_pre_domain,omitempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omitempty" json:"groups_id,omitempty"`
	Help                          *String    `xmlrpc:"help,omitempty" json:"help,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	LastRun                       *Time      `xmlrpc:"last_run,omitempty" json:"last_run,omitempty"`
	LinkFieldId                   *Many2One  `xmlrpc:"link_field_id,omitempty" json:"link_field_id,omitempty"`
	ModelId                       *Many2One  `xmlrpc:"model_id,omitempty" json:"model_id,omitempty"`
	ModelName                     *String    `xmlrpc:"model_name,omitempty" json:"model_name,omitempty"`
	MrrChangeAmount               *Float     `xmlrpc:"mrr_change_amount,omitempty" json:"mrr_change_amount,omitempty"`
	MrrChangePeriod               *Selection `xmlrpc:"mrr_change_period,omitempty" json:"mrr_change_period,omitempty"`
	MrrChangeUnit                 *Selection `xmlrpc:"mrr_change_unit,omitempty" json:"mrr_change_unit,omitempty"`
	MrrMax                        *Float     `xmlrpc:"mrr_max,omitempty" json:"mrr_max,omitempty"`
	MrrMin                        *Float     `xmlrpc:"mrr_min,omitempty" json:"mrr_min,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OnChangeFields                *String    `xmlrpc:"on_change_fields,omitempty" json:"on_change_fields,omitempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omitempty" json:"partner_ids,omitempty"`
	ProductIds                    *Relation  `xmlrpc:"product_ids,omitempty" json:"product_ids,omitempty"`
	RatingOperator                *Selection `xmlrpc:"rating_operator,omitempty" json:"rating_operator,omitempty"`
	RatingPercentage              *Int       `xmlrpc:"rating_percentage,omitempty" json:"rating_percentage,omitempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	SmsMassKeepLog                *Bool      `xmlrpc:"sms_mass_keep_log,omitempty" json:"sms_mass_keep_log,omitempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omitempty" json:"sms_template_id,omitempty"`
	StageFromId                   *Many2One  `xmlrpc:"stage_from_id,omitempty" json:"stage_from_id,omitempty"`
	StageId                       *Many2One  `xmlrpc:"stage_id,omitempty" json:"stage_id,omitempty"`
	StageToId                     *Many2One  `xmlrpc:"stage_to_id,omitempty" json:"stage_to_id,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omitempty" json:"subscription_count,omitempty"`
	SubscriptionTemplateIds       *Relation  `xmlrpc:"subscription_template_ids,omitempty" json:"subscription_template_ids,omitempty"`
	TagId                         *Many2One  `xmlrpc:"tag_id,omitempty" json:"tag_id,omitempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omitempty" json:"template_id,omitempty"`
	TrgDateCalendarId             *Many2One  `xmlrpc:"trg_date_calendar_id,omitempty" json:"trg_date_calendar_id,omitempty"`
	TrgDateId                     *Many2One  `xmlrpc:"trg_date_id,omitempty" json:"trg_date_id,omitempty"`
	TrgDateRange                  *Int       `xmlrpc:"trg_date_range,omitempty" json:"trg_date_range,omitempty"`
	TrgDateRangeType              *Selection `xmlrpc:"trg_date_range_type,omitempty" json:"trg_date_range_type,omitempty"`
	TrgDateResourceFieldId        *Many2One  `xmlrpc:"trg_date_resource_field_id,omitempty" json:"trg_date_resource_field_id,omitempty"`
	Trigger                       *Selection `xmlrpc:"trigger,omitempty" json:"trigger,omitempty"`
	TriggerCondition              *Selection `xmlrpc:"trigger_condition,omitempty" json:"trigger_condition,omitempty"`
	TriggerFieldIds               *Relation  `xmlrpc:"trigger_field_ids,omitempty" json:"trigger_field_ids,omitempty"`
	Type                          *String    `xmlrpc:"type,omitempty" json:"type,omitempty"`
	Usage                         *Selection `xmlrpc:"usage,omitempty" json:"usage,omitempty"`
	WebsitePath                   *String    `xmlrpc:"website_path,omitempty" json:"website_path,omitempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omitempty" json:"website_published,omitempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omitempty" json:"website_url,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
	XmlId                         *String    `xmlrpc:"xml_id,omitempty" json:"xml_id,omitempty"`
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
	return c.Create(SaleSubscriptionAlertModel, vv, nil)
}

// UpdateSaleSubscriptionAlert updates an existing sale.subscription.alert record.
func (c *Client) UpdateSaleSubscriptionAlert(ssa *SaleSubscriptionAlert) error {
	return c.UpdateSaleSubscriptionAlerts([]int64{ssa.Id.Get()}, ssa)
}

// UpdateSaleSubscriptionAlerts updates existing sale.subscription.alert records.
// All records (represented by ids) will be updated by ssa values.
func (c *Client) UpdateSaleSubscriptionAlerts(ids []int64, ssa *SaleSubscriptionAlert) error {
	return c.Update(SaleSubscriptionAlertModel, ids, ssa, nil)
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
	return &((*ssas)[0]), nil
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
	return &((*ssas)[0]), nil
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
	return c.Search(SaleSubscriptionAlertModel, criteria, options)
}

// FindSaleSubscriptionAlertId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
