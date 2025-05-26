let currentIndex = 0; // Текущий индекс слайда
const slides = document.querySelectorAll(".slide");
const dotsContainer = document.getElementById("dots-container");

// Создаём индикаторы
slides.forEach((_, index) => {
  const dot = document.createElement("div");
  dot.classList.add("dot");
  if (index === 0) dot.classList.add("active");
  dot.addEventListener("click", () => goToSlide(index));
  dotsContainer.appendChild(dot);
});

const dots = document.querySelectorAll(".dot");

// Функция перемещения слайдов
function moveSlide(step) {
  currentIndex += step;
  if (currentIndex < 0) currentIndex = slides.length - 1; // Зацикливание назад
  if (currentIndex >= slides.length) currentIndex = 0; // Зацикливание вперёд
  updateSlider();
}

// Переход к конкретному слайду
function goToSlide(index) {
  currentIndex = index;
  updateSlider();
}

// Обновление слайдера
function updateSlider() {
  const slideWidth = slides[0].clientWidth;
  const slidesContainer = document.getElementById("slides");
  slidesContainer.style.transform = `translateX(-${currentIndex * slideWidth}px)`;

  // Обновление активного индикатора
  dots.forEach((dot, index) => {
    dot.classList.toggle("active", index === currentIndex);
  });
}
