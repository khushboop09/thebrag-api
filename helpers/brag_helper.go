package helpers

import (
	b64 "encoding/base64"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"thebrag/models"
	"thebrag/requests"
	"thebrag/s3"
	"time"

	"github.com/sendgrid/sendgrid-go"
)

func FormatDataForCSV(brags []models.Brag) [][]string {
	var records [][]string
	records = append(records, []string{"ID", "Title", "Details", "Category Name", "Created On"})
	for _, brag := range brags {
		item := []string{strconv.FormatUint(uint64(brag.ID), 10), brag.Title, brag.Details, brag.Category.Name, brag.CreatedAt.Format(time.RFC822)}
		records = append(records, item)
	}
	return records
}

func WriteToCSVFileAndEmail(records [][]string, request requests.ExportBragRequest, user models.User) bool {
	filePath := fmt.Sprintf("brags_%d_%s_%s.csv", user.ID, request.From, request.To)
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(records)
	//TODO:upload to cloud object storage
	UploadFileToObjectStorage(filePath)

	b, err := os.ReadFile(filePath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(b)
	status := SendBragsOnEmail(request, user, filePath, fileContent)
	os.Remove(filePath)
	return status
}

func SendBragsOnEmail(request requests.ExportBragRequest, user models.User, fileName string, fileContent string) bool {
	csvContentb64 := b64.StdEncoding.EncodeToString([]byte(fileContent))
	subject := fmt.Sprintf("Your brags from %s to %s", request.From, request.To)
	body := fmt.Sprintf("Your brags from %s to %s", request.From, request.To)
	sendgridRequest := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	sendgridRequest.Method = "POST"
	apiRequest := fmt.Sprintf(`{"personalizations": [{"to": [{"email": "%s"}]}],"subject": "%s","from": {"email": "%s"},"content": [{"type": "text/csv","value": "%s"}], "attachments": [{"content": "%s", "disposition": "attachment", "filename": "%s", "type": "text/csv"}]}`, user.Email, subject, os.Getenv("SENDER_EMAIL"), body, csvContentb64, fileName)
	sendgridRequest.Body = []byte(apiRequest)
	response, err := sendgrid.API(sendgridRequest)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return response.StatusCode == 202
	}
}

func UploadFileToObjectStorage(filePath string) {
	awsConfig := requests.AWSConfig{
		AccessKeyID:     os.Getenv("LINODE_ACCESS_KEY"),
		AccessKeySecret: os.Getenv("LINODE_SECRET_KEY"),
		Region:          os.Getenv("LINODE_BUCKET_REGION"),
		BucketName:      os.Getenv("BUCKET_NAME"),
		BaseURL:         os.Getenv("LINODE_BASE_URL"),
	}
	error := s3.UploadFile(filePath, awsConfig)
	if error != nil {
		fmt.Println(error.Error())
	}
}
