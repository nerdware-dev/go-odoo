package odoo

import (
	"fmt"
)

// ProjectTaskCreateSaleOrder represents project.task.create.sale.order model.
type ProjectTaskCreateSaleOrder struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omptempty"`
	PriceUnit   *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	TaskId      *Many2One `xmlrpc:"task_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskCreateSaleOrders represents array of project.task.create.sale.order model.
type ProjectTaskCreateSaleOrders []ProjectTaskCreateSaleOrder

// ProjectTaskCreateSaleOrderModel is the odoo model name.
const ProjectTaskCreateSaleOrderModel = "project.task.create.sale.order"

// Many2One convert ProjectTaskCreateSaleOrder to *Many2One.
func (ptcso *ProjectTaskCreateSaleOrder) Many2One() *Many2One {
	return NewMany2One(ptcso.Id.Get(), "")
}

// CreateProjectTaskCreateSaleOrder creates a new project.task.create.sale.order model and returns its id.
func (c *Client) CreateProjectTaskCreateSaleOrder(ptcso *ProjectTaskCreateSaleOrder) (int64, error) {
	ids, err := c.CreateProjectTaskCreateSaleOrders([]*ProjectTaskCreateSaleOrder{ptcso})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskCreateSaleOrder creates a new project.task.create.sale.order model and returns its id.
func (c *Client) CreateProjectTaskCreateSaleOrders(ptcsos []*ProjectTaskCreateSaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcsos {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskCreateSaleOrderModel, vv)
}

// UpdateProjectTaskCreateSaleOrder updates an existing project.task.create.sale.order record.
func (c *Client) UpdateProjectTaskCreateSaleOrder(ptcso *ProjectTaskCreateSaleOrder) error {
	return c.UpdateProjectTaskCreateSaleOrders([]int64{ptcso.Id.Get()}, ptcso)
}

// UpdateProjectTaskCreateSaleOrders updates existing project.task.create.sale.order records.
// All records (represented by ids) will be updated by ptcso values.
func (c *Client) UpdateProjectTaskCreateSaleOrders(ids []int64, ptcso *ProjectTaskCreateSaleOrder) error {
	return c.Update(ProjectTaskCreateSaleOrderModel, ids, ptcso)
}

// DeleteProjectTaskCreateSaleOrder deletes an existing project.task.create.sale.order record.
func (c *Client) DeleteProjectTaskCreateSaleOrder(id int64) error {
	return c.DeleteProjectTaskCreateSaleOrders([]int64{id})
}

// DeleteProjectTaskCreateSaleOrders deletes existing project.task.create.sale.order records.
func (c *Client) DeleteProjectTaskCreateSaleOrders(ids []int64) error {
	return c.Delete(ProjectTaskCreateSaleOrderModel, ids)
}

// GetProjectTaskCreateSaleOrder gets project.task.create.sale.order existing record.
func (c *Client) GetProjectTaskCreateSaleOrder(id int64) (*ProjectTaskCreateSaleOrder, error) {
	ptcsos, err := c.GetProjectTaskCreateSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptcsos != nil && len(*ptcsos) > 0 {
		return &((*ptcsos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.create.sale.order not found", id)
}

// GetProjectTaskCreateSaleOrders gets project.task.create.sale.order existing records.
func (c *Client) GetProjectTaskCreateSaleOrders(ids []int64) (*ProjectTaskCreateSaleOrders, error) {
	ptcsos := &ProjectTaskCreateSaleOrders{}
	if err := c.Read(ProjectTaskCreateSaleOrderModel, ids, nil, ptcsos); err != nil {
		return nil, err
	}
	return ptcsos, nil
}

// FindProjectTaskCreateSaleOrder finds project.task.create.sale.order record by querying it with criteria.
func (c *Client) FindProjectTaskCreateSaleOrder(criteria *Criteria) (*ProjectTaskCreateSaleOrder, error) {
	ptcsos := &ProjectTaskCreateSaleOrders{}
	if err := c.SearchRead(ProjectTaskCreateSaleOrderModel, criteria, NewOptions().Limit(1), ptcsos); err != nil {
		return nil, err
	}
	if ptcsos != nil && len(*ptcsos) > 0 {
		return &((*ptcsos)[0]), nil
	}
	return nil, fmt.Errorf("project.task.create.sale.order was not found with criteria %v", criteria)
}

// FindProjectTaskCreateSaleOrders finds project.task.create.sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskCreateSaleOrders(criteria *Criteria, options *Options) (*ProjectTaskCreateSaleOrders, error) {
	ptcsos := &ProjectTaskCreateSaleOrders{}
	if err := c.SearchRead(ProjectTaskCreateSaleOrderModel, criteria, options, ptcsos); err != nil {
		return nil, err
	}
	return ptcsos, nil
}

// FindProjectTaskCreateSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskCreateSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskCreateSaleOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskCreateSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskCreateSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskCreateSaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.create.sale.order was not found with criteria %v and options %v", criteria, options)
}
