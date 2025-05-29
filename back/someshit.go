package main

import (
	"context"
	"fmt"
)

// PGFindUsers ищет пользователей по email
func PGFindUsers(email string) ([]User, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM users WHERE (email ILIKE $1) OR ($1 = '')`,
		"%"+email+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	var user User
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// PGFindCatalogCategories ищет категории каталога по имени
func PGFindCatalogCategories(name string) ([]CatalogCategory, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM catalog_categories WHERE (name ILIKE $1) OR ($1 = '')`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []CatalogCategory
	var category CatalogCategory
	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.URL); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// PGFindShops ищет магазины по имени
func PGFindShops(name string) ([]Shop, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM shops WHERE (name ILIKE $1) OR ($1 = '')`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []Shop
	var shop Shop
	for rows.Next() {
		if err = rows.Scan(&shop.ID, &shop.CategoryID, &shop.Name, &shop.ShortDescription,
			&shop.DetailedDescription, &shop.WorkingHours, &shop.Floor, &shop.ImageURL,
			&shop.Features, &shop.VisitorTips, &shop.Characteristic); err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

// PGFindFoodItems ищет заведения общепита по имени
func PGFindFoodItems(name string) ([]FoodItem, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM food WHERE (name ILIKE $1) OR ($1 = '')`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foodItems []FoodItem
	var foodItem FoodItem
	for rows.Next() {
		if err = rows.Scan(&foodItem.ID, &foodItem.CategoryID, &foodItem.Name, &foodItem.ShortDescription,
			&foodItem.DetailedDescription, &foodItem.WorkingHours, &foodItem.Floor, &foodItem.ImageURL,
			&foodItem.CuisineType, &foodItem.Features, &foodItem.VisitorTips, &foodItem.Characteristic); err != nil {
			return nil, err
		}
		foodItems = append(foodItems, foodItem)
	}
	return foodItems, nil
}

// PGFindKidsZones ищет детские зоны по имени
func PGFindKidsZones(name string) ([]KidsZone, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM kids_zones WHERE (name ILIKE $1) OR ($1 = '')`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kidsZones []KidsZone
	var kidsZone KidsZone
	for rows.Next() {
		if err = rows.Scan(&kidsZone.ID, &kidsZone.CategoryID, &kidsZone.Name, &kidsZone.ShortDescription, &kidsZone.DetailedDescription, &kidsZone.WorkingHours, &kidsZone.Floor, &kidsZone.ImageURL, &kidsZone.AgeRange, &kidsZone.Features, &kidsZone.VisitorTips, &kidsZone.Characteristic); err != nil {
			return nil, err
		}
		fmt.Printf("kids_zone: %+v\n", kidsZone)
		kidsZones = append(kidsZones, kidsZone)
	}
	return kidsZones, nil
}

// PGFindCinemas ищет кинотеатры по имени
func PGFindCinemas(name string) ([]Cinema, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM cinemas WHERE (name ILIKE $1) OR ($1 = '');`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []Cinema
	var cinema Cinema
	for rows.Next() {
		if err = rows.Scan(&cinema.ID, &cinema.CategoryID, &cinema.Name, &cinema.ShortDescription,
			&cinema.DetailedDescription, &cinema.WorkingHours, &cinema.Floor, &cinema.ImageURL,
			&cinema.Features, &cinema.VisitorTips, &cinema.Characteristic); err != nil {
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}
	return cinemas, nil
}

// PGFindCinemas ищет фильм по имени
func PGFindMovies(title string) ([]Movies, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM movies WHERE (name ILIKE $1) OR ($1 = '');`,
		"%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movies
	var movie Movies
	for rows.Next() {
		if err = rows.Scan(&movie.ID, &movie.CinemaID, &movie.Title, &movie.Genre,
			&movie.Year, &movie.Country); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// PGFindGuestbook ищет записи гостевой книги по комментарию
func PGFindGuestbook(comment string) ([]Guestbook, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM guestbook WHERE (comment ILIKE $1) OR ($1 = '')`,
		"%"+comment+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Guestbook
	var entry Guestbook
	for rows.Next() {
		if err = rows.Scan(&entry.ID, &entry.UserID, &entry.Comment, &entry.CreatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

// PGFindContacts ищет контакты по адресу или email
func PGFindContacts(search string) ([]Contact, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM contacts WHERE (address ILIKE $1 OR email ILIKE $1) OR ($1 = '')`,
		"%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	var contact Contact
	for rows.Next() {
		if err = rows.Scan(&contact.ID, &contact.Address, &contact.Phone, &contact.Email,
			&contact.WorkingHours, &contact.MapURL); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

// PGFindAbouts ищет записи "О нас" по заголовку
func PGFindAbouts(title string) ([]About, error) {
	rows, err := pg.Query(
		context.TODO(),
		`SELECT * FROM about WHERE (title ILIKE $1) OR ($1 = '')`,
		"%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var abouts []About
	var about About
	for rows.Next() {
		if err = rows.Scan(&about.ID, &about.Title, &about.Content, &about.ImageURL); err != nil {
			return nil, err
		}
		abouts = append(abouts, about)
	}
	return abouts, nil
}
