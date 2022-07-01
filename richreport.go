package richreport

import (
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/richoffice/richreport/xlsx2map"
)

func LoadDataFrames(excelFile, excelDefFile string, options *xlsx2map.Options) (map[string]dataframe.DataFrame, error) {
	data, err := xlsx2map.LoadFromFile(excelFile, excelDefFile, options)
	if err != nil {
		return nil, err
	}

	dfs := make(map[string]dataframe.DataFrame)

	for key, sheetRows := range data {
		dfs[key] = dataframe.LoadMaps(sheetRows, dataframe.DetectTypes(true))
	}

	return dfs, nil
}

func StoreDataFrames(frames map[string]dataframe.DataFrame, outExcelFile, excelDefFile string, opts *xlsx2map.Options) error {
	data := make(map[string][]map[string]interface{})
	for key, frame := range frames {
		data[key] = frame.Maps()
	}

	return xlsx2map.ExportToFile(data, outExcelFile, excelDefFile, opts)
}

//2022-05-09 08:02:00 +0000 UTC
func ParseDate(dataStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 -0700 UTC", dataStr)
}
