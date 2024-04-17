package league_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func GetLeagueStats(c *gin.Context) {

	leagueID := c.Query("id")
	season := c.Query("season")

	id, err := strconv.Atoi(leagueID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/standings?season=%s&league=%d", season, id)

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
