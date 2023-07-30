package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     uint `gorm:"column:problem_id;type:varchar(36);" json:"problemId"` // 问题的ID
	CategoryId    uint
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"` // 关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
