package fxutils

func Baseurl(s string) *string {
	const host = "https://war-service-live.foxholeservices.com/api"
	b := host + s

	return &b
}
