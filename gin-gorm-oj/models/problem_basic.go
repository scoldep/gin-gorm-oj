package models

import (
	"gorm.io/gorm"
)

// 问题创建的步骤
// 1.ProblemBasic的创建
// 2.关联问题分类
// 3.关联测试案例
// 现在少一张测试案例的表，需要新建
// 一个问题有多个测试案例

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`  // 问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`                // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`       // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`           // 文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"maxRuntime"` // 最大的运行时间
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`        // 最大的运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity"`    // 关联测试用例表
	PassNum           int64              `gorm:"column:pass_num;type:int(11);" json:"pass_num"`      // 通过的问题个数
	SubmitNum         int64              `gorm:"column:pass_num;type:int(11);" json:"submit_num"`    // 提交次数
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (select cb.id from category_basic cb where cb.identity = ?)", categoryIdentity)
	}
	return tx
}
