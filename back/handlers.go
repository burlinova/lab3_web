package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func renderHTML(htmlName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, path.Join("templates/pages", htmlName))
		if err != nil {
			fmt.Printf("error parsing fs: %s", err.Error())
		}
		uname := r.Context().Value(key("username"))
		if uname == nil {
			uname = ""
		}

		tmpl.Execute(w, map[string]any{
			"User": uname,
		})
	}
}

func routereg(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFS(content, path.Join("templates", "register.html"))
			if err != nil {
				fmt.Printf("error routereg ParseFS: %s", err.Error())
			}
			tmpl.Execute(w, nil)
		}(w, r)
	} else if r.Method == http.MethodPost {
		POSTregHandler(w, r)
	}
}

func POSTregHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
		return
	}

	// Извлекаем значения по ключам
	username := r.FormValue("login")
	password := r.FormValue("password")
	email := r.FormValue("email")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Ошибка при хешировании пароля", http.StatusInternalServerError)
		return
	}
	if err = PGUserCreate(username, string(hash), email); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при сохранении пользователя. %s", err.Error()), http.StatusInternalServerError)
		return
	}
	// Устанавливаем куку для авторизации
	cookie := http.Cookie{
		Name:     "access-key",
		Value:    username,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	err = WriteEncrypted(w, cookie, secretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Обработчик логина для установки куки
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
		return
	}

	// Извлекаем значения по ключам
	username := r.FormValue("username")
	password := r.FormValue("password")
	hash, err := PGFindPassword(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Неверный логин. %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		return
	}
	// Устанавливаем куку для авторизации
	cookie := http.Cookie{
		Name:     "access-key",
		Value:    username,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	err = WriteEncrypted(w, cookie, secretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

type key string

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value, _ := ReadEncrypted(r, "access-key", secretKey)
		ctx := context.WithValue(context.Background(), key("username"), value)
		// Если кука есть, вызываем следующий обработчик
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" || !(token == expectedToken) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

const expectedToken = "supersecret"

// Обработчик логаута для удаления куки
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "access-key",
		Value:  "",
		MaxAge: -1, // Удаление куки
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

// Admin handlers

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	EMail    string `json:"email"`
}

// GetUserNameByID godoc
//
//		@Summary		Получить имя пользователя, по его айди
//		@Description	Ищет пользователя в базе данных по айди, и возвращает его имя
//		@Produce		plain/text
//		@Param       id             query    string true "ID пользователя"
//		@Success		200
//		@Router			/admin/get-username-by-id [get]
//	 @Security BearerAuth
func GetUserNameByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := pg.QueryRow(context.TODO(), "SELECT username FROM users WHERE id = $1", id)
	var name string
	if err := row.Scan(&name); err != nil {
		w.Write([]byte("User not founded"))
		return
	}
	w.Write([]byte(name))
}

// GetAllUsers godoc
//
//		@Summary		Получить всех пользователей из базы
//		@Description	Достает из базы всех пользователей
//		@Produce		application/json
//		@Success		200
//		@Router			/admin/get-all-users [get]
//	 @Security BearerAuth
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := pg.Query(context.TODO(), "SELECT id, username, email FROM users")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	users := make([]user, 0)
	for rows.Next() {
		var user user
		rows.Scan(&user.ID, &user.Username, &user.EMail)
		users = append(users, user)
	}
	bytes, _ := json.Marshal(&users)
	w.Header().Set("Content-Type", "json/application")
	w.Write(bytes)
}

