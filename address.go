package main

import (
	"address/database"
	"address/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("LocList.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Location{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}


	for _, item := range v.CountryRegion {
		cAddress := models.Address{
			PID:         0,
			Name:        item.Name,
			Code:        item.Code,
			Country:     item.Name,
			CountryCode: item.Code,
			State:       "",
			StateCode:   "",
			City:        "",
			CityCode:    "",
			Region:      "",
			RegionCode:  "",
		}
		database.MysqlDB.Create(&cAddress)
		for _, state := range item.State {
			sAddress := models.Address{
				PID:         cAddress.ID,
				Name:        state.Name,
				Code:        cAddress.Code+"."+state.Code,
				Country:     item.Name,
				CountryCode: item.Code,
				State:       state.Name,
				StateCode:   state.Code,
				City:        "",
				CityCode:    "",
				Region:      "",
				RegionCode:  "",
			}
			if len(state.Code) == 0 {
				sAddress.ID = cAddress.ID
			}else {
				database.MysqlDB.Create(&sAddress)
			}
			for _, city := range state.City {
				cityAddress := models.Address{
					PID:         sAddress.ID,
					Name:        city.Name,
					Code:        sAddress.Code+"."+city.Code,
					Country:     item.Name,
					CountryCode: item.Code,
					State:       state.Name,
					StateCode:   state.Code,
					City:        city.Name,
					CityCode:    city.Code,
					Region:      "",
					RegionCode:  "",
				}
				database.MysqlDB.Create(&cityAddress)
				for _, region := range city.Region {
					rAddress := models.Address{
						PID:         cityAddress.ID,
						Name:        region.Name,
						Code:        cityAddress.Code+"."+region.Code,
						Country:     item.Name,
						CountryCode: item.Code,
						State:       state.Name,
						StateCode:   state.Code,
						City:        city.Name,
						CityCode:    city.Code,
						Region:      region.Name,
						RegionCode:  region.Code,
					}
					database.MysqlDB.Create(&rAddress)
					fmt.Print(".")
				}
			}
		}
	}
}
