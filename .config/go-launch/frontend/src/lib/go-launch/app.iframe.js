/**
 * @typedef {import('./app').GoLaunch.DesktopEntry} DesktopEntry
 * @typedef {import('./app').GoLaunch.GoMethodResponse} GoMethodResponse
 * @typedef {import('./app').GoLaunch.GoMethod} GoMethod
 */

const PROD = 'wails://wails';
const DEV = 'http://localhost:34115';

/** @returns {string} uuid*/
function generateUUID() {
	return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
		const r = (Math.random() * 16) | 0;
		const v = c === 'x' ? r : (r & 0x3) | 0x8;
		return v.toString(16);
	});
}
/**
 * @param {GoMethod} method
 * @param {any[]} args
 * @returns {Promise<any>}
 */
function callGoMethod(method, ...args) {
	return new Promise((resolve, _) => {
		const responseId = generateUUID();

		function handleResponse(event) {
			// if (event.origin !== window.location.origin) return; // Security check
			console.log(event);

			/** @type{GoMethodResponse} */
			const data = event.data;
			const { messageId, result } = data;
			if (responseId === messageId) {
				window.removeEventListener('message', handleResponse);
				resolve(result);
			}
		}

		window.addEventListener('message', handleResponse);
		window.parent.postMessage({ method, args, messageId: responseId }, PROD);
	});
}

/**
 * @param {string} prompt
 * @returns {Promise<DesktopEntry[][]>}
 */

export async function FuzzyFindDesktopEntry(prompt) {
	return await callGoMethod('FuzzyFindDesktopEntry', prompt);
}

/**
 * @param {string} appId
 * @returns {Promise<void>}
 */
export async function LaunchApp(appId) {
	return await callGoMethod('LaunchApp', appId);
}
