package handler

import (
	"fmt"
	"net/http"
)

const html = `<!DOCTYPE html>
<html>
	<head>
		<title>Short URL</title>
		<style>
			* {
				font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
			}

			.container {
				display: table;
  				margin: 0 auto;
			}

			.short {
				display: block;
				width: 230px;
			}

			input {
				font-size: 15px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Short URL</h1>
			<form>
				<input id="url" type="url" />
				<input id="shorten" type="submit" value="Shorten" />
				<input class="short" id="short" type="text" value="" readonly />
			</form>
		</div>
		<script>
			(function() {
				function shorten(event) {
					var $url = document.getElementById('url');
					if ($url.value == '') {
						event.preventDefault();
						return;
					}

					var xhr = new XMLHttpRequest();
					xhr.open("POST", window.location.href, true);
					xhr.setRequestHeader("Content-Type", "application/json");
					xhr.onreadystatechange = function () {
						if (xhr.readyState === 4 && xhr.status === 200) {
							var json = JSON.parse(xhr.responseText);

							var $short = document.getElementById('short');
							$short.value = json.short_url;
						}
					};
					var data = JSON.stringify({"long_url": $url.value});
					xhr.send(data);

					event.preventDefault();
				}

				var $shorten = document.getElementById('shorten');
				$shorten.addEventListener('click', shorten);
			})();
		</script>
	</body>
</html>`

// Index index handler func for router
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, html)
}
