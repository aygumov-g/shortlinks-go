package templates

const FORM_HTML_PAGE = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Form</title>
	<style>
		body { display: flex; justify-content: center; align-items: center; height: 100vh; font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; }
		.form-container { background-color: #fff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); box-sizing: border-box; }
		.form-container h2 {margin-bottom: 20px; font-size: 24px; text-align: center;}
		.form-group { margin-bottom: 15px; }
		.form-group:last-of-type { margin-bottom: 20px; }
		.form-group label { display: block; margin-bottom: 5px; font-weight: bold; }
		.form-group input { width: 100%; padding: 10px; border: 1px solid #ccc; border-radius: 4px; font-size: 16px; box-sizing: border-box; }
		.form-group input[type="submit"] { background-color: #28a475; color: #fff; border: none; cursor: pointer; transition: background-color 0.3s; }
		.form-group input[type="submit"]:hover { background-color: #218838; }
		.form-group input[type="submit"]:active { background-color: #085c1a; }
		.example-text { color: #FF4500; margin-top: 15px; padding-bottom: 10px; }
	</style>
</head>
<body>
	<div class="form-container">
		<h2></h2>
		<form id="form">
			<div class="form-group" id="oneGroup">
				<label for="one" id="oneLabel"></label>
				<input type="text" id="oneInput" name="one">
			</div>
			<div class="example-text" id="exampleText">Пример: https://google.com</div>
			<div class="form-group">
				<input type="submit" id="submit">
			</div>
		</form>
	</div>
	<script>
		let params = new URLSearchParams(window.location.search);
		let param1 = params.get("type");
		if (param1 == "CR") {
			document.querySelector("h2").textContent = "Создание короткой ссылки";
			document.getElementById("oneLabel").innerHTML = "Вставь сюда ссылку, короткую версию которой требуется получить:";
			document.getElementById("submit").value = "Получить";
			document.getElementById("submit").onclick = async function() {
				event.preventDefault();
				if (document.getElementById("oneInput").value != "") {
					(async function() {
						try {
							let response = await fetch("{BACKEND_SERVER_URL}", {
								method: "POST",
								body: JSON.stringify({
									"link_addr_out": document.getElementById("oneInput").value
								}),
								headers: {
									"Content-Type": "application/json",
								}
							});
							text = await response.text();
							if (text.includes("{")) {
								sessionStorage.setItem("text", text)
								window.location.href = "/form?type=INF";
							} else {
								alert("Ошибка: " + text);
								window.location.href = "/";
							}
						} catch {
							alert("Не могу подключиться к серверу.");
							window.location.href = "/";
						}
					}())
				}
			}
		} else if (param1 == "INF") {
			document.querySelector("h2").textContent = "Информация о ссылке";
			const text = sessionStorage.getItem("text")
			if (text != null) {
				var data = JSON.parse(text)
				const domain = window.location.protocol + "//" + window.location.hostname.replace(/xn--s1a/, "у") + "/";
				const text1 = "<pre>Основная: " + data["link_addr_out"] + "</pre>"
				const text2 = "<pre>Короткая версия: " + domain + data["link_addr_in"] + "</pre>"

				document.getElementById("oneLabel").innerHTML = text1 + text2
			} else {
				window.location.href = "/";
			}
			document.getElementById("oneInput").style.display = "none";
			document.getElementById("exampleText").style.display = "none";
			document.getElementById("submit").value = "Окей";
			document.getElementById("submit").onclick = function() {
				event.preventDefault();
				window.location.href = "/";
			}
		} else if (!param1) {
			window.location.href = "/";
		}
	</script>
</body>
</html>
`
