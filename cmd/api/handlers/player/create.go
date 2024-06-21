package player

import (
	"net/http"

	"github.com/HexArq/internal/domain"
	"github.com/HexArq/internal/services/player"
	"github.com/gin-gonic/gin"
)

/*
	A handler must have
	- Translate the request
	- Validation
	- Consume a service
	- Translate the response
*/

func CreatePlayer(ginCtx *gin.Context) {
	var dPlayer domain.Player
	if errBindJSON := ginCtx.BindJSON(&dPlayer); errBindJSON != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": errBindJSON.Error()})
		return
	}

	insertedID, errCreatePLayer := player.CreatePLayerService(dPlayer)
	if errCreatePLayer != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	ginCtx.JSON(http.StatusOK, gin.H{"player_id": insertedID})
}
