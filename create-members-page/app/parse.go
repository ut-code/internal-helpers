package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func ParseFile(ac AppContext, path string) ([]Member, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file at %s: %w", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("error closing file:", err)
		}
	}()

	var sheetName = "フォームの回答 1"
	var members []Member
	// it's 1-indexed in excelize and there's also a header row
	headers, err := getRow(f, sheetName, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to get headers: %w", err)
	}
	for ri := 2; ; ri++ {
		rowData, err := getRow(f, sheetName, ri)
		if err != nil {
			return nil, fmt.Errorf("failed to get row %d: %w", ri, err)
		}
		if len(rowData) == 0 {
			break
		}
		row := parseRow(ac, rowData, headers)

		members = append(members, row)
	}

	return members, nil
}

func getRow(f *excelize.File, sheet string, row int) ([]string, error) {
	var cells []string
	var col int = 1

	for cell := "-"; cell != ""; col++ {
		// convert column number to letter: 1 -> "A", 2 -> "B", etc.
		colString, err := excelize.ColumnNumberToName(col)
		if err != nil {
			return nil, err
		}
		cell, err := f.GetCellValue(sheet, colString+strconv.Itoa(row))
		if err != nil {
			return nil, err
		}
		if cell == "" {
			break
		}
		cells = append(cells, cell)
	}

	return cells, nil
}

func parseRow(ac AppContext, row []string, headers []string) Member {
	member := Member{
		Metadata: Metadata{},
	}

	member.Metadata.JoinYear = 2025

	for i, cell := range row {
		switch headers[i] {
		case "列 1", "メールアドレス":
			// ignore
		case "タイムスタンプ":
			ts, err := time.Parse("1/2/2006 15:04:05", cell)
			if err != nil {
				log.Fatalln("failed to parse timestamp:", err)
			}
			member.Metadata.Timestamp = ts
		case "名前":
			member.Metadata.NameJa = cell
		case "名前 (ローマ字)":
			member.Metadata.NameEn = cell
		case "簡単な一言":
			member.Metadata.Description = cell
		case "写真":
			member.Metadata.PicturePath = ac.PicturesDirectory + cell
		case "自分のGithubのURL  (あれば)":
			member.Metadata.GitHub = cell
		case "自分のTwitter(X)のID (載せたければ)":
			member.Metadata.Twitter = cell
		case "自分のWebsiteのURL  (あれば)":
			member.Metadata.Website = cell
		case "自己紹介文":
			member.Body = cell
		default:
			log.Fatalln("Unknown header:", headers[i])
		}
	}

	if err := MemberPreprocess(&member); err != nil {
		log.Fatalln("[parser] failed to preprocess member:", err)
	}

	return member
}

func NameToFileName(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
