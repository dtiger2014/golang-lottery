package models

// ObjLoginuser : 站点中与浏览器交互的用户模型
type ObjLoginuser struct {
	UID int
	Username string
	Now int
	IP string
	Sign string 
}