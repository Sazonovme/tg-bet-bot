package handler

type Handlers struct {
	Services Services
}

type Services interface {
}

func NewHandler(services Services) *Handlers {
	return &Handlers{
		Services: services,
	}
}

func (h *Handlers) Start() {
	// b.SendText("без полезное хуйня")

	// 1 привет как дела
	// 2 определить админ ли юзер
	// 3 меняем ui кнопок на доступный фун-ции взависимости от админа или не админа

	// // b.SendText("регистарция пользователя")
	// // 1 загрузка о пользователе в бд через сервис
	// // 2 поменять
}

// admin
func (h *Handlers) CreateEvent() {
	// 1 получение списка в формате [{name: navi-faze, date: 21.01.2000T16:30}]
	// 2 парсинг списка
	// 3 добавление в бд списка событий  ( формат события id name date result )
	// 4 отдача ответа
}

// adimn
func (h *Handlers) AddResult() {
	// 1 получение результат по id { id: 123123, result: '0:1' }
	// 2 парсинг
	// 3 добавление в бд 
}

// admin
func (h *Handlers) FinishTournament() {

}

// user
func (h *Handlers) MyPredictions() {

}

// user
func (h *Handlers) MakePrediction() {

}
