package local

import (
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceLocalDirectory() *schema.Resource {
	return &schema.Resource{
		Create: resourceLocalDirectoryCreate,
		Read:   resourceLocalDirectoryRead,
		Update: resourceLocalDirectoryUpdate,
		Delete: resourceLocalDirectoryDelete,

		Schema: map[string]*schema.Schema{
			"directory": {
				Type:        schema.TypeString,
				Description: "Path to the directory",
				Required:    true,
				ForceNew:    true,
			},
			"directory_permission": {
				Type:         schema.TypeString,
				Description:  "Permissions to set for directories created",
				Optional:     true,
				ForceNew:     false,
				Default:      "0777",
				ValidateFunc: validateMode,
			},
		},
	}
}

func resourceLocalDirectoryRead(d *schema.ResourceData, _ interface{}) error {
	// If the output directory doesn't exist, mark the resource for creation.
	wantedDirectory := d.Get("directory").(string)
	dirInfo, err := os.Stat(wantedDirectory)
	switch {
	case os.IsNotExist(err):
		d.SetId("")
		return nil
	case err != nil:
		return err
	}
	d.SetId(wantedDirectory)

	// The directory might have been modified externally and we might have to reconcile.
	dirPermission := dirInfo.Mode().Perm()
	log.Printf("[INFO] wanted %d, current %d", d.Get("directory_permission"), dirPermission)
	d.Set("directory_permission", dirPermission)
	return nil
}

func resourceLocalDirectoryUpdate(d *schema.ResourceData, _ interface{}) error {
	wantedDirectory := d.Get("directory").(string)
	wantedPermission := os.FileMode(d.Get("directory_permission").(int))

	if dirInfo, _ := os.Stat(wantedDirectory); dirInfo.Mode().Perm() != wantedPermission {
		if err := os.Chmod(wantedDirectory, wantedPermission); err != nil {
			log.Printf("[ERROR] error trying to modify permissions of directory %s to %d", wantedDirectory, wantedPermission)
			return err
		}
	}

	d.Set("directory_permission", wantedPermission)

	return nil
}

func resourceLocalDirectoryCreate(d *schema.ResourceData, _ interface{}) error {
	wantedDirectory := d.Get("directory").(string)
	wantedPermissionStr := d.Get("directory_permission").(string)
	wantedPermissionInt, _ := strconv.ParseInt(wantedPermissionStr, 8, 64)
	wantedPermission := os.FileMode(wantedPermissionInt)

	_, errStat := os.Stat(wantedDirectory)
	if errStat != nil && !os.IsNotExist(errStat) {
		log.Printf("[ERROR] error trying to stat directory %s", wantedDirectory)
		return errStat
	}

	if os.IsNotExist(errStat) {
		if err := os.MkdirAll(wantedDirectory, wantedPermission); err != nil {
			log.Printf("[ERROR] error trying to create directory %s", wantedDirectory)
			return err
		}
	}

	if errStat == nil {
		dirInfo, err := os.Stat(wantedDirectory)
		if err != nil {
			return err
		}
		if dirInfo.Mode().Perm() != wantedPermission {
			if err := os.Chmod(wantedDirectory, wantedPermission); err != nil {
				log.Printf("[ERROR] error trying to modify permissions of directory %s to %d", wantedDirectory, wantedPermission)
				return err
			}
		}
	}

	d.Set("directory_permission", wantedPermission)
	d.SetId(wantedDirectory)

	return nil
}

func resourceLocalDirectoryDelete(d *schema.ResourceData, _ interface{}) error {
	os.Remove(d.Get("directory").(string))
	return nil
}
