package transport

import (
	"log"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/rogudator/tuvebes-backend/internal/repository"
	"github.com/rogudator/tuvebes-backend/internal/entity"
)

type Transport struct {
	repo *repository.Repository
}

func NewTransport(repo *repository.Repository) *Transport {
	return &Transport{
		repo: repo,
	}
}

func (h *Transport) CreateTuvebe(ctx *gin.Context) {
	var t entity.Tuvebe
	if err := ctx.BindJSON(&t); err != nil {
		b, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(b))
		log.Println("failed to unmarshall tuvebe:", err.Error())
		return
	}
	h.repo.DB = append(h.repo.DB, t)
}

func (h *Transport) GetTuvebes(ctx *gin.Context) {
	ctx.JSON(200, h.repo.GetTuvebes())
}

func (h *Transport) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "helloooo!!!!!!")
	})

	router.POST("tuvebes", h.CreateTuvebe)
	router.GET("tuvebes", h.GetTuvebes)

	return router
}