package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type User struct {
	ID   string
	Name string
	Age  int
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/").To(u.findAllUsers))
	return ws
}

func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	list := []User{}
	list = append(list, User{"1", "yang", 10})
	response.WriteEntity(list)
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	usr := User{"1", "yang", 10}
	if len(usr.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

func main() {
	u := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(u.WebService())
	log.Printf("Get the API using http://localhost:8080/users")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
