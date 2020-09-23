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

// resourceOrderCreate creates a new HashiCups order
func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 1. Retrieve API client from meta parameter

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// 2. Map the order schema.Resource to []hc.OrderItems{}

	// 3. Invoke the CreateOrder function on the HashiCups client

	// 4. Set order ID as resource ID

	// 5. Map response (hc.Order) to order schema.Resource (done through resourceOrderRead)
	resourceOrderRead(ctx, d, m)

	return diags
}

// resourceOrderRead retrieves information for the specified HashiCups order and
// maps it to the order schema.Resource
func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 1. Retrieve API client from meta parameter

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// 2. Retrieve order ID

	// 3. Invoke the GetOrder function on the HashiCups client

	// 4. Maps response (hc.Order) to order schema.Resource through flattening functions

	return diags
}

// resourceOrderUpdate detects whether there is a difference between the state
// and configuration. If there is, it will update the HashiCups order.
func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 1. Retrieve API client from meta parameter

	// 2. Retrieve order ID

	// 3. Map the order schema.Resource to []hc.OrderItems{} if changes to "items" are detected

	// 4. Invoke the UpdateOrder function on the HashiCups client

	// 5. Map response (hc.Order) to order schema.Resource (done through resourceOrderRead)
	return resourceOrderRead(ctx, d, m)
}

// resourceOrderDelete deletes the specified HashiCups order
func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 1. Retrieve API client from meta parameter

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// 2. Retrieve order ID

	// 3. Invoke the UpdateOrder function on the HashiCups client

	return diags
}

// flattenOrderItems maps a *[]hc.OrderItem object to a object of type []interface{}
// A hc.orderItem contains of an hc.Coffee and a quantity. The flattenCoffee function
// can be used to map hc.Coffee to a []inferface{}
// If orderItems is empty, the function should return an empty []interface{}
func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
	return make([]interface{}, 0)
}

// flattenCoffee maps a hc.Coffee object to a object of type []interface{}
// The coffee object must be mapped to a []interface{} because the coffee
// schema.Schema is defined as a list containing a single coffee object
//
// This approach was taken because you cannot currently nest objects in schema.Schema
func flattenCoffee(coffee hc.Coffee) []interface{} {
	return make([]interface{}, 0)
}
