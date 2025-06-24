import htmx from "htmx.org";

htmx.defineExtension("daisyui-theme-controller", {
	onEvent: function (name, _evt) {
		if (name === "htmx:load") {
			const savedTheme = localStorage.getItem("theme-controller");
			if (savedTheme) {
				/** @type {NodeListOf<HTMLInputElement>} */
				const themeControllers =
					document.querySelectorAll(".theme-controller");
				themeControllers.forEach((controller) => {
					if (controller.value === savedTheme) {
						controller.checked = true;
					}
				});
			}
			document.body.classList.add("opacity-100!");
		}
		return true;
	},
});
