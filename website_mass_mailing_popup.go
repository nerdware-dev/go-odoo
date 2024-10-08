package odoo

// WebsiteMassMailingPopup represents website.mass_mailing.popup model.
type WebsiteMassMailingPopup struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	MailingListId *Many2One `xmlrpc:"mailing_list_id,omitempty"`
	PopupContent  *String   `xmlrpc:"popup_content,omitempty"`
	WebsiteId     *Many2One `xmlrpc:"website_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteMassMailingPopups represents array of website.mass_mailing.popup model.
type WebsiteMassMailingPopups []WebsiteMassMailingPopup

// WebsiteMassMailingPopupModel is the odoo model name.
const WebsiteMassMailingPopupModel = "website.mass_mailing.popup"

// Many2One convert WebsiteMassMailingPopup to *Many2One.
func (wmp *WebsiteMassMailingPopup) Many2One() *Many2One {
	return NewMany2One(wmp.Id.Get(), "")
}

// CreateWebsiteMassMailingPopup creates a new website.mass_mailing.popup model and returns its id.
func (c *Client) CreateWebsiteMassMailingPopup(wmp *WebsiteMassMailingPopup) (int64, error) {
	ids, err := c.CreateWebsiteMassMailingPopups([]*WebsiteMassMailingPopup{wmp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteMassMailingPopup creates a new website.mass_mailing.popup model and returns its id.
func (c *Client) CreateWebsiteMassMailingPopups(wmps []*WebsiteMassMailingPopup) ([]int64, error) {
	var vv []interface{}
	for _, v := range wmps {
		vv = append(vv, v)
	}
	return c.Create(WebsiteMassMailingPopupModel, vv, nil)
}

// UpdateWebsiteMassMailingPopup updates an existing website.mass_mailing.popup record.
func (c *Client) UpdateWebsiteMassMailingPopup(wmp *WebsiteMassMailingPopup) error {
	return c.UpdateWebsiteMassMailingPopups([]int64{wmp.Id.Get()}, wmp)
}

// UpdateWebsiteMassMailingPopups updates existing website.mass_mailing.popup records.
// All records (represented by ids) will be updated by wmp values.
func (c *Client) UpdateWebsiteMassMailingPopups(ids []int64, wmp *WebsiteMassMailingPopup) error {
	return c.Update(WebsiteMassMailingPopupModel, ids, wmp, nil)
}

// DeleteWebsiteMassMailingPopup deletes an existing website.mass_mailing.popup record.
func (c *Client) DeleteWebsiteMassMailingPopup(id int64) error {
	return c.DeleteWebsiteMassMailingPopups([]int64{id})
}

// DeleteWebsiteMassMailingPopups deletes existing website.mass_mailing.popup records.
func (c *Client) DeleteWebsiteMassMailingPopups(ids []int64) error {
	return c.Delete(WebsiteMassMailingPopupModel, ids)
}

// GetWebsiteMassMailingPopup gets website.mass_mailing.popup existing record.
func (c *Client) GetWebsiteMassMailingPopup(id int64) (*WebsiteMassMailingPopup, error) {
	wmps, err := c.GetWebsiteMassMailingPopups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wmps)[0]), nil
}

// GetWebsiteMassMailingPopups gets website.mass_mailing.popup existing records.
func (c *Client) GetWebsiteMassMailingPopups(ids []int64) (*WebsiteMassMailingPopups, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.Read(WebsiteMassMailingPopupModel, ids, nil, wmps); err != nil {
		return nil, err
	}
	return wmps, nil
}

// FindWebsiteMassMailingPopup finds website.mass_mailing.popup record by querying it with criteria.
func (c *Client) FindWebsiteMassMailingPopup(criteria *Criteria) (*WebsiteMassMailingPopup, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.SearchRead(WebsiteMassMailingPopupModel, criteria, NewOptions().Limit(1), wmps); err != nil {
		return nil, err
	}
	return &((*wmps)[0]), nil
}

// FindWebsiteMassMailingPopups finds website.mass_mailing.popup records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMassMailingPopups(criteria *Criteria, options *Options) (*WebsiteMassMailingPopups, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.SearchRead(WebsiteMassMailingPopupModel, criteria, options, wmps); err != nil {
		return nil, err
	}
	return wmps, nil
}

// FindWebsiteMassMailingPopupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMassMailingPopupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteMassMailingPopupModel, criteria, options)
}

// FindWebsiteMassMailingPopupId finds record id by querying it with criteria.
func (c *Client) FindWebsiteMassMailingPopupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteMassMailingPopupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
