package local

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestLocalDirectory_Basic(t *testing.T) {
	var tests = []struct {
		directory            string
		directory_permission os.FileMode
		config               string
	}{
		{
			"local_directory",
			0750,
			`resource "local_directory" "directory" {
         directory            = "local_directory"
         directory_permission = 0750
      }`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run("", func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				Providers: testProviders,
				Steps: []resource.TestStep{
					{
						Config: test.config,
						Check: func(s *terraform.State) error {
							dirInfo, err := os.Stat(test.directory)
							if err != nil {
								return fmt.Errorf("config:\n%s\n,got: %s\n", test.config, err)
							}
							if dirInfo.Mode().Perm() != test.directory_permission {
								return fmt.Errorf("config:\n%s\ngot:\n%+v\nwant:\n%+v\n", test.config, dirInfo.Mode().Perm(), test.directory_permission)
							}
							return nil
						},
					},
				},
				CheckDestroy: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						if _, err := os.Stat(test.directory); os.IsNotExist(err) {
							return nil
						}
						return errors.New("local_directory did not get destroyed")
					},
				),
			})
		})
	}
}
