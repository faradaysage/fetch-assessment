package server

import (
	"context"
	"errors"
	"fetch-assessment/api"
	mapper "fetch-assessment/mappers"
	"fetch-assessment/repository"
	"fetch-assessment/rules"
	"fetch-assessment/validation"
	"fmt"
)

// implement the generated StrictServerInterface
type Server struct {
	Repository repository.Repository
}

// create an instance of Server
func NewServer(repository repository.Repository) *Server {
	return &Server{
		Repository: repository,
	}
}

// if we had excessive handlers, we'd break each out out into a separate file
// but for cohesion, and because the endpoints represent a logical grouping, i'm implementing both handlers in this one file

// Submits a receipt for processing.
// (POST /receipts/process)
func (s *Server) PostReceiptsProcess(ctx context.Context, request api.PostReceiptsProcessRequestObject) (api.PostReceiptsProcessResponseObject, error) {
	// ensure the receipt exists
	if request.Body == nil {
		return nil, errors.New("invalid request: missing receipt")
	}

	validationEngine := validation.NewReceiptValidationEngine()
	if !validationEngine.IsValid(*request.Body) {
		return nil, fmt.Errorf("receipt failed validation")
	}

	// api.Receipt to rules.Receipt (business object)
	receipt, err := mapper.MapToReceipt(*request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to map receipt: %w", err)
	}

	// save the receipt to the repository
	id, err := s.Repository.SaveReceipt(receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to save receipt: %w", err)
	}

	// return the database ID for the saved receipt
	response := api.PostReceiptsProcess200JSONResponse{Id: id}
	return response, nil

}

// Returns the points awarded for the receipt.
// (GET /receipts/{id}/points)
func (s *Server) GetReceiptsIdPoints(ctx context.Context, request api.GetReceiptsIdPointsRequestObject) (api.GetReceiptsIdPointsResponseObject, error) {
	// no need to validate id because it's validated by the generated handler
	// load the receipt from the repository
	receipt, err := s.Repository.LoadReceipt(request.Id)
	if err != nil {
		// receipt was not found
		return CustomGetReceiptsIdPoints404Response{}, nil
	}

	// create a rules engine to process the receipt
	rulesEngine := rules.NewRulesEngine()

	// calculate the total points
	points := rulesEngine.CalculateTotalPoints(receipt)

	// return the total points calculated
	response := api.GetReceiptsIdPoints200JSONResponse{Points: &points}
	return response, nil
}
