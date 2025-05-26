let discountRate = 0.1; // Пример: скидка 10%

// let initialCartItems = [
//     { id: 1, name: "Товар 1", price: 1599 },
//     { id: 2, name: "Товар 2", price: 2599 },
//     { id: 3, name: "Товар 3", price: 999 },
// ];
// localStorage.setItem('cartItems', JSON.stringify(initialCartItems));

function addToCart(id, name, price) {
  // Загружаем текущие товары из LocalStorage
  let cartItems = JSON.parse(localStorage.getItem("cartItems")) || [];

  // // Проверяем, существует ли товар в корзине
  let existingItem = cartItems.find((item) => item.id === id);
  if (existingItem) {
    existingItem.count++;
  } else {
    let count = 1;
    // Добавляем новый товар в массив корзины
    const newItem = { id, name, price, count };
    cartItems.push(newItem);
  }

  // Сохраняем обновленный массив корзины в LocalStorage
  localStorage.setItem("cartItems", JSON.stringify(cartItems));

  // Обновляем количество товаров и общую сумму
  updateCartTotals();
  updateCartCount();
}

// Функция для обновления итогов корзины
function updateCartTotals() {
  let cartItems = JSON.parse(localStorage.getItem("cartItems")) || [];
  let totalCost = cartItems.reduce(
    (sum, item) => sum + item.price * item.count,
    0,
  );
  let discount = totalCost * discountRate;
  let finalCost = totalCost - discount;

  // Сохраняем итоговые данные в LocalStorage
  const cartTotals = { totalCost, discount, finalCost };
  localStorage.setItem("cartTotals", JSON.stringify(cartTotals));

  // Обновляем отображение на экране
  displayCartTotals();
}

function updateCartCount() {
  const cartItems = JSON.parse(localStorage.getItem("cartItems")) || [];
  // const cartCount = cartItems.length;
  const cartCount = cartItems.reduce((sum, item) => sum + item.count, 0);
  document.querySelector(".cart-count").textContent = cartCount;
}

// Функция для отображения итогов корзины из LocalStorage
function displayCartTotals() {
  const cartTotals = JSON.parse(localStorage.getItem("cartTotals")) || {
    totalCost: 0,
    discount: 0,
    finalCost: 0,
  };

  document.getElementById("totalCost").textContent =
    `${cartTotals.totalCost.toFixed(2)} ₽`;
  document.getElementById("discount").textContent =
    `${cartTotals.discount.toFixed(2)} ₽`;
  document.getElementById("finalCost").textContent =
    `${cartTotals.finalCost.toFixed(2)} ₽`;
}

// Функция отображения товаров в модальном окне
function displayCartItems() {
  const cartItemsContainer = document.getElementById("cartItems");
  cartItemsContainer.innerHTML = "";

  let cartItems = JSON.parse(localStorage.getItem("cartItems")) || [];

  cartItems.forEach((item) => {
    const itemDiv = document.createElement("div");
    itemDiv.classList.add("cart-item");
    itemDiv.innerHTML = `
            <span>${item.name}</span>
            <span>${item.price.toFixed(2)} ₽</span>
            <span> x${item.count.toFixed(0)} </span>
            <button onclick="removeItem(${item.id})">Удалить</button>
        `;
    cartItemsContainer.appendChild(itemDiv);
  });

  // Обновляем отображение итогов
  displayCartTotals();
  updateCartCount();
}

// Функция для удаления товара из корзины
function removeItem(itemId) {
  let cartItems = JSON.parse(localStorage.getItem("cartItems")) || [];
  let existingItem = cartItems.find((item) => item.id === itemId);
  if (existingItem.count == 1) {
    cartItems = cartItems.filter((item) => item.id !== itemId);
  } else {
    existingItem.count--;
  }

  localStorage.setItem("cartItems", JSON.stringify(cartItems));

  // Обновляем итоговые значения после удаления товара
  updateCartTotals();
  displayCartItems();
}

// Настройка событий для открытия и закрытия модального окна
const cartModal = document.getElementById("cartModal");
const cartButton = document.getElementById("cartButton");
const closeBtn = document.querySelector(".close");

// Открытие модального окна и отображение корзины
cartButton.onclick = () => {
  displayCartItems(); // Обновляем содержимое корзины
};

// Вызов функции обновления итогов при загрузке страницы
window.onload = () => {
  updateCartTotals(); // Вызывается для расчета и сохранения итогов при первой загрузке
  displayCartItems(); // Обновляет отображение корзины
  updateCartCount();
};
