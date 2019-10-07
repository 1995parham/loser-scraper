package mail

const tmplString = `
<html>
        <body>
		<h1>Tweets from {{.Target}}</h1>
		<p>Dear user these are the tweets from your target:</p>
		{{range $t := .Timelines}}
		<h4>{{$t.User}}</h4>
		<p>
		{{$t.Content}}
		</p>
		{{end}}
        </body>
</html>
`
