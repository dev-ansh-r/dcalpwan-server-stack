package assertions

import "fmt"

func need(needed int, expected []any) string {
	if len(expected) != needed {
		return fmt.Sprintf(needExactValues, needed, len(expected))
	}
	return success
}
