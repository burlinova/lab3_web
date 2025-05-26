package main

import (
	"database/sql"
	"time"
)

// User представляет таблицу users для входа и регистрации
type User struct {
	ID           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// CatalogCategory представляет таблицу catalog_categories для категорий каталога
type CatalogCategory struct {
	ID          int64          `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description sql.NullString `json:"description" db:"description"`
	URL         string         `json:"url" db:"url"`
}

// Shop представляет таблицу shops для магазинов
type Shop struct {
	ID                  int64          `json:"id" db:"id"`
	CategoryID          int64          `json:"category_id" db:"category_id"`
	Name                string         `json:"name" db:"name"`
	ShortDescription    string         `json:"short_description" db:"short_description"`
	DetailedDescription sql.NullString `json:"detailed_description" db:"detailed_description"`
	WorkingHours        string         `json:"working_hours" db:"working_hours"`
	Floor               sql.NullString `json:"floor" db:"floor"`
	ImageURL            sql.NullString `json:"image_url" db:"image_url"`
	Features            sql.NullString `json:"features" db:"features"`
	VisitorTips         sql.NullString `json:"visitor_tips" db:"visitor_tips"`
}

// FoodItem представляет таблицу food_items для заведений общепита
type FoodItem struct {
	ID                  int64          `json:"id" db:"id"`
	CategoryID          int64          `json:"category_id" db:"category_id"`
	Name                string         `json:"name" db:"name"`
	ShortDescription    string         `json:"short_description" db:"short_description"`
	DetailedDescription sql.NullString `json:"detailed_description" db:"detailed_description"`
	WorkingHours        string         `json:"working_hours" db:"working_hours"`
	Floor               sql.NullString `json:"floor" db:"floor"`
	ImageURL            sql.NullString `json:"image_url" db:"image_url"`
	CuisineType         sql.NullString `json:"cuisine_type" db:"cuisine_type"`
	Features            sql.NullString `json:"features" db:"features"`
	VisitorTips         sql.NullString `json:"visitor_tips" db:"visitor_tips"`
}

// KidsZone представляет таблицу kids_zones для детских зон
type KidsZone struct {
	ID                  int64          `json:"id" db:"id"`
	CategoryID          int64          `json:"category_id" db:"category_id"`
	Name                string         `json:"name" db:"name"`
	ShortDescription    string         `json:"short_description" db:"short_description"`
	DetailedDescription sql.NullString `json:"detailed_description" db:"detailed_description"`
	WorkingHours        string         `json:"working_hours" db:"working_hours"`
	Floor               sql.NullString `json:"floor" db:"floor"`
	ImageURL            sql.NullString `json:"image_url" db:"image_url"`
	AgeRange            sql.NullString `json:"age_range" db:"age_range"`
	Features            sql.NullString `json:"features" db:"features"`
	VisitorTips         sql.NullString `json:"visitor_tips" db:"visitor_tips"`
}

// Cinema представляет таблицу cinemas для кинотеатров
type Cinema struct {
	ID                  int64          `json:"id" db:"id"`
	CategoryID          int64          `json:"category_id" db:"category_id"`
	Name                string         `json:"name" db:"name"`
	ShortDescription    string         `json:"short_description" db:"short_description"`
	DetailedDescription sql.NullString `json:"detailed_description" db:"detailed_description"`
	WorkingHours        string         `json:"working_hours" db:"working_hours"`
	Floor               sql.NullString `json:"floor" db:"floor"`
	ImageURL            sql.NullString `json:"image_url" db:"image_url"`
	Features            sql.NullString `json:"features" db:"features"`
	VisitorTips         sql.NullString `json:"visitor_tips" db:"visitor_tips"`
}

// Guestbook представляет таблицу guestbook для отзывов
type Guestbook struct {
	ID        int64         `json:"id" db:"id"`
	UserID    sql.NullInt64 `json:"user_id" db:"user_id"`
	Comment   string        `json:"comment" db:"comment"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
}

// Contact представляет таблицу contacts для контактной информации
type Contact struct {
	ID           int64          `json:"id" db:"id"`
	Address      string         `json:"address" db:"address"`
	Phone        sql.NullString `json:"phone" db:"phone"`
	Email        sql.NullString `json:"email" db:"email"`
	WorkingHours sql.NullString `json:"working_hours" db:"working_hours"`
	MapURL       sql.NullString `json:"map_url" db:"map_url"`
}

// About представляет таблицу about для информации о ТРЦ
type About struct {
	ID       int64          `json:"id" db:"id"`
	Title    string         `json:"title" db:"title"`
	Content  string         `json:"content" db:"content"`
	ImageURL sql.NullString `json:"image_url" db:"image_url"`
}
