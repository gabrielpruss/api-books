package models

/*
type JsonDateTime struct {
	Dataproc time.Time `json:"data leitura"`
}*/

type Ebook struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Lido     uint   `json:"lido"`
	Dataproc string `json:"data leitura"`
	Paginas  uint   `json:"paginas"`
}

type SetLido struct {
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Lido     uint   `json:"lido"`
	Dataproc string `json:"data leitura"`
	Paginas  uint   `json:"paginas"`
}

type Livros_fisicos struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Lido     uint   `json:"lido"`
	Dataproc string `json:"data leitura"`
	Paginas  uint   `json:"paginas"`
}

/*
// Implement Marshaler and Unmarshaler interface
func (j *JsonDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDateTime(t)
	return nil
}

func (j JsonDateTime) MarshalJSON() ([]byte, error) {
	//return json.Marshal(time.Time(j))
	return []byte("\"" + time.Time(j).Format("2006-01-02") + "\""), nil
}

// Maybe a Format function for printing your date
func (j JsonDateTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
*/
