package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"my-module/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Seed - struct
func Seed(db *gorm.DB) error {
	log.Print("Seeding Database....")
	// LoadSeed/Populate Database from JSON file
	err := seedSecurityCheckSpecs(db)
	if err != nil {
		return fmt.Errorf("Failed to seed SecurityCheckSpecs: %v", err)
	}
	err = seedMobileApplications(db)
	if err != nil {
		return fmt.Errorf("Failed to seed SecurityCheckSpecs: %v", err)
	}
	return nil
}

func seedSecurityCheckSpecs(db *gorm.DB) error {
	log.Print("Seeding SecurityCheckSpecs....")
	data, err := ioutil.ReadFile("./database/security_check_specs.json")
	if err != nil {
		return fmt.Errorf("Failed to read seed json data file: %v", err)
	}
	// unmarshall it
	var securityCheckSpecs model.SecurityCheckSpecs
	err = json.Unmarshal(data, &securityCheckSpecs)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json data into go struct: %v", err)
	}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(securityCheckSpecs, 100) // batch size 100
	if result.Error != nil {
		return fmt.Errorf("Failed to seed data using gorm: %v", result.Error)
	}
	log.Printf("Seeded SecurityCheckSpec(s): %v", result.RowsAffected)
	return nil
}

func seedMobileApplications(db *gorm.DB) error {
	log.Print("Seeding MobileApplications....")
	data, err := ioutil.ReadFile("./database/mobile_applications.json")
	if err != nil {
		return fmt.Errorf("Failed to read seed json data file: %v", err)
	}
	// unmarshall it
	var mobileApplications model.MobileApplications
	err = json.Unmarshal(data, &mobileApplications)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json data into go struct: %v", err)
	}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(mobileApplications, 100) // batch size 100
	if result.Error != nil {
		return fmt.Errorf("Failed to seed data using gorm: %v", result.Error)
	}
	log.Printf("Seeded MobileApplications(s): %v", result.RowsAffected)
	return nil
}
