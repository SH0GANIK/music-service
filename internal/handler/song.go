package handler

import (
	"github.com/gin-gonic/gin"
	"music-service/internal/model"
	"net/http"
	"strconv"
)

//@Summary Create song
//@Tags song
//@Accept json
//@Produce json
//@Param song body model.Song true "Song"
//@Success 200 {object} model.Song
//@Failure 400 {object} errorResponse
//@Failure 500 {object} errorResponse
//@Router /api/songs [post]

func (h *Handler) CreateSong(c *gin.Context) {
	ctx := c.Request.Context()
	var song model.Song
	if err := c.BindJSON(&song); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if song.Group == "" || song.Song == "" {
		NewErrorResponse(c, http.StatusBadRequest, "group or song is empty")
		return
	}
	id, err := h.service.CreateSong(ctx, &song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get all songs
// @Tags songs
// @Description Get a list of songs with optional pagination and filtering.
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Page size (default: 10)"
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song"
// @Param releaseDate query string false "Filter by release date"
// @Success 200 {object} model.PaginatedSongsResponse
// @Failure 500 {object} errorResponse
// @Router /songs [get]

func (h *Handler) GetSongs(c *gin.Context) {
	ctx := c.Request.Context()
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	filters := make(map[string]string)
	if c.Query("group") != "" {
		filters["group"] = c.Query("group")
	}
	if c.Query("song") != "" {
		filters["song"] = c.Query("song")
	}
	if c.Query("releaseDate") != "" {
		filters["releaseDate"] = c.Query("releaseDate")
	}
	paginatedResponse, err := h.service.GetSongs(ctx, page, pageSize, filters)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, paginatedResponse)
}

// @Summary Get song text
// @Tags songs
// @Description Get the text of a song with optional pagination.
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Page size (default: 10)"
// @Param id path string true "Song ID"
// @Success 200 {object} model.PaginatedSongTextResponse
// @Failure 500 {object} errorResponse
// @Router /songs/{id}/text [get]

func (h *Handler) GetSongText(c *gin.Context) {
	ctx := c.Request.Context()
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	paginatedResponse, err := h.service.GetSongText(ctx, page, pageSize, c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, paginatedResponse)
}

// @Summary Update song
// @Tags song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param song body model.Song true "Song"
// @Success 200
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/songs/{id} [put]

func (h *Handler) UpdateSong(c *gin.Context) {
	ctx := c.Request.Context()
	var song model.Song
	if err := c.BindJSON(&song); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	song.ID = c.Param("id")
	err := h.service.UpdateSong(ctx, &song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Delete song
// @Tags song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200
// @Failure 500 {object} errorResponse
// @Router /api/songs/{id} [delete]

func (h *Handler) DeleteSong(c *gin.Context) {
	ctx := c.Request.Context()
	err := h.service.DeleteSong(ctx, c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
