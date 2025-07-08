import htmx from "htmx.org";

htmx.defineExtension("form-reset", {
	onEvent: function (name, evt) {
		if (name === "htmx:afterRequest" && evt.target.tagName === "FORM") {
			const status = evt.detail.xhr.status;
			if (status >= 200 && status < 300) {
				evt.target.reset();
			}
		}
		return true;
	},
});
