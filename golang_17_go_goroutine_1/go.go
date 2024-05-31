package main

import (
 "encoding/json"
 "fmt"
 "github.com/360EntSecGroup-Skylar/excelize"
 "os"
)

type Data struct {
 ID         int    
 FirstName  string json:"first_name"
 LastName   string json:"last_name"
 Email      string json:"email"
 Age        string json:"age"
}

func main() {
 jsonFile, err := os.Open("data.json")
 if err != nil {
  fmt.Println("Error opening JSON file:", err)
  return
 }
 defer jsonFile.Close()

 var data []Data
 jsonDecoder := json.NewDecoder(jsonFile)
 err = jsonDecoder.Decode(&data)
 if err != nil {
  fmt.Println("Error decoding JSON data:", err)
  return
 }

 xlsx := excelize.NewFile()

 sheetName := "Sheet1"

 headers := []string{"id", "first_name", "last_name", "email", "age"}
 for col, header := range headers {
  cell := excelize.ToAlphaString(col) + "1"
  xlsx.SetCellValue(sheetName, cell, header)
 }

 for row, item := range data {
  rowIndex := row + 2 
  xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ID)
  xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.FirstName)
  xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.LastName)
  xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Email)
  xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex), item.Age)
 }

 err = xlsx.SaveAs("data.xlsx")
 if err != nil {
  fmt.Println("Error saving Excel file:", err)
  return
 }

 fmt.Println("Data from JSON file has been written to data.xlsx")
}
