package odoo

// PlanningSlotReportAnalysis represents planning.slot.report.analysis model.
type PlanningSlotReportAnalysis struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omitempty"`
	EntryDate   *Time     `xmlrpc:"entry_date,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	NumberHours *Float    `xmlrpc:"number_hours,omitempty"`
	RoleId      *Many2One `xmlrpc:"role_id,omitempty"`
}

// PlanningSlotReportAnalysiss represents array of planning.slot.report.analysis model.
type PlanningSlotReportAnalysiss []PlanningSlotReportAnalysis

// PlanningSlotReportAnalysisModel is the odoo model name.
const PlanningSlotReportAnalysisModel = "planning.slot.report.analysis"

// Many2One convert PlanningSlotReportAnalysis to *Many2One.
func (psra *PlanningSlotReportAnalysis) Many2One() *Many2One {
	return NewMany2One(psra.Id.Get(), "")
}

// CreatePlanningSlotReportAnalysis creates a new planning.slot.report.analysis model and returns its id.
func (c *Client) CreatePlanningSlotReportAnalysis(psra *PlanningSlotReportAnalysis) (int64, error) {
	ids, err := c.CreatePlanningSlotReportAnalysiss([]*PlanningSlotReportAnalysis{psra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSlotReportAnalysis creates a new planning.slot.report.analysis model and returns its id.
func (c *Client) CreatePlanningSlotReportAnalysiss(psras []*PlanningSlotReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range psras {
		vv = append(vv, v)
	}
	return c.Create(PlanningSlotReportAnalysisModel, vv, nil)
}

// UpdatePlanningSlotReportAnalysis updates an existing planning.slot.report.analysis record.
func (c *Client) UpdatePlanningSlotReportAnalysis(psra *PlanningSlotReportAnalysis) error {
	return c.UpdatePlanningSlotReportAnalysiss([]int64{psra.Id.Get()}, psra)
}

// UpdatePlanningSlotReportAnalysiss updates existing planning.slot.report.analysis records.
// All records (represented by ids) will be updated by psra values.
func (c *Client) UpdatePlanningSlotReportAnalysiss(ids []int64, psra *PlanningSlotReportAnalysis) error {
	return c.Update(PlanningSlotReportAnalysisModel, ids, psra, nil)
}

// DeletePlanningSlotReportAnalysis deletes an existing planning.slot.report.analysis record.
func (c *Client) DeletePlanningSlotReportAnalysis(id int64) error {
	return c.DeletePlanningSlotReportAnalysiss([]int64{id})
}

// DeletePlanningSlotReportAnalysiss deletes existing planning.slot.report.analysis records.
func (c *Client) DeletePlanningSlotReportAnalysiss(ids []int64) error {
	return c.Delete(PlanningSlotReportAnalysisModel, ids)
}

// GetPlanningSlotReportAnalysis gets planning.slot.report.analysis existing record.
func (c *Client) GetPlanningSlotReportAnalysis(id int64) (*PlanningSlotReportAnalysis, error) {
	psras, err := c.GetPlanningSlotReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*psras)[0]), nil
}

// GetPlanningSlotReportAnalysiss gets planning.slot.report.analysis existing records.
func (c *Client) GetPlanningSlotReportAnalysiss(ids []int64) (*PlanningSlotReportAnalysiss, error) {
	psras := &PlanningSlotReportAnalysiss{}
	if err := c.Read(PlanningSlotReportAnalysisModel, ids, nil, psras); err != nil {
		return nil, err
	}
	return psras, nil
}

// FindPlanningSlotReportAnalysis finds planning.slot.report.analysis record by querying it with criteria.
func (c *Client) FindPlanningSlotReportAnalysis(criteria *Criteria) (*PlanningSlotReportAnalysis, error) {
	psras := &PlanningSlotReportAnalysiss{}
	if err := c.SearchRead(PlanningSlotReportAnalysisModel, criteria, NewOptions().Limit(1), psras); err != nil {
		return nil, err
	}
	return &((*psras)[0]), nil
}

// FindPlanningSlotReportAnalysiss finds planning.slot.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotReportAnalysiss(criteria *Criteria, options *Options) (*PlanningSlotReportAnalysiss, error) {
	psras := &PlanningSlotReportAnalysiss{}
	if err := c.SearchRead(PlanningSlotReportAnalysisModel, criteria, options, psras); err != nil {
		return nil, err
	}
	return psras, nil
}

// FindPlanningSlotReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningSlotReportAnalysisModel, criteria, options)
}

// FindPlanningSlotReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindPlanningSlotReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSlotReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
