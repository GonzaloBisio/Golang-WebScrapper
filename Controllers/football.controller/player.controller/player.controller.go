package player_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func GetPalyerStats(c *gin.Context) {
	// Extract player ID and season from query parameters
	playerID := c.Query("id")
	season := c.Query("season")

	// Convert playerID to an integer
	id, err := strconv.Atoi(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	// Now you can use the `id` and `season` variables to fetch player stats
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=%s", id, season)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "2511100361msh02d2482535b154cp1331f6jsn9a404f42fd7a")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	// Check if the response status code is not OK
	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, gin.H{"error": fmt.Sprintf("Received non-200 status code: %d", res.StatusCode)})
		return
	}

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response to client
	c.Data(res.StatusCode, res.Header.Get("Content-Type"), body)
}
