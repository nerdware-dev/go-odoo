package odoo

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                             *Time      `xmlrpc:"__last_update,omitempty"`
	Active                                 *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline                   *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration            *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                  *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                            *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                          *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                        *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                         *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                         *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttributeLineIds                       *Relation  `xmlrpc:"attribute_line_ids,omitempty"`
	Barcode                                *String    `xmlrpc:"barcode,omitempty"`
	CanBeExpensed                          *Bool      `xmlrpc:"can_be_expensed,omitempty"`
	CanImage1024BeZoomed                   *Bool      `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CanImageVariant1024BeZoomed            *Bool      `xmlrpc:"can_image_variant_1024_be_zoomed,omitempty"`
	CanPublish                             *Bool      `xmlrpc:"can_publish,omitempty"`
	CategId                                *Many2One  `xmlrpc:"categ_id,omitempty"`
	Code                                   *String    `xmlrpc:"code,omitempty"`
	Color                                  *Int       `xmlrpc:"color,omitempty"`
	CombinationIndices                     *String    `xmlrpc:"combination_indices,omitempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omitempty"`
	CostCurrencyId                         *Many2One  `xmlrpc:"cost_currency_id,omitempty"`
	CostMethod                             *Selection `xmlrpc:"cost_method,omitempty"`
	CreateDate                             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                             *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCode                            *String    `xmlrpc:"default_code,omitempty"`
	Description                            *String    `xmlrpc:"description,omitempty"`
	DescriptionPicking                     *String    `xmlrpc:"description_picking,omitempty"`
	DescriptionPickingin                   *String    `xmlrpc:"description_pickingin,omitempty"`
	DescriptionPickingout                  *String    `xmlrpc:"description_pickingout,omitempty"`
	DescriptionPurchase                    *String    `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                        *String    `xmlrpc:"description_sale,omitempty"`
	DisplayName                            *String    `xmlrpc:"display_name,omitempty"`
	DisplayPrice                           *String    `xmlrpc:"display_price,omitempty"`
	ExpensePolicy                          *Selection `xmlrpc:"expense_policy,omitempty"`
	ExtraDaily                             *Float     `xmlrpc:"extra_daily,omitempty"`
	ExtraHourly                            *Float     `xmlrpc:"extra_hourly,omitempty"`
	FreeQty                                *Float     `xmlrpc:"free_qty,omitempty"`
	HasConfigurableAttributes              *Bool      `xmlrpc:"has_configurable_attributes,omitempty"`
	HsCode                                 *String    `xmlrpc:"hs_code,omitempty"`
	Id                                     *Int       `xmlrpc:"id,omitempty"`
	Image1024                              *String    `xmlrpc:"image_1024,omitempty"`
	Image128                               *String    `xmlrpc:"image_128,omitempty"`
	Image1920                              *String    `xmlrpc:"image_1920,omitempty"`
	Image256                               *String    `xmlrpc:"image_256,omitempty"`
	Image512                               *String    `xmlrpc:"image_512,omitempty"`
	ImageVariant1024                       *String    `xmlrpc:"image_variant_1024,omitempty"`
	ImageVariant128                        *String    `xmlrpc:"image_variant_128,omitempty"`
	ImageVariant1920                       *String    `xmlrpc:"image_variant_1920,omitempty"`
	ImageVariant256                        *String    `xmlrpc:"image_variant_256,omitempty"`
	ImageVariant512                        *String    `xmlrpc:"image_variant_512,omitempty"`
	IncomingQty                            *Float     `xmlrpc:"incoming_qty,omitempty"`
	IntrastatId                            *Many2One  `xmlrpc:"intrastat_id,omitempty"`
	IntrastatOriginCountryId               *Many2One  `xmlrpc:"intrastat_origin_country_id,omitempty"`
	InvoicePolicy                          *Selection `xmlrpc:"invoice_policy,omitempty"`
	IsProductVariant                       *Bool      `xmlrpc:"is_product_variant,omitempty"`
	IsPublished                            *Bool      `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized                         *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	ListPrice                              *Float     `xmlrpc:"list_price,omitempty"`
	LocationId                             *Many2One  `xmlrpc:"location_id,omitempty"`
	LstPrice                               *Float     `xmlrpc:"lst_price,omitempty"`
	MessageAttachmentCount                 *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds                      *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds                     *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                        *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                 *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                     *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                             *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                      *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId                *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                      *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter               *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                      *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                          *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter                   *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                                   *String    `xmlrpc:"name,omitempty"`
	NbrReorderingRules                     *Int       `xmlrpc:"nbr_reordering_rules,omitempty"`
	OrderpointIds                          *Relation  `xmlrpc:"orderpoint_ids,omitempty"`
	OutgoingQty                            *Float     `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation  `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                             *String    `xmlrpc:"partner_ref,omitempty"`
	PreparationTime                        *Float     `xmlrpc:"preparation_time,omitempty"`
	Price                                  *Float     `xmlrpc:"price,omitempty"`
	PriceExtra                             *Float     `xmlrpc:"price_extra,omitempty"`
	PricelistId                            *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	PricelistItemCount                     *Int       `xmlrpc:"pricelist_item_count,omitempty"`
	ProductTemplateAttributeValueIds       *Relation  `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTmplId                          *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductVariantCount                    *Int       `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                       *Many2One  `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                      *Relation  `xmlrpc:"product_variant_ids,omitempty"`
	ProjectId                              *Many2One  `xmlrpc:"project_id,omitempty"`
	ProjectTemplateId                      *Many2One  `xmlrpc:"project_template_id,omitempty"`
	PropertyAccountCreditorPriceDifference *Many2One  `xmlrpc:"property_account_creditor_price_difference,omitempty"`
	PropertyAccountExpenseId               *Many2One  `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId                *Many2One  `xmlrpc:"property_account_income_id,omitempty"`
	PropertyStockInventory                 *Many2One  `xmlrpc:"property_stock_inventory,omitempty"`
	PropertyStockProduction                *Many2One  `xmlrpc:"property_stock_production,omitempty"`
	PurchaseLineWarn                       *Selection `xmlrpc:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                    *String    `xmlrpc:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                         *Selection `xmlrpc:"purchase_method,omitempty"`
	PurchaseOk                             *Bool      `xmlrpc:"purchase_ok,omitempty"`
	PurchasedProductQty                    *Float     `xmlrpc:"purchased_product_qty,omitempty"`
	PutawayRuleIds                         *Relation  `xmlrpc:"putaway_rule_ids,omitempty"`
	QtyAvailable                           *Float     `xmlrpc:"qty_available,omitempty"`
	QtyInRent                              *Float     `xmlrpc:"qty_in_rent,omitempty"`
	QuantitySvl                            *Float     `xmlrpc:"quantity_svl,omitempty"`
	RatingAvg                              *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingCount                            *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                              *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback                     *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage                        *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastValue                        *Float     `xmlrpc:"rating_last_value,omitempty"`
	RecurringInvoice                       *Bool      `xmlrpc:"recurring_invoice,omitempty"`
	RentOk                                 *Bool      `xmlrpc:"rent_ok,omitempty"`
	Rental                                 *Bool      `xmlrpc:"rental,omitempty"`
	RentalPricingIds                       *Relation  `xmlrpc:"rental_pricing_ids,omitempty"`
	ReorderingMaxQty                       *Float     `xmlrpc:"reordering_max_qty,omitempty"`
	ReorderingMinQty                       *Float     `xmlrpc:"reordering_min_qty,omitempty"`
	ResponsibleId                          *Many2One  `xmlrpc:"responsible_id,omitempty"`
	RouteFromCategIds                      *Relation  `xmlrpc:"route_from_categ_ids,omitempty"`
	RouteIds                               *Relation  `xmlrpc:"route_ids,omitempty"`
	SaleDelay                              *Float     `xmlrpc:"sale_delay,omitempty"`
	SaleLineWarn                           *Selection `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                        *String    `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleOk                                 *Bool      `xmlrpc:"sale_ok,omitempty"`
	SalesCount                             *Float     `xmlrpc:"sales_count,omitempty"`
	SellerIds                              *Relation  `xmlrpc:"seller_ids,omitempty"`
	Sequence                               *Int       `xmlrpc:"sequence,omitempty"`
	ServicePolicy                          *Selection `xmlrpc:"service_policy,omitempty"`
	ServiceToPurchase                      *Bool      `xmlrpc:"service_to_purchase,omitempty"`
	ServiceTracking                        *Selection `xmlrpc:"service_tracking,omitempty"`
	ServiceType                            *Selection `xmlrpc:"service_type,omitempty"`
	StandardPrice                          *Float     `xmlrpc:"standard_price,omitempty"`
	StockMoveIds                           *Relation  `xmlrpc:"stock_move_ids,omitempty"`
	StockQuantIds                          *Relation  `xmlrpc:"stock_quant_ids,omitempty"`
	StockValuationLayerIds                 *Relation  `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	SubscriptionTemplateId                 *Many2One  `xmlrpc:"subscription_template_id,omitempty"`
	SupplierTaxesId                        *Relation  `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxesId                                *Relation  `xmlrpc:"taxes_id,omitempty"`
	Tracking                               *Selection `xmlrpc:"tracking,omitempty"`
	Type                                   *Selection `xmlrpc:"type,omitempty"`
	UomId                                  *Many2One  `xmlrpc:"uom_id,omitempty"`
	UomName                                *String    `xmlrpc:"uom_name,omitempty"`
	UomPoId                                *Many2One  `xmlrpc:"uom_po_id,omitempty"`
	ValidProductTemplateAttributeLineIds   *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	Valuation                              *Selection `xmlrpc:"valuation,omitempty"`
	ValueSvl                               *Float     `xmlrpc:"value_svl,omitempty"`
	VariantSellerIds                       *Relation  `xmlrpc:"variant_seller_ids,omitempty"`
	VirtualAvailable                       *Float     `xmlrpc:"virtual_available,omitempty"`
	VisibleExpensePolicy                   *Bool      `xmlrpc:"visible_expense_policy,omitempty"`
	VisibleQtyConfigurator                 *Bool      `xmlrpc:"visible_qty_configurator,omitempty"`
	Volume                                 *Float     `xmlrpc:"volume,omitempty"`
	VolumeUomName                          *String    `xmlrpc:"volume_uom_name,omitempty"`
	WarehouseId                            *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds                      *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription                 *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords                    *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg                       *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle                       *String    `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished                       *Bool      `xmlrpc:"website_published,omitempty"`
	Weight                                 *Float     `xmlrpc:"weight,omitempty"`
	WeightUomName                          *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv, nil)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp, nil)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductProductModel, criteria, options)
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
