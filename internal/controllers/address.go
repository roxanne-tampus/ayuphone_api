package controllers

import (
	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac ApiController) GetRegions(c *gin.Context) {
	var regions []models.Region
	if err := ac.DbService.GetAllRegions(c, &regions); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to fetch regions")
		return
	}
	utils.JSONResponse(c, true, "Regions retrieved successfully", regions)
}

func (ac ApiController) GetProvinces(c *gin.Context) {
	var provinces []models.Province
	regionID := c.Param("region_id")

	// Validate regionID
	if regionID == "" {
		utils.JSONResponse(c, false, "Region ID is required", nil)
		return
	}

	// Fetch from DB
	if err := ac.DbService.GetProvincesByRegionID(c, regionID, &provinces); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to fetch provinces")
		return
	}
	utils.JSONResponse(c, true, "Provinces retrieved successfully", provinces)
}

func (ac ApiController) GetMunicipalities(c *gin.Context) {
	var municipalities []models.Municipality
	provinceID := c.Param("province_id")

	// Validate provinceID
	if provinceID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Province ID is required")
		return
	}

	// Fetch from DB
	if err := ac.DbService.GetMunicipalitiesByProvinceID(c, provinceID, &municipalities); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to fetch municipalities")
		return
	}
	utils.JSONResponse(c, true, "Municipalities retrieved successfully", municipalities)
}

func (ac ApiController) GetBarangays(c *gin.Context) {
	var barangays []models.Barangay
	municipalityID := c.Param("municipality_id")

	// Validate municipalityID
	if municipalityID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Municipality ID is required")
		return
	}

	// Fetch from DB
	if err := ac.DbService.GetBarangaysByMunicipalityID(c, municipalityID, &barangays); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to fetch barangays")
		return
	}
	utils.JSONResponse(c, true, "Barangays retrieved successfully", barangays)
}
