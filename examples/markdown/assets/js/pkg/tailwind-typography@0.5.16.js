/**
 * Minified by jsDelivr using Terser v5.37.0.
 * Original file: /npm/@tailwindcss/typography@0.5.16/src/index.js
 *
 * Do NOT use SRI with dynamically generated files! More information: https://www.jsdelivr.com/using-sri-with-dynamic-files
 */
const plugin = require("tailwindcss/plugin"),
	merge = require("lodash.merge"),
	castArray = require("lodash.castarray"),
	styles = require("./styles"),
	{ commonTrailingPseudos: commonTrailingPseudos } = require("./utils"),
	computed = {};
function inWhere(e, { className: r, modifier: t, prefix: s }) {
	let i = s(`.not-${r}`).slice(1),
		a = e.startsWith(">")
			? ("DEFAULT" === t ? `.${r}` : `.${r}-${t}`) + " "
			: "",
		[o, n] = commonTrailingPseudos(e);
	return o
		? `:where(${a}${n}):not(:where([class~="${i}"],[class~="${i}"] *))${o}`
		: `:where(${a}${e}):not(:where([class~="${i}"],[class~="${i}"] *))`;
}
function isObject(e) {
	return "object" == typeof e && null !== e;
}
function configToCss(
	e = {},
	{ target: r, className: t, modifier: s, prefix: i },
) {
	function a(e, o) {
		if ("legacy" === r) return [e, o];
		if (Array.isArray(o)) return [e, o];
		if (isObject(o)) {
			return Object.values(o).some(isObject)
				? [
						inWhere(e, { className: t, modifier: s, prefix: i }),
						o,
						Object.fromEntries(
							Object.entries(o).map(([e, r]) => a(e, r)),
						),
					]
				: [inWhere(e, { className: t, modifier: s, prefix: i }), o];
		}
		return [e, o];
	}
	return Object.fromEntries(
		Object.entries(
			merge(
				{},
				...Object.keys(e)
					.filter((e) => computed[e])
					.map((r) => computed[r](e[r])),
				...castArray(e.css || {}),
			),
		).map(([e, r]) => a(e, r)),
	);
}
module.exports = plugin.withOptions(
	({ className: e = "prose", target: r = "modern" } = {}) =>
		function ({ addVariant: t, addComponents: s, theme: i, prefix: a }) {
			let o = i("typography"),
				n = { className: e, prefix: a };
			for (let [s, ...i] of [
				["headings", "h1", "h2", "h3", "h4", "h5", "h6", "th"],
				["h1"],
				["h2"],
				["h3"],
				["h4"],
				["h5"],
				["h6"],
				["p"],
				["a"],
				["blockquote"],
				["figure"],
				["figcaption"],
				["strong"],
				["em"],
				["kbd"],
				["code"],
				["pre"],
				["ol"],
				["ul"],
				["li"],
				["table"],
				["thead"],
				["tr"],
				["th"],
				["td"],
				["img"],
				["video"],
				["hr"],
				["lead", '[class~="lead"]'],
			]) {
				i = 0 === i.length ? [s] : i;
				let a = "legacy" === r ? i.map((e) => `& ${e}`) : i.join(", ");
				t(`${e}-${s}`, "legacy" === r ? a : `& :is(${inWhere(a, n)})`);
			}
			s(
				Object.keys(o).map((t) => ({
					["DEFAULT" === t ? `.${e}` : `.${e}-${t}`]: configToCss(
						o[t],
						{ target: r, className: e, modifier: t, prefix: a },
					),
				})),
			);
		},
	() => ({ theme: { typography: styles } }),
);
//# sourceMappingURL=/sm/5719670467f9d2b5ca1f6b7d8792d9eecfc2184cf518b0a1ba0a72886dbfe6da.map
