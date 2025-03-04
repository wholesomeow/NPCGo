package grammar

import (
	"fmt"
	"io"
	"os"

	"github.com/qmuntal/opc"
)

// Defines the grammar in a custom .cfg file from structs using OPC

func writeCFG(path string) (string, error) {
	// Create a file to write our archive to.
	f, err := os.Create(path)
	if err != nil {
		return "Error creating file ", err
	}

	// Create a new OPC archive.
	w := opc.NewWriter(f)

	// Create a new OPC part.
	name := opc.NormalizePartName("docs\\readme.txt")
	part, err := w.Create(name, "text/plain")
	if err != nil {
		return "Error creating part ", err
	}

	// Write content to the part.
	part.Write([]byte("<start> -> <expr>"))

	// Make sure to check the error on Close.
	w.Close()

	return "Error creating file ", err
}

func readCFG(path string) (string, error) {
	r, err := opc.OpenReader(path)
	if err != nil {
		return "Error reading file ", err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.Files {
		fmt.Printf("Contents of %s with type %s :\n", f.Name, f.ContentType)
		rc, _ := f.Open()
		io.CopyN(os.Stdout, rc, 68)
		rc.Close()
		fmt.Println()

	}

	return "Error reading file ", err
}
