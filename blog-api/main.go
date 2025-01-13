package main

import (
	"blog-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Post{}, &models.Comment{})
	log.Println("Database initialized and models migrated!")

}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.Create(&post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []models.Post
	db.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var updatedPost models.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	db.Save(&post)

	json.NewEncoder(w).Encode(post)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	db.Delete(&post)
	w.WriteHeader(http.StatusNoContent)
}

func addCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.Create(&comment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postID, _ := strconv.Atoi(mux.Vars(r)["post_id"])
	var comments []models.Comment
	db.Where("post_id = ?", postID).Find(&comments)
	json.NewEncoder(w).Encode(comments)
}

func setupRoutes() {
	router := mux.NewRouter()

	// Post routes
	router.HandleFunc("/posts", createPostHandler).Methods("POST")
	router.HandleFunc("/posts", getPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePostHandler).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePostHandler).Methods("DELETE")

	// Comment routes
	router.HandleFunc("/posts/{post_id}/comments", addCommentHandler).Methods("POST")
	router.HandleFunc("/posts/{post_id}/comments", getCommentsHandler).Methods("GET")

	http.Handle("/", router)
}

func main() {
	initDB()
	setupRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
