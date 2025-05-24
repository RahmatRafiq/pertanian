package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"golang_starter_kit_2025/app/helpers"
	"golang_starter_kit_2025/app/models"
	"golang_starter_kit_2025/app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
// @Success		200	{object}	helpers.ResponseParams[models.Farmer]{data=[]models.Farmer}
// @Router		/farmers [get]
func (c *FarmerController) List(ctx *gin.Context) {
	farmers, err := c.service.GetAllFarmers()
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal mendapatkan daftar Farmer",
			Reference: "ERROR-FARMER-1",
		}, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Farmer]{Data: &farmers}, http.StatusOK)
}

// @Summary		Show a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"Farmer ID"
// @Success		200	{object}	helpers.ResponseParams[models.Farmer]{item=models.Farmer}
// @Router		/farmers/{id} [get]
func (c *FarmerController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	farmerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"id": "Invalid id"},
			Message:   "Parameter tidak valid",
			Reference: "ERROR-FARMER-2",
		}, http.StatusBadRequest)
		return
	}
	farmer, err := c.service.Find(farmerID)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": "Farmer not found"},
			Message:   "Farmer tidak ditemukan",
			Reference: "ERROR-FARMER-3",
		}, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Farmer]{Item: &farmer}, http.StatusOK)
}

// @Summary		Upsert a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		JSON	body	models.Farmer	true	"Farmer object"
// @Success		201	{object}	helpers.ResponseParams[models.Farmer]{item=models.Farmer}
// @Router		/farmers [put]
func (c *FarmerController) Put(ctx *gin.Context) {
	var farmer models.Farmer
	if err := ctx.ShouldBindJSON(&farmer); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
				Errors:    helpers.ValidationError(verr),
				Message:   "Parameter tidak valid",
				Reference: "ERROR-FARMER-4",
			}, http.StatusBadRequest)
			return
		}
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal memproses data Farmer",
			Reference: "ERROR-FARMER-5",
		}, http.StatusBadRequest)
		return
	}
	updatedFarmer, err := c.service.Put(farmer)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal menyimpan data Farmer",
			Reference: "ERROR-FARMER-6",
		}, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Farmer]{Item: &updatedFarmer}, http.StatusCreated)
}

// @Summary		Delete a farmer
// @Tags		farmers
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"Farmer ID"
// @Success		200	{object}	helpers.ResponseParams[models.Farmer]{message=string}
// @Router		/farmers/{id} [delete]
func (c *FarmerController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	farmerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"id": "Invalid id"},
			Message:   "Parameter tidak valid",
			Reference: "ERROR-FARMER-7",
		}, http.StatusBadRequest)
		return
	}
	if err := c.service.Delete(farmerID); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": "Farmer not found"},
			Message:   "Farmer tidak ditemukan",
			Reference: "ERROR-FARMER-8",
		}, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Farmer]{Message: "Farmer deleted"}, http.StatusOK)
}
