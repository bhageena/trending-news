package main


/*Main Sends the news articles to the Slack channel

SLACK_HOOK: Entire URL of the Slack hook (Ex: https://hooks.slack.com/services/TTFF7GGAA/DEHHHHKJL/FdfEdfjdkrueiso) */
func Main(args map[string]interface{}) map[string]interface{} {
	msg := make(map[string]interface{})
	msg["body"] = ""
	return msg
}
