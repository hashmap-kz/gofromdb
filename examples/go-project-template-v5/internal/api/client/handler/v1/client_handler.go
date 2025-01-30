package v1

import (
	"fmt"
	"net/http"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/client/dto"
	"go-project-template-v5/internal/api/client/service"

	"go-project-template-v5/pkg/httputils"
	"go-project-template-v5/pkg/validator"
)

type ClientHTTPHandler struct {
	clientService service.ClientService
}

func NewClientHTTPHandler(clientService service.ClientService) *ClientHTTPHandler {
	return &ClientHTTPHandler{
		clientService: clientService,
	}
}

// Save
//
// @Summary Create new Client
// @Description Create Client handler
// @Tags Client
// @Accept json
// @Produce json
// @Param request body clientCreateRequest true "Create input"
// @Success 201 {object} clientResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/clients [post]
func (h *ClientHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	// read RequestBody
	req := &clientCreateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// check RequestBody
	if err := validator.ValidateStruct(req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, err := h.clientService.Save(r.Context(), &dto.ClientCreateDto{
		Email: req.Email,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

// GetAll retrieves all Client.
//
// @Summary Get all Client
// @Description Retrieves a list of all Client without pagination.
// @Tags Client
// @Accept json
// @Produce  json
// @Success 200 {object} clientResponseList "List of all BuyItems"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/clients [get]
func (h *ClientHTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call service
	resp, err := h.clientService.GetAll(r.Context())
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtosToPayloads, err := mapDtosToPayloads(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, clientResponseList{
		Data: dtosToPayloads,
	})
}

// GetAllPaginated retrieves a paginated list of Client.
//
// @Summary Get paginated list of Client
// @Description Retrieves a paginated list of Client with pagination parameters.
// @Tags Client
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Number of items per page (default: 10)"
// @Param sort query string false "Sort order, e.g., 'name,asc'"
// @Success 200 {object} clientResponseList "Paginated list of Client"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid pagination parameters or service failure)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Data processing failure)"
// @Router /api/v1/clients/pageable [get]
func (h *ClientHTTPHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
	pq, err := pageable.GetPaginationFromCtx(r)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// call service
	resp, page, err := h.clientService.GetAllPaginated(r.Context(), pq)
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtosToPayloads, err := mapDtosToPayloads(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 200 OK
	httputils.WriteJSON(w, http.StatusOK, clientResponseList{
		Data: dtosToPayloads,
		Page: &page,
	})
}

// Update
//
// @Summary Update existing Client
// @Description Update Client handler
// @Tags Client
// @Accept json
// @Produce json
// @Param request body clientUpdateRequest true "Update input"
// @Success 201 {object} clientResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error"
// @Router /api/v1/clients [put]
func (h *ClientHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	req := &clientUpdateRequest{}
	if err := httputils.ReadJSON(r, &req); err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// TODO: types - int(id)
	resp, err := h.clientService.Update(r.Context(), int(id), &dto.ClientUpdateDto{
		Email: req.Email,
	})
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// convert service-model to handler-payload
	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 201 OK
	httputils.WriteJSON(w, http.StatusCreated, dtoToPayload)
}

// Delete handles the deletion of a Client.
//
// @Summary Delete a Client
// @Description Deletes a Client by its ID
// @Tags Client
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Success 204 "No Content (Successfully deleted)"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request (Invalid ID format)"
// @Failure 500 {object} httputils.ErrorResponse "Internal Server Error (Deletion failed)"
// @Router /api/v1/clients [delete]
func (h *ClientHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.clientService.Delete(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	// 204 OK
	w.WriteHeader(http.StatusNoContent)
}

func (h *ClientHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.PathValueI64(r, "id")
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := h.clientService.GetByID(r.Context(), int(id))
	if err != nil {
		httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	dtoToPayload, err := mapDtoToPayload(resp)
	if err != nil {
		httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
		return
	}

	httputils.WriteJSON(w, http.StatusOK, dtoToPayload)
}

// mappers

func mapDtosToPayloads(inputDtos []dto.ClientDto) ([]clientResponse, error) {
	var outputResponses []clientResponse
	for _, inputDto := range inputDtos {
		toPayload, err := mapDtoToPayload(&inputDto)
		if err != nil {
			return nil, err
		}
		outputResponses = append(outputResponses, toPayload)
	}
	return outputResponses, nil
}

func mapDtoToPayload(inputDto *dto.ClientDto) (clientResponse, error) {
	if inputDto == nil {
		return clientResponse{}, fmt.Errorf("unexpected nil input for mapping between ClientDto->clientResponse")
	}
	return clientResponse{
		RecordID:  inputDto.RecordID,
		Email:     inputDto.Email,
		CreatedAt: inputDto.CreatedAt,
		UpdatedAt: inputDto.UpdatedAt,
		Guid:      inputDto.Guid,
	}, nil
}
