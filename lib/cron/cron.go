package cron

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/edenorte"
	"cortes-programados-api/scrapers/edesur"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron"
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	Cron     *cron.Cron
	db       *lib.DBLib
	dbConfig *models.DatabaseConfig
}

func NewJob(db *models.DatabaseConfig) *Job {
	return &Job{
		Cron:     cron.New(),
		db:       lib.NewDBLib(db, "outages"),
		dbConfig: db,
	}
}

func (j *Job) UpdateOutagesCollection(schedule string) {
	fmt.Println("Running:", schedule)
	j.Cron.AddFunc(schedule, func() {
		started := time.Now()
		fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' started ***")
		fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' start time: %v ***\n", started)

		outages, err := GetOutagesScrapeData(j.dbConfig)
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'UpdateOutagesCollection' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'UpdateOutagesCollection' time elapsed: %v ***\n", ended.Sub(started))
		}

		err = SaveOutageCollection(j.db, outages)
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

func GetOutagesScrapeData(dbConfig *models.DatabaseConfig) ([]*models.Outage, error) {

	norte, err := edenorte.ReadOutageAnouncement()
	if err != nil {
		log.Println("Error fetching north", err)
		// return nil, err
	}

	sur, err := edesur.GetOutageAnouncement()
	if err != nil {
		log.Println("Error fetching south", err)
		// return nil, err
	}

	all := append(norte, sur...)

	setCircuitToOutage(dbConfig, all)

	return all, nil
}

func SaveOutageCollection(db *lib.DBLib, outages []*models.Outage) error {

	if len(outages) == 0 {
		log.Println("No outages")
		return nil
	}

	for _, o := range outages {
		log.Println("Outage: ", o.Province, o.Circuit)
		o.ID = bson.NewObjectId()
		err := db.Insert(o)
		if err != nil {
			log.Println("Error saving outage =>", err)
			if !strings.Contains(err.Error(), "outage_unq dup key") {
				return err
			}
		}
	}

	return nil
}

// Set Circuit name to outages that doesn't have it initially
// Right now the only ones without this field are the ones being
// taken from EDESur.
func setCircuitToOutage(db *models.DatabaseConfig, outages []*models.Outage) error {

	for _, o := range outages {

		if len(o.AffectedZones) < 1 {
			continue
		}

		query := bson.M{
			"province":        o.Province,
			"influence_zones": o.AffectedZones[0],
		}

		res := models.Circuit{}
		err := db.DB.C("circuits").Find(query).One(&res)
		if err != nil {
			log.Println("Error getting this circuit => ", o.Province, o.AffectedZones[0])
			log.Println("Error getting this circuit => ", query)
			continue
		}

		o.Circuit = res.Name
	}

	return nil
}
