package odoo

// ForumPostVote represents forum.post.vote model.
type ForumPostVote struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	ForumId     *Many2One  `xmlrpc:"forum_id,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	PostId      *Many2One  `xmlrpc:"post_id,omitempty"`
	RecipientId *Many2One  `xmlrpc:"recipient_id,omitempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omitempty"`
	Vote        *Selection `xmlrpc:"vote,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ForumPostVotes represents array of forum.post.vote model.
type ForumPostVotes []ForumPostVote

// ForumPostVoteModel is the odoo model name.
const ForumPostVoteModel = "forum.post.vote"

// Many2One convert ForumPostVote to *Many2One.
func (fpv *ForumPostVote) Many2One() *Many2One {
	return NewMany2One(fpv.Id.Get(), "")
}

// CreateForumPostVote creates a new forum.post.vote model and returns its id.
func (c *Client) CreateForumPostVote(fpv *ForumPostVote) (int64, error) {
	ids, err := c.CreateForumPostVotes([]*ForumPostVote{fpv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateForumPostVote creates a new forum.post.vote model and returns its id.
func (c *Client) CreateForumPostVotes(fpvs []*ForumPostVote) ([]int64, error) {
	var vv []interface{}
	for _, v := range fpvs {
		vv = append(vv, v)
	}
	return c.Create(ForumPostVoteModel, vv, nil)
}

// UpdateForumPostVote updates an existing forum.post.vote record.
func (c *Client) UpdateForumPostVote(fpv *ForumPostVote) error {
	return c.UpdateForumPostVotes([]int64{fpv.Id.Get()}, fpv)
}

// UpdateForumPostVotes updates existing forum.post.vote records.
// All records (represented by ids) will be updated by fpv values.
func (c *Client) UpdateForumPostVotes(ids []int64, fpv *ForumPostVote) error {
	return c.Update(ForumPostVoteModel, ids, fpv, nil)
}

// DeleteForumPostVote deletes an existing forum.post.vote record.
func (c *Client) DeleteForumPostVote(id int64) error {
	return c.DeleteForumPostVotes([]int64{id})
}

// DeleteForumPostVotes deletes existing forum.post.vote records.
func (c *Client) DeleteForumPostVotes(ids []int64) error {
	return c.Delete(ForumPostVoteModel, ids)
}

// GetForumPostVote gets forum.post.vote existing record.
func (c *Client) GetForumPostVote(id int64) (*ForumPostVote, error) {
	fpvs, err := c.GetForumPostVotes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fpvs)[0]), nil
}

// GetForumPostVotes gets forum.post.vote existing records.
func (c *Client) GetForumPostVotes(ids []int64) (*ForumPostVotes, error) {
	fpvs := &ForumPostVotes{}
	if err := c.Read(ForumPostVoteModel, ids, nil, fpvs); err != nil {
		return nil, err
	}
	return fpvs, nil
}

// FindForumPostVote finds forum.post.vote record by querying it with criteria.
func (c *Client) FindForumPostVote(criteria *Criteria) (*ForumPostVote, error) {
	fpvs := &ForumPostVotes{}
	if err := c.SearchRead(ForumPostVoteModel, criteria, NewOptions().Limit(1), fpvs); err != nil {
		return nil, err
	}
	return &((*fpvs)[0]), nil
}

// FindForumPostVotes finds forum.post.vote records by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPostVotes(criteria *Criteria, options *Options) (*ForumPostVotes, error) {
	fpvs := &ForumPostVotes{}
	if err := c.SearchRead(ForumPostVoteModel, criteria, options, fpvs); err != nil {
		return nil, err
	}
	return fpvs, nil
}

// FindForumPostVoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPostVoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ForumPostVoteModel, criteria, options)
}

// FindForumPostVoteId finds record id by querying it with criteria.
func (c *Client) FindForumPostVoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumPostVoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
