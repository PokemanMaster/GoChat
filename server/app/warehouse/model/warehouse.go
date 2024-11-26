package model

// Warehouse 仓库
type Warehouse struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	CityID  uint   `gorm:"not null;comment:'城市ID';index:idx_city_id"` // 城市编号；仓库在哪个城市，有索引，提高查询速度。
	Address string `gorm:"type:varchar(200);not null;comment:'地址'"`   // 仓库地址
	Tel     string `gorm:"type:varchar(20);not null;comment:'电话'"`    // 联系电话
}
