/**
 * @typedef {import('./app').GoLaunch.GoMethodRequest} GoMethodRequest
 */

export async function handleIframeMessage(event) {
	/** @type {GoMethodRequest} */
	const data = event.data;
	const { method, messageId, args } = data;
	if (method === 'FuzzyFindDesktopEntry') {
		const result = await FuzzyFindDesktopEntry(args);
		event.source.postMessage({ messageId, result }, event.origin);
	}
	if (method === 'LaunchApp') {
		await LaunchApp(args);
		event.source.postMessage({ messageId }, event.origin);
	}
}

window.addEventListener('message', handleIframeMessage);
