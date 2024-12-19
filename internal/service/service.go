package service

import (
	"github.com/google/uuid"
)

type Record struct {
	RecordId *uuid.UUID `json:"record_id"`
	Text     string     `json:"text"`
}

var record = []*Record{}

func GetRecordCount() (int, error) {
	return len(record), nil
}
func GetRecords() ([]*Record, error) {
	return record, nil
}
func AddRecord(text string) (*Record, error) {
	recordId := uuid.New()
	newRecord := &Record{
		RecordId: &recordId,
		Text:     text,
	}
	record = append(record, newRecord)
	return newRecord, nil
}
func RemoveRecords() error {
	record = []*Record{}
	return nil
}
