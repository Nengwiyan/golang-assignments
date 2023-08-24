package helpers

import (
	"bytes"
	"context"
	"final-project/configs"
	"io"
	"mime/multipart"
	"path"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(configs.EnvCloudName(), configs.EnvCloudAPIKey(), configs.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	//convert file
	fileReader, err := convertFile(fileHeader)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   configs.EnvCloudUploadFolder(),
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//read file into memory buffer
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	//create byte.Reader from the buffer
	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

func RemoveExtention(fileName string) string {
	return path.Base(fileName[:len(fileName)-len(path.Ext(fileName))])
}
