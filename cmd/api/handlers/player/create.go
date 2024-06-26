package player

import (
	"net/http"

	"github.com/HexArq/internal/domain"
	"github.com/gin-gonic/gin"
)

/*
	A handler must have
	- Translate the request
	- Validation
	- Consume a service
	- Translate the response
*/

func (h Handler) Create(ginCtx *gin.Context) {
	var dPlayer domain.Player
	if errBindJSON := ginCtx.BindJSON(&dPlayer); errBindJSON != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": errBindJSON.Error()})
		return
	}

	insertedID, errCreatePLayer := h.PlayerService.Create(dPlayer)
	if errCreatePLayer != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	ginCtx.JSON(http.StatusOK, gin.H{"player_id": insertedID})
}
