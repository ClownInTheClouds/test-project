package repository

import (
	"errors"
	"log"
	"strconv"

	"test-project/internal/mapper"
	"test-project/internal/models"
	"test-project/internal/service"
)

type TaskRepository interface {
	CrudRepository[models.Task, uint64]
}

type CsvTaskRepository struct {
	TaskRepository
	Path   []string
	Mapper mapper.TaskCsvMapper
	id     uint64
}

func (tr *CsvTaskRepository) Create(task *models.Task) (bool, error) {
	file, err := service.GetFile(tr.Path...)
	if err != nil {
		log.Printf("Can not get repository csv-file: %v\n", err)
		return false, err
	}
	defer service.CloseFile(file)

	reader := service.CreateCsvReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Can not read repository csv-file: %v\n", err)
		return false, err
	}

	err = tr.defineId(records)
	if err != nil {
		log.Printf("Can not define repository id: %v\n", err)
		return false, err
	}

	writer := service.CreateCsvWriter(file)
	defer writer.Flush()

	if len(records) == 0 {
		if err := writer.Write(getCsvHeader()); err != nil {
			log.Printf("Can not write header to csv-file: %v\n", err)
			return false, err
		}
	}

	task.SetId(tr.nextId())
	taskAsCsv := tr.Mapper.ToCsv(task)
	if err := writer.Write(taskAsCsv); err != nil {
		log.Printf("Can not write task to csv-file: %v\n", err)
		return false, err
	}

	return true, nil
}

func (tr *CsvTaskRepository) Read(id uint64) (*models.Task, error) {
	file, err := service.GetFile(tr.Path...)
	if err != nil {
		log.Printf("Can not get repository csv-file: %v\n", err)
		return nil, err
	}
	defer service.CloseFile(file)

	reader := service.CreateCsvReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Can not read repository csv-file: %v\n", err)
		return nil, err
	}

	idStr := strconv.FormatUint(id, 10)
	for _, record := range records {
		if idStr == record[0] {
			task, err := tr.Mapper.MapRow(record)
			if err != nil {
				log.Printf("Can not map row: %v\n", err)
				return nil, err
			}
			return task, nil
		}
	}

	return nil, errors.New("can not find task")
}

func (tr *CsvTaskRepository) ReadAll() ([]*models.Task, error) {
	file, err := service.GetFile(tr.Path...)
	if err != nil {
		log.Printf("Can not get repository csv-file: %v\n", err)
		return nil, err
	}
	defer service.CloseFile(file)

	reader := service.CreateCsvReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Can not read repository csv-file: %v\n", err)
		return nil, err
	}

	taskList := make([]*models.Task, len(records)-1)
	for i, record := range records {
		if i == 0 {
			continue
		}
		task, err := tr.Mapper.MapRow(record)
		if err != nil {
			log.Printf("Can not map row: %v\n", err)
			return nil, err
		}
		taskList[i-1] = task
	}

	return taskList, nil
}

func (tr *CsvTaskRepository) defineId(records [][]string) error {
	if len(records) > 1 {
		record := records[len(records)-1]
		latestId, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Printf("Can not parse id: %v\n", err)
			return err
		}
		tr.id = latestId
	} else {
		tr.id = 0
	}
	return nil
}

func getCsvHeader() []string {
	return []string{"ID", "Description", "CreatedAt", "IsCompleted"}
}

func (tr *CsvTaskRepository) nextId() uint64 {
	tr.id++
	return tr.id
}
