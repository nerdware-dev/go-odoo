package odoo

import (
	"fmt"
)

// TungstenExporterWizard represents tungsten.exporter.wizard model.
type TungstenExporterWizard struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	TungstenWizardLineIds *Relation `xmlrpc:"tungsten_wizard_line_ids,omptempty"`
	TungstenZipBinary     *String   `xmlrpc:"tungsten_zip_binary,omptempty"`
	TungstenZipFilename   *String   `xmlrpc:"tungsten_zip_filename,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// TungstenExporterWizards represents array of tungsten.exporter.wizard model.
type TungstenExporterWizards []TungstenExporterWizard

// TungstenExporterWizardModel is the odoo model name.
const TungstenExporterWizardModel = "tungsten.exporter.wizard"

// Many2One convert TungstenExporterWizard to *Many2One.
func (tew *TungstenExporterWizard) Many2One() *Many2One {
	return NewMany2One(tew.Id.Get(), "")
}

// CreateTungstenExporterWizard creates a new tungsten.exporter.wizard model and returns its id.
func (c *Client) CreateTungstenExporterWizard(tew *TungstenExporterWizard) (int64, error) {
	ids, err := c.CreateTungstenExporterWizards([]*TungstenExporterWizard{tew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTungstenExporterWizard creates a new tungsten.exporter.wizard model and returns its id.
func (c *Client) CreateTungstenExporterWizards(tews []*TungstenExporterWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range tews {
		vv = append(vv, v)
	}
	return c.Create(TungstenExporterWizardModel, vv)
}

// UpdateTungstenExporterWizard updates an existing tungsten.exporter.wizard record.
func (c *Client) UpdateTungstenExporterWizard(tew *TungstenExporterWizard) error {
	return c.UpdateTungstenExporterWizards([]int64{tew.Id.Get()}, tew)
}

// UpdateTungstenExporterWizards updates existing tungsten.exporter.wizard records.
// All records (represented by ids) will be updated by tew values.
func (c *Client) UpdateTungstenExporterWizards(ids []int64, tew *TungstenExporterWizard) error {
	return c.Update(TungstenExporterWizardModel, ids, tew)
}

// DeleteTungstenExporterWizard deletes an existing tungsten.exporter.wizard record.
func (c *Client) DeleteTungstenExporterWizard(id int64) error {
	return c.DeleteTungstenExporterWizards([]int64{id})
}

// DeleteTungstenExporterWizards deletes existing tungsten.exporter.wizard records.
func (c *Client) DeleteTungstenExporterWizards(ids []int64) error {
	return c.Delete(TungstenExporterWizardModel, ids)
}

// GetTungstenExporterWizard gets tungsten.exporter.wizard existing record.
func (c *Client) GetTungstenExporterWizard(id int64) (*TungstenExporterWizard, error) {
	tews, err := c.GetTungstenExporterWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if tews != nil && len(*tews) > 0 {
		return &((*tews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of tungsten.exporter.wizard not found", id)
}

// GetTungstenExporterWizards gets tungsten.exporter.wizard existing records.
func (c *Client) GetTungstenExporterWizards(ids []int64) (*TungstenExporterWizards, error) {
	tews := &TungstenExporterWizards{}
	if err := c.Read(TungstenExporterWizardModel, ids, nil, tews); err != nil {
		return nil, err
	}
	return tews, nil
}

// FindTungstenExporterWizard finds tungsten.exporter.wizard record by querying it with criteria.
func (c *Client) FindTungstenExporterWizard(criteria *Criteria) (*TungstenExporterWizard, error) {
	tews := &TungstenExporterWizards{}
	if err := c.SearchRead(TungstenExporterWizardModel, criteria, NewOptions().Limit(1), tews); err != nil {
		return nil, err
	}
	if tews != nil && len(*tews) > 0 {
		return &((*tews)[0]), nil
	}
	return nil, fmt.Errorf("tungsten.exporter.wizard was not found with criteria %v", criteria)
}

// FindTungstenExporterWizards finds tungsten.exporter.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTungstenExporterWizards(criteria *Criteria, options *Options) (*TungstenExporterWizards, error) {
	tews := &TungstenExporterWizards{}
	if err := c.SearchRead(TungstenExporterWizardModel, criteria, options, tews); err != nil {
		return nil, err
	}
	return tews, nil
}

// FindTungstenExporterWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTungstenExporterWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TungstenExporterWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTungstenExporterWizardId finds record id by querying it with criteria.
func (c *Client) FindTungstenExporterWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TungstenExporterWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("tungsten.exporter.wizard was not found with criteria %v and options %v", criteria, options)
}
