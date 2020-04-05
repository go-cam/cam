package camStructs

import "github.com/go-cam/cam/base/camBase"

// validation's rule
type Rule struct {
	camBase.ValidRuleInterface

	fields   []string
	handlers []camBase.ValidHandler
}

// new rule
func NewRule(fields []string, handlers ...camBase.ValidHandler) *Rule {
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
func (rule *Rule) Handlers() []camBase.ValidHandler {
	return rule.handlers
}
