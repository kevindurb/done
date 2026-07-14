POST {{base_url}}/signup
[Form]
Email: test@example.com
Password: abc123

POST {{base_url}}/login
[Form]
Email: test@example.com
Password: abc123
HTTP 302
[Captures]
session: cookie "done-session"

GET {{base_url}}/
[Cookies]
session: {{session}}
HTTP 200
[Asserts]
xpath "string(/html/head/title)" contains "Done"
xpath "string(/html/body/main/h1)" contains "Home"
