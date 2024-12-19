package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"rafli", "bima", "pratandra"})
	_ = writer.Write([]string{"raka", "nugraha", "pratandra"})
	_ = writer.Write([]string{"adi", "riyadi", "pratandra"})

	writer.Flush()
}
