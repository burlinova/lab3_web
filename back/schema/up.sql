-- Таблица категорий каталога (магазины, еда, дети, кинотеатр)
CREATE TABLE catalog_categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL, -- Например, "Магазины", "Еда", "Детям", "Кинотеатр"
    description TEXT,
    url TEXT NOT NULL -- Например, "shops_item.html", "food_item.html"
);

-- Таблица пользователей (для входа и регистрации)
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Таблица магазинов
CREATE TABLE shops (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    name TEXT NOT NULL, -- Например, "Zara", "Nike"
    short_description TEXT NOT NULL, -- Краткое описание для карточки
    detailed_description TEXT, -- Подробное описание для модального окна
    working_hours TEXT NOT NULL, -- Например, "10:00–22:00"
    floor TEXT, -- Например, "2-й этаж"
    image_url TEXT, -- Например, "zara_store.jpg"
    features TEXT, -- Особенности, например, "Еженедельное обновление коллекций"
    visitor_tips TEXT, -- Советы посетителям
    FOREIGN KEY (category_id) REFERENCES catalog_categories(id) ON DELETE CASCADE
);

-- Таблица заведений общепита
CREATE TABLE food (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    name TEXT NOT NULL, -- Например, "Вкусно — и точка", "Starbucks"
    short_description TEXT NOT NULL, -- Краткое описание для карточки
    detailed_description TEXT, -- Подробное описание для модального окна
    working_hours TEXT NOT NULL, -- Например, "10:00–22:00"
    floor TEXT, -- Например, "3-й этаж"
    image_url TEXT, -- Например, "vkusno_store.jpg"
    cuisine_type TEXT, -- Например, "Фастфуд", "Японская"
    features TEXT, -- Особенности, например, "Терминалы самообслуживания"
    visitor_tips TEXT, -- Советы посетителям
    FOREIGN KEY (category_id) REFERENCES catalog_categories(id) ON DELETE CASCADE
);

-- Таблица детских зон
CREATE TABLE kids_zones (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    name TEXT NOT NULL, -- Например, "Лабиринт", "Прыг-Скок"
    short_description TEXT NOT NULL, -- Краткое описание для карточки
    detailed_description TEXT, -- Подробное описание для модального окна
    working_hours TEXT NOT NULL, -- Например, "10:00–20:00"
    floor TEXT, -- Например, "3-й этаж"
    image_url TEXT, -- Например, "labyrinth_zone.jpg"
    age_range TEXT, -- Например, "3–10 лет"
    features TEXT, -- Особенности, например, "Безопасные материалы"
    visitor_tips TEXT, -- Советы посетителям
    FOREIGN KEY (category_id) REFERENCES catalog_categories(id) ON DELETE CASCADE
);

-- Таблица кинотеатров
CREATE TABLE cinemas (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    name TEXT NOT NULL, -- Например, "KAPO Max"
    short_description TEXT NOT NULL, -- Краткое описание для карточки
    detailed_description TEXT, -- Подробное описание для модального окна
    working_hours TEXT NOT NULL, -- Например, "10:00–23:00"
    floor TEXT, -- Например, "4-й этаж"
    image_url TEXT, -- Например, "kapo_max.jpg"
    features TEXT, -- Особенности, например, "IMAX, Dolby Atmos"
    visitor_tips TEXT, -- Советы посетителям
    FOREIGN KEY (category_id) REFERENCES catalog_categories(id) ON DELETE CASCADE
);

-- Таблица гостевой книги
CREATE TABLE guestbook (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    comment TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);


-- Индексы для оптимизации поиска
/*
CREATE INDEX idx_shops_name ON shops(name);
CREATE INDEX idx_food_name ON food(name);
CREATE INDEX idx_kids_zones_name ON kids_zones(name);
CREATE INDEX idx_cinemas_name ON cinemas(name);
*/