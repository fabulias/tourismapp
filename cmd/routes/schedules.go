type schedules struct { //Horarios, 7 días de la semana, open-Close por cada día
	Id int64  `json:"id_place" binding:"required"`
	o1 string `json:"o1" binding:"required"`
	c1 string `json:"c1" binding:"required"`
	o2 string `json:"o2" binding:"required"`
	c2 string `json:"c2" binding:"required"`
	o3 string `json:"o3" binding:"required"`
	c3 string `json:"c3" binding:"required"`
	o4 string `json:"o4" binding:"required"`
	c4 string `json:"c4" binding:"required"`
	o5 string `json:"o5" binding:"required"`
	c5 string `json:"c5" binding:"required"`
	o6 string `json:"o6" binding:"required"`
	c6 string `json:"c6" binding:"required"`
	o7 string `json:"o7" binding:"required"`
	c7 string `json:"c7" binding:"required"`
}
