package version

import (
	"gorm.io/gorm"
	"runtime"
	"shimmer/cmd/migrate/migration/models"
	common "shimmer/common/models"
	"strconv"

	"shimmer/cmd/migrate/migration"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1653638869132Test)
}

func _1653638869132Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var list []models.SysMenu
		err := tx.Model(&models.SysMenu{}).Order("parent_id,menu_id").Find(&list).Error
		if err != nil {
			return err
		}
		for _, v := range list {
			if v.ParentId == 0 {
				v.Paths = "/0/" + strconv.Itoa(v.MenuId)
			} else {
				var e models.SysMenu
				err = tx.Model(&models.SysMenu{}).Where("menu_id=?", v.ParentId).First(&e).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						continue
					}
					return err
				}
				v.Paths = e.Paths + "/" + strconv.Itoa(v.MenuId)
			}
			err = tx.Model(&v).Update("paths", v.Paths).Error
			if err != nil {
				return err
			}
		}
		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
