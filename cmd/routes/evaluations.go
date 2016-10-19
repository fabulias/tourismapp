type evaluation struct {
	Id       int64     `json:"id" binding:"required"`
	Id_user  string    `json:"id_user" binding:"required"`
	Id_place string    `json:"id_place" binding:"required"`
	Score    int64     `json:"score" binding:"required"`
	Comment  string    `json:"comment" binding:"required"`
	Date     time.Time `json:"date" binding:required`
}
