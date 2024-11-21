package odoo

// IrUiView represents ir.ui.view model.
type IrUiView struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Active                 *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	Arch                   *String    `xmlrpc:"arch,omitempty" json:"arch,omitempty"`
	ArchBase               *String    `xmlrpc:"arch_base,omitempty" json:"arch_base,omitempty"`
	ArchDb                 *String    `xmlrpc:"arch_db,omitempty" json:"arch_db,omitempty"`
	ArchFs                 *String    `xmlrpc:"arch_fs,omitempty" json:"arch_fs,omitempty"`
	ArchPrev               *String    `xmlrpc:"arch_prev,omitempty" json:"arch_prev,omitempty"`
	ArchUpdated            *Bool      `xmlrpc:"arch_updated,omitempty" json:"arch_updated,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CustomizeShow          *Bool      `xmlrpc:"customize_show,omitempty" json:"customize_show,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	FieldParent            *String    `xmlrpc:"field_parent,omitempty" json:"field_parent,omitempty"`
	FirstPageId            *Many2One  `xmlrpc:"first_page_id,omitempty" json:"first_page_id,omitempty"`
	GroupsId               *Relation  `xmlrpc:"groups_id,omitempty" json:"groups_id,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	InheritChildrenIds     *Relation  `xmlrpc:"inherit_children_ids,omitempty" json:"inherit_children_ids,omitempty"`
	InheritId              *Many2One  `xmlrpc:"inherit_id,omitempty" json:"inherit_id,omitempty"`
	IsSeoOptimized         *Bool      `xmlrpc:"is_seo_optimized,omitempty" json:"is_seo_optimized,omitempty"`
	Key                    *String    `xmlrpc:"key,omitempty" json:"key,omitempty"`
	Mode                   *Selection `xmlrpc:"mode,omitempty" json:"mode,omitempty"`
	Model                  *String    `xmlrpc:"model,omitempty" json:"model,omitempty"`
	ModelDataId            *Many2One  `xmlrpc:"model_data_id,omitempty" json:"model_data_id,omitempty"`
	ModelIds               *Relation  `xmlrpc:"model_ids,omitempty" json:"model_ids,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	PageIds                *Relation  `xmlrpc:"page_ids,omitempty" json:"page_ids,omitempty"`
	Priority               *Int       `xmlrpc:"priority,omitempty" json:"priority,omitempty"`
	ThemeTemplateId        *Many2One  `xmlrpc:"theme_template_id,omitempty" json:"theme_template_id,omitempty"`
	Track                  *Bool      `xmlrpc:"track,omitempty" json:"track,omitempty"`
	Type                   *Selection `xmlrpc:"type,omitempty" json:"type,omitempty"`
	WebsiteId              *Many2One  `xmlrpc:"website_id,omitempty" json:"website_id,omitempty"`
	WebsiteMetaDescription *String    `xmlrpc:"website_meta_description,omitempty" json:"website_meta_description,omitempty"`
	WebsiteMetaKeywords    *String    `xmlrpc:"website_meta_keywords,omitempty" json:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg       *String    `xmlrpc:"website_meta_og_img,omitempty" json:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle       *String    `xmlrpc:"website_meta_title,omitempty" json:"website_meta_title,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
	XmlId                  *String    `xmlrpc:"xml_id,omitempty" json:"xml_id,omitempty"`
}

// IrUiViews represents array of ir.ui.view model.
type IrUiViews []IrUiView

// IrUiViewModel is the odoo model name.
const IrUiViewModel = "ir.ui.view"

// Many2One convert IrUiView to *Many2One.
func (iuv *IrUiView) Many2One() *Many2One {
	return NewMany2One(iuv.Id.Get(), "")
}

// CreateIrUiView creates a new ir.ui.view model and returns its id.
func (c *Client) CreateIrUiView(iuv *IrUiView) (int64, error) {
	ids, err := c.CreateIrUiViews([]*IrUiView{iuv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrUiView creates a new ir.ui.view model and returns its id.
func (c *Client) CreateIrUiViews(iuvs []*IrUiView) ([]int64, error) {
	var vv []interface{}
	for _, v := range iuvs {
		vv = append(vv, v)
	}
	return c.Create(IrUiViewModel, vv, nil)
}

// UpdateIrUiView updates an existing ir.ui.view record.
func (c *Client) UpdateIrUiView(iuv *IrUiView) error {
	return c.UpdateIrUiViews([]int64{iuv.Id.Get()}, iuv)
}

// UpdateIrUiViews updates existing ir.ui.view records.
// All records (represented by ids) will be updated by iuv values.
func (c *Client) UpdateIrUiViews(ids []int64, iuv *IrUiView) error {
	return c.Update(IrUiViewModel, ids, iuv, nil)
}

// DeleteIrUiView deletes an existing ir.ui.view record.
func (c *Client) DeleteIrUiView(id int64) error {
	return c.DeleteIrUiViews([]int64{id})
}

// DeleteIrUiViews deletes existing ir.ui.view records.
func (c *Client) DeleteIrUiViews(ids []int64) error {
	return c.Delete(IrUiViewModel, ids)
}

// GetIrUiView gets ir.ui.view existing record.
func (c *Client) GetIrUiView(id int64) (*IrUiView, error) {
	iuvs, err := c.GetIrUiViews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iuvs)[0]), nil
}

// GetIrUiViews gets ir.ui.view existing records.
func (c *Client) GetIrUiViews(ids []int64) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.Read(IrUiViewModel, ids, nil, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiView finds ir.ui.view record by querying it with criteria.
func (c *Client) FindIrUiView(criteria *Criteria) (*IrUiView, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, NewOptions().Limit(1), iuvs); err != nil {
		return nil, err
	}
	return &((*iuvs)[0]), nil
}

// FindIrUiViews finds ir.ui.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViews(criteria *Criteria, options *Options) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, options, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrUiViewModel, criteria, options)
}

// FindIrUiViewId finds record id by querying it with criteria.
func (c *Client) FindIrUiViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrUiViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
