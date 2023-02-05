package entity

type User struct {
	Id   int    `json:"id""`
	Name string `json:"name""`
	Sex  int    `json:"sex""`
	Url  string `json:"url"`
}

//func (u User) GetId() int {
//	return u.id
//}
//func (u User) SetId(id int) {
//	u.id = id
//}
//
//func (u User) GetName() string {
//	return u.name
//}
//func (u User) SetName(name string) {
//	u.name = name
//}
//
//func (u User) GetSex() int {
//	return u.sex
//}
//func (u User) SetSex(sex int) {
//	u.sex = sex
//}
//
//func (u User) GetUrl() string {
//	return u.url
//}
//func (u User) SetUrl(url string) {
//	u.url = url
//}
