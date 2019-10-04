package mail

const tmplString = `
<html>
        <body>
		<h1>You are a Loser!</h1>
		<p>Dear user these are the tweets from your target:</p>
		{{range $t := .Timelines}}
		<p>
		{{$t.Content}}
		</p>
		{{end}}
        </body>
</html>
`
