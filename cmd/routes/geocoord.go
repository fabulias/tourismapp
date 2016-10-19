type geocord struct {
	Id_place int64 `json:"id_place" binding:"required"`
	pos      int64 `json:"latitud" binding :"required"`
}
