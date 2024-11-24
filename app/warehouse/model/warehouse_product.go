package model

import (
	"errors"
	"gorm.io/gorm"
)

// WarehouseProduct 仓库商品库存
type WarehouseProduct struct {
	WarehouseID uint   `gorm:"primaryKey;comment:'仓库ID'"`
	ProductID   uint   `gorm:"primaryKey;comment:'商品ID'"`
	Num         uint   `gorm:"not null;comment:'库存数量'"`                  // 库存数量
	Unit        string `gorm:"type:varchar(20);not null;comment:'库存单位'"` // 库存的单位
}

// ReduceStock 减少库存
func ReduceStock(db *gorm.DB, warehouseID uint, productID uint, reduceNum uint) error {
	var product WarehouseProduct

	// 事务操作
	err := db.Transaction(func(tx *gorm.DB) error {
		// 查询当前库存
		if err := tx.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&product).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("库存商品未找到")
			}
			return err
		}

		// 检查库存是否足够
		if product.Num < reduceNum {
			return errors.New("库存不足")
		}

		// 减少库存，使用乐观锁
		if err := tx.Model(&WarehouseProduct{}).
			Where("warehouse_id = ? AND product_id = ? AND num >= ?", warehouseID, productID, reduceNum).
			Update("num", gorm.Expr("num - ?", reduceNum)).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
