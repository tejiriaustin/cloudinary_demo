package database

type repo struct{}

//TODO: FLESH OUT THE DATABASE TO SUPPORT CLOUDINARY

type Storage interface {
	Save(int642 int64) error
	FindAll() error
	Find() error
}

func NewCloudStore() Storage {
	return &repo{}
}
