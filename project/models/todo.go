package models

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// 增删改查相关的操作也放在这个model里面(具体操作)
// db.Create(&todo).Error
// db.Find(&todoList).Error
// db.Save(&todo)
