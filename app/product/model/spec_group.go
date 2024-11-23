package model

// SpecGroup 品类表
type SpecGroup struct {
	ID    uint   `gorm:"primaryKey;autoIncrement;not null;comment:'主键'"`
	SpgID uint   `gorm:"type:int unsigned;not null;uniqueIndex:unq_spg_id;comment:'品类编号'"`
	Name  string `gorm:"type:varchar(200);not null;uniqueIndex:unq_name;comment:'品类名称'"`
}
