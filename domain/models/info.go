package models

import "time"




type Record struct {
	name              string  
	code              string  
	date              time.Time  
	prSystem          float64  
	previous          float64 
	withdrawal        float64 
	slip              float64 
	remainingOnSystem float64 
	uncollected       float64 
}

type RecordConfig struct {
	Name              string  
	Code              string  
	Date              time.Time  
	PRSystem          float64  
	Previous          float64 
	Withdrawal        float64 
	Slip              float64 
	RemainingOnSystem float64 
	Uncollected       float64 
}

func NewRecord(record *RecordConfig) *Record {
	return &Record{
		name:              record.Name,
		code:              record.Code,
		date:              record.Date,
		prSystem:          record.PRSystem,
		previous:          record.Previous,
		withdrawal:        record.Withdrawal,
		slip:              record.Slip,
		remainingOnSystem: record.RemainingOnSystem,
		uncollected:       record.Uncollected,
	}
}

type SystemLog struct {
	records   []*Record
	updatedAt time.Time
	date      time.Time
}

func (s *SystemLog) GetRecords() []*Record {
	return s.records
}

func (s *SystemLog) SetRecords(records []*Record) {
	s.records = records
}

func (s *SystemLog) GetUpdatedAt() time.Time {
	return s.updatedAt
}

func (s *SystemLog) SetUpdatedAt(updatedAt time.Time) {
	s.updatedAt = updatedAt
}

func (s *SystemLog) GetDate() time.Time {
	return s.date
}

func (s *SystemLog) SetDate(date time.Time) {
	s.date = date
}

type SystemLogConfig struct {
	Records   []*RecordConfig
	UpdatedAt time.Time
	Date      time.Time
}

func NewSystemLog(config *SystemLogConfig) *SystemLog {
	records := make([]*Record, len(config.Records))
	for i, recordConfig := range config.Records {
		records[i] = NewRecord(recordConfig)
	}
	return &SystemLog{
		records:   records,
		updatedAt: config.UpdatedAt,
		date:      config.Date,
	}
}


func (r *Record) GetName() string {
	return r.name
}

func (r *Record) SetName(name string) {
	r.name = name
}

func (r *Record) GetCode() string {
	return r.code
}

func (r *Record) SetCode(code string) {
	r.code = code
}

func (r *Record) GetDate() time.Time {
	return r.date
}

func (r *Record) SetDate(date time.Time) {
	r.date = date
}

func (r *Record) GetPRSystem() float64 {
	return r.prSystem
}

func (r *Record) SetPRSystem(prSystem float64) {
	r.prSystem = prSystem
}

func (r *Record) GetPrevious() float64 {
	return r.previous
}

func (r *Record) SetPrevious(previous float64) {
	r.previous = previous
}

func (r *Record) GetWithdrawal() float64 {
	return r.withdrawal
}

func (r *Record) SetWithdrawal(withdrawal float64) {
	r.withdrawal = withdrawal
}

func (r *Record) GetSlip() float64 {
	return r.slip
}

func (r *Record) SetSlip(slip float64) {
	r.slip = slip
}

func (r *Record) GetRemainingOnSystem() float64 {
	return r.remainingOnSystem
}

func (r *Record) SetRemainingOnSystem(remainingOnSystem float64) {
	r.remainingOnSystem = remainingOnSystem
}

func (r *Record) GetUncollected() float64 {
	return r.uncollected
}

func (r *Record) SetUncollected(uncollected float64) {
	r.uncollected = uncollected
}

