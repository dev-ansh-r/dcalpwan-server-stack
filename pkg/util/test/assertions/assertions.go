// Package assertions contains custom assertions compatible with github.com/smarty/assertions.
package assertions

const (
	success           = ""
	needExactValues   = "This assertion requires exactly %d comparison values (you provided %d)."
	needAtLeastValues = "This assertion requires at least %d comparison values (you provided %d)."

	shouldNotHaveResembled = "Expected        '%#v'\nto NOT resemble '%#v'\n(but it did)!"
)
