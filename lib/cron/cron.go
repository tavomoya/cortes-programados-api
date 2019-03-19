package cron

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/edenorte"
	"cortes-programados-api/scrapers/edesur"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron"
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	Cron *cron.Cron
	db   *lib.DBLib
}

func NewJob(db *models.DatabaseConfig) *Job {
	return &Job{
		Cron: cron.New(),
		db:   lib.NewDBLib(db),
	}
}

func (j *Job) UpdateOutagesCollection(schedule string) {
	fmt.Println(":)", schedule)
	j.Cron.AddFunc(schedule, func() {
		started := time.Now()
		fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' started ***")
		fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' start time: %v ***\n", started)

		outages, err := getOutagesScrapeData()
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' time elapsed: %v ***\n", ended.Sub(started))
		}

		err = saveOutageCollection(j.db, outages)
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' time elapsed: %v ***\n", ended.Sub(started))
		}

		ended := time.Now()
		fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' finished succesfully ***")
		fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' end time: %v ***\n", ended)
		fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' time elapsed: %v ***\n", ended.Sub(started))
	})

	j.Cron.Start()
}

func getOutagesScrapeData() ([]*models.Outage, error) {

	norte, err := edenorte.ReadOutageAnouncement()
	if err != nil {
		return nil, err
	}

	sur, err := edesur.GetOutageAnouncement()
	if err != nil {
		return nil, err
	}

	return append(norte, sur...), nil
}

func saveOutageCollection(db *lib.DBLib, outages []*models.Outage) error {

	if len(outages) == 0 {
		return nil
	}

	for _, o := range outages {
		o.ID = bson.NewObjectId()
		err := db.Insert(o)
		if err != nil {
			if !strings.Contains(err.Error(), "outage_unq dup key") {
				return err
			}
		}
	}

	return nil
}
