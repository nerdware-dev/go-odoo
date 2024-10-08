package odoo

// TungstenExporterWizardLine represents tungsten.exporter.wizard.line model.
type TungstenExporterWizardLine struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	Data               *String   `xmlrpc:"data,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	ExceptionMsg       *String   `xmlrpc:"exception_msg,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	SourceDocumentName *String   `xmlrpc:"source_document_name,omitempty"`
	TungstenExporterId *Many2One `xmlrpc:"tungsten_exporter_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// TungstenExporterWizardLines represents array of tungsten.exporter.wizard.line model.
type TungstenExporterWizardLines []TungstenExporterWizardLine

// TungstenExporterWizardLineModel is the odoo model name.
const TungstenExporterWizardLineModel = "tungsten.exporter.wizard.line"

// Many2One convert TungstenExporterWizardLine to *Many2One.
func (tewl *TungstenExporterWizardLine) Many2One() *Many2One {
	return NewMany2One(tewl.Id.Get(), "")
}

// CreateTungstenExporterWizardLine creates a new tungsten.exporter.wizard.line model and returns its id.
func (c *Client) CreateTungstenExporterWizardLine(tewl *TungstenExporterWizardLine) (int64, error) {
	ids, err := c.CreateTungstenExporterWizardLines([]*TungstenExporterWizardLine{tewl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTungstenExporterWizardLine creates a new tungsten.exporter.wizard.line model and returns its id.
func (c *Client) CreateTungstenExporterWizardLines(tewls []*TungstenExporterWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range tewls {
		vv = append(vv, v)
	}
	return c.Create(TungstenExporterWizardLineModel, vv, nil)
}

// UpdateTungstenExporterWizardLine updates an existing tungsten.exporter.wizard.line record.
func (c *Client) UpdateTungstenExporterWizardLine(tewl *TungstenExporterWizardLine) error {
	return c.UpdateTungstenExporterWizardLines([]int64{tewl.Id.Get()}, tewl)
}

// UpdateTungstenExporterWizardLines updates existing tungsten.exporter.wizard.line records.
// All records (represented by ids) will be updated by tewl values.
func (c *Client) UpdateTungstenExporterWizardLines(ids []int64, tewl *TungstenExporterWizardLine) error {
	return c.Update(TungstenExporterWizardLineModel, ids, tewl, nil)
}

// DeleteTungstenExporterWizardLine deletes an existing tungsten.exporter.wizard.line record.
func (c *Client) DeleteTungstenExporterWizardLine(id int64) error {
	return c.DeleteTungstenExporterWizardLines([]int64{id})
}

// DeleteTungstenExporterWizardLines deletes existing tungsten.exporter.wizard.line records.
func (c *Client) DeleteTungstenExporterWizardLines(ids []int64) error {
	return c.Delete(TungstenExporterWizardLineModel, ids)
}

// GetTungstenExporterWizardLine gets tungsten.exporter.wizard.line existing record.
func (c *Client) GetTungstenExporterWizardLine(id int64) (*TungstenExporterWizardLine, error) {
	tewls, err := c.GetTungstenExporterWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tewls)[0]), nil
}

// GetTungstenExporterWizardLines gets tungsten.exporter.wizard.line existing records.
func (c *Client) GetTungstenExporterWizardLines(ids []int64) (*TungstenExporterWizardLines, error) {
	tewls := &TungstenExporterWizardLines{}
	if err := c.Read(TungstenExporterWizardLineModel, ids, nil, tewls); err != nil {
		return nil, err
	}
	return tewls, nil
}

// FindTungstenExporterWizardLine finds tungsten.exporter.wizard.line record by querying it with criteria.
func (c *Client) FindTungstenExporterWizardLine(criteria *Criteria) (*TungstenExporterWizardLine, error) {
	tewls := &TungstenExporterWizardLines{}
	if err := c.SearchRead(TungstenExporterWizardLineModel, criteria, NewOptions().Limit(1), tewls); err != nil {
		return nil, err
	}
	return &((*tewls)[0]), nil
}

// FindTungstenExporterWizardLines finds tungsten.exporter.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTungstenExporterWizardLines(criteria *Criteria, options *Options) (*TungstenExporterWizardLines, error) {
	tewls := &TungstenExporterWizardLines{}
	if err := c.SearchRead(TungstenExporterWizardLineModel, criteria, options, tewls); err != nil {
		return nil, err
	}
	return tewls, nil
}

// FindTungstenExporterWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTungstenExporterWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TungstenExporterWizardLineModel, criteria, options)
}

// FindTungstenExporterWizardLineId finds record id by querying it with criteria.
func (c *Client) FindTungstenExporterWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TungstenExporterWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
