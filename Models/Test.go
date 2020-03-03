package Models
type Test struct{
	TestID int `gorm:"column:test_id;AUTO_INCREMENT;PRIMARY_KEY" form:"test_id"`
	TestName string `gorm:"column:test_name;type:varchar(50)" form:"test_name"`

}
