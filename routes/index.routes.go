package routes

import "net/http"

func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello my love, this is my first program in golang, I love you so much ❤️"))
	if err != nil {
		return
	}
}
