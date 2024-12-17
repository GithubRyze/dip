package pkg

import (
	"fmt"
	"regexp"
)

func MatchExactPath(path string, pattern string) bool {
	// 正则表达式确保路径精确匹配
	regex := fmt.Sprintf("^%s$", regexp.QuoteMeta(pattern))
	matched, _ := regexp.MatchString(regex, path)
	return matched

}
