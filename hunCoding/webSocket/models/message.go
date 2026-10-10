package models


/*1º criaremos  um tipo**/
type Message struct {
	UserName string `json:"userName"`
	Message string `json:"message"`

}