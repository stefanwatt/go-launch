import { WindowSetSize } from '$lib/wailsjs/runtime/runtime';

export function setup() {
	const ro = new ResizeObserver((entries) => {
		for (let entry of entries) {
			const cr = entry.contentRect;
			if (!cr?.width || !cr?.height || !cr.width > 0 || !cr.height > 0) return;
			WindowSetSize(Math.round(cr.width), Math.round(cr.height));
		}
	});
	const mainElement = document.getElementById('main');
	if (!mainElement) {
		console.log("couldn't find main element");
		return;
	}
	ro.observe(mainElement);
}
