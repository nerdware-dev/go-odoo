package odoo

// FollowupSend represents followup.send model.
type FollowupSend struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	InvalidAddresses  *Int      `xmlrpc:"invalid_addresses,omitempty"`
	InvalidPartnerIds *Relation `xmlrpc:"invalid_partner_ids,omitempty"`
	LettersQty        *Int      `xmlrpc:"letters_qty,omitempty"`
	PartnerIds        *Relation `xmlrpc:"partner_ids,omitempty"`
	SnailmailCost     *Float    `xmlrpc:"snailmail_cost,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// FollowupSends represents array of followup.send model.
type FollowupSends []FollowupSend

// FollowupSendModel is the odoo model name.
const FollowupSendModel = "followup.send"

// Many2One convert FollowupSend to *Many2One.
func (fs *FollowupSend) Many2One() *Many2One {
	return NewMany2One(fs.Id.Get(), "")
}

// CreateFollowupSend creates a new followup.send model and returns its id.
func (c *Client) CreateFollowupSend(fs *FollowupSend) (int64, error) {
	ids, err := c.CreateFollowupSends([]*FollowupSend{fs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFollowupSend creates a new followup.send model and returns its id.
func (c *Client) CreateFollowupSends(fss []*FollowupSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range fss {
		vv = append(vv, v)
	}
	return c.Create(FollowupSendModel, vv, nil)
}

// UpdateFollowupSend updates an existing followup.send record.
func (c *Client) UpdateFollowupSend(fs *FollowupSend) error {
	return c.UpdateFollowupSends([]int64{fs.Id.Get()}, fs)
}

// UpdateFollowupSends updates existing followup.send records.
// All records (represented by ids) will be updated by fs values.
func (c *Client) UpdateFollowupSends(ids []int64, fs *FollowupSend) error {
	return c.Update(FollowupSendModel, ids, fs, nil)
}

// DeleteFollowupSend deletes an existing followup.send record.
func (c *Client) DeleteFollowupSend(id int64) error {
	return c.DeleteFollowupSends([]int64{id})
}

// DeleteFollowupSends deletes existing followup.send records.
func (c *Client) DeleteFollowupSends(ids []int64) error {
	return c.Delete(FollowupSendModel, ids)
}

// GetFollowupSend gets followup.send existing record.
func (c *Client) GetFollowupSend(id int64) (*FollowupSend, error) {
	fss, err := c.GetFollowupSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fss)[0]), nil
}

// GetFollowupSends gets followup.send existing records.
func (c *Client) GetFollowupSends(ids []int64) (*FollowupSends, error) {
	fss := &FollowupSends{}
	if err := c.Read(FollowupSendModel, ids, nil, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFollowupSend finds followup.send record by querying it with criteria.
func (c *Client) FindFollowupSend(criteria *Criteria) (*FollowupSend, error) {
	fss := &FollowupSends{}
	if err := c.SearchRead(FollowupSendModel, criteria, NewOptions().Limit(1), fss); err != nil {
		return nil, err
	}
	return &((*fss)[0]), nil
}

// FindFollowupSends finds followup.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFollowupSends(criteria *Criteria, options *Options) (*FollowupSends, error) {
	fss := &FollowupSends{}
	if err := c.SearchRead(FollowupSendModel, criteria, options, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFollowupSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFollowupSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FollowupSendModel, criteria, options)
}

// FindFollowupSendId finds record id by querying it with criteria.
func (c *Client) FindFollowupSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FollowupSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
