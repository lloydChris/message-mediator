package main

func GetPipeline(email Email) Pipeline {
	if email.From == "me@wildbit.com" {
		return PipeA
	}
	return PipeB
}
