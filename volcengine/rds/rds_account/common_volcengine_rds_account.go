package rds_account

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var rdsAccountImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("account_name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
