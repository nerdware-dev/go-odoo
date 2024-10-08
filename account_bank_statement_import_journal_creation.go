package odoo

// AccountBankStatementImportJournalCreation represents account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreation struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccountControlIds           *Relation  `xmlrpc:"account_control_ids,omitempty"`
	AccountOnlineJournalId      *Many2One  `xmlrpc:"account_online_journal_id,omitempty"`
	AccountOnlineProviderId     *Many2One  `xmlrpc:"account_online_provider_id,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AliasDomain                 *String    `xmlrpc:"alias_domain,omitempty"`
	AliasId                     *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasName                   *String    `xmlrpc:"alias_name,omitempty"`
	AtLeastOneInbound           *Bool      `xmlrpc:"at_least_one_inbound,omitempty"`
	AtLeastOneOutbound          *Bool      `xmlrpc:"at_least_one_outbound,omitempty"`
	BankAccNumber               *String    `xmlrpc:"bank_acc_number,omitempty"`
	BankAccountId               *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	BankId                      *Many2One  `xmlrpc:"bank_id,omitempty"`
	BankStatementCreation       *Selection `xmlrpc:"bank_statement_creation,omitempty"`
	BankStatementsSource        *Selection `xmlrpc:"bank_statements_source,omitempty"`
	Code                        *String    `xmlrpc:"code,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyPartnerId            *Many2One  `xmlrpc:"company_partner_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCreditAccountId      *Many2One  `xmlrpc:"default_credit_account_id,omitempty"`
	DefaultDebitAccountId       *Many2One  `xmlrpc:"default_debit_account_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InboundPaymentMethodIds     *Relation  `xmlrpc:"inbound_payment_method_ids,omitempty"`
	InvoiceReferenceModel       *Selection `xmlrpc:"invoice_reference_model,omitempty"`
	InvoiceReferenceType        *Selection `xmlrpc:"invoice_reference_type,omitempty"`
	JournalGroupIds             *Relation  `xmlrpc:"journal_group_ids,omitempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omitempty"`
	JsonActivityData            *String    `xmlrpc:"json_activity_data,omitempty"`
	KanbanDashboard             *String    `xmlrpc:"kanban_dashboard,omitempty"`
	KanbanDashboardGraph        *String    `xmlrpc:"kanban_dashboard_graph,omitempty"`
	LossAccountId               *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NextSynchronization         *Time      `xmlrpc:"next_synchronization,omitempty"`
	OutboundPaymentMethodIds    *Relation  `xmlrpc:"outbound_payment_method_ids,omitempty"`
	PostAt                      *Selection `xmlrpc:"post_at,omitempty"`
	ProfitAccountId             *Many2One  `xmlrpc:"profit_account_id,omitempty"`
	RefundSequence              *Bool      `xmlrpc:"refund_sequence,omitempty"`
	RefundSequenceId            *Many2One  `xmlrpc:"refund_sequence_id,omitempty"`
	RefundSequenceNumberNext    *Int       `xmlrpc:"refund_sequence_number_next,omitempty"`
	RestrictModeHashTable       *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	SecureSequenceId            *Many2One  `xmlrpc:"secure_sequence_id,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	SequenceId                  *Many2One  `xmlrpc:"sequence_id,omitempty"`
	SequenceNumberNext          *Int       `xmlrpc:"sequence_number_next,omitempty"`
	ShowOnDashboard             *Bool      `xmlrpc:"show_on_dashboard,omitempty"`
	SynchronizationStatus       *String    `xmlrpc:"synchronization_status,omitempty"`
	Type                        *Selection `xmlrpc:"type,omitempty"`
	TypeControlIds              *Relation  `xmlrpc:"type_control_ids,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatementImportJournalCreations represents array of account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreations []AccountBankStatementImportJournalCreation

// AccountBankStatementImportJournalCreationModel is the odoo model name.
const AccountBankStatementImportJournalCreationModel = "account.bank.statement.import.journal.creation"

// Many2One convert AccountBankStatementImportJournalCreation to *Many2One.
func (absijc *AccountBankStatementImportJournalCreation) Many2One() *Many2One {
	return NewMany2One(absijc.Id.Get(), "")
}

// CreateAccountBankStatementImportJournalCreation creates a new account.bank.statement.import.journal.creation model and returns its id.
func (c *Client) CreateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) (int64, error) {
	ids, err := c.CreateAccountBankStatementImportJournalCreations([]*AccountBankStatementImportJournalCreation{absijc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementImportJournalCreation creates a new account.bank.statement.import.journal.creation model and returns its id.
func (c *Client) CreateAccountBankStatementImportJournalCreations(absijcs []*AccountBankStatementImportJournalCreation) ([]int64, error) {
	var vv []interface{}
	for _, v := range absijcs {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementImportJournalCreationModel, vv, nil)
}

// UpdateAccountBankStatementImportJournalCreation updates an existing account.bank.statement.import.journal.creation record.
func (c *Client) UpdateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) error {
	return c.UpdateAccountBankStatementImportJournalCreations([]int64{absijc.Id.Get()}, absijc)
}

// UpdateAccountBankStatementImportJournalCreations updates existing account.bank.statement.import.journal.creation records.
// All records (represented by ids) will be updated by absijc values.
func (c *Client) UpdateAccountBankStatementImportJournalCreations(ids []int64, absijc *AccountBankStatementImportJournalCreation) error {
	return c.Update(AccountBankStatementImportJournalCreationModel, ids, absijc, nil)
}

// DeleteAccountBankStatementImportJournalCreation deletes an existing account.bank.statement.import.journal.creation record.
func (c *Client) DeleteAccountBankStatementImportJournalCreation(id int64) error {
	return c.DeleteAccountBankStatementImportJournalCreations([]int64{id})
}

// DeleteAccountBankStatementImportJournalCreations deletes existing account.bank.statement.import.journal.creation records.
func (c *Client) DeleteAccountBankStatementImportJournalCreations(ids []int64) error {
	return c.Delete(AccountBankStatementImportJournalCreationModel, ids)
}

// GetAccountBankStatementImportJournalCreation gets account.bank.statement.import.journal.creation existing record.
func (c *Client) GetAccountBankStatementImportJournalCreation(id int64) (*AccountBankStatementImportJournalCreation, error) {
	absijcs, err := c.GetAccountBankStatementImportJournalCreations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*absijcs)[0]), nil
}

// GetAccountBankStatementImportJournalCreations gets account.bank.statement.import.journal.creation existing records.
func (c *Client) GetAccountBankStatementImportJournalCreations(ids []int64) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.Read(AccountBankStatementImportJournalCreationModel, ids, nil, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreation finds account.bank.statement.import.journal.creation record by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreation(criteria *Criteria) (*AccountBankStatementImportJournalCreation, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, NewOptions().Limit(1), absijcs); err != nil {
		return nil, err
	}
	return &((*absijcs)[0]), nil
}

// FindAccountBankStatementImportJournalCreations finds account.bank.statement.import.journal.creation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreations(criteria *Criteria, options *Options) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, options, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
}

// FindAccountBankStatementImportJournalCreationId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
