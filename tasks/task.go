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
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Data struct {
	Data []Task `json:"task"`
}

func openJSON() (*os.File, error) {
	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func AddJSON(arg1 string, arg2 string) error {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
	} else {
		data = Data{Data: []Task{}}
	}

	lastid := 0
	if len(data.Data) > 0 {
		lastid = data.Data[len(data.Data)-1].Id
	}

	if arg1 == "add" {
		updatedTime := time.Now()

		newTask := Task{
			Id:          lastid + 1,
			Description: arg2,
			Status:      "To-Do",
			CreatedAt:   time.Now(),
			UpdatedAt:   updatedTime,
		}
		data.Data = append(data.Data, newTask)

		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return err
		}

		err = os.WriteFile("tasks.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
		return nil
	}
	return fmt.Errorf("incorrect argument")
}

func UpdateJSON(arg1 string, arg2 int, arg3 string) error {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
	} else {
		return fmt.Errorf("no data found")
	}

	if arg1 == "update" {
		updatedTime := time.Now()
		found := false
		for i := range data.Data {
			if data.Data[i].Id == arg2 {
				data.Data[i].Description = arg3
				data.Data[i].UpdatedAt = updatedTime
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("task with id %d not found", arg2)
		}

		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return err
		}
		err = os.WriteFile("tasks.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
		return nil
	}
	return fmt.Errorf("incorrect argument")
}

func DeleteJSON(arg1 string, arg2 int) error {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
	} else {
		return fmt.Errorf("no data found")
	}

	if arg1 == "delete" {
		indextodelete := -1
		for i := range data.Data {
			if data.Data[i].Id == arg2 {
				indextodelete = i
				break
			}
		}
		if indextodelete == -1 {
			return fmt.Errorf("task with id %d not found", arg2)
		}

		data.Data = append(data.Data[:indextodelete], data.Data[indextodelete+1:]...)
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return err
		}
		err = os.WriteFile("tasks.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
		return nil
	}
	return fmt.Errorf("incorrect argument")
}

func Change_status(arg1 string, arg2 int) error {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
	} else {
		return fmt.Errorf("no data found")
	}
	updatedTime := time.Now()
	found := false
	for i := range data.Data {
		if data.Data[i].Id == arg2 {
			data.Data[i].Status = arg1
			data.Data[i].UpdatedAt = updatedTime
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with id %d not found", arg2)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func ListAll() (Data, error) {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return Data{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
		return data, nil
	} else {
		return Data{}, fmt.Errorf("no data found")
	}
}

func Liststatuswise(arg1 string) (Data, error) {
	jsonFile, err := openJSON()
	if err != nil {
		fmt.Println(err)
		return Data{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	if len(byteValue) > 0 {
		json.Unmarshal(byteValue, &data)
	} else {
		return Data{}, fmt.Errorf("no data found")
	}

	var datatoreturn Data

	for i := range data.Data {
		if data.Data[i].Status == arg1 {
			datatoreturn.Data = append(datatoreturn.Data, data.Data[i])
		}
	}

	return datatoreturn, nil

}
