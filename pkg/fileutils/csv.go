package fileutils

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func CSVNewFile(filePathName string) (*csv.Writer, error) {
	if !strings.HasSuffix(filePathName, ".csv") {
		filePathName += ".csv"
	}
	csvFile, err := os.Create(filePathName)
	if err != nil {
		log.Fatalf("fileutils.WriteNewCSVFile - Failed creating file.")
		return nil, err
	}
	csvwriter := csv.NewWriter(csvFile)

	return csvwriter, nil
}

func CSVAppendLine(csvFile *csv.Writer, line []string) error {
	err := csvFile.Write(line)
	if err != nil {
		log.Fatalf("fileutils.AppendLine - Failed writing to file.")
		return err
	}
	csvFile.Flush()

	return nil
}

func CSVAppendLines(csvFile *csv.Writer, lines [][]string) error {
	err := csvFile.WriteAll(lines)
	if err != nil {
		log.Fatalf("fileutils.AppendLines - Failed writing to file.")
		return err
	}
	csvFile.Flush()

	return nil
}
