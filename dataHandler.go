package main

import (
	"github.com/go-pg/pg"
	"log"
)

var db *pg.DB

func isDbAlive() {
	if db == nil {
		Db_connect()
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&RelationShip{}} {
		//todo 判断表是否已存在
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func Db_connect() {
	db = pg.Connect(&pg.Options{
		// Host:     "127.0.0.1",
		// Port:     "5432",
		User:     "postgres",
		Password: "123456",
		Database: "postgres",
	})
}

func InsertUser(user *User) (bool, error) {
	isDbAlive()
	err := db.Insert(user)
	if err != nil {
		//TODO recover this panic
		log.Println(err)
		return false, err
	}
	return true, nil
}

func InsertRelationShip(relationShip *RelationShip) (bool, error) {
	isDbAlive()

	getship := &RelationShip{}
	getship.User_id = relationShip.To_uid
	getship.To_uid = relationShip.User_id

	if "liked" == relationShip.State {
		searchShip(getship)
		//如果已经有喜欢，like修改为match。
		if "liked" == getship.State {
			relationShip.State = "matched"
			getship.State = "matched"
			//updateresult := UpdateRelationShip(getship)
			// if !updateresult {
			// 	fmt.Println("UpdateRelationaShip error")
			// 	panic("UpdateRelationShip error")
			// }
			tx, err := db.Begin()
			if err != nil {
				log.Println(err)
				return false, err
			}
			// rollback on error
			defer tx.Rollback()
			_, err = tx.Exec(`UPDATE relation_ships SET state = ? where user_id=? and to_uid=?`,
				getship.State, getship.User_id, getship.To_uid)

			if err != nil {
				log.Println(err)
				return false, err
			}
			_, err = tx.Exec(`INSERT into relation_ships(user_id,to_uid,state,type) values(?,?,?,?)`,
				relationShip.User_id, relationShip.To_uid, relationShip.State, relationShip.Type)
			if err != nil {
				log.Println(err)
				return false, err
			}
			return true, tx.Commit()
		}
	}
	err := db.Insert(relationShip)
	if err != nil {
		//TODO recover this panic
		//panic(err)
		log.Println(err)
		return false, err
	}
	return true, nil
}

func searchShip(relationShip *RelationShip) {
	isDbAlive()
	db.Model(relationShip).
		Where(" user_id=? ", relationShip.User_id).
		Where(" to_uid=?", relationShip.To_uid).
		Select()

	//查询是否存在，这里err是否可以不处理，为何取不到也要报错
	// if err != nil {
	// 	panic(err)
	// }
}

func getAllRelationShip(relationShips *[]RelationShip, user int) {
	isDbAlive()
	db.Model(relationShips).Column("user_id", "state", "type").Where("to_uid=?", user).Select()
}

func UpdateRelationShip(relationShip *RelationShip) bool {
	isDbAlive()
	err := db.Update(relationShip)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetAllUser(users *[]User) {
	isDbAlive()
	err := db.Model(users).Select()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Db_connect_test() {

	isDbAlive()
	// err := createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }
	// user1 := &User{
	// 	Name: "admin3",
	// 	Type: "user",
	// }
	// InsertUser(user1)

	// ship := &RelationShip{
	// 	User_id: 22000,
	// 	To_uid:  23000,
	// 	State:   "liked",
	// 	Type:    "relationship",
	// }

	// fmt.Println(getship.User_id)
	// fmt.Println(getship.To_uid)
	// fmt.Println("开始的State" + getship.State)

	// result := InsertRelationShip(ship)
	// fmt.Println(result)

	// var relationShips []RelationShip
	// getAllRelationShip(&relationShips, 22)
	// fmt.Println(len(relationShips))
	//selects the model by primary key.
	// err := db.Select(&user)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(ship.Id)

}

// func main() {
// 	Db_connect_test()
// }

