package model

// SpecParam 参数表
type SpecParam struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null;comment:'主键'"`
	SpgID     uint   `gorm:"type:int unsigned;not null;index:idx_spg_id;comment:'品类编号'"`
	SppID     uint   `gorm:"type:int unsigned;not null;index:idx_spp_id;comment:'参数编号'"`
	Name      string `gorm:"type:varchar(200);not null;comment:'参数名称'"`
	Numeric   bool   `gorm:"type:tinyint(1);not null;comment:'是否为数字参数'"`
	Unit      string `gorm:"type:varchar(200);comment:'单位（量词语）'"`
	Generic   bool   `gorm:"type:tinyint(1);not null;comment:'是否为通用参数'"`
	Searching bool   `gorm:"not null;comment:'是否用于通用搜索'"`
	Segments  string `gorm:"type:varchar(500);comment:'参数值'"`
}
