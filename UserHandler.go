package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	//info := fmt.Sprintf("%s %s %s %s %s", r.URL, r.Host, r.Method, r.RequestURI, r.URL.RawQuery)
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	//fmt.Fprintf(w, "User: %v\n", vars["user"])
	//TODO get userInfo from database
	var users []User
	GetAllUser(&users)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func SetUserHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	//get method by r.Method
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//TODO while insert duplicate name cause panic
	user.Type = "user"
	result, insert_err := InsertUser(&user)
	if !result {
		var parametererror Error
		parametererror.ErrorCode = 1002
		parametererror.Contet = fmt.Sprintln("error:\t", insert_err)
		if err := json.NewEncoder(w).Encode(parametererror); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	//return json result

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}

func GetRelationShipHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//fmt.Fprintf(w, "User: %v\n", vars["to_uid"])
	//TODO get userInfo from database
	var relationShips []RelationShip
	to_uid, _ := strconv.Atoi(vars["to_uid"])
	//strconv.ParseInt(vars["to_uid"], 10, 64)
	getAllRelationShip(&relationShips, to_uid)
	//fmt.Println(len(relationShips))
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	// if relationShips== nil {
	// 	relationShips = {}
	// }
	if err := json.NewEncoder(w).Encode(relationShips); err != nil {
		panic(err)
	}

}

func SetRelationShipHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var relationShip RelationShip
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	//get method by r.Method
	//fmt.Fprintf(w, "method: %v\n", r.Method)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	if err := json.Unmarshal(body, &relationShip); err != nil {
		//fmt.Fprintf(w, "3")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//fmt.Println(relationShip.State)

	w.WriteHeader(http.StatusCreated)
	if "liked" != relationShip.State && "disliked" != relationShip.State {
		var parametererror Error
		parametererror.ErrorCode = 1001
		parametererror.Contet = "parameter is illegal" + relationShip.State
		if err := json.NewEncoder(w).Encode(parametererror); err != nil {
			panic(err)
		}

	} else {
		to_uid, _ := strconv.ParseInt(vars["to_uid"], 10, 64)
		user_id, _ := strconv.ParseInt(vars["user_id"], 10, 64)
		relationShip.To_uid = to_uid
		relationShip.User_id = user_id
		relationShip.Type = "relationship"
		result, relation_err := InsertRelationShip(&relationShip)
		//insert error
		if relation_err != nil || !result {
			var parametererror Error
			parametererror.ErrorCode = 1002
			parametererror.Contet = fmt.Sprintln("error:\t", relation_err)
			if err := json.NewEncoder(w).Encode(parametererror); err != nil {
				panic(err)
			}
			return
		}
		//return json
		if err := json.NewEncoder(w).Encode(relationShip); err != nil {
			panic(err)
		}
	}

}

