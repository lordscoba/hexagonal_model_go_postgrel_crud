package todo_service

import (
	"errors"
	"strconv"

	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage/postgresql"
	"gorm.io/gorm"
)

func CreateToDOService(data *models.Todo, Db *gorm.DB) (models.Todo, string, int, error) {

	err := postgresql.CreateOneRecord(Db, data)
	if err != nil {
		return models.Todo{}, "", 0, err
	}

	return *data, "", 0, nil
}

func GetToDosService(Db *gorm.DB) ([]models.Todo, string, int, error) {

	var todos []models.Todo

	err := postgresql.SelectAllFromDb(Db, "", &todos, nil)
	if err != nil {
		return []models.Todo{}, "", 0, err
	}

	return todos, "", 0, nil
}

func GetToDoService(id string, Db *gorm.DB) (models.Todo, string, int, error) {

	var todo models.Todo

	ID, err := strconv.Atoi(id)
	if err != nil {
		return models.Todo{}, "", 0, err
	}

	err, nilErr := postgresql.SelectOneFromDb(Db, &todo, "id = ?", uint(ID))

	if err != nil {
		return models.Todo{}, "", 0, err
	}

	if nilErr != nil {
		return models.Todo{}, "", 0, nilErr
	}

	return todo, "", 0, nil
}

func UpdateToDOService(id string, data *models.Todo, Db *gorm.DB) (models.Todo, string, int, error) {

	ID, err := strconv.Atoi(id)
	if err != nil {
		return models.Todo{}, "", 0, err
	}

	status := postgresql.CheckExists(Db, &models.Todo{}, "id = ?", uint(ID))

	if !status {
		return models.Todo{}, "record not found", 404, errors.New("record not found")
	}

	data.ID = uint(ID)

	_, err = postgresql.SaveAllFields(Db, data)

	if err != nil {
		return models.Todo{}, "", 0, err
	}

	return *data, "", 0, nil
}

func DeleteToDOService(id string, Db *gorm.DB) (models.Todo, string, int, error) {

	ID, err := strconv.Atoi(id)
	if err != nil {
		return models.Todo{}, "", 0, err
	}

	status := postgresql.CheckExists(Db, &models.Todo{}, "id = ?", uint(ID))

	if !status {
		return models.Todo{}, "record not found", 404, errors.New("record not found")
	}

	return models.Todo{}, "", 0, nil
}
