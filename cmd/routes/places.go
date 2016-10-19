package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
)

type places struct {
	Id          int64     `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Score       int64     `json:"score" binding:"required"`
	User_c      string    `json:"user_c" binding:"required"`
	Date_c      time.Time `json:"fecha_c" binding:"required"`
	Descripcion string    `json: "descripcion" binding:"required"`
}
