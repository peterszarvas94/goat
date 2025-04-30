import htmx from "htmx.org";

/* THEME SELECTOR */
function getSystemTheme() {
	return window.matchMedia("(prefers-color-scheme: dark)").matches
		? "dark"
		: "light";
}

/** @param {string} value */
function setAttribute(value) {
	const realValue = value === "system" ? getSystemTheme() : value;
	document.body.setAttribute("data-theme", realValue);
}

/** @param {string} value */
function setLocalTheme(value) {
	return localStorage.setItem("theme", value);
}

/** @returns {string | null} */
function getLocalTheme() {
	return localStorage.getItem("theme");
}

function initializeThemeSelectors() {
	/** @type {NodeListOf<HTMLSelectElement>} */
	const themeSelectors = document.querySelectorAll(
		"select.goat-theme-selector",
	);

	themeSelectors.forEach((select) => {
		// Set the value to match saved theme
		select.value = getLocalTheme();

		// Add change listener if it doesn't already have one
		if (!select.dataset.goatInitialized) {
			select.addEventListener("change", (event) => {
				const newValue = /** @type {HTMLSelectElement} */ (event.target)
					.value;
				setAttribute(newValue);
				setLocalTheme(newValue);
			});
			select.dataset.goatInitialized = "true";
		}
	});
}

/* PASSWORD REVEAL */
function initializePasswordReveals() {
	const passwordReveals = Array.from(
		document.querySelectorAll(".goat-pw-reveal"),
	);
	passwordReveals.forEach(
		/** @param {HTMLButtonElement} button */ (button) => {
			const id = button.dataset.goatPw;
			// Remove the # from the ID if it's included in the data attribute
			const cleanId = id.startsWith("#") ? id.substring(1) : id;

			const inputElement = /** @type {HTMLInputElement | null} */ (
				document.getElementById(cleanId)
			);
			if (!inputElement) {
				console.error(
					`There is no input element with the id ${cleanId}, which is defined at:`,
					button,
				);
				return;
			}

			const pwVisibleElement = /** @type {HTMLElement | null} */ (
				button.querySelector(".goat-pw-visible")
			);

			const pwHiddenElement = /** @type {HTMLElement | null} */ (
				button.querySelector(".goat-pw-hidden")
			);

			if (!button.dataset.goatInitialized) {
				button.addEventListener("click", () => {
					if (inputElement.type === "password") {
						inputElement.type = "text";
						if (pwVisibleElement) {
							pwVisibleElement.style.display = "none";
						}

						if (pwHiddenElement) {
							pwHiddenElement.style.display = "block";
						}
					} else if (inputElement.type === "text") {
						inputElement.type = "password";
						if (pwVisibleElement) {
							pwVisibleElement.style.display = "block";
						}

						if (pwHiddenElement) {
							pwHiddenElement.style.display = "none";
						}
					}
				});
			}
			button.dataset.goatInitialized = "true";
		},
	);
}

// handle page load via navigating or boosted links
htmx.defineExtension("init-goatui", {
	onEvent: (name, event) => {
		if (
			(name === "htmx:load" || name === "htmx:afterSettle") &&
			event.target === document.body
		) {
			const localTheme = getLocalTheme();
			if (!localTheme) {
				setLocalTheme("system");
			}
			setAttribute(localTheme || "system");
			initializeThemeSelectors();
			initializePasswordReveals();
		}
		return true;
	},
});

// For system settings change
window
	.matchMedia("(prefers-color-scheme: dark)")
	.addEventListener("change", () => {
		const localTheme = getLocalTheme();
		if (localTheme === "system") {
			setAttribute(localTheme);
		}
	});
