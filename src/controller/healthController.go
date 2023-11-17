package controller

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}