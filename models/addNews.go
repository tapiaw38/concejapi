package models

type AddNews struct {
	Title    string `bson:"title" json:"title,omitempty"`
	Category string `bson:"category" json:"category,omitempty"`
	Body     string `bson:"body" json:"body,omitempty"`
	Picture  string `bson:"body" json:"picture,omitempty"`
}
