package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Rating      int    `json:"rating"` // Changed to int to match database
	Genre       string `json:"genre"`
}

func getDbConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:cept123*@localhost:5432/movie")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getMovies(c *gin.Context) {
	conn, err := getDbConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, title, description, year, ratings, genre FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.Rating, &movie.Genre)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse database response"})
			return
		}
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, movies)
}

func addMovie(c *gin.Context) {
	var movie Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	conn, err := getDbConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		"INSERT INTO movies (title, description, year, ratings, genre) VALUES ($1, $2, $3, $4, $5)",
		movie.Title, movie.Description, movie.Year, movie.Rating, movie.Genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie added successfully"})
}

func editMovie(c *gin.Context) {
	id := c.Param("id")

	var movie Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	conn, err := getDbConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		"UPDATE movies SET title = $1, description = $2, year = $3, ratings = $4, genre = $5 WHERE id = $6",
		movie.Title, movie.Description, movie.Year, movie.Rating, movie.Genre, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func deleteMovie(c *gin.Context) {
	id := c.Param("id")

	conn, err := getDbConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

func main() {
	router := gin.Default()

	router.GET("/movies", getMovies)
	router.POST("/add_movie", addMovie)
	router.PUT("/edit_movie/:id", editMovie)
	router.DELETE("/delete_movie/:id", deleteMovie)

	log.Fatal(router.Run(":8080"))
}
