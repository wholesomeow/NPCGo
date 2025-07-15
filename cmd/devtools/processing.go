package main

import (
	"log"
	"time"

	rawdataproc "github.com/wholesomeow/npcGo/internal/rawdataProcessing"
)

func main() {
	// Process JSONL Files - Adjectives
	start_proc := time.Now()
	log.Print("starting JSONL processing - 'pos-adj.jsonl'")

	// Get first example of file
	// err = rawdataproc.ExtractFirstJSONL(
	// 	"../../../data/rawdata/jsonl/pos-adj.jsonl",
	// 	"../../../data/rawdata/json/pos-adj.json",
	// )
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	err := rawdataproc.ProcessJSONL("data/raw/jsonl/pos-adj.jsonl")
	// err := rawdataproc.ProcessJSONL("../../../data/rawdata/jsonl/pos-verb.jsonl")
	if err != nil {
		log.Fatalf("failure in JSONL data processing: %v", err)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("processing completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	// Process JSONL Files - Verbs
	start_proc = time.Now()
	log.Print("starting JSONL processing - 'pos-adj.jsonl'")

	// Get first example of file
	// err = rawdataproc.ExtractFirstJSONL(
	// 	"../../../data/rawdata/jsonl/pos-verb.jsonl",
	// 	"../../../data/rawdata/json/pos-verb.json",
	// )
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	err = rawdataproc.ProcessJSONL("data/raw/jsonl/pos-verb.jsonl")
	// err = rawdataproc.ProcessJSONL("../../../data/rawdata/jsonl/pos-verb.jsonl")
	if err != nil {
		log.Fatalf("failure in JSONL data processing: %v", err)
	}
	end_proc = time.Now()
	elapsed_proc = end_proc.Sub(start_proc)
	log.Printf("processing completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	// Process JSONL Files - Nouns
	start_proc = time.Now()
	log.Print("starting JSONL processing - 'pos-noun.jsonl'")

	// Get first example of file
	// err = rawdataproc.ExtractFirstJSONL(
	// 	"../../../data/rawdata/jsonl/pos-noun.jsonl",
	// 	"../../../data/rawdata/json/pos-noun.json",
	// )
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	err = rawdataproc.ProcessJSONL("data/rawd/jsonl/pos-noun.jsonl")
	// err = rawdataproc.ProcessJSONL("../../../data/rawdata/jsonl/pos-noun.jsonl")
	if err != nil {
		log.Fatalf("failure in JSONL data processing: %v", err)
	}
	end_proc = time.Now()
	elapsed_proc = end_proc.Sub(start_proc)
	log.Printf("processing completed... elapsed time: %s", time.Duration.String(elapsed_proc))
}
