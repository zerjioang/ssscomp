package helper

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadCsv(path string, fields int, skipHeader bool, callback func(obs float64, vars []float64)) error {
	// we open the csv file from the disk
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// we create a new csv reader specifying
	// the number of columns it has
	csvData := csv.NewReader(f)
	csvData.FieldsPerRecord = fields
	// we read all the records
	records, err := csvData.ReadAll()
	if err != nil {
		return err
	}
	if skipHeader {
		// by slicing the records we skip the header
		records = records[1:]
	}

	// Loop of records in the CSV, adding the training data to the regressionvalue.
	for _, record := range records {
		speed, _ := strconv.Atoi(record[1])
		dist, _ := strconv.Atoi(record[2])
		// Add these points to the regression value.
		callback(float64(dist), []float64{float64(speed)})
	}
	return nil
}
