import { WindowSetSize } from '$lib/wailsjs/runtime/runtime';

export function setup() {
	const ro = new ResizeObserver((entries) => {
		for (let entry of entries) {
			const cr = entry.contentRect;
			console.log('resize window', cr.width, cr.height);
			WindowSetSize(cr.width, cr.height);
		}
	});
	const mainElement = document.getElementById('main');
	if (!mainElement) {
		console.log("couldn't find main element");
		return;
	}
	ro.observe(mainElement);
}
