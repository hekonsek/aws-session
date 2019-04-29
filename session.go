package awsom_session

import "github.com/aws/aws-sdk-go/aws/session"

// NewSession returns session which respects:
// - environment variables
// - `~/aws/.config` and `~/aws/.credentials` files
//
// Example:
//
//     import "github.com/hekonsek/awsom-session"
//     ...
//     err, sess := awsom_session.NewSession()
func NewSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
