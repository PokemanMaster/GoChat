package model

// ProductBrand 商品品牌
type ProductBrand struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	Name   string `gorm:"type:varchar(200);not null;comment:'名称';uniqueIndex:unq_name"` // 品牌名唯一
	Image  string `gorm:"type:varchar(500);comment:'图片网址'"`
	Letter string `gorm:"type:char(1);not null;comment:'单位（量词语）';index:idx_letter"` // 由于需要按首字母搜索品牌，所以给一个索引
}
