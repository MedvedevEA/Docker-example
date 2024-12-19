package service

import "testing"

func TestGetRecordCount(t *testing.T) {
	_, err := GetRecordCount()
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
func TestGetRecords(t *testing.T) {
	_, err := GetRecords()
	if err != nil {
		t.Errorf("error %s", err)
	}
}
func TestAddRecord(t *testing.T) {
	recordText := "any text"
	record, err := AddRecord(recordText)
	if err != nil {
		t.Errorf("error %s", err)
	}
	if record.Text != recordText {
		t.Errorf("record text expected to be %s; got %s", recordText, record.Text)
	}
}
func TestRemoveRecords(t *testing.T) {
	err := RemoveRecords()
	if err != nil {
		t.Errorf("error %s", err)
	}
}
