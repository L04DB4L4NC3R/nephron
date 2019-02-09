package model

type StringReturn struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

type ByteReturn struct {
	Rs  []byte
	Err error
}

type ESdata struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
