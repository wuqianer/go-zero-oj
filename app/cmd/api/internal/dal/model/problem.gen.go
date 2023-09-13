// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameProblem = "problem"

// Problem mapped from table <problem>
type Problem struct {
	ID         int32          `gorm:"column:id;primaryKey" json:"id"`
	Identity   string         `gorm:"column:identity;comment:唯一标识" json:"identity"`          // 唯一标识
	Title      string         `gorm:"column:title;comment:标题" json:"title"`                  // 标题
	MaxRuntime int32          `gorm:"column:max_runtime;comment:最大运行时间" json:"max_runtime"`  // 最大运行时间
	MaxMem     int32          `gorm:"column:max_mem;comment:最大运行内存" json:"max_mem"`          // 最大运行内存
	Content    string         `gorm:"column:content;comment:题目内容" json:"content"`            // 题目内容
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`      // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:修改时间" json:"updated_at"`      // 修改时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间(软删除)" json:"deleted_at"` // 删除时间(软删除)
}

// TableName Problem's table name
func (*Problem) TableName() string {
	return TableNameProblem
}
