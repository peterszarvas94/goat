import htmx from "htmx.org";

htmx.defineExtension("auto-remove", {
	/**
	 * @param {string} name - Event name
	 * @param {Event} evt - Event object
	 */
	onEvent: function (name, evt) {
		if (name === "htmx:afterProcessNode") {
			var target = /** @type {Element} */ (evt.target);

			// Find all elements with hx-auto-remove attribute
			var elements = target.querySelectorAll("[hx-auto-remove]");

			var elementsArray = Array.from(elements);
			if (target.hasAttribute && target.hasAttribute("hx-auto-remove")) {
				elementsArray = [target].concat(elementsArray);
			}

			elementsArray.forEach(function (element) {
				var delay =
					parseInt(element.getAttribute("hx-auto-remove")) || 3000;

				setTimeout(function () {
					if (element && element.parentNode) {
						var htmlElement = /** @type {HTMLElement} */ (element);
						htmlElement.style.transition = "opacity 0.5s ease-out";
						htmlElement.style.opacity = "0";
						setTimeout(function () {
							if (element && element.parentNode) {
								element.remove();
							}
						}, 500);
					}
				}, delay);
			});
		}

		return true;
	},
});
