package odoo

// L10NDeReportPartnerVatIntra represents l10n.de.report.partner.vat.intra model.
type L10NDeReportPartnerVatIntra struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// L10NDeReportPartnerVatIntras represents array of l10n.de.report.partner.vat.intra model.
type L10NDeReportPartnerVatIntras []L10NDeReportPartnerVatIntra

// L10NDeReportPartnerVatIntraModel is the odoo model name.
const L10NDeReportPartnerVatIntraModel = "l10n.de.report.partner.vat.intra"

// Many2One convert L10NDeReportPartnerVatIntra to *Many2One.
func (ldrpvi *L10NDeReportPartnerVatIntra) Many2One() *Many2One {
	return NewMany2One(ldrpvi.Id.Get(), "")
}

// CreateL10NDeReportPartnerVatIntra creates a new l10n.de.report.partner.vat.intra model and returns its id.
func (c *Client) CreateL10NDeReportPartnerVatIntra(ldrpvi *L10NDeReportPartnerVatIntra) (int64, error) {
	ids, err := c.CreateL10NDeReportPartnerVatIntras([]*L10NDeReportPartnerVatIntra{ldrpvi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NDeReportPartnerVatIntra creates a new l10n.de.report.partner.vat.intra model and returns its id.
func (c *Client) CreateL10NDeReportPartnerVatIntras(ldrpvis []*L10NDeReportPartnerVatIntra) ([]int64, error) {
	var vv []interface{}
	for _, v := range ldrpvis {
		vv = append(vv, v)
	}
	return c.Create(L10NDeReportPartnerVatIntraModel, vv, nil)
}

// UpdateL10NDeReportPartnerVatIntra updates an existing l10n.de.report.partner.vat.intra record.
func (c *Client) UpdateL10NDeReportPartnerVatIntra(ldrpvi *L10NDeReportPartnerVatIntra) error {
	return c.UpdateL10NDeReportPartnerVatIntras([]int64{ldrpvi.Id.Get()}, ldrpvi)
}

// UpdateL10NDeReportPartnerVatIntras updates existing l10n.de.report.partner.vat.intra records.
// All records (represented by ids) will be updated by ldrpvi values.
func (c *Client) UpdateL10NDeReportPartnerVatIntras(ids []int64, ldrpvi *L10NDeReportPartnerVatIntra) error {
	return c.Update(L10NDeReportPartnerVatIntraModel, ids, ldrpvi, nil)
}

// DeleteL10NDeReportPartnerVatIntra deletes an existing l10n.de.report.partner.vat.intra record.
func (c *Client) DeleteL10NDeReportPartnerVatIntra(id int64) error {
	return c.DeleteL10NDeReportPartnerVatIntras([]int64{id})
}

// DeleteL10NDeReportPartnerVatIntras deletes existing l10n.de.report.partner.vat.intra records.
func (c *Client) DeleteL10NDeReportPartnerVatIntras(ids []int64) error {
	return c.Delete(L10NDeReportPartnerVatIntraModel, ids)
}

// GetL10NDeReportPartnerVatIntra gets l10n.de.report.partner.vat.intra existing record.
func (c *Client) GetL10NDeReportPartnerVatIntra(id int64) (*L10NDeReportPartnerVatIntra, error) {
	ldrpvis, err := c.GetL10NDeReportPartnerVatIntras([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ldrpvis)[0]), nil
}

// GetL10NDeReportPartnerVatIntras gets l10n.de.report.partner.vat.intra existing records.
func (c *Client) GetL10NDeReportPartnerVatIntras(ids []int64) (*L10NDeReportPartnerVatIntras, error) {
	ldrpvis := &L10NDeReportPartnerVatIntras{}
	if err := c.Read(L10NDeReportPartnerVatIntraModel, ids, nil, ldrpvis); err != nil {
		return nil, err
	}
	return ldrpvis, nil
}

// FindL10NDeReportPartnerVatIntra finds l10n.de.report.partner.vat.intra record by querying it with criteria.
func (c *Client) FindL10NDeReportPartnerVatIntra(criteria *Criteria) (*L10NDeReportPartnerVatIntra, error) {
	ldrpvis := &L10NDeReportPartnerVatIntras{}
	if err := c.SearchRead(L10NDeReportPartnerVatIntraModel, criteria, NewOptions().Limit(1), ldrpvis); err != nil {
		return nil, err
	}
	return &((*ldrpvis)[0]), nil
}

// FindL10NDeReportPartnerVatIntras finds l10n.de.report.partner.vat.intra records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NDeReportPartnerVatIntras(criteria *Criteria, options *Options) (*L10NDeReportPartnerVatIntras, error) {
	ldrpvis := &L10NDeReportPartnerVatIntras{}
	if err := c.SearchRead(L10NDeReportPartnerVatIntraModel, criteria, options, ldrpvis); err != nil {
		return nil, err
	}
	return ldrpvis, nil
}

// FindL10NDeReportPartnerVatIntraIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NDeReportPartnerVatIntraIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NDeReportPartnerVatIntraModel, criteria, options)
}

// FindL10NDeReportPartnerVatIntraId finds record id by querying it with criteria.
func (c *Client) FindL10NDeReportPartnerVatIntraId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NDeReportPartnerVatIntraModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
