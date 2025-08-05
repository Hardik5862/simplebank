package mail

import (
	"testing"

	"github.com/Hardik5862/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..", "app")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://hardiksachan.in">Simple bank golang</a>.</p>
	`
	to := []string{"hardik0casr@gmail.com"}

	err = sender.SendMail(subject, content, to, nil, nil)
	require.NoError(t, err)
}
