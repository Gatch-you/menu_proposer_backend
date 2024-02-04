package controllers

import (
	database "backend/src/database"
	"backend/src/middlewares"
	"backend/src/models"

	"github.com/gofiber/fiber/v2"
)

func Foods(c *fiber.Ctx) error {
	var foods []models.Food

	database.DB.Find(&foods)

	return c.JSON(foods)
}

func CreateFoods(c *fiber.Ctx) error {
	var food models.Food

	if err := c.BodyParser(&food); err != nil {
		return err
	}

	user_id, _ := middlewares.GetUserId(c)
	food.UserId = user_id

	database.DB.Create(&food)

	return c.JSON(food)
}

// import (
// 	db "backend/src/database"
// 	middleware "backend/src/middlewares"
// 	model "backend/src/models"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func FetchFoods(w http.ResponseWriter, r *http.Request) {
// 	db := db.Connect()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM foods")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	foodArgs := make([]model.Food, 0)
// 	for rows.Next() {
// 		var food model.Food
// 		err = rows.Scan(&food.ID, &food.Name, &food.Quantity, &food.Unit, &food.ExpirationDate, &food.Type)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		food.FormattedDate = food.ExpirationDate.Format("2006-01-02")

// 		foodArgs = append(foodArgs, food)
// 	}

// 	middleware.CorsMiddleware(http.DefaultServeMux)
// 	json.NewEncoder(w).Encode(foodArgs)
// }

// // curl http://localhost:8080/backend/foods/search_name%q=
// func SearchFoodsName(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query().Get("q")

// 	db := db.Connect()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, quantity, unit, expiration_date FROM foods WHERE name LIKE ?", "%"+query+"%")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	foodArgs := make([]model.Food, 0)
// 	for rows.Next() {
// 		var food model.Food
// 		err = rows.Scan(&food.Name, &food.Quantity, &food.Unit, &food.ExpirationDate)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		foodArgs = append(foodArgs, food)
// 	}

// 	middleware.CorsMiddleware(http.DefaultServeMux)
// 	json.NewEncoder(w).Encode(foodArgs)

// }

// // 新しい食品の項目追加
// // ↓実行コマンド
// // curl -X POST -d '{"name": "豚肉", "quantity": 250, "unit": "g", "expiration_date": "2023-06-10T00:00:00Z", "type": "精肉"}' http://localhost:8080/backend/insert_food
// func InsertFoods(w http.ResponseWriter, r *http.Request) {
// 	db := db.Connect()
// 	defer db.Close()

// 	var food model.Food
// 	err := json.NewDecoder(r.Body).Decode(&food)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	expirationDate, err := time.Parse("2006-01-02", food.ExpirationDate.Format("2006-01-02"))
// 	if err != nil {
// 		log.Fatal(food.ExpirationDate)
// 	}
// 	food.ExpirationDate = expirationDate

// 	inst, err := db.Prepare("INSERT INTO foods (id, name, quantity, unit, expiration_date, type) VALUES (?, ?, ?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	_, err = inst.Exec(food.ID, food.Name, food.Quantity, food.Unit, food.ExpirationDate, food.Type)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	middleware.CorsMiddleware(http.DefaultServeMux)
// 	json.NewEncoder(w).Encode(food)

// }

// // 食品の数量、個数の変化をこのコードにて処理する。0の量もこのデータにて扱う

// // curl -X PUT -H "Content-Type: application/json" -d '{"name": "キャベツ", "quantity": 0.3, "unit": " 個", "expiration_date": "2023-05-02T00:00:00Z", "type": "野菜"}' http://localhost:8080/backend/update_food
// func UpdateFoods(w http.ResponseWriter, r *http.Request) {
// 	db := db.Connect()
// 	defer db.Close()

// 	var food model.Food
// 	err := json.NewDecoder(r.Body).Decode(&food)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	update, err := db.Prepare("UPDATE foods SET id = ?, name = ?, quantity = ?, unit = ?, expiration_date = ?, type = ? WHERE id = ?")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	_, err = update.Exec(food.ID, food.Name, food.Quantity, food.Unit, food.ExpirationDate, food.Type, food.ID)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	middleware.CorsMiddleware(http.DefaultServeMux)
// 	json.NewEncoder(w).Encode(food)
// }

// // 食品のデータベースのフィールドそのものを削除する。再びその食品を使うには再度InsertFoodsを叩かなければならなくなる。
// // curl -X DELETE -H "Content-Type: application.json" -d '{"food_id": 1}' http://localhost:8080/backend/delete_food
// func DeleteFoods(w http.ResponseWriter, r *http.Request) {
// 	db := db.Connect()
// 	defer db.Close()

// 	var food model.Food
// 	err := json.NewDecoder(r.Body).Decode(&food)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	delete, err := db.Prepare("DELETE FROM foods WHERE id = ?")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	_, err = delete.Exec(food.ID)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	middleware.CorsMiddleware(http.DefaultServeMux)
// 	json.NewEncoder(w).Encode(food)
// }
