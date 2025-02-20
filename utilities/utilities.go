package utilities

import (
	"encoding/csv"
	"fmt"
	"go/npcGen/configuration"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadCSV(path string, header bool) [][]string {
	// Open CSV File
	log.Printf("reading %s file", path)
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("error opening CSV: %s", err)
	}
	defer f.Close()

	// Read CSV File
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()

	// If csv has a header row, remove that row from the parsed data
	if header {
		data = data[1:]
	}
	if err != nil {
		log.Printf("error reading CSV: %s", err)
	}

	return data
}

func WriteCSV(path string, filename string, data [][]string) {
	// Create CSV file
	location := fmt.Sprintf("%s/%s", path, filename)
	log.Printf("writing data to %s", location)
	file, err := os.Create(location)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	// Write to CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, val := range data {
		var row []string
		row = append(row, val...)
		writer.Write(row)
	}
}

func ReadJSON(path string) []byte {
	// Open JSON File
	log.Printf("reading %s file", path)
	f, err := os.Open(path)
	if err != nil {
		log.Printf("error opening JSON: %s", err)
	}
	defer f.Close()

	// Read JSON File
	byte_value, _ := io.ReadAll(f)

	return byte_value
}

func ReadConfig(path string, config *configuration.Config) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse file based on file extention
	// TODO(wholesomeow): Implement Environment variables and JSON
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		yaml_decoder := yaml.NewDecoder(f)
		err := yaml_decoder.Decode(config)
		if err != nil {
			// TODO(wholesomeow): Figure out better logging/error handling for known things like this
			log.Fatal(err)
		}
	default:
		return fmt.Errorf("file format '%s' not supported by parser", ext)
	}

	return nil
}

func CheckFilePath(path string, required bool) bool {
	// TODO(wholesomeow): Maybe rewrite to remove error handling and make function more flexible
	found := true
	_, err := os.Stat(path)
	if err == nil {
		log.Printf("file %s exists", path)
		return found
	} else if os.IsNotExist(err) {
		if !required {
			log.Printf("file %s doesn't exist in expected location", path)
		}
		log.Fatalf("file %s does not exist", path)
	} else {
		log.Fatalf("file %s stat error: %v", path, err)
	}
	return !found
}

func SliceContainsString(str string, slc []string) bool {
	for _, char := range slc {
		if char == str {
			return true
		}
	}

	return false
}

// TODO(wholesomeow): Implement RandomRange function that uses generics and optional parameters to return random value in a range
//                    This function -> r_val := rand.Intn(len(npc.Pronouns)) + 1

func ImperialToMetric(inches int, lbs int) (float64, float64) {
	cm := float64(inches) * 2.54
	kg := float64(lbs) * 0.453592

	return cm, kg
}

func RoundToDecimal(number float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(number*multiplier) / multiplier
}

func ChangeWorkingDir(path_diff string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Change to project root before executing test
	root := filepath.Join(cwd, path_diff)
	err = os.Chdir(root)
	if err != nil {
		return err
	}

	return nil
}
