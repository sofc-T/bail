package logsrepo

import (
	"bail/domain/models"
	"time"

	"github.com/google/uuid"
)

type recordDto struct {
	id 			  uuid.UUID    `bson:"_id"`
	name              string    `bson:"name"`
	code              string    `bson:"code"`
	date              time.Time `bson:"date"`
	prSystem          float64   `bson:"pr_system"`
	previous          float64   `bson:"previous"`
	withdrawal        float64   `bson:"withdrawal"`
	slip              float64   `bson:"slip"`
	remainingOnSystem float64   `bson:"remaining_on_system"`
	uncollected       float64   `bson:"uncollected"`
}

func mapRecords(records []*models.Record) []*models.RecordConfig {
	mappedRecords := make([]*models.RecordConfig, len(records))
	for i, record := range records {
		mappedRecords[i] = &models.RecordConfig{
			Name:              record.GetName(),
			Code:              record.GetCode(),
			Date:              record.GetDate(),
			PRSystem:          record.GetPRSystem(),
			Previous:          record.GetPrevious(),
			Withdrawal:        record.GetWithdrawal(),
			Slip:              record.GetSlip(),
			Uncollected:       record.GetUncollected(),
		}
	}
	return mappedRecords
}

type sytemLogDto struct {
	id        uuid.UUID          `bson:"_id"`
	records   []*models.Record `bson:"records"`
	updatedAt time.Time        `bson:"updated_at"`
	date      time.Time        `bson:"date"`
}

      

func fromSystemlogDto(dto sytemLogDto) *models.SystemLog {
	return models.MapSystemLog(
		models.SystemLogConfig{
			Id:        dto.id,
			Records:   dto.records,
			UpdatedAt: dto.updatedAt,
			Date:      dto.date,
		},
	)
}

func fromSystemLog(systemLog *models.SystemLog) sytemLogDto {
	return sytemLogDto{
		id:        systemLog.GetID(),
		records:   systemLog.GetRecords(),
		updatedAt: systemLog.GetUpdatedAt(),
		date:      systemLog.GetDate(),
	}
}




func fromRecordDto(dto recordDto) *models.Record {
	return models.MapRecord(
		models.RecordConfig{
			Name:              dto.name,
			Code:              dto.code,
			Date:              dto.date,
			PRSystem:          dto.prSystem,
			Previous:          dto.previous,
			Withdrawal:        dto.withdrawal,
			Slip:              dto.slip,
			RemainingOnSystem: dto.remainingOnSystem,
			Uncollected:       dto.uncollected,
		},
	)
}

func fromRecord(record *models.Record) recordDto{
	return recordDto{
		name:              record.GetName(),
		code:              record.GetCode(),
		date:              record.GetDate(),
		prSystem:          record.GetPRSystem(),
		previous:          record.GetPrevious(),
		withdrawal:        record.GetWithdrawal(),
		slip:              record.GetSlip(),
		remainingOnSystem: record.GetRemainingOnSystem(),
		uncollected:       record.GetUncollected(),
	}
}