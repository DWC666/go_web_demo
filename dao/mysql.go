package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web_demo/models"
)

var(
	DB *gorm.DB
)

func InitMysql()(err error){
	des := "root:admin@(127.0.0.1:3306)/godb1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", des)
	if err != nil {
		return
	}

	return DB.DB().Ping()
}


func InitModel(){
	//模型绑定
	DB.AutoMigrate(&models.Todo{})
}


func Close(){
	DB.Close()
}

func CreateTodo(todo *models.Todo)(err error) {
	err = DB.Create(&todo).Error
	return
}


func GetTodoList()(todoList []models.Todo, err error) {
	err = DB.Find(&todoList).Error
	return
}

func GetTodoById(id string)(todo *models.Todo, err error){
	todo = new(models.Todo) //？
	if err := DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *models.Todo)(err error){
	err = DB.Save(todo).Error
	return
}


func DeleteTodoById(id string)(err error){
	err = DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return
}
