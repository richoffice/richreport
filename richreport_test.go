package richreport

import (
	"fmt"
	"testing"
)

func TestLoadDataFrames(t *testing.T) {

	got, err := LoadDataFrames("test/sample_file.xlsx", "test/sample_def.json", nil)
	if err != nil {
		t.Errorf("LoadDataFrames() error = %v, want no error", err)
		return
	}
	fmt.Println(got)

}

func TestLoadDateDataFrames(t *testing.T) {

	got, err := LoadDataFrames("xlsx2map/testfiles/input.xlsx", "xlsx2map/testfiles/wh-def.json", nil)
	if err != nil {
		t.Errorf("LoadDataFrames() error = %v, want no error", err)
		return
	}
	frame := got["workhours"]
	fmt.Println(frame)
	fmt.Println(frame.Col("dueDispatchTime").Str())

	fmt.Println(frame.Col("dueDispatchTime").Records())

	create, err := ParseDate(frame.Col("dueDispatchTime").Records()[0])
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	fmt.Println(create)

	// fmt.Println(got)

}

func TestStoreDataFrames(t *testing.T) {
	got, err := LoadDataFrames("test/sample_file.xlsx", "test/sample_def.json", nil)
	if err != nil {
		t.Errorf("LoadDataFrames() error = %v, want no error", err)
		return
	}

	err = StoreDataFrames(got, "test/sample_out.xlsx", "test/sample_def.json", nil)
	if err != nil {
		t.Errorf("StoreDataFrames() error = %v, want no error", err)
		return
	}
}
