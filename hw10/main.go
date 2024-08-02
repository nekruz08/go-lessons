// Задания для 10 – го урока

// Оценки студента
// Описание: Реализуйте структуру Student с полем grades (срез оценок). Реализуйте метод AddGrade, добавляющий оценку, и метод AverageGrade, возвращающий среднее значение оценок.
// Методы:
// AddGrade(grade int)
// AverageGrade() float64
// Круг и Площадь
// Описание: Реализуйте структуру Circle с полем radius. Реализуйте методы Area и Circumference для вычисления площади и периметра круга.
// Методы:
// Area() float64
// Circumference() float64
// Контейнер для товаров
// Описание: Реализуйте структуру Item с полями name и price. Реализуйте структуру Inventory с срезом товаров и методы для добавления товара и получения общего значения товаров в инвентаре.
// Методы:
// AddItem(item Item)
// TotalValue() float64
// Обработка транзакций
// Описание: Реализуйте структуру Transaction с полями amount и description. Реализуйте метод Process, который выводит информацию о транзакции.
// Методы:
// Process()
// AddTransaction(amount float64, description string)
// Управление задачами
// Описание: Реализуйте структуру Task с полями title и completed. Реализуйте структуру TaskList с срезом задач и методы для добавления задачи и получения количества завершённых задач.
// Методы:
// AddTask(title string)
// CompletedTasks() int
// Работа с температурой
// Описание: Реализуйте структуру Temperature с полем celsius. Реализуйте методы для преобразования температуры в Фаренгейт и Кельвин.
// Методы:
// ToFahrenheit() float64
// ToKelvin() float64
// Управление списком студентов
// Описание: Реализуйте структуру Student с полем name и age. Реализуйте структуру Classroom с срезом студентов и методы для добавления студента и получения средней возрастной группы.
// Методы:
// AddStudent(name string, age int)
// AverageAge() float64
// Склад товаров
// Описание: Реализуйте структуру Product с полями name и quantity. Реализуйте структуру Warehouse с срезом продуктов и методы для добавления продукта и получения общего количества товаров на складе.
// Методы:
// AddProduct(name string, quantity int)
// TotalQuantity() int
// Заметки и метки
// Описание: Реализуйте структуру Note с полем content и tags. Реализуйте структуру Notebook с срезом заметок и методы для добавления заметки и получения всех заметок с указанной меткой.
// Методы:
// AddNote(content string, tags []string)
// NotesWithTag(tag string) []Note
// Управление зарплатами
// Описание: Реализуйте структуру Employee с полями name и salary. Реализуйте структуру Payroll с срезом сотрудников и методы для добавления сотрудника и получения средней зарплаты.
// Методы:
// AddEmployee(name string, salary float64)
// AverageSalary() float64
// Разделение бюджета
// Описание: Реализуйте структуру Budget с полем amount. Реализуйте метод Allocate, который выделяет сумму из бюджета, если сумма доступна. Реализуйте метод Remaining для получения оставшегося бюджета.
// Методы:
// Allocate(amount float64) bool
// Remaining() float64
// Отслеживание заказов
// Описание: Реализуйте структуру Order с полями id и totalAmount. Реализуйте структуру OrderManager с срезом заказов и методы для добавления заказа и получения общей суммы всех заказов.
// Методы:
// AddOrder(id int, totalAmount float64)
// TotalOrdersAmount() float64
// Управление аккаунтами
// Описание: Реализуйте структуру Account с полем balance. Реализуйте методы Deposit и Withdraw, а также метод GetBalance для получения текущего баланса.
// Методы:
// Deposit(amount float64)
// Withdraw(amount float64) bool
// GetBalance() float64
// Операции над строками
// Описание: Реализуйте структуру Text с полем content. Реализуйте методы для добавления текста, замены слова и получения длины текста.
// Методы:
// AddText(text string)
// ReplaceWord(oldWord, newWord string)
// Length() int
// Управление списком покупок
// Описание: Реализуйте структуру ShoppingItem с полями name и price. Реализуйте структуру ShoppingList с срезом покупок и методы для добавления покупки и получения общего значения покупок.
// Методы:
// AddItem(name string, price float64)
// TotalPrice() float64
// Учет времени
// Описание: Реализуйте структуру Event с полями name и date (в формате строк). Реализуйте методы для добавления события и получения всех событий после указанной даты.
// Методы:
// AddEvent(name string, date string)
// EventsAfterDate(date string) []Event
// Управление заказами в магазине
// Описание: Реализуйте структуру Order с полями orderID, product и quantity. Реализуйте структуру Store с срезом заказов и методы для добавления заказа и получения общего количества товаров по каждому продукту.
// Методы:
// AddOrder(orderID int, product string, quantity int)
// TotalQuantityByProduct(product string) int
// Расчет стоимости поездки
// Описание: Реализуйте структуру Trip с полями distance и costPerMile. Реализуйте методы для установки стоимости за милю и расчета общей стоимости поездки.
// Методы:
// SetCostPerMile(cost float64)
// TotalCost() float64
// Рейтинг фильмов
// Описание: Реализуйте структуру Movie с полями title и rating. Реализуйте структуру MovieList с срезом фильмов и методы для добавления фильма и получения среднего рейтинга.
// Методы:
// AddMovie(title string, rating float64)
// AverageRating() float64
// Система учёта заказов с учетом скидок
// Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру Discount с полем percentage. Реализуйте методы для применения скидки и получения общей суммы заказа с учетом скидки.
// Методы:
// ApplyDiscount(discount Discount)
// TotalAmount() float64
// Многократные переводы денег
// Описание: Реализуйте структуру Account с полями balance и history (срез транзакций). Реализуйте методы для пополнения, снятия и получения истории транзакций.
// Методы:
// Deposit(amount float64)
// Withdraw(amount float64) bool
// TransactionHistory() []string
// Статистика по продажам
// Описание: Реализуйте структуру Sale с полями item и amount. Реализуйте структуру SalesReport с методами для добавления продажи и получения статистики по общим продажам.
// Методы:
// AddSale(item string, amount float64)
// TotalSales() float64
// Учет кредитов и дебетов
// Описание: Реализуйте структуру Transaction с полями transactionType (кредит или дебет) и amount. Реализуйте структуру Account с методами для добавления транзакций и получения баланса.
// Методы:
// AddTransaction(transactionType string, amount float64)
// Balance() float64
// Учет заказов с учетом возвратов
// Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру OrderManager с методами для добавления заказа, обработки возвратов и получения общего значения активных заказов.
// Методы:
// AddOrder(id int, product string, quantity int, price float64)
// ReturnOrder(id int, quantity int)
// TotalActiveOrders() float64
// Управление проектами
// Описание: Реализуйте структуру Task с полями title, description и status. Реализуйте структуру Project с срезом задач и методы для добавления задачи и получения количества задач в каждом статусе.
// Методы:
// AddTask(title, description, status string)
// CountTasksByStatus(status string) int
// Управление пользовательскими данными
// Описание: Реализуйте структуру User с полями username, email и age. Реализуйте структуру UserManager с срезом пользователей и методы для добавления пользователя и получения всех пользователей старше определенного возраста.
// Методы:
// AddUser(username, email string, age int)
// UsersOlderThan(age int) []User
// Управление контрактами
// Описание: Реализуйте структуру Contract с полями contractID, client и amount. Реализуйте структуру ContractManager с срезом контрактов и методы для добавления контракта и получения общего объема контрактов для клиента.
// Методы:
// AddContract(contractID int, client string, amount float64)
// TotalAmountForClient(client string) float64
// Управление списком книг
// Описание: Реализуйте структуру Book с полями title, author и year. Реализуйте структуру Library с срезом книг и методы для добавления книги и получения всех книг авторов, опубликованных после определенного года.
// Методы:
// AddBook(title, author string, year int)
// BooksByAuthorsAfterYear(year int) []Book
// Отслеживание активностей пользователя
// Описание: Реализуйте структуру Activity с полями activityType и timestamp. Реализуйте структуру UserActivityTracker с срезом активностей и методы для добавления активности и получения всех активностей после определенного времени.
// Методы:
// AddActivity(activityType string, timestamp time.Time)
// ActivitiesAfterTime(timestamp time.Time) []Activity
// -------------------------------------------------------------------------------------------------
// Книга и Автор
// Описание: Реализуйте структуру Book с полями title и author,
// а также метод GetDetails, который возвращает строку с деталями книги.

// Реализуйте структуру Library с массивом книг
// и метод AddBook, добавляющий книгу в библиотеку.
// Методы:
// GetDetails() string для структуры Book
// AddBook(book Book) для структуры Library
package main

import "fmt"

type Book struct {
	title  string
	author string
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("title:  %s\nauthor: %s", b.title, b.author)
}

type Library struct {
	book []Book
}

func (l *Library) AddBook(book Book) {
	l.book = append(l.book, book)
}

func main() {
	book1 := Book{
		title:  "Jam Jou",
		author: "Din Dan",
	}
	result := book1.GetDetails()
	fmt.Println(result)

	fmt.Println()

	l:=Library{
		book: []Book{},
	}
	fmt.Printf("before: %v\n",l.book) //before
	l.AddBook(book1)
	fmt.Printf("after: %v\n",l.book) //after
}
