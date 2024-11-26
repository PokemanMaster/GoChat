package model

// 一个品牌可以有多个分类，一个分类可以分给多个品牌。多对多关系

// ProductCategoryBrand  分类与品牌关联表
type ProductCategoryBrand struct {
	CategoryID uint `gorm:"not null;comment:'分类 ID';primaryKey"`
	BrandID    uint `gorm:"not null;comment:'品牌 ID';primaryKey"`
}
