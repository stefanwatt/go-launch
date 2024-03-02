import { WindowSetSize } from '$lib/wailsjs/runtime/runtime';

/** @param {HTMLElement} node */
export function setupResizeObserver(node) {
	const { clientWidth, clientHeight } = node;
	WindowSetSize(Math.ceil(clientWidth), Math.ceil(clientHeight)) + 15;
	const ro = new ResizeObserver((entries) => {
		for (let entry of entries) {
			const cr = entry.contentRect;
			if (!cr?.width || !cr?.height || !cr.width > 0 || !cr.height > 0) return;
			WindowSetSize(Math.ceil(cr.width), Math.ceil(cr.height)) + 15;
		}
	});
	ro.observe(node);
}
