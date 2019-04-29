# awsom session

**awsom session** makes working with AWS SDK session easier.

## Usage

 NewSession returns session which respects:
 - environment variables
 - `~/aws/.config` and `~/aws/.credentials` files

 Example:

     import "github.com/hekonsek/awsom_session"
     ...
     err, sess := awsom_session.NewSession()
     
## License

This project is distributed under Apache 2.0 license.