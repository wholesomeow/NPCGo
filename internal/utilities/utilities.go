package utilities

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

func ReadCSV(path string, header bool) ([][]string, error) {
	var nil_data [][]string

	// Open CSV File
	log.Printf("reading %s file", path)
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil_data, err
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
		return data, err
	}

	return data, nil
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

func ReadJSON(path string) ([]byte, error) {
	var nil_data []byte

	// Open JSON File
	log.Printf("reading %s file", path)
	f, err := os.Open(path)
	if err != nil {
		return nil_data, err
	}
	defer f.Close()

	// Read JSON File
	byte_value, err := io.ReadAll(f)
	if err != nil {
		return nil_data, err
	}

	return byte_value, nil
}

func parseJSONToSlice(data interface{}) [][]string {
	var result [][]string

	// Check the type of the data and handle accordingly
	switch v := data.(type) {
	case map[string]interface{}:
		// If it's a map, traverse the keys and values recursively
		for _, value := range v {
			// Recursively process each value in the map
			result = append(result, parseJSONToSlice(value)...)
		}

	case []interface{}:
		// If it's a slice, iterate through the elements and process each one
		for _, item := range v {
			// Recursively process each item in the slice
			result = append(result, parseJSONToSlice(item)...)
		}

	case string:
		// If it's a string, add it as a single-element slice
		result = append(result, []string{v})

	case float64:
		// If it's a number, convert it to string and add it as a single-element slice
		result = append(result, []string{fmt.Sprintf("%f", v)})

	case bool:
		// If it's a boolean, convert it to string and add it as a single-element slice
		result = append(result, []string{fmt.Sprintf("%t", v)})

	default:
		// For any other data type, handle it by converting it to string
		result = append(result, []string{fmt.Sprintf("%v", v)})
	}

	return result
}

func JSONToSlice(data []byte) ([][]string, error) {
	// Use a generic map to unmarshal the JSON data
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Parse the data into a nested slice of strings
	return parseJSONToSlice(data), nil
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

// Returns random value within a range. Wrapper for rand.Intn() function.
func RandomRange(min int, max int) int {
	return rand.Intn((max - min + 1)) + min
}

func RemapInt(value float64, minInput float64, maxInput float64, minOutput float64, maxOutput float64) float64 {
	var part_1 float64
	var part_2 float64
	part_1 = (value - minInput) / (maxInput - minInput)
	part_2 = (maxOutput - minOutput) + minOutput
	return part_1 * part_2
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
