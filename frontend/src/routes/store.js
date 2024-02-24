import { writable, derived, get } from 'svelte/store';

/**
 * @typedef {import('svelte/store').Writable} Writable
 * @typedef {import('svelte/store').Derived} Derived
 * @typedef {App.DesktopEntry} DesktopEntry
 * @type {Writable<App.DesktopEntry[]>}
 */
export const desktopEntries = writable([]);

export const selectionPosition = writable({
	row: 0,
	col: 0
});

/** @type {Writable<HTMLInputElement>}*/
export const keyboardNavigationInput = writable();

/** @type {Writable<HTMLInputElement>}*/
export const promptInput = writable();

/** @type {Writable<string>}*/
export const searchTerm = writable('');

/** @type {Writable<App.DesktopEntry[]>}*/
export const searchResults = writable([]);

/** @type {Writable<App.DesktopEntry[]>}*/
export const selectedEntry = derived(selectionPosition, ($selectionPosition) => {
	const $searchResults = get(searchResults);
	if (!$selectionPosition || $searchResults.every((row) => row.length === 0)) return null;
	if ($searchResults[0].length === 1) return $searchResults[0][0];
	return $searchResults[$selectionPosition.row][$selectionPosition.col];
});
