package odoo

// ResPartner represents res.partner model.
type ResPartner struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omitempty" json:"active_lang_count,omitempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omitempty" json:"bank_account_count,omitempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omitempty" json:"bank_ids,omitempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omitempty" json:"calendar_last_notif_ack,omitempty"`
	CanPublish                    *Bool      `xmlrpc:"can_publish,omitempty" json:"can_publish,omitempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omitempty" json:"category_id,omitempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omitempty" json:"channel_ids,omitempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	City                          *String    `xmlrpc:"city,omitempty" json:"city,omitempty"`
	Color                         *Int       `xmlrpc:"color,omitempty" json:"color,omitempty"`
	Comment                       *String    `xmlrpc:"comment,omitempty" json:"comment,omitempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omitempty" json:"commercial_company_name,omitempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CommunicationMedias           *Relation  `xmlrpc:"communication_medias,omitempty" json:"communication_medias,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omitempty" json:"company_name,omitempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omitempty" json:"company_type,omitempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omitempty" json:"contact_address,omitempty"`
	ContactAddressComplete        *String    `xmlrpc:"contact_address_complete,omitempty" json:"contact_address_complete,omitempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omitempty" json:"contract_ids,omitempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Credit                        *Float     `xmlrpc:"credit,omitempty" json:"credit,omitempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omitempty" json:"credit_limit,omitempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omitempty" json:"currency_id,omitempty"`
	CustomerRank                  *Int       `xmlrpc:"customer_rank,omitempty" json:"customer_rank,omitempty"`
	DataProtectionEnabled         *Bool      `xmlrpc:"data_protection_enabled,omitempty" json:"data_protection_enabled,omitempty"`
	Date                          *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Debit                         *Float     `xmlrpc:"debit,omitempty" json:"debit,omitempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omitempty" json:"debit_limit,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentCount                 *Int       `xmlrpc:"document_count,omitempty" json:"document_count,omitempty"`
	Email                         *String    `xmlrpc:"email,omitempty" json:"email,omitempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omitempty" json:"email_formatted,omitempty"`
	EmailNormalized               *String    `xmlrpc:"email_normalized,omitempty" json:"email_normalized,omitempty"`
	Employee                      *Bool      `xmlrpc:"employee,omitempty" json:"employee,omitempty"`
	FollowupLevel                 *Many2One  `xmlrpc:"followup_level,omitempty" json:"followup_level,omitempty"`
	FollowupStatus                *Selection `xmlrpc:"followup_status,omitempty" json:"followup_status,omitempty"`
	Function                      *String    `xmlrpc:"function,omitempty" json:"function,omitempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omitempty" json:"has_unreconciled_entries,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omitempty" json:"im_status,omitempty"`
	Image1024                     *String    `xmlrpc:"image_1024,omitempty" json:"image_1024,omitempty"`
	Image128                      *String    `xmlrpc:"image_128,omitempty" json:"image_128,omitempty"`
	Image1920                     *String    `xmlrpc:"image_1920,omitempty" json:"image_1920,omitempty"`
	Image256                      *String    `xmlrpc:"image_256,omitempty" json:"image_256,omitempty"`
	Image512                      *String    `xmlrpc:"image_512,omitempty" json:"image_512,omitempty"`
	ImageMedium                   *String    `xmlrpc:"image_medium,omitempty" json:"image_medium,omitempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omitempty" json:"industry_id,omitempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omitempty" json:"invoice_ids,omitempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omitempty" json:"invoice_warn,omitempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omitempty" json:"invoice_warn_msg,omitempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omitempty" json:"is_blacklisted,omitempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omitempty" json:"is_company,omitempty"`
	IsPublished                   *Bool      `xmlrpc:"is_published,omitempty" json:"is_published,omitempty"`
	IsSeoOptimized                *Bool      `xmlrpc:"is_seo_optimized,omitempty" json:"is_seo_optimized,omitempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omitempty" json:"journal_item_count,omitempty"`
	L10NDeDatevIdentifier         *Int       `xmlrpc:"l10n_de_datev_identifier,omitempty" json:"l10n_de_datev_identifier,omitempty"`
	Lang                          *Selection `xmlrpc:"lang,omitempty" json:"lang,omitempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omitempty" json:"last_time_entries_checked,omitempty"`
	MeetingCount                  *Int       `xmlrpc:"meeting_count,omitempty" json:"meeting_count,omitempty"`
	MeetingIds                    *Relation  `xmlrpc:"meeting_ids,omitempty" json:"meeting_ids,omitempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omitempty" json:"message_bounce,omitempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omitempty" json:"message_channel_ids,omitempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omitempty" json:"message_main_attachment_id,omitempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omitempty" json:"message_unread,omitempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omitempty" json:"message_unread_counter,omitempty"`
	Mobile                        *String    `xmlrpc:"mobile,omitempty" json:"mobile,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OcnToken                      *String    `xmlrpc:"ocn_token,omitempty" json:"ocn_token,omitempty"`
	OnlinePartnerBankAccount      *String    `xmlrpc:"online_partner_bank_account,omitempty" json:"online_partner_bank_account,omitempty"`
	OnlinePartnerVendorName       *String    `xmlrpc:"online_partner_vendor_name,omitempty" json:"online_partner_vendor_name,omitempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omitempty" json:"opportunity_count,omitempty"`
	OpportunityCountIds           *Relation  `xmlrpc:"opportunity_count_ids,omitempty" json:"opportunity_count_ids,omitempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omitempty" json:"opportunity_ids,omitempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omitempty" json:"parent_id,omitempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omitempty" json:"parent_name,omitempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omitempty" json:"partner_latitude,omitempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omitempty" json:"partner_longitude,omitempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omitempty" json:"partner_share,omitempty"`
	PaymentNextActionDate         *Time      `xmlrpc:"payment_next_action_date,omitempty" json:"payment_next_action_date,omitempty"`
	PaymentResponsibleId          *Many2One  `xmlrpc:"payment_responsible_id,omitempty" json:"payment_responsible_id,omitempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omitempty" json:"payment_token_count,omitempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omitempty" json:"payment_token_ids,omitempty"`
	Phone                         *String    `xmlrpc:"phone,omitempty" json:"phone,omitempty"`
	PhoneBlacklisted              *Bool      `xmlrpc:"phone_blacklisted,omitempty" json:"phone_blacklisted,omitempty"`
	PhoneSanitized                *String    `xmlrpc:"phone_sanitized,omitempty" json:"phone_sanitized,omitempty"`
	PickingWarn                   *Selection `xmlrpc:"picking_warn,omitempty" json:"picking_warn,omitempty"`
	PickingWarnMsg                *String    `xmlrpc:"picking_warn_msg,omitempty" json:"picking_warn_msg,omitempty"`
	PlanToChangeCar               *Bool      `xmlrpc:"plan_to_change_car,omitempty" json:"plan_to_change_car,omitempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omitempty" json:"property_account_payable_id,omitempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omitempty" json:"property_account_position_id,omitempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omitempty" json:"property_account_receivable_id,omitempty"`
	PropertyDeliveryCarrierId     *Many2One  `xmlrpc:"property_delivery_carrier_id,omitempty" json:"property_delivery_carrier_id,omitempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omitempty" json:"property_payment_term_id,omitempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omitempty" json:"property_product_pricelist,omitempty"`
	PropertyPurchaseCurrencyId    *Many2One  `xmlrpc:"property_purchase_currency_id,omitempty" json:"property_purchase_currency_id,omitempty"`
	PropertyStockCustomer         *Many2One  `xmlrpc:"property_stock_customer,omitempty" json:"property_stock_customer,omitempty"`
	PropertyStockSupplier         *Many2One  `xmlrpc:"property_stock_supplier,omitempty" json:"property_stock_supplier,omitempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omitempty" json:"property_supplier_payment_term_id,omitempty"`
	PurchaseOrderCount            *Int       `xmlrpc:"purchase_order_count,omitempty" json:"purchase_order_count,omitempty"`
	PurchaseWarn                  *Selection `xmlrpc:"purchase_warn,omitempty" json:"purchase_warn,omitempty"`
	PurchaseWarnMsg               *String    `xmlrpc:"purchase_warn_msg,omitempty" json:"purchase_warn_msg,omitempty"`
	Ref                           *String    `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omitempty" json:"ref_company_ids,omitempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omitempty" json:"sale_order_ids,omitempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omitempty" json:"sale_warn,omitempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omitempty" json:"sale_warn_msg,omitempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omitempty" json:"same_vat_partner_id,omitempty"`
	SddCount                      *Int       `xmlrpc:"sdd_count,omitempty" json:"sdd_count,omitempty"`
	SddMandateIds                 *Relation  `xmlrpc:"sdd_mandate_ids,omitempty" json:"sdd_mandate_ids,omitempty"`
	Self                          *Many2One  `xmlrpc:"self,omitempty" json:"self,omitempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omitempty" json:"signup_expiration,omitempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omitempty" json:"signup_token,omitempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omitempty" json:"signup_type,omitempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omitempty" json:"signup_url,omitempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omitempty" json:"signup_valid,omitempty"`
	SlaCompleted                  *Bool      `xmlrpc:"sla_completed,omitempty" json:"sla_completed,omitempty"`
	SlaDate                       *Time      `xmlrpc:"sla_date,omitempty" json:"sla_date,omitempty"`
	SlaTicketCreated              *Bool      `xmlrpc:"sla_ticket_created,omitempty" json:"sla_ticket_created,omitempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omitempty" json:"state_id,omitempty"`
	Street                        *String    `xmlrpc:"street,omitempty" json:"street,omitempty"`
	Street2                       *String    `xmlrpc:"street2,omitempty" json:"street2,omitempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omitempty" json:"subscription_count,omitempty"`
	SupplierInvoiceCount          *Int       `xmlrpc:"supplier_invoice_count,omitempty" json:"supplier_invoice_count,omitempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omitempty" json:"supplier_rank,omitempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omitempty" json:"task_count,omitempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omitempty" json:"task_ids,omitempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	TicketCount                   *Int       `xmlrpc:"ticket_count,omitempty" json:"ticket_count,omitempty"`
	Title                         *Many2One  `xmlrpc:"title,omitempty" json:"title,omitempty"`
	TotalDue                      *Float     `xmlrpc:"total_due,omitempty" json:"total_due,omitempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omitempty" json:"total_invoiced,omitempty"`
	TotalOverdue                  *Float     `xmlrpc:"total_overdue,omitempty" json:"total_overdue,omitempty"`
	Trust                         *Selection `xmlrpc:"trust,omitempty" json:"trust,omitempty"`
	Type                          *Selection `xmlrpc:"type,omitempty" json:"type,omitempty"`
	Tz                            *Selection `xmlrpc:"tz,omitempty" json:"tz,omitempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omitempty" json:"tz_offset,omitempty"`
	UnpaidInvoices                *Relation  `xmlrpc:"unpaid_invoices,omitempty" json:"unpaid_invoices,omitempty"`
	UnreconciledAmlIds            *Relation  `xmlrpc:"unreconciled_aml_ids,omitempty" json:"unreconciled_aml_ids,omitempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omitempty" json:"user_id,omitempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omitempty" json:"user_ids,omitempty"`
	Vat                           *String    `xmlrpc:"vat,omitempty" json:"vat,omitempty"`
	VisitorIds                    *Relation  `xmlrpc:"visitor_ids,omitempty" json:"visitor_ids,omitempty"`
	Website                       *String    `xmlrpc:"website,omitempty" json:"website,omitempty"`
	WebsiteDescription            *String    `xmlrpc:"website_description,omitempty" json:"website_description,omitempty"`
	WebsiteId                     *Many2One  `xmlrpc:"website_id,omitempty" json:"website_id,omitempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WebsiteMetaDescription        *String    `xmlrpc:"website_meta_description,omitempty" json:"website_meta_description,omitempty"`
	WebsiteMetaKeywords           *String    `xmlrpc:"website_meta_keywords,omitempty" json:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg              *String    `xmlrpc:"website_meta_og_img,omitempty" json:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle              *String    `xmlrpc:"website_meta_title,omitempty" json:"website_meta_title,omitempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omitempty" json:"website_published,omitempty"`
	WebsiteShortDescription       *String    `xmlrpc:"website_short_description,omitempty" json:"website_short_description,omitempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omitempty" json:"website_url,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
	Zip                           *String    `xmlrpc:"zip,omitempty" json:"zip,omitempty"`
}

// ResPartners represents array of res.partner model.
type ResPartners []ResPartner

// ResPartnerModel is the odoo model name.
const ResPartnerModel = "res.partner"

// Many2One convert ResPartner to *Many2One.
func (rp *ResPartner) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartner(rp *ResPartner) (int64, error) {
	ids, err := c.CreateResPartners([]*ResPartner{rp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartners(rps []*ResPartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range rps {
		vv = append(vv, v)
	}
	return c.Create(ResPartnerModel, vv, nil)
}

// UpdateResPartner updates an existing res.partner record.
func (c *Client) UpdateResPartner(rp *ResPartner) error {
	return c.UpdateResPartners([]int64{rp.Id.Get()}, rp)
}

// UpdateResPartners updates existing res.partner records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateResPartners(ids []int64, rp *ResPartner) error {
	return c.Update(ResPartnerModel, ids, rp, nil)
}

// DeleteResPartner deletes an existing res.partner record.
func (c *Client) DeleteResPartner(id int64) error {
	return c.DeleteResPartners([]int64{id})
}

// DeleteResPartners deletes existing res.partner records.
func (c *Client) DeleteResPartners(ids []int64) error {
	return c.Delete(ResPartnerModel, ids)
}

// GetResPartner gets res.partner existing record.
func (c *Client) GetResPartner(id int64) (*ResPartner, error) {
	rps, err := c.GetResPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rps)[0]), nil
}

// GetResPartners gets res.partner existing records.
func (c *Client) GetResPartners(ids []int64) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.Read(ResPartnerModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartner finds res.partner record by querying it with criteria.
func (c *Client) FindResPartner(criteria *Criteria) (*ResPartner, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	return &((*rps)[0]), nil
}

// FindResPartners finds res.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartners(criteria *Criteria, options *Options) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResPartnerModel, criteria, options)
}

// FindResPartnerId finds record id by querying it with criteria.
func (c *Client) FindResPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
