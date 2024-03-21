// import { WindowSetSize } from '$lib/wailsjs/runtime/runtime';

//HACK: fix incorrect window height
const plus = 4;

/** @param {HTMLElement} node */
export function setupResizeObserver(node) {
	// const { clientWidth, clientHeight } = node;
	// WindowSetSize(Math.ceil(clientWidth), Math.ceil(clientHeight)) + plus;
	// const ro = new ResizeObserver((entries) => {
	// 	const cr = entries[0].contentRect;
	// 	if (!cr?.width || !cr?.height || !cr.width > 0 || !cr.height > 0) return;
	// 	WindowSetSize(Math.ceil(cr.width), Math.ceil(cr.height) + plus);
	// });
	// ro.observe(node);
}
