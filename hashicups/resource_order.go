package hashicups

import (
	"context"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema:        map[string]*schema.Schema{},
	}
}

// resourceOrderCreate creates a new HashiCups order by doing the following:
// 1) maps the order schema.Resource to []hc.OrderItems{}
// 2) invokes the CreateOrder function on the HashiCups client
// 3) maps response (hc.Order) to order schema.Resource (similar to resourceOrderRead)
func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceOrderRead(ctx, d, m)

	return diags
}

// resourceOrderRead retrieves information for the specified HashiCups order and
// maps it to the order schema.Resource by doing the following:
// 1) retrieves order ID
// 2) invokes the GetOrder function on the HashiCups client
// 3) maps response (hc.Order) to order schema.Resource
// TIP: since hc.Order is a nested object, use flattening functions to map to schema
func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

// resourceOrderUpdate detects whether there is a difference between the state
// and configuration. If there is, it will update the HashiCups order by
// by doing the following:
// 1) retrieves order ID
// 2) detects changes in "items"
// 3) maps the order schema.Resource to []hc.OrderItems{}
// 4) invokes the UpdateOrder function on the HashiCups client
// 5) updates "last_updated" field to current datetime
// 6) maps response (hc.Order) to order schema.Resource (similar to resourceOrderRead)
func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceOrderRead(ctx, d, m)
}

// resourceOrderDelete deletes the specified HashiCups order by doing the following:
// 1) retrieves order ID
// 2) invokes the DeleteOrder function on the HashiCups client
func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

// flattenOrderItems maps a *[]hc.OrderItem object to a object of type []interface{}
// A hc.orderItem contains of an hc.Coffee and a quantity. The flattenCoffee function
// can be used to map hc.Coffee to a []inferface{}
// If orderItems is empty, the function should return an empty []interface{}
func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
}

// flattenCoffee maps a hc.Coffee object to a object of type []interface{}
// The coffee object must be mapped to a []interface{} because the coffee
// schema.Schema is defined as a list containing a single coffee object
//
// This approach was taken because you cannot currently nest objects in schema.Schema
func flattenCoffee(coffee hc.Coffee) []interface{} {
}
