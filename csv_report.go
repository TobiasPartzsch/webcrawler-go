package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

const warningMsgNoData = "No data to write to CSV"
const errMsgCsvCreate = "create csv: %w"
const errMsgWriteHeader = "write header: %w"
const errMsgWriteRow = "write row for %s: %w"
const msgWriteSuccess = "Report written to %s\n"

var headerLine = []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}

func writeCSVRport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println(warningMsgNoData)
		return nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf(errMsgCsvCreate, err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	if err := w.Write(headerLine); err != nil {
		return fmt.Errorf(errMsgWriteHeader, err)
	}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, normalizedURL := range keys {
		pd := pages[normalizedURL]
		if err := w.Write([]string{
			pd.URL,
			pd.H1,
			pd.FirstParagraph,
			strings.Join(pd.OutgoingLinks, ";"),
			strings.Join(pd.ImageURLs, ";"),
		}); err != nil {
			return fmt.Errorf(errMsgWriteRow, pd.URL, err)
		}
	}

	fmt.Printf(msgWriteSuccess, filename)
	return nil
}
