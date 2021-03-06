package model

const (
	NotifyGrade = 1 << iota
	NotifyLibrary
	NotifyExam
)

// 默认有10位1，当设置项超过10个之后需要重新设置
const NotifyAll = (1 << 10) - 1

const (
	// RegularStudent 本科
	RegularStudent = 0
	// GraduateStudent 研究生
	GraduateStudent = 1
)

type UserConfig struct {
	Model
	UserID   uint
	Notify   int // Notify 通知设置
	UserType int `gorm:"default:0"` // 用户类型
}
