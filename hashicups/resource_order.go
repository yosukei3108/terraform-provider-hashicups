package hashicups

import (
	"context"
	"strconv"

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
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Define Coffee Schema
						"coffee": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								// ** | Coffee attributes
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"teaser": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"price": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"image": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

// resourceOrderCreate creates a new HashiCups order
func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Retrieve API client from meta parameter
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Map the order schema.Resource to []hc.OrderItems{}
	items := d.Get("items").([]interface{})
	ois := []hc.OrderItem{}

	for _, item := range items {
		i := item.(map[string]interface{})

		co := i["coffee"].([]interface{})[0]
		coffee := co.(map[string]interface{})

		oi := hc.OrderItem{
			Coffee: hc.Coffee{
				ID: coffee["id"].(int),
			},
			Quantity: i["quantity"].(int),
		}

		ois = append(ois, oi)
	}

	// Invoke the CreateOrder function on the HashiCups client
	o, err := c.CreateOrder(ois)
	if err != nil {
		return diag.FromErr(err)
	}

	// ** | Set order ID as resource ID
	d.SetId(strconv.Itoa(o.ID))

	// Map response (hc.Order) to order schema.Resource (done through resourceOrderRead)
	resourceOrderRead(ctx, d, m)

	return diags
}

// resourceOrderRead retrieves information for the specified HashiCups order and
// maps it to the order schema.Resource
func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Retrieve API client from meta parameter
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve order ID
	orderID := d.Id()

	// Invoke the GetOrder function on the HashiCups client
	order, err := c.GetOrder(orderID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Maps response (hc.Order) to order schema.Resource through flattening functions
	orderItems := flattenOrderItems(&order.Items)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// resourceOrderUpdate detects whether there is a difference between the state
// and configuration. If there is, it will update the HashiCups order.
func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Retrieve API client from meta parameter
	c := m.(*hc.Client)

	// Retrieve order ID
	orderID := d.Id()

	// Detect whether "items" has been changed
	if d.HasChange("items") {
		items := d.Get("items").([]interface{})
		ois := []hc.OrderItem{}

		// ** | Map the order schema.Resource to []hc.OrderItems{}
		for _, item := range items {
			i := item.(map[string]interface{})

			co := i["coffee"].([]interface{})[0]
			coffee := co.(map[string]interface{})

			oi := hc.OrderItem{
				Coffee: hc.Coffee{
					ID: coffee["id"].(int),
				},
				Quantity: i["quantity"].(int),
			}
			ois = append(ois, oi)
		}

		// Invoke the UpdateOrder function on the HashiCups client (this should only be invoked when "items" has been changed)
		_, err := c.UpdateOrder(orderID, ois)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Map response (hc.Order) to order schema.Resource (done through resourceOrderRead)
	return resourceOrderRead(ctx, d, m)
}

// resourceOrderDelete deletes the specified HashiCups order
func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Retrieve API client from meta parameter
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve order ID
	orderID := d.Id()

	// Invoke the UpdateOrder function on the HashiCups client
	err := c.DeleteOrder(orderID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// flattenOrderItems maps a *[]hc.OrderItem object to a object of type []interface{}
// A hc.orderItem contains of an hc.Coffee and a quantity. The flattenCoffee function
// can be used to map hc.Coffee to a []inferface{}
// If orderItems is empty, the function should return an empty []interface{}
func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems), len(*orderItems))

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})

			oi["coffee"] = flattenCoffee(orderItem.Coffee)
			oi["quantity"] = orderItem.Quantity

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

// flattenCoffee maps a hc.Coffee object to a object of type []interface{}
// The coffee object must be mapped to a []interface{} because the coffee
// schema.Schema is defined as a list containing a single coffee object
//
// This approach was taken because you cannot currently nest objects in schema.Schema
func flattenCoffee(coffee hc.Coffee) []interface{} {
	c := make(map[string]interface{})

	// ** | Map Coffee attributes
	c["id"] = coffee.ID
	c["name"] = coffee.Name
	c["teaser"] = coffee.Teaser
	c["description"] = coffee.Description
	c["price"] = coffee.Price
	c["image"] = coffee.Image

	return []interface{}{c}
}
