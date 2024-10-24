var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Winardi_1004@tcp(127.0.0.1:3306)/store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}