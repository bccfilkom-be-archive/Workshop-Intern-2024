package supabase

import (
	"mime/multipart"
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type supabaseStorage struct {
	client *supabasestorageuploader.Client
}

func Init() Interface {
	supClient := supabasestorageuploader.New(
		os.Getenv("SUPABASE_LINK"),
		os.Getenv("SUPABASE_TOKEN"),
		os.Getenv("SUPABASE_BUCKET"),
	)

	return &supabaseStorage{
		client: supClient,
	}
}

func (s *supabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	link, err := s.client.Upload(file)
	if err != nil {
		return link, err
	}

	return link, nil
}

func (s *supabaseStorage) Delete(link string) error {
	err := s.client.Delete(link)
	if err != nil {
		return err
	}

	return nil
}
