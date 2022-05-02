package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var debug_disable_sendgrid bool = true

func main() {

	client, _ := reddit.NewReadonlyClient()

	posts, _, _ := client.Subreddit.TopPosts(context.Background(), "hiring", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "all",
	})
	println(posts)
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
	// from := "go.get.jobs.oss@gmail.com"
	// password := os.Getenv("GMAIL_PASSWORD")
	// fmt.Println(password)

	// Receiver email address.
	// to := []string{
	// 	"go.get.jobs.oss+dev@gmail.com",
	// }

	// Message.
	if !debug_disable_sendgrid {
		message_text := "test message with \n" + posts_string

		from := mail.NewEmail(os.Getenv("SENDGRID_FROM_NAME"), os.Getenv("SENDGRID_FROM_EMAIL"))
		subject := "Go Get Jobs notification"
		to := mail.NewEmail(os.Getenv("SENDGRID_TO_NAME"), os.Getenv("SENDGRID_TO_EMAIL"))
		plainTextContent := message_text
		htmlContent := message_text
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		sendgridclient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := sendgridclient.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}
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
