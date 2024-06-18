package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// User представляет структуру пользователя
type User struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}

var (
	users      = make(map[int]*User) // мапа для хранения пользователей
	usersMutex sync.Mutex            // мьютекс для безопасного доступа к мапе
	idCounter  = 0                   // счетчик для генерации уникальных ID пользователей
)

// CreateHandler обработчик для создания нового пользователя
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	idCounter++
	newUser.ID = idCounter
	users[idCounter] = &newUser
	usersMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User ID: %d\n", newUser.ID)
}

// MakeFriendsHandler обработчик для создания друзей
func MakeFriendsHandler(w http.ResponseWriter, r *http.Request) {
	var friendData struct {
		SourceID int `json:"source_id"`
		TargetID int `json:"target_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&friendData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sourceUser, ok := users[friendData.SourceID]
	if !ok {
		http.Error(w, "Source user not found", http.StatusNotFound)
		return
	}

	targetUser, ok := users[friendData.TargetID]
	if !ok {
		http.Error(w, "Target user not found", http.StatusNotFound)
		return
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.Name)
	targetUser.Friends = append(targetUser.Friends, sourceUser.Name)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s and %s are now friends\n", sourceUser.Name, targetUser.Name)
}

// DeleteUserHandler обработчик для удаления пользователя
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData struct {
		TargetID int `json:"target_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, ok := users[userData.TargetID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Удаление пользователя из мапы
	delete(users, userData.TargetID)

	for _, u := range users {
		for i, friend := range u.Friends {
			if friend == user.Name {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				break
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s deleted\n", user.Name)
}

// GetFriendsHandler обработчик для получения списка друзей пользователя
func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/friends/"):]
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, ok := users[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Friends of user %d: %v\n", id, user.Friends)
}

// UpdateAgeHandler обработчик для обновления возраста пользователя
func UpdateAgeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[1:]
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var ageData struct {
		NewAge int `json:"new_age"`
	}

	err = json.NewDecoder(r.Body).Decode(&ageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, ok := users[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Age = ageData.NewAge

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User's age successfully updated")
}

func main() {
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/make_friends", MakeFriendsHandler)
	http.HandleFunc("/delete", DeleteUserHandler)
	http.HandleFunc("/friends/", GetFriendsHandler)
	http.HandleFunc("/", UpdateAgeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
