package camStructs

import (
	"github.com/go-cam/cam/base/camStatics"
)

// validation's rule
type Rule struct {
	camStatics.RuleInterface

	fields   []string
	handlers []camStatics.ValidHandler
}

// new rule
func NewRule(fields []string, handlers ...camStatics.ValidHandler) *Rule {
	rule := new(Rule)
	rule.fields = fields
	rule.handlers = handlers
	return rule
}

// get fields
func (rule *Rule) Fields() []string {
	return rule.fields
}

// get handlers
func (rule *Rule) Handlers() []camStatics.ValidHandler {
	return rule.handlers
}
