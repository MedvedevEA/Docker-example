package service

import (
	"github.com/google/uuid"
)

type Record struct {
	RecordId *uuid.UUID `json:"record_id"`
	Text     string     `json:"text"`
}

var record = []*Record{}

func GetRecordCount() int {
	return len(record)
}
func GetRecords() []*Record {
	return record
}
func AddRecord(text string) *Record {
	recordId := uuid.New()
	newRecord := &Record{
		RecordId: &recordId,
		Text:     text,
	}
	record = append(record, newRecord)
	return newRecord
}
func DeleteRecords() int {
	record = []*Record{}
	return len(record)
}
