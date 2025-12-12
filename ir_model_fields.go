package odoo

// IrModelFields represents ir.model.fields model.
type IrModelFields struct {
	Column1             *String    `xmlrpc:"column1,omitempty" json:"column1,omitempty"`
	Column2             *String    `xmlrpc:"column2,omitempty" json:"column2,omitempty"`
	CompanyDependent    *Bool      `xmlrpc:"company_dependent,omitempty" json:"company_dependent,omitempty"`
	CompleteName        *String    `xmlrpc:"complete_name,omitempty" json:"complete_name,omitempty"`
	Compute             *String    `xmlrpc:"compute,omitempty" json:"compute,omitempty"`
	Copied              *Bool      `xmlrpc:"copied,omitempty" json:"copied,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyField       *String    `xmlrpc:"currency_field,omitempty" json:"currency_field,omitempty"`
	Depends             *String    `xmlrpc:"depends,omitempty" json:"depends,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	Domain              *String    `xmlrpc:"domain,omitempty" json:"domain,omitempty"`
	FieldDescription    *String    `xmlrpc:"field_description,omitempty" json:"field_description,omitempty"`
	GroupExpand         *Bool      `xmlrpc:"group_expand,omitempty" json:"group_expand,omitempty"`
	Groups              *Relation  `xmlrpc:"groups,omitempty" json:"groups,omitempty"`
	Help                *String    `xmlrpc:"help,omitempty" json:"help,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Index               *Bool      `xmlrpc:"index,omitempty" json:"index,omitempty"`
	Model               *String    `xmlrpc:"model,omitempty" json:"model,omitempty"`
	ModelId             *Many2One  `xmlrpc:"model_id,omitempty" json:"model_id,omitempty"`
	Modules             *String    `xmlrpc:"modules,omitempty" json:"modules,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OnDelete            *Selection `xmlrpc:"on_delete,omitempty" json:"on_delete,omitempty"`
	Readonly            *Bool      `xmlrpc:"readonly,omitempty" json:"readonly,omitempty"`
	Related             *String    `xmlrpc:"related,omitempty" json:"related,omitempty"`
	RelatedFieldId      *Many2One  `xmlrpc:"related_field_id,omitempty" json:"related_field_id,omitempty"`
	Relation            *String    `xmlrpc:"relation,omitempty" json:"relation,omitempty"`
	RelationField       *String    `xmlrpc:"relation_field,omitempty" json:"relation_field,omitempty"`
	RelationFieldId     *Many2One  `xmlrpc:"relation_field_id,omitempty" json:"relation_field_id,omitempty"`
	RelationTable       *String    `xmlrpc:"relation_table,omitempty" json:"relation_table,omitempty"`
	Required            *Bool      `xmlrpc:"required,omitempty" json:"required,omitempty"`
	Sanitize            *Bool      `xmlrpc:"sanitize,omitempty" json:"sanitize,omitempty"`
	SanitizeAttributes  *Bool      `xmlrpc:"sanitize_attributes,omitempty" json:"sanitize_attributes,omitempty"`
	SanitizeForm        *Bool      `xmlrpc:"sanitize_form,omitempty" json:"sanitize_form,omitempty"`
	SanitizeOverridable *Bool      `xmlrpc:"sanitize_overridable,omitempty" json:"sanitize_overridable,omitempty"`
	SanitizeStyle       *Bool      `xmlrpc:"sanitize_style,omitempty" json:"sanitize_style,omitempty"`
	SanitizeTags        *Bool      `xmlrpc:"sanitize_tags,omitempty" json:"sanitize_tags,omitempty"`
	Selectable          *Bool      `xmlrpc:"selectable,omitempty" json:"selectable,omitempty"`
	Selection           *String    `xmlrpc:"selection,omitempty" json:"selection,omitempty"`
	SelectionIds        *Relation  `xmlrpc:"selection_ids,omitempty" json:"selection_ids,omitempty"`
	Size                *Int       `xmlrpc:"size,omitempty" json:"size,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	Store               *Bool      `xmlrpc:"store,omitempty" json:"store,omitempty"`
	StripClasses        *Bool      `xmlrpc:"strip_classes,omitempty" json:"strip_classes,omitempty"`
	StripStyle          *Bool      `xmlrpc:"strip_style,omitempty" json:"strip_style,omitempty"`
	Tracking            *Int       `xmlrpc:"tracking,omitempty" json:"tracking,omitempty"`
	Translate           *Bool      `xmlrpc:"translate,omitempty" json:"translate,omitempty"`
	Ttype               *Selection `xmlrpc:"ttype,omitempty" json:"ttype,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// IrModelFieldss represents array of ir.model.fields model.
type IrModelFieldss []IrModelFields

// IrModelFieldsModel is the odoo model name.
const IrModelFieldsModel = "ir.model.fields"

// Many2One convert IrModelFields to *Many2One.
func (imf *IrModelFields) Many2One() *Many2One {
	return NewMany2One(imf.Id.Get(), "")
}

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFields(imf *IrModelFields) (int64, error) {
	ids, err := c.CreateIrModelFieldss([]*IrModelFields{imf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFieldss(imfs []*IrModelFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range imfs {
		vv = append(vv, v)
	}
	return c.Create(IrModelFieldsModel, vv, nil)
}

// UpdateIrModelFields updates an existing ir.model.fields record.
func (c *Client) UpdateIrModelFields(imf *IrModelFields) error {
	return c.UpdateIrModelFieldss([]int64{imf.Id.Get()}, imf)
}

// UpdateIrModelFieldss updates existing ir.model.fields records.
// All records (represented by ids) will be updated by imf values.
func (c *Client) UpdateIrModelFieldss(ids []int64, imf *IrModelFields) error {
	return c.Update(IrModelFieldsModel, ids, imf, nil)
}

// DeleteIrModelFields deletes an existing ir.model.fields record.
func (c *Client) DeleteIrModelFields(id int64) error {
	return c.DeleteIrModelFieldss([]int64{id})
}

// DeleteIrModelFieldss deletes existing ir.model.fields records.
func (c *Client) DeleteIrModelFieldss(ids []int64) error {
	return c.Delete(IrModelFieldsModel, ids)
}

// GetIrModelFields gets ir.model.fields existing record.
func (c *Client) GetIrModelFields(id int64) (*IrModelFields, error) {
	imfs, err := c.GetIrModelFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imfs)[0]), nil
}

// GetIrModelFieldss gets ir.model.fields existing records.
func (c *Client) GetIrModelFieldss(ids []int64) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.Read(IrModelFieldsModel, ids, nil, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFields finds ir.model.fields record by querying it with criteria.
func (c *Client) FindIrModelFields(criteria *Criteria) (*IrModelFields, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, NewOptions().Limit(1), imfs); err != nil {
		return nil, err
	}
	return &((*imfs)[0]), nil
}

// FindIrModelFieldss finds ir.model.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldss(criteria *Criteria, options *Options) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, options, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelFieldsModel, criteria, options)
}

// FindIrModelFieldsId finds record id by querying it with criteria.
func (c *Client) FindIrModelFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
