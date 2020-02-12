package main

//Simulates an external service this could be a different Go package, or an API
func GetPipeline(email Email) Pipeline {
	if email.From == "me@wildbit.com" {
		return PipeA
	}
	return PipeB
}
