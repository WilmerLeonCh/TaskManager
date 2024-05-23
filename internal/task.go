package tasks

import "gorm.io/gorm"

type MTask struct {
	gorm.Model
	ID          int `gorm:"primaryKey,autoIncrement"`
	Name        string
	Description string
	Completed   bool
}

func Add(db *gorm.DB, task *MTask) *MTask {
	task.Completed = false
	db.Create(&task)
	return task
}

func GetAll(db *gorm.DB) []MTask {
	var tasks []MTask
	db.Find(&tasks)
	return tasks
}

func GetById(db *gorm.DB, id int) *MTask {
	var task MTask
	res := db.First(&task, id)
	if res.RowsAffected == 0 {
		return nil
	}
	return &task
}

func DeleteById(db *gorm.DB, id int) {
	var taskToDelete MTask
	// verify if the task exists
	existTask := db.First(&taskToDelete, id)
	if existTask.Error != nil {
		panic(existTask.Error)
	}
	res := db.Delete(&MTask{}, id)
	if res.Error != nil {
		panic(res.Error)
	}
}

func UpdateById(db *gorm.DB, id int, task *MTask) *MTask {
	var taskToUpdate *MTask
	// verify if the task exists
	existTask := db.First(&taskToUpdate, id)
	if existTask.Error != nil {
		panic(existTask.Error)
	}
	taskToUpdate = task
	resultSave := db.Save(&taskToUpdate)
	if resultSave.Error != nil {
		panic(resultSave.Error)
	}
	return taskToUpdate
}
