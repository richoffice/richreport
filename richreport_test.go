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
