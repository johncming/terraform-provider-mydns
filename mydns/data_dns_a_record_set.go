package mydns

import (
	"fmt"
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/miekg/dns"
)

func dataSourceDnsARecordSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDnsARecordSetRead,
		Schema: map[string]*schema.Schema{
			"resolver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "223.5.5.5:53",
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceDnsARecordSetRead(d *schema.ResourceData, meta interface{}) error {
	host := d.Get("host").(string)

	resolver := d.Get("resolver").(string)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(host), dns.TypeA)
	resposeMsg, err := dns.Exchange(m, resolver)
	if err != nil {
		return fmt.Errorf("error looking up A records for %q: %s", host, err)
	}

	addrs := make([]string, 0)

	for _, rr := range resposeMsg.Answer {
		ip := rr.(*dns.A).A
		addrs = append(addrs, ip.String())
	}

	sort.Strings(addrs)

	d.Set("addrs", addrs)
	d.SetId(host)

	return nil
}
