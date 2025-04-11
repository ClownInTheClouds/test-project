package mapper

import (
	"fmt"
	"strconv"
	"time"

	"test-project/internal/models"
)

type TaskCsvMapper struct {
	CsvMapper[models.Task]
}

func (taskCsvMapper TaskCsvMapper) MapRow(row []string) (*models.Task, error) {
	return taskCsvMapper.fromCsv(row)
}

func (taskCsvMapper TaskCsvMapper) fromCsv(row []string) (*models.Task, error) {
	if len(row) != 4 {
		return nil, fmt.Errorf("Wrong columns number: %v\n", len(row))
	}

	var task models.Task
	var err error

	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return nil, err
	}
	task.SetId(id)
	task.Description = row[1]
	task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", row[2])
	if err != nil {
		return nil, err
	}
	task.IsCompleted, err = strconv.ParseBool(row[3])
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (taskCsvMapper TaskCsvMapper) ToCsv(task *models.Task) []string {
	return []string{
		strconv.FormatUint(task.GetId(), 10),
		task.Description,
		task.CreatedAt.Format("2006-01-02 15:04:05"),
		strconv.FormatBool(task.IsCompleted),
	}
}
