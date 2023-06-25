package ecs_invocation

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EcsInvocation can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_invocation.default ivk-ychnxnm45dl8j0mm****
```

*/

func ResourceVolcengineEcsInvocation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsInvocationCreate,
		Read:   resourceVolcengineEcsInvocationRead,
		Delete: resourceVolcengineEcsInvocationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"command_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The command id of the ecs invocation.",
			},
			"instance_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of ECS instance IDs.",
			},
			"invocation_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the ecs invocation.",
			},
			"invocation_description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the ecs invocation.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.",
			},
			"working_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.",
			},
			"timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(10, 600),
				Description:  "The timeout of the ecs command. Valid value range: 10-600. When this field is not specified, use the value of the field with the same name in ecs command as the default value.",
			},
			"repeat_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Once",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Once",
					"Rate",
					"Fixed",
				}, false),
				Description: "The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.",
			},
			"frequency": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Once" {
						return true
					}
					return false
				},
				Description: "The frequency of the ecs invocation. This field is valid when the value of the repeat_mode field is `Rate` or `Fixed`.",
			},
			"launch_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Rate" {
						return false
					}
					return true
				},
				Description: "The launch time of the ecs invocation. RFC3339 format. This field is valid when the value of the repeat_mode field is `Rate`.",
			},
			"recurrence_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Rate" {
						return false
					}
					return true
				},
				Description: "The recurrence end time of the ecs invocation. RFC3339 format. This field is valid when the value of the repeat_mode field is `Rate`.",
			},

			"invocation_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the ecs invocation.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The start time of the ecs invocation.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The end time of the ecs invocation.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsInvocationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on creating ecs invocation %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInvocationRead(d, meta)
}

func resourceVolcengineEcsInvocationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on reading ecs invocation %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInvocationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on deleting ecs invocation %q, %s", d.Id(), err)
	}
	return err
}
