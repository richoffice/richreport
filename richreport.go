package richreport

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/rocksun/xlsx2map"
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
