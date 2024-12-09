package serializer

import (
	Mcategory "github.com/PokemanMaster/GoChat/v1/server/app/category/model"
	Mproduct "github.com/PokemanMaster/GoChat/v1/server/app/product/model"
)

// ProductSerialization 商品序列化
type ProductSerialization struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`        // 商品标题
	SubTitle   string `json:"subTitle"`    // 商品副标题
	CategoryID uint   `json:"category_id"` // 分类ID
	BrandID    uint   `json:"brandID"`     // 一些散装的可能没有品牌，比如花生之类的
	SpgID      uint   `json:"spgID"`       // 电子产品可能包含：电脑、电视等等，相当于二级索引
	Saleable   bool   `json:"saleable"`    // 是否上架
	Valid      bool   `json:"valid"`       // 是否有效
}

// BuildProduct 序列化商品
func BuildProduct(item Mproduct.Product) ProductSerialization {
	return ProductSerialization{
		ID:         item.ID,
		CategoryID: item.CategoryID,
		Name:       item.Name,
		BrandID:    item.BrandID,
		Saleable:   item.Saleable,
	}
}

// BuildProducts 序列化商品列表
func BuildProducts(items []Mproduct.Product) (products []ProductSerialization) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ProductBrandSerialization 商品品牌
type ProductBrandSerialization struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"` // 品牌名唯一
	Image  string `json:"image"`
	Letter string `json:"letter"` // 由于需要按首字母搜索品牌，所以给一个索引
}

// BuildProductBrand 序列化商品品牌
func BuildProductBrand(item Mproduct.ProductBrand) ProductBrandSerialization {
	return ProductBrandSerialization{
		ID:     item.ID,
		Name:   item.Name,
		Image:  item.Image,
		Letter: item.Letter,
	}
}

// BuildProductBrands 序列化商品品牌
func BuildProductBrands(items []Mproduct.ProductBrand) (categories []ProductBrandSerialization) {
	for _, item := range items {
		category := BuildProductBrand(item)
		categories = append(categories, category)
	}
	return categories
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ProductCategorySerialization 分类序列化器
type ProductCategorySerialization struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`      //分类名称：可重复
	ParentID uint   `json:"parent_id"` //父节点：顶级节点没有父
	IfParent bool   `json:"if_parent"` //是否为父节点：如果有子，则为 false
	Sort     uint   `json:"sort"`      //排名指数：相当于搜索的权重
}

// BuildProductCategory 序列化分类
func BuildProductCategory(item Mcategory.ProductCategory) ProductCategorySerialization {
	return ProductCategorySerialization{
		ID:       item.ID,
		Name:     item.Name,
		ParentID: item.ParentID,
		IfParent: item.IfParent,
		Sort:     item.Sort,
	}
}

// BuildProductCategorys 序列化分类列表
func BuildProductCategorys(items []Mcategory.ProductCategory) (categories []ProductCategorySerialization) {
	for _, item := range items {
		category := BuildProductCategory(item)
		categories = append(categories, category)
	}
	return categories
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ProductParamSerialization 商品参数序列化
type ProductParamSerialization struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	Saleable  bool    `json:"saleable"`
	SoldCount uint    `json:"sold_count"`
}

// BuildProductParam 序列化商品图片
func BuildProductParam(item Mproduct.Product, item2 Mproduct.ProductParam) ProductParamSerialization {
	return ProductParamSerialization{
		ID:        item2.ID,
		Name:      item.Name,
		ProductID: item2.ProductID,
		Image:     item.Image,
		Price:     item2.Price,
		Saleable:  item.Saleable,
	}
}
