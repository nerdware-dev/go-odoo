package odoo

import (
	"fmt"
)

// TungstenExporterWizardLine represents tungsten.exporter.wizard.line model.
type TungstenExporterWizardLine struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	Data               *String   `xmlrpc:"data,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	ExceptionMsg       *String   `xmlrpc:"exception_msg,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	Name               *String   `xmlrpc:"name,omptempty"`
	SourceDocumentName *String   `xmlrpc:"source_document_name,omptempty"`
	TungstenExporterId *Many2One `xmlrpc:"tungsten_exporter_id,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(TungstenExporterWizardLineModel, vv)
}

// UpdateTungstenExporterWizardLine updates an existing tungsten.exporter.wizard.line record.
func (c *Client) UpdateTungstenExporterWizardLine(tewl *TungstenExporterWizardLine) error {
	return c.UpdateTungstenExporterWizardLines([]int64{tewl.Id.Get()}, tewl)
}

// UpdateTungstenExporterWizardLines updates existing tungsten.exporter.wizard.line records.
// All records (represented by ids) will be updated by tewl values.
func (c *Client) UpdateTungstenExporterWizardLines(ids []int64, tewl *TungstenExporterWizardLine) error {
	return c.Update(TungstenExporterWizardLineModel, ids, tewl)
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
	if tewls != nil && len(*tewls) > 0 {
		return &((*tewls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of tungsten.exporter.wizard.line not found", id)
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
	if tewls != nil && len(*tewls) > 0 {
		return &((*tewls)[0]), nil
	}
	return nil, fmt.Errorf("tungsten.exporter.wizard.line was not found with criteria %v", criteria)
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
	ids, err := c.Search(TungstenExporterWizardLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTungstenExporterWizardLineId finds record id by querying it with criteria.
func (c *Client) FindTungstenExporterWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TungstenExporterWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("tungsten.exporter.wizard.line was not found with criteria %v and options %v", criteria, options)
}
