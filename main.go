package main

import (
	"context"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func main() {

	client, _ := reddit.NewReadonlyClient()

	posts, _, _ := client.Subreddit.TopPosts(context.Background(), "hiring", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "all",
	})

	//= [post for post in posts if post.title.lower().contains('[hiring]')]
	var filteredPosts []*reddit.Post

	for _, post := range posts {
		if strings.Contains(strings.ToLower(post.Title), "[hiring]") {
			filteredPosts = append(filteredPosts, post)
		}

	}

	posts_string := ""
	fmt.Printf("Received %d posts.\n", len(posts))
	for i, s := range filteredPosts {
		fmt.Println(i, s)
		posts_string = posts_string + strconv.Itoa(i) + s.Title + " " + s.URL + " \n"
	}

	// Sender data.
	from := "go.get.jobs.oss@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	fmt.Println(password)

	// Receiver email address.
	to := []string{
		"go.get.jobs.oss+dev@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("test message with \n" + posts_string)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
