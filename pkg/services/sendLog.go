package services

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"loggerservice/pkg/config"
	"loggerservice/pkg/data"
	"loggerservice/pkg/data/entity"
	"loggerservice/pkg/utils"
	"time"
)

type SuccessLogRequest struct {
	Collection     string `json:"collection" validate:"required"`
	Source         string `json:"source"`
	Request        string `json:"request"`
	RequestHeader  string `json:"request_header"`
	Response       string `json:"response"`
	ResponseHeader string `json:"response_header"`
}

type ErrLogRequest struct {
	Collection string `json:"collection" validate:"required"`
	Level      string `json:"level"`
	Source     string `json:"source"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

// SendSuccessLog sends success log to mongo
func SendSuccessLog(c *fiber.Ctx) error {

	request := c.Request()

	var req SuccessLogRequest

	err := json.Unmarshal(request.Body(), &req)
	utils.LogErr("Error while unmarshal request body", err)

	validate := validator.New()
	err = validate.Struct(req)
	utils.LogErr("Error while validate request body", err)

	client := data.InitDB()
	defer data.CloseDB(client)

	db := client.Database(config.EnvConfigs.MongoDb)
	collection := db.Collection(req.Collection)

	log := entity.SuccessLog{
		Source:         req.Source,
		Request:        req.Request,
		RequestHeader:  req.RequestHeader,
		Response:       req.Response,
		ResponseHeader: req.ResponseHeader,
		Timestamp:      time.Now(),
	}

	result, err := collection.InsertOne(c.Context(), log)
	utils.LogErr("Error while insert success log", err)

	insertedID := result.InsertedID

	response := Response{
		Status: "success",
		Data:   make(map[string]interface{}),
	}

	response.Data["inserted_id"] = insertedID
	response.Data["collection"] = req.Collection

	return c.Status(fiber.StatusOK).JSON(response)
}

// SendErrLog sends error log to mongo
func SendErrLog(c *fiber.Ctx) error {

	request := c.Request()

	var req ErrLogRequest

	err := json.Unmarshal(request.Body(), &req)
	utils.LogErr("Error while unmarshal request body", err)

	validate := validator.New()
	err = validate.Struct(req)
	utils.LogErr("Error while validate request body", err)

	client := data.InitDB()
	defer data.CloseDB(client)

	db := client.Database(config.EnvConfigs.MongoDb)
	collection := db.Collection(req.Collection)

	log := entity.ErrLog{
		Level:     req.Level,
		Source:    req.Source,
		Message:   req.Message,
		Error:     req.Error,
		Timestamp: time.Now(),
	}

	result, err := collection.InsertOne(c.Context(), log)
	utils.LogErr("Error while insert success log", err)

	insertedID := result.InsertedID

	response := Response{
		Status: "success",
		Data:   make(map[string]interface{}),
	}

	response.Data["inserted_id"] = insertedID
	response.Data["collection"] = req.Collection

	return c.Status(fiber.StatusOK).JSON(response)
}
