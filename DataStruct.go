package main

type User struct {
	Id   int64  `json:"id,string"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Users []User

type RelationShip struct {
	Id      int64  `json:"-"`              //- ignore this column
	User_id int64  `json:"user_id,string"` //，string convert to string ,Only userd on string, floating-point and integer data
	To_uid  int64  `json:"-"`
	State   string `json:"state"`
	Type    string `json:"type"`
}
type Relations []RelationShip

type Error struct {
	ErrorCode int    `json:"errorCode,string"`
	Contet    string `json:"content"`
}