// GetAllNames godoc
//
//		@Summary		Получить всех имена пользователей из базы
//		@Description	Достает из БД всех имена пользователей
//		@Produce		text/csv
//		@Success		200
//		@Router			/admin/get-all-names [get]
//	 @Security BearerAuth
func GetAllNames(w http.ResponseWriter, r *http.Request) {
	rows, err := pg.Query(context.TODO(), "SELECT  username FROM users")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	usersnames := make([]string, 0)
	for rows.Next() {
		var username string
		rows.Scan(&username)
		usersnames = append(usersnames, username)
	}
	bytes := make([]byte, 0)
	for _, u := range usersnames {
		bytes = append(bytes, []byte(u)...)
		bytes = append(bytes, []byte("\n")...)
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Write([]byte(bytes))
}

// GetAllIDs godoc
//
//		@Summary		Получить все ID пользователей из базы
//		@Description	Достает из БД всех ID пользователей
//		@Produce		text/csv
//		@Success		200
//		@Router			/admin/get-all-ids [get]
//	 @Security BearerAuth
func GetAllIDs(w http.ResponseWriter, r *http.Request) {
	rows, err := pg.Query(context.TODO(), "SELECT  username FROM users")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	ids := make([]int, 0)
	for rows.Next() {
		var id int
		rows.Scan(&ids)
		ids = append(ids, id)
	}

	bytes := make([]byte, 0)
	for _, u := range ids {
		bytes = append(bytes, []byte(fmt.Sprintf("%d", u))...)
		bytes = append(bytes, []byte("\n")...)
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Write([]byte(bytes))
}

// GetAllEMails godoc
//
//		@Summary		Получить все почты пользователей из базы
//		@Description	Достает из БД все почты пользователей
//		@Produce		text/csv
//		@Success		200
//		@Router			/admin/get-all-emails [get]
//	 @Security BearerAuth
func GetAllEMails(w http.ResponseWriter, r *http.Request) {
	rows, err := pg.Query(context.TODO(), "SELECT  email FROM users")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	emails := make([]string, 0)
	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails = append(emails, email)
	}

	bytes := make([]byte, 0)
	for _, u := range emails {
		bytes = append(bytes, []byte(u)...)
		bytes = append(bytes, []byte("\n")...)
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Write([]byte(bytes))
}

// GetUsersCount godoc
//
//		@Summary		Получить количество всех пользователей из базы
//		@Description	Достает из базы количество всех пользовательей
//		@Produce		plain/text
//		@Success		200
//		@Router			/admin/get-all-users-count [get]
//	 @Security BearerAuth
func GetUsersCount(w http.ResponseWriter, r *http.Request) {
	row := pg.QueryRow(context.TODO(), "SELECT COUNT(*) FROM users")
	var count int
	row.Scan(&count)
	w.Write([]byte(fmt.Sprintf("%d", count)))
}

// GetUserByID godoc
//
//		@Summary		Получить пользователя по айди
//		@Description	Ищет пользователя в базе данных по айди, и возвращает всю информацию о нем
//		@Produce		application/json
//		@Param       id             query    string true "ID пользователя"
//		@Success		200		{object}	user
//		@Router			/admin/get-user-by-id [get]
//	 	@Security BearerAuth
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := pg.QueryRow(context.TODO(), "SELECT id, username, email FROM users WHERE id = $1", id)
	var user user
	row.Scan(&user.ID, &user.Username, &user.EMail)
	bytes, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "json/application")
	w.Write(bytes)
}

// GetUserByName godoc
//
//		@Summary		Получить пользователя по имени
//		@Description	Ищет пользователя в базе данных по имени, и возвращает всю информацию о нем
//		@Produce		application/json
//		@Param       name             query    string true "Имя пользователя"
//		@Success		200		{object}	user
//		@Router			/admin/get-user-by-name [get]
//	 	@Security BearerAuth
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	row := pg.QueryRow(context.TODO(), "SELECT id, username, email FROM users WHERE name = $1", name)
	var user user
	row.Scan(&user.ID, &user.Username, &user.EMail)
	bytes, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "json/application")
	w.Write(bytes)
}

// GetEMailByID godoc
//
//		@Summary		Получить email пользователя, по его айди
//		@Description	Ищет пользователя в базе данных по айди, и возвращает его почту
//		@Produce		plain/text
//		@Param       id             query    string true "ID пользователя"
//		@Success		200
//		@Router			/admin/get-email-by-id [get]
//	 @Security BearerAuth
func GetEMailByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := pg.QueryRow(context.TODO(), "SELECT email FROM users WHERE id = $1", id)
	var email string
	if err := row.Scan(&email); err != nil {
		w.Write([]byte("User not founded"))
		return
	}
	w.Write([]byte(email))
}

// GetUserByEMail godoc
//
//		@Summary		Получить пользователя по почте
//		@Description	Ищет пользователя в базе данных по почте, и возвращает всю информацию о нем
//		@Produce		application/json
//		@Param       email             query    string true "EMail пользователя"
//		@Success		200		{object}	user
//		@Router			/admin/get-user-by-email [get]
//	 	@Security BearerAuth
func GetUserByEMail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	row := pg.QueryRow(context.TODO(), "SELECT id, username, email FROM users WHERE email = $1", email)
	var user user
	row.Scan(&user.ID, &user.Username, &user.EMail)
	bytes, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "json/application")
	w.Write(bytes)
}
