package model

// ProductCategory  商品分类
type ProductCategory struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	Name     string `gorm:"type:varchar(200);not null;comment:'分类名称'"`   //分类名称：可重复
	ParentID uint   `gorm:"comment:'上级分类ID';index:idx_parent_id"`        //父节点：顶级节点没有父
	IfParent bool   `gorm:"type:tinyint(1);not null;comment:'是否包含下级分类'"` //是否为父节点：如果有子，则为 false
	Sort     uint   `gorm:"not null;comment:'排名指数';index:idx_sort"`      //排名指数：相当于搜索的权重
}
