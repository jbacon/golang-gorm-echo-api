package handlers

import (
	"fmt"
	"log"
	"my-module/middleware"
	"my-module/model"
	"net/http"

	"github.com/labstack/echo"
)

// GetReports - "/apps/:app/version/:version/reports"
func GetReports(context echo.Context) error {
	customContext := context.(*middleware.DatabaseEchoContext)
	app := customContext.Param("app")
	version := customContext.Param("version")
	log.Printf("App: %v.. Version: %v", app, version)
	var reports model.SecurityReports
	result := customContext.DB.
		Where(&model.SecurityReport{
			App:     app,
			Version: version,
		}).
		Preload("Tests.SecurityCheckSpec").
		Omit("MobileApplication").
		Omit("MobileAppVersion").
		Find(&reports)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to query database")
	}
	log.Printf("Results Rows Affected: %v. Error: %v. ", result.RowsAffected, result.Error)
	// Get Data from Database using GORM
	return customContext.JSON(http.StatusOK, reports)
}

// Not being used at the moment, I had some questions regarding requirement's structures
type postReportRequestBody struct {
	ID          string `json:"id"`
	BundleID    string `json:"bundle_id"`
	Name        string `json:"name"`
	Platform    string `json:"platform"`
	CreatedDate string `json:"createdDate"`
	Tests       struct {
		ShortName  string      `json:"short_name"`
		Vulnerable bool        `json:"vulnerable"`
		Data       interface{} `json:"data"`
	} `json:"tests"`
}

// PostReport - "/apps/:app/versions/:version/reports"
func PostReport(context echo.Context) error {
	customContext := context.(*middleware.DatabaseEchoContext)
	app := customContext.Param("app")
	version := customContext.Param("version")
	// Find Corresponding Mobile App/Version (check existance)
	var mobileAppVersion model.MobileAppVersion
	result := customContext.DB.
		Where("MobileApplication.name=?", app).
		Where("mobile_app_versions.appstore_version=?", version).
		Joins("MobileApplication").
		Find(&mobileAppVersion)
	if result.Error != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to query database, error: %v", result.Error))
	}
	if result.RowsAffected == 0 {
		return echo.NewHTTPError(
			http.StatusNotFound,
			fmt.Sprintf(
				"Failed to find app %v with version %v,m error: %v",
				app,
				version,
				result.Error))
	}
	// Bind Request Body to struct
	var requestBody model.SecurityReport
	if err := customContext.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Request Body invalid format: %v", err))
	}
	requestBody.App = app
	requestBody.Version = version
	// Add Database Records
	result = customContext.DB.
		Omit("MobileApplication").
		Omit("MobileAppVersion").
		Omit("SecurityCheckSpec").
		Create(&requestBody)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to add report: %v", result.Error))
	}
	return customContext.String(
		http.StatusOK,
		fmt.Sprintf("Created New Security Report w/ ID: %v", requestBody.ID),
	)
}
