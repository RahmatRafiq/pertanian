package controllers

import (
	"net/http"
	"strconv"

	"golang_starter_kit_2025/app/models"
	"golang_starter_kit_2025/app/services"

	"github.com/gin-gonic/gin"
)

type FarmerController struct {
	service services.FarmerService
}

func NewFarmerController(service services.FarmerService) *FarmerController {
	return &FarmerController{service: service}
}

// @Summary		Show all farmers
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Success		200	{array}	models.Farmer
// @Router		/farmers [get]
func (c *FarmerController) List(ctx *gin.Context) {
	farmers, err := c.service.GetAllFarmers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, farmers)
}

// @Summary		Show a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"User ID"
// @Success		200	{object}	models.Farmer
// @Router		/farmers/{id} [get]
func (c *FarmerController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	farmerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	farmer, err := c.service.Find(farmerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Farmer not found"})
		return
	}
	ctx.JSON(http.StatusOK, farmer)
}

// @Summary		Upsert a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		JSON	body	models.Farmer	true	"Farmer object"
// @Success		201	{object}	models.Farmer
// @Router		/farmers [put]
func (c *FarmerController) Put(ctx *gin.Context) {
	var farmer models.Farmer
	if err := ctx.ShouldBindJSON(&farmer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedFarmer, err := c.service.Put(farmer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, updatedFarmer)
}

// @Summary		Delete a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"User ID"
// @Success		200	{string}	string	"Farmer deleted"
// @Router		/farmers/{id} [delete]
func (c *FarmerController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	farmerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := c.service.Delete(farmerID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Farmer not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Farmer deleted"})
}
