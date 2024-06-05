package odoo

import (
	"fmt"
)

// PlanningSlot represents planning.slot model.
type PlanningSlot struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	AllocatedHours          *Float     `xmlrpc:"allocated_hours,omptempty"`
	AllocatedPercentage     *Float     `xmlrpc:"allocated_percentage,omptempty"`
	AllocationType          *Selection `xmlrpc:"allocation_type,omptempty"`
	AllowSelfUnassign       *Bool      `xmlrpc:"allow_self_unassign,omptempty"`
	Color                   *Int       `xmlrpc:"color,omptempty"`
	CompanyId               *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId              *Many2One  `xmlrpc:"employee_id,omptempty"`
	EndDatetime             *Time      `xmlrpc:"end_datetime,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	IsAssignedToMe          *Bool      `xmlrpc:"is_assigned_to_me,omptempty"`
	IsPublished             *Bool      `xmlrpc:"is_published,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	OverlapSlotCount        *Int       `xmlrpc:"overlap_slot_count,omptempty"`
	PublicationWarning      *Bool      `xmlrpc:"publication_warning,omptempty"`
	RecurrencyId            *Many2One  `xmlrpc:"recurrency_id,omptempty"`
	Repeat                  *Bool      `xmlrpc:"repeat,omptempty"`
	RepeatInterval          *Int       `xmlrpc:"repeat_interval,omptempty"`
	RepeatType              *Selection `xmlrpc:"repeat_type,omptempty"`
	RepeatUntil             *Time      `xmlrpc:"repeat_until,omptempty"`
	RoleId                  *Many2One  `xmlrpc:"role_id,omptempty"`
	StartDatetime           *Time      `xmlrpc:"start_datetime,omptempty"`
	TemplateAutocompleteIds *Relation  `xmlrpc:"template_autocomplete_ids,omptempty"`
	TemplateCreation        *Bool      `xmlrpc:"template_creation,omptempty"`
	TemplateId              *Many2One  `xmlrpc:"template_id,omptempty"`
	UserId                  *Many2One  `xmlrpc:"user_id,omptempty"`
	WasCopied               *Bool      `xmlrpc:"was_copied,omptempty"`
	WorkingDaysCount        *Int       `xmlrpc:"working_days_count,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PlanningSlots represents array of planning.slot model.
type PlanningSlots []PlanningSlot

// PlanningSlotModel is the odoo model name.
const PlanningSlotModel = "planning.slot"

// Many2One convert PlanningSlot to *Many2One.
func (ps *PlanningSlot) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePlanningSlot creates a new planning.slot model and returns its id.
func (c *Client) CreatePlanningSlot(ps *PlanningSlot) (int64, error) {
	ids, err := c.CreatePlanningSlots([]*PlanningSlot{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSlot creates a new planning.slot model and returns its id.
func (c *Client) CreatePlanningSlots(pss []*PlanningSlot) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PlanningSlotModel, vv)
}

// UpdatePlanningSlot updates an existing planning.slot record.
func (c *Client) UpdatePlanningSlot(ps *PlanningSlot) error {
	return c.UpdatePlanningSlots([]int64{ps.Id.Get()}, ps)
}

// UpdatePlanningSlots updates existing planning.slot records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePlanningSlots(ids []int64, ps *PlanningSlot) error {
	return c.Update(PlanningSlotModel, ids, ps)
}

// DeletePlanningSlot deletes an existing planning.slot record.
func (c *Client) DeletePlanningSlot(id int64) error {
	return c.DeletePlanningSlots([]int64{id})
}

// DeletePlanningSlots deletes existing planning.slot records.
func (c *Client) DeletePlanningSlots(ids []int64) error {
	return c.Delete(PlanningSlotModel, ids)
}

// GetPlanningSlot gets planning.slot existing record.
func (c *Client) GetPlanningSlot(id int64) (*PlanningSlot, error) {
	pss, err := c.GetPlanningSlots([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of planning.slot not found", id)
}

// GetPlanningSlots gets planning.slot existing records.
func (c *Client) GetPlanningSlots(ids []int64) (*PlanningSlots, error) {
	pss := &PlanningSlots{}
	if err := c.Read(PlanningSlotModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSlot finds planning.slot record by querying it with criteria.
func (c *Client) FindPlanningSlot(criteria *Criteria) (*PlanningSlot, error) {
	pss := &PlanningSlots{}
	if err := c.SearchRead(PlanningSlotModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("planning.slot was not found with criteria %v", criteria)
}

// FindPlanningSlots finds planning.slot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlots(criteria *Criteria, options *Options) (*PlanningSlots, error) {
	pss := &PlanningSlots{}
	if err := c.SearchRead(PlanningSlotModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSlotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PlanningSlotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPlanningSlotId finds record id by querying it with criteria.
func (c *Client) FindPlanningSlotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSlotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("planning.slot was not found with criteria %v and options %v", criteria, options)
}
