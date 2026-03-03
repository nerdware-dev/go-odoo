package odoo

// ProductTemplate represents product.template model.
type ProductTemplate struct {
	AccountTagIds                          *Relation   `xmlrpc:"account_tag_ids,omitempty" json:"account_tag_ids,omitempty"`
	Active                                 *Bool       `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId                *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                   *Time       `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration            *Selection  `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                  *String     `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                            *Relation   `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                          *Selection  `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary                        *String     `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon                       *String     `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId                         *Many2One   `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                         *Many2One   `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AttributeLineIds                       *Relation   `xmlrpc:"attribute_line_ids,omitempty" json:"attribute_line_ids,omitempty"`
	Barcode                                *String     `xmlrpc:"barcode,omitempty" json:"barcode,omitempty"`
	CanBeExpensed                          *Bool       `xmlrpc:"can_be_expensed,omitempty" json:"can_be_expensed,omitempty"`
	CanImage1024BeZoomed                   *Bool       `xmlrpc:"can_image_1024_be_zoomed,omitempty" json:"can_image_1024_be_zoomed,omitempty"`
	CategId                                *Many2One   `xmlrpc:"categ_id,omitempty" json:"categ_id,omitempty"`
	Color                                  *Int        `xmlrpc:"color,omitempty" json:"color,omitempty"`
	CompanyId                              *Many2One   `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CostCurrencyId                         *Many2One   `xmlrpc:"cost_currency_id,omitempty" json:"cost_currency_id,omitempty"`
	CostMethod                             *Selection  `xmlrpc:"cost_method,omitempty" json:"cost_method,omitempty"`
	CreateDate                             *Time       `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                              *Many2One   `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	CurrencyId                             *Many2One   `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	DefaultCode                            *String     `xmlrpc:"default_code,omitempty" json:"default_code,omitempty"`
	Description                            *String     `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DescriptionPicking                     *String     `xmlrpc:"description_picking,omitempty" json:"description_picking,omitempty"`
	DescriptionPickingin                   *String     `xmlrpc:"description_pickingin,omitempty" json:"description_pickingin,omitempty"`
	DescriptionPickingout                  *String     `xmlrpc:"description_pickingout,omitempty" json:"description_pickingout,omitempty"`
	DescriptionPurchase                    *String     `xmlrpc:"description_purchase,omitempty" json:"description_purchase,omitempty"`
	DescriptionSale                        *String     `xmlrpc:"description_sale,omitempty" json:"description_sale,omitempty"`
	DetailedType                           *Selection  `xmlrpc:"detailed_type,omitempty" json:"detailed_type,omitempty"`
	DisplayName                            *String     `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentsAllowedCompanyId              *Many2One   `xmlrpc:"documents_allowed_company_id,omitempty" json:"documents_allowed_company_id,omitempty"`
	ExpensePolicy                          *Selection  `xmlrpc:"expense_policy,omitempty" json:"expense_policy,omitempty"`
	ExpensePolicyTooltip                   *String     `xmlrpc:"expense_policy_tooltip,omitempty" json:"expense_policy_tooltip,omitempty"`
	FiscalCountryCodes                     *String     `xmlrpc:"fiscal_country_codes,omitempty" json:"fiscal_country_codes,omitempty"`
	HasAvailableRouteIds                   *Bool       `xmlrpc:"has_available_route_ids,omitempty" json:"has_available_route_ids,omitempty"`
	HasConfigurableAttributes              *Bool       `xmlrpc:"has_configurable_attributes,omitempty" json:"has_configurable_attributes,omitempty"`
	HasMessage                             *Bool       `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                                     *Int        `xmlrpc:"id,omitempty" json:"id,omitempty"`
	Image1024                              *String     `xmlrpc:"image_1024,omitempty" json:"image_1024,omitempty"`
	Image128                               *String     `xmlrpc:"image_128,omitempty" json:"image_128,omitempty"`
	Image1920                              *String     `xmlrpc:"image_1920,omitempty" json:"image_1920,omitempty"`
	Image256                               *String     `xmlrpc:"image_256,omitempty" json:"image_256,omitempty"`
	Image512                               *String     `xmlrpc:"image_512,omitempty" json:"image_512,omitempty"`
	IncomingQty                            *Float      `xmlrpc:"incoming_qty,omitempty" json:"incoming_qty,omitempty"`
	InvoicePolicy                          *Selection  `xmlrpc:"invoice_policy,omitempty" json:"invoice_policy,omitempty"`
	IsProductVariant                       *Bool       `xmlrpc:"is_product_variant,omitempty" json:"is_product_variant,omitempty"`
	ListPrice                              *Float      `xmlrpc:"list_price,omitempty" json:"list_price,omitempty"`
	LocationId                             *Many2One   `xmlrpc:"location_id,omitempty" json:"location_id,omitempty"`
	MessageAttachmentCount                 *Int        `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds                     *Relation   `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError                        *Bool       `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter                 *Int        `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError                     *Bool       `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                             *Relation   `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower                      *Bool       `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction                      *Bool       `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter               *Int        `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds                      *Relation   `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MyActivityDateDeadline                 *Time       `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                                   *String     `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NbrMovesIn                             *Int        `xmlrpc:"nbr_moves_in,omitempty" json:"nbr_moves_in,omitempty"`
	NbrMovesOut                            *Int        `xmlrpc:"nbr_moves_out,omitempty" json:"nbr_moves_out,omitempty"`
	NbrReorderingRules                     *Int        `xmlrpc:"nbr_reordering_rules,omitempty" json:"nbr_reordering_rules,omitempty"`
	OptionalProductIds                     *Relation   `xmlrpc:"optional_product_ids,omitempty" json:"optional_product_ids,omitempty"`
	OutgoingQty                            *Float      `xmlrpc:"outgoing_qty,omitempty" json:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation   `xmlrpc:"packaging_ids,omitempty" json:"packaging_ids,omitempty"`
	PricelistItemCount                     *Int        `xmlrpc:"pricelist_item_count,omitempty" json:"pricelist_item_count,omitempty"`
	Priority                               *Selection  `xmlrpc:"priority,omitempty" json:"priority,omitempty"`
	ProductDocumentCount                   *Int        `xmlrpc:"product_document_count,omitempty" json:"product_document_count,omitempty"`
	ProductDocumentIds                     *Relation   `xmlrpc:"product_document_ids,omitempty" json:"product_document_ids,omitempty"`
	ProductProperties                      interface{} `xmlrpc:"product_properties,omitempty" json:"product_properties,omitempty"`
	ProductSubscriptionPricingIds          *Relation   `xmlrpc:"product_subscription_pricing_ids,omitempty" json:"product_subscription_pricing_ids,omitempty"`
	ProductTagIds                          *Relation   `xmlrpc:"product_tag_ids,omitempty" json:"product_tag_ids,omitempty"`
	ProductTooltip                         *String     `xmlrpc:"product_tooltip,omitempty" json:"product_tooltip,omitempty"`
	ProductVariantCount                    *Int        `xmlrpc:"product_variant_count,omitempty" json:"product_variant_count,omitempty"`
	ProductVariantId                       *Many2One   `xmlrpc:"product_variant_id,omitempty" json:"product_variant_id,omitempty"`
	ProductVariantIds                      *Relation   `xmlrpc:"product_variant_ids,omitempty" json:"product_variant_ids,omitempty"`
	ProjectId                              *Many2One   `xmlrpc:"project_id,omitempty" json:"project_id,omitempty"`
	ProjectTemplateId                      *Many2One   `xmlrpc:"project_template_id,omitempty" json:"project_template_id,omitempty"`
	ProjectTemplateUseDocuments            *Bool       `xmlrpc:"project_template_use_documents,omitempty" json:"project_template_use_documents,omitempty"`
	PropertyAccountCreditorPriceDifference *Many2One   `xmlrpc:"property_account_creditor_price_difference,omitempty" json:"property_account_creditor_price_difference,omitempty"`
	PropertyAccountExpenseId               *Many2One   `xmlrpc:"property_account_expense_id,omitempty" json:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId                *Many2One   `xmlrpc:"property_account_income_id,omitempty" json:"property_account_income_id,omitempty"`
	PropertyStockInventory                 *Many2One   `xmlrpc:"property_stock_inventory,omitempty" json:"property_stock_inventory,omitempty"`
	PropertyStockProduction                *Many2One   `xmlrpc:"property_stock_production,omitempty" json:"property_stock_production,omitempty"`
	PurchaseLineWarn                       *Selection  `xmlrpc:"purchase_line_warn,omitempty" json:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                    *String     `xmlrpc:"purchase_line_warn_msg,omitempty" json:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                         *Selection  `xmlrpc:"purchase_method,omitempty" json:"purchase_method,omitempty"`
	PurchaseOk                             *Bool       `xmlrpc:"purchase_ok,omitempty" json:"purchase_ok,omitempty"`
	PurchasedProductQty                    *Float      `xmlrpc:"purchased_product_qty,omitempty" json:"purchased_product_qty,omitempty"`
	QtyAvailable                           *Float      `xmlrpc:"qty_available,omitempty" json:"qty_available,omitempty"`
	RatingIds                              *Relation   `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RecurringInvoice                       *Bool       `xmlrpc:"recurring_invoice,omitempty" json:"recurring_invoice,omitempty"`
	ReorderingMaxQty                       *Float      `xmlrpc:"reordering_max_qty,omitempty" json:"reordering_max_qty,omitempty"`
	ReorderingMinQty                       *Float      `xmlrpc:"reordering_min_qty,omitempty" json:"reordering_min_qty,omitempty"`
	ResponsibleId                          *Many2One   `xmlrpc:"responsible_id,omitempty" json:"responsible_id,omitempty"`
	RouteFromCategIds                      *Relation   `xmlrpc:"route_from_categ_ids,omitempty" json:"route_from_categ_ids,omitempty"`
	RouteIds                               *Relation   `xmlrpc:"route_ids,omitempty" json:"route_ids,omitempty"`
	SaleDelay                              *Int        `xmlrpc:"sale_delay,omitempty" json:"sale_delay,omitempty"`
	SaleLineWarn                           *Selection  `xmlrpc:"sale_line_warn,omitempty" json:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                        *String     `xmlrpc:"sale_line_warn_msg,omitempty" json:"sale_line_warn_msg,omitempty"`
	SaleOk                                 *Bool       `xmlrpc:"sale_ok,omitempty" json:"sale_ok,omitempty"`
	SalesCount                             *Float      `xmlrpc:"sales_count,omitempty" json:"sales_count,omitempty"`
	SellerIds                              *Relation   `xmlrpc:"seller_ids,omitempty" json:"seller_ids,omitempty"`
	Sequence                               *Int        `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	ServicePolicy                          *Selection  `xmlrpc:"service_policy,omitempty" json:"service_policy,omitempty"`
	ServiceToPurchase                      *Bool       `xmlrpc:"service_to_purchase,omitempty" json:"service_to_purchase,omitempty"`
	ServiceTracking                        *Selection  `xmlrpc:"service_tracking,omitempty" json:"service_tracking,omitempty"`
	ServiceType                            *Selection  `xmlrpc:"service_type,omitempty" json:"service_type,omitempty"`
	ServiceUpsellThreshold                 *Float      `xmlrpc:"service_upsell_threshold,omitempty" json:"service_upsell_threshold,omitempty"`
	ServiceUpsellThresholdRatio            *String     `xmlrpc:"service_upsell_threshold_ratio,omitempty" json:"service_upsell_threshold_ratio,omitempty"`
	ShowForecastedQtyStatusButton          *Bool       `xmlrpc:"show_forecasted_qty_status_button,omitempty" json:"show_forecasted_qty_status_button,omitempty"`
	ShowOnHandQtyStatusButton              *Bool       `xmlrpc:"show_on_hand_qty_status_button,omitempty" json:"show_on_hand_qty_status_button,omitempty"`
	StandardPrice                          *Float      `xmlrpc:"standard_price,omitempty" json:"standard_price,omitempty"`
	SupplierTaxesId                        *Relation   `xmlrpc:"supplier_taxes_id,omitempty" json:"supplier_taxes_id,omitempty"`
	TaxString                              *String     `xmlrpc:"tax_string,omitempty" json:"tax_string,omitempty"`
	TaxesId                                *Relation   `xmlrpc:"taxes_id,omitempty" json:"taxes_id,omitempty"`
	TemplateFolderId                       *Many2One   `xmlrpc:"template_folder_id,omitempty" json:"template_folder_id,omitempty"`
	Tracking                               *Selection  `xmlrpc:"tracking,omitempty" json:"tracking,omitempty"`
	Type                                   *Selection  `xmlrpc:"type,omitempty" json:"type,omitempty"`
	UomId                                  *Many2One   `xmlrpc:"uom_id,omitempty" json:"uom_id,omitempty"`
	UomName                                *String     `xmlrpc:"uom_name,omitempty" json:"uom_name,omitempty"`
	UomPoId                                *Many2One   `xmlrpc:"uom_po_id,omitempty" json:"uom_po_id,omitempty"`
	ValidProductTemplateAttributeLineIds   *Relation   `xmlrpc:"valid_product_template_attribute_line_ids,omitempty" json:"valid_product_template_attribute_line_ids,omitempty"`
	Valuation                              *Selection  `xmlrpc:"valuation,omitempty" json:"valuation,omitempty"`
	VariantSellerIds                       *Relation   `xmlrpc:"variant_seller_ids,omitempty" json:"variant_seller_ids,omitempty"`
	VirtualAvailable                       *Float      `xmlrpc:"virtual_available,omitempty" json:"virtual_available,omitempty"`
	VisibleExpensePolicy                   *Bool       `xmlrpc:"visible_expense_policy,omitempty" json:"visible_expense_policy,omitempty"`
	Volume                                 *Float      `xmlrpc:"volume,omitempty" json:"volume,omitempty"`
	VolumeUomName                          *String     `xmlrpc:"volume_uom_name,omitempty" json:"volume_uom_name,omitempty"`
	WarehouseId                            *Many2One   `xmlrpc:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`
	WebsiteMessageIds                      *Relation   `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	Weight                                 *Float      `xmlrpc:"weight,omitempty" json:"weight,omitempty"`
	WeightUomName                          *String     `xmlrpc:"weight_uom_name,omitempty" json:"weight_uom_name,omitempty"`
	WriteDate                              *Time       `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                               *Many2One   `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// ProductTemplates represents array of product.template model.
type ProductTemplates []ProductTemplate

// ProductTemplateModel is the odoo model name.
const ProductTemplateModel = "product.template"

// Many2One convert ProductTemplate to *Many2One.
func (pt *ProductTemplate) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProductTemplate creates a new product.template model and returns its id.
func (c *Client) CreateProductTemplate(pt *ProductTemplate) (int64, error) {
	ids, err := c.CreateProductTemplates([]*ProductTemplate{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductTemplate creates a new product.template model and returns its id.
func (c *Client) CreateProductTemplates(pts []*ProductTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProductTemplateModel, vv, nil)
}

// UpdateProductTemplate updates an existing product.template record.
func (c *Client) UpdateProductTemplate(pt *ProductTemplate) error {
	return c.UpdateProductTemplates([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTemplates updates existing product.template records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTemplates(ids []int64, pt *ProductTemplate) error {
	return c.Update(ProductTemplateModel, ids, pt, nil)
}

// DeleteProductTemplate deletes an existing product.template record.
func (c *Client) DeleteProductTemplate(id int64) error {
	return c.DeleteProductTemplates([]int64{id})
}

// DeleteProductTemplates deletes existing product.template records.
func (c *Client) DeleteProductTemplates(ids []int64) error {
	return c.Delete(ProductTemplateModel, ids)
}

// GetProductTemplate gets product.template existing record.
func (c *Client) GetProductTemplate(id int64) (*ProductTemplate, error) {
	pts, err := c.GetProductTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*pts) == 0 {
		return nil, nil
	}
	return &((*pts)[0]), nil
}

// GetProductTemplates gets product.template existing records.
func (c *Client) GetProductTemplates(ids []int64) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.Read(ProductTemplateModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplate finds product.template record by querying it with criteria.
func (c *Client) FindProductTemplate(criteria *Criteria) (*ProductTemplate, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if len(*pts) == 0 {
		return nil, nil
	}
	return &((*pts)[0]), nil
}

// FindProductTemplates finds product.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplates(criteria *Criteria, options *Options) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductTemplateModel, criteria, options)
}

// FindProductTemplateId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
