package medicine

import (
  "errors"
  "github.com/gbrayhan/microservices-go/controllers"
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"

  "github.com/gbrayhan/microservices-go/models"
  errorModels "github.com/gbrayhan/microservices-go/models/errors"
)

// NewMedicine godoc
// @Tags medicine
// @Summary Create New Medicine
// @Description Create new medicine on the system
// @Accept  json
// @Produce  json
// @Param data body medicineController true "body data"
// @Success 200 {object} models.Medicine
// @Failure 400 {object} generalResponse
// @Failure 500 {object} generalResponse
// @Router /medicine/new [post]
func NewMedicine(c *gin.Context) {
  request := struct {
    Name        string `json:"name" example:"Paracetamol" gorm:"unique" binding:"required"`
    Description string `json:"description" example:"Something" binding:"required"`
    Laboratory  string `json:"laboratory" example:"Roche" binding:"required"`
    EanCode     string `json:"ean_code" example:"122000000021" gorm:"unique" binding:"required"`
  }{}

  if err := controllers.BindJSON(c, &request); err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    //appError := domainErrors.NewAppError(errors.New(strings.Join(messages, ", ")), domainErrors.ValidationError)
    _ = c.Error(appError)
    return
  }
  medicine := models.Medicine{
    Name:        request.Name,
    Description: request.Description,
    Laboratory:  request.Laboratory,
    EANCode:     request.EanCode,
  }

  err := models.CreateMedicine(&medicine)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, medicine)
}

// GetAllMedicines godoc
// @Tags medicine
// @Summary Get all Medicines
// @Description Get all Medicines on the system
// @Success 200 {object} []models.Medicine
// @Failure 400 {object} generalResponse
// @Failure 500 {object} generalResponse
// @Router /medicine/get-all [get]
func GetAllMedicines(c *gin.Context) {
  var medicines []models.Medicine
  err := models.GetAllMedicines(&medicines)
  if err != nil {
    controllers.ServerError(c, err)
    return
  }
  c.JSON(http.StatusOK, medicines)
}

// GetMedicinesByID godoc
// @Tags medicine
// @Summary Get medicines by ID
// @Description Get Medicines by ID on the system
// @Param medicine_id path int true "id of medicine"
// @Success 200 {object} models.Medicine
// @Failure 400 {object} generalResponse
// @Failure 500 {object} generalResponse
// @Router /medicine/get-by-id/{medicine_id} [get]
func GetMedicinesByID(c *gin.Context) {
  var medicine models.Medicine
  medicineID, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    appError := errorModels.NewAppError(errors.New("medicine id is invalid"), errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  err = models.GetMedicineByID(&medicine, medicineID)
  if err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  c.JSON(http.StatusOK, medicine)
}

func UpdateMedicine(c *gin.Context) {
  medicineID, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }
  var requestMap map[string]interface{}

  err = controllers.BindJSONMap(c, &requestMap)
  if err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  err = updateValidation(requestMap)
  if err != nil {
    _ = c.Error(err)
    return
  }



  medicine, err := models.UpdateMedicine(medicineID, requestMap)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, medicine)

}
