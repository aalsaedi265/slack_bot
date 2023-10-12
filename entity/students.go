
package entity

type Student struct {
	Id           int    `gorm:"primary_key;auto_increment;not_null"`
	StudentName  string `json:"studentName"`
	Grade        string `json:"grade"`
	ClassDetails string `json:"classDetails"`
}