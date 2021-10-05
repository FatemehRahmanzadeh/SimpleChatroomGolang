package repository
import(
    "database/sql"
    "log"
    "github.com/FatemehRahmanzadeh/chat_sample/models"
)
// parsing json data to the user structure
type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) GetId() string {
	return user.Id
}

func (user *User) GetUserame() string {
	return user.Username
}

type UserRepository struct {
	Db *sql.DB
}


func (repo *UserRepository) FindUserById(ID string) models.User {

	row := repo.Db.QueryRow("SELECT id, username FROM user where id = ? LIMIT 1", ID)

	var user User

	if err := row.Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user

}

func (repo *UserRepository) GetAllUsers() []models.User {

	rows, err := repo.Db.Query("SELECT id, name FROM user")

	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username)
		users = append(users, &user)
	}

	return users
}

func (repo *UserRepository) FindUserByUsername(username string) *User {

	row := repo.Db.QueryRow("SELECT id, username, password FROM user where username = ? LIMIT 1", username)

	var user User

	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user

}