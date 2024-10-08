package odoo

// Website represents website model.
type Website struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
	AuthSignupUninvited          *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutoRedirectLang             *Bool      `xmlrpc:"auto_redirect_lang,omitempty"`
	CdnActivated                 *Bool      `xmlrpc:"cdn_activated,omitempty"`
	CdnFilters                   *String    `xmlrpc:"cdn_filters,omitempty"`
	CdnUrl                       *String    `xmlrpc:"cdn_url,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryGroupIds              *Relation  `xmlrpc:"country_group_ids,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmDefaultTeamId             *Many2One  `xmlrpc:"crm_default_team_id,omitempty"`
	CrmDefaultUserId             *Many2One  `xmlrpc:"crm_default_user_id,omitempty"`
	DefaultLangId                *Many2One  `xmlrpc:"default_lang_id,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	Domain                       *String    `xmlrpc:"domain,omitempty"`
	Favicon                      *String    `xmlrpc:"favicon,omitempty"`
	ForumsCount                  *Int       `xmlrpc:"forums_count,omitempty"`
	GoogleAnalyticsKey           *String    `xmlrpc:"google_analytics_key,omitempty"`
	GoogleManagementClientId     *String    `xmlrpc:"google_management_client_id,omitempty"`
	GoogleManagementClientSecret *String    `xmlrpc:"google_management_client_secret,omitempty"`
	GoogleMapsApiKey             *String    `xmlrpc:"google_maps_api_key,omitempty"`
	HomepageId                   *Many2One  `xmlrpc:"homepage_id,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	KarmaProfileMin              *Int       `xmlrpc:"karma_profile_min,omitempty"`
	LanguageIds                  *Relation  `xmlrpc:"language_ids,omitempty"`
	Logo                         *String    `xmlrpc:"logo,omitempty"`
	MenuId                       *Many2One  `xmlrpc:"menu_id,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omitempty"`
	SocialDefaultImage           *String    `xmlrpc:"social_default_image,omitempty"`
	SocialFacebook               *String    `xmlrpc:"social_facebook,omitempty"`
	SocialGithub                 *String    `xmlrpc:"social_github,omitempty"`
	SocialInstagram              *String    `xmlrpc:"social_instagram,omitempty"`
	SocialLinkedin               *String    `xmlrpc:"social_linkedin,omitempty"`
	SocialTwitter                *String    `xmlrpc:"social_twitter,omitempty"`
	SocialYoutube                *String    `xmlrpc:"social_youtube,omitempty"`
	SpecificUserAccount          *Bool      `xmlrpc:"specific_user_account,omitempty"`
	ThemeId                      *Many2One  `xmlrpc:"theme_id,omitempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteFormEnableMetadata    *Bool      `xmlrpc:"website_form_enable_metadata,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// Websites represents array of website model.
type Websites []Website

// WebsiteModel is the odoo model name.
const WebsiteModel = "website"

// Many2One convert Website to *Many2One.
func (w *Website) Many2One() *Many2One {
	return NewMany2One(w.Id.Get(), "")
}

// CreateWebsite creates a new website model and returns its id.
func (c *Client) CreateWebsite(w *Website) (int64, error) {
	ids, err := c.CreateWebsites([]*Website{w})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsite creates a new website model and returns its id.
func (c *Client) CreateWebsites(ws []*Website) ([]int64, error) {
	var vv []interface{}
	for _, v := range ws {
		vv = append(vv, v)
	}
	return c.Create(WebsiteModel, vv, nil)
}

// UpdateWebsite updates an existing website record.
func (c *Client) UpdateWebsite(w *Website) error {
	return c.UpdateWebsites([]int64{w.Id.Get()}, w)
}

// UpdateWebsites updates existing website records.
// All records (represented by ids) will be updated by w values.
func (c *Client) UpdateWebsites(ids []int64, w *Website) error {
	return c.Update(WebsiteModel, ids, w, nil)
}

// DeleteWebsite deletes an existing website record.
func (c *Client) DeleteWebsite(id int64) error {
	return c.DeleteWebsites([]int64{id})
}

// DeleteWebsites deletes existing website records.
func (c *Client) DeleteWebsites(ids []int64) error {
	return c.Delete(WebsiteModel, ids)
}

// GetWebsite gets website existing record.
func (c *Client) GetWebsite(id int64) (*Website, error) {
	ws, err := c.GetWebsites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ws)[0]), nil
}

// GetWebsites gets website existing records.
func (c *Client) GetWebsites(ids []int64) (*Websites, error) {
	ws := &Websites{}
	if err := c.Read(WebsiteModel, ids, nil, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsite finds website record by querying it with criteria.
func (c *Client) FindWebsite(criteria *Criteria) (*Website, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, NewOptions().Limit(1), ws); err != nil {
		return nil, err
	}
	return &((*ws)[0]), nil
}

// FindWebsites finds website records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsites(criteria *Criteria, options *Options) (*Websites, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, options, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsiteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteModel, criteria, options)
}

// FindWebsiteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
