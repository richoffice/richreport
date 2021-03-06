package xlsx2map

import (
	"fmt"
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {

	def := &XlsxFileDef{}
	file, err := os.Open("sample_def.json")

	if err != nil {
		t.Errorf("expected no err, but got %v", err)
	}

	defer file.Close()

	loadErr := LoadXlsxFileDef(file, def)
	if loadErr != nil {
		t.Errorf("expected no err, but got %v", loadErr)
	}

	xlsxFile := "sample_file.xlsx"

	xlsxMaps := make(map[string][]map[string]interface{})

	err = Unmarshal(xlsxFile, xlsxMaps, def, nil)
	if err != nil {
		t.Errorf("Unmarshal() error = %v, wantErr %v", err, nil)
	}

	if xlsxMaps["visitors"][1]["name"] != "Tom Hanks" {
		t.Errorf("Expected 'Tom Hanks', but got %v", xlsxMaps["visitors"][1]["name"])

	}

	fmt.Println(xlsxMaps)
	// if !reflect.DeepEqual(got, tt.want) {
	// 	t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
	// }

	outErr := Marshal("test_out.xlsx", xlsxMaps, def)
	if outErr != nil {
		t.Errorf("Expected no output err, but got %v", outErr)
	}

}

func TestSheet1Test(t *testing.T) {

	def := &XlsxFileDef{}
	file, err := os.Open("./testfiles/wh-def.json")

	if err != nil {
		t.Errorf("expected no err, but got %v", err)
	}

	defer file.Close()

	loadErr := LoadXlsxFileDef(file, def)
	if loadErr != nil {
		t.Errorf("expected no err, but got %v", loadErr)
	}

	xlsxFile := "testfiles/input.xlsx"

	xlsxMaps := make(map[string][]map[string]interface{})

	err = Unmarshal(xlsxFile, xlsxMaps, def, nil)
	if err != nil {
		t.Errorf("Unmarshal() error = %v, wantErr %v", err, nil)
	}

	fmt.Println(xlsxMaps)
	// if !reflect.DeepEqual(got, tt.want) {
	// 	t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
	// }

	outErr := Marshal("testfiles/output.xlsx", xlsxMaps, def)
	if outErr != nil {
		t.Errorf("Expected no output err, but got %v", outErr)
	}

}

func TestLoadFromFile(t *testing.T) {
	data, err := LoadFromFile("sample_file.xlsx", "sample_def.json", nil)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	// fmt.Println(data)

	err = ExportToFile(data, "test_out.xlsx", "sample_def.json", nil)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
