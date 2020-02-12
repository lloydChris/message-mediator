package main

import "fmt"

//Type for pipeline enum
type Pipeline string

//Enum for the different pipelines available
const (
	PipeA = "pipea"
	PipeB = "pipeb"
)

//Attachments as found in the Email object
type Attachments struct {
	Name        string
	Content     string //Assume this a base64 encoded string of data
	ContentType string
	ContentID   string
}

//Email object to parse queue message into
type Email struct {
	From        string
	To          string
	CC          string
	BCC         string
	Subject     string
	Tag         string
	HTMLBody    string
	TextBody    string
	ReplyTo     string
	Headers     []map[string]string //Make this an array of Key Value pairs
	TrackOpens  bool
	TrackLinks  string
	Metadata    []map[string]string //Make this a defined keyValue pair
	Attachments []Attachments
	messageID   string //GUID that can be used to prevent duplicate sendings
}

func main() {

	//Subscribe to queue
	//TODO: find queue reading library based on what queuing service is chosen
	//TODO: parse queue message into email object

	//mocked email for demonstration
	inboundEmail := Email{
		From: "some@wildbit.com",
		To:   "other@wildbit.com",
	}
	processMessage(inboundEmail)
}

func processMessage(email Email) {

	//Call Qualifying status logic
	isQualified := IsQualified(email)

	//If Message is to be rejected, log it and end execution
	if !isQualified {
		fmt.Println("message is unqualified, terminating")
		log(email, "unqualified")
		return
	}

	fmt.Println("message is qualified")

	//Call whatever other business logic is needed
	fmt.Println("calling some other business logic.  maybe an external api or just another module")

	//Call Pipeline business logic
	pipeline := GetPipeline(email)
	fmt.Println("Determined pipeline: ", pipeline)

	//Chose pipeline based on pipeline busnesslogic response
	//TODO: implement Pipeline handoff
	switch pipeline {
	case PipeA:
		fmt.Println("Sending email to pipeline A")
		//Do something PipeA Specific
	case PipeB:
		fmt.Println("Sending email to pipelline B")
		//Do something PipeB Specifc
	}
}

func log(email Email, message string) {
	//plug in actual logging method
	fmt.Print(message, email)
}
