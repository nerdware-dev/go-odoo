package odoo

// HelpdeskSlaReportAnalysis represents helpdesk.sla.report.analysis model.
type HelpdeskSlaReportAnalysis struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omitempty"`
	Priority               *Selection `xmlrpc:"priority,omitempty"`
	SlaDeadline            *Time      `xmlrpc:"sla_deadline,omitempty"`
	SlaExceededDays        *Int       `xmlrpc:"sla_exceeded_days,omitempty"`
	SlaId                  *Many2One  `xmlrpc:"sla_id,omitempty"`
	SlaReachedDatetime     *Time      `xmlrpc:"sla_reached_datetime,omitempty"`
	SlaStageId             *Many2One  `xmlrpc:"sla_stage_id,omitempty"`
	SlaStatus              *Selection `xmlrpc:"sla_status,omitempty"`
	SlaStatusFailed        *Bool      `xmlrpc:"sla_status_failed,omitempty"`
	TeamId                 *Many2One  `xmlrpc:"team_id,omitempty"`
	TicketAssignationHours *Int       `xmlrpc:"ticket_assignation_hours,omitempty"`
	TicketCloseHours       *Int       `xmlrpc:"ticket_close_hours,omitempty"`
	TicketClosed           *Bool      `xmlrpc:"ticket_closed,omitempty"`
	TicketDeadline         *Time      `xmlrpc:"ticket_deadline,omitempty"`
	TicketFailed           *Bool      `xmlrpc:"ticket_failed,omitempty"`
	TicketId               *Many2One  `xmlrpc:"ticket_id,omitempty"`
	TicketOpenHours        *Int       `xmlrpc:"ticket_open_hours,omitempty"`
	TicketStageId          *Many2One  `xmlrpc:"ticket_stage_id,omitempty"`
	TicketTypeId           *Many2One  `xmlrpc:"ticket_type_id,omitempty"`
	UserId                 *Many2One  `xmlrpc:"user_id,omitempty"`
}

// HelpdeskSlaReportAnalysiss represents array of helpdesk.sla.report.analysis model.
type HelpdeskSlaReportAnalysiss []HelpdeskSlaReportAnalysis

// HelpdeskSlaReportAnalysisModel is the odoo model name.
const HelpdeskSlaReportAnalysisModel = "helpdesk.sla.report.analysis"

// Many2One convert HelpdeskSlaReportAnalysis to *Many2One.
func (hsra *HelpdeskSlaReportAnalysis) Many2One() *Many2One {
	return NewMany2One(hsra.Id.Get(), "")
}

// CreateHelpdeskSlaReportAnalysis creates a new helpdesk.sla.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskSlaReportAnalysis(hsra *HelpdeskSlaReportAnalysis) (int64, error) {
	ids, err := c.CreateHelpdeskSlaReportAnalysiss([]*HelpdeskSlaReportAnalysis{hsra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskSlaReportAnalysis creates a new helpdesk.sla.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskSlaReportAnalysiss(hsras []*HelpdeskSlaReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsras {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskSlaReportAnalysisModel, vv, nil)
}

// UpdateHelpdeskSlaReportAnalysis updates an existing helpdesk.sla.report.analysis record.
func (c *Client) UpdateHelpdeskSlaReportAnalysis(hsra *HelpdeskSlaReportAnalysis) error {
	return c.UpdateHelpdeskSlaReportAnalysiss([]int64{hsra.Id.Get()}, hsra)
}

// UpdateHelpdeskSlaReportAnalysiss updates existing helpdesk.sla.report.analysis records.
// All records (represented by ids) will be updated by hsra values.
func (c *Client) UpdateHelpdeskSlaReportAnalysiss(ids []int64, hsra *HelpdeskSlaReportAnalysis) error {
	return c.Update(HelpdeskSlaReportAnalysisModel, ids, hsra, nil)
}

// DeleteHelpdeskSlaReportAnalysis deletes an existing helpdesk.sla.report.analysis record.
func (c *Client) DeleteHelpdeskSlaReportAnalysis(id int64) error {
	return c.DeleteHelpdeskSlaReportAnalysiss([]int64{id})
}

// DeleteHelpdeskSlaReportAnalysiss deletes existing helpdesk.sla.report.analysis records.
func (c *Client) DeleteHelpdeskSlaReportAnalysiss(ids []int64) error {
	return c.Delete(HelpdeskSlaReportAnalysisModel, ids)
}

// GetHelpdeskSlaReportAnalysis gets helpdesk.sla.report.analysis existing record.
func (c *Client) GetHelpdeskSlaReportAnalysis(id int64) (*HelpdeskSlaReportAnalysis, error) {
	hsras, err := c.GetHelpdeskSlaReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hsras)[0]), nil
}

// GetHelpdeskSlaReportAnalysiss gets helpdesk.sla.report.analysis existing records.
func (c *Client) GetHelpdeskSlaReportAnalysiss(ids []int64) (*HelpdeskSlaReportAnalysiss, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.Read(HelpdeskSlaReportAnalysisModel, ids, nil, hsras); err != nil {
		return nil, err
	}
	return hsras, nil
}

// FindHelpdeskSlaReportAnalysis finds helpdesk.sla.report.analysis record by querying it with criteria.
func (c *Client) FindHelpdeskSlaReportAnalysis(criteria *Criteria) (*HelpdeskSlaReportAnalysis, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.SearchRead(HelpdeskSlaReportAnalysisModel, criteria, NewOptions().Limit(1), hsras); err != nil {
		return nil, err
	}
	return &((*hsras)[0]), nil
}

// FindHelpdeskSlaReportAnalysiss finds helpdesk.sla.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaReportAnalysiss(criteria *Criteria, options *Options) (*HelpdeskSlaReportAnalysiss, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.SearchRead(HelpdeskSlaReportAnalysisModel, criteria, options, hsras); err != nil {
		return nil, err
	}
	return hsras, nil
}

// FindHelpdeskSlaReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskSlaReportAnalysisModel, criteria, options)
}

// FindHelpdeskSlaReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskSlaReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskSlaReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
