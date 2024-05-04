package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"backend/domain"
	"backend/usecase"
)

type MemoHandler interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type memoHandler struct {
	memoUsecase usecase.MemoUsecase
}

func NewMemoHandler(memoUsecase usecase.MemoUsecase) MemoHandler {
	return &memoHandler{memoUsecase}
}

type requestMemo struct {
	Title   string       `json:"title"`
	Content string       `json:"content"`
	Pinned  bool         `json:"pinned"`
	Parent  *domain.Memo `json:"parent"`
}

type responseMemo struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	ShortHash string        `json:"short_hash"`
	Pinned    bool          `json:"pinned"`
	ViewCount uint          `json:"view_count"`
	Parent    *domain.Memo  `json:"parent"`
	Children  []domain.Memo `json:"children"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (h *memoHandler) Create(c *gin.Context) {
	var req requestMemo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memo := domain.Memo{
		Title:   req.Title,
		Content: req.Content,
		Pinned:  req.Pinned,
		Parent:  req.Parent,
	}

	if err := h.memoUsecase.Create(&memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := responseMemo{
		ID:        memo.ID,
		Title:     memo.Title,
		Content:   memo.Content,
		ShortHash: memo.ShortHash,
		Pinned:    memo.Pinned,
		ViewCount: memo.ViewCount,
		Parent:    memo.Parent,
		Children:  memo.Children,
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *memoHandler) Read(c *gin.Context) {
	hash := c.Param("hash")

	memo, err := h.memoUsecase.Read(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := responseMemo{
		ID:        memo.ID,
		Title:     memo.Title,
		Content:   memo.Content,
		ShortHash: memo.ShortHash,
		Pinned:    memo.Pinned,
		ViewCount: memo.ViewCount,
		Parent:    memo.Parent,
		Children:  memo.Children,
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

func (h *memoHandler) ReadAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	per, _ := strconv.Atoi(c.DefaultQuery("per", "10"))

	memos, err := h.memoUsecase.ReadAll(page, per)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ress := make([]responseMemo, len(memos))
	for i, memo := range memos {
		ress[i] = responseMemo{
			ID:        memo.ID,
			Title:     memo.Title,
			Content:   memo.Content,
			ShortHash: memo.ShortHash,
			Pinned:    memo.Pinned,
			ViewCount: memo.ViewCount,
			Parent:    memo.Parent,
			Children:  memo.Children,
			CreatedAt: memo.CreatedAt,
			UpdatedAt: memo.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, ress)
}

func (h *memoHandler) Update(c *gin.Context) {
	hash := c.Param("hash")

	var req requestMemo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memo, err := h.memoUsecase.Read(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	memo.Title = req.Title
	memo.Content = req.Content
	memo.Pinned = req.Pinned
	memo.Parent = req.Parent

	if err := h.memoUsecase.Update(memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := responseMemo{
		ID:        memo.ID,
		Title:     memo.Title,
		Content:   memo.Content,
		ShortHash: memo.ShortHash,
		Pinned:    memo.Pinned,
		ViewCount: memo.ViewCount,
		Parent:    memo.Parent,
		Children:  memo.Children,
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

func (h *memoHandler) Delete(c *gin.Context) {
	hash := c.Param("hash")

	memo, err := h.memoUsecase.Read(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.memoUsecase.Delete(memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
