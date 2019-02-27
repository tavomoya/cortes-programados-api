package lib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/goodsign/monday"
)

var timeMap = map[float64]string{
	1:    "1:00 AM",
	1.5:  "1:30 AM",
	2:    "2:00 AM",
	2.5:  "2:30 AM",
	3:    "3:00 AM",
	3.5:  "3:30 AM",
	4:    "4:00 AM",
	4.5:  "4:30 AM",
	5:    "5:00 AM",
	5.5:  "5:30 AM",
	6:    "6:00 AM",
	6.5:  "6:30 AM",
	7:    "7:00 AM",
	7.5:  "7:30 AM",
	8:    "8:00 AM",
	8.5:  "8:30 AM",
	9:    "9:00 AM",
	9.5:  "9:30 AM",
	10:   "10:00 AM",
	10.5: "10:30 AM",
	11:   "11:00 AM",
	11.5: "11:30 AM",
	12:   "12:00 AM",
	12.5: "12:30 AM",
	13:   "1:00 PM",
	13.5: "1:30 PM",
	14:   "2:00 PM",
	14.5: "2:30 PM",
	15:   "3:00 PM",
	15.5: "3:30 PM",
	16:   "4:00 PM",
	16.5: "4:30 PM",
	17:   "5:00 PM",
	17.5: "5:30 PM",
	18:   "6:00 PM",
	18.5: "6:30 PM",
	19:   "7:00 PM",
	19.5: "7:30 PM",
	20:   "8:00 PM",
	20.5: "8:30 PM",
	21:   "9:00 PM",
	21.5: "9:30 PM",
	22:   "10:00 PM",
	22.5: "10:30 PM",
	23:   "11:00 PM",
	23.5: "11:30 PM",
	24:   "12:00 PM",
	24.5: "12:30 PM",
}

// GetTimeString returns the string associated
// with the hour passed as a float number
func GetTimeString(time float64) string {
	return timeMap[time]
}

func ParseLocalTimeString(dateStr string) time.Time {
	d := strings.Replace(dateStr, "de", "", -1)
	d = strings.Replace(d, ",", "", -1)
	date, err := monday.ParseInLocation("2 January 2006", d, time.UTC, monday.LocaleEsES)
	if err != nil {
		fmt.Println(err)
	}

	return date
}

func ParseInterfaceToStruct(src interface{}, dst interface{}) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Error encoding source: %v", err)
	}
	err = json.Unmarshal(raw, dst)
	if err != nil {
		return fmt.Errorf("Error decoding object: %v", err)
	}
	return nil
}

func ParseMapToStruct(src map[string]interface{}, dst interface{}) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Error encoding source: %v", err)
	}
	err = json.Unmarshal(raw, dst)
	if err != nil {
		return fmt.Errorf("Error decoding object: %v", err)
	}
	return nil
}

func ParseMapToSlice(src []*map[string]interface{}, dst interface{}) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Error encoding source: %v", err)
	}
	err = json.Unmarshal(raw, dst)
	if err != nil {
		return fmt.Errorf("Error decoding object: %v", err)
	}
	return nil
}

func StructToMap(src interface{}, dst *map[string]interface{}) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Error encoding source: %v", err)
	}
	err = json.Unmarshal(raw, dst)
	if err != nil {
		return fmt.Errorf("Error decoding source: %v", err)
	}
	return nil
}
