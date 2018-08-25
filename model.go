package imagesearch

type File struct {
	Filename  string `sql:"not null; index;" json:"filename"`
	Hash      string `sql:"not null; index;" json:"hash"`
	MediaType string `sql:"not null" json:"media_type"`
}

func NewFile(hash string, mediaType string, fileName string) *File {
	return &File{Filename: fileName, Hash: hash, MediaType: mediaType}
}