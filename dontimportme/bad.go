// Don't import this package
package dontimportme

import "crypto/rsa"

func init() {
	rsa.ErrVerification = nil
}
