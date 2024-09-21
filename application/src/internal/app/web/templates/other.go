package templates

const OTHER_HTML_PAGE = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Perelink</title>
</head>
<body>
    <script>
        (async function() {
            try {
                let response = await fetch("{BACKEND_SERVER_URL}" + window.location.pathname, {
                    method: "GET"
                });
                text = await response.text();
                if (text.includes("{")) {
                    window.location.href = JSON.parse(text)["link_addr_out"];
                } else {
                    alert("Ошибка: " + text);
                    window.location.href = "/";
                }
            } catch {
                alert("Не могу подключиться к серверу.");
                window.location.href = "/";
            }
        }())
    </script>
</body>
</html>
`
