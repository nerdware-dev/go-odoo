package odoo

// BaseAutomation represents base.automation model.
type BaseAutomation struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	ActionServerId                *Many2One  `xmlrpc:"action_server_id,omitempty" json:"action_server_id,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omitempty" json:"activity_date_deadline_range,omitempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omitempty" json:"activity_date_deadline_range_type,omitempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omitempty" json:"activity_note,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omitempty" json:"activity_user_field_name,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omitempty" json:"activity_user_type,omitempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omitempty" json:"binding_model_id,omitempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omitempty" json:"binding_type,omitempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omitempty" json:"binding_view_types,omitempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omitempty" json:"channel_ids,omitempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	Code                          *String    `xmlrpc:"code,omitempty" json:"code,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omitempty" json:"crud_model_id,omitempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omitempty" json:"crud_model_name,omitempty"`
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
	Name                          *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OnChangeFields                *String    `xmlrpc:"on_change_fields,omitempty" json:"on_change_fields,omitempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omitempty" json:"partner_ids,omitempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	SmsMassKeepLog                *Bool      `xmlrpc:"sms_mass_keep_log,omitempty" json:"sms_mass_keep_log,omitempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omitempty" json:"sms_template_id,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omitempty" json:"template_id,omitempty"`
	TrgDateCalendarId             *Many2One  `xmlrpc:"trg_date_calendar_id,omitempty" json:"trg_date_calendar_id,omitempty"`
	TrgDateId                     *Many2One  `xmlrpc:"trg_date_id,omitempty" json:"trg_date_id,omitempty"`
	TrgDateRange                  *Int       `xmlrpc:"trg_date_range,omitempty" json:"trg_date_range,omitempty"`
	TrgDateRangeType              *Selection `xmlrpc:"trg_date_range_type,omitempty" json:"trg_date_range_type,omitempty"`
	TrgDateResourceFieldId        *Many2One  `xmlrpc:"trg_date_resource_field_id,omitempty" json:"trg_date_resource_field_id,omitempty"`
	Trigger                       *Selection `xmlrpc:"trigger,omitempty" json:"trigger,omitempty"`
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

// BaseAutomations represents array of base.automation model.
type BaseAutomations []BaseAutomation

// BaseAutomationModel is the odoo model name.
const BaseAutomationModel = "base.automation"

// Many2One convert BaseAutomation to *Many2One.
func (ba *BaseAutomation) Many2One() *Many2One {
	return NewMany2One(ba.Id.Get(), "")
}

// CreateBaseAutomation creates a new base.automation model and returns its id.
func (c *Client) CreateBaseAutomation(ba *BaseAutomation) (int64, error) {
	ids, err := c.CreateBaseAutomations([]*BaseAutomation{ba})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseAutomation creates a new base.automation model and returns its id.
func (c *Client) CreateBaseAutomations(bas []*BaseAutomation) ([]int64, error) {
	var vv []interface{}
	for _, v := range bas {
		vv = append(vv, v)
	}
	return c.Create(BaseAutomationModel, vv, nil)
}

// UpdateBaseAutomation updates an existing base.automation record.
func (c *Client) UpdateBaseAutomation(ba *BaseAutomation) error {
	return c.UpdateBaseAutomations([]int64{ba.Id.Get()}, ba)
}

// UpdateBaseAutomations updates existing base.automation records.
// All records (represented by ids) will be updated by ba values.
func (c *Client) UpdateBaseAutomations(ids []int64, ba *BaseAutomation) error {
	return c.Update(BaseAutomationModel, ids, ba, nil)
}

// DeleteBaseAutomation deletes an existing base.automation record.
func (c *Client) DeleteBaseAutomation(id int64) error {
	return c.DeleteBaseAutomations([]int64{id})
}

// DeleteBaseAutomations deletes existing base.automation records.
func (c *Client) DeleteBaseAutomations(ids []int64) error {
	return c.Delete(BaseAutomationModel, ids)
}

// GetBaseAutomation gets base.automation existing record.
func (c *Client) GetBaseAutomation(id int64) (*BaseAutomation, error) {
	bas, err := c.GetBaseAutomations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bas)[0]), nil
}

// GetBaseAutomations gets base.automation existing records.
func (c *Client) GetBaseAutomations(ids []int64) (*BaseAutomations, error) {
	bas := &BaseAutomations{}
	if err := c.Read(BaseAutomationModel, ids, nil, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBaseAutomation finds base.automation record by querying it with criteria.
func (c *Client) FindBaseAutomation(criteria *Criteria) (*BaseAutomation, error) {
	bas := &BaseAutomations{}
	if err := c.SearchRead(BaseAutomationModel, criteria, NewOptions().Limit(1), bas); err != nil {
		return nil, err
	}
	return &((*bas)[0]), nil
}

// FindBaseAutomations finds base.automation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomations(criteria *Criteria, options *Options) (*BaseAutomations, error) {
	bas := &BaseAutomations{}
	if err := c.SearchRead(BaseAutomationModel, criteria, options, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBaseAutomationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseAutomationModel, criteria, options)
}

// FindBaseAutomationId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
