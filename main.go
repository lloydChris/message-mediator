package main

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
}

func main() {

	//Subscribe to queue
	//call function on message read

}

func processMessage(email Email) {

}
