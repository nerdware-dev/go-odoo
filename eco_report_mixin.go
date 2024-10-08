package odoo

// EcoReportMixin represents eco_report.mixin model.
type EcoReportMixin struct {
	LastUpdate        *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName       *String `xmlrpc:"display_name,omitempty"`
	Id                *Int    `xmlrpc:"id,omitempty"`
	ReportComputeDate *Time   `xmlrpc:"report_compute_date,omitempty"`
	ReportLineIndex   *Int    `xmlrpc:"report_line_index,omitempty"`
}

// EcoReportMixins represents array of eco_report.mixin model.
type EcoReportMixins []EcoReportMixin

// EcoReportMixinModel is the odoo model name.
const EcoReportMixinModel = "eco_report.mixin"

// Many2One convert EcoReportMixin to *Many2One.
func (em *EcoReportMixin) Many2One() *Many2One {
	return NewMany2One(em.Id.Get(), "")
}

// CreateEcoReportMixin creates a new eco_report.mixin model and returns its id.
func (c *Client) CreateEcoReportMixin(em *EcoReportMixin) (int64, error) {
	ids, err := c.CreateEcoReportMixins([]*EcoReportMixin{em})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEcoReportMixin creates a new eco_report.mixin model and returns its id.
func (c *Client) CreateEcoReportMixins(ems []*EcoReportMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ems {
		vv = append(vv, v)
	}
	return c.Create(EcoReportMixinModel, vv, nil)
}

// UpdateEcoReportMixin updates an existing eco_report.mixin record.
func (c *Client) UpdateEcoReportMixin(em *EcoReportMixin) error {
	return c.UpdateEcoReportMixins([]int64{em.Id.Get()}, em)
}

// UpdateEcoReportMixins updates existing eco_report.mixin records.
// All records (represented by ids) will be updated by em values.
func (c *Client) UpdateEcoReportMixins(ids []int64, em *EcoReportMixin) error {
	return c.Update(EcoReportMixinModel, ids, em, nil)
}

// DeleteEcoReportMixin deletes an existing eco_report.mixin record.
func (c *Client) DeleteEcoReportMixin(id int64) error {
	return c.DeleteEcoReportMixins([]int64{id})
}

// DeleteEcoReportMixins deletes existing eco_report.mixin records.
func (c *Client) DeleteEcoReportMixins(ids []int64) error {
	return c.Delete(EcoReportMixinModel, ids)
}

// GetEcoReportMixin gets eco_report.mixin existing record.
func (c *Client) GetEcoReportMixin(id int64) (*EcoReportMixin, error) {
	ems, err := c.GetEcoReportMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// GetEcoReportMixins gets eco_report.mixin existing records.
func (c *Client) GetEcoReportMixins(ids []int64) (*EcoReportMixins, error) {
	ems := &EcoReportMixins{}
	if err := c.Read(EcoReportMixinModel, ids, nil, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEcoReportMixin finds eco_report.mixin record by querying it with criteria.
func (c *Client) FindEcoReportMixin(criteria *Criteria) (*EcoReportMixin, error) {
	ems := &EcoReportMixins{}
	if err := c.SearchRead(EcoReportMixinModel, criteria, NewOptions().Limit(1), ems); err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// FindEcoReportMixins finds eco_report.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEcoReportMixins(criteria *Criteria, options *Options) (*EcoReportMixins, error) {
	ems := &EcoReportMixins{}
	if err := c.SearchRead(EcoReportMixinModel, criteria, options, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEcoReportMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEcoReportMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EcoReportMixinModel, criteria, options)
}

// FindEcoReportMixinId finds record id by querying it with criteria.
func (c *Client) FindEcoReportMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EcoReportMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
