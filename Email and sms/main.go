package main

// // uiqordbpeimpbpqm

// // Sending Email Using Smtp in Golang
import (
	"log"
	"net/smtp"
)

// func main() {
// 	send("Day 2 Testing")
// }

func send(body string) {
	from := "bihnam998@gmail.com"
	pass := "ipzkdahqvdsdfhgx"
	to := "smaeil.mazahery@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: TestDay2\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, Done")
}

// // ------------------------------------------- using template

// // package main

// // import (
// // 	"bytes"
// // 	"fmt"
// // 	"html/template"
// // 	"net/smtp"
// // 	"os"
// // )

// func main() {
// 	// send without template
// 	send("before template test 2")

// 	GMAIL_USERNAME := "bihnam998@gmail.com"
// 	GMAIL_PASSWORD := "ipzkdahqvdsdfhgx"
// 	gmailAuth := smtp.PlainAuth("", GMAIL_USERNAME, GMAIL_PASSWORD, "smtp.gmail.com")

// 	t, err := template.ParseFiles("email-template.html")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var body bytes.Buffer

// 	headers := "MINE-version: 1.0;\nContent-Type: text/html;"
// 	_, err3 := body.Write([]byte(fmt.Sprintf("Subject: html test\n%s\n\n", headers)))

// 	if err3 != nil {
// 		fmt.Println(err3)
// 	}

// 	err4 := t.Execute(&body, struct {
// 		Name    string
// 		Email   string
// 		Message string
// 	}{
// 		Name:    "Essi Kuni",
// 		Email:   "bihnam998@gmail.com",
// 		Message: "Suck it!\n8=============================================================D",
// 	})
// 	if err4 != nil {
// 		fmt.Println(err4)
// 	}

// 	err2 := smtp.SendMail("smtp.gmail.com:587", gmailAuth, GMAIL_USERNAME, []string{"smaeil.mazahery@gmail.com", "bihnam998@gmail.com"}, body.Bytes())

// 	if err2 != nil {
// 		fmt.Println(err2)
// 	} else {
// 		fmt.Println("done")
// 	}
// }

// test send sms

// package main

// import (
// 	"fmt"

// 	"github.com/kavenegar/kavenegar-go"
// )

// func main() {
// 	api := kavenegar.New("663171736B65374D61476E67475957743543326F474F3254777943347469675063307845302F54384736673D")
// 	sender := "10008663"
// 	receptor := []string{"09224534117"}
// 	message := "با سلام\nجناب آقای ابراهیم سلطانی\nلطفا هر چه سریع تر نسبت به تصویه حساب خود با آقای مظاهری اقدام فرمائید."
// 	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
// 		switch err := err.(type) {
// 		case *kavenegar.APIError:
// 			fmt.Println(err.Error())
// 		case *kavenegar.HTTPError:
// 			fmt.Println(err.Error())
// 		default:
// 			fmt.Println(err.Error())
// 		}
// 	} else {
// 		for _, r := range res {
// 			fmt.Println("MessageID 	= ", r.MessageID)
// 			fmt.Println("Status    	= ", r.Status)
// 			//...
// 		}
// 	}
// }
