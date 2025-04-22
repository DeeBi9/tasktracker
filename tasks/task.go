package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"string"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type Data struct {
	Data []Task `json:"task"`
}

func openJSON() os.File {
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
	}

	return *jsonFile
}

func AddJSON(arg1 string, arg2 string) error {
	jsonFile := openJSON()

	byteValue, _ := ioutil.ReadAll(&jsonFile)

	var data Data
	json.Unmarshal(byteValue, &data)
	lastid := data.Data[len(data.Data)-1].Id

	defer jsonFile.Close()

	if arg1 == "add" {
		updatedTime := time.Now()
		if len(data.Data) == 0 {
			lastid = 1
		}

		d := Data{
			Data: []Task{
				{
					Id:          lastid + 1,
					Description: arg2,
					Status:      "To-Do",
					CreatedAt:   time.Now(),
					UpdatedAt:   updatedTime,
				},
			},
		}
		jsonData, err := json.Marshal(d)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return err
		}
		_, err = jsonFile.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
		return nil
	} else {
		return fmt.Errorf("incorrect argument")
	}

}

func UpdateJSON(arg1 string, arg2 int, arg3 string) error {
	jsonFile := openJSON()

	byteValue, _ := ioutil.ReadAll(&jsonFile)

	var data Data
	json.Unmarshal(byteValue, &data)

	defer jsonFile.Close()

	if arg1 == "update" {
		updatedTime := time.Now()
		d := Data{
			Data: []Task{
				{
					Description: arg3,
					UpdatedAt:   updatedTime,
				},
			},
		}
		for i := range data.Data {
			if data.Data[i].Id == arg2 {
				jsonData, err := json.Marshal(d)
				if err != nil {
					fmt.Println("Error encoding JSON:", err)
					return err
				}
				_, err = jsonFile.Write(jsonData)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return err
				}
			}
		}
		return nil
	} else {
		return fmt.Errorf("incorrect argument")
	}
}

func DeleteJSON(arg1 string, arg2 int) error {
	jsonFile := openJSON()

	byteValue, _ := ioutil.ReadAll(&jsonFile)

	var data Data
	json.Unmarshal(byteValue, &data)

	defer jsonFile.Close()

	if arg1 == "delete" {
		indextodelete := -1
		for i := range data.Data {
			if data.Data[i].Id == arg2 {
				indextodelete = i
				break
			}
		}
		if indextodelete != -1 {
			data.Data = append(data.Data[:indextodelete], data.Data[indextodelete+1:]...)
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error encoding JSON:", err)
				return err
			}
			_, err = jsonFile.Write(jsonData)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return err
			}
		}
		return nil
	} else {
		return fmt.Errorf("incorrect argument")
	}
}
