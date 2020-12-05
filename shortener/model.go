package shortener

type Redirect struct {
	Code     string `json:"code" bson:"code" msgpack:"code"`
	URL      string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	CriadoEm int64  `json:"criado_em" bson:"criado_em" msgpack:"criado_em"`
}
