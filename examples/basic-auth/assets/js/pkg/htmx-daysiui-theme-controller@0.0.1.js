import htmx from "htmx.org";

/**
 * Apply theme to theme controllers
 * @param {string} theme - Theme name to apply
 */
function applyTheme(theme) {
	/** @type {NodeListOf<HTMLInputElement>} */
	const themeControllers = document.querySelectorAll(".theme-controller");
	themeControllers.forEach((controller) => {
		controller.checked = controller.value === theme;
	});
}

/**
 * Get system theme preference
 * @returns {string} "dark" or "light"
 */
function getSystemTheme() {
	const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
	return prefersDark ? "dark" : "light";
}

// handle page load via navigating or boosted links
htmx.defineExtension("daisyui-theme-controller", {
	onEvent: function (name, event) {
		if (
			(name === "htmx:load" || name === "htmx:afterSettle") &&
			event.target === document.body
		) {
			const savedTheme = localStorage.getItem("theme-controller");
			const themeToApply = savedTheme || getSystemTheme();

			applyTheme(themeToApply);
			if (!savedTheme) {
				localStorage.setItem("theme-controller", themeToApply);
			}
			
			document.body.removeAttribute("style");
			document.body.classList.add("opacity-100!");
		}
		return true;
	},
});

// Handle system preference changes
window
	.matchMedia("(prefers-color-scheme: dark)")
	.addEventListener("change", () => {
		const savedTheme = localStorage.getItem("theme-controller");
		if (!savedTheme) {
			// No theme saved, follow system preference
			const systemTheme = getSystemTheme();
			applyTheme(systemTheme);
		}
	});
