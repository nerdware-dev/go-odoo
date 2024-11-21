package odoo

// ResUsers represents res.users model.
type ResUsers struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty" json:"__last_update,omitempty"`
	AccessesCount                 *Int       `xmlrpc:"accesses_count,omitempty" json:"accesses_count,omitempty"`
	ActionId                      *Many2One  `xmlrpc:"action_id,omitempty" json:"action_id,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omitempty" json:"active_lang_count,omitempty"`
	ActivePartner                 *Bool      `xmlrpc:"active_partner,omitempty" json:"active_partner,omitempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AdditionalNote                *String    `xmlrpc:"additional_note,omitempty" json:"additional_note,omitempty"`
	AddressHomeId                 *Many2One  `xmlrpc:"address_home_id,omitempty" json:"address_home_id,omitempty"`
	AddressId                     *Many2One  `xmlrpc:"address_id,omitempty" json:"address_id,omitempty"`
	AliasContact                  *Selection `xmlrpc:"alias_contact,omitempty" json:"alias_contact,omitempty"`
	AliasId                       *Many2One  `xmlrpc:"alias_id,omitempty" json:"alias_id,omitempty"`
	AllocationCount               *Float     `xmlrpc:"allocation_count,omitempty" json:"allocation_count,omitempty"`
	AllocationDisplay             *String    `xmlrpc:"allocation_display,omitempty" json:"allocation_display,omitempty"`
	AllocationUsedCount           *Float     `xmlrpc:"allocation_used_count,omitempty" json:"allocation_used_count,omitempty"`
	AllocationUsedDisplay         *String    `xmlrpc:"allocation_used_display,omitempty" json:"allocation_used_display,omitempty"`
	AttendanceState               *Selection `xmlrpc:"attendance_state,omitempty" json:"attendance_state,omitempty"`
	BadgeIds                      *Relation  `xmlrpc:"badge_ids,omitempty" json:"badge_ids,omitempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omitempty" json:"bank_account_count,omitempty"`
	BankAccountId                 *Many2One  `xmlrpc:"bank_account_id,omitempty" json:"bank_account_id,omitempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omitempty" json:"bank_ids,omitempty"`
	Barcode                       *String    `xmlrpc:"barcode,omitempty" json:"barcode,omitempty"`
	Birthday                      *Time      `xmlrpc:"birthday,omitempty" json:"birthday,omitempty"`
	BronzeBadge                   *Int       `xmlrpc:"bronze_badge,omitempty" json:"bronze_badge,omitempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omitempty" json:"calendar_last_notif_ack,omitempty"`
	CanEdit                       *Bool      `xmlrpc:"can_edit,omitempty" json:"can_edit,omitempty"`
	CanPublish                    *Bool      `xmlrpc:"can_publish,omitempty" json:"can_publish,omitempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omitempty" json:"category_id,omitempty"`
	CategoryIds                   *Relation  `xmlrpc:"category_ids,omitempty" json:"category_ids,omitempty"`
	Certificate                   *Selection `xmlrpc:"certificate,omitempty" json:"certificate,omitempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omitempty" json:"channel_ids,omitempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	Children                      *Int       `xmlrpc:"children,omitempty" json:"children,omitempty"`
	City                          *String    `xmlrpc:"city,omitempty" json:"city,omitempty"`
	CoachId                       *Many2One  `xmlrpc:"coach_id,omitempty" json:"coach_id,omitempty"`
	Color                         *Int       `xmlrpc:"color,omitempty" json:"color,omitempty"`
	Comment                       *String    `xmlrpc:"comment,omitempty" json:"comment,omitempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omitempty" json:"commercial_company_name,omitempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omitempty" json:"commercial_partner_id,omitempty"`
	CommunicationMedias           *Relation  `xmlrpc:"communication_medias,omitempty" json:"communication_medias,omitempty"`
	CompaniesCount                *Int       `xmlrpc:"companies_count,omitempty" json:"companies_count,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CompanyIds                    *Relation  `xmlrpc:"company_ids,omitempty" json:"company_ids,omitempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omitempty" json:"company_name,omitempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omitempty" json:"company_type,omitempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omitempty" json:"contact_address,omitempty"`
	ContactAddressComplete        *String    `xmlrpc:"contact_address_complete,omitempty" json:"contact_address_complete,omitempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omitempty" json:"contract_ids,omitempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omitempty" json:"country_id,omitempty"`
	CountryOfBirth                *Many2One  `xmlrpc:"country_of_birth,omitempty" json:"country_of_birth,omitempty"`
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
	DepartmentId                  *Many2One  `xmlrpc:"department_id,omitempty" json:"department_id,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentCount                 *Int       `xmlrpc:"document_count,omitempty" json:"document_count,omitempty"`
	DocumentIds                   *Relation  `xmlrpc:"document_ids,omitempty" json:"document_ids,omitempty"`
	Email                         *String    `xmlrpc:"email,omitempty" json:"email,omitempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omitempty" json:"email_formatted,omitempty"`
	EmailNormalized               *String    `xmlrpc:"email_normalized,omitempty" json:"email_normalized,omitempty"`
	EmergencyContact              *String    `xmlrpc:"emergency_contact,omitempty" json:"emergency_contact,omitempty"`
	EmergencyPhone                *String    `xmlrpc:"emergency_phone,omitempty" json:"emergency_phone,omitempty"`
	Employee                      *Bool      `xmlrpc:"employee,omitempty" json:"employee,omitempty"`
	EmployeeBankAccountId         *Many2One  `xmlrpc:"employee_bank_account_id,omitempty" json:"employee_bank_account_id,omitempty"`
	EmployeeCarsCount             *Int       `xmlrpc:"employee_cars_count,omitempty" json:"employee_cars_count,omitempty"`
	EmployeeCount                 *Int       `xmlrpc:"employee_count,omitempty" json:"employee_count,omitempty"`
	EmployeeCountryId             *Many2One  `xmlrpc:"employee_country_id,omitempty" json:"employee_country_id,omitempty"`
	EmployeeId                    *Many2One  `xmlrpc:"employee_id,omitempty" json:"employee_id,omitempty"`
	EmployeeIds                   *Relation  `xmlrpc:"employee_ids,omitempty" json:"employee_ids,omitempty"`
	EmployeeParentId              *Many2One  `xmlrpc:"employee_parent_id,omitempty" json:"employee_parent_id,omitempty"`
	EmployeePhone                 *String    `xmlrpc:"employee_phone,omitempty" json:"employee_phone,omitempty"`
	EmployeeSkillIds              *Relation  `xmlrpc:"employee_skill_ids,omitempty" json:"employee_skill_ids,omitempty"`
	ExpenseManagerId              *Many2One  `xmlrpc:"expense_manager_id,omitempty" json:"expense_manager_id,omitempty"`
	FavoriteLunchProductIds       *Relation  `xmlrpc:"favorite_lunch_product_ids,omitempty" json:"favorite_lunch_product_ids,omitempty"`
	FollowupLevel                 *Many2One  `xmlrpc:"followup_level,omitempty" json:"followup_level,omitempty"`
	FollowupStatus                *Selection `xmlrpc:"followup_status,omitempty" json:"followup_status,omitempty"`
	ForumWaitingPostsCount        *Int       `xmlrpc:"forum_waiting_posts_count,omitempty" json:"forum_waiting_posts_count,omitempty"`
	Function                      *String    `xmlrpc:"function,omitempty" json:"function,omitempty"`
	Gender                        *Selection `xmlrpc:"gender,omitempty" json:"gender,omitempty"`
	GoalIds                       *Relation  `xmlrpc:"goal_ids,omitempty" json:"goal_ids,omitempty"`
	GoldBadge                     *Int       `xmlrpc:"gold_badge,omitempty" json:"gold_badge,omitempty"`
	GroupsCount                   *Int       `xmlrpc:"groups_count,omitempty" json:"groups_count,omitempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omitempty" json:"groups_id,omitempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omitempty" json:"has_unreconciled_entries,omitempty"`
	HelpdeskTargetClosed          *Float     `xmlrpc:"helpdesk_target_closed,omitempty" json:"helpdesk_target_closed,omitempty"`
	HelpdeskTargetRating          *Float     `xmlrpc:"helpdesk_target_rating,omitempty" json:"helpdesk_target_rating,omitempty"`
	HelpdeskTargetSuccess         *Float     `xmlrpc:"helpdesk_target_success,omitempty" json:"helpdesk_target_success,omitempty"`
	HoursLastMonth                *Float     `xmlrpc:"hours_last_month,omitempty" json:"hours_last_month,omitempty"`
	HoursLastMonthDisplay         *String    `xmlrpc:"hours_last_month_display,omitempty" json:"hours_last_month_display,omitempty"`
	HrPresenceState               *Selection `xmlrpc:"hr_presence_state,omitempty" json:"hr_presence_state,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IdentificationId              *String    `xmlrpc:"identification_id,omitempty" json:"identification_id,omitempty"`
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
	IsAbsent                      *Bool      `xmlrpc:"is_absent,omitempty" json:"is_absent,omitempty"`
	IsAddressHomeACompany         *Bool      `xmlrpc:"is_address_home_a_company,omitempty" json:"is_address_home_a_company,omitempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omitempty" json:"is_blacklisted,omitempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omitempty" json:"is_company,omitempty"`
	IsModerator                   *Bool      `xmlrpc:"is_moderator,omitempty" json:"is_moderator,omitempty"`
	IsPublished                   *Bool      `xmlrpc:"is_published,omitempty" json:"is_published,omitempty"`
	IsSeoOptimized                *Bool      `xmlrpc:"is_seo_optimized,omitempty" json:"is_seo_optimized,omitempty"`
	JobTitle                      *String    `xmlrpc:"job_title,omitempty" json:"job_title,omitempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omitempty" json:"journal_item_count,omitempty"`
	Karma                         *Int       `xmlrpc:"karma,omitempty" json:"karma,omitempty"`
	KmHomeWork                    *Int       `xmlrpc:"km_home_work,omitempty" json:"km_home_work,omitempty"`
	L10NDeDatevIdentifier         *Int       `xmlrpc:"l10n_de_datev_identifier,omitempty" json:"l10n_de_datev_identifier,omitempty"`
	Lang                          *Selection `xmlrpc:"lang,omitempty" json:"lang,omitempty"`
	LastActivity                  *Time      `xmlrpc:"last_activity,omitempty" json:"last_activity,omitempty"`
	LastActivityTime              *String    `xmlrpc:"last_activity_time,omitempty" json:"last_activity_time,omitempty"`
	LastCheckIn                   *Time      `xmlrpc:"last_check_in,omitempty" json:"last_check_in,omitempty"`
	LastCheckOut                  *Time      `xmlrpc:"last_check_out,omitempty" json:"last_check_out,omitempty"`
	LastLunchLocationId           *Many2One  `xmlrpc:"last_lunch_location_id,omitempty" json:"last_lunch_location_id,omitempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omitempty" json:"last_time_entries_checked,omitempty"`
	LeaveDateTo                   *Time      `xmlrpc:"leave_date_to,omitempty" json:"leave_date_to,omitempty"`
	LeaveManagerId                *Many2One  `xmlrpc:"leave_manager_id,omitempty" json:"leave_manager_id,omitempty"`
	LogIds                        *Relation  `xmlrpc:"log_ids,omitempty" json:"log_ids,omitempty"`
	Login                         *String    `xmlrpc:"login,omitempty" json:"login,omitempty"`
	LoginDate                     *Time      `xmlrpc:"login_date,omitempty" json:"login_date,omitempty"`
	Marital                       *Selection `xmlrpc:"marital,omitempty" json:"marital,omitempty"`
	MedicExam                     *Time      `xmlrpc:"medic_exam,omitempty" json:"medic_exam,omitempty"`
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
	MobilePhone                   *String    `xmlrpc:"mobile_phone,omitempty" json:"mobile_phone,omitempty"`
	ModerationChannelIds          *Relation  `xmlrpc:"moderation_channel_ids,omitempty" json:"moderation_channel_ids,omitempty"`
	ModerationCounter             *Int       `xmlrpc:"moderation_counter,omitempty" json:"moderation_counter,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	NewPassword                   *String    `xmlrpc:"new_password,omitempty" json:"new_password,omitempty"`
	NextRankId                    *Many2One  `xmlrpc:"next_rank_id,omitempty" json:"next_rank_id,omitempty"`
	NotificationType              *Selection `xmlrpc:"notification_type,omitempty" json:"notification_type,omitempty"`
	OauthAccessToken              *String    `xmlrpc:"oauth_access_token,omitempty" json:"oauth_access_token,omitempty"`
	OauthProviderId               *Many2One  `xmlrpc:"oauth_provider_id,omitempty" json:"oauth_provider_id,omitempty"`
	OauthUid                      *String    `xmlrpc:"oauth_uid,omitempty" json:"oauth_uid,omitempty"`
	OcnToken                      *String    `xmlrpc:"ocn_token,omitempty" json:"ocn_token,omitempty"`
	OdooComUid                    *String    `xmlrpc:"odoo_com_uid,omitempty" json:"odoo_com_uid,omitempty"`
	OdoobotState                  *Selection `xmlrpc:"odoobot_state,omitempty" json:"odoobot_state,omitempty"`
	OnlinePartnerBankAccount      *String    `xmlrpc:"online_partner_bank_account,omitempty" json:"online_partner_bank_account,omitempty"`
	OnlinePartnerVendorName       *String    `xmlrpc:"online_partner_vendor_name,omitempty" json:"online_partner_vendor_name,omitempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omitempty" json:"opportunity_count,omitempty"`
	OpportunityCountIds           *Relation  `xmlrpc:"opportunity_count_ids,omitempty" json:"opportunity_count_ids,omitempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omitempty" json:"opportunity_ids,omitempty"`
	OutOfOfficeMessage            *String    `xmlrpc:"out_of_office_message,omitempty" json:"out_of_office_message,omitempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omitempty" json:"parent_id,omitempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omitempty" json:"parent_name,omitempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omitempty" json:"partner_latitude,omitempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omitempty" json:"partner_longitude,omitempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omitempty" json:"partner_share,omitempty"`
	PassportId                    *String    `xmlrpc:"passport_id,omitempty" json:"passport_id,omitempty"`
	Password                      *String    `xmlrpc:"password,omitempty" json:"password,omitempty"`
	PaymentNextActionDate         *Time      `xmlrpc:"payment_next_action_date,omitempty" json:"payment_next_action_date,omitempty"`
	PaymentResponsibleId          *Many2One  `xmlrpc:"payment_responsible_id,omitempty" json:"payment_responsible_id,omitempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omitempty" json:"payment_token_count,omitempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omitempty" json:"payment_token_ids,omitempty"`
	PermitNo                      *String    `xmlrpc:"permit_no,omitempty" json:"permit_no,omitempty"`
	Phone                         *String    `xmlrpc:"phone,omitempty" json:"phone,omitempty"`
	PhoneBlacklisted              *Bool      `xmlrpc:"phone_blacklisted,omitempty" json:"phone_blacklisted,omitempty"`
	PhoneSanitized                *String    `xmlrpc:"phone_sanitized,omitempty" json:"phone_sanitized,omitempty"`
	PickingWarn                   *Selection `xmlrpc:"picking_warn,omitempty" json:"picking_warn,omitempty"`
	PickingWarnMsg                *String    `xmlrpc:"picking_warn_msg,omitempty" json:"picking_warn_msg,omitempty"`
	Pin                           *String    `xmlrpc:"pin,omitempty" json:"pin,omitempty"`
	PlaceOfBirth                  *String    `xmlrpc:"place_of_birth,omitempty" json:"place_of_birth,omitempty"`
	PlanToChangeCar               *Bool      `xmlrpc:"plan_to_change_car,omitempty" json:"plan_to_change_car,omitempty"`
	PrivateEmail                  *String    `xmlrpc:"private_email,omitempty" json:"private_email,omitempty"`
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
	RankId                        *Many2One  `xmlrpc:"rank_id,omitempty" json:"rank_id,omitempty"`
	Ref                           *String    `xmlrpc:"ref,omitempty" json:"ref,omitempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omitempty" json:"ref_company_ids,omitempty"`
	ResourceCalendarId            *Many2One  `xmlrpc:"resource_calendar_id,omitempty" json:"resource_calendar_id,omitempty"`
	ResourceIds                   *Relation  `xmlrpc:"resource_ids,omitempty" json:"resource_ids,omitempty"`
	ResumeLineIds                 *Relation  `xmlrpc:"resume_line_ids,omitempty" json:"resume_line_ids,omitempty"`
	RulesCount                    *Int       `xmlrpc:"rules_count,omitempty" json:"rules_count,omitempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omitempty" json:"sale_order_count,omitempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omitempty" json:"sale_order_ids,omitempty"`
	SaleTeamId                    *Many2One  `xmlrpc:"sale_team_id,omitempty" json:"sale_team_id,omitempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omitempty" json:"sale_warn,omitempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omitempty" json:"sale_warn_msg,omitempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omitempty" json:"same_vat_partner_id,omitempty"`
	SddCount                      *Int       `xmlrpc:"sdd_count,omitempty" json:"sdd_count,omitempty"`
	SddMandateIds                 *Relation  `xmlrpc:"sdd_mandate_ids,omitempty" json:"sdd_mandate_ids,omitempty"`
	Self                          *Many2One  `xmlrpc:"self,omitempty" json:"self,omitempty"`
	Share                         *Bool      `xmlrpc:"share,omitempty" json:"share,omitempty"`
	ShowLeaves                    *Bool      `xmlrpc:"show_leaves,omitempty" json:"show_leaves,omitempty"`
	Signature                     *String    `xmlrpc:"signature,omitempty" json:"signature,omitempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omitempty" json:"signup_expiration,omitempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omitempty" json:"signup_token,omitempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omitempty" json:"signup_type,omitempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omitempty" json:"signup_url,omitempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omitempty" json:"signup_valid,omitempty"`
	SilverBadge                   *Int       `xmlrpc:"silver_badge,omitempty" json:"silver_badge,omitempty"`
	SlaCompleted                  *Bool      `xmlrpc:"sla_completed,omitempty" json:"sla_completed,omitempty"`
	SlaDate                       *Time      `xmlrpc:"sla_date,omitempty" json:"sla_date,omitempty"`
	SlaTicketCreated              *Bool      `xmlrpc:"sla_ticket_created,omitempty" json:"sla_ticket_created,omitempty"`
	SpouseBirthdate               *Time      `xmlrpc:"spouse_birthdate,omitempty" json:"spouse_birthdate,omitempty"`
	SpouseCompleteName            *String    `xmlrpc:"spouse_complete_name,omitempty" json:"spouse_complete_name,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omitempty" json:"state_id,omitempty"`
	Street                        *String    `xmlrpc:"street,omitempty" json:"street,omitempty"`
	Street2                       *String    `xmlrpc:"street2,omitempty" json:"street2,omitempty"`
	StudyField                    *String    `xmlrpc:"study_field,omitempty" json:"study_field,omitempty"`
	StudySchool                   *String    `xmlrpc:"study_school,omitempty" json:"study_school,omitempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omitempty" json:"subscription_count,omitempty"`
	SupplierInvoiceCount          *Int       `xmlrpc:"supplier_invoice_count,omitempty" json:"supplier_invoice_count,omitempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omitempty" json:"supplier_rank,omitempty"`
	TargetSalesDone               *Int       `xmlrpc:"target_sales_done,omitempty" json:"target_sales_done,omitempty"`
	TargetSalesInvoiced           *Int       `xmlrpc:"target_sales_invoiced,omitempty" json:"target_sales_invoiced,omitempty"`
	TargetSalesWon                *Int       `xmlrpc:"target_sales_won,omitempty" json:"target_sales_won,omitempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omitempty" json:"task_count,omitempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omitempty" json:"task_ids,omitempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omitempty" json:"team_id,omitempty"`
	TicketCount                   *Int       `xmlrpc:"ticket_count,omitempty" json:"ticket_count,omitempty"`
	TimesheetManagerId            *Many2One  `xmlrpc:"timesheet_manager_id,omitempty" json:"timesheet_manager_id,omitempty"`
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
	Vehicle                       *String    `xmlrpc:"vehicle,omitempty" json:"vehicle,omitempty"`
	VisaExpire                    *Time      `xmlrpc:"visa_expire,omitempty" json:"visa_expire,omitempty"`
	VisaNo                        *String    `xmlrpc:"visa_no,omitempty" json:"visa_no,omitempty"`
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
	WorkEmail                     *String    `xmlrpc:"work_email,omitempty" json:"work_email,omitempty"`
	WorkLocation                  *String    `xmlrpc:"work_location,omitempty" json:"work_location,omitempty"`
	WorkPhone                     *String    `xmlrpc:"work_phone,omitempty" json:"work_phone,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
	Zip                           *String    `xmlrpc:"zip,omitempty" json:"zip,omitempty"`
}

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	ids, err := c.CreateResUserss([]*ResUsers{ru})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUserss(rus []*ResUsers) ([]int64, error) {
	var vv []interface{}
	for _, v := range rus {
		vv = append(vv, v)
	}
	return c.Create(ResUsersModel, vv, nil)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru, nil)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rus)[0]), nil
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	return &((*rus)[0]), nil
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersModel, criteria, options)
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
