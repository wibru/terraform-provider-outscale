package outscale

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func TestAccOutscaleVM_tags(t *testing.T) {
	var v fcu.Instance

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOutscaleVMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckInstanceConfigTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleVMExists("outscale_vm.foo", &v),
					// testAccCheckTags(&v.Tags, "foo", "bar"),
					// testAccCheckTags(&v.Tags, "#", ""),
				),
			},
		},
	})
}

// func testAccCheckTags(
// 	ts *[]*fcu.Tag, key string, value string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		m := tagsToMap(*ts)
// 		v, ok := m[key]
// 		if value != "" && !ok {
// 			return fmt.Errorf("Missing tag: %s", key)
// 		} else if value == "" && ok {
// 			return fmt.Errorf("Extra tag: %s", key)
// 		}
// 		if value == "" {
// 			return nil
// 		}

// 		if v != value {
// 			return fmt.Errorf("%s: bad value: %s", key, v)
// 		}

// 		return nil
// 	}
// }

const testAccCheckInstanceConfigTags = `
resource "outscale_vm" "foo" {
	image_id = "ami-8a6a0120"
	instance_type = "m1.small"
	tags {
		foo = "bar"
	}
}

resource "outscale_tag" "foo" {
	resource_ids = ["${outscale_vm.foo.id}"]
	tags {
		faz = "baz"
	}
}
`
