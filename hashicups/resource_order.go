package hashicups

import (
	"strconv"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrderCreate,
		Read:   resourceOrderRead,
		Update: resourceOrderUpdate,
		Delete: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
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
										Type:     schema.TypeInt,
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

func resourceOrderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*hc.Client)

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

	o, err := c.CreateOrder(ois)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(o.ID))

	resourceOrderRead(d, m)

	return nil
}

func resourceOrderRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*hc.Client)

	orderID := d.Id()

	order, err := c.GetOrder(orderID)
	if err != nil {
		return err
	}

	orderItems := flattenOrderItems(&order.Items)
	if err := d.Set("items", orderItems); err != nil {
		return err
	}

	return nil
}

func resourceOrderUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceOrderRead(d, m)
}

func resourceOrderDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
