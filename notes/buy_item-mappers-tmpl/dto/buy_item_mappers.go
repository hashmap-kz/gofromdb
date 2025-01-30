package dto

import (
	"fmt"

	dbModel "go-project-template-v5/internal/api/buy_item/entity/postgres"
)

// Service layer

func FromCreateDtoToEntity(input *BuyItemCreateDto) (*dbModel.BuyItem, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyItemCreateDto->BuyItem: input dto cannot be nil")
	}
	return &dbModel.BuyItem{
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func FromUpdateDtoToEntity(input *BuyItemUpdateDto) (*dbModel.BuyItem, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyItemUpdateDto->BuyItem: input dto cannot be nil")
	}
	return &dbModel.BuyItem{
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func FromEntitiesToDtos(inputEntities []dbModel.BuyItem) ([]BuyItemDto, error) {
	var outputDtos []BuyItemDto
	for _, inputEntity := range inputEntities {
		toDto, err := FromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func FromEntityToDto(inputEntity *dbModel.BuyItem) (BuyItemDto, error) {
	if inputEntity == nil {
		return BuyItemDto{}, fmt.Errorf("unexpected nil input for mapping between BuyItem->BuyItemDto")
	}
	return BuyItemDto{
		RecordID:  inputEntity.RecordID,
		BuyID:     inputEntity.BuyID,
		ProductID: inputEntity.ProductID,
		Quantity:  inputEntity.Quantity,
		Price:     inputEntity.Price,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}

// Handler layer

func FromDtosToPayloads(inputDtos []BuyItemDto) ([]BuyItemResponse, error) {
	var outputResponses []BuyItemResponse
	for _, inputDto := range inputDtos {
		toPayload, err := FromDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func FromDtoToPayload(inputDto *BuyItemDto) (BuyItemResponse, error) {
	if inputDto == nil {
		return BuyItemResponse{}, fmt.Errorf("unexpected nil input for mapping between BuyItemDto->buyItemResponse")
	}
	return BuyItemResponse{
		RecordID:  inputDto.RecordID,
		BuyID:     inputDto.BuyID,
		ProductID: inputDto.ProductID,
		Quantity:  inputDto.Quantity,
		Price:     inputDto.Price,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		Guid:      inputDto.Guid,
	}, nil
}
