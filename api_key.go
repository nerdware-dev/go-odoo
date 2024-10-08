package odoo

// ApiKey represents api.key model.
type ApiKey struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *String   `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Key         *String   `xmlrpc:"key,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ApiKeys represents array of api.key model.
type ApiKeys []ApiKey

// ApiKeyModel is the odoo model name.
const ApiKeyModel = "api.key"

// Many2One convert ApiKey to *Many2One.
func (ak *ApiKey) Many2One() *Many2One {
	return NewMany2One(ak.Id.Get(), "")
}

// CreateApiKey creates a new api.key model and returns its id.
func (c *Client) CreateApiKey(ak *ApiKey) (int64, error) {
	ids, err := c.CreateApiKeys([]*ApiKey{ak})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApiKey creates a new api.key model and returns its id.
func (c *Client) CreateApiKeys(aks []*ApiKey) ([]int64, error) {
	var vv []interface{}
	for _, v := range aks {
		vv = append(vv, v)
	}
	return c.Create(ApiKeyModel, vv, nil)
}

// UpdateApiKey updates an existing api.key record.
func (c *Client) UpdateApiKey(ak *ApiKey) error {
	return c.UpdateApiKeys([]int64{ak.Id.Get()}, ak)
}

// UpdateApiKeys updates existing api.key records.
// All records (represented by ids) will be updated by ak values.
func (c *Client) UpdateApiKeys(ids []int64, ak *ApiKey) error {
	return c.Update(ApiKeyModel, ids, ak, nil)
}

// DeleteApiKey deletes an existing api.key record.
func (c *Client) DeleteApiKey(id int64) error {
	return c.DeleteApiKeys([]int64{id})
}

// DeleteApiKeys deletes existing api.key records.
func (c *Client) DeleteApiKeys(ids []int64) error {
	return c.Delete(ApiKeyModel, ids)
}

// GetApiKey gets api.key existing record.
func (c *Client) GetApiKey(id int64) (*ApiKey, error) {
	aks, err := c.GetApiKeys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aks)[0]), nil
}

// GetApiKeys gets api.key existing records.
func (c *Client) GetApiKeys(ids []int64) (*ApiKeys, error) {
	aks := &ApiKeys{}
	if err := c.Read(ApiKeyModel, ids, nil, aks); err != nil {
		return nil, err
	}
	return aks, nil
}

// FindApiKey finds api.key record by querying it with criteria.
func (c *Client) FindApiKey(criteria *Criteria) (*ApiKey, error) {
	aks := &ApiKeys{}
	if err := c.SearchRead(ApiKeyModel, criteria, NewOptions().Limit(1), aks); err != nil {
		return nil, err
	}
	return &((*aks)[0]), nil
}

// FindApiKeys finds api.key records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApiKeys(criteria *Criteria, options *Options) (*ApiKeys, error) {
	aks := &ApiKeys{}
	if err := c.SearchRead(ApiKeyModel, criteria, options, aks); err != nil {
		return nil, err
	}
	return aks, nil
}

// FindApiKeyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApiKeyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ApiKeyModel, criteria, options)
}

// FindApiKeyId finds record id by querying it with criteria.
func (c *Client) FindApiKeyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApiKeyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
