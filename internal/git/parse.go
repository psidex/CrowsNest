package git

import "regexp"

// TODO: More &| better regex, maybe .git dir will provide better insights?

var pulledRegex *regexp.Regexp = regexp.MustCompile(
	`Updating [A-Za-z0-9_]{7}..[A-Za-z0-9_]{7}\nFast-forward`,
)

// PullPulled determines if the output from running `git pull` shows changes made.
func PullPulled(str string) bool {
	return pulledRegex.Match([]byte(str))
}
