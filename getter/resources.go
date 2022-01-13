package getter

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gogetter "github.com/hashicorp/go-getter"
	"fmt"
	"strings"
	"crypto/sha1"
	"encoding/hex"
	

)

func resourceGetter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: false,
				// ForceNew:  true,
			},

			"dest": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: false,
				// ForceNew:  true,
			},
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	url := d.Get("url").(string)
	dest := d.Get("dest").(string)
	// insecure := d.Get("insecure").(bool)
	opts := []gogetter.ClientOption{}
	// if insecure {
	// 	opts = append(opts, gogetter.WithInsecure())
	// }
	
	client := &gogetter.Client{
		Ctx:     ctx,
		Src:     url,
		Dst:     dest,
		Mode:    gogetter.ClientModeAny,
		Options: opts,
	}
	err := client.Get()
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error downloading: %s", err))
	}
	d.SetId(hashForState("getter_" + url + dest ))
	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func hashForState(value string) string {
	if value == "" {
		return ""
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(value)))
	return hex.EncodeToString(hash[:])
}

