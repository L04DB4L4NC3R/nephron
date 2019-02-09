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

type FuzzyReturn struct {
	Rs  []map[string]interface{} `json:"rs"`
	Err error                    `json:"err"`
}

type Query struct {
	Query string `json:"query"`
}

type LogType struct {
	Query    string  `json:"query"`
	Hits     int     `json:"hits"`
	Time     float64 `json:"time"`
	MaxScore float64 `json:"maxScore"`
}
